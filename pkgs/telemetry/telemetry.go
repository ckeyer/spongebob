package telemetry

import (
	"fmt"
	"os"

	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

// Metrics
func NewRegistry() (*prometheus.Registry, error) {
	prometheusRegistry := prometheus.NewRegistry()
	err := prometheusRegistry.Register(prometheus.NewGoCollector())
	if err != nil {
		return nil, fmt.Errorf("failed to register Go runtime metrics: %v", err)
	}

	err = prometheusRegistry.Register(prometheus.NewProcessCollector(os.Getpid(), ""))
	if err != nil {
		return nil, fmt.Errorf("failed to register process metrics: %v", err)
	}

	grpcMetrics := grpcprometheus.NewServerMetrics()
	err = prometheusRegistry.Register(grpcMetrics)
	if err != nil {
		return nil, fmt.Errorf("failed to register gRPC server metrics: %v", err)
	}

	return prometheusRegistry, nil
}
