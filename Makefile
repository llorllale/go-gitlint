#
# Copyright 2019 George Aristy
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

#
# Supported Targets:
#
# build:   builds the binary
# test:    runs tests
# lint:    runs linters
# checks:  runs build+test+lint
# release: releases to GitHub (requires GitHub token)
#

build:
	go build -o gitlint cmd/go-gitlint/main.go

test:
	go test -count=1 -race -cover -coverprofile=coverage.txt -covermode=atomic ./...

lint:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.14.0
	./bin/golangci-lint run --enable-all

checks: build lint test

release:
	./release.sh
