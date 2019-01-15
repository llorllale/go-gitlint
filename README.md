[![Release](https://img.shields.io/github/release/llorllale/go-gitlint.svg?style=flat-square)](https://github.com/llorllale/go-gitlint/releases/latest)
[![codecov](https://codecov.io/gh/llorllale/go-gitlint/branch/master/graph/badge.svg)](https://codecov.io/gh/llorllale/go-gitlint)
[![Go Report Card](https://goreportcard.com/badge/github.com/llorllale/go-gitlint?style=flat-square)](https://goreportcard.com/report/github.com/llorllale/go-gitlint)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/llorllale/go-gitlint)

# go-gitlint
Go lint your commit messages!

## Requirements

As an architect of other (not necessarily golang) projects hosted on GitHub I need:

* Commit titles and bodies merged to the development branch conform to an arbitrary regex
* Commit titles to include the relevant issue's ID
* Lint commit msgs on varios development platforms (Windows, Linux, Mac)
* (BONUS) Pre-commit hook to validate my commit's msg
* (BONUS) Performance (because a slow pre-commit hook would render the git workflow unmanageable)
