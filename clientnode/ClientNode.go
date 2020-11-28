package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	pb "github.com/dnorambu/tarea2sd/bibliotecadn"
	"google.golang.org/grpc"
)

// Funcion para dividir libros en chunks de 250 KB
func splitFile(cliente pb.DataNodeServiceClient, algoritmo int64, nombreLibro string) {

	fileToBeChunked := "books/" + nombreLibro

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
		fileName := "parte_" + nombreLibro + "_" + strconv.FormatUint(i, 10)
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
		// Problema?
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

// Funcion para iniciar la conexion con algun dataNode
func conectarConDn(maquinas []string) (*pb.DataNodeServiceClient, *grpc.ClientConn) {

	var random int
	for {
		random = rand.Intn(len(maquinas))

		//Para realizar pruebas locales
		conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
		// conn, err := grpc.Dial(maquinas[random], grpc.WithInsecure())

		if err != nil {
			maquinas = append(maquinas[:random], maquinas[random+1:]...)
		} else {
			c := pb.NewDataNodeServiceClient(conn)
			fmt.Println("Conectado a :", maquinas[random])
			return &c, conn
		}
	}
}
func main() {
	maquinas := []string{"10.10.28.140:9000",
		"10.10.28.141:9000",
		"10.10.28.142:9000"}

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
				clienteDn, conexionDn := conectarConDn(maquinas)
				defer conexionDn.Close()
				splitFile(*clienteDn, 0, nombre)
			} else if opcion2 == 1 {
				//Termina proceso de inputs caso 2 (Subir libro con algoritmo distribuido)
			} else {
				fmt.Println("Se introdujo una opcion no valida. Intente de nuevo")
			}
		} else if opcion1 == 1 {
			//Acá se pregunta el nombre del libro que se quiere descargar
			fmt.Println("Indique el nombre del libro que desea descargar:")
			fmt.Scan(&nombre)
			//Termina proceso de inputs caso 3 (Descargar libro)

		} else if opcion1 == 2 {
			fmt.Println("Adios")
			return
		} else {
			fmt.Println("Se introdujo una opcion no valida")
		}

	}
}
