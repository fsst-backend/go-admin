FROM docker.fastdocker.com:5000/violet-base:3.20.3

# 替换 apk 源、安装包并设置 UTC 时区
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk add --no-cache \
        ca-certificates \
        tzdata \
        gcc \
        g++ \
        libc6-compat \
    && ln -sf /usr/share/zoneinfo/UTC /etc/localtime \
    && echo "UTC" > /etc/timezone


ENV TZ=UTC

WORKDIR /app

# 复制程序和配置
COPY build/linux/go-admin ./go-admin
COPY config/db.sql ./db.sql
COPY config/db-begin-mysql.sql ./db-begin-mysql.sql
COPY config/db-end-mysql.sql ./db-end-mysql.sql
COPY docs ./docs
COPY config/menu.json ./menu.json
COPY config/settings.template.yaml ./settings.yaml

EXPOSE 13348
RUN  chmod +x /app/go-admin
CMD ["/app/go-admin","server","-c", "/app/settings.yaml"]
