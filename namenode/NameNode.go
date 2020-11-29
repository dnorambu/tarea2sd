package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	nn "github.com/dnorambu/tarea2sd/bibliotecann"

	"google.golang.org/grpc"
)

var (
	localnn  = "localhost:9000"
	localdn1 = "localhost:9001"
	localdn2 = "localhost:9002"
	localdn3 = "localhost:9003"
	nameNode = "10.10.28.14:9000"
	dn1      = "10.10.28.140:9000"
	dn2      = "10.10.28.141:9000"
	dn3      = "10.10.28.142:9000"
)

//Server Se declara la estructura del servidor
type Server struct {
	nn.UnimplementedNameNodeServiceServer

	//
	Mu sync.Mutex
}

func conectarConDnDesdeNn(ipDestino string) bool {

	conn, err := grpc.Dial(ipDestino, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

//SendPropuesta implementada para que DataNode pueda enviar propuesta inicial a NameNode y este ultimo
//devuelve la misma propuesta si es aceptada, y en el caso de que no se acepte se envia otra propuesta distinta
func (s *Server) SendPropuesta(ctx context.Context, propuesta *nn.Propuesta) (*nn.Propuesta, error) {
	//BORRAR
	fmt.Println("La propuesta recibida desde el DN es: (1,2,3) ",
		propuesta.Chunksmaquina1, propuesta.Chunksmaquina2, propuesta.Chunksmaquina3)
	var err error
	maquinas := map[string]int64{
		"maquina3": propuesta.Chunksmaquina3,
		"maquina2": propuesta.Chunksmaquina2,
		"maquina1": propuesta.Chunksmaquina1,
	}
	//Barrido inicial, queremos saber de antemano que DataNodes estan activos
	// estaVivo3 := conectarConDnDesdeNn("10.10.28.142:9000")
	// estaVivo2 := conectarConDnDesdeNn("10.10.28.141:9000")
	// estaVivo1 := conectarConDnDesdeNn("10.10.28.140:9000")

	//Para test local
	estaVivo3 := conectarConDnDesdeNn(localdn3)
	estaVivo2 := conectarConDnDesdeNn(localdn2)
	estaVivo1 := conectarConDnDesdeNn(localdn1)
	fmt.Println("Estados de las maquinas (vivas o muertas segun NN):\n",
		"M1", estaVivo1, "\n",
		"M2", estaVivo2, "\n",
		"M3", estaVivo3)
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
	//Generamos un nuevo map que contendra la cantidad de chunks asignados a cada maquina para la nueva propuesta
	propuestaNueva := map[string]int64{
		"maquina3": 0,
		"maquina2": 0,
		"maquina1": 0,
	}
	//Se procede a asignar la cantidad de chunks por maquina siguiendo el orden de mayor a menor como se hizo en DataNode
	//Las maquinas caidas siempre van a quedar con una cantidad de chunks igual a 0
	for totalChunks >= 1 {
		if totalChunks >= 1 && estadoDeMaquina["maquina3"] {
			propuestaNueva["maquina3"]++
			totalChunks--
		}
		if totalChunks >= 1 && estadoDeMaquina["maquina2"] {
			propuestaNueva["maquina2"]++
			totalChunks--
		}
		if totalChunks >= 1 && estadoDeMaquina["maquina1"] {
			propuestaNueva["maquina1"]++
			totalChunks--
		}
	}
	//Ahora se genera el mensaje a retornar al DataNode con la propuesta final generada
	propuestaFinal := &nn.Propuesta{
		Chunksmaquina1: propuestaNueva["maquina1"],
		Chunksmaquina2: propuestaNueva["maquina2"],
		Chunksmaquina3: propuestaNueva["maquina3"],
	}
	fmt.Println("La nueva propuesta es: ",
		propuestaNueva["maquina1"],
		propuestaNueva["maquina2"],
		propuestaNueva["maquina3"])
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
