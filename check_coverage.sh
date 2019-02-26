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

let THRESHOLD=80
IGNORE=github.com/llorllale/go-gitlint/cmd/go-gitlint

let exit_code=0

while read line; do
	pkg=$(echo $line | sed 's/\s\+/ /g' | sed 's/%//' | cut -d ' ' -f 2)
	if [[ "$(echo $line | grep 'no test files')" != "" && "$pkg" != "$IGNORE" ]]; then
		echo "No coverage for package [$pkg]"
		let exit_code++
	elif [[ "$(echo $line | grep coverage)" != "" ]]; then
		cov=$(echo $line | sed 's/\s\+/ /g' | sed 's/%//' | cut -d ' ' -f 5)
		if [ 1 -eq $(echo "$THRESHOLD > $cov" | bc) ]; then
			echo "Coverage [$cov] for package [$pkg] is below threshold [$THRESHOLD]"
			let exit_code++
		fi
	fi
done < ./cov_check.txt

exit $exit_code
