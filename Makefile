#   Copyright 2020 Pokémon GO Nancy
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.

VERSION=0.1.0-dev

DOCKER=docker
GO=go

include .makerc

define docker_build
	@echo "Building $(1)..."
	@$(DOCKER) build \
		--rm \
		--label "project=cgear-go" \
		--build-arg "VERSION=$(VERSION)" \
		--tag  github.com/cgear-go/bot/$(1):$(VERSION) \
		--file docker/$(1)/Dockerfile . > /dev/null
endef

all: run

.PHONY: containers run test

run:
	@$(GO) run .

containers: test
	$(call docker_build,cgear-go-bot)

test:
	@echo "Running tests..."
	@$(GO) test ./... | grep -v "no test files"

test-deps:
	@go install github.com/golang/mock/mockgen@v1.5.0
