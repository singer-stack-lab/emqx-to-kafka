FROM crpi-enwh9j989b0fsg8b.cn-guangzhou.personal.cr.aliyuncs.com/alexxqli/base:golang_base_1.25.0-v0.0.4 as builder

WORKDIR /backend
COPY . .

RUN rm -f config.yaml

# 构建参数
ARG ENV=test
COPY --from=0 /backend ./
#COPY --from=0 /backend/resource ./resource/
COPY --from=0 /backend/config.${ENV}.yaml ./config.yaml

# 安装 git
RUN apk add --no-cache git ca-certificates tzdata build-base


RUN go env -w GO111MODULE=on \
    && go env -w CGO_ENABLED=0 \
    && go env -w GOPRIVATE=github.com/singer-stack-lab \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/library/alpine:latest

LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

# 设置时区为上海
RUN apk update && apk add --no-cache tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

WORKDIR /backend

# 挂载目录：如果使用了sqlite数据库，容器命令示例：docker run -d -v /宿主机路径/gva.db:/go/src/github.com/flipped-aurora/gin-vue-admin/server/gva.db -p 8888:8888 --name gva-server-v1 gva-server:1.0
# VOLUME ["/go/src/github.com/flipped-aurora/gin-vue-admin/server"]

EXPOSE 9999
ENTRYPOINT ./server