// Copyright 2019 George Aristy
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

	"github.com/llorllale/go-gitlint/internal/commits"
	"github.com/llorllale/go-gitlint/internal/commits/filter"
	"github.com/llorllale/go-gitlint/internal/commits/issues"
	"github.com/llorllale/go-gitlint/internal/repo"
)

// @todo #4 The path should be passed in as a command line
//  flag instead of being hard coded. All other configuration
//  options should be passed in through CLI as well.
func main() {
	os.Exit(
		len(
			issues.Printed(
				os.Stdout, "\n",
				issues.Collected(
					[]func(*commits.Commit) issues.Issue{
						filter.OfSubjectRegex(".{,1}"),
						filter.OfBodyRegex(".{,1}"),
					},
					commits.In(
						repo.Filesystem("."),
					),
				),
			)(),
		),
	)
}
