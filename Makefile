.PHONY: test build run vet clean release

ifndef VERBOSE
.SILENT:
endif

APP=./cmd/ynot
OUTPUT_FOLDER=./bin
BINARY_NAME=$(OUTPUT_FOLDER)/ynot

SOURCE=$(shell find ./ -type f -name '*.go')
MODULES=go.mod go.sum

build: $(BINARY_NAME) ;

$(BINARY_NAME): $(SOURCE) $(MODULES)
	mkdir -p $(OUTPUT_FOLDER)
	go build -o $(BINARY_NAME) $(APP)

release:
	mkdir -p $(OUTPUT_FOLDER)
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux $(APP)

test:
	go test -v ./...

run: build
	$(BINARY_NAME)

vet:
	go vet $(APP)

clean:
	go clean
	rm -rf $(OUTPUT_FOLDER)
