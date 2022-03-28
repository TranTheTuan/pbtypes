#!/usr/bin/env bash

find * -type f ! -name third_party -name *.proto | while read -r d; do
    echo ${d}
    protoc -I . \
   --go_out ./build/go \
   --go-grpc_out ./build/go \
   --grpc-gateway_out ./build/go --grpc-gateway_opt paths=source_relative \
    ${d}

done
