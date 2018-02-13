FROM golang:1.9

WORKDIR /go/src/clippy-bot
COPY . .

RUN go install -v ./...

CMD ["clippy-bot"]
