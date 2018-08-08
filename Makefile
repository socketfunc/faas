proto-build:
	protoc --go_out=plugins=grpc:. runtime/proto/runtime.proto
