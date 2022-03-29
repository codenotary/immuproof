SHELL=/bin/bash -o pipefail

VERSION=0.0.5

GO ?= go

GIT_REV := $(shell git rev-parse HEAD 2> /dev/null || true)
GIT_COMMIT := $(if $(shell git status --porcelain --untracked-files=no),${GIT_REV}-dirty,${GIT_REV})
GIT_BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null)

LDFLAGS := -s -X github.com/codenotary/immuproof/meta.version=v${VERSION} \
			  -X github.com/codenotary/immuproof/meta.gitCommit=${GIT_COMMIT} \
			  -X github.com/codenotary/immuproof/meta.gitBranch=${GIT_BRANCH}
LDFLAGS_STATIC := ${LDFLAGS} \
				  -X github.com/codenotary/immuproof/meta.static=static \
				  -extldflags "-static"
TEST_FLAGS ?= -v -race

.PHONY: immuproof
immuproof:
	$(GO) build -ldflags '${LDFLAGS} -X github.com/codenotary/immuproof/meta.version=v${VERSION}-dev' -o immuproof ./main.go

.PHONY: immuproof-release
immuproof-release:
	$(GO) build -ldflags '${LDFLAGS}' -o immuproof ./main.go

.PHONY: test
test:
	$(GO) vet ./...
	$(GO) test ${TEST_FLAGS} ./...

.PHONY: static
static:
	CGO_ENABLED=0 $(GO) build -ldflags '${LDFLAGS_STATIC}' -o immuproof ./main.go

.PHONY: clean/dist
clean/dist:
	rm -Rf ./dist

.PHONY: clean
clean: clean/dist
	rm -f ./vcn

.PHONY: CHANGELOG.md
CHANGELOG.md:
	git-chglog -o CHANGELOG.md

.PHONY: CHANGELOG.md.next-tag
CHANGELOG.md.next-tag:
	git-chglog -o CHANGELOG.md --next-tag v${VERSION}
