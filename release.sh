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

docker run --rm --privileged \
	-v $PWD:/go/src/github.com/user/repo \
	-v /var/run/docker.sock:/var/run/docker.sock \
	-w /go/src/github.com/user/repo \
	-e GITHUB_TOKEN \
	-e GO111MODULE=on \
	goreleaser/goreleaser:v0.101 release --rm-dist
