package main

import (
	"log"
	"net/http"
	"org/shared/components/go/loglevel"
	"os"
	"strconv"
	"time"

	"org/shared/components/go/endpoints"
	"org/shared/components/go/version"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/tylerb/graceful.v1"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.Println("poller service starting...")

	// get required env variables to run
	httpPort := os.Getenv("APP_SERVER_PORT")
	if httpPort == "" {
		log.Fatal("env APP_SERVER_PORT must be specified (e.g. 8080).")
	}
	qps, err := strconv.Atoi(os.Getenv("APP_POLLER_QPS"))
	if err != nil || qps < 1 {
		log.Fatalf("env APP_POLLER_QPS must be specified | %s", err)
	}
	target := os.Getenv("APP_POLLER_TARGET")
	if target == "" {
		log.Fatal("env APP_POLLER_TARGET must be specified (e.g. http://domain.told/endpoint).")
	}
	method := os.Getenv("APP_POLLER_METHOD")
	if method == "" {
		log.Fatal("env APP_POLLER_METHOD must be specified (e.g. GET).")
	}
	body := os.Getenv("env APP_POLLER_BODY")

	// initialize http server
	e := echo.New()
	e.Logger.SetLevel(loglevel.GetLogLevel())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// mount standard handlers
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	e.GET("/version", echo.WrapHandler(endpoints.NewVersionHandler(version.Identifier())))

	// because this service does not receive inbound traffic we can always
	// consider it "ready"
	ready := true

	// a kubernetes pod configured with a readiness probe on this endpoint will
	// not allow traffic on any ports unless this is returning a 200 status code
	e.GET("/readiness", echo.WrapHandler(endpoints.NewReadinessHandler(&ready)))

	// a kubernetes pod configured with a liveness probe will signal the scheduler
	// to restart this service after a specified number of failures.
	e.GET("/liveness", echo.WrapHandler(http.HandlerFunc(endpoints.LivenessHandler)))

	client := http.Client{
		Timeout: 1 * time.Second,
	}
	poll(&client, qps, target, method, body)

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
	srv.ListenAndServe()
}
