FROM golang:1.10.1-alpine AS builder
RUN apk add curl git --update --no-cache; \
    curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64; \
    chmod +x /usr/local/bin/dep

RUN mkdir -p /go/src/github.com/tormath1/tweetogo
WORKDIR /go/src/github.com/tormath1/tweetogo
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only
COPY . ./
RUN GOOS=linux GOARCH=amd64 go build -o main

FROM alpine:3.6
COPY --from=builder /go/src/github.com/tormath1/tweetogo/main /opt/
RUN apk add ca-certificates --update; \
    chmod u+x /opt/main && mkdir -p /opt/secrets
CMD ["./opt/main"]