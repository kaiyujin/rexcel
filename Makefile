OUTPUT := rexcel

build: ## Build files
	go build -o $(OUTPUT) main.go

fmt: ## Format go files
	gofmt -l -s -w .

clean: ## Clean files
	rm $(OUTPUT)
	go clean

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
