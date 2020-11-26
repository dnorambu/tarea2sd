package main

import (
	"io"
	"log"
	"net"
	"sync"

	pb "github.com/dnorambu/tarea2sd/biblioteca/bibliotecadn"

	"github.com/dnorambu/pruebas/courier"
	"google.golang.org/grpc"
)

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

//Server Se declara la estructura del servidor
type Server struct {
	pb.UnimplementedDataNodeServiceServer
	//Slice que guarda en memoria RAM los chunks que un cliente me envia
	ChunksRecibidos []*pb.UploadBookRequest

	//mutex que protegerá las variables compartidas
	Mu sync.Mutex
}

// UploadBook sirve para subir chunks de libros a traves de un stream
func (s *Server) UploadBookCentralizado(stream pb.DataNodeService_UploadBookServer) error {
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		s.ChunksRecibidos = append(s.ChunksRecibidos, chunk)
	}
	// Implementar propuesta y posterior distribucion
	return stream.SendAndClose(&pb.UploadBookResponse{
		Respuesta: "Libro enviado exitosamente",
	})
}
func (s *Server) UploadBookDistribuido(stream pb.DataNodeService_UploadBookServer) error {
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		s.ChunksRecibidos = append(s.ChunksRecibidos, chunk)
	}
	// Implementar propuesta y posterior distribucion
	return stream.SendAndClose(&pb.UploadBookResponse{
		Respuesta: "Libro enviado exitosamente",
	})
}

// mustEmbedUnimplementedCourierServiceServer solo se añadio por compatibilidad
// y evitar warnings al compilar
func (s *Server) mustEmbedUnimplementedDataNodeServiceServer() {}

func newServer() *Server {
	s := &Server{
		ChunksRecibidos: make([]*pb.UploadBookRequest, 0),
	}
	return s
}

func main() {
	/*
		Implementar el input del usuario aquí, para elegir el método de eleccion
		de propuestas: distribuido o centralizado.
	*/
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterDataNodeServiceServer(grpcServer, newServer())
	//Instanciar variables de la estructura Server

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}

}
