package middlewares

import (
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	mu        sync.Mutex
	lastVisit map[string]time.Time
	interval  time.Duration
}

func NewRateLimiter(interval time.Duration) *RateLimiter {
	return &RateLimiter{
		lastVisit: make(map[string]time.Time),
		interval:  interval,
	}
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		rl.mu.Lock()
		last, exists := rl.lastVisit[ip]
		now := time.Now()
		if exists && now.Sub(last) < rl.interval {
			rl.mu.Unlock()
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		rl.lastVisit[ip] = now
		rl.mu.Unlock()
		next.ServeHTTP(w, r)
	})
}