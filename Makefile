.PHONY: compile

compile:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biblioteca/DataNodeServer.proto
#cliente:
#	go run cliente_courier/cliente.go
