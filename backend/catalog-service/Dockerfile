FROM rust:1.85-alpine AS builder

RUN apk add --no-cache musl-dev gcc

WORKDIR /app

COPY Cargo.toml Cargo.lock ./

RUN mkdir src && echo "fn main() {}" > src/main.rs && cargo build --release && rm -rf src

COPY . .

RUN cargo build --release

FROM alpine:latest
WORKDIR /root/

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/target/release/catalog-service .

EXPOSE 8082

CMD ["./catalog-service"]