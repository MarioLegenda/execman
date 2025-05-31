# EXAMPLE USAGE:  docker run --rm codewars/python-runner run -l python -c "print 1+1 "
#                 docker run --rm codewars/python-runner run -l python3 -c "print(1+1)"

# Pull base image.
FROM rust:slim-buster

ENV APP_DIR /app

WORKDIR $APP_DIR
RUN mkdir -p $APP_DIR/src

RUN chmod g+r $APP_DIR
RUN useradd -d $APP_DIR -s /bin/bash dockeruser
RUN chown -R dockeruser:dockeruser $APP_DIR && chown dockeruser:dockeruser /usr/local/rustup && chown dockeruser:dockeruser /usr/local/rustup

USER dockeruser
ENV USER=dockeruser HOME=$APP_DIR