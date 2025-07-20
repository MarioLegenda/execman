FROM php:8.2-cli

ENV APP_DIR /app
#
WORKDIR $APP_DIR

RUN adduser --disabled-password --gecos "" --home /app --shell /bin/bash dockeruser
RUN chown -R dockeruser:dockeruser $APP_DIR

RUN chmod -R g+r $APP_DIR
RUN chown -R dockeruser:dockeruser /usr/local/bin/php && chown -R dockeruser /usr/local/bin/php

USER dockeruser
ENV USER=dockeruser HOME=$APP_DIR