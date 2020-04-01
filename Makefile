# Version config
TAG := $(shell git tag --points-at HEAD | sort --version-sort | tail -n 1)
LASTTAG := $(or $(shell git tag -l | sort -r -V | head -n 1),0.1.0)
SNAPINFO := $(shell date +%Y%m%d%H%M%S)git$(shell git log -1 --pretty=%h)
RELEASE := $(or $(BUILD_NUMBER), 1)
VERSION := $(or $(TAG:v%=%),$(LASTTAG:v%=%))-$(or $(BUILD_NUMBER), 1)$(if $(TAG),,.$(SNAPINFO))

# Build config
TESTFLAGS := -v -race
TESTENV :=
BUILDDIR := build
DISTDIR := $(BUILDDIR)/dist
CMDDIR := cmd
REPORTDIR := $(BUILDDIR)/test-reports
PKGS := $(shell go list ./... | grep -v /vendor/)
LINUX_BUILDENV := GOOS=linux GOARCH=amd64
WIN_BUILDENV := GOOS=windows GOARCH=386
SOURCES := $(shell find . -name '*.go')
TAR_ARCHIVE := blockchyp-cli-$(VERSION).tar.gz
ZIP_ARCHIVE := blockchyp-cli-$(VERSION).zip
ICON := assets/blockchyp.ico
LDFLAGS := -s -w -extldflags '-static' \
	-X github.com/blockchyp/blockchyp-go.Version=$(VERSION)
BUILDFLAGS := -v -trimpath -ldflags "$(LDFLAGS)"

# Executables
DOCKER = docker
GO = go
GOJUNITREPORT = $(GO) run github.com/jstemmer/go-junit-report
GOLINT := $(GO) run golang.org/x/lint/golint
GOVERSIONINFO := $(GO) run github.com/josephspurrier/goversioninfo/cmd/goversioninfo
REVIVE := $(GO) run github.com/mgechev/revive
TAR := tar
ZIP := zip

# Integration test config
export BC_TEST_DELAY := 5
IMAGE := golang:1.14-buster
SCMROOT := $(shell git rev-parse --show-toplevel)
PWD := $(shell pwd)
CACHE := $(HOME)/.local/share/blockchyp/itest-cache
CONFIGFILE := $(HOME)/.config/blockchyp/sdk-itest-config.json
CACHEPATHS := $(dir $(CONFIGFILE)) $(HOME)/.cache $(HOME)/go
ifeq ($(shell uname -s), Linux)
HOSTIP = $(shell ip -4 addr show docker0 | grep -Po 'inet \K[\d.]+')
else
HOSTIP = host.docker.internal
endif

# Default target
.PHONY: all
all: clean lint tidy dist

# Build compiles the packages
.PHONY: build
build: lint dist

# Runs unit tests
.PHONY: test
test:
	mkdir -p $(REPORTDIR)/xUnit
	$(TESTENV) $(GO) test $(TESTFLAGS) $(if $(TEST), -run=$(TEST),) $(PKGS) \
		| tee -i /dev/stderr \
		| $(GOJUNITREPORT) -set-exit-code >$(REPORTDIR)/xUnit/test-report.xml

# Runs integration tests
.PHONY: integration
integration:
	$(if $(LOCALBUILD),, \
		$(foreach path,$(CACHEPATHS),mkdir -p $(CACHE)/$(path) ; ) \
		sed 's/localhost/$(HOSTIP)/' $(CONFIGFILE) >$(CACHE)/$(CONFIGFILE) ; \
		$(DOCKER) run \
		-u $(shell id -u):$(shell id -g) \
		-v $(SCMROOT):$(SCMROOT):Z \
		-v /etc/passwd:/etc/passwd:ro \
		$(foreach path,$(CACHEPATHS),-v $(CACHE)/$(path):$(path):Z) \
		-e BC_TEST_DELAY=$(BC_TEST_DELAY) \
		-e HOME=$(HOME) \
		-e GOPATH=$(HOME)/go \
		-w $(PWD) \
		--rm -it $(IMAGE)) \
	$(GO) test $(TESTFLAGS) $(if $(TEST), -run=$(TEST),) -tags=integration $(PKGS)

# Performs any tasks necessary before a release build
.PHONY: stage
stage:

# Pushes build archives to github tag
.PHONY: publish
publish:
	ghr \
		-n "BlockChyp CLI ${CIRCLE_TAG#v}" \
		-u ${CIRCLE_PROJECT_USERNAME} \
		-r ${CIRCLE_PROJECT_REPONAME} -c ${CIRCLE_SHA1} \
		-recreate \
		${CIRCLE_TAG} \
		build/dist

# Runs go lint and revive linter
.PHONY: lint
lint:
	$(GOLINT) -set_exit_status $(PKGS)
	$(REVIVE) -config revive.toml  -formatter friendly $(PKGS)

# Runs mod tidy to remove unused dependencies
.PHONY: tidy
tidy:
	$(GO) mod tidy

# Builds the linux CLI executable
.PHONY: cli-linux
cli-linux:
	GOOS=linux GOARCH=amd64 $(MAKE) $(BUILDDIR)/blockchyp/linux/amd64/blockchyp

# Builds the windows CLI executable
.PHONY: cli-windows
cli-windows:
	$(GOVERSIONINFO) -icon=$(ICON)
	GOOS=windows GOARCH=386 $(MAKE) $(BUILDDIR)/blockchyp/windows/386/blockchyp.exe
	GOOS=windows GOARCH=386 $(MAKE) $(BUILDDIR)/blockchyp/windows/386/blockchyp-headless.exe LDFLAGS="$(LDFLAGS) -H=windowsgui"
	rm *.syso

# Builds distribution archives
.PHONY: dist
dist: $(DISTDIR)/$(TAR_ARCHIVE) $(DISTDIR)/$(ZIP_ARCHIVE)

.PHONY: clean
clean:
	$(GO) clean -cache $(PKGS)
	rm -rf $(BUILDDIR)

$(DISTDIR)/$(TAR_ARCHIVE): cli-linux cli-windows
	mkdir -p $(DISTDIR)
	cd $(BUILDDIR); $(TAR) -czvf ../$(DISTDIR)/$(TAR_ARCHIVE) ./blockchyp/

$(DISTDIR)/$(ZIP_ARCHIVE): cli-linux cli-windows
	mkdir -p $(DISTDIR)
	cd $(BUILDDIR); $(ZIP) -r ../$(DISTDIR)/$(ZIP_ARCHIVE) ./blockchyp/

$(BUILDDIR)/%: $(wildcard $(CMDDIR)/**/*) $(SOURCES)
	$(BUILDENV) $(GO) build $(BUILDFLAGS) -o $@ $<
