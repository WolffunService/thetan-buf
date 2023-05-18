install :
#if macos - Darwin
	brew install protobuf
	brew install buf

build :
	#buf generate proto
	protoc --proto_path=./proto/ --go_out=./gen/go/. --go-grpc_out=./gen/go/. ./proto/*.proto