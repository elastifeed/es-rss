# Just for building
FROM golang:1.12-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/github.com/elastifeed/es-pusher

# Enable go Modules
ENV GO111MODULE=on

# Copy source files
COPY . .

# Fetch deps dependencies
RUN go get -d -v ./...

# Build and Install executables
RUN CGO_ENABLED=0 GOOS=linux go build ./poc.go && mkdir -p /go/bin/ && mv poc /go/bin/es-rss

# Create smallest possible docker image for production
FROM alpine:3.9

LABEL maintainer="Matthias Riegler <me@xvzf.tech>"

RUN apk update && apk --no-cache add ca-certificates

COPY --from=builder /go/bin/es-rss /go/bin/es-rss

# Entrypoint for the elasticsearch gateway
ENTRYPOINT ["/go/bin/es-rss"]

EXPOSE 8080