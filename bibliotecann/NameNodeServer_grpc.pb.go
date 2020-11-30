// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package bibliotecann

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// NameNodeServiceClient is the client API for NameNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NameNodeServiceClient interface {
	SendPropuesta(ctx context.Context, in *Propuesta, opts ...grpc.CallOption) (*Propuesta, error)
	EscribirenLog(ctx context.Context, opts ...grpc.CallOption) (NameNodeService_EscribirenLogClient, error)
	Quelibroshay(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Consultalista, error)
	Descargar(ctx context.Context, in *Ubicacionlibro, opts ...grpc.CallOption) (NameNodeService_DescargarClient, error)
	Saladeespera(ctx context.Context, in *Consultaacceso, opts ...grpc.CallOption) (*Permisoacceso, error)
}

type nameNodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNameNodeServiceClient(cc grpc.ClientConnInterface) NameNodeServiceClient {
	return &nameNodeServiceClient{cc}
}

func (c *nameNodeServiceClient) SendPropuesta(ctx context.Context, in *Propuesta, opts ...grpc.CallOption) (*Propuesta, error) {
	out := new(Propuesta)
	err := c.cc.Invoke(ctx, "/NameNodeService/SendPropuesta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameNodeServiceClient) EscribirenLog(ctx context.Context, opts ...grpc.CallOption) (NameNodeService_EscribirenLogClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NameNodeService_serviceDesc.Streams[0], "/NameNodeService/EscribirenLog", opts...)
	if err != nil {
		return nil, err
	}
	x := &nameNodeServiceEscribirenLogClient{stream}
	return x, nil
}

type NameNodeService_EscribirenLogClient interface {
	Send(*Logchunk) error
	CloseAndRecv() (*Confirmacion, error)
	grpc.ClientStream
}

type nameNodeServiceEscribirenLogClient struct {
	grpc.ClientStream
}

func (x *nameNodeServiceEscribirenLogClient) Send(m *Logchunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nameNodeServiceEscribirenLogClient) CloseAndRecv() (*Confirmacion, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Confirmacion)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nameNodeServiceClient) Quelibroshay(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Consultalista, error) {
	out := new(Consultalista)
	err := c.cc.Invoke(ctx, "/NameNodeService/Quelibroshay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameNodeServiceClient) Descargar(ctx context.Context, in *Ubicacionlibro, opts ...grpc.CallOption) (NameNodeService_DescargarClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NameNodeService_serviceDesc.Streams[1], "/NameNodeService/Descargar", opts...)
	if err != nil {
		return nil, err
	}
	x := &nameNodeServiceDescargarClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NameNodeService_DescargarClient interface {
	Recv() (*Respuesta, error)
	grpc.ClientStream
}

type nameNodeServiceDescargarClient struct {
	grpc.ClientStream
}

func (x *nameNodeServiceDescargarClient) Recv() (*Respuesta, error) {
	m := new(Respuesta)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nameNodeServiceClient) Saladeespera(ctx context.Context, in *Consultaacceso, opts ...grpc.CallOption) (*Permisoacceso, error) {
	out := new(Permisoacceso)
	err := c.cc.Invoke(ctx, "/NameNodeService/Saladeespera", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NameNodeServiceServer is the server API for NameNodeService service.
// All implementations must embed UnimplementedNameNodeServiceServer
// for forward compatibility
type NameNodeServiceServer interface {
	SendPropuesta(context.Context, *Propuesta) (*Propuesta, error)
	EscribirenLog(NameNodeService_EscribirenLogServer) error
	Quelibroshay(context.Context, *Empty) (*Consultalista, error)
	Descargar(*Ubicacionlibro, NameNodeService_DescargarServer) error
	Saladeespera(context.Context, *Consultaacceso) (*Permisoacceso, error)
	mustEmbedUnimplementedNameNodeServiceServer()
}

// UnimplementedNameNodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNameNodeServiceServer struct {
}

func (UnimplementedNameNodeServiceServer) SendPropuesta(context.Context, *Propuesta) (*Propuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPropuesta not implemented")
}
func (UnimplementedNameNodeServiceServer) EscribirenLog(NameNodeService_EscribirenLogServer) error {
	return status.Errorf(codes.Unimplemented, "method EscribirenLog not implemented")
}
func (UnimplementedNameNodeServiceServer) Quelibroshay(context.Context, *Empty) (*Consultalista, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Quelibroshay not implemented")
}
func (UnimplementedNameNodeServiceServer) Descargar(*Ubicacionlibro, NameNodeService_DescargarServer) error {
	return status.Errorf(codes.Unimplemented, "method Descargar not implemented")
}
func (UnimplementedNameNodeServiceServer) Saladeespera(context.Context, *Consultaacceso) (*Permisoacceso, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Saladeespera not implemented")
}
func (UnimplementedNameNodeServiceServer) mustEmbedUnimplementedNameNodeServiceServer() {}

// UnsafeNameNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NameNodeServiceServer will
// result in compilation errors.
type UnsafeNameNodeServiceServer interface {
	mustEmbedUnimplementedNameNodeServiceServer()
}

func RegisterNameNodeServiceServer(s *grpc.Server, srv NameNodeServiceServer) {
	s.RegisterService(&_NameNodeService_serviceDesc, srv)
}

func _NameNodeService_SendPropuesta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Propuesta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServiceServer).SendPropuesta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NameNodeService/SendPropuesta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServiceServer).SendPropuesta(ctx, req.(*Propuesta))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameNodeService_EscribirenLog_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NameNodeServiceServer).EscribirenLog(&nameNodeServiceEscribirenLogServer{stream})
}

type NameNodeService_EscribirenLogServer interface {
	SendAndClose(*Confirmacion) error
	Recv() (*Logchunk, error)
	grpc.ServerStream
}

type nameNodeServiceEscribirenLogServer struct {
	grpc.ServerStream
}

func (x *nameNodeServiceEscribirenLogServer) SendAndClose(m *Confirmacion) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nameNodeServiceEscribirenLogServer) Recv() (*Logchunk, error) {
	m := new(Logchunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NameNodeService_Quelibroshay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServiceServer).Quelibroshay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NameNodeService/Quelibroshay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServiceServer).Quelibroshay(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameNodeService_Descargar_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Ubicacionlibro)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NameNodeServiceServer).Descargar(m, &nameNodeServiceDescargarServer{stream})
}

type NameNodeService_DescargarServer interface {
	Send(*Respuesta) error
	grpc.ServerStream
}

type nameNodeServiceDescargarServer struct {
	grpc.ServerStream
}

func (x *nameNodeServiceDescargarServer) Send(m *Respuesta) error {
	return x.ServerStream.SendMsg(m)
}

func _NameNodeService_Saladeespera_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Consultaacceso)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServiceServer).Saladeespera(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NameNodeService/Saladeespera",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServiceServer).Saladeespera(ctx, req.(*Consultaacceso))
	}
	return interceptor(ctx, in, info, handler)
}

var _NameNodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "NameNodeService",
	HandlerType: (*NameNodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPropuesta",
			Handler:    _NameNodeService_SendPropuesta_Handler,
		},
		{
			MethodName: "Quelibroshay",
			Handler:    _NameNodeService_Quelibroshay_Handler,
		},
		{
			MethodName: "Saladeespera",
			Handler:    _NameNodeService_Saladeespera_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EscribirenLog",
			Handler:       _NameNodeService_EscribirenLog_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Descargar",
			Handler:       _NameNodeService_Descargar_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bibliotecann/NameNodeServer.proto",
}
