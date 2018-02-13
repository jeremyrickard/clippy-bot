FROM golang:1.9

WORKDIR /go/src/clippy-bot
COPY . .

RUN go build -o clippy-bot .

CMD ["./clippy-bot"]
