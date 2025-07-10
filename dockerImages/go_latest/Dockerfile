FROM golang:latest

ENV APP_DIR /app

WORKDIR $APP_DIR

RUN useradd -d /app -s /bin/bash dockeruser
RUN chmod -R u+rw /app
RUN chown -R dockeruser:dockeruser /app && chown dockeruser:dockeruser /usr/local/go/bin && chown dockeruser:dockeruser /usr/local/go/bin

USER dockeruser

ENV HOME=$APP_DIR