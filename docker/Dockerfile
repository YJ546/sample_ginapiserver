# Building stage
FROM golang:1.11-alpine3.9

RUN apk add --update make

WORKDIR /go/src/github.com/duyanghao/sample_apiserver

COPY . /go/src/github.com/duyanghao/sample_apiserver

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GO111MODULE=on

ENV TIMEZONE "Asia/Shanghai"

RUN make src.build
RUN mv build/pkg/cmd/sample_apiserver/sample_apiserver /go/bin

# Production stage
FROM alpine:3.9
WORKDIR /go/bin

LABEL maintainer="gaoyujian <1194373959@qq.com>"

# copy the go binaries from the building stage
COPY  --from=0 /go/bin /go/bin

COPY etc/sample_apiserver.yaml /etc/sample_apiserver/sample_iapiserver.yaml

EXPOSE 8081

ENTRYPOINT ["./sample_apiserver"]
