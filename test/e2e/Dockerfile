# Copyright 2018 The Kubernetes Authors.
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

ARG E2EIMAGE

FROM --platform=${TARGETPLATFORM} hub.byted.org/compile/bytedance.golang.compile_tango_1_21_buster:b2ff06836c4c599928d691bf82350e76 as builder

ADD . /go/src/github.com/openkruise/kruise
WORKDIR /go/src/github.com/openkruise/kruise

#ARG CODEBASE
#ENV GIT_SSH_COMMAND $CODEBASE
#RUN apt-get update -y \
#    && apt-get install -y ssh sshpass git \
#    && rm -rf /var/lib/apt/list/*
#COPY gitconfig /etc/gitconfig

RUN make ginkgo && CGO_ENABLED=0 bin/ginkgo build /go/src/github.com/openkruise/kruise/test/e2e

FROM ${E2EIMAGE}


COPY --from=builder /go/src/github.com/openkruise/kruise/test/e2e/e2e.test /usr/local/bin/
COPY --from=builder /go/src/github.com/openkruise/kruise/test/e2e/testing-manifests /
