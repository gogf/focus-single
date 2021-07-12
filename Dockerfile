###############################################################################
#                                BUILDE
###############################################################################

# 打包依赖阶段使用golang作为基础镜像
FROM golang:1.15.7-alpine as builder

# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN cp docker/config.docker.toml config/config.toml
# 指定OS等，并go build
RUN wget https://goframe.org/cli/linux_amd64/gf && chmod +x gf
RUN rm -rf packed/data.go && \
    ./gf install && \
    gf build

# 由于我不止依赖二进制文件，还依赖views文件夹下的html文件还有assets文件夹下的一些静态文件
# 所以我将这些文件放到了publish文件夹
RUN mkdir -p publish/config && \
    cp bin/linux_amd64/focus publish && \
    cp -r docker/config.docker.toml publish/config/config.toml

###############################################################################
#                                   START
###############################################################################
FROM alpine:3.13

LABEL maintainer="focus"

WORKDIR /app
# 将上一个阶段publish文件夹下的所有文件复制进来
COPY --from=builder /app/publish .
RUN chmod +x focus

EXPOSE 8199

CMD ./focus
