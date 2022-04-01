FROM golang:alpine

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /go/src/seckill-jiujia
COPY . .
RUN go env && go build -buildvcs=false -o seckill-jiujia .

FROM alpine:latest
LABEL MAINTAINER="ruitao.Yang"

WORKDIR /
COPY --from=0 /go/src/seckill-jiujia/seckill-jiujia .
COPY --from=0 /go/src/seckill-jiujia/config.toml .

# 修改时区
RUN apk add --update tzdata && \
    rm -rf /etc/localtime && \
    ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

EXPOSE 18803

ENTRYPOINT ./seckill-jiujia
