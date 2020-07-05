FROM ubuntu:devel
RUN apt update \
    && apt install -y lsb-release
WORKDIR /app
COPY rhinoctl /app
RUN /app/rhinoctl -f -d
