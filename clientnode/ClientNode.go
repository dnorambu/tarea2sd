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
		fmt.Println("Se reconocio al algoritmo distribuido, se crea el stream")
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

		reply, err := stream.CloseAndRecv()
		if err != nil {
			log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
		}
		log.Printf("Se ha cerrado el stream %v", reply)
	}
}

//DescargarPartes se explica solo, pero ... es para descargar los chunks
func DescargarPartes(archivos *[]byte, maqSlice []string, c pb.DataNodeServiceClient) {
	archivos2 := *archivos
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := c.DownloadBook(ctx) //hay que definir el rpc (descomentar)
	if err != nil {
		fmt.Printf("Ciertos chunks del libro que quieres bajar estan en un DataNode caido: %v", err)
		return
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

			*archivos = append(archivos2, parte.Chunkdata...)
			archivos2 = *archivos
		}
	}()
	for _, maquina := range maqSlice {
		sendMaquina := &pb.PartName{
			Nombre: maquina,
		}
		if err := stream.Send(sendMaquina); err != nil {
			log.Fatalf("Error al enviar el nombre de una parte: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
}

// Funcion para iniciar la conexion con algun dataNode al azar y que no esté caido
func conectarConDnAleatorio() (*pb.DataNodeServiceClient, *grpc.ClientConn) {
	maquinas := []string{dn1, dn2, dn3}

	var random int
	for {
		rand.Seed(time.Now().Unix())
		random = rand.Intn(len(maquinas))
		conn, err := grpc.Dial(maquinas[random], grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*60))

		if err != nil {
			//Esta linea borra la maquina caida de los posibles candidatos a conectarse
			maquinas = append(maquinas[:random], maquinas[random+1:]...)
		} else {
			//Nodo vivo, pero necesitamos saber si su lista de chunks está vacía para poder hacer la conexion
			c := pb.NewDataNodeServiceClient(conn)
			listaVacia, _ := c.ListaVacia(context.Background(), &pb.Empty{})
			if listaVacia.Okay {
				fmt.Println("Conectado a :", maquinas[random])
				return &c, conn
			}
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
	fmt.Println("Conectado a DataNode: " + datanode)
	return c, conn
}

func conectarConNn() (nn.NameNodeServiceClient, *grpc.ClientConn) {
	conn, err := grpc.Dial(nameNode, grpc.WithInsecure())
	if err != nil {
		//OJO, si el NN no esta funcionando, el programa terminara la ejecucion
		log.Fatalf("Se cayo el name node, adios: %s", err)
	}
	c := nn.NewNameNodeServiceClient(conn)
	fmt.Println("Conectado a NameNode:", nameNode)
	return c, conn
}
func main() {

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
				clienteDn, conexionDn := conectarConDnAleatorio()
				defer conexionDn.Close()
				splitFile(*clienteDn, 0, nombre)
			} else if opcion2 == 1 {
				//Termina proceso de inputs caso 2 (Subir libro con algoritmo distribuido)
				clienteDn, conexionDn := conectarConDnAleatorio()
				defer conexionDn.Close()
				splitFile(*clienteDn, 1, nombre)
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
					if value == dn1 {
						maq1Slice = append(maq1Slice, key)
					}
					if value == dn2 {
						maq2Slice = append(maq2Slice, key)
					}
					if value == dn3 {
						maq3Slice = append(maq3Slice, key)
					}
				}
				var bytesLibro = make([]byte, 0)
				if len(maq1Slice) != 0 {
					clienteDn, conexionDn := conectarConDn(dn1)
					defer conexionDn.Close()
					DescargarPartes(&bytesLibro, maq1Slice, clienteDn)
				}
				if len(maq2Slice) != 0 {
					clienteDn, conexionDn := conectarConDn(dn2)
					defer conexionDn.Close()
					DescargarPartes(&bytesLibro, maq2Slice, clienteDn)
				}
				if len(maq3Slice) != 0 {
					clienteDn, conexionDn := conectarConDn(dn3)
					defer conexionDn.Close()
					DescargarPartes(&bytesLibro, maq3Slice, clienteDn)
				}
				nuevoLibro := nombre + ".pdf"
				_, err = os.Create("descargas/" + nuevoLibro)
				if err != nil {
					fmt.Println("No se pudo crear el nuevo archivo", err)
					os.Exit(1)
				}
				archivo, err := os.OpenFile("descargas/"+nuevoLibro, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
				if err != nil {
					fmt.Println("No se pudo abrir nuevo archivo en modo APPEND", err)
					os.Exit(1)
				}
				_, err = archivo.Write(bytesLibro)
				if err != nil {
					fmt.Println("No se pudieron escribir los bytes en el archivo", err)
					os.Exit(1)
				}
				archivo.Sync()
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
