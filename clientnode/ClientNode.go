package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/dnorambu/pruebas/courier"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Funcion para dividir libros en chunks de 250 KB
func splitFile() {

	fileToBeChunked := "./somebigfile"

	file, err := os.Open(fileToBeChunked)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()

	const fileChunk = 250000 // 250 KB, change this to your requirement

	// calculate total number of parts the file will be chunked into

	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

	for i := uint64(0); i < totalPartsNum; i++ {
		// O el chunk pesa 250KB o es el último y puede que pese menos
		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)

		//
		// // write to disk
		fileName := "somebigfile_" + strconv.FormatUint(i, 10)
		// _, err := os.Create(fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// write/save buffer to disk
		ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)

		fmt.Println("Split to : ", fileName)
	}
}

func main() {
	/*
		***************************
		Generamos la conexión al servidor
		***************************
	*/
	var conn *grpc.ClientConn

	//Para realizar pruebas locales
	//conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())

	// Descomentar para testear en MV 3 (logistica)
	conn, err := grpc.Dial("10.10.28.141:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := courier.NewCourierServiceClient(conn)

	/*
		****************************
			Conexión establecida
		****************************
	*/

	fmt.Println("#-#-#-#-#-#-#-# Bienvenido #-#-#-#-#-#-#-#-#")
	
	var opcion1, opcion2, opcion3 int64
	var nombre string
	//Preguntamos si desea subir o descargar
	/*
		Aquí se ingresa el NUMERO de la opcion
		Ingresar 0 para Subir
		Ingresar 1 para Descargar
	*/
	fmt.Println("Indique la opcion (numero) que desea realizar:\n0 Subir libro\n1 Descargar libro")
	fmt.Scan(&opcion1)
	
	if opcion1 == 0 {
		//Acá se pregunta el nombre del libro que se quiere subir
		fmt.Println("Indique el nombre del libro que desea subir:\n")
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
		} else if opcion2 == 1 {
			//Termina proceso de inputs caso 2 (Subir libro con algoritmo distribuido)
		} else {
			fmt.Println("Se introdujo una opcion no valida")
			return
		}
	} else if opcion1 == 1 {
		//Acá se pregunta el nombre del libro que se quiere descargar
		fmt.Println("Indique el nombre del libro que desea descargar:\n")
		fmt.Scan(&nombre)
		//Termina proceso de inputs caso 3 (Descargar libro)
		
	} else {
		fmt.Println("Se introdujo una opcion no valida")
		return
	}

	fmt.Println("Ingrese el tiempo entre ordenes (en segundos): ")
	fmt.Scan(&tiempo)

	fmt.Println("Ingrese el nombre del archivo (ej: orden.csv) ")
	fmt.Scan(&file)

	// Abrimos el csv correspondiente según el tipo de cliente
	// Los archivos se guardan en arch/
	var recordFile *os.File
	if tipo == 0 {
		recordFile, err = os.Open("arch/" + file)
		if err != nil { //Chequear  si hubo error
			fmt.Println("Error al abrir el archivo: ", err)
			return
		}
		sendRetail(recordFile, tiempo, c)
		//Se termina la ejecucion pues retail no pide seguimiento
		return
	} else if tipo == 1 {
		if err != nil { //Chequear  si hubo error
			fmt.Println("Error al abrir el archivo: ", err)
			return
		}
		recordFile, err = os.Open("arch/" + file)
		sendPyme(recordFile, tiempo, c)
	} else {
		fmt.Println("Se introdujo un tipo no valido de cliente")
		return
	}

	//Ciclo para pedir seguimientos
	var opcion, codigoSeg int64
	fmt.Println("Ingrese la accion (numero) que desee: ")
	for {
		fmt.Println("1 Para solicitar seguimiento\n2 Para terminar la ejecucion")
		fmt.Scan(&opcion)
		if opcion == 1 {
			fmt.Println("Ingrese el codigo de seguimiento: ")
			fmt.Scan(&codigoSeg)
			SolicitudSeg := courier.Codigo{Cod: codigoSeg}
			estado, err := c.Seguimiento(context.Background(), &SolicitudSeg)
			if err != nil {
				fmt.Println("El codigo no esta registrado en el servidor: ", err)
			}
			fmt.Println("Estado de la orden:", estado.Body)
		} else if opcion == 2 {
			fmt.Println("Adios")
			break
		} else {
			fmt.Println("Opcion invalida, intente de nuevo")
		}
	}
}
