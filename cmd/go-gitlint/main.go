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
	"math"
	"os"
	"strconv"

	"github.com/llorllale/go-gitlint/internal/commits"
	"github.com/llorllale/go-gitlint/internal/issues"
	"github.com/llorllale/go-gitlint/internal/repo"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// @todo #9 Global variables are a code smell (especially those in filterse.go).
//  They promote coupling across different components inside the same package.
//  Figure out a way to remove these global variables. Whatever command line
//  parser we choose should be able to auto-generate usage.
var (
	path          = kingpin.Flag("path", `Path to the git repo (default: ".").`).Default(".").String()                                                                         //nolint[gochecknoglobals]
	subjectRegex  = kingpin.Flag("subject-regex", `Commit subject line must conform to this regular expression (default: ".*").`).Default(".*").String()                       //nolint[gochecknoglobals]
	subjectLength = kingpin.Flag("subject-len", "Commit subject line cannot exceed this length (default: math.MaxUint32).").Default(strconv.Itoa(math.MaxUint32)).Int()        //nolint[gochecknoglobals]
	bodyRegex     = kingpin.Flag("body-regex", `Commit message body must conform to this regular expression (default: ".*").`).Default(".*").String()                          //nolint[gochecknoglobals]
	since         = kingpin.Flag("since", `A date in "yyyy-MM-dd" format starting from which commits will be analyzed (default: "1970-01-01")`).Default("1970-01-01").String() //nolint[gochecknoglobals]
)

func main() {
	configure()
	os.Exit(
		len(
			issues.Printed(
				os.Stdout, "\n",
				issues.Collected(
					[]issues.Filter{
						issues.OfSubjectRegex(*subjectRegex),
						issues.OfBodyRegex(*bodyRegex),
						issues.OfSubjectLength(*subjectLength),
					},
					commits.Since(
						*since,
						commits.In(
							repo.Filesystem(*path),
						),
					),
				),
			)(),
		),
	)
}

func configure() {
	args := os.Args[1:]
	const file = ".gitlint"
	if _, err := os.Stat(file); err == nil {
		config, err := kingpin.ExpandArgsFromFile(file)
		if err != nil {
			panic(err)
		}
		args = append(args, config...)
	}
	if _, err := kingpin.CommandLine.Parse(args); err != nil {
		panic(err)
	}
}
