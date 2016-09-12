#
#  Makefile
#
#  The kickoff point for all project management commands.
#

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

INSTALL_PATH=$(GOPATH)/${REPO_HOST_URL}/${OWNER}/${PROJECT_NAME}

FIND_DIST:=find * -type d -exec

BUILD_DIR=dist

default: build

help:
	@echo 'Management commands for git-user:'
	@echo
	@echo 'Usage:'
	@echo '    make build    Compile the project'
	@echo '    make link     Symlink this project into the GOPATH'
	@echo '    make test     Run tests on a compiled project'
	@echo '    make install  Install binary'
	@echo '    make depends  Download dependencies'
	@echo '    make docs     Creates documentation'
	@echo '    make fmt      Reformat the source tree with gofmt'
	@echo '    make clean    Clean the directory tree'
	@echo '    make dist     Cross compile the full distribution'
	@echo

build:
	@echo "building ${OWNER} ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	${GOCC} build -ldflags "-X main.version=${VERSION} -X main.dirty=${GIT_DIRTY}" -o ${BIN_NAME}

install: build
	install -d ${DESTDIR}/usr/local/bin/
	install -m 755 ./${BIN_NAME} ${DESTDIR}/usr/local/bin/${BIN_NAME}
	install -m 644 ./man/*.1 ${DESTDIR}/usr/local/share/man/man1/

depends:
	${GOCC} get -u github.com/Masterminds/glide
	glide install

test:
	${GOCC} test ./...

clean:
	${GOCC} clean
	rm -f ./${BIN_NAME}.test
	rm -f ./${BIN_NAME}
	rm -rf ./dist
	rm -f ./genman/genman

bootstrap-dist:
	${GOCC} get -u github.com/mitchellh/gox

build-all: bootstrap-dist
	gox -verbose \
	-ldflags "-X main.version=${VERSION} -X main.dirty=${GIT_DIRTY}" \
	-os="linux darwin windows " \
	-arch="amd64 386" \
	-output="dist/{{.OS}}-{{.Arch}}/{{.Dir}}" .

dist: build-all
	install/dist.sh "linux-386" "${PROJECT_NAME}-${VERSION}-linux-x32"
	install/dist.sh "linux-amd64" "${PROJECT_NAME}-${VERSION}-linux-x64"
	install/dist.sh "darwin-386" "${PROJECT_NAME}-${VERSION}-osx-x32"
	install/dist.sh "darwin-amd64" "${PROJECT_NAME}-${VERSION}-osx-x64"
	install/dist.sh "windows-386" "${PROJECT_NAME}-${VERSION}-windows-x32"
	install/dist.sh "windows-amd64" "${PROJECT_NAME}-${VERSION}-windows-x64"

docs:
	cd genman && ${GOCC} build -ldflags "-X main.version=${VERSION}"
	mkdir -p man
	genman/genman ./man

fmt:
	find . -name '*.go' -not -path './.vendor/*' -exec gofmt -w=true {} ';'

link:
	# relink into the go path
	if [ ! $(INSTALL_PATH) -ef . ]; then \
		mkdir -p `dirname $(INSTALL_PATH)`; \
		ln -s $(PWD) $(INSTALL_PATH); \
	fi


.PHONY: build help test install depends clean bootstrap-dist build-all dist docs fmt link
