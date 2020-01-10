package endpoints

import (
	"net/http"
)

// LivenessHandler is a http.HandlerFunc that simply returns status OK.
func LivenessHandler(res http.ResponseWriter, req *http.Request) {
	status := http.StatusOK
	//log.Printf("Liveness probe received - Health is %d...", status)
	res.WriteHeader(status)
	res.Write([]byte(http.StatusText(status)))
	return
}
