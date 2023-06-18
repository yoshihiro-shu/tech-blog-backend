# syntax=docker/dockerfile:1
FROM golang:1.20 as builder

WORKDIR /usr/src/app

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
CMD ["./binary"]
