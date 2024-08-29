package ratelimiting

import (
	"net"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimit() gin.HandlerFunc {
	limiter := rate.NewLimiter(2, 5)

	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		c.Next()
	}
}

func RateLimitWithIP() gin.HandlerFunc {
	var (
		mu      sync.Mutex
		clients = make(map[string]*rate.Limiter)
	)

	return func(c *gin.Context) {
		ip, _, err := net.SplitHostPort(c.Request.RemoteAddr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		mu.Lock()
		if _, found := clients[ip]; !found {
			clients[ip] = rate.NewLimiter(2, 4)
		}
		mu.Unlock()

		if !clients[ip].Allow() {
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}

		c.Next()
	}
}
