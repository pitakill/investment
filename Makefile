run-test:
	go test ./... -v -cover
run-development-with-air:
	air
run-development:
	go run cmd/main/main.go
