# 1. Etapa de compilación
FROM rust:1.92-alpine AS builder

# Instalamos dependencias de compilación
RUN apk add --no-cache musl-dev gcc

WORKDIR /app

# Copiamos archivos de configuración y la carpeta de SQLx
COPY Cargo.toml Cargo.lock ./
COPY .sqlx/ ./.sqlx/

# Paso de caché (compila dependencias)
RUN mkdir src && echo "fn main() {}" > src/main.rs && \
    cargo build --release && \
    rm -rf src

# Copiamos el código fuente y compilamos el binario final
COPY src/ ./src/
ENV SQLX_OFFLINE=true
RUN cargo build --release

# 2. Etapa final (Imagen de producción)
FROM alpine:latest
RUN apk add --no-cache ca-certificates

WORKDIR /root/
# Copiamos el binario desde la etapa builder
COPY --from=builder /app/target/release/catalog-service .

# Puerto estándar para microservicios
EXPOSE 8080

CMD ["./catalog-service"]