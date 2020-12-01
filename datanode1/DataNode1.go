package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"

	pb "github.com/dnorambu/tarea2sd/bibliotecadn"
	nn "github.com/dnorambu/tarea2sd/bibliotecann"

	"google.golang.org/grpc"
)

//Server Se declara la estructura del servidor
type Server struct {
	pb.UnimplementedDataNodeServiceServer
	//Slice que guarda en memoria RAM los chunks que un cliente me envia
	ChunksRecibidos []*pb.UploadBookRequest
	//Slice que guardara datos necesarios para hacer un stream hacia el NN
	Chunksaescribir []*nn.Logchunk
	Estado          string
	Mu 				sync.Mutex
}

/*
Esta funcion permite crear una conexion entre el DataNode1 y algun otro
DataNode, de tal manera que DataNode1 sera el "cliente" y el otro DataNode
sera el "servidor" que recibe los chunks luego de que la propuesta de
distribucion fuera aceptada
*/
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

func conectarConDn(maquina string) (pb.DataNodeServiceClient, *grpc.ClientConn) {
	//Para testear en local
	conn, err := grpc.Dial(maquina, grpc.WithInsecure())

	//Para testear en MV
	// conn, err := grpc.Dial(maquina, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Se cayo el DataNode durante la ejecucion: %s", err)
	}
	c := pb.NewDataNodeServiceClient(conn)
	fmt.Println("Conectado a DataNode:" + maquina)
	return c, conn
}

// Aclaracion: podra parecer que conectarConDnDesdeDn y conectarConDn hacen lo mismo, pero no. Sus propositos
// son diferentes y por ello se implementan de manera separada.
func conectarConDnDesdeDn(ipDestino string) bool {

	conn, err := grpc.Dial(ipDestino, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func conectarConNn() (nn.NameNodeServiceClient, *grpc.ClientConn) {

	//Para realizar pruebas locales
	conn, err := grpc.Dial(localnn, grpc.WithInsecure())
	// conn, err := grpc.Dial(10.10.28.14:9000, grpc.WithInsecure())

	if err != nil {
		//OJO, si el NN no esta funcionando, el programa terminara la ejecucion
		log.Fatalf("Se cayo el name node, adios: %s", err)
	}
	c := nn.NewNameNodeServiceClient(conn)
	fmt.Println("Conectado a NameNode: 10.10.28.14:9000")
	return c, conn
}

//Aqui empiezan las implementaciones de los rpc

//RequestCompetencia me permite saber si otros DN estan usando el LOG
func (s *Server) RequestCompetencia(ctx context.Context, rct *pb.Ricart) (*pb.Okrespondido, error) {
	var err error
	for {
		if s.Estado == "HELD" {
			//LOG Ocupado por el DN1, no necesita preguntar si su ID es mayor que otro DN porque eso nunca ocurrira
		} else {
			//DN1 es un caso especial porque su prioridad es la mas baja en el algoritmo distribuido
			return &pb.Okrespondido{Okay: true}, err
		}
	}
}

func (s *Server) saladeEsperaDistribuida() {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	//LOGIC
	s.Estado = "WANTED"
	//Verificar por diferentes casos
	//Preguntar si alguien mas tiene usado el LOG
	clienteDn2, conn2 := conectarConDn(localdn2)
	clienteDn3, conn3 := conectarConDn(localdn3)
	defer conn2.Close()
	defer conn3.Close()
	r1 := &pb.Ricart{
		Ip: localdn1,
		Id: 1,
	}
	awa, err := clienteDn2.RequestCompetencia(context.Background(), r1)
	if err != nil {
		fmt.Println("Ups! Error en la sala de espera: ", err)
	}
	owo, err := clienteDn3.RequestCompetencia(context.Background(), r1)
	if err != nil {
		fmt.Println("Ups! Error en la sala de espera: ", err)
	}
	if awa.Okay && owo.Okay {
		//NODO 2 Y 3 Me autorizaron para usar el LOG
		s.Estado = "HELD"
	}
}

//SendPropuestaDistribuida sirve para x cosa
func (s *Server) SendPropuestaDistribuida(ctx context.Context, prop *pb.Propuesta) (*pb.Okrespondido, error) {
	var err error

	if prop.Ipaverificar == 1 {
		dnverificado := conectarConDnDesdeDn(localdn1)
		if dnverificado {
			return &pb.Okrespondido{Okay: true}, err
		}
		return &pb.Okrespondido{Okay: false}, err
	}

	if prop.Ipaverificar == 2 {
		dnverificado := conectarConDnDesdeDn(localdn2)
		if dnverificado {
			return &pb.Okrespondido{Okay: true}, err
		}
		return &pb.Okrespondido{Okay: false}, err
	}

	if prop.Ipaverificar == 3 {
		dnverificado := conectarConDnDesdeDn(localdn3)
		if dnverificado {
			return &pb.Okrespondido{Okay: true}, err
		}
		return &pb.Okrespondido{Okay: false}, err
	}
	return &pb.Okrespondido{Okay: false}, err
}

// UploadBookCentralizado sirve para subir chunks de libros a traves de un stream
func (s *Server) UploadBookCentralizado(stream pb.DataNodeService_UploadBookCentralizadoServer) error {
	contador := 0
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			//Aqui es donde el cliente espera que su libro sea distribuido correctamente antes de poder
			//hacer otra cosa (se puede goroutines!)
			go s.crearPropuesta()
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

// UploadBookDistribuido para recibir chunks por medio de un stream
func (s *Server) UploadBookDistribuido(stream pb.DataNodeService_UploadBookDistribuidoServer) error {
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			go s.crearPropuestaDistribuida()
			return stream.SendAndClose(&pb.UploadBookResponse{
				Respuesta: "Libro enviado exitosamente",
			})
		}
		s.ChunksRecibidos = append(s.ChunksRecibidos, chunk)
	}
}

//DistributeBook es la funcion que recibe el stream de chunks (despues de propuesta) para guardar asi en disco para la maquina correspondiente
func (s *Server) DistributeBook(stream pb.DataNodeService_DistributeBookServer) error {
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.UploadBookResponse{
				Respuesta: "Libro recibido y guardado exitosamente",
			})
		}
		//Le damos el nombre
		fileName := chunk.Nombre
		_, err = os.Create(fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ioutil.WriteFile(fileName, chunk.Chunkdata, os.ModeAppend)
	}
}

//DownloadBook es para recibir el stream de nombres de partes y enviar un stream de las partes
func (s *Server) DownloadBook(stream pb.DataNodeService_DownloadBookServer) error {
	//var slicedePartes []*pb.PartChunk Creo que no sirve esto...
	for {
		nombreParte, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		parteActual := nombreParte.Nombre
		newFile, err := os.Open(parteActual)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer newFile.Close()
		parteInfo, err := newFile.Stat()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var parteSize int64 = parteInfo.Size()
		parteBufferBytes := make([]byte, parteSize)

		reader := bufio.NewReader(newFile)
		_, err = reader.Read(parteBufferBytes)

		parte := &pb.PartChunk{
			Chunkdata: parteBufferBytes,
		}
		if err := stream.Send(parte); err != nil {
			return err
		}
	}
}

func (s *Server) envChunks(dataNode string, cantidadDechunks int64) {
	var x *pb.UploadBookRequest
	var i int64

	clienteDn, conexionDn := conectarConDn(dataNode)
	defer conexionDn.Close() //no se si sea bueno usar un defer o simplemente hacer el .close al final del bloque if
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := clienteDn.DistributeBook(ctx)
	if err != nil {
		//Termina la ejecucion del programa por un error de stream
		log.Fatalf("No se pudo obtener el stream %v", err)
	}
	for i = 1; i <= cantidadDechunks; i++ {
		x, s.ChunksRecibidos = s.ChunksRecibidos[0], s.ChunksRecibidos[1:]
		logchunk := &nn.Logchunk{
			Nombre:    x.Nombre,
			Ipmaquina: dataNode,
		}
		s.Chunksaescribir = append(s.Chunksaescribir, logchunk)
		if err := stream.Send(x); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, x, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Se ha cerrado el stream hacia %v. %v", dataNode, reply.Respuesta)
	return
}

func (s *Server) crearPropuestaDistribuida() {
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
	propuesta1 := &pb.Propuesta{
		Chunksadn1:   chunksDataNode1,
		Chunksadn2:   chunksDataNode2,
		Chunksadn3:   chunksDataNode3,
		Ipaverificar: 3,
	}
	propuesta2 := &pb.Propuesta{
		Chunksadn1:   chunksDataNode1,
		Chunksadn2:   chunksDataNode2,
		Chunksadn3:   chunksDataNode3,
		Ipaverificar: 2,
	}

	clienteDn2, conn2 := conectarConDn(localdn2)
	clienteDn3, conn3 := conectarConDn(localdn3)
	defer conn2.Close()
	defer conn3.Close()

	aceptadoPorDn2, err2 := clienteDn2.SendPropuestaDistribuida(context.Background(), propuesta1)
	if err2 != nil {
		log.Fatalf("Paso alguna cosa %v", err2)
	}
	aceptadoPorDn3, err3 := clienteDn3.SendPropuestaDistribuida(context.Background(), propuesta2)
	if err3 != nil {
		log.Fatalf("Paso alguna cosa %v", err3)
	}
	// caso bonito donde te aceptan todo
	if aceptadoPorDn3.Okay && aceptadoPorDn2.Okay {
		if chunksDataNode1 != 0 {
			s.envChunks(localdn1, chunksDataNode1)
		}
		if chunksDataNode2 != 0 {
			s.envChunks(localdn2, chunksDataNode2)
		}
		if chunksDataNode3 != 0 {
			s.envChunks(localdn3, chunksDataNode3)
		}
		//Agrawala y luego escribir en log
		s.saladeEsperaDistribuida()
		clienteNn, conexionNn := conectarConNn()
		defer conexionNn.Close() 
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		stream, err := clienteNn.EscribirenLog(ctx)
		if err != nil {
			//Termina la ejecucion del programa por un error de stream
			log.Fatalf("No se pudo obtener el stream %v", err)
		}
		for i := 0; i < len(s.Chunksaescribir); i++ {
			if err := stream.Send(s.Chunksaescribir[i]); err != nil {
				log.Fatalf(".Send(%v) = %v", stream, err)
			}
		}
		//Se limpia el slice para una futura subida de libro de otro cliente
		s.Chunksaescribir = make([]*nn.Logchunk, 0)
		reply, err := stream.CloseAndRecv()
		if err != nil {
			log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
		}
		s.Estado = "RELEASED"
		log.Printf("Se ha cerrado el stream hacia %v, %v", localnn, reply.Mensaje)
	} else {
		//Ahora entramos al caso en que la propuesta original no sirve y DataNode debera considerar otra nueva
		totalChunks := chunksDataNode1 + chunksDataNode2 + chunksDataNode3
		//Generamos un nuevo map que contendra la cantidad de chunks asignados a cada maquina para la nueva propuesta
		propuestaNueva := map[string]int64{
			"maquina3": 0,
			"maquina2": 0,
			"maquina1": 0,
		}
		estadoDeMaquina := map[string]bool{
			"maquina3": aceptadoPorDn3.Okay,
			"maquina2": aceptadoPorDn2.Okay,
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
			if totalChunks >= 1 { //Estando en el DN1 no necesitamos verificar su estado actual porque esta viva si o si
				propuestaNueva["maquina1"]++
				totalChunks--
			}
		}
		if propuestaNueva["maquina1"] != 0 {
			s.envChunks(localdn1, propuestaNueva["maquina1"])
		}
		if propuestaNueva["maquina2"] != 0 {
			s.envChunks(localdn2, propuestaNueva["maquina2"])
		}
		if propuestaNueva["maquina3"] != 0 {
			s.envChunks(localdn3, propuestaNueva["maquina3"])
		}
	}
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
	if err != nil {
		//Caso en que no haya funcionado la conexion con Nn al momento de hacer SendPropuesta
		log.Fatalf("Se cayo el name node, adios %s", err)
	}
	if (confirmacion.Chunksmaquina1 == propuesta.Chunksmaquina1) && (confirmacion.Chunksmaquina2 == propuesta.Chunksmaquina2) && (confirmacion.Chunksmaquina3 == propuesta.Chunksmaquina3) {
		fmt.Println("La propuesta ha sido aceptada.")
		//Eliminar este caso de else if porque no aporta a la solucion
	} else {
		fmt.Println("La propuesta ha sido rechazada y se ha generado una nueva propuesta por el NameNode.")
	}
	//Ahora se procede a enviar (y escribir en disco) los chunks a los datanodes correspondientes.
	if confirmacion.Chunksmaquina1 != 0 {
		// Para pruebas locales
		s.envChunks(localdn1, confirmacion.Chunksmaquina1)
		// s.envChunks(dn1, confirmacion.Chunksmaquina1)
	}
	if confirmacion.Chunksmaquina2 != 0 {
		// Para pruebas locales
		s.envChunks(localdn2, confirmacion.Chunksmaquina2)
		// s.envChunks(dn2, confirmacion.Chunksmaquina2)
	}
	if confirmacion.Chunksmaquina3 != 0 {
		// Para pruebas locales
		s.envChunks(localdn3, confirmacion.Chunksmaquina3)
		// s.envChunks(dn3, confirmacion.Chunksmaquina3)
	}
	//Como ya sabemos que chunks estan repartidos a cada maquina, podemos escribir
	//finalmente en el log. Pero primero debemos consultar al NN si está libre el log

	acceso := &nn.Consultaacceso{
		Ipmaq: localdn1,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	respAcceso, err := clienteNn.Saladeespera(ctx, acceso)
	if respAcceso.Permiso {
		//LOG libre, procedemos a escribir en el DN
		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		stream, err := clienteNn.EscribirenLog(ctx)
		if err != nil {
			//Termina la ejecucion del programa por un error de stream
			log.Fatalf("No se pudo obtener el stream %v", err)
		}
		for i := 0; i < len(s.Chunksaescribir); i++ {
			if err := stream.Send(s.Chunksaescribir[i]); err != nil {
				log.Fatalf(".Send(%v) = %v", stream, err)
			}
		}
		//Se limpia el slice para una futura subida de libro de otro cliente
		s.Chunksaescribir = make([]*nn.Logchunk, 0)
		reply, err := stream.CloseAndRecv()
		if err != nil {
			log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
		}
		log.Printf("Se ha cerrado el stream hacia %v, %v", localnn, reply.Mensaje)
	} else {
		//En caso de llegar al timeout
		fmt.Println("Se acabó el tiempo de espera en cola: ", err)
	}
}

// mustEmbedUnimplementedCourierServiceServer solo se añadio por compatibilidad
// y evitar warnings al compilar
func (s *Server) mustEmbedUnimplementedDataNodeServiceServer() {}

func newServer() *Server {
	s := &Server{
		ChunksRecibidos: make([]*pb.UploadBookRequest, 0),
		Chunksaescribir: make([]*nn.Logchunk, 0),
		Estado:          "RELEASED",
	}
	return s
}

func main() {
	lis, err := net.Listen("tcp", ":9001")
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