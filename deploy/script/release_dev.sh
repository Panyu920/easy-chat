#!/bin/bash

need_start_sh=(
    user_rpc_dev.sh
)

for sh in ${need_start_sh[*]}; do
    chmod +x ${sh}
    ./${sh} 
done

docker ps 

# 查看etcd
docker exec -it etcd etcdctl get --prefix ""
