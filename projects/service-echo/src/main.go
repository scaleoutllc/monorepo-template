package main

import (
	"log"
	"math/rand"
	"net/http"
	"org/shared/components/go/loglevel"
	"os"
	"time"

	"org/shared/components/go/endpoints"
	"org/shared/components/go/version"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/tylerb/graceful.v1"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.Println("echo service starting...")

	// get required env variables to run
	httpPort := os.Getenv("APP_SERVER_PORT")
	if httpPort == "" {
		log.Fatal("env APP_SERVER_PORT must be specified (e.g. 8080).")
	}

	// for the purposes of demonstration leave this value as false for 5 + random
	// 5 seconds during startup. this value is used by a readiness probe below to
	// simulate startup time (when the service is running but not ready to serve
	// traffic). a more realistic use-case would be flagging this to true after a
	// cache has been warmed.
	ready := false
	go func() {
		time.Sleep((5 + time.Duration(rand.Intn(5))) * time.Second)
		ready = true
		log.Printf("echo service ready...")
	}()

	// initialize http server
	e := echo.New()
	e.Logger.SetLevel(loglevel.GetLogLevel())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// mount standard handlers
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	e.GET("/version", echo.WrapHandler(endpoints.NewVersionHandler(version.Identifier())))

	// a kubernetes pod configured with a readiness probe on this endpoint will
	// not allow traffic on any ports unless this is returning a 200 status code
	e.GET("/readiness", echo.WrapHandler(endpoints.NewReadinessHandler(&ready)))

	// a kubernetes pod configured with a liveness probe will signal the scheduler
	// to restart this service after a specified number of failures.
	e.GET("/liveness", echo.WrapHandler(http.HandlerFunc(endpoints.LivenessHandler)))

	// mount service specific handlers
	e.Any("/v1/echo", echo.WrapHandler(http.HandlerFunc(EchoHandler)))

	// create http server that will drain connections gracefully when shut down
	srv := &graceful.Server{
		Timeout: 10 * time.Second,
		Server: &http.Server{
			Handler:           e.Server.Handler,
			ReadTimeout:       1 * time.Minute,
			ReadHeaderTimeout: 1 * time.Minute,
			WriteTimeout:      10 * time.Second,
			Addr:              ":" + httpPort,
		},
	}
	e.Logger.Fatal(srv.ListenAndServe())
}
