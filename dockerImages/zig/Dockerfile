FROM alpine:latest

# Install dependencies
RUN apk add --no-cache wget tar xz libstdc++ musl

# Set Zig version
ENV ZIG_VERSION=0.12.0
ENV ZIG_HOME=/opt/zig

# Download and install Zig
RUN wget https://ziglang.org/download/${ZIG_VERSION}/zig-linux-x86_64-${ZIG_VERSION}.tar.xz && \
    tar -xf zig-linux-x86_64-${ZIG_VERSION}.tar.xz && \
    mkdir -p ${ZIG_HOME} && \
    mv zig-linux-x86_64-${ZIG_VERSION}/* ${ZIG_HOME} && \
    ln -s ${ZIG_HOME}/zig /usr/local/bin/zig && \
    rm -rf zig-linux-x86_64-${ZIG_VERSION}*

# Add Zig to PATH
ENV PATH="${ZIG_HOME}:${PATH}"

ENV APP_DIR /app

WORKDIR $APP_DIR

RUN adduser -D -h $APP_DIR -s /bin/bash dockeruser
RUN chmod -R u+rw $APP_DIR
RUN chown -R dockeruser:dockeruser $APP_DIR && chown dockeruser:dockeruser /usr/local/bin/zig

USER dockeruser

ENV USER=dockeruser HOME=$APP_DIR