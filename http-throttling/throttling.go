package throttling

import (
	"net/http"

	"golang.org/x/time/rate"
)

const (
	DefaultRate      = 100 // Default 'rate' to use if application doesn't configure it
	DefaultBurstSize = 10  // Default 'burst' to use if application doesn't configure it
)

var limiter *rate.Limiter

// Application must configure 'rate' and 'burstSize'
func ConfigureLimiter(r rate.Limit, b int) {
	limiter = rate.NewLimiter(r, b)
}

// To be called by application to check rate limit
func LimitRate(next http.Handler) http.Handler {
	if limiter == nil {
		// Create a limiter with defaults as application has not configured it
		limiter = rate.NewLimiter(DefaultRate, DefaultBurstSize)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// return with 'Too Many Requests' when rate is exceeded.
		if !limiter.Allow() {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		// Rate limit allowed handle request as usual
		next.ServeHTTP(w, r)
	})
}
