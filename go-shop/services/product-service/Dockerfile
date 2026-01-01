# --- 第一阶段：编译阶段 (命名为 builder) ---
# 使用 golang 1.24 的 alpine 版本作为编译环境，体积小且包含必要的 Go 工具链
FROM golang:1.24-alpine AS builder

# 设置容器内的当前工作目录为 /app
WORKDIR /app
# --- 关键：添加下面这一行 ---
ENV GOPROXY=https://goproxy.cn,direct
# 先把依赖管理文件复制到容器中
# 这样做是为了利用 Docker 的镜像层缓存：只要 go.mod/sum 没变，就不需要重新下载依赖
COPY go.mod go.sum ./

# 下载项目依赖（利用 GOPROXY 如果在国内构建可能需要配置）
RUN go mod download

# 复制项目的所有源代码到容器的工作目录
COPY . .

# 执行编译命令
# CGO_ENABLED=0: 禁用 CGO，实现静态编译，防止在 alpine 运行阶段缺少 glibc 库
# GOOS=linux: 目标操作系统为 Linux
# GOARCH=amd64: 目标架构为 x86_64
# -o product-service: 指定生成的二进制文件名为 product-service
# ./cmd: 告诉 Go 编译器 main 函数在 cmd 目录下
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o product-service ./cmd


# --- 第二阶段：运行阶段 ---
# 使用轻量级的 alpine 3.19 作为基础镜像，这只是一个极小的 Linux 运行环境
FROM alpine:3.19

# 设置运行时的基础目录
WORKDIR /app

# 从第一个阶段 (builder) 中只把编译好的二进制文件拷贝过来
# 这样最终镜像里就不会包含源代码、Go SDK 等几十 GB 的多余内容
COPY --from=builder /app/product-service .

# 声明容器内部监听的端口（仅起标识作用，实际映射看 docker-compose 或 -p 参数）
EXPOSE 9001

# 容器启动时默认执行的命令
CMD ["./product-service"]
