# Use a minimal base image
FROM alpine:latest

# Set environment variables
ENV KOTLIN_VERSION=1.9.24

# Install dependencies and Kotlin compiler
RUN apk add --no-cache curl openjdk17 bash && \
    curl -sSL https://github.com/JetBrains/kotlin/releases/download/v${KOTLIN_VERSION}/kotlin-compiler-${KOTLIN_VERSION}.zip -o kotlin-compiler.zip && \
    unzip kotlin-compiler.zip -d /opt && \
    rm kotlin-compiler.zip && \
    ln -s /opt/kotlinc/bin/* /usr/local/bin/ && \
    apk del curl

ENV APP_DIR /app

WORKDIR $APP_DIR

RUN adduser -D -h $APP_DIR -s /bin/bash dockeruser
RUN chmod -R u+rw $APP_DIR
RUN chown -R dockeruser:dockeruser $APP_DIR && chown dockeruser:dockeruser /usr/local/bin && chown dockeruser:dockeruser /usr/local/bin

USER dockeruser

ENV USER=dockeruser HOME=$APP_DIR