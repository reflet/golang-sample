FROM golang:1.11.4

# locale & timezone (Asia/Tokyo)
# https://github.com/moby/moby/issues/12084
ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

# system update
RUN apt-get update && \
    apt-get install -y vim less

# copy application code from host.
WORKDIR /go
