// Copyright 2015 Jake Dahn
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

package main

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEcho(t *testing.T) {

	Convey("Given a `GOECHO` environment variable", t, func() {
		os.Setenv("GOECHO", "this is a test")

		Convey("the output of Echo should be identical to the env var", func() {
			So(Echo(), ShouldEqual, "this is a test")
			os.Setenv("GOECHO", "and another")
			So(Echo(), ShouldEqual, "and another")
		})
	})

	Convey("Given an empty `GOECHO` environment variable", t, func() {
		os.Clearenv()

		Convey("Echo should be an empty string", func() {
			So(Echo(), ShouldEqual, "")
		})
	})

}
