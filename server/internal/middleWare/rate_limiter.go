package middleware

import (
	"errors"
	"net/http"
	"time"

	"forum/server/internal/data"
	"forum/server/internal/utils"
)

type MiddleWareLayer struct {
	MiddlewareData data.DataLayer
}

func (db *MiddleWareLayer) Allow(ip string) bool {
	now := time.Now()
	refill, lastRefill := db.MiddlewareData.ExtractBucketDate(ip)
	tokensToAdd := int(now.Sub(lastRefill) / refill)
	if tokensToAdd > 0 {
		db.MiddlewareData.RefillTokens(tokensToAdd, ip)
	}
	return db.MiddlewareData.TakeToken(ip)
}

func (db *MiddleWareLayer) RateLimiter(next http.Handler, maxTokens int, refillTime time.Duration) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db.MiddlewareData.GiveBucket(r.RemoteAddr, maxTokens, refillTime, r.URL.Path)
		if !db.Allow(r.RemoteAddr) {
			utils.SendResponseStatus(w, http.StatusTooManyRequests, errors.New("too many requests"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
