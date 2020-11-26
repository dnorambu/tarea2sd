package bibliotecaNN

/*
Recordar, este archivo guarda la implementación de las funciones que se utilizarán
junto a las estructuras de datos que necesitamos en NameNode.go
*/

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/net/context"
)

//Server Se declara la estructura del servidor
type Server struct {
	//mutex que protegerá las variables compartidas
	Mu sync.Mutex
}

func (s *Server) 
// mustEmbedUnimplementedCourierServiceServer solo se añadio por compatibilidad
// y evitar warnings al compilar
func (s *Server) mustEmbedUnimplementedCourierServiceServer() {}
