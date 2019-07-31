package opencensusexporter

import (
	"github.com/open-telemetry/opentelemetry-service/config/configmodels"
	"github.com/open-telemetry/opentelemetry-service/consumer"
	"github.com/open-telemetry/opentelemetry-service/exporter"
	"github.com/open-telemetry/opentelemetry-service/exporter/opencensusexporter"
	"go.uber.org/zap"
)

const (
	typeStr = "opencensus"
)

type Factory struct {
	factory opencensusexporter.Factory
}

func (f *Factory) Type() string {
	return typeStr
}

func (f *Factory) CreateDefaultConfig() configmodels.Exporter {
	cfg := f.factory.CreateDefaultConfig()
	c := cfg.(*opencensusexporter.Config)
	return &Config{Config: *c, ExtraField: true}
}

func (f *Factory) CreateTraceExporter(logger *zap.Logger, config configmodels.Exporter) (consumer.TraceConsumer, exporter.StopFunc, error) {
	ocac := config.(*Config)
	// 1. can use this or custom method to generate options
	opts, err := f.factory.OCAgentOptions(logger, &ocac.Config)
	if err != nil {
		return nil, nil, err
	}
	// 2. can remove, add or change ocagent options here

	// 3. can use this or custom method to generate final exporter
	return f.factory.CreateOCAgent(logger, &ocac.Config, opts)
}

func (f *Factory) CreateMetricsExporter(logger *zap.Logger, config configmodels.Exporter) (consumer.MetricsConsumer, exporter.StopFunc, error) {
	return f.factory.CreateMetricsExporter(logger, config)
}
