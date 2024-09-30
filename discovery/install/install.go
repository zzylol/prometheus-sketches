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
	_ "github.com/zzylol/prometheus-sketches/discovery/aws"          // register aws
	_ "github.com/zzylol/prometheus-sketches/discovery/azure"        // register azure
	_ "github.com/zzylol/prometheus-sketches/discovery/consul"       // register consul
	_ "github.com/zzylol/prometheus-sketches/discovery/digitalocean" // register digitalocean
	_ "github.com/zzylol/prometheus-sketches/discovery/dns"          // register dns
	_ "github.com/zzylol/prometheus-sketches/discovery/eureka"       // register eureka
	_ "github.com/zzylol/prometheus-sketches/discovery/file"         // register file
	_ "github.com/zzylol/prometheus-sketches/discovery/gce"          // register gce
	_ "github.com/zzylol/prometheus-sketches/discovery/hetzner"      // register hetzner
	_ "github.com/zzylol/prometheus-sketches/discovery/http"         // register http
	_ "github.com/zzylol/prometheus-sketches/discovery/ionos"        // register ionos
	_ "github.com/zzylol/prometheus-sketches/discovery/kubernetes"   // register kubernetes
	_ "github.com/zzylol/prometheus-sketches/discovery/linode"       // register linode
	_ "github.com/zzylol/prometheus-sketches/discovery/marathon"     // register marathon
	_ "github.com/zzylol/prometheus-sketches/discovery/moby"         // register moby
	_ "github.com/zzylol/prometheus-sketches/discovery/nomad"        // register nomad
	_ "github.com/zzylol/prometheus-sketches/discovery/openstack"    // register openstack
	_ "github.com/zzylol/prometheus-sketches/discovery/ovhcloud"     // register ovhcloud
	_ "github.com/zzylol/prometheus-sketches/discovery/puppetdb"     // register puppetdb
	_ "github.com/zzylol/prometheus-sketches/discovery/scaleway"     // register scaleway
	_ "github.com/zzylol/prometheus-sketches/discovery/triton"       // register triton
	_ "github.com/zzylol/prometheus-sketches/discovery/uyuni"        // register uyuni
	_ "github.com/zzylol/prometheus-sketches/discovery/vultr"        // register vultr
	_ "github.com/zzylol/prometheus-sketches/discovery/xds"          // register xds
	_ "github.com/zzylol/prometheus-sketches/discovery/zookeeper"    // register zookeeper
)
