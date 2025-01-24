package middleware

import (
	"net/http"
	"time"

	"forum/server/internal/data"
	"forum/server/internal/types"
	"forum/server/internal/utils"
)

type MiddleWareLayer struct {
	MiddlewareData data.DataLayer
}

func (db *MiddleWareLayer) Allow(ip string) bool {
	now := time.Now()
	refill, lastRefill := db.MiddlewareData.ExtractBucketDate(ip)
	tokensToAdd := int(now.Sub(lastRefill) / (time.Second * time.Duration(refill)))
	if tokensToAdd > 0 {
		db.MiddlewareData.RefillTokens(tokensToAdd, ip)
	}
	return db.MiddlewareData.TakeToken(ip)
}

func (db *MiddleWareLayer) RateLimiter(next http.Handler, maxTokens int, refillTime time.Duration) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := db.MiddlewareData.GiveBucket(r.RemoteAddr, maxTokens, refillTime, r.URL.Path)
		if !db.Allow(r.RemoteAddr) || err != nil {
			utils.SendResponseStatus(w, http.StatusTooManyRequests, types.ErrTooManyRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
