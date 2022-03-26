# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:latest

ENV GO114MODULE=on
# Copy the local package files to the container's workspace.

WORKDIR /app
ADD . /app

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
RUN chmod +x /app

# Document that the service listens on port 8085.
EXPOSE 8080
ENTRYPOINT ["/app/users-report"]