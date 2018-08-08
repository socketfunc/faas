proto-build:
	protoc --go_out=plugins=grpc,import_prefix=github.com/socketfunc/:. -I=".:${GOPATH}/src/github.com/socketfunc" ./proto/*.proto
	sed -i.bak -E 's/github.com\/socketfunc\/(github\.com|golang\.org|google\.golang\.org)/\1/g' ./proto/*.pb.go
	rm -f ./proto/*.bak
