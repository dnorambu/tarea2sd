package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
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
	//Libros ofrecidos al publico que fueron subidos por clientes
	Librosdescargables []string
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

//EscribirenLog usada para escribir la distribución de los chunks en el log. Un DN ejecuta el rpc y
//el NN se encarga de aceptar la solicitud.
func (s *Server) EscribirenLog(stream nn.NameNodeService_EscribirenLogServer) error {
	//Se procesa el primer mensaje para encontrar el nombre del libro y guardarlo en
	//un slice que se mostrara al cliente cuando quiera descargarlos
	s.Mu.Lock()
	defer s.Mu.Unlock()
	chunk, err := stream.Recv()
	nombreLibro := chunk.Nombre[:strings.IndexByte(chunk.Nombre, '_')]
	//Agregamos el nombre del libro que esta siendo subido al archivo "libros.txt",
	//de esta forma facilitamos el proceso de descarga cuando el cliente lo pida.
	file, err := os.OpenFile("libros.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if isError(err) {
		return err
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString(nombreLibro + "\n")
	if isError(err) {
		return err
	}

	// Slice que almacena los mensajes recibidos en el stream
	var Sliceaux = make([]*nn.Logchunk, 0)
	Sliceaux = append(Sliceaux, chunk)
	for {
		chunk, err = stream.Recv()
		if err == io.EOF {
			s.escribir(Sliceaux, nombreLibro)
			//BORRAR
			for i := 0; i < len(s.Librosdescargables); i++ {
				fmt.Println(s.Librosdescargables[i])
			}
			return stream.SendAndClose(&nn.Confirmacion{
				Mensaje: "Distribucion de chunks terminada",
			})
		}
		Sliceaux = append(Sliceaux, chunk)
	}
}
func (s *Server) escribir(chunks []*nn.Logchunk, nombreLibro string) {
	//Obtuvimos el codigo siguiente desde:
	//https://www.golangprograms.com/golang-read-write-create-and-delete-text-file.html
	// ya que solo sirve para abrir archivos de manera correcta

	// Open file using READ & WRITE permission.

	var file, err = os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString(nombreLibro + " " + strconv.Itoa(len(chunks)) + "_partes\n")
	if isError(err) {
		return
	}
	for i := 0; i < len(chunks); i++ {
		_, err = file.WriteString(chunks[i].Nombre + " " + chunks[i].Ipmaquina + "\n")
		if isError(err) {
			return
		}
	}

	fmt.Println("Log modificado exitosamente")
}

//Descargar para tomar los chunks de los diferentes data nodes y permitir que el cliente los pueda descargar
func (s *Server) Descargar(libro *nn.Ubicacionlibro, stream nn.NameNodeService_DescargarServer) error {
	//Primero parsear el log
	var i1 int
	file, err := os.Open("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nombreLibro := scanner.Text()[:strings.IndexByte(scanner.Text(), ' ')]
		if nombreLibro == libro.Nombre {
			aux := strings.Split(scanner.Text(), " ")[1]
			total := aux[:strings.IndexByte(aux, '_')]
			i1, _ = strconv.Atoi(total)
			for i := 0; i < i1; i++ {
				scanner.Scan()
				linea := strings.Split(scanner.Text(), " ")
				response := &nn.Respuesta{
					NombreParte: linea[0],
					Maquina:     linea[1],
				}
				if err := stream.Send(response); err != nil {
					return err
				}
			}
			return nil
		}
	}
	log.Fatalf("No se encontro el libro buscado")
	return nil
}

//Quelibroshay nos muestra que libros hay (xd) disponibles para descargar por el
//cliente
func (s *Server) Quelibroshay(ctx context.Context, emp *nn.Empty) (*nn.Consultalista, error) {
	var temporal *nn.Consultalista
	var libros string
	file, err := os.Open("libros.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		libros = scanner.Text() + "\n" + libros
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	temporal = &nn.Consultalista{
		Listadelibros: libros,
	}
	return temporal, err
}
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
func newServer() *Server {
	s := &Server{
		Librosdescargables: make([]string, 0),
	}
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