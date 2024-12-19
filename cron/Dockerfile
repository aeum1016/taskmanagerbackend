FROM ubuntu:latest

RUN apt-get update && apt-get -y install cron curl

COPY jobs /etc/cron.d/jobs

RUN chmod 0644 /etc/cron.d/jobs

RUN crontab /etc/cron.d/jobs

CMD printenv > /etc/environment && cron -f