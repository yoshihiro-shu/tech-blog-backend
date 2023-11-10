# syntax=docker/dockerfile:1
FROM golang:1.21.4-alpine3.18 as builder

WORKDIR /usr/src/app

# tzdata パッケージのインストール
RUN apt-get update && apt-get install -y tzdata

COPY ./backend/go.mod ./backend/go.sum ./
RUN go mod download && go mod verify
COPY ./backend .

RUN CGO_ENABLED=0 go build -o binary ./cmd/main.go

FROM scratch

WORKDIR /usr/src/app

# scratch doesn't have timezone.
COPY --from=builder /usr/share/zoneinfo/Asia/Tokyo /usr/share/zoneinfo/Asia/Tokyo

COPY --from=builder /usr/src/app/binary /usr/src/app/binary
COPY ./backend/configs.yaml .

EXPOSE 8080
ENTRYPOINT ["./binary"]
