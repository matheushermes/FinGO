#Caminho do entrypoint
MAIN=cmd/main.go

#Subir os containers
up:
	docker-compose up -d

#Derrubar os containers
down:
	docker-compose down

#Ver status dos containers
ps:
	docker-compose ps

#Rodar a aplicação principal
run:
	go run $(MAIN)