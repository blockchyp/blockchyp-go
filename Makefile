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

# Executables
GO = $(MODSUPPORT) go
GOLINT = $(MODSUPPORT) golint
REVIVE = $(MODSUPPORT) revive
XUNIT = $(MODSUPPORT) go2xunit

# Default target
.PHONY: all
all: clean lint test tidy

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
	$(LINUX_BUILDENV) $(MAKE) $(BUILDDIR)/blockchyp

# Builds the windows CLI executable
.PHONY: cli-windows
cli-windows:
	$(WINDOWS_BUILDENV) $(MAKE) $(BUILDDIR)/blockchyp.exe

.PHONY: clean
clean:
	$(GO) clean -cache $(PKGS)
	rm -f $(BUILDDIR)/core
	rm -f $(BUILDDIR)/blockchyp*

$(BUILDDIR)/%: $(wildcard $(CMDDIR)/**/*)
	$(BUILDENV) $(GO) build -o $@ $<
