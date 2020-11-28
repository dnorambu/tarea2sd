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

func conectarConDnDesdeNn(ipDestino string) bool {
	//Para realizar pruebas locales
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	// conn, err := grpc.Dial(ipDestino, grpc.WithInsecure())
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

//SendPropuesta implementada
func (s *Server) SendPropuesta(ctx context.Context, propuesta *nn.Propuesta) (*nn.Propuesta, error) {
	var err error
	maquinas := map[string]int64{
		"maquina3": propuesta.Chunksmaquina3,
		"maquina2": propuesta.Chunksmaquina2,
		"maquina1": propuesta.Chunksmaquina1,
	}
	//Barrido inicial, queremos saber de antemano que DataNodes estan activos
	estaVivo3 := conectarConDnDesdeNn("10.10.28.142:9000")
	estaVivo2 := conectarConDnDesdeNn("10.10.28.141:9000")
	estaVivo1 := conectarConDnDesdeNn("10.10.28.140:9000")

	estadoDeMaquina := map[string]bool{
		"maquina3": estaVivo3,
		"maquina2": estaVivo2,
		"maquina1": estaVivo1,
	}
	// Borramos del primer mapa aquellas maquinas no consideradas
	// en la propuesta inicial (que no recibian chunks)
	for maquina, chunks := range maquinas {
		if chunks == 0 {
			delete(maquinas, maquina)
		}
	}
	// Ahora hacemos una interseccion entre maquinas (mapa que considera
	// los DN de la propuesta original) y los DN vivos (mapa estadoDeMaquina)
	exitoPropuestaOriginal := true
	for maquina := range maquinas {
		if !estadoDeMaquina[maquina] {
			exitoPropuestaOriginal = false
		}
	}
	//Aprobar la propuesta porque todos estan vivos y pueden recibir chunks
	//Se retorna la misma propuesta que recibio la funcion
	if exitoPropuestaOriginal {
		return propuesta, err
	}
	//Ahora entramos al caso en que la propuesta original no sirve y NameNode debera considerar otra
	totalChunks := propuesta.Chunksmaquina1 + propuesta.Chunksmaquina2 + propuesta.Chunksmaquina3
	propuestaNueva := map[string]int64{
		"maquina3": 0,
		"maquina2": 0,
		"maquina1": 0,
	}
	for totalChunks >= 1 {
		if totalChunks >= 1 && estadoDeMaquina["maquina3"]{
			propuestaNueva["maquina3"]++
			totalChunks--
		}
		if totalChunks >= 1 && estadoDeMaquina["maquina2"]{
			propuestaNueva["maquina2"]++
			totalChunks--
		}
		if totalChunks >= 1 && estadoDeMaquina["maquina1"]{
			propuestaNueva["maquina1"]++
			totalChunks--
		}
	}

	propuestaFinal := &nn.Propuesta{
		Chunksmaquina1: propuestaNueva["maquina1"],
		Chunksmaquina2: propuestaNueva["maquina2"],
		Chunksmaquina3: propuestaNueva["maquina3"],
	}

	return propuestaFinal, err
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
