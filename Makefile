BINARY := godo
COVER  := coverage.out

.PHONY: all build install run test lint coverage clean tidy

all: tidy lint test build

## build: compile the binary locally
build:
	go build -o $(BINARY) .

## install: install the CLI into Go's user binary directory
install:
	@GOBIN="$$(go env GOPATH)/bin"; \
	mkdir -p "$$GOBIN"; \
	go build -o "$$GOBIN/$(BINARY)" .

## run: run without producing a binary
run:
	go run . $(ARGS)

## test: run unit tests
test:
	go test -v ./...

## coverage: run tests and show coverage report
coverage:
	go test -coverprofile=$(COVER) ./...
	go tool cover -func=$(COVER)
	@THRESHOLD=90; \
	COVERAGE=$$(go tool cover -func=$(COVER) | \
		awk '/^total:/ {gsub("%","",$$3); print $$3}'); \
	awk -v coverage="$$COVERAGE" -v threshold="$$THRESHOLD" 'BEGIN { \
		printf "Coverage: %.1f%% (threshold: %d%%)\n", coverage, threshold; \
		if (coverage < threshold) { \
			print "Error: coverage below threshold"; \
			exit 1; \
		} \
	}'

## lint: run golangci-lint
lint:
	golangci-lint run ./...

## tidy: tidy and verify go modules
tidy:
	go mod tidy
	go mod verify

## clean: remove build artifacts
clean:
	rm -f $(BINARY) $(COVER)