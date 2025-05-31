FROM gcc:4.9

ENV APP_DIR /app

WORKDIR $APP_DIR

RUN useradd -M -d $APP_DIR --shell /bin/bash dockeruser
RUN chown -R dockeruser:dockeruser $APP_DIR

USER dockeruser
ENV USER=dockeruser HOME=$APP_DIR