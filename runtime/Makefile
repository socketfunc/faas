proto-build:
	protoeasy --grpc --go ./proto

build-arm:
	GOOS=linux GOARCH=arm GOARM=7 go build -o runtime main.go
