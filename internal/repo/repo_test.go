package repo

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
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func TestFilesystem(t *testing.T) {
	msgs := []string{"commit1", "commit2", "commit3"}
	r, path, cleanup := tmpGitRepo(msgs...)
	defer cleanup()
	test := Filesystem(path)()
	head, err := test.Head()
	require.NoError(t, err)
	iter, err := r.Log(&git.LogOptions{From: head.Hash()})
	require.NoError(t, err)
	_ = iter.ForEach(func(c *object.Commit) error { //nolint[errcheck]
		assert.Contains(t, msgs, c.Message,
			"repo.Filesystem() did not return all commits")
		return nil
	})
}

func tmpGitRepo(msgs ...string) (r *git.Repository, folder string, cleanup func()) {
	var err error
	folder, err = ioutil.TempDir(
		"",
		strings.Replace(uuid.New().String(), "-", "", -1), //nolint[gocritic]
	)
	panicIf(err)
	cleanup = func() {
		panicIf(os.RemoveAll(folder))
	}
	r, err = git.PlainInit(folder, false)
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
	return r, folder, cleanup
}

// panics if err is not nil.
func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
