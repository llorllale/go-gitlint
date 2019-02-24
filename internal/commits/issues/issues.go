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

package issues

import (
	"fmt"
	"io"

	"github.com/llorllale/go-gitlint/internal/commits"
)

// Issue is a problem found with a commit.
type Issue struct {
	Desc   string
	Commit commits.Commit
}

func (i *Issue) String() string {
	return fmt.Sprintf("Issue{Desc=%s Commit=%+v}", i.Desc, i.Commit)
}

// Issues is a collection of Issues.
type Issues func() []Issue

// Collected returns a collection of issues identified.
func Collected(filters []func(c *commits.Commit) Issue, cmts commits.Commits) Issues {
	return func() []Issue {
		issues := make([]Issue, 0)
		for _, c := range cmts() {
			for _, f := range filters {
				if issue := f(c); issue != (Issue{}) {
					issues = append(issues, issue)
					break
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
		for i := range iss {
			_, err := w.Write(
				[]byte(fmt.Sprintf("%s%s", iss[i].String(), sep)),
			)
			if err != nil {
				panic(err)
			}
		}
		return iss
	}
}
