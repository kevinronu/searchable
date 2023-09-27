#build stage
FROM golang:1.21.0-alpine3.18 AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY ["./server/go.mod", "./server/go.sum", "./"]
RUN go mod download
COPY ./server .
RUN go build -o server -v .

#final stage
FROM alpine:3.18
RUN apk update
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/server .
ENTRYPOINT ["./server"]
LABEL Name="email-indexer-server"