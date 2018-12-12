# Build config
MODSUPPORT = GO111MODULE=on # TODO: Remove this when on is default
TESTFLAGS = -v -race
TESTENV =
BUILDDIR = build
REPORTDIR = $(BUILDDIR)/test-reports
PKGS = $(shell go list ./... | grep -v /vendor/)

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


.PHONY: cli-linux
cli-linux:
	go build -o build/blockchyp cmd/main.go

.PHONY: cli-windows
cli-windows:
	GOOS=windows GOARCH=386 go build -o build/blockchyp.exe cmd/main.go

GOOS=windows GOARCH=386 go

.PHONY: clean
clean:
	$(GO) clean -cache $(PKGS)
	rm -f $(BUILDDIR)/core

.PHONY: all lint test integration
