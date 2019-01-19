FROM golang:1.11.4

RUN apt-get update && \
    apt-get install -y vim less

# copy application code from host.
WORKDIR /go
