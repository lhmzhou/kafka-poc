FROM golang:1.13.1 as builder

RUN apt-get -y update \
        && apt-get install -y --no-install-recommends upx-ucl zip libssl-dev \
        && apt-get clean \
        && rm -rf /var/lib/apt/lists/*

WORKDIR /src/
COPY . .
ENV HTTP_PROXY ${CURRENT_PROXY}
RUN go build -o app main.go
ENTRYPOINT ["/bin/sh", "-c", "./app"]