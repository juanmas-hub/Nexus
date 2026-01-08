# --- Etapa de Compilación ---
FROM rust:1.92-alpine AS builder

# Agregamos dependencias mínimas para compilar en Alpine
RUN apk add --no-cache musl-dev gcc

WORKDIR /app

# 1. Copiamos los archivos de dependencias
COPY Cargo.toml Cargo.lock ./
# 2. COPIAMOS LA CARPETA .sqlx (Esto es lo que permite el modo offline)
COPY .sqlx ./.sqlx 

# 3. Cache de dependencias (usando un main.rs vacío)
RUN mkdir src && echo "fn main() {}" > src/main.rs && \
    cargo build --release && \
    rm -rf src

# 4. Ahora copiamos el resto del código y compilamos de verdad
COPY . .
ENV SQLX_OFFLINE=true 
RUN cargo build --release

WORKDIR /root/
# El nombre coincide con tu [package] name = "catalog-service"
COPY --from=builder /app/target/release/catalog-service .

# Puerto estándar para microservicios
EXPOSE 8080

CMD ["./catalog-service"]