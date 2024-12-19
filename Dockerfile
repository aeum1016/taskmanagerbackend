FROM golang:1.23-alpine

WORKDIR /usr/local/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/ ./...

FROM ubuntu:latest

RUN apt-get update && apt-get -y install cron curl

COPY ./cron/removeSessions /etc/cron.d/removeSessions
COPY ./cron/removeTasks /etc/cron.d/removeTasks

# Give execution rights on the cron job
RUN chmod 0644 /etc/cron.d/removeSessions
RUN chmod 0644 /etc/cron.d/removeTasks

# Apply cron job
RUN crontab /etc/cron.d/removeSessions
RUN crontab /etc/cron.d/removeTasks

# Run the command on container startup
COPY --from=0 /usr/local/bin/ /usr/local/bin/

EXPOSE 8080

CMD cron && taskmanagerbackend