FROM golang:1.15

WORKDIR /go/src/app

COPY . .

RUN go build -o math
CMD ["./math"]