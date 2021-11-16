FROM rust:latest as builder
WORKDIR /app
COPY Cargo.* /app/
COPY src/ /app/src/
# RUN docker run --rm -it -v "$(pwd)":/app/src ekidd/rust-musl-builder cargo build --release
RUN cargo build --release

FROM debian:latest  
ENV ROCKET_ENV=release
WORKDIR /app/
COPY Rocket.toml /app/
COPY --from=builder /app/target/release/app-rocket .
EXPOSE 8000
CMD ["/app/app-rocket"]

