package middlewares

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type AccessCountMiddleware struct {
	metric *prometheus.CounterVec
}

func NewAccessCountMiddleware() *AccessCountMiddleware {
	return &AccessCountMiddleware{
		metric: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "urlshortener_access_url",
			Help: "The total number of access in the shorturl",
		}, []string{"shorturl"}),
	}
}

func (m *AccessCountMiddleware) Collect(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		go m.collect(r)
		next(w, r)
	}
}

func (m *AccessCountMiddleware) collect(r *http.Request) {
	vars := mux.Vars(r)
	shorturl := vars["shorturl"]

	m.metric.WithLabelValues(shorturl).Inc()
}
