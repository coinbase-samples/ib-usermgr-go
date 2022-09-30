package config

import (
	xray "contrib.go.opencensus.io/exporter/aws"
	"go.opentelemetry.io/otel"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// Init configures an OpenTelemetry exporter and trace provider.
func Init(app AppConfig) (*sdktrace.TracerProvider, error) {
	var exporter sdktrace.SpanExporter
	var sampler sdktrace.Sampler
	var err error
	if app.IsLocalEnv() {
		exporter, err = stdout.New()
		if err != nil {
			return nil, err
		}
		sampler = sdktrace.AlwaysSample()
	} else {
		exporter, err := xray.NewExporter(
			xray.WithVersion("latest"),
			// Add your AWS region.
			xray.WithRegion(app.Region),
		)
		if err != nil {
			// Handle any error.
			return nil, err
		}
		// Do not forget to call Flush() before the application terminates.
		defer exporter.Flush()

		sampler = sdktrace.AlwaysSample()
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sampler),
		sdktrace.WithBatcher(exporter),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}
