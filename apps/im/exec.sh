goctl model mongo --type ChatLog --dir "./apps/im/im_model"  
goctl model mongo --type Conversations --dir "./apps/im/im_model"  
goctl model mongo --type Conversation --dir "./apps/im/im_model"  

goctl rpc  protoc ./apps/im/rpc/proto/im.proto --go_out=./apps/im/rpc --go-grpc_out=./apps/im/rpc --zrpc_out=./apps/im/rpc --client=true

