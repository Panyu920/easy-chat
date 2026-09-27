goctl rpc  protoc ./apps/social/rpc/proto/social.proto --go_out=./apps/social/rpc --go-grpc_out=./apps/social/rpc --zrpc_out=./apps/social/rpc --client=true

goctl model mysql ddl -src="./deploy/sql/social.sql" -dir="./apps/social/socialmodels/" -c
