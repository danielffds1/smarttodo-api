.PHONY: help run build test clean docker-up docker-down migrate-up migrate-down

help: ## Mostra este help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run: ## Roda a aplicação
	go run cmd/api/main.go

build: ## Builda a aplicação
	go build -o bin/api cmd/api/main.go

test: ## Roda os testes
	go test -v ./...

clean: ## Remove arquivos de build
	rm -rf bin/

docker-up: ## Sobe o Docker Compose
	docker-compose up -d

docker-down: ## Para o Docker Compose
	docker-compose down

deps: ## Instala dependências
	go mod download
	go mod tidy

.DEFAULT_GOAL := help