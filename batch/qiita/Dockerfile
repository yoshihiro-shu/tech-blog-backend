FROM rust:1.76.0 AS builder
COPY . .
RUN cargo build --release

# FROM debian:buster-slim
# COPY --from=builder ./target/release/qiita ./target/release/qiita
CMD ["./target/release/qiita"]