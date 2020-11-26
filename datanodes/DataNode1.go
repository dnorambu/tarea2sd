package main

import (
	"log"
	"net"

	"github.com/dnorambu/pruebas/courier"
	"github.com/dnorambu/tarea2sd/biblioteca"
	"google.golang.org/grpc"
)

func main() {
	/*
		Implementar el input del usuario aquí, para elegir el método de eleccion
		de propuestas: distribuido o centralizado.
	*/
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := biblioteca.Server{}

	grpcServer := grpc.NewServer()

	biblioteca.RegisterDataNodeServerServiceServer(grpcServer, &s)
	//Instanciar variables de la estructura Server

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}

}

/*
Esta funcion permite crear una conexion entre el DataNode1 y algun otro
DataNode, de tal manera que DataNode1 sera el "cliente" y el otro DataNode
sera el "servidor" que recibe los chunks luego de que la propuesta de
distribucion fuera aceptada
*/
func conexionDatanodeCliente() {
	var conn *grpc.ClientConn
	//Para testear en local
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())

	//Para testear en MV
	// conn, err := grpc.Dial("10.10.28.141:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := courier.NewCourierServiceClient(conn)
}
