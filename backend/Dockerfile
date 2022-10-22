FROM golang:1.19 as dev
# LABEL maintainer "draft-backend"

ENV ROOT=/usr/src/app
WORKDIR ${ROOT}

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY . ./
RUN go mod download && go mod verify
EXPOSE 8000

CMD ["go", "run", "./cmd/main.go"]


# FROM golang:1.19-alpine as builder

# ENV ROOT=/usr/src/app
# WORKDIR ${ROOT}

# RUN apk update && apk add git
# COPY go.mod go.sum ./
# RUN go mod download

# COPY . ${ROOT}
# RUN CGO_ENABLED=0 GOOS=linux go build -o ${ROOT}/binary ${ROOT}/cmd/main.go


# FROM ubuntu:20.04 as prd
# # FROM scratch as prd

# ENV ROOT=/usr/src/app
# ENV CGO_ENABLED 0
# WORKDIR ${ROOT}

# COPY --from=builder ${ROOT}/binary ${ROOT}

# EXPOSE 8000

# CMD ["/usr/src/app/binary"]