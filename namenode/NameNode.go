package main

import (
	"log"
	"net"

	"github.com/dnorambu/tarea2sd/biblioteca"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := biblioteca.Server{}

	grpcServer := grpc.NewServer()

	biblioteca.RegisterNameNodeServerServiceServer(grpcServer, &s)
	//Instanciar variables de la estructura Server

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
