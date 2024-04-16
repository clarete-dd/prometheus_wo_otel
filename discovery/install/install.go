// Copyright 2020 The Prometheus Authors
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

// Package install has the side-effect of registering all builtin
// service discovery config types.
package install

import (
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/aws"          // register aws
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/azure"        // register azure
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/consul"       // register consul
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/digitalocean" // register digitalocean
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/dns"          // register dns
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/eureka"       // register eureka
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/file"         // register file
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/gce"          // register gce
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/hetzner"      // register hetzner
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/http"         // register http
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/ionos"        // register ionos
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/kubernetes"   // register kubernetes
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/linode"       // register linode
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/marathon"     // register marathon
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/nomad"        // register nomad
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/openstack"    // register openstack
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/ovhcloud"     // register ovhcloud
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/puppetdb"     // register puppetdb
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/scaleway"     // register scaleway
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/triton"       // register triton
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/uyuni"        // register uyuni
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/vultr"        // register vultr
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/xds"          // register xds
	_ "github.com/clarete-dd/prometheus_wo_otel/discovery/zookeeper"    // register zookeeper
)
