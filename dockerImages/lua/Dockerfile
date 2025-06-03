FROM woahbase/alpine-lua:x86_64

ENV APP_DIR /app

WORKDIR $APP_DIR

RUN adduser -D -h $APP_DIR -s /bin/bash dockeruser
RUN chown -R dockeruser:dockeruser $APP_DIR

RUN chmod -R g+r $APP_DIR
RUN chown -R dockeruser:dockeruser /usr/bin/lua && chown -R dockeruser /usr/bin/lua

USER dockeruser
ENV USER=dockeruser HOME=$APP_DIR