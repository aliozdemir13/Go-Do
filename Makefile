BINARY  := go-do
COVER   := coverage.out

.PHONY: all build run test lint coverage clean tidy

all: tidy lint test build

## build: compile the binary
build:
	go build -o $(BINARY) .

## run: run without producing a binary
run:
	go run . $(ARGS)

## test: run unit tests
test:
	go test -v ./...

## coverage: run tests and show coverage report (fails below 90%)
coverage:
	go test -coverprofile=$(COVER) ./...
	go tool cover -func=$(COVER)
	@THRESHOLD=90; \
	COVERAGE=$$(go tool cover -func=$(COVER) | grep total | awk '{print $$3}' | sed 's/%//'); \
	echo "Coverage: $$COVERAGE%  (threshold: $$THRESHOLD%)"; \
	VALID=$$(echo "$$COVERAGE >= $$THRESHOLD" | bc -l); \
	if [ "$$VALID" -eq 0 ]; then \
		echo "Error: coverage below $$THRESHOLD%"; exit 1; \
	fi

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
