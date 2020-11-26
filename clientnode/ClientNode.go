package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dnorambu/pruebas/courier"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//Funcion para enviar las ordenes de las pymes
func sendPyme(recordFile *os.File, tiempo int64, c courier.CourierServiceClient) {
	// Iniciar el reader
	reader := csv.NewReader(recordFile)
	//Sacamos el Header para no considerarlo en el registro
	reader.Read()

	//Comenzamos a leer las entradas del csv para enviar
	//las ordenes al servidor
	for i := 0; ; i++ {
		registro, err := reader.Read()
		if err == io.EOF {
			fmt.Println("Fin de archivo ", err)
			break
		} else if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		//Se convierten los campos necesarios (valor, prioridad)
		//de tipo string a tipo int64
		valorTemp, _ := strconv.Atoi(registro[2])
		prioridadTemp, _ := strconv.Atoi(registro[5])

		//Se crea una orden de tipo OrdenPyme que será enviada al server
		//con los datos del registro extraido del csv
		orden := courier.OrdenPyme{
			Id:        registro[0],
			Producto:  registro[1],
			Valor:     int64(valorTemp),
			Tienda:    registro[3],
			Destino:   registro[4],
			Prioridad: int64(prioridadTemp),
		}
		//Se recibe el codigo de seguimiento generado por el server
		cod, err := c.EnviarOrdenPyme(context.Background(), &orden)
		if err != nil {
			fmt.Println("Error, error: ", err)
			return
		}
		// Se muestra al cliente el ID de la orden junto al
		// codigo de seguimiento generado por el servidor
		log.Println("ID producto:", registro[0], " Seguimiento:", cod.Cod)
		time.Sleep(time.Duration(tiempo) * time.Second)
	}

	if recordFile.Close() != nil {
		fmt.Println("Error al cerrar")
		return
	}
}

//Funcion para enviar las ordenes de Retail
func sendRetail(recordFile *os.File, tiempo int64, c courier.CourierServiceClient) {
	reader := csv.NewReader(recordFile) // Iniciar el reader

	reader.Read() //Sacamos el Header para no considerarlo en el registro
	for i := 0; ; i++ {
		registro, err := reader.Read()
		if err == io.EOF {
			fmt.Println("Fin de archivo ", err)
			break
		} else if err != nil {
			fmt.Println("Error, error: ", err)
			return
		}

		valorTemp, _ := strconv.Atoi(registro[2])

		orden := courier.OrdenRetail{
			Id:       registro[0],
			Producto: registro[1],
			Valor:    int64(valorTemp),
			Tienda:   registro[3],
			Destino:  registro[4],
		}
		//Puesto que para retail no se genera un codigo de seguimiento
		//solo se almacena un posible error luego de enviar la orden
		_, err = c.EnviarOrdenRetail(context.Background(), &orden)
		if err != nil {
			fmt.Println("Error, error: ", err)
			return
		}
		fmt.Println("Orden:", registro[0], " entregada a logistica")
		time.Sleep(time.Duration(tiempo) * time.Second)
	}
	if recordFile.Close() != nil {
		fmt.Println("Error al cerrar")
		return
	}
	fmt.Println("Todas las ordenes fueron enviadas\nTermina la sesion")
}

// Funcion para dividir libros en chunks de 250 KB
func splitFile (name string){
	fileToBeChunked := "./somebigfile"

	file, err := os.Open(fileToBeChunked)

	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()

	const fileChunk = 1 * (1 << 20) // 1 MB, change this to your requirement

	// calculate total number of parts the file will be chunked into

	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

	for i := uint64(0); i < totalPartsNum; i++ {

			partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
			partBuffer := make([]byte, partSize)

			file.Read(partBuffer)

			// write to disk
			fileName := "somebigfile_" + strconv.FormatUint(i, 10)
			_, err := os.Create(fileName)

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
	fmt.Println("#-#-#-#-#-#-#-# PRESTIGIO - EXPRESS #-#-#-#-#-#-#-#-#")

	//Preguntamos el tipo de cliente y los segundos entre ordenes
	var tipo, tiempo int64
	var file string
	/*
		Aquí se ingresa el NUMERO de la opcion
		Ingresar 0 para retail
		Ingresar 1 para pyme
	*/
	fmt.Println("Indique su tipo (numero) de cliente:\n0 Retail\n1 Pyme")
	fmt.Scan(&tipo)

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
