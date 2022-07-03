include .env.example
-include .env

export BIN=kaguya

.PHONY: run
run:
	@echo "Running ${BIN}..."
	@cat .env.example > .env
	go run cmd/main.go

.PHONY: protoc
protoc:
	@echo "Generating protobuf files..."
	@set -e
	protoc ./hayasaka/kaguya/*.proto --go_out=. --go-grpc_out=.
	@echo "Done!"
