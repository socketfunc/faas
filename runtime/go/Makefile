.PHONY: run
run:
	_RUNTIME_PORT=8080 go run main.go

.PHONY: debug
debug:
	dlv debug main.go

.PHONY: build
build:
	go build -buildmode=plugin -o handler.so ./_example/main.go
