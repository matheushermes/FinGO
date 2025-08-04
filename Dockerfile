# Etapa 1: build da aplicação
FROM golang:1.24.4-alpine AS builder

# Instala dependências básicas
RUN apk add --no-cache git

# Define diretório de trabalho
WORKDIR /app

# Copia os arquivos go.mod e go.sum e instala dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código da aplicação
COPY . .

# Compila o binário
RUN go build -o fingo ./cmd/main.go

# Etapa 2: imagem final minimalista
FROM alpine:latest

# Cria diretório de trabalho
WORKDIR /root/

# Copia o binário do estágio anterior
COPY --from=builder /app/fingo .

# Expõe a porta que sua app usa
EXPOSE 8080

# Executa o binário
CMD ["./fingo"]
