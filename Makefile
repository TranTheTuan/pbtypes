.PHONY: build install clean upload

install:
	go get -u \
		github.com/gogo/protobuf/protoc-gen-gogo \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
		github.com/mwitkow/go-proto-validators/protoc-gen-govalidators \
		github.com/gogo/protobuf/jsonpb \
		github.com/gogo/protobuf/gogoproto \
		github.com/gogo/protobuf/proto

build: 
	mkdir -pv build/swagger/OpenAPI
	./regenerate_go.sh

clean:
	find ./build ./dist ! -name ".gitkeep" -delete
