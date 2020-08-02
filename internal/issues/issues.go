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

// Package issues provides filters for problems found in commit messages.
package issues

import (
	"io"

	"github.com/fatih/color"
	"github.com/llorllale/go-gitlint/internal/commits"
)

// Issue is a problem found with a commit.
type Issue struct {
	Desc   string
	Commit commits.Commit
}

// Issues is a collection of Issues.
type Issues func() []Issue

// Collected returns a collection of issues identified.
func Collected(filters []Filter, cmts commits.Commits) Issues {
	return func() []Issue {
		issues := make([]Issue, 0)

		for _, c := range cmts() {
			for _, f := range filters {
				if issue := f(c); issue != (Issue{}) {
					issues = append(issues, issue)
				}
			}
		}

		return issues
	}
}

// Printed prints the issues to the writer.
func Printed(w io.Writer, sep string, issues Issues) Issues {
	return func() []Issue {
		iss := issues()

		for idx := range iss {
			i := iss[idx]

			_, err := color.New(color.Bold).Fprintf(w, "%s: ", i.Commit.ShortID())
			if err != nil {
				panic(err)
			}

			_, err = color.New(color.FgRed).Fprintf(w, "%s%s", i.Desc, sep)
			if err != nil {
				panic(err)
			}
		}

		return iss
	}
}
