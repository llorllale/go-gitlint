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

package repo_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"

	"github.com/llorllale/go-gitlint/internal/repo"
)

func TestFilesystem(t *testing.T) {
	msgs := []string{"commit1", "commit2", "commit3"}
	r, path := tmpGitRepo(t, msgs...)
	test := repo.Filesystem(path)()

	head, err := test.Head()
	require.NoError(t, err)

	iter, err := r.Log(&git.LogOptions{From: head.Hash()})
	require.NoError(t, err)

	err = iter.ForEach(func(c *object.Commit) error {
		assert.Contains(t, msgs, c.Message,
			"repo.Filesystem() did not return all commits")

		return nil
	})
	require.NoError(t, err)
}

func tmpGitRepo(t *testing.T, msgs ...string) (r *git.Repository, folder string) {
	var err error

	folder, err = ioutil.TempDir(
		"",
		strings.ReplaceAll(uuid.New().String(), "-", ""),
	)
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, os.RemoveAll(folder))
	})

	r, err = git.PlainInit(folder, false)
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

	return r, folder
}
