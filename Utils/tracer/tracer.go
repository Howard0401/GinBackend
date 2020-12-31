package tracer

import (
	"VueGin/global"
	"io"
	"time"

	"github.com/opentracing/opentracing-go"              //需要手動加
	"github.com/uber/jaeger-client-go/config"            //這個可能要手動加
	jaegerZap "github.com/uber/jaeger-client-go/log/zap" //需要手動加
	"go.uber.org/zap"
)

func NewJaegerTracer(name, port string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: name,
		Sampler: &config.SamplerConfig{ //
			Type:  "const",
			Param: 1, // Param is a value passed to the sampler
		},
		Reporter: &config.ReporterConfig{
			// LogSpans, when true, enables LoggingReporter that runs in parallel with the main reporter and logs all submitted spans.
			LogSpans: true,
			// BufferFlushInterval controls how often the buffer is force-flushed, even if it's not full.
			BufferFlushInterval: 1 * time.Second,
			// LocalAgentHostPort instructs reporter to send spans to jaeger-agent at this address.
			LocalAgentHostPort: port,
		},
	}
	tracer, closer, err := cfg.NewTracer(
		config.Logger(jaegerZap.NewLogger(global.Global_Logger.With(zap.String("type", "tracing")))),
	)
	if err != nil {
		return nil, nil, err
	}
	// opentracing.SetGlobalTracer(tracer) //全域已建立，不需再加入
	return tracer, closer, nil

}
