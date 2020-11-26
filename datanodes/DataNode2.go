package main

import (
	"log"

	"github.com/dnorambu/pruebas/courier"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	//Para testear en local
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())

	//Para testear en MV
	// conn, err := grpc.Dial("10.10.28.141:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := courier.NewCourierServiceClient(conn)

}
