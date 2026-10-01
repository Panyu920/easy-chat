SERVER_NAME=im
SERVER_TYPE=ws

CONF_PATH=./apps/${SERVER_NAME}/${SERVER_TYPE}/etc/${SERVER_NAME}.yaml
run-test:
	@echo "run ${SERVER_NAME} ${SERVER_TYPE}"
	@go run ./apps/im/ws/im.go -f ${CONF_PATH}
