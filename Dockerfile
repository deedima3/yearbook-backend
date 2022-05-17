FROM golang:1.17-alpine3.15

WORKDIR /build

COPY ./go.mod ./go.sum /build/
RUN go mod download

COPY . .

RUN go build ./cmd/main/main.go

EXPOSE 8080

CMD ["./main"]
