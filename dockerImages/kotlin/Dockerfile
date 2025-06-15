FROM zenika/kotlin

ENV APP_DIR /app

WORKDIR $APP_DIR

RUN useradd -d $APP_DIR -s /bin/bash dockeruser
RUN chmod -R u+rw $APP_DIR
RUN chown -R dockeruser:dockeruser $APP_DIR && chown dockeruser:dockeruser /usr/lib/kotlinc/bin && chown dockeruser:dockeruser /usr/lib/kotlinc/bin

USER dockeruser

ENV USER=dockeruser HOME=$APP_DIR