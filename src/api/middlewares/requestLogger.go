package middlewares

import (
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
}

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Fatalf("Error: %v. Stack Trace: %s", err, debug.Stack())
			}
		}()

		start := time.Now()
		wrapped := wrapResponseWriter(w)
		next(wrapped, r)
		t := time.Since(start)
		duration := t.Milliseconds()
		n := t.Nanoseconds()
		log.Printf("\nStatus: %d. Method: %s. Path: %s. Duration: %d Miliseconds (%d nanoseconds)", wrapped.status, r.Method, r.URL.EscapedPath(), duration, n)
	}
}
