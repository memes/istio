// Copyright 2019 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"istio.io/istio/pkg/config/protocol"
	"istio.io/istio/pkg/test/framework/components/echo"
	"istio.io/istio/pkg/test/framework/components/galley"
	"istio.io/istio/pkg/test/framework/components/namespace"
	"istio.io/istio/pkg/test/framework/components/pilot"
)

func EchoConfig(name string, ns namespace.Instance, headless bool, annos echo.Annotations, g galley.Instance, p pilot.Instance) echo.Config {
	return echo.Config{
		Service:        name,
		Namespace:      ns,
		ServiceAccount: true,
		Headless:       headless,
		Annotations:    annos,
		Ports: []echo.Port{
			{
				Name:     "http",
				Protocol: protocol.HTTP,
			},
			{
				Name:     "tcp",
				Protocol: protocol.TCP,
			},
			{
				Name:     "grpc",
				Protocol: protocol.GRPC,
			},
		},
		Galley: g,
		Pilot:  p,
	}
}
