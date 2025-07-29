# Caminho do entrypoint
MAIN=cmd/main.go

.PHONY: up down ps run db-shell

# Subir os containers
up:
	docker-compose up -d

# Derrubar os containers
down:
	docker-compose down

# Ver status dos containers
ps:
	docker-compose ps

# Rodar a aplicação principal
run:
	go run $(MAIN)

# Abrir terminal do banco PostgreSQL
db-shell:
	docker exec -it fingo_postgres psql -U fingo -d fingodb -c "SELECT * FROM users;"

