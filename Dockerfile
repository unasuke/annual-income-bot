FROM golang:1.8-alpine3.6 as builder

WORKDIR /go/src/github.com/unasuke/annual-income-bot

RUN apk --no-cache add git
RUN go get -d -v github.com/dghubble/go-twitter/twitter github.com/dghubble/oauth1
COPY post.go /go/src/github.com/unasuke/annual-income-bot/
RUN go build -o post post.go

FROM alpine:3.6
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /go/src/github.com/unasuke/annual-income-bot/post /app/post
CMD ./post
