.PHONY: compile

compilenn:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative bibliotecann/NameNodeServer.proto
compiledn:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative bibliotecadn/DataNodeServer.proto
cliente:
	go run clientnode/ClientNode.go
dn1:
	cd datanode1; \
	go run DataNode1.go
dn2:
	cd datanode2; \
	go run DataNode2.go
dn3:
	cd datanode3; \
	go run DataNode3.go
nn: 
	cd namenode; \
	go run NameNode.go