package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	ops = promauto.NewCounter(prometheus.CounterOpts{
		Name: "processed_ops_total",
		Help: "The total number of processed events",
	})
)

// Handler responds to requests with the content of the request.
func EchoHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Range, Content-Disposition, Content-Type, ETag")
	req.Write(res)
	ops.Inc()
}
