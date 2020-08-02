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

package issues_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/llorllale/go-gitlint/internal/commits"
	"github.com/llorllale/go-gitlint/internal/issues"
)

func TestOfSubjectRegexMatch(t *testing.T) {
	assert.Zero(t,
		issues.OfSubjectRegex(`\(#\d+\) [\w ]{10,50}`)(
			&commits.Commit{
				Message: "(#123) This is a good subject",
			},
		),
		"filter.OfSubjectRegex() must match if the commit's subject matches the regex",
	)
}

func TestOfSubjectRegexNonMatch(t *testing.T) {
	assert.NotZero(t,
		issues.OfSubjectRegex(`\(#\d+\) [\w ]{,50}`)(
			&commits.Commit{
				Message: "I break all the rules!",
			},
		),
		"filter.OfSubjectRegex() must not match if the commit's subject does not match the regex",
	)
}

func TestOfBodyRegexMatch(t *testing.T) {
	assert.Zero(t,
		issues.OfBodyRegex(`^.{10,20}$`)(
			&commits.Commit{
				Message: "subject\n\nBetween 10 and 20",
			},
		),
		"filter.OfBodyRegex() must match if the commit's subject matches the regex",
	)
}

func TestOfBodyRegexNonMatch(t *testing.T) {
	assert.NotZero(t,
		issues.OfBodyRegex(`^.{10,20}$`)(
			&commits.Commit{
				Message: "subject\n\nMore than twenty characters!",
			},
		),
		"filter.OfBodyRegex() must not match if the commit's subject does not match the regex",
	)
}

func TestOfSubjectMaxLengthMatch(t *testing.T) {
	assert.NotZero(t,
		issues.OfSubjectMaxLength(5)(
			&commits.Commit{
				Message: "very very very VERY long subject\n\nand body",
			},
		),
		"filter.OfSubjectMaxLength() must match if the commit's subject is too long",
	)
}

func TestOfSubjectMaxLengthNonMatch(t *testing.T) {
	assert.Zero(t,
		issues.OfSubjectMaxLength(10)(
			&commits.Commit{
				Message: "short\n\nmessage",
			},
		),
		"filter.OfSubjectMaxLength() must not match if the commit's subject is not too long",
	)
}

func TestOfSubjectMinLengthMatch(t *testing.T) {
	assert.NotZero(t,
		issues.OfSubjectMinLength(10)(
			&commits.Commit{
				Message: "short\n\nand body",
			},
		),
		"filter.OfSubjectMinLength() must match if the commit's subject is too short",
	)
}

func TestOfSubjectMinLengthNonMatch(t *testing.T) {
	assert.Zero(t,
		issues.OfSubjectMinLength(10)(
			&commits.Commit{
				Message: "not too short subject\n\nmessage",
			},
		),
		"filter.OfSubjectMinLength() must not match if the commit's subject is not too short",
	)
}

func TestOfBodyMaxLengthMatch(t *testing.T) {
	assert.NotZero(t,
		issues.OfBodyMaxLength(1)(
			&commits.Commit{
				Message: "subject\n\nclearly, this commit has a long body",
			},
		),
	)
}

func TestOfBodyMaxLengthNonMatch(t *testing.T) {
	assert.Zero(t,
		issues.OfBodyMaxLength(math.MaxInt32)(
			&commits.Commit{
				Message: "subject\n\nclearly, this commit cannot exceed this max",
			},
		),
	)
}
