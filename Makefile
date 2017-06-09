GOCC := go

# Program version
VERSION ?= $(shell git describe --always --tags)

# Binary name for bintray
BIN_NAME=git-user

# Project owner for bintray
OWNER=gesquive

# Project name for bintray
PROJECT_NAME=git-user

# Project url used for builds
# examples: github.com, bitbucket.org
REPO_HOST_URL=github.com

# Grab the current commit
GIT_COMMIT=$(shell git rev-parse HEAD)

# Check if there are uncommited changes
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)

# Use a local vendor directory for any dependencies; comment this out to
# use the global GOPATH instead
# GOPATH=$(PWD)

INSTALL_PATH=$(GOPATH)/src/${REPO_HOST_URL}/${OWNER}/${PROJECT_NAME}

BUILD_DIR=dist

default: test build

.PHONY: help
help:
	@echo 'Management commands for ${PROJECT_NAME}:'
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
	 awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Compile the project
	@echo "building ${OWNER} ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	${GOCC} build -ldflags "-X main.version=${VERSION} -X main.dirty=${GIT_DIRTY}" -o ${BIN_NAME}

.PHONY: install
install: build ## Install the binaries on this computer
	install -d ${DESTDIR}/usr/local/bin/
	install -m 755 ./${BIN_NAME} ${DESTDIR}/usr/local/bin/${BIN_NAME}
	install -m 644 ./man/*.1 ${DESTDIR}/usr/local/share/man/man1/

.PHONY: deps
deps: glide ## Download project dependencies
	glide install

.PHONY: test
test: glide ## Run golang tests
	${GOCC} test $(shell glide novendor)

.PHONY: bench
bench: glide ## Run golang benchmarks
	${GOCC} test -benchmem -bench=. $(shell glide novendor)

.PHONY: clean
clean: ## Clean the directory tree of artifacts
	${GOCC} clean
	rm -f ./${BIN_NAME}.test
	rm -f ./${BIN_NAME}
	rm -rf ./dist
	rm -f ./genman/genman

.PHONY: build-all
build-all: gox
	gox -verbose \
	-ldflags "-X main.version=${VERSION} -X main.dirty=${GIT_DIRTY}" \
	-os="linux darwin windows" \
	-arch="amd64 386" \
	-output="dist/{{.OS}}-{{.Arch}}/{{.Dir}}" .

.PHONY: dist
dist: build-all ## Cross compile the full distribution
	pkg/dist.sh "linux-386" "${PROJECT_NAME}-${VERSION}-linux-x32"
	pkg/dist.sh "linux-amd64" "${PROJECT_NAME}-${VERSION}-linux-x64"
	pkg/dist.sh "darwin-386" "${PROJECT_NAME}-${VERSION}-osx-x32"
	pkg/dist.sh "darwin-amd64" "${PROJECT_NAME}-${VERSION}-osx-x64"
	pkg/dist.sh "windows-386" "${PROJECT_NAME}-${VERSION}-windows-x32"
	pkg/dist.sh "windows-amd64" "${PROJECT_NAME}-${VERSION}-windows-x64"
	cd dist && find . -mindepth 1 -maxdepth 1 -type d -exec rm -rf "{}" \;

.PHONY: docs
docs: ## Compile the documentation
	cd genman && ${GOCC} build -ldflags "-X main.version=${VERSION}"
	mkdir -p man
	genman/genman ./man

.PHONY: fmt
fmt: ## Reformat the source tree with gofmt
	find . -name '*.go' -not -path './.vendor/*' -exec gofmt -w=true {} ';'

.PHONY: link
link: $(INSTALL_PATH) ## Symlink this project into the GOPATH
$(INSTALL_PATH):
	@mkdir -p `dirname $(INSTALL_PATH)`
	@ln -s $(PWD) $(INSTALL_PATH) >/dev/null 2>&1

.PHONY: glide
glide:
	@command -v glide >/dev/null 2>&1 || \
	echo "Installing glide" && ${GOCC} get -u github.com/Masterminds/glide

.PHONY: gox
gox:
	@command -v gox >/dev/null 2>&1 || \
	echo "Installing gox" && ${GOCC} get -u github.com/mitchellh/gox
