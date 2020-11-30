package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	pb "github.com/dnorambu/tarea2sd/bibliotecadn"
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

// Funcion para dividir libros en chunks de 250 KB
func splitFile(cliente pb.DataNodeServiceClient, algoritmo int64, nombreLibro string) {

	fileToBeChunked := "books/" + nombreLibro + ".pdf"

	file, err := os.Open(fileToBeChunked)

	if err != nil {
		fmt.Println("No se pudo dividir el archivo en chunks :c", err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()

	const fileChunk = 250000 // 250 KB, change this to your requirement

	// calculate total number of parts the file will be chunked into

	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)
	var sliceDeChunks []*pb.UploadBookRequest
	for i := uint64(0); i < totalPartsNum; i++ {

		// O el chunk pesa 250KB o es el último y puede que pese menos
		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)

		//
		// // write to disk
		fileName := nombreLibro + "_parte_" + strconv.FormatUint(i, 10)
		// _, err := os.Create(fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// write/save buffer to disk
		//ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)
		chunk := &pb.UploadBookRequest{
			Nombre:    fileName,
			Chunkdata: partBuffer,
		}

		sliceDeChunks = append(sliceDeChunks, chunk)
		fmt.Println("Largo del sliceDeChunks", len(sliceDeChunks))
	}
	// ahora podemos usar el stream para enviar los chunks
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//Se define el stream, el codigo se repite para solucionar problemas
	//de variables definidas pero no usadas (segun Golang)

	if algoritmo == 0 {
		fmt.Println("Se reconocio al algoritmo centralizado, se crea el stream")
		stream, err := cliente.UploadBookCentralizado(ctx)
		if err != nil {
			//Termina la ejecucion del programa por un error de stream
			log.Fatalf("No se pudo obtener el stream %v", err)
		}

		for _, chunk := range sliceDeChunks {
			if err := stream.Send(chunk); err != nil {
				log.Fatalf("%v.Send(%v) = %v", stream, chunk, err)
			}
		}

		reply, err := stream.CloseAndRecv()
		if err != nil {
			log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
		}
		log.Printf("Se ha cerrado el stream %v", reply.Respuesta)
	} else {
		stream, err := cliente.UploadBookDistribuido(ctx)
		if err != nil {
			//Termina la ejecucion del programa por un error de stream
			log.Fatalf("No se pudo obtener el stream %v", err)
		}
		for _, chunk := range sliceDeChunks {
			if err := stream.Send(chunk); err != nil {
				log.Fatalf("%v.Send(%v) = %v", stream, chunk, err)
			}
		}
		fmt.Println("Sali del for con Send()")
		reply, err := stream.CloseAndRecv()
		if err != nil {
			log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
		}
		log.Printf("Se ha cerrado el stream %v", reply)
	}
}

//DescargarPartes es para inciar el proceso de descargas de partes
func DescargarPartes(maqSlice []string, c pb.DataNodeServiceClient){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := c.DownloadBook(ctx) //hay que definir el rpc (descomentar)
	if err != nil {
		log.Fatalf("%v.DownloadBook(_) = _, %v", c, err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			parte, err := stream.Recv()
			if err == io.EOF {
				// Lectura lista.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Error al recibir una parte: %v", err)
			}
			log.Printf("Se obtuvo una parte del DataNode exitosamente.")
			//ACA HAY QUE PROCESAR LA PARTE RECIBIDA DE ALGUNA FORMA.
		}
	}()
	for _, maqSlice := range maqSlice {
		if err := stream.Send(maqSlice); err != nil {
			log.Fatalf("Error al enviar el nombre de una parte: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
}


// Funcion para iniciar la conexion con algun dataNode
func conectarConDnDistr(maquinas []string) (*pb.DataNodeServiceClient, *grpc.ClientConn) {

	var random int
	for {
		rand.Seed(time.Now().Unix())
		random = rand.Intn(len(maquinas))

		//Para realizar pruebas locales
		conn, err := grpc.Dial(maquinas[random], grpc.WithInsecure())

		if err != nil {
			maquinas = append(maquinas[:random], maquinas[random+1:]...)
		} else {
			c := pb.NewDataNodeServiceClient(conn)
			fmt.Println("Conectado a :", maquinas[random])
			return &c, conn
		}
	}
}

func conectarConDn(datanode string) (pb.DataNodeServiceClient, *grpc.ClientConn) {

	conn, err := grpc.Dial(datanode, grpc.WithInsecure())

	if err != nil {
		//OJO, si el DN no esta funcionando, el programa terminara la ejecucion
		log.Fatalf("No esta disponible el DataNode con partes a descargar: %s", err)
	}
	c := pb.NewDataNodeServiceClient(conn)
	fmt.Println("Conectado a DataNode: "+ datanode)
	return c, conn
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
func main() {
	// maquinas := []string{"10.10.28.140:9000",
	// 	"10.10.28.141:9000",
	// 	"10.10.28.142:9000"}

	//Para probar en local
	maquinas := []string{"localhost:9001",
		"localhost:9002",
		"localhost:9003"}

	fmt.Println("#-#-#-#-#-#-#-# Bienvenido #-#-#-#-#-#-#-#-#")

	var opcion1, opcion2 int64
	var nombre string
	//Preguntamos si desea subir o descargar
	/*
		Aquí se ingresa el NUMERO de la opcion
		Ingresar 0 para Subir
		Ingresar 1 para Descargar
	*/

	for {
		fmt.Println("Indique la opcion (numero) que desea realizar:\n0 Subir libro\n1 Descargar libro\n2 Salir")
		fmt.Scan(&opcion1)
		if opcion1 == 0 {
			//Acá se pregunta el nombre del libro que se quiere subir
			fmt.Println("Indique el nombre del libro que desea subir:")
			fmt.Scan(&nombre)
			/*
				Acá se pregunta el qué tipo de algoritmo se va a utilizar al subir el libro
				Ingresar 0 para Centralizado
				Ingresar 1 para Distribuido
			*/
			fmt.Println("Indique la opcion (numero) para tipo de algoritmo:\n0 Centralizado\n1 Distribuido")
			fmt.Scan(&opcion2)
			//Se separan para luego hacer el llamado a función correspondiente para cada algoritmo.
			if opcion2 == 0 {
				//Termina proceso de inputs caso 1 (Subir libro con algoritmo centralizado)
				clienteDnD, conexionDnD := conectarConDnDistr(maquinas)
				defer conexionDnD.Close()
				splitFile(*clienteDnD, 0, nombre)
			} else if opcion2 == 1 {
				//Termina proceso de inputs caso 2 (Subir libro con algoritmo distribuido)
			} else {
				fmt.Println("Se introdujo una opcion no valida. Intente de nuevo")
			}
		} else if opcion1 == 1 {
			//Listamos los libros descargables mediante grpc con NameNode
			clienteNn, conexionNn := conectarConNn()
			defer conexionNn.Close()
			librosDescargables, err := clienteNn.Quelibroshay(context.Background(), &nn.Empty{})
			if err != nil {
				log.Println("No se pudo leer el archivo libros.txt del name node", err)
			}
			if len(librosDescargables.Listadelibros) == 0 {
				fmt.Println("No hay libros disponibles para su descarga en este momento")
			} else {
				fmt.Println("Libros disponibles:\n", librosDescargables.Listadelibros)
				//Acá se pregunta el nombre del libro que se quiere descargar
				fmt.Println("Indique el nombre del libro que desea descargar:")
				fmt.Scan(&nombre)
				ul := &nn.Ubicacionlibro{
					Nombre: nombre,
				}
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()
				var mapaTemp = make(map[string]string)
				stream, err := clienteNn.Descargar(ctx, ul)
				if err != nil {
					//Termina la ejecucion del programa por un error de stream
					log.Fatalf("No se pudo obtener el stream %v", err)
				}
				for {
					parteLibro, err := stream.Recv()
					if err == io.EOF {
						break
					}
					if err != nil {
						log.Fatalf("%v.ListFeatures() = , %v", clienteNn, err)
					}
					//Este mapa nos guardará los chunks para pedirlos a los diferentes Dn
					mapaTemp[parteLibro.NombreParte] = parteLibro.Maquina
				}
				fmt.Println("MAPA: ", mapaTemp)
				var maq1Slice []string
				var maq2Slice []string
				var maq3Slice []string
				for key, value := range mapaTemp {
					if value == localdn1 {
						maq1Slice = append(maq1Slice, key)
					}
					if value == localdn2 {
						maq2Slice = append(maq2Slice, key)
					}
					if value == localdn3 {
						maq3Slice = append(maq3Slice, key)
					}
				}
				if len(maq1Slice) != 0{
					clienteDn, conexionDn := conectarConDn(localdn1)
					defer conexionDn.Close()
					DescargarPartes(maq1Slice, clienteDn)
				}
				if len(maq2Slice) != 0{
					clienteDn, conexionDn := conectarConDn(localdn2)
					defer conexionDn.Close()
					DescargarPartes(maq2Slice, clienteDn)
				}
				if len(maq3Slice) != 0{
					clienteDn, conexionDn := conectarConDn(localdn3)
					defer conexionDn.Close()
					DescargarPartes(maq3Slice, clienteDn)
				}
				//Termina proceso de inputs caso 3 (Descargar libro)
			}
		} else if opcion1 == 2 {
			fmt.Println("Adios")
			return
		} else {
			fmt.Println("Se introdujo una opcion no valida")
		}
	}
}