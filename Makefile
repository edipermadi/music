GO		:= $(shell which go)
GOFMT		:= $(shell which gofmt)
GOPATH		:= $(shell $(GO) env GOPATH)
GOLINT		:= $(GOPATH)/bin/golint
RM		:= $(shell which rm)
TEST_OUT_DIR	:= test_out
WIKI_DIR	:= $(CURDIR)/docs
DOT_IN_FILES	:= $(wildcard docs/*.dot)
DOT_OUT_FILES	:= $(patsubst %.dot,%.svg,$(DOT_IN_FILES))

.PHONY: build
build: test docs

.PHONY: format
format:
	$(GOFMT) -w .

.PHONY: format-check
format-check:
	@test -z $($(GOFMT) -l .) || { echo "code is not formatted"; exit 1; } && { exit 0; }

.PHONY: lint-install
lint-install:
	test -e $(GOLINT) || $(GO) get -u golang.org/x/lint/golint

.PHONY: lint
lint: format-check lint-install
	$(GOLINT) ./...

.PHONY: vet
vet:
	$(GO) vet ./...

.PHONY: unit
unit: $(TEST_OUT_DIR)
	TEST_OUT_DIR=$(CURDIR)/$(TEST_OUT_DIR) $(GO) test -v ./... -coverprofile=coverage.out

.PHONY: coverage
coverage: unit
	$(GO) tool cover -func=coverage.out

.PHONY: test
test: lint vet coverage

.PHONY: xref
xref:
	mkdir -p $(WIKI_DIR)
	WIKI_DIR=$(WIKI_DIR) $(GO) run ./cmd/$@

%.svg: %.dot
	dot -Tsvg $< -o $@

.PHONY: diagrams
diagrams: $(DOT_OUT_FILES)

.PHONY: docs
docs: xref

.PHONY: db-seed
db-seed:
	$(GO) run ./cmd/db-seed > ./db/seed.sql

.PHONY: clean
clean:
	$(RM) -rf $(TEST_OUT_DIR)

$(TEST_OUT_DIR):
	mkdir -p $@

.PHONY: chordgen
chordgen:
	$(GO) run ./cmd/chordgen > chordgen-data.go
