SERVER_NAME=social
SERVER_TYPE=rpc

CONF_PATH=./apps/${SERVER_NAME}/${SERVER_TYPE}/etc/${SERVER_NAME}.yaml
run-test:
	@echo "run ${SERVER_NAME} ${SERVER_TYPE}"
	@go run ./apps/social/rpc/social.go -f ${CONF_PATH}
