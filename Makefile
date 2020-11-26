.PHONY: compile

compile:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative bibliotecadn/DataNodeServer.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative bibliotecann/NameNodeServer.proto
#cliente:
#	go run cliente_courier/cliente.go
