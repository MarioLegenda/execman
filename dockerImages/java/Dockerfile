FROM debian:bullseye-slim

# Install required tools
RUN apt-get update && apt-get install -y curl tar ca-certificates && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Define environment variables
ENV JAVA_VERSION=22 \
    JAVA_BUILD=36 \
    JAVA_URL=https://github.com/adoptium/temurin22-binaries/releases/download/jdk-22%2B36/OpenJDK22U-jdk_x64_linux_hotspot_22_36.tar.gz

# Download and extract the JDK
RUN curl -L "$JAVA_URL" | tar -xz -C /opt && \
    ln -s /opt/jdk-22+36 /opt/jdk

# Add javac and java to /usr/local/bin
RUN ln -s /opt/jdk/bin/java /usr/local/bin/java && \
    ln -s /opt/jdk/bin/javac /usr/local/bin/javac

# Set JAVA_HOME and PATH
ENV JAVA_HOME=/opt/jdk
ENV PATH="${JAVA_HOME}/bin:${PATH}"

ENV APP_DIR /app

RUN useradd -m -d $APP_DIR -s /bin/bash dockeruser
RUN chown -R dockeruser:dockeruser $APP_DIR

RUN chmod -R g+r $APP_DIR
RUN chown -R dockeruser:dockeruser /usr/local/bin/javac && chown -R dockeruser /usr/local/bin/javac

USER dockeruser
ENV USER=dockeruser HOME=$APP_DIR

# Default command
CMD ["javac", "-version"]