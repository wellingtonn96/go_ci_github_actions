# Use uma imagem base oficial do Go para compilar o binário
FROM golang:1.16 as builder

# Defina o diretório de trabalho dentro do container
WORKDIR /app

# Copie os arquivos go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixe as dependências do módulo
RUN go mod download

# Copie o código-fonte do projeto para o diretório de trabalho
COPY . .

# Compile a aplicação
RUN go build -o main .

# Use uma imagem base menor para a execução da aplicação
FROM alpine:latest

# Instale o certificado SSL
RUN apk --no-cache add ca-certificates

# Defina o diretório de trabalho
WORKDIR /root/

# Copie o binário compilado da fase anterior
COPY --from=builder /app/main .

# Exponha a porta que a aplicação irá rodar
EXPOSE 8000

# Comando para rodar a aplicação
CMD ["./main"]