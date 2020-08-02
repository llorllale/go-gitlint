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
# build:        builds the binary
# test:         runs tests
# coverage:     verifies test coverage
# dependencies: ensures build dependencies
# lint:         runs linters
# pdd:          runs pdd (see https://github.com/yegor256/pdd)
# checks:       runs build+test+pdd+lint
# release:      releases to GitHub (requires GitHub token)
#

build:
	@echo "Building..."
	@go build -o gitlint cmd/go-gitlint/main.go

test:
	@echo "Running unit tests..."
	@go test -count=1 -race -cover -coverprofile=coverage.txt -covermode=atomic ./... | tee cov_check.txt

coverage:
	@echo "Verifying test coverage..."
	@./check_coverage.sh

dependencies:
	@echo "Ensuring dependencies..."
	@./dependencies.sh

lint: dependencies
	@echo "Running linter..."
	@golangci-lint run

pdd: dependencies
	@echo "Scanning for puzzles..."
	@pdd -q --file=puzzles.xml

license: dependencies
	@echo "Verifying license headers..."
	@weasel

checks: build lint pdd license test coverage

release:
	@echo "Releasing..."
	@./release.sh
