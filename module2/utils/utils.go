package utils

import "net/http"

func WriteReqHeadersToResponse(w http.ResponseWriter, r *http.Request) {
	for key, values := range r.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
}
