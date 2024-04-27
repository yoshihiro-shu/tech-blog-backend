# syntax=docker/dockerfile:1
FROM golang:1.22-alpine3.18 as builder

WORKDIR /usr/src/app

# tzdata パッケージのインストール
RUN apk add --no-cache tzdata
RUN apk add --no-cache ca-certificates openssl
# RUN apt-get install -y ca-certificates openssl

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download && go mod verify
COPY ./src .

RUN CGO_ENABLED=0 go build -o binary ./cmd/main.go

FROM scratch

WORKDIR /usr/src/app

# scratch doesn't have timezone.
COPY --from=builder /usr/share/zoneinfo/Asia/Tokyo /usr/share/zoneinfo/Asia/Tokyo

COPY --from=builder /usr/src/app/binary /usr/src/app/binary
COPY ./src/configs.yaml .

EXPOSE 8080
ENTRYPOINT ["./binary"]
