#!/bin/bash
#
# Copyright 2026 George Aristy
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

installWeasel() {
  installed=$(weasel -v && echo yes)

  if [ -z "$installed" ]; then
    echo "weasel not found. Installing..."
    (cd $(mktemp -d) && go install github.com/comcast/weasel@latest)
  fi
}

main() {
  ensureRuby2xInstalled
  installPDD
  installWeasel
}

main
