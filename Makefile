generate: ## To generate file using protoc.
	mkdir Generate
	protoc -I Protoc --go_out=./Generate --go-grpc_out=./Generate --go_opt=module=grpc_dummy --go-grpc_opt=module=grpc_dummy ./Protoc/dummy.proto

build:
	go build -o bin/Server ./Server
	go build -o bin/Client ./Client

run_server:
	./bin/Server

run_client:
	./bin/Client

help:
	echo "Will setup help soon..."