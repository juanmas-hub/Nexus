FROM rust:1.92-alpine AS builder

RUN apk add --no-cache musl-dev gcc

WORKDIR /app

COPY Cargo.toml Cargo.lock ./

RUN mkdir src && echo "fn main() {}" > src/main.rs && cargo build --release && rm -rf src

COPY . .
RUN cargo build --release

FROM alpine:latest
RUN apk add --no-cache ca-certificates

WORKDIR /root/
COPY --from=builder /app/target/release/catalog-service .

EXPOSE 8080

CMD ["./catalog-service"]