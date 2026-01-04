# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copiamos dependencias primero para aprovechar la caché de Docker
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compilamos el binario (lo llamamos 'main' para que coincida con auth-service)
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

# Stage 2: Runtime
FROM alpine:latest
WORKDIR /root/

# Copiamos solo el binario desde la etapa anterior
COPY --from=builder /app/main .

# Exponemos el puerto que usa Render por defecto
EXPOSE 10000

CMD ["./main"]