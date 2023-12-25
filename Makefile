
dev: tidy
	@go run cmd/main.go


test: tidy
	@go test -cover -v ./...


tidy:
	@go mod tidy

.PHONY: dev test tidy