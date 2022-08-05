FROM golang:1.18

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
# COPY go.mod go.sum ./
COPY . ./
RUN go mod download && go mod verify

CMD ["go", "run", "main.go"]
