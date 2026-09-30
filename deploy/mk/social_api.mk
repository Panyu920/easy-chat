SERVER_NAME=social
SERVER_TYPE=api

CONF_PATH=./apps/${SERVER_NAME}/${SERVER_TYPE}/etc/${SERVER_NAME}.yaml
run-test:
	@echo "run social api"
	@go run apps/social/api/social.go -f ${CONF_PATH}
