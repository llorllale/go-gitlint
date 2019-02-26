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
	"testing"

	"github.com/llorllale/go-gitlint/internal/commits"
	"github.com/stretchr/testify/assert"
)

func TestCollected(t *testing.T) {
	expected := []*commits.Commit{
		{Hash: "123"},
		{Hash: "456"},
	}
	issues := Collected(
		[]Filter{
			func(c *commits.Commit) Issue {
				var issue Issue
				if c.Hash == "123" || c.Hash == "456" {
					issue = Issue{
						Desc:   "test",
						Commit: *c,
					}
				}
				return issue
			},
		},
		func() []*commits.Commit {
			return append(expected, &commits.Commit{Hash: "789"})
		},
	)()
	assert.Len(t,
		issues,
		2,
		"issues.Collected() must return the filtered commits")
	for _, i := range issues {
		assert.Contains(t,
			expected, &i.Commit,
			"issues.Collected() must return the filtered commits")
	}
}

func TestPrinted(t *testing.T) {
	const sep = "-"
	issues := []Issue{
		{
			Desc: "issueA",
			Commit: commits.Commit{
				Hash:    "1",
				Message: "first commit",
			},
		},
		{
			Desc: "issueB",
			Commit: commits.Commit{
				Hash:    "2",
				Message: "second commit",
			},
		},
	}
	var expected string
	for _, i := range issues {
		expected = expected + i.String() + sep
	}
	writer := &mockWriter{}
	Printed(
		writer, sep,
		func() []Issue {
			return issues
		},
	)()
	assert.Equal(t,
		expected, writer.msg,
		"issues.Printed() must concatenate Issue.String() with the separator")
}

type mockWriter struct {
	msg string
}

func (m *mockWriter) Write(b []byte) (int, error) {
	m.msg += string(b)
	return len(b), nil
}
