[![Release](https://img.shields.io/github/release/llorllale/go-gitlint.svg?style=flat-square)](https://github.com/llorllale/go-gitlint/releases/latest)
[![Build Status](https://travis-ci.org/llorllale/go-gitlint.svg?branch=master)](https://travis-ci.org/llorllale/go-gitlint)
[![codecov](https://codecov.io/gh/llorllale/go-gitlint/branch/master/graph/badge.svg)](https://codecov.io/gh/llorllale/go-gitlint)
[![Go Report Card](https://goreportcard.com/badge/github.com/llorllale/go-gitlint?style=flat-square)](https://goreportcard.com/report/github.com/llorllale/go-gitlint)
[![codebeat](https://codebeat.co/badges/16512202-9758-4e2e-b0b8-4121724680b8)](https://codebeat.co/projects/github-com-llorllale-go-gitlint-master)
[![GolangCI](https://golangci.com/badges/github.com/llorllale/go-gitlint.svg)](https://golangci.com/r/github.com/llorllale/go-gitlint)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg)](https://github.com/goreleaser)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/llorllale/go-gitlint)
[![PDD status](http://www.0pdd.com/svg?name=llorllale/go-gitlint)](http://www.0pdd.com/p?name=llorllale/go-gitlint)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://raw.githubusercontent.com/llorllale/go-gitlint/master/LICENSE)

# go-gitlint
Lint your git!

## Usage
```
$ ./gitlint --help
usage: gitlint [<flags>]

Flags:
  --help                    Show context-sensitive help (also try --help-long and --help-man).
  --path="."                Path to the git repo (default: ".").
  --subject-regex=".*"      Commit subject line must conform to this regular expression (default: ".*").
  --subject-len=2147483646  Commit subject line cannot exceed this length (default: math.MaxInt32 - 1).
  --body-regex=".*"         Commit message body must conform to this regular expression (default: ".*").
  --since="1970-01-01"      A date in "yyyy-MM-dd" format starting from which commits will be analyzed (default: "1970-01-01")
```
Additionally, it will look for configurations in a file `.gitlint` in the current directory if it exists. This file's format is just the same command line flags but each on a separate line. *Flags passed through the command line take precedence.*

## Integrate to your CI

Use the [`download-gitlint.sh`](https://raw.githubusercontent.com/llorllale/go-gitlint/master/download-gitlint.sh) script:

`curl -sfL https://raw.githubusercontent.com/llorllale/go-gitlint/master/download-gitlint.sh | sh -s 1.0.0`

Specifying the version is optional; if you just want the latest, omit the `-s <version>` part.

In both cases the correct version for your platform will be downloaded and installed at `$GOPATH/bin`.

## Motivation

- [X] Validate format of commit message subject and body
- [X] Lint commit msgs on varios development platforms (Windows, Linux, Mac)
- [X] Configuration from file with cli args taking precedence
- [ ] [`commit-msg`](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks) hook to validate my commit's msg
- [X] Performance (because a slow pre-commit hook would render the git workflow unmanageable)
- [X] My first Go project :)

## Contributing
Fork this repo, make sure `make checks` works, **and then** open a PR.

## Build dependencies
To run `make checks` you will need:

* Go `1.11.x`
* Ruby 2.x (for `pdd`)
* [pdd](https://github.com/yegor256/pdd) (a ruby gem - `gem install pdd`)
* [golangci-lint](https://github.com/golangci/golangci-lint) v1.14.0 (expected to be in the `./bin` folder)
* [weasel](https://github.com/comcast/weasel) (`go get github.com/comcast/weasel`)

