SERVER_NAME=task
SERVER_TYPE=mq

CONF_PATH=./apps/${SERVER_NAME}/${SERVER_TYPE}/etc/dev/${SERVER_NAME}.yaml
run-test:
	@echo "run task mq"
	@go run apps/${SERVER_NAME}/${SERVER_TYPE}/${SERVER_NAME}.go -f ${CONF_PATH}
