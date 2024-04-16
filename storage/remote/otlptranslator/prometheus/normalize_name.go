// DO NOT EDIT. COPIED AS-IS. SEE ../README.md

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheus // import "github.com/clarete-dd/prometheus_wo_otel/storage/remote/otlptranslator/prometheus"

import (
	"strings"
	"unicode"
)

// The map to translate OTLP units to Prometheus units
// OTLP metrics use the c/s notation as specified at https://ucum.org/ucum.html
// Prometheus best practices for units: https://prometheus.io/docs/practices/naming/#base-units
// OpenMetrics specification for units: https://github.com/OpenObservability/OpenMetrics/blob/main/specification/OpenMetrics.md#units-and-base-units
var unitMap = map[string]string{

	// Time
	"d":   "days",
	"h":   "hours",
	"min": "minutes",
	"s":   "seconds",
	"ms":  "milliseconds",
	"us":  "microseconds",
	"ns":  "nanoseconds",

	// Bytes
	"By":   "bytes",
	"KiBy": "kibibytes",
	"MiBy": "mebibytes",
	"GiBy": "gibibytes",
	"TiBy": "tibibytes",
	"KBy":  "kilobytes",
	"MBy":  "megabytes",
	"GBy":  "gigabytes",
	"TBy":  "terabytes",

	// SI
	"m": "meters",
	"V": "volts",
	"A": "amperes",
	"J": "joules",
	"W": "watts",
	"g": "grams",

	// Misc
	"Cel": "celsius",
	"Hz":  "hertz",
	"1":   "",
	"%":   "percent",
}

// The map that translates the "per" unit
// Example: s => per second (singular)
var perUnitMap = map[string]string{
	"s":  "second",
	"m":  "minute",
	"h":  "hour",
	"d":  "day",
	"w":  "week",
	"mo": "month",
	"y":  "year",
}

func removeUnitSuffixes(nameTokens []string, unit string) []string {
	l := len(nameTokens)
	unitTokens := strings.Split(unit, "_")
	lu := len(unitTokens)

	if lu == 0 || l <= lu {
		return nameTokens
	}

	suffixed := true
	for i := range unitTokens {
		if nameTokens[l-i-1] != unitTokens[lu-i-1] {
			suffixed = false
			break
		}
	}

	if suffixed {
		return nameTokens[:l-lu]
	}

	return nameTokens
}

func removeSuffix(tokens []string, suffix string) []string {
	l := len(tokens)
	if tokens[l-1] == suffix {
		return tokens[:l-1]
	}

	return tokens
}

// Clean up specified string so it's Prometheus compliant
func CleanUpString(s string) string {
	return strings.Join(strings.FieldsFunc(s, func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsDigit(r) }), "_")
}

func RemovePromForbiddenRunes(s string) string {
	return strings.Join(strings.FieldsFunc(s, func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' && r != ':' }), "_")
}

// Retrieve the Prometheus "basic" unit corresponding to the specified "basic" unit
// Returns the specified unit if not found in unitMap
func unitMapGetOrDefault(unit string) string {
	if promUnit, ok := unitMap[unit]; ok {
		return promUnit
	}
	return unit
}

// Retrieve the Prometheus "per" unit corresponding to the specified "per" unit
// Returns the specified unit if not found in perUnitMap
func perUnitMapGetOrDefault(perUnit string) string {
	if promPerUnit, ok := perUnitMap[perUnit]; ok {
		return promPerUnit
	}
	return perUnit
}

// Returns whether the slice contains the specified value
func contains(slice []string, value string) bool {
	for _, sliceEntry := range slice {
		if sliceEntry == value {
			return true
		}
	}
	return false
}

// Remove the specified value from the slice
func removeItem(slice []string, value string) []string {
	newSlice := make([]string, 0, len(slice))
	for _, sliceEntry := range slice {
		if sliceEntry != value {
			newSlice = append(newSlice, sliceEntry)
		}
	}
	return newSlice
}
