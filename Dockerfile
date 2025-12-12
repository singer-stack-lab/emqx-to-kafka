FROM crpi-enwh9j989b0fsg8b.cn-guangzhou.personal.cr.aliyuncs.com/alexxqli/base:golang_base_1.25.0-v0.0.4 as builder

WORKDIR /backend
COPY . .

RUN rm -f config.yaml

# 安装依赖
RUN apk add --no-cache git ca-certificates tzdata build-base


ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOPRIVATE=github.com/singer-stack-lab

RUN go mod tidy \
    && go build -o server .

# =============================

FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/library/alpine:latest

LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

# 设置时区
RUN apk add --no-cache tzdata \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

WORKDIR /backend

ARG ENV=test

# 只拷贝「产物」
COPY --from=builder /backend/server ./server
COPY --from=builder /backend/config.${ENV}.yaml ./config.yaml
RUN cat config.yaml
RUN ls

EXPOSE 9999
ENTRYPOINT ["./server"]
