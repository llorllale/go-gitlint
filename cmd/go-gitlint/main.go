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
	"strings"

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
	path             = kingpin.Flag("path", `Path to the git repo (default: ".").`).Default(".").String()                                                                                         //nolint[gochecknoglobals]
	subjectRegex     = kingpin.Flag("subject-regex", `Commit subject line must conform to this regular expression (default: ".*").`).Default(".*").String()                                       //nolint[gochecknoglobals]
	subjectMaxLength = kingpin.Flag("subject-maxlen", "Max length for commit subject line (default: math.MaxInt32 - 1).").Default(strconv.Itoa(math.MaxInt32 - 1)).Int()                          //nolint[gochecknoglobals]
	subjectMinLength = kingpin.Flag("subject-minlen", "Min length for commit subject line (default: 0).").Default("0").Int()                                                                      //nolint[gochecknoglobals]
	bodyRegex        = kingpin.Flag("body-regex", `Commit message body must conform to this regular expression (default: ".*").`).Default(".*").String()                                          //nolint[gochecknoglobals]
	bodyMaxLength    = kingpin.Flag("body-maxlen", `Max length for commit body (default: math.MaxInt32 - 1)`).Default(strconv.Itoa(math.MaxInt32 - 1)).Int()                                      //nolint[gochecknoglobals]
	since            = kingpin.Flag("since", `A date in "yyyy-MM-dd" format starting from which commits will be analyzed (default: "1970-01-01").`).Default("1970-01-01").String()                //nolint[gochecknoglobals]
	msgFile          = kingpin.Flag("msg-file", `Only analyze the commit message found in this file (default: "").`).Default("").String()                                                         //nolint[gochecknoglobals]
	maxParents       = kingpin.Flag("max-parents", `Max number of parents a commit can have in order to be analyzed (default: 1). Useful for excluding merge commits.`).Default("1").Int()        //nolint[gochecknoglobals]
	authorNames      = kingpin.Flag("excl-author-names", "Don't lint commits with authors whose names match these comma-separated regular expressions (default: '$a').").Default("$a").String()   //nolint[gochecknoglobals]
	authorEmails     = kingpin.Flag("excl-author-emails", "Don't lint commits with authors whose emails match these comma-separated regular expressions (default: '$a').").Default("$a").String() //nolint[gochecknoglobals]
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
						issues.OfSubjectMaxLength(*subjectMaxLength),
						issues.OfSubjectMinLength(*subjectMinLength),
						issues.OfBodyRegex(*bodyRegex),
						issues.OfBodyMaxLength(*bodyMaxLength),
					},
					try(
						len(*msgFile) > 0,
						func() commits.Commits {
							file, err := os.Open(*msgFile)
							if err != nil {
								panic(err)
							}
							return commits.MsgIn(file)
						},
						func() commits.Commits {
							return commits.NotAuthoredByNames(
								strings.Split(*authorNames, ","),
								commits.NotAuthoredByEmails(
									strings.Split(*authorEmails, ","),
									commits.WithMaxParents(
										*maxParents,
										commits.Since(
											*since,
											commits.In(
												repo.Filesystem(*path),
											),
										),
									),
								),
							)
						},
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
	if _, err := kingpin.CommandLine.Parse(unique(args)); err != nil {
		panic(err)
	}
}

func unique(args []string) []string {
	u := make([]string, 0)
	flags := make([]string, 0)
	for _, a := range args {
		name := strings.Split(a, "=")[0]
		if !contains(name, flags) {
			u = append(u, a)
			flags = append(flags, name)
		}
	}
	return u
}

func contains(s string, strs []string) bool {
	for _, str := range strs {
		if s == str {
			return true
		}
	}
	return false
}

func try(cond bool, actual, dflt func() commits.Commits) commits.Commits {
	if cond {
		return actual()
	}
	return dflt()
}
