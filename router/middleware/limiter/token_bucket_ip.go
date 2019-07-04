package limiter

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
	"time"
)

func TBIP() gin.HandlerFunc {
	options := limiter.ExpirableOptions{
		DefaultExpirationTTL: time.Hour,
	}
	lmt := tollbooth.NewLimiter(1, &options)
	lmt.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"})

	// Set a custom expiration TTL for token bucket.
	lmt.SetTokenBucketExpirationTTL(time.Hour)
	// Set a custom expiration TTL for basic auth users.
	lmt.SetBasicAuthExpirationTTL(time.Hour)
	// Set a custom expiration TTL for header entries.
	lmt.SetHeaderEntryExpirationTTL(time.Hour)
	middleware := tollbooth_gin.LimitHandler(lmt)
	return middleware
}
