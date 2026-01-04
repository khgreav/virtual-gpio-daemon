PROGRAM_NAME=virtual-gpio-daemon

.PHONY: all
all: build

build:
	@go build -o $(PROGRAM_NAME) .

clean:
	@rm -f $(PROGRAM_NAME)

update:
	@go get -u ./...
	@go mod tidy