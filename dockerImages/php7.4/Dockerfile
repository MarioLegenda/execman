FROM alpine:3.11

ENV APP_DIR /app

WORKDIR $APP_DIR

RUN apk add php7-cli

RUN adduser -D -h $APP_DIR -s /bin/bash dockeruser
RUN chown -R dockeruser:dockeruser $APP_DIR

RUN chmod -R g+r $APP_DIR
RUN chown -R dockeruser:dockeruser /usr/bin/php && chown -R dockeruser /usr/bin/php

RUN php -r "ini_set('upload_max_filesize', '0M');"
RUN php -r "ini_set('memory_limit', '8M');"
RUN php -r "ini_set('post_max_size', '0M');"
RUN php -r "ini_set('file_uploads', 'Off');"
RUN php -r "ini_set('expose_php', 'Off');"
RUN php -r "ini_set('register_globals', 'Off');"
RUN php -r "ini_set('allow_url_fopen', 'Off');"
RUN php -r "ini_set('safe_mode', 'On');"
RUN php -r "ini_set('sql.safe_mode', 'On');"

USER dockeruser
ENV USER=dockeruser HOME=$APP_DIR