#!/usr/bin/env bash

find ./proto -type f -name *.proto | while read -r d; do
    echo ${d}
    protoc -I . \
   --go_out ./build/go --go_opt=module=github.com/TranTheTuan/pbtypes/build/go \
   --go-grpc_out ./build/go --go-grpc_opt=module=github.com/TranTheTuan/pbtypes/build/go \
   --grpc-gateway_out ./build/go --grpc-gateway_opt=paths=import --grpc-gateway_opt=module=github.com/TranTheTuan/pbtypes/build/go \
    ${d}

done
