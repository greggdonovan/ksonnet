// Copyright 2018 The kubecfg authors
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// +build e2e

package e2e

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("ks registry", func() {
	var a app

	BeforeEach(func() {
		a = e.initApp("")
		a.generateDeployedService()
	})

	Describe("list", func() {
		It("lists the currently configured registries", func() {
			o := a.runKs("registry", "list")
			assertExitStatus(o, 0)
			assertOutput("registry/list/output.txt", o.stdout)
		})
	})

	Describe("describe", func() {
		Context("incubator", func() {
			It("describe a registry", func() {
				o := a.runKs("registry", "describe", "incubator")
				assertExitStatus(o, 0)
				assertOutput("registry/describe/output.txt", o.stdout)
			})
		})
	})
})
