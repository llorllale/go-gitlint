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

package commits_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/llorllale/go-gitlint/internal/commits"
	"github.com/llorllale/go-gitlint/internal/repo"
)

func TestCommitID(t *testing.T) {
	const ID = "test ID"

	assert.Equal(t,
		(&commits.Commit{Hash: ID}).ID(), ID,
		"Commit.ID() must return the commit's hash")
}

func TestCommitShortID(t *testing.T) {
	const ID = "c26cf8af130955c5c67cfea96f9532680b963628"

	assert.Equal(t,
		(&commits.Commit{Hash: ID}).ShortID(),
		ID[:7],
		"Commit.ShortID() must equal the first 7 characters of the commit's hash")
}

func TestCommitSubject(t *testing.T) {
	const subject = "test subject"

	assert.Equal(t,
		(&commits.Commit{Message: subject + "\n\ntest body"}).Subject(),
		subject,
		`Commit.Subject() must return the substring before the first \n`)
}

func TestCommitBody(t *testing.T) {
	const body = "test body"

	assert.Equal(t,
		(&commits.Commit{Message: "test subject\n\n" + body}).Body(),
		body,
		`Commit.Body() must return the substring after the first \n\n`)
}

func TestIn(t *testing.T) {
	msgs := []string{"subject1\n\nbody1", "subject2\n\nbody2", "subject3\n\nbody3"}
	r := tmpRepo(t, msgs...)
	cmits := commits.In(r)()

	assert.Len(t, cmits, len(msgs),
		"commits.In() did not return the correct number of commits")

	for i, msg := range msgs {
		commit := cmits[len(cmits)-i-1]
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

	cmits := commits.Since(
		"2019-01-01",
		func() []*commits.Commit {
			return []*commits.Commit{
				{Date: before},
				{Date: since},
				{Date: after},
			}
		},
	)()

	assert.Len(t, cmits, 2)
	assert.Contains(t, cmits, &commits.Commit{Date: since})
	assert.Contains(t, cmits, &commits.Commit{Date: after})
}

func TestMsgIn(t *testing.T) {
	const message = "test subject\n\ntest body"

	cmits := commits.MsgIn(strings.NewReader(message))()

	assert.Len(t, cmits, 1)
	assert.Equal(t, "test subject", cmits[0].Subject())
	assert.Equal(t, "test body", cmits[0].Body())
}

func TestWithMaxParents(t *testing.T) {
	const max = 1

	cmits := commits.WithMaxParents(max, func() []*commits.Commit {
		return []*commits.Commit{
			{NumParents: max},
			{NumParents: 2},
			{NumParents: 3},
		}
	})()

	assert.Len(t, cmits, 1)
	assert.Equal(t, cmits[0].NumParents, max)
}

func TestNotAuthored(t *testing.T) {
	filtered := &commits.Commit{Author: randomAuthor()}
	expected := []*commits.Commit{
		{Author: randomAuthor()},
		{Author: randomAuthor()},
		{Author: randomAuthor()},
		{Author: randomAuthor()},
	}

	actual := commits.NotAuthoredByNames(
		[]string{filtered.Author.Name},
		func() []*commits.Commit { return append(expected, filtered) },
	)()

	assert.Equal(t, expected, actual)

	actual = commits.NotAuthoredByEmails(
		[]string{filtered.Author.Email},
		func() []*commits.Commit { return append(expected, filtered) },
	)()

	assert.Equal(t, expected, actual)
}

func randomAuthor() *commits.Author {
	return &commits.Author{
		Name:  uuid.New().String(),
		Email: uuid.New().String() + "@test.com",
	}
}

// A git repo initialized and with one commit per each of the messages provided.
// This repo is created in a temporary directory; use the cleanup function
// to delete it afterwards.
func tmpRepo(t *testing.T, msgs ...string) repo.Repo {
	folder, err := ioutil.TempDir(
		"",
		strings.ReplaceAll(uuid.New().String(), "-", ""),
	)
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, os.RemoveAll(folder))
	})

	return func() *git.Repository {
		r, err := git.PlainInit(folder, false)
		require.NoError(t, err)

		wt, err := r.Worktree()
		require.NoError(t, err)

		for i, msg := range msgs {
			file := fmt.Sprintf("msg%d.txt", i)

			err = ioutil.WriteFile(filepath.Join(folder, file), []byte(msg), 0600)
			require.NoError(t, err)

			_, err = wt.Add(file)
			require.NoError(t, err)

			_, err = wt.Commit(msg, &git.CommitOptions{
				Author: &object.Signature{
					Name:  "John Doe",
					Email: "john@doe.org",
					When:  time.Now(),
				},
			})
			require.NoError(t, err)
		}

		return r
	}
}
