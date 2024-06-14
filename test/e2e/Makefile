# Copyright 2016 The Kubernetes Authors.
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

# Build the conformance image.
#
# Usage:
#   [ARCH=amd64] [REGISTRY="registry.k8s.io"] make (build|push) VERSION={some_released_version_of_kubernetes}

REGISTRY?=registry.k8s.io
ARCH?=amd64
OUT_DIR?=_output

LOCAL_OUTPUT_PATH=$(shell pwd)/../../../$(OUT_DIR)/local/bin/linux/$(ARCH)
DOCKERIZED_OUTPUT_PATH=$(shell pwd)/../../../$(OUT_DIR)/dockerized/bin/linux/$(ARCH)

# Keep debian releases (e.g. debian 11 == bullseye) consistent 
# between BASE_IMAGE_VERSION and DISTROLESS_IMAGE images
DISTROLESS_IMAGE?=conformance:v1.28.0
E2EIMAGE?=${REGISTRY}/${DISTROLESS_IMAGE}


all: build

build:

ifndef VERSION
    $(error VERSION is undefined)
endif
	DOCKER_CLI_EXPERIMENTAL=enabled docker buildx build \
		--platform linux/${ARCH} \
		--load \
		--pull \
		-t ${REGISTRY}/openkrusie-conformance-${ARCH}:${VERSION} \
		--build-arg E2EIMAGE=$(E2EIMAGE) \
		-f test/e2e/Dockerfile .

push: build
	docker push ${REGISTRY}/openkrusie-conformance-${ARCH}:${VERSION}
ifeq ($(ARCH),amd64)
	docker rmi ${REGISTRY}/openkrusie-conformance:${VERSION} 2>/dev/null || true
	docker tag ${REGISTRY}/openkrusie-conformance-${ARCH}:${VERSION} ${REGISTRY}/openkrusie-conformance:${VERSION}
	docker push ${REGISTRY}/openkrusie-conformance:${VERSION}
endif

.PHONY: build push all
