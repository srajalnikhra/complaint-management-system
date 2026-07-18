package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/srajalnikhra/complaint-management-system/internal/utils"
	"golang.org/x/time/rate"
)

// Visitor tracks rate-limiting statistics for a single client IP.
type Visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	visitors = make(map[string]*Visitor)
	mu       sync.Mutex
)

func getLimiter(ip string) *rate.Limiter {

	// Acquire lock to inspect the visitor map.
	mu.Lock()
	defer mu.Unlock()

	// Check if the limiter for this IP already exists.
	visitor, exists := visitors[ip]

	if !exists {

		// Create a new limiter permitting up to 3 requests every 10 seconds.
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

// StartRateLimiterCleanup runs a background loop to clean up stale visitors
// from the in-memory map to prevent memory leaks.
func StartRateLimiterCleanup() {

	// Tick once every minute.
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for range ticker.C {

		// Acquire lock to clean up the map.
		mu.Lock()

		// Delete IPs that haven't made a request in the last 3 minutes.
		for ip, visitor := range visitors {

			if time.Since(visitor.lastSeen) > 3*time.Minute {
				delete(visitors, ip)
			}
		}

		mu.Unlock()
	}
}

// RateLimitMiddleware restricts request rates from clients based on IP.
func RateLimitMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Extract the host IP, removing port details.
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ip = r.RemoteAddr
		}

		// Retrieve the rate limiter corresponding to this IP.
		limiter := getLimiter(ip)

		// Reject the request if the visitor has exceeded their limit.
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
