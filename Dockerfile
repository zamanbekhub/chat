FROM golang:1.21.4

WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod tidy

ADD ./ /app

RUN go build -o index

EXPOSE 8000

ENTRYPOINT ["/app/index"]
