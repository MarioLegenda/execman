FROM amazoncorretto:24-alpine-full

ENV APP_DIR /app

RUN adduser -D -h $APP_DIR -s /bin/bash dockeruser
RUN chown -R dockeruser:dockeruser $APP_DIR

RUN chmod -R g+r $APP_DIR
RUN chown -R dockeruser:dockeruser /usr/bin/javac && chown -R dockeruser /usr/bin/javac

USER dockeruser
ENV USER=dockeruser HOME=$APP_DIR