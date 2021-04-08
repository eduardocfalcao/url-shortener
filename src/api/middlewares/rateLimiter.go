package middlewares

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
)

type RateLimiterMiddleware struct {
	accessPerSecond int
	availableTokens int
	m               map[string]*rate.Limiter
	mutex           sync.Mutex
}

func NewApiRateLimiter(accessPerSecond, availableTokens int) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{
		accessPerSecond: accessPerSecond,
		availableTokens: availableTokens,
		m:               make(map[string]*rate.Limiter),
	}
}

func (m *RateLimiterMiddleware) getLimiter(r *http.Request) *rate.Limiter {
	vars := mux.Vars(r)
	key := vars["shorturl"]

	m.mutex.Lock()
	defer m.mutex.Unlock()

	if limiter, ok := m.m[key]; ok {
		return limiter
	} else {
		l := rate.NewLimiter(rate.Limit(m.accessPerSecond), m.availableTokens)
		m.m[key] = l
		return l
	}
}

func (m *RateLimiterMiddleware) Limit(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := m.getLimiter(r)
		if limiter.Allow() {
			next(w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
		}
	})
}
