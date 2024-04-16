// Copyright 2021 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tracing

import (
	"testing"

	"github.com/go-kit/log"
	config_util "github.com/prometheus/common/config"
	"github.com/stretchr/testify/require"

	"github.com/prometheus/prometheus/config"
)

func TestInstallingNewTracerProvider(t *testing.T) {
	m := NewManager(log.NewNopLogger())
	cfg := config.Config{
		TracingConfig: config.TracingConfig{
			Endpoint:   "localhost:1234",
			ClientType: config.TracingClientGRPC,
		},
	}

	require.NoError(t, m.ApplyConfig(&cfg))
}

func TestReinstallingTracerProvider(t *testing.T) {
	m := NewManager(log.NewNopLogger())
	cfg := config.Config{
		TracingConfig: config.TracingConfig{
			Endpoint:   "localhost:1234",
			ClientType: config.TracingClientGRPC,
			Headers:    map[string]string{"foo": "bar"},
		},
	}

	require.NoError(t, m.ApplyConfig(&cfg))

	// Trying to apply the same config should not reinstall provider.
	require.NoError(t, m.ApplyConfig(&cfg))

	cfg2 := config.Config{
		TracingConfig: config.TracingConfig{
			Endpoint:   "localhost:1234",
			ClientType: config.TracingClientHTTP,
			Headers:    map[string]string{"bar": "foo"},
		},
	}

	require.NoError(t, m.ApplyConfig(&cfg2))

	// Setting previously unset option should reinstall provider.
	cfg2.TracingConfig.Compression = "gzip"
	require.NoError(t, m.ApplyConfig(&cfg2))
}

func TestReinstallingTracerProviderWithTLS(t *testing.T) {
	m := NewManager(log.NewNopLogger())
	cfg := config.Config{
		TracingConfig: config.TracingConfig{
			Endpoint:   "localhost:1234",
			ClientType: config.TracingClientGRPC,
			TLSConfig: config_util.TLSConfig{
				CAFile: "testdata/ca.cer",
			},
		},
	}

	require.NoError(t, m.ApplyConfig(&cfg))

	// Trying to apply the same config with TLS should reinstall provider.
	require.NoError(t, m.ApplyConfig(&cfg))
}

func TestUninstallingTracerProvider(t *testing.T) {
	m := NewManager(log.NewNopLogger())
	cfg := config.Config{
		TracingConfig: config.TracingConfig{
			Endpoint:   "localhost:1234",
			ClientType: config.TracingClientGRPC,
		},
	}

	require.NoError(t, m.ApplyConfig(&cfg))

	// Uninstall by passing empty config.
	cfg2 := config.Config{
		TracingConfig: config.TracingConfig{},
	}

	require.NoError(t, m.ApplyConfig(&cfg2))
}

func TestTracerProviderShutdown(t *testing.T) {
	m := NewManager(log.NewNopLogger())
	cfg := config.Config{
		TracingConfig: config.TracingConfig{
			Endpoint:   "localhost:1234",
			ClientType: config.TracingClientGRPC,
		},
	}

	require.NoError(t, m.ApplyConfig(&cfg))
	m.Stop()

	// Check if we closed the done channel.
	_, ok := <-m.done
	require.False(t, ok)
}
