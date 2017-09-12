// Copyright 2017 CoreOS, Inc.
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

package tests

import (
	"flag"
	"os"
	"testing"

	"github.com/coreos/init/tests/register"

	_ "github.com/coreos/init/tests/registry"
)

var flagBinaryPath string

func init() {
	flag.StringVar(&flagBinaryPath, "coreos-install", "coreos-install", "path to coreos-install binary")
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func TestCoreosInstall(t *testing.T) {
	for _, test := range register.Tests {
		t.Run(test.Name, func(t *testing.T) {
			test.BinaryPath = flagBinaryPath
			test.Run(t)
		})
	}
}
