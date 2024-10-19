# Stage 1: Build the Go program with Cobra
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY haproxyctl.go .
RUN go mod init haproxyctl && go get -u github.com/spf13/cobra && go build -o haproxyctl haproxyctl.go

# Stage 2: Set up HAProxy and include the built Go program
FROM alpine:latest
RUN apk update && apk add --no-cache haproxy bash
COPY haproxy.cfg /usr/local/etc/haproxy/haproxy.cfg
COPY --from=builder /app/haproxyctl /usr/local/bin/haproxyctl
RUN mkdir -p /var/run && touch /var/run/haproxy.sock && chmod 600 /var/run/haproxy.sock
EXPOSE 80
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
