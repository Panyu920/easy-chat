.PHONY: make_user_rpc

# 生成用户RPC服务的pb文件
user-proto:
	chmod +x ./apps/user/make_rpc.sh && ./apps/user/make_rpc.sh

# 编译并发布用户RPC服务的开发环境镜像
user-rpc-dev:
	@make -f ./deploy/mk/user_rpc.mk release-test 


release-dev:
	chmod 777 ./components -R && cd ./deploy/script && chmod +x release_dev.sh && ./release_dev.sh


db_schema:
	dbml2sql  ./doc/schema.dbml -o ./doc/schema.sql  --mysql

db_create:
	@make -f ./deploy/mk/db.mk create_db

db_new_migration:
	@make -f ./deploy/mk/db.mk new_migration
db_migrate_up:
	@make -f ./deploy/mk/db.mk migrate_up
db_migrate_down:
	@make -f ./deploy/mk/db.mk migrate_down
db_sqlc:
	@make -f ./deploy/mk/db.mk sqlc

run_user_rpc:
	@make -f ./deploy/mk/user_rpc.mk run-test

run_user_api:
	@make -f ./deploy/mk/user_api.mk run-test
