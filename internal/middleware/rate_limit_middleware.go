package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/srajalnikhra/complaint-management-system/internal/utils"
	"golang.org/x/time/rate"
)

type Visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	visitors = make(map[string]*Visitor)
	mu       sync.Mutex
)

func getLimiter(ip string) *rate.Limiter {

	mu.Lock()
	defer mu.Unlock()

	visitor, exists := visitors[ip]

	if !exists {

		limiter := rate.NewLimiter(rate.Every(10*time.Second), 3)

		visitors[ip] = &Visitor{
			limiter:  limiter,
			lastSeen: time.Now(),
		}

		return limiter
	}

	visitor.lastSeen = time.Now()

	return visitor.limiter
}

func StartRateLimiterCleanup() {

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for range ticker.C {

		mu.Lock()

		for ip, visitor := range visitors {

			if time.Since(visitor.lastSeen) > 3*time.Minute {
				delete(visitors, ip)
			}
		}

		mu.Unlock()
	}
}

func RateLimitMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ip = r.RemoteAddr
		}

		limiter := getLimiter(ip)

		if !limiter.Allow() {
			utils.Error(
				w,
				http.StatusTooManyRequests,
				"Too many requests. Please try again later.",
			)
			return
		}

		next.ServeHTTP(w, r)
	})
}
