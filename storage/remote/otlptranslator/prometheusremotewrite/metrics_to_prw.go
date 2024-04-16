// DO NOT EDIT. COPIED AS-IS. SEE ../README.md

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewrite // import "github.com/clarete-dd/prometheus_wo_otel/storage/remote/otlptranslator/prometheusremotewrite"

type Settings struct {
	Namespace           string
	ExternalLabels      map[string]string
	DisableTargetInfo   bool
	ExportCreatedMetric bool
	AddMetricSuffixes   bool
	SendMetadata        bool
}
