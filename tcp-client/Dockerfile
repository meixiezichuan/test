# 使用轻量级 alpine 镜像作为最终镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /

# 复制编译好的可执行文件
COPY client .

# 启动客户端
CMD ["/client"]
