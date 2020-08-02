#!/bin/bash
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

ensureRuby2xInstalled() {
  result=$(ruby --version)

  if [ -z "$result" ]; then
    echo "Please install Ruby 2.x!"
    exit 1
  fi

  version=$(echo $result | cut -d " " -f 2)
  matched=$([[ $version =~ 2\..* ]] && echo matched)

  if [ -z "$matched" ]; then
    echo "You have Ruby $version installed - please install a 2.x version."
    exit 1
  fi
}

installPDD() {
  installed=$(pdd -h && echo yes)

  if [ -z "$installed" ]; then
    echo "PDD not found. Installing..."
    gem install pdd
  fi
}

installGolangCILint() {
  VERSION=1.29.0
  installed=$(golangci-lint --version)

  if [ -z "$installed" ]; then
    echo "golangci-lint not found. Installing..."
    curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(go env GOPATH)/bin v$VERSION

    return 0
  fi

  version=$(echo $installed | cut -d " " -f 4)
  matched=$([[ $version =~ $VERSION ]] && echo matched)

  if [ -z "$matched" ]; then
    echo "You have golangci-lint $version installed. Please install version v$VERSION."
    exit 1
  fi
}

installWeasel() {
  installed=$(weasel -v && echo yes)

  if [ -z "$installed" ]; then
    echo "weasel not found. Installing..."
    (cd $(mktemp -d) && go get github.com/comcast/weasel)
  fi
}

main() {
  ensureRuby2xInstalled
  installPDD
  installGolangCILint
  installWeasel
}

main
