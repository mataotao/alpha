package middleware

import (
	log "alpha/config"
	"alpha/handler"
	"alpha/pkg/errno"
	"alpha/pkg/jaeger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Jaeger() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestTrace := &jaeger.RequestTrace{
			Request:  c.Request,
			ClientIP: c.ClientIP(),
		}
		err := requestTrace.Init()
		if err != nil {
			log.Logger.Error("jaeger InIt",
				zap.Error(err),
			)
			handler.SendBandResponse(c, errno.ErrJaegerInit, nil)
			c.Abort()
			return
		}
		defer requestTrace.Closer.Close()
		defer requestTrace.Span.Finish()
		defer requestTrace.SetResponseStatus(*c)
		c.Set("request_trace", requestTrace)

		c.Next()
	}
}
