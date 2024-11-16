FROM golang:1.22.8-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o ./bin/runner ./server/main.go

FROM alpine:latest

RUN apk update && apk add bash
RUN apk add --no-cache bash

WORKDIR /app

COPY --from=builder /app/bin/runner .
COPY --from=builder /app/docker/wait-for-it.sh .

ENTRYPOINT [ "/bin/bash", "-c" ]
EXPOSE 8080
CMD ["./runner"]