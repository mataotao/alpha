package limiter

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"

	"fmt"
	"net/http"
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

	// Set a custom message.
	lmt.SetMessage("You have reached maximum request limit.")

	// Set a custom content-type.
	lmt.SetMessageContentType("application/json; charset=utf-8")

	// Set a custom function for rejection.
	lmt.SetOnLimitReached(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("A request was rejected")
	})
	middleware := tollbooth_gin.LimitHandler(lmt)
	return middleware
}
