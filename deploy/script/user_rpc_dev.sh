#!/bin/bash

reso_addr="ccr.ccs.tencentyun.com/pan-chat/user-rpc-dev"
tag="latest"
container_name="easy-chat-user-rpc-dev"

# 停止并删除容器
docker stop $container_name
docker rm $container_name

# 拉取新的镜像
docker pull ${reso_addr}:${tag}

# 启动新的容器
docker run -d --name $container_name -p 10000:8080 ${reso_addr}:${tag}
