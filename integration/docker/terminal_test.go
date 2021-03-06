// Copyright (c) 2017 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docker

import (
	. "github.com/clearcontainers/tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("terminal", func() {
	var (
		id string
	)

	BeforeEach(func() {
		id = randomDockerName()
	})

	AfterEach(func() {
		Expect(RemoveDockerContainer(id)).To(BeTrue())
		Expect(ExistDockerContainer(id)).NotTo(BeTrue())
	})

	Describe("terminal with docker", func() {
		Context("TERM env variable is set when allocating a tty", func() {
			It("should display the terminal's name", func() {
				stdout, _, exitCode := DockerRun("--name", id, "-t", Image, "env")
				Expect(exitCode).To(Equal(0))
				Expect(stdout).To(MatchRegexp("TERM=" + `[[:alnum:]]`))
			})
		})

		Context("TERM env variable is not set when not allocating a tty", func() {
			It("should not display the terminal's name", func() {
				stdout, _, exitCode := DockerRun( "--name", id, Image, "env")
				Expect(exitCode).To(Equal(0))
				Expect(stdout).NotTo(ContainSubstring("TERM"))
			})
		})

		Context("Check that pseudo tty is setup properly when allocating a tty", func() {
			It("should display the pseudo tty's name", func() {
				stdout, _, exitCode := DockerRun("--name", id, "-t", Image, "tty")
				Expect(exitCode).To(Equal(0))
				Expect(stdout).To(MatchRegexp("/dev/pts/" + `[[:alnum:]]`))
			})
		})
	})
})
