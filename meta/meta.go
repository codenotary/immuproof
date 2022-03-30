/*
Copyright © 2022 CodeNotary, Inc. All rights reserved
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package meta

import (
	"time"

	"github.com/adrg/xdg"
)

var version = ""

var static = ""

var gitCommit = ""
var gitBranch = ""

const AppName = "immuproof"
const LedgerHeaderName = "lc-ledger"

var DefaultConfigFolder = xdg.ConfigHome + "/" + AppName
var DefaultStateFolder = xdg.StateHome + "/" + AppName
var DefaultStateMapFileName = DefaultStateFolder + "/" + "state-history.json"

var DefaultCNCPort = 443
var DefaultCNCHost = "localhost"

var DefaultAuditInterval, _ = time.ParseDuration("1h")

// Version returns the current Immuproof version string
func Version() string {
	return version
}

// StaticBuild returns when the current Immuproof executable has been statically linked against libraries
func StaticBuild() bool {
	return static == "static"
}

// GitRevision returns the current CodeNotary Immuproof git revision string
func GitRevision() string {
	rev := gitCommit
	if gitBranch != "" {
		rev += " (" + gitBranch + ")"
	}
	return rev
}
