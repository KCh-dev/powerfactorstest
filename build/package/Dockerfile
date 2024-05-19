FROM golang:1.22.0 AS build

ARG VERSION
ARG DEBUG

COPY . /go/src/goapp
WORKDIR /go/src/goapp

RUN \
    mkdir -p bin && \
    make goapp

FROM debian:12-slim

RUN \
    apt-get -y update && \
    apt-get -y upgrade && \
    apt-get -y install ca-certificates iputils-ping net-tools netcat-traditional procps tzdata && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

COPY --from=build /go/src/goapp/bin /goapp

# Set executable permissions for the script
RUN chmod +x /goapp/scripts/run_tests.sh && \
    chmod +x /goapp/scripts/run_client.sh

ENTRYPOINT ["/goapp/goapp"]