# EXAMPLE USAGE:  docker run --rm codewars/python-runner run -l python -c "print 1+1 "
#                 docker run --rm codewars/python-runner run -l python3 -c "print(1+1)"

# Pull base image.
FROM python:3.8.2-alpine3.11

ENV APP_DIR /app

WORKDIR $APP_DIR

RUN adduser -D -h $APP_DIR -s /bin/bash dockeruser
RUN chown -R dockeruser:dockeruser $APP_DIR

RUN chmod -R g+r $APP_DIR
RUN chown -R dockeruser:dockeruser /usr/local/bin/python3 && chown -R dockeruser /usr/local/bin/python3

USER dockeruser
ENV USER=dockeruser HOME=$APP_DIR