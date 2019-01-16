# Build config
MODSUPPORT = GO111MODULE=on # TODO: Remove this when on is default
TESTFLAGS = -v -race
TESTENV =
BUILDDIR = build
CMDDIR = cmd
REPORTDIR = $(BUILDDIR)/test-reports
PKGS = $(shell go list ./... | grep -v /vendor/)
LINUX_BUILDENV = GOOS=linux GOARCH=amd64
WIN_BUILDENV = GOOS=windows GOARCH=386
SOURCES = $(shell find . -name '*.go')
HASH = $(shell git log -1 --pretty=%h)
TAG = $(shell git tag --points-at HEAD | sort --version-sort | tail -n 1)
TAR_ARCHIVE = blockchyp-cli-$(or $(TAG:v%=%), $(HASH)).tar.gz
ZIP_ARCHIVE = blockchyp-cli-$(or $(TAG:v%=%), $(HASH)).zip
ICON = assets/blockchyp.ico

# Executables
GO = $(MODSUPPORT) go
GOLINT = $(GO) run github.com/golang/lint/golint
GOVERSIONINFO = $(GO) run github.com/josephspurrier/goversioninfo/cmd/goversioninfo
REVIVE = $(MODSUPPORT) $(GO) run github.com/mgechev/revive
XUNIT = $(GO) run github.com/tebeka/go2xunit
ZIP = zip
TAR = tar

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
	@mkdir -p $(REPORTDIR)
	$(TESTENV) $(GO) test $(TESTFLAGS) $(if $(TEST), -run=$(TEST),) $(PKGS) \
		| tee -i /dev/stderr \
		| $(XUNIT) -fail -output $(REPORTDIR)/unit.xml

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
	$(GOVERSIONINFO) -icon=$(ICON) -manifest=blockchyp.exe.manifest
	GOOS=windows GOARCH=386 $(MAKE) $(BUILDDIR)/blockchyp/windows/386/blockchyp.exe
	rm *.syso

# Builds distribution archives
.PHONY: dist
dist: $(BUILDDIR)/$(TAR_ARCHIVE) $(BUILDDIR)/$(ZIP_ARCHIVE)

.PHONY: clean
clean:
	$(GO) clean -cache $(PKGS)
	rm -f $(BUILDDIR)/core
	rm -f $(BUILDDIR)/blockchyp*
	rm -f $(BUILDDIR)/*.tar.gz
	rm -f $(BUILDDIR)/*.zip

$(BUILDDIR)/$(TAR_ARCHIVE): cli-linux cli-windows
	cd $(BUILDDIR); $(TAR) -czvf $(TAR_ARCHIVE) ./blockchyp/

$(BUILDDIR)/$(ZIP_ARCHIVE): cli-linux cli-windows
	cd $(BUILDDIR); $(ZIP) -r $(ZIP_ARCHIVE) ./blockchyp/

$(BUILDDIR)/%: $(wildcard $(CMDDIR)/**/*) $(SOURCES)
	$(BUILDENV) $(GO) build -o $@ $<
