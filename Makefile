#Caminho do entrypoint da aplicação Go
MAIN=cmd/main.go

#Nome dos containers (mesmo do docker-compose.yml)
APP_CONTAINER=fingo_app
DB_CONTAINER=fingo_postgres
REDIS_CONTAINER=fingo_redis


# Configuração do banco
DB_USER=fingo
DB_NAME=fingodb

# Docker Compose
DC=docker-compose

.PHONY: up down ps run logs restart db-shell redis-shell app-shell build lint test

## Subir os containers
up:
	$(DC) up -d --build

## Derrubar os containers
down:
	$(DC) down

## Ver status dos containers
ps:
	$(DC) p

## Rodar a aplicação principal (fora do Docker)
run:
	go run $(MAIN)

## Recompilar e reiniciar containers
restart:
	$(DC) down
	$(DC) up -d --build

## Ver logs da aplicação
logs:
	$(DC) logs -f $(APP_CONTAINER)

## Abrir terminal do banco PostgreSQL
db-shell:
	docker exec -it $(DB_CONTAINER) psql -U $(DB_USER) -d $(DB_NAME)

## Abrir terminal do Redis
redis-shell:
	docker exec -it $(REDIS_CONTAINER) redis-cli -a $$(grep REDIS_PASSWORD .env | cut -d '=' -f2)

## Abrir shell da aplicação
app-shell:
	docker exec -it $(APP_CONTAINER) sh

## Rodar linter (go vet + golangci-lint se disponível)
lint:
	go vet ./...
	@if command -v golangci-lint >/dev/null 2>&1; then golangci-lint run ./...; fi

## Rodar testes
test:
	go test ./... -v