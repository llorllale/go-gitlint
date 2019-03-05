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

package commits

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/llorllale/go-gitlint/internal/repo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func TestCommitID(t *testing.T) {
	const ID = "test ID"
	assert.Equal(t,
		(&Commit{Hash: ID}).ID(), ID,
		"Commit.ID() must return the commit's hash")
}

func TestCommitShortID(t *testing.T) {
	const ID = "c26cf8af130955c5c67cfea96f9532680b963628"
	assert.Equal(t,
		(&Commit{Hash: ID}).ShortID(),
		ID[:7],
		"Commit.ShortID() must equal the first 7 characters of the commit's hash")
}

func TestCommitSubject(t *testing.T) {
	const subject = "test subject"
	assert.Equal(t,
		(&Commit{Message: subject + "\n\ntest body"}).Subject(),
		subject,
		`Commit.Subject() must return the substring before the first \n`)
}

func TestCommitBody(t *testing.T) {
	const body = "test body"
	assert.Equal(t,
		(&Commit{Message: "test subject\n\n" + body}).Body(),
		body,
		`Commit.Body() must return the substring after the first \n\n`)
}

func TestIn(t *testing.T) {
	msgs := []string{"subject1\n\nbody1", "subject2\n\nbody2", "subject3\n\nbody3"}
	r, cleanup := tmpRepo(msgs...)
	defer cleanup()
	commits := In(r)()
	assert.Len(t, commits, len(msgs),
		"commits.In() did not return the correct number of commits")
	for i, msg := range msgs {
		commit := commits[len(commits)-i-1]
		assert.Equal(t, msg, commit.Subject()+"\n\n"+commit.Body(),
			"commits.In() returned commits with incorrect message subjects or bodies")
	}
}

func TestSince(t *testing.T) {
	before, err := time.Parse("2006-01-02", "2017-10-25")
	require.NoError(t, err)
	since, err := time.Parse("2006-01-02", "2019-01-01")
	require.NoError(t, err)
	after, err := time.Parse("2006-01-02", "2019-03-03")
	require.NoError(t, err)
	commits := Since(
		"2019-01-01",
		func() []*Commit {
			return []*Commit{
				{Date: before},
				{Date: since},
				{Date: after},
			}
		},
	)()
	assert.Len(t, commits, 2)
	assert.Contains(t, commits, &Commit{Date: since})
	assert.Contains(t, commits, &Commit{Date: after})
}

func TestMsgIn(t *testing.T) {
	const message = "test subject\n\ntest body"
	commits := MsgIn(strings.NewReader(message))()
	assert.Len(t, commits, 1)
	assert.Equal(t, "test subject", commits[0].Subject())
	assert.Equal(t, "test body", commits[0].Body())
}

func TestWithMaxParents(t *testing.T) {
	const max = 1
	commits := WithMaxParents(max, func() []*Commit {
		return []*Commit{
			{NumParents: max},
			{NumParents: 2},
			{NumParents: 3},
		}
	})()
	assert.Len(t, commits, 1)
	assert.Equal(t, commits[0].NumParents, max)
}

// A git repo initialized and with one commit per each of the messages provided.
// This repo is created in a temporary directory; use the cleanup function
// to delete it afterwards.
func tmpRepo(msgs ...string) (r repo.Repo, cleanup func()) {
	folder, err := ioutil.TempDir(
		"",
		strings.Replace(uuid.New().String(), "-", "", -1), //nolint[gocritic]
	)
	panicIf(err)
	cleanup = func() {
		panicIf(os.RemoveAll(folder))
	}
	r = func() *git.Repository {
		r, err := git.PlainInit(folder, false)
		panicIf(err)
		wt, err := r.Worktree()
		panicIf(err)
		for i, msg := range msgs {
			file := fmt.Sprintf("msg%d.txt", i)
			panicIf(ioutil.WriteFile(filepath.Join(folder, file), []byte(msg), 0644))
			_, err = wt.Add(file)
			panicIf(err)
			_, err = wt.Commit(msg, &git.CommitOptions{
				Author: &object.Signature{
					Name:  "John Doe",
					Email: "john@doe.org",
					When:  time.Now(),
				},
			})
			panicIf(err)
		}
		return r
	}
	return r, cleanup
}

// panics if err is not nil.
func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
