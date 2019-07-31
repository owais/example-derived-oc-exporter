package opencensusexporter

import (
	"github.com/open-telemetry/opentelemetry-service/exporter/opencensusexporter"
)

type Config struct {
	opencensusexporter.Config `mapstructure:",squash"`

	ExtraField bool `mapstructure:"extra-field,omitempty"`
}
