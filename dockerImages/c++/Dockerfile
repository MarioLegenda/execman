FROM mhart/alpine-node:14.1.0

ENV APP_DIR /app

RUN apk add g++

WORKDIR $APP_DIR

RUN adduser -D -h $APP_DIR -s /bin/bash dockeruser
RUN chown -R dockeruser:dockeruser $APP_DIR

RUN chmod -R g+r $APP_DIR
RUN chown -R dockeruser:dockeruser /usr/bin/g++

USER dockeruser
ENV USER=dockeruser HOME=$APP_DIR