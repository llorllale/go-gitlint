[![Release](https://img.shields.io/github/release/llorllale/go-gitlint.svg?style=flat-square)](https://github.com/llorllale/go-gitlint/releases/latest)
[![Build Status](https://travis-ci.org/llorllale/go-gitlint.svg?branch=master)](https://travis-ci.org/llorllale/go-gitlint)
[![codecov](https://codecov.io/gh/llorllale/go-gitlint/branch/master/graph/badge.svg)](https://codecov.io/gh/llorllale/go-gitlint)
[![Go Report Card](https://goreportcard.com/badge/github.com/llorllale/go-gitlint?style=flat-square)](https://goreportcard.com/report/github.com/llorllale/go-gitlint)
[![GolangCI](https://golangci.com/badges/github.com/llorllale/go-gitlint.svg)](https://golangci.com/r/github.com/llorllale/go-gitlint)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/llorllale/go-gitlint)
[![PDD status](http://www.0pdd.com/svg?name=llorllale/go-gitlint)](http://www.0pdd.com/p?name=llorllale/go-gitlint)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://raw.githubusercontent.com/llorllale/go-gitlint/master/LICENSE)

# go-gitlint
Go lint your commit messages!

## Requirements

As an architect of other (not necessarily golang) projects hosted on GitHub I need:

* Commit titles and bodies merged to the development branch conform to an arbitrary regex
* Commit titles to include the relevant issue's ID
* Lint commit msgs on varios development platforms (Windows, Linux, Mac)
* (BONUS) Pre-commit hook to validate my commit's msg
* (BONUS) Performance (because a slow pre-commit hook would render the git workflow unmanageable)
