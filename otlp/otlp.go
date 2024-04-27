package otlp

import (
	"context"
	"database/sql"

	. "github.com/ZenLiuCN/ote/common"
	otlp "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/trace"
	"log/slog"
	"time"
)

type TraceConfig struct {
	Endpoint           string
	Compress           string
	Insecure           sql.NullBool
	Reconnect          time.Duration
	Timeout            time.Duration
	Retry              *otlp.RetryConfig
	Headers            map[string]string
	ExportTimeout      time.Duration
	ExportBatchSize    sql.NullInt32
	ExportBatchTimeout time.Duration
	QueueSize          sql.NullInt32
	QueueBlocking      sql.NullBool
	Sampler            *SamplerConfig
	*ResourceConfig
}
type SamplerConfig struct {
	Name    string
	Based   string
	Ratio   sql.NullFloat64
	Options []string
}

func NewTraceProvider(ctx context.Context, cfg *TraceConfig) (*trace.TracerProvider, error) {
	var opt []otlp.Option
	{
		opt = append(opt, otlp.WithEndpointURL(cfg.Endpoint))
		if d := cfg.Compress; d != "" {
			opt = append(opt, otlp.WithCompressor(d))
		}
		if cfg.Insecure.Valid && cfg.Insecure.Bool {
			opt = append(opt, otlp.WithInsecure())
		}
		if d := cfg.Reconnect; d != time.Duration(0) {
			opt = append(opt, otlp.WithReconnectionPeriod(d))
		}
		if d := cfg.Timeout; d != time.Duration(0) {
			opt = append(opt, otlp.WithTimeout(d))
		}
		if d := cfg.Retry; d != nil {
			opt = append(opt, otlp.WithRetry(*d))
		}
		if d := cfg.Headers; len(d) > 0 {
			opt = append(opt, otlp.WithHeaders(d))
		}
	}
	traceExporter, err := otlp.New(ctx, opt...)
	if err != nil {
		return nil, err
	}
	var opts []trace.TracerProviderOption
	{
		//!! batch
		{
			var spanOpt []trace.BatchSpanProcessorOption
			if cfg.ExportTimeout != 0 {
				spanOpt = append(spanOpt, trace.WithExportTimeout(cfg.ExportTimeout))
			}
			if cfg.ExportBatchSize.Valid {
				spanOpt = append(spanOpt, trace.WithMaxExportBatchSize(int(cfg.ExportBatchSize.Int32)))
			}

			if cfg.ExportBatchTimeout != 0 {
				spanOpt = append(spanOpt, trace.WithBatchTimeout(cfg.ExportBatchTimeout))
			}
			if cfg.QueueSize.Valid {
				spanOpt = append(spanOpt, trace.WithMaxQueueSize(int(cfg.QueueSize.Int32)))
			}
			if cfg.QueueBlocking.Valid && cfg.QueueBlocking.Bool {
				spanOpt = append(spanOpt, trace.WithBlocking())
			}
			opts = append(opts, trace.WithBatcher(traceExporter, spanOpt...))
		}
		//!! resource
		{
			res, err := ParseResource(ctx, cfg.ResourceConfig)
			if err != nil {
				return nil, err
			}
			opts = append(opts, trace.WithResource(res))
		}
		//!! sampler
		{
			if sampler := cfg.Sampler; sampler != nil {
				switch sampler.Name {
				case "always":
					opts = append(opts, trace.WithSampler(trace.AlwaysSample()))
				case "never":
					opts = append(opts, trace.WithSampler(trace.NeverSample()))
				case "parent":
					var options []trace.ParentBasedSamplerOption
					var sam trace.Sampler
					switch sampler.Based {
					case "always":
						sam = trace.AlwaysSample()
					case "ratio":
						ratio := 1.0
						if sampler.Ratio.Valid {
							ratio = sampler.Ratio.Float64
						}
						sam = trace.TraceIDRatioBased(ratio)
					case "never":
						sam = trace.NeverSample()
					default:
						slog.Error("telemetry.oltp.trace.sample.based not one of always|never|ratio, will use never as default",
							"based", sampler.Based,
						)
						sam = trace.NeverSample()
					}
					for _, s := range sampler.Options {
						switch s {
						case "withRemote":
							options = append(options, trace.WithRemoteParentSampled(sam))
						case "withoutRemote":
							options = append(options, trace.WithRemoteParentNotSampled(sam))
						case "withLocal":
							options = append(options, trace.WithLocalParentSampled(sam))
						case "withoutLocal":
							options = append(options, trace.WithLocalParentNotSampled(sam))
						default:
							slog.Warn("unknown options", "name", s)
						}
					}
					opts = append(opts, trace.WithSampler(trace.ParentBased(sam, options...)))
				default:
					slog.Error("sampler.name not one of always|never|parent, will use system default",
						"name", sampler.Name,
					)
				}
			}
		}

	}
	traceProvider := trace.NewTracerProvider(opts...)
	return traceProvider, nil
}
