.PHONY: make_user_rpc

user-proto:
	chmod +x ./apps/user/make_rpc.sh && ./apps/user/make_rpc.sh

user-rpc-dev:
	@make -f ./deploy/mk/user_rpc.mk release-test 