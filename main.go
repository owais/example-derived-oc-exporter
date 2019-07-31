package main

import (
	"log"

	"github.com/open-telemetry/opentelemetry-service/defaults"
	"github.com/open-telemetry/opentelemetry-service/service"

	"github.com/owais/example-derived-oc-exporter/opencensusexporter"
)

func main() {
	handleErr := func(err error) {
		if err != nil {
			log.Fatalf("Failed to run the service: %v", err)
		}
	}

	receivers, processors, exporters, err := defaults.Components()
	handleErr(err)

	exporters["opencensus"] = &opencensusexporter.Factory{}

	svc := service.New(receivers, processors, exporters)
	err = svc.StartUnified()
	handleErr(err)
}
