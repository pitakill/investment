.PHONY: build
run-test:
	go test ./... -v -cover
run-development-with-air:
	air
run-development:
	go run cmd/main/main.go
build:
	go build -o ./main ./cmd/main
