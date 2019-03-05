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
	"regexp"

	"github.com/llorllale/go-gitlint/internal/commits"
)

// Filter identifies an issue with a commit.
// A filter returning a zero-valued Issue signals that it found no issue
// with the commit.
type Filter func(*commits.Commit) Issue

// OfSubjectRegex tests a commit's subject with the regex.
func OfSubjectRegex(regex string) Filter {
	return func(c *commits.Commit) Issue {
		var issue Issue
		matched, err := regexp.MatchString(regex, c.Subject())
		if err != nil {
			panic(err)
		}
		if !matched {
			issue = Issue{
				Desc:   fmt.Sprintf("subject does not match regex [%s]", regex),
				Commit: *c,
			}
		}
		return issue
	}
}

// OfBodyRegex tests a commit's body with the regex.
func OfBodyRegex(regex string) Filter {
	return func(c *commits.Commit) Issue {
		var issue Issue
		matched, err := regexp.MatchString(regex, c.Body())
		if err != nil {
			panic(err)
		}
		if !matched {
			issue = Issue{
				Desc:   fmt.Sprintf("body does not conform to regex [%s]", regex),
				Commit: *c,
			}
		}
		return issue
	}
}

// OfSubjectMaxLength checks that a commit's subject does not exceed this length.
func OfSubjectMaxLength(length int) Filter {
	return func(c *commits.Commit) Issue {
		var issue Issue
		if len(c.Subject()) > length {
			issue = Issue{
				Desc:   fmt.Sprintf("subject length exceeds max [%d]", length),
				Commit: *c,
			}
		}
		return issue
	}
}

// OfSubjectMinLength checks that a commit's subject's length is at least
// of length min.
func OfSubjectMinLength(min int) Filter {
	return func(c *commits.Commit) Issue {
		var issue Issue
		if len(c.Subject()) < min {
			issue = Issue{
				Desc:   fmt.Sprintf("subject length less than min [%d]", min),
				Commit: *c,
			}
		}
		return issue
	}
}
