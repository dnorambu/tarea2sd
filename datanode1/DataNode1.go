package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"sync"

	pb "github.com/dnorambu/tarea2sd/bibliotecadn"
	nn "github.com/dnorambu/tarea2sd/bibliotecann"

	"google.golang.org/grpc"
)

//Server Se declara la estructura del servidor
type Server struct {
	pb.UnimplementedDataNodeServiceServer
	//Slice que guarda en memoria RAM los chunks que un cliente me envia
	ChunksRecibidos []*pb.UploadBookRequest
	//
	Mu sync.Mutex
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

	//c := courier.NewCourierServiceClient(conn)
}
func conectarConNn() (nn.NameNodeServiceClient, *grpc.ClientConn) {

	//Para realizar pruebas locales
	conn, err := grpc.Dial("localhost:9001", grpc.WithInsecure())
	// conn, err := grpc.Dial(10.10.28.14:9000, grpc.WithInsecure())

	if err != nil {
		//OJO, si el NN no esta funcionando, el programa terminara la ejecucion
		log.Fatalf("Se cayo el name node, adios", err)
	}
	c := nn.NewNameNodeServiceClient(conn)
	fmt.Println("Conectado a NameNode: 10.10.28.14:9000")
	return c, conn
}

// UploadBookCentralizado sirve para subir chunks de libros a traves de un stream
func (s *Server) UploadBookCentralizado(stream pb.DataNodeService_UploadBookCentralizadoServer) error {
	contador := 0
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.UploadBookResponse{
				Respuesta: "Libro enviado exitosamente",
			})
		}
		contador++
		fmt.Println("Numero de partes recibidas: ", contador)
		s.ChunksRecibidos = append(s.ChunksRecibidos, chunk)
	}
	// Implementar propuesta y posterior distribucion
}
func (s *Server) crearPropuesta() {
	largo := len(s.ChunksRecibidos)
	var chunksDataNode1, chunksDataNode2, chunksDataNode3 int64
	var aleatorio int

	if largo == 1 {
		aleatorio = rand.Intn(2) + 1 //DN etiquetados de 1 a 3
		if aleatorio == 1 {
			chunksDataNode1++

		} else if aleatorio == 2 {
			chunksDataNode2++

		} else {
			chunksDataNode3++

		}
	} else if largo == 2 {
		//En este caso, predefinimos los envios segun prioridad por ID (de mayor a menor)
		chunksDataNode3++
		chunksDataNode2++

	} else if largo >= 3 {
		// Se distribuyen los chunks siguiendo la prioridad (mayor a menor)
		// La distribucion funciona similar al modulo 3
		for i := 0; i < largo; {
			if largo-i == 1 {
				chunksDataNode3++
				i++
			}
			if largo-i == 2 {
				chunksDataNode3++
				chunksDataNode2++
				i += 2
			}
			if largo-i >= 3 {
				chunksDataNode3++
				chunksDataNode2++
				chunksDataNode1++
				i += 3
			}
		}
	} else {
		fmt.Println("No hay chunks, raro pero cierto")
		return
	}
	//Hacer y enviar la propuesta
	propuesta := &nn.Propuesta{
		Chunksmaquina1: chunksDataNode1,
		Chunksmaquina2: chunksDataNode2,
		Chunksmaquina3: chunksDataNode3,
	}
	//Nos conectamos al NN
	clienteNn, conexionNn := conectarConNn()
	defer conexionNn.Close()
	confirmacion, err := clienteNn.SendPropuesta(context.Background(), propuesta)
	if confirmacion {

	}
}

// UploadBookDistribuido para recibir chunks por medio de un stream
func (s *Server) UploadBookDistribuido(stream pb.DataNodeService_UploadBookDistribuidoServer) error {
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
