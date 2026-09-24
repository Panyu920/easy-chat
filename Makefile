.PHONY: make_user_rpc

# 生成用户RPC服务的pb文件
user-proto:
	chmod +x ./apps/user/make_rpc.sh && ./apps/user/make_rpc.sh

# 编译并发布用户RPC服务的开发环境镜像
user-rpc-dev:
	@make -f ./deploy/mk/user_rpc.mk release-test 


release-dev:
	chmod 777 ./components -R && cd ./deploy/script && chmod +x release_dev.sh && ./release_dev.sh
