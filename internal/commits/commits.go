package commits

import (
	"fmt"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// Commits returns commits.
// @todo #4 Figure out how to disable the golint check that
//  forces us to write redundant comments of the form
//  'comment on exported type Commits should be of the form
//  "Commits ..." (with optional leading article)' and rewrite
//  all comments.
type Commits func() []*Commit

// Commit holds data for a single git commit.
type Commit struct {
	Hash    string
	Subject string
	Body    string
}

// In returns commits in the path.
// @todo #4 These err checks are extremely annoying. Figure out
//  how to handle them elegantly and reduce the cyclo complexity
//  of this function (currently at 4).
func In(path string) Commits {
	repo, err := git.PlainOpen(path)
	if err != nil {
		panic(err)
	}
	ref, err := repo.Head()
	if err != nil {
		panic(err)
	}
	iter, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		panic(err)
	}
	commits := make([]*Commit, 0)
	_ = iter.ForEach(func(c *object.Commit) error { //nolint[errcheck]
		commits = append(
			commits,
			&Commit{
				Hash: c.Hash.String(),
				// @todo #4 Figure out how to split the commit's message into a subject
				//  and a body. I believe the separator is the first instance of the
				//  CRLFCRLF sequence.
				Subject: "don't know",
				Body:    c.Message,
			},
		)
		return nil
	})
	return func() []*Commit { return commits }
}

// Printed prints the commits to stdout.
func Printed(commits Commits) Commits {
	return func() []*Commit {
		input := commits()
		for _, c := range input {
			fmt.Printf("%s\n", &pretty{c})
		}
		return input
	}
}

// a Stringer for pretty-printing the commit.
type pretty struct {
	*Commit
}

func (p *pretty) String() string {
	return fmt.Sprintf(
		"hash: %s subject=%s body=%s",
		p.Commit.Hash, p.Commit.Subject, p.Commit.Body,
	)
}
