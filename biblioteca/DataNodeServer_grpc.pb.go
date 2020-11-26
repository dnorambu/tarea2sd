// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package biblioteca

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DataNodeServiceClient is the client API for DataNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataNodeServiceClient interface {
	UploadBook(ctx context.Context, opts ...grpc.CallOption) (DataNodeService_UploadBookClient, error)
}

type dataNodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataNodeServiceClient(cc grpc.ClientConnInterface) DataNodeServiceClient {
	return &dataNodeServiceClient{cc}
}

func (c *dataNodeServiceClient) UploadBook(ctx context.Context, opts ...grpc.CallOption) (DataNodeService_UploadBookClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataNodeService_serviceDesc.Streams[0], "/DataNodeService/UploadBook", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataNodeServiceUploadBookClient{stream}
	return x, nil
}

type DataNodeService_UploadBookClient interface {
	Send(*UploadBookRequest) error
	CloseAndRecv() (*UploadBookResponse, error)
	grpc.ClientStream
}

type dataNodeServiceUploadBookClient struct {
	grpc.ClientStream
}

func (x *dataNodeServiceUploadBookClient) Send(m *UploadBookRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataNodeServiceUploadBookClient) CloseAndRecv() (*UploadBookResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadBookResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataNodeServiceServer is the server API for DataNodeService service.
// All implementations must embed UnimplementedDataNodeServiceServer
// for forward compatibility
type DataNodeServiceServer interface {
	UploadBook(DataNodeService_UploadBookServer) error
	mustEmbedUnimplementedDataNodeServiceServer()
}

// UnimplementedDataNodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataNodeServiceServer struct {
}

func (UnimplementedDataNodeServiceServer) UploadBook(DataNodeService_UploadBookServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadBook not implemented")
}
func (UnimplementedDataNodeServiceServer) mustEmbedUnimplementedDataNodeServiceServer() {}

// UnsafeDataNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataNodeServiceServer will
// result in compilation errors.
type UnsafeDataNodeServiceServer interface {
	mustEmbedUnimplementedDataNodeServiceServer()
}

func RegisterDataNodeServiceServer(s *grpc.Server, srv DataNodeServiceServer) {
	s.RegisterService(&_DataNodeService_serviceDesc, srv)
}

func _DataNodeService_UploadBook_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataNodeServiceServer).UploadBook(&dataNodeServiceUploadBookServer{stream})
}

type DataNodeService_UploadBookServer interface {
	SendAndClose(*UploadBookResponse) error
	Recv() (*UploadBookRequest, error)
	grpc.ServerStream
}

type dataNodeServiceUploadBookServer struct {
	grpc.ServerStream
}

func (x *dataNodeServiceUploadBookServer) SendAndClose(m *UploadBookResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataNodeServiceUploadBookServer) Recv() (*UploadBookRequest, error) {
	m := new(UploadBookRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DataNodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DataNodeService",
	HandlerType: (*DataNodeServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadBook",
			Handler:       _DataNodeService_UploadBook_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "biblioteca/DataNodeServer.proto",
}
