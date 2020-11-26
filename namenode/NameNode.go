package main

import (
	"log"
	"net"

	"github.com/dnorambu/pruebas/courier"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := courier.Server{}

	grpcServer := grpc.NewServer()

	courier.RegisterCourierServiceServer(grpcServer, &s)
	//Instanciar variables de la estructura Server
	s.MapaSeguimiento = make(map[int64]string)
	s.OrdenesP = make([]*courier.OrdenPyme, 0)
	s.OrdenesR = make([]*courier.OrdenRetail, 0)
	s.ColaNormal = make([]*courier.Paquete, 0)
	s.ColaPrioritario = make([]*courier.Paquete, 0)
	s.ColaRetail = make([]*courier.Paquete, 0)
	s.EntregasC = make([]*courier.Entrega, 0)
	s.MapaSegRetail = make(map[string]string)
	s.MapaIntermedio = make(map[string]int64)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
