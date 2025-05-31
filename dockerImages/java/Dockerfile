# RUN-USING:      docker run --rm --name=node codewars/node-runner --help
# EXAMPLE USAGE:  docker run --rm codewars/node-runner run -l javascript -c "console.log(1+1)"

# Pull base image.
FROM ubuntu:bionic

ENV APP_DIR /app

RUN apt-get update -qq && apt-get install -y --no-install-recommends \
        software-properties-common \
        build-essential \
		ca-certificates \
		curl \
		netbase \
		unzip \
		tar \
		xz-utils \
		wget \
	&& rm -rf /var/lib/apt/lists/*

WORKDIR $APP_DIR

COPY jdk-8u231-linux-x64.tar.gz .
RUN mkdir -p /usr/lib/jvm/jdk/
RUN mkdir -p /usr/lib/jvm/jdk/bin/java
RUN mkdir -p /usr/lib/jvm/jdk/bin/javac

RUN tar -zxvf jdk-8u231-linux-x64.tar.gz --one-top-level=java_tar --strip-components=1

RUN mv java_tar /usr/lib/jvm/jdk/

RUN update-alternatives --install /usr/bin/java java /usr/lib/jvm/jdk/java_tar/bin/java 100
RUN update-alternatives --install /usr/bin/javac javac /usr/lib/jvm/jdk/java_tar/bin/javac 100

RUN update-alternatives --config java

RUN addgroup dockerusergroup
RUN chmod g+x /usr/bin/java
RUN chmod g+x /usr/bin/javac
RUN chmod g+rwx $APP_DIR
RUN chgrp -R dockerusergroup $APP_DIR
RUN useradd -G dockerusergroup -d $APP_DIR -s /bin/bash dockeruser
RUN chown -R dockeruser:dockerusergroup $APP_DIR && chown -R dockeruser:dockerusergroup /usr/lib/jvm/jdk/ && chown -R dockeruser:dockerusergroup /usr/lib/jvm/jdk/

RUN javac -version

#RUN cp -v /bin/bash $BIN_DIR
#RUN cp -v /bin/ls $BIN_DIR

#RUN list="$(ldd /bin/bash | egrep -o '/lib.*\.[0-9]')"
#RUN for i in $list; do cp -v --parents "$i" "${APP_DIR}"; done
#RUN list="$(ldd /bin/ls | egrep -o '/lib.*\.[0-9]')"
#RUN for i in $list; do cp -v --parents "$i" "${APP_DIR}"; done

#RUN chroot --userspec=dockeruser:dockerusergroup /app /bin/bash

USER dockeruser

ENV USER=dockeruser HOME=$APP_DIR