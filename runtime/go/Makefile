run:
	_RUNTIME_PORT=8080 go run main.go

debug:
	dlv debug main.go

build:
	go build -buildmode=plugin -o handler.so ./test/main.go

lint:
	go tool vet main.go
