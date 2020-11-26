package bibliotecadn

/*
Recordar, este archivo guarda la implementación de las funciones que se utilizarán
junto a las estructuras de datos que necesitamos en NameNode.go
*/

import (
	"sync"
)

//Server Se declara la estructura del servidor
type Server struct {
	//mutex que protegerá las variables compartidas
	Mu sync.Mutex
}

func (s *Server) UploadBook(stream) (DataNodeService_UploadBookClient, error)
func (server *LaptopServer) UploadImage(stream pb.LaptopService_UploadImageServer) error {
	return nil
}

// mustEmbedUnimplementedCourierServiceServer solo se añadio por compatibilidad
// y evitar warnings al compilar
func (s *Server) mustEmbedUnimplementedDataNodeServiceServer() {}
