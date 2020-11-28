package main

import (
	"context"
	"log"
	"net"
	"sync"

	nn "github.com/dnorambu/tarea2sd/bibliotecann"
	"google.golang.org/grpc"
)

//Server Se declara la estructura del servidor
type Server struct {
	nn.UnimplementedNameNodeServiceServer

	//
	Mu sync.Mutex
}

//SendPropuesta implementada
func (s *Server) SendPropuesta(ctx context.Context, propuesta *nn.Propuesta) (*nn.Confirmacion, error) {
	var err error
	return nil, err
}
func newServer() *Server {
	s := &Server{}
	return s
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}
	grpcServer := grpc.NewServer()

	nn.RegisterNameNodeServiceServer(grpcServer, newServer())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
