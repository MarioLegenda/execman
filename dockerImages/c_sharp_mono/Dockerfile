FROM ubuntu:bionic

ENV APP_DIR /app
ARG DEBIAN_FRONTEND=noninteractive

RUN apt update
RUN apt -y install gnupg ca-certificates
RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 3FA7E0328081BFF6A14DA29AA6A19B38D3D831EF
RUN "deb https://download.mono-project.com/repo/ubuntu stable-focal main" | tee /etc/apt/sources.list.d/mono-official-stable.list
RUN apt update

RUN apt -y install mono-complete

WORKDIR /app

RUN useradd -d $APP_DIR -s /bin/bash dockeruser
RUN chown -R dockeruser:dockeruser $APP_DIR
RUN chmod -R 777 $APP_DIR

USER dockeruser

ENV USER=dockeruser HOME=$APP_DIR