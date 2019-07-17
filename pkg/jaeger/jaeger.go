package jaeger

import (
	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
	"net/http"
	"strings"
	"time"
)

type RequestTrace struct {
	Span          opentracing.Span
	OperationName string
	Request       *http.Request
	Closer        io.Closer
	ClientIP      string
}

func (t *RequestTrace) Init() error {
	LocalAgentHostPortConf := viper.GetString("jaeger.agent_host")
	cfg := jaegercfg.Configuration{
		ServiceName: "service-alpha", //自定义服务名称
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  LocalAgentHostPortConf, //jaeger agent
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return err
	}
	t.Closer = closer

	ancestorSpanContext, err := tracer.Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(t.Request.Header),
	)
	if err != nil {
		return err
	}
	if t.OperationName == "" {
		routeSlice := strings.Split(t.Request.RequestURI, "?")
		t.OperationName = routeSlice[0]
	}
	span := tracer.StartSpan(t.OperationName, opentracing.ChildOf(ancestorSpanContext))
	t.Span = span
	t.Span.SetTag("remote-addr", t.Request.RemoteAddr)
	t.Span.SetTag("ip", t.ClientIP)
	t.Span.SetTag("referer", t.Request.Referer())
	t.Span.SetTag("user_agent", t.Request.UserAgent())
	t.Span.SetTag("method", t.Request.Method)

	return nil
}

func (t *RequestTrace) SetResponseStatus(c gin.Context) {
	t.Span.SetTag("status", c.Writer.Status())
}
