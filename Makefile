# Build config
MODSUPPORT := GO111MODULE=on # TODO: Remove this when on is default
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
HASH := $(shell git log -1 --pretty=%h)
TAG := $(shell git tag --points-at HEAD | sort --version-sort | tail -n 1)
TAR_ARCHIVE := blockchyp-cli-$(or $(TAG:v%=%), $(HASH)).tar.gz
ZIP_ARCHIVE := blockchyp-cli-$(or $(TAG:v%=%), $(HASH)).zip
ICON := assets/blockchyp.ico
BUILDFLAGS := -ldflags "-X main.compileTimeVersion=$(or $(TAG:v%=%), $(HASH))"

# Executables
GO := $(MODSUPPORT) go
GOJUNITREPORT := $(GO) run github.com/jstemmer/go-junit-report
GOLINT := $(GO) run golang.org/x/lint/golint
GOVERSIONINFO := $(GO) run github.com/josephspurrier/goversioninfo/cmd/goversioninfo
REVIVE := $(MODSUPPORT) $(GO) run github.com/mgechev/revive
ZIP := zip
TAR := tar

# Default target
.PHONY: all
all: clean lint tidy dist

# Runs go lint and revive linter
.PHONY: lint
lint:
	$(GOLINT) -set_exit_status $(PKGS)
	$(REVIVE) -config revive.toml  -formatter friendly $(PKGS)

# Runs unit tests
.PHONY: test
test:
	mkdir -p $(REPORTDIR)/xUnit
	$(TESTENV) $(GO) test $(TESTFLAGS) $(if $(TEST), -run=$(TEST),) $(PKGS) \
		| tee -i /dev/stderr \
		| $(GOJUNITREPORT) -set-exit-code >$(REPORTDIR)/xUnit/test-report.xml

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
