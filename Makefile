all: deps build

.PHONY: build
build: deps
	go build -o asso_api main.go handlers/associations.go