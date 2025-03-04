PACKAGE=github.com/tksasha/month
MODFILE=-modfile go.tool.mod

.PHONY: default
default: vet fix fmt lint test

.PHONY: vet
vet:
	@echo "go vet"
	@go vet ./...

.PHONY: fix
fix:
	@echo "go fix"
	@go fix ./...

.PHONY: fmt
fmt:
	@echo "go fmt"
	@go tool $(MODFILE) gofumpt -l -w .

.PHONY: lint
lint:
	@echo "go lint"
	@go tool $(MODFILE) golangci-lint run

.PHONY: test
test:
	@echo "go test"
	@go test ./...

.PHONY: prepare
prepare:
	@if [ ! -f go.mod ]; then go mod init $(PACKAGE) && go mod -tidy; fi
	@if [ ! -f go.tool.mod ]; then go mod init $(MODFILE) $(PACKAGE) && go mod -tidy $(MODFILE); fi
	@go get -tool $(MODFILE) mvdan.cc/gofumpt@latest
	@go get -tool $(MODFILE) github.com/golangci/golangci-lint/cmd/golangci-lint@latest
