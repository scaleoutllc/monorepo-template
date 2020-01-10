package endpoints

import (
	"net/http"
)

func NewReadinessHandler(ready *bool) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		status := http.StatusServiceUnavailable
		if *ready {
			status = http.StatusOK
		}
		//log.Printf("Readiness probe received - Health is %d...", status)
		res.WriteHeader(status)
		res.Write([]byte(http.StatusText(status)))
	}
}
