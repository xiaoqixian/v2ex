# Date:   Sun Sep 21 04:00:05 PM 2025
# Mail:   lunar_ubuntu@qq.com
# Author: https://github.com/xiaoqixian

import os

TEMPLATE = """FROM golang:1.24 AS builder

WORKDIR /app
WORKDIR /rpc_gen

COPY app/go.work app/go.work
COPY app/go.work.sum app/go.work.sum

COPY app/{service} /app/{service}
COPY app/common /app/common
{post_copy}
COPY rpc_gen /rpc_gen

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

RUN cd /app/{service} && go mod tidy

# static link
RUN cd /app/{service} && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o {service}_service .

FROM alpine:latest

COPY --from=builder /app/{service}/{service}_service /{service}_service
COPY --from=builder /app/{service}/conf/conf.yaml /conf/conf.yaml

WORKDIR /

{expose}

ENV MYSQLADDR=mysql
ENV REDISADDR=my_redis
ENV KAFKAADDR=kafka
ENV CONSULADDR=my_consul
ENV ESADDR=es01
CMD ["/{service}_service"]
"""

services = ["post", "home", "rec", "user", "comment"]

for svc in services:
    target_dir = f"app/{svc}"
    os.makedirs(target_dir, exist_ok=True)

    fmt_args = {
        "service": svc,
        "expose": "" if svc != "home" else "EXPOSE 8080",
        "post_copy": "" if svc != "rec" else "COPY app/post /app/post"
    }

    dockerfile_path = os.path.join(target_dir, "Dockerfile")
    with open(dockerfile_path, "w") as f:
        f.write(TEMPLATE.format(**fmt_args))
    
    print(f"✅ Generated {dockerfile_path}")
