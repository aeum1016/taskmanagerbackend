FROM golang:1.23-alpine

WORKDIR /usr/local/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/ ./...

FROM alpine:latest

# Run the command on container startup
COPY --from=0 /usr/local/bin/ /usr/local/bin/

EXPOSE 8080

CMD taskmanagerbackend