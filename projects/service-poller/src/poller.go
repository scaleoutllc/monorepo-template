package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	ops = promauto.NewCounter(prometheus.CounterOpts{
		Name: "processed_ops_total",
		Help: "The total number of processed events",
	})
)

func poll(client *http.Client, qps int, target string, method string, body string) {
	request, err := http.NewRequest(method, target, strings.NewReader(body))
	if err != nil {
		log.Fatalln(err)
	}
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			for i := 1; i <= qps; i++ {
				go client.Do(request)
				ops.Inc()
			}
		}
	}()
}
