FROM debian:stretch-slim
LABEL version="1.0" maintainer="docker user <kevinacoder@github>"
RUN apt-get update && \
    apt-get install -y python3 && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*