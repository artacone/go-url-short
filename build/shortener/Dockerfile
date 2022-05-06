FROM golang:alpine

ENV GOPATH=/
COPY ./ ./

RUN go mod download
RUN go build -o shortener cmd/server/main.go

CMD ["./shortener"]