package middlewares

import (
	"net/http"
)

func EnsureHttpMethod(method string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		handler(w, r)
	}
}

func EnsurePost(handler http.HandlerFunc) http.HandlerFunc {
	return EnsureHttpMethod("POST", handler)
}
