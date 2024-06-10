package main

import (
	"context"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.uber.org/zap"
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	_, span := otel.Tracer("example-tracer").Start(r.Context(), "handleRequest")
	defer span.End()

	zap.L().Info("Handling request")
	w.Write([]byte("Hello, World!"))
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	ctx := context.Background()
	traceExporter, err := otlptrace.New(ctx, otlptracehttp.NewClient())
	if err != nil {
		sugar.Fatal("failed to create trace exporter: ", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("ExampleService"),
		)),
	)
	otel.SetTracerProvider(tp)

	http.Handle("/", otelhttp.NewHandler(http.HandlerFunc(exampleHandler), "Example"))
	sugar.Fatal(http.ListenAndServe(":8080", nil))
}
