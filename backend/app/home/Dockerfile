FROM golang:1.24 AS builder

WORKDIR /app
WORKDIR /rpc_gen

COPY app/go.work app/go.work
COPY app/go.work.sum app/go.work.sum

COPY app/home /app/home
COPY app/common /app/common

COPY rpc_gen /rpc_gen

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

RUN cd /app/home && go mod tidy

# static link
RUN cd /app/home && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o home_service .

FROM alpine:latest

COPY --from=builder /app/home/home_service /home_service
COPY --from=builder /app/home/conf/conf.yaml /conf/conf.yaml

WORKDIR /

EXPOSE 8080

ENV MYSQLADDR=mysql
ENV REDISADDR=my_redis
ENV KAFKAADDR=kafka
ENV CONSULADDR=my_consul
ENV ESADDR=es01
CMD ["/home_service"]
