FROM dart:stable

ENV APP_DIR /app

WORKDIR $APP_DIR

RUN adduser --disabled-password --home /app --shell /bin/bash dockeruser
RUN chown -R dockeruser:dockeruser /app

RUN chmod -R g+r /app
RUN chown -R dockeruser:dockeruser /usr/lib/dart/bin/dart && chown -R dockeruser /usr/lib/dart/bin/dart
USER dockeruser

ENV HOME=$APP_DIR