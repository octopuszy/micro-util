package util

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	config2 "github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)


// NewTrancer 创建链路追踪实例
func NewTrancer(serviceName, addr string)  (opentracing.Tracer, io.Closer, error){
	config := &config2.Configuration{
		ServiceName: serviceName,
		Sampler: &config2.SamplerConfig{
			Type: jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config2.ReporterConfig{
			QueueSize:                  0,
			BufferFlushInterval:        1*time.Second,
			LogSpans:                   true,
			LocalAgentHostPort:         addr,
		},
	}
	return config.NewTracer()
}
