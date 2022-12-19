.PHONY: all
all: tidy fmt vet golangci-lint test-cover

.PHONY: tidy
tidy:
	@echo "Executing go mod tidy..."
	@go mod tidy

.PHONY: fmt
fmt:
	@echo "Executing go fmt.."
	@go fmt ./...

.PHONY: vet
vet:
	@echo "Executing go vet..."
	@go vet ./...

.PHONY: golangci-lint
golangci-lint:
	@echo "Executing golangci-lint..."
	@golangci-lint run

.PHONY: test-cover
test-cover:
	@echo "Running tests and generating report"
	@go test -v -coverpkg=./... -coverprofile=profile.cov ./...
	@go tool cover -func profile.cov
	@rm profile.cov