from golang:1.9

WORKDIR /go/src/clippy-bot
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["clippy-bot"]
