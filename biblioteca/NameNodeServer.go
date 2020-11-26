package biblioteca

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
	//Mapa para guardar los pares "CodigoSeguimiento-Estado" de cada paquete PYME
	MapaSeguimiento map[int64]string

	//Mapa para guardar los pares "ID-Estado" de cada paquete RETAIL
	MapaSegRetail map[string]string

	//Mapa para guardar referencias entre "ID-CodigoSeguimiento"
	//Sirve como man in the middle para mantener consistencia en el estado del paquete
	MapaIntermedio map[string]int64
	//Slice que tendrá las órdenes de pymes que lleguen al servidor
	OrdenesP []*OrdenPyme

	//Slice que tendrá las órdens de retail que lleguen al servidor
	OrdenesR []*OrdenRetail
	//mutex que protegerá las variables compartidas
	Mu sync.Mutex

	//Slices que representan a las 3 colas
	ColaNormal, ColaPrioritario, ColaRetail []*Paquete

	//Slice que guardará las entregas de camiones.
	EntregasC []*Entrega
}

//Seguimiento Esta es la prueba inicial de mi función para devolver código de seguimiento
func (s *Server) Seguimiento(ctx context.Context, seguim *Codigo) (*Mensaje, error) {
	//Buscamos el codigo en el mapa de seguimiento
	i, ok := s.MapaSeguimiento[seguim.Cod]
	if !ok {
		err := errors.New("Error: missing key in MapaSeguimiento")
		return &Mensaje{Body: "Ese codigo no está registrado"}, err
	}
	return &Mensaje{Body: i}, nil
}

//EnviarOrdenPyme se recibe una orden, se guarda en el slice "ordenes"
//
func (s *Server) EnviarOrdenPyme(ctx context.Context, ordenP *OrdenPyme) (*Codigo, error) {

	//Implementacion de una fuente de numeros aleatorios
	//para dar códigos diferentes en cada instancia
	seed := rand.NewSource(time.Now().UnixNano())
	source := rand.New(seed)

	var num int64
	var err error
	for {
		// Se usa un set de numeros (Cardinalidad 1x10^13) para sacar codigos de seguimiento
		num = int64(source.Intn(10000000000000))

		//Manejar acceso concurrente al mapa para evitar errores de
		//sobre escritura o inconsistencias
		s.Mu.Lock()
		_, err := s.MapaSeguimiento[num]
		s.MapaIntermedio[ordenP.Id] = num
		s.Mu.Unlock()
		if !err {
			s.Mu.Lock()
			s.MapaSeguimiento[num] = "En bodega"
			s.Mu.Unlock()
			break
		}
	}
	//Definir el tipo según la prioridad
	var prioridad string
	if ordenP.Prioridad == 1 {
		prioridad = "prioritario"
	} else if ordenP.Prioridad == 0 {
		prioridad = "normal"
	} else {
		fmt.Println("No coincide con ninguna prioridad")
	}

	Pkg := Paquete{
		Id:      ordenP.Id,
		Tipo:    prioridad,
		Valor:   ordenP.Valor,
		Tienda:  ordenP.Tienda,
		Estado:  s.MapaSeguimiento[num],
		Destino: ordenP.Destino,
	}

	if prioridad == "prioritario" {
		s.ColaPrioritario = append(s.ColaPrioritario, &Pkg)
	} else {
		s.ColaNormal = append(s.ColaNormal, &Pkg)
	}

	s.OrdenesP = append(s.OrdenesP, ordenP)

	return &Codigo{Cod: num}, err
}

//EnviarOrdenRetail para manejar las ordenes desde retail hasta servidor
func (s *Server) EnviarOrdenRetail(ctx context.Context, ordenR *OrdenRetail) (*Empty, error) {
	s.OrdenesR = append(s.OrdenesR, ordenR)
	var err error
	dummy := &Empty{}
	Pkg := Paquete{
		Id:      ordenR.Id,
		Tipo:    "retail",
		Valor:   ordenR.Valor,
		Tienda:  ordenR.Tienda,
		Estado:  "En bodega",
		Destino: ordenR.Destino,
	}
	s.Mu.Lock()
	s.ColaRetail = append(s.ColaRetail, &Pkg)
	s.Mu.Unlock()
	return dummy, err
}

//PedirRetail para que camion obtenga paquete retail.
func (s *Server) PedirRetail(ctx context.Context, msj *Mensaje) (*Paquete, error) {
	var err error
	s.Mu.Lock()
	if len(s.ColaRetail) >= 1 {
		nuevopaquete := s.ColaRetail[0]
		s.ColaRetail = s.ColaRetail[1:]
		nuevopaquete.Estado = "En Camino"
		s.Mu.Unlock()
		return nuevopaquete, err
	}
	s.Mu.Unlock()
	return &Paquete{Id: ""}, err
}

//PedirPrioritario para que camion obtenga paquete prioritario.
func (s *Server) PedirPrioritario(ctx context.Context, msj *Mensaje) (*Paquete, error) {
	var err error
	s.Mu.Lock()
	if len(s.ColaPrioritario) >= 1 {
		nuevopaquete := s.ColaPrioritario[0]
		s.ColaPrioritario = s.ColaPrioritario[1:]
		nuevopaquete.Estado = "En Camino"
		//Interaccion entre mapas
		idAux := nuevopaquete.Id
		codigoAux, ok := s.MapaIntermedio[idAux]
		if !ok {
			fmt.Println("No está la llave pedida")
			s.Mu.Unlock()
			return &Paquete{}, err
		}
		s.MapaSeguimiento[codigoAux] = "En Camino"
		s.Mu.Unlock()
		return nuevopaquete, err
	}
	s.Mu.Unlock()
	return &Paquete{Id: ""}, err
}

//PedirNormal para que camion obtenga paquete normal.
func (s *Server) PedirNormal(ctx context.Context, msj *Mensaje) (*Paquete, error) {
	var err error
	s.Mu.Lock()
	if len(s.ColaNormal) >= 1 {
		nuevopaquete := s.ColaNormal[0]
		s.ColaNormal = s.ColaNormal[1:]
		nuevopaquete.Estado = "En Camino"
		//Interaccion entre mapas
		idAux := nuevopaquete.Id
		codigoAux, ok := s.MapaIntermedio[idAux]
		if !ok {
			fmt.Println("No está la llave pedida")
			s.Mu.Unlock()
			return &Paquete{}, err
		}
		s.MapaSeguimiento[codigoAux] = "En Camino"
		s.Mu.Unlock()
		return nuevopaquete, err
	}
	s.Mu.Unlock()
	return &Paquete{Id: ""}, err
}

//ResultadoEntrega para que camión entregue resultados.
func (s *Server) ResultadoEntrega(ctx context.Context, ent *Entrega) (*Empty, error) {
	var err error
	s.Mu.Lock()
	dummy := &Empty{}
	s.EntregasC = append(s.EntregasC, ent)
	//Interaccion entre mapas
	idAux := ent.Id
	codigoAux, ok := s.MapaIntermedio[idAux]
	if !ok {
		fmt.Println("Entrega completa: ", ent)
		s.Mu.Unlock()
		return dummy, err
	}
	s.MapaSeguimiento[codigoAux] = ent.Estado
	s.Mu.Unlock()
	fmt.Println("Entrega completa: ", ent)
	return dummy, err
}

// mustEmbedUnimplementedCourierServiceServer solo se añadio por compatibilidad
// y evitar warnings al compilar
func (s *Server) mustEmbedUnimplementedCourierServiceServer() {}
