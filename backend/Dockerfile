FROM golang:1.19

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
# COPY go.mod go.sum ./
COPY . ./
RUN go mod download && go mod verify

EXPOSE 8000

CMD ["go", "run", "main.go"]
