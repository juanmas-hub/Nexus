# 1. Etapa de compilación
FROM rust:1.92-alpine AS builder

# Instalamos dependencias necesarias para compilar SQLx con soporte TLS/SSL
RUN apk add --no-cache musl-dev gcc pkgconfig openssl-dev openssl-libs-static

WORKDIR /app

# Copiamos solo los archivos de dependencias para aprovechar el caché de Docker
COPY Cargo.toml Cargo.lock ./

# --- PASO DE CACHÉ ---
# Creamos un proyecto dummy para compilar las dependencias antes que el código
RUN mkdir src && echo "fn main() {}" > src/main.rs && \
    cargo build --release && \
    rm -rf src

# --- COMPILACIÓN REAL ---
# Copiamos la carpeta .sqlx (necesaria para el modo offline) y el código fuente
COPY .sqlx/ ./.sqlx/
COPY src/ ./src/

# FORZAMOS LA RE-COMPILACIÓN:
# 1. 'touch' actualiza la fecha de main.rs para que Cargo sepa que cambió.
# 2. 'rm' elimina el binario dummy por seguridad.
RUN touch src/main.rs && rm -f target/release/catalog-service*

# Usamos los metadatos de SQLx para no necesitar conexión a la DB en build-time
ENV SQLX_OFFLINE=true
RUN cargo build --release

# 2. Etapa final (Imagen de producción ligera)
FROM alpine:latest
# Librerías necesarias para ejecutar binarios de Rust en Alpine
RUN apk add --no-cache ca-certificates libgcc libstdc++ gcompat

WORKDIR /app/

# Copiamos el binario real desde la etapa builder
COPY --from=builder /app/target/release/catalog-service .

# Exponemos el puerto (Render lo usará para el tráfico interno)
EXPOSE 8082

CMD ["./catalog-service"]