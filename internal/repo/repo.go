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

// Package repo is the API for fetching repos.
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
