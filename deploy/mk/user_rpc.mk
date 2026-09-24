# 版本号
VERSION=latest

# 服务名称
SERVER_NAME=user
# 服务类型
SERVER_TYPE=rpc

# 测试环境配置
# docker的镜像发布地址（镜像仓库地址）
DOCKER_REPO_TEST=ccr.ccs.tencentyun.com/pan-chat/${SERVER_NAME}-${SERVER_TYPE}-dev
# 测试版本
VERSION_TEST=$(VERSION)
# 编译的程序名称（本地镜像名称）
APP_NAME_TEST=easy-im-${SERVER_NAME}-${SERVER_TYPE}-test

# 测试下的编译文件（Dockerfile 路径）
DOCKER_FILE_TEST=./deploy/dockerfile/Dockerfile.${SERVER_NAME}_${SERVER_TYPE}_dev

# 测试环境的编译发布
build-test:
	@# 编译Go二进制文件 -mod=readonly：使用离线模式（不检查依赖更新）
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/${SERVER_NAME}-${SERVER_TYPE} ./apps/${SERVER_NAME}/${SERVER_TYPE}/${SERVER_NAME}.go
	
	@# 优化体积方法：
	@# 选择构建模式 默认模式（动态链接glibc）：适用Linux，静态编译模式（CGO_ENABLED=0）：适用容器、服务器端
	@# -ldflags="-s -w" -s移除符号表，-w移除DWARF调试信息（生产环境推荐，可减少10%~30%体积）
	@# -trimpath 去除路径信息（只能减少约 2-3M，编译慢很多）
	@# 编译后进一步优化体积：
	@# strip xx 移除额外符号信息
	@# UPX （运行时需要解压，启动稍慢）
	# GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o bin/${SERVER_NAME}-${SERVER_TYPE} ./apps/${SERVER_NAME}/${SERVER_TYPE}/${SERVER_NAME}.go
	
	@# 构建 Docker 镜像（不使用缓存）
	docker build . -f ${DOCKER_FILE_TEST} --no-cache -t ${APP_NAME_TEST}

# 镜像的测试标签
tag-test:

	@echo 'create tag ${VERSION_TEST}'
	@# 给镜像打上仓库标签
	docker tag ${APP_NAME_TEST} ${DOCKER_REPO_TEST}:${VERSION_TEST}

publish-test:

	@echo 'publish ${VERSION_TEST} to ${DOCKER_REPO_TEST}'
	@# 推送到镜像仓库
	docker push $(DOCKER_REPO_TEST):${VERSION_TEST}

# 串联多个步骤
release-test: build-test tag-test publish-test