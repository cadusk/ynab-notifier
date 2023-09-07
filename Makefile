.PHONY: test build run vet clean

ifndef VERBOSE
.SILENT:
endif

OUTPUT_FOLDER=./bin
BINARY_NAME=$(OUTPUT_FOLDER)/ynot

SOURCE=$(wildcard *.go)
MODULES=go.mod go.sum

build: $(BINARY_NAME) ;

$(BINARY_NAME): $(SOURCE) $(MODULES)
	mkdir -p $(OUTPUT_FOLDER)
	go build -o $(BINARY_NAME) $(SOURCE)

test:
	go test -v $(SOURCE)

run: $(BINARY_NAME)
	sh -c "source .env && $(BINARY_NAME)"

vet:
	go vet

clean:
	go clean
	rm -rf $(OUTPUT_FOLDER)
