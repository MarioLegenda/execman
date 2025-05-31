FROM ruby:3.0-slim-buster

ENV APP_DIR /app

WORKDIR $APP_DIR

RUN useradd -M -d $APP_DIR --shell /bin/bash dockeruser
RUN chown -R dockeruser:dockeruser $APP_DIR

RUN chmod -R g+r $APP_DIR
RUN chown -R dockeruser:dockeruser /usr/local/bin/ruby && chown -R dockeruser /usr/local/bin/ruby

USER dockeruser
ENV USER=dockeruser HOME=$APP_DIR