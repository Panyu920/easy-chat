SERVER_NAME=im
SERVER_TYPE=rpc

CONF_PATH=./apps/${SERVER_NAME}/${SERVER_TYPE}/etc/dev/${SERVER_NAME}.yaml
run-test:
	@echo "run ${SERVER_NAME} ${SERVER_TYPE}"
	@go run ./apps/${SERVER_NAME}/${SERVER_TYPE}/${SERVER_NAME}.go -f ${CONF_PATH}
