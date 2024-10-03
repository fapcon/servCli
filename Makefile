.PHONY: run gen

run:
	go run grpcServ/cmd/main.go

gen:
	protoc --go_out=. --go-grpc_out=. proto/*.proto