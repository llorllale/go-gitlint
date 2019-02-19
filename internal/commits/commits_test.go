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
	"bytes"
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
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func TestIn(t *testing.T) {
	msgs := []string{"commit1", "commit2", "commit3"}
	r, cleanup := tmpRepo(msgs...)
	defer cleanup()
	commits := In(r)()
	assert.Len(t, commits, len(msgs),
		"commits.In() did not return the correct number of commits")
	for i, msg := range msgs {
		assert.Equal(t, msg, commits[len(commits)-i-1].Body,
			"commits.In() returned commits with incorrect body messages")
	}
}

func TestPrinted(t *testing.T) {
	commit := &Commit{
		Hash:    "abc123",
		Subject: "subject",
		Body:    "body",
	}
	const sep = " "
	buffer := &bytes.Buffer{}
	_ = Printed(
		func() []*Commit { return []*Commit{commit} },
		buffer,
		sep,
	)()
	assert.Equal(t, (&pretty{commit}).String()+sep, string(buffer.Bytes()),
		"commits.Printed() did not pretty-print the commit correctly")
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
