FROM golang:1.20 as dev
# LABEL maintainer "draft-backend"

ENV ROOT=/usr/src/app
WORKDIR ${ROOT}

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY . ./
RUN go mod download && go mod verify
EXPOSE 8000

CMD ["go", "run", "./cmd/main.go"]
