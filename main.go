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
	"fmt"
	"os"
	"strconv"
	"time"
)

func Echo() string {
	msg := os.Getenv("GOECHO")
	return msg
}

func main() {
	msg := Echo()
	if msg == "" {
		fmt.Println("The `GOECHO` variable was empty")
		os.Exit(1)
	}
	fmt.Println(msg)

	if os.Getenv("GOSLEEP") != "" {
		sleepSeconds, err := strconv.Atoi(os.Getenv("GOSLEEP"))
		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("Sleeping for %d seconds.", sleepSeconds))
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}

	os.Exit(0)
}
