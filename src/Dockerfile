FROM golang:1.22.2-alpine3.18 as dev

ENV ROOT=/usr/src/app
WORKDIR ${ROOT}

RUN go install github.com/cosmtrek/air@latest

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY . ./
RUN go mod download && go mod verify

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]
