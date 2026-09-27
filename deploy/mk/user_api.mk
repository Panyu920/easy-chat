SERVER_NAME=user
SERVER_TYPE=api

CONF_PATH=./apps/${SERVER_NAME}/${SERVER_TYPE}/etc/${SERVER_NAME}.yaml
run-test:
	@echo "run user api"
	@go run apps/user/api/user.go -f ${CONF_PATH}
