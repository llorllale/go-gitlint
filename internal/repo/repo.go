package repo

import (
	git "gopkg.in/src-d/go-git.v4"
)

// Repo is an initialized git repository.
type Repo func() *git.Repository

// Filesystem is a pre-existing git repository on the filesystem
// with directory as root.
func Filesystem(directory string) Repo {
	return func() *git.Repository {
		repo, err := git.PlainOpen(directory)
		if err != nil {
			panic(err)
		}
		return repo
	}
}
