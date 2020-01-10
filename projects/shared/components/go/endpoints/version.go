package endpoints

import "net/http"

// NewVersionHandler builds a http.HandlerFunc that response with the supplied version.
func NewVersionHandler(version string) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		status := http.StatusOK
		res.WriteHeader(status)
		res.Write([]byte(version))
		return
	}
}
