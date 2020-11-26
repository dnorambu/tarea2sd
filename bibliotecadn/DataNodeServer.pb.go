// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: bibliotecadn/DataNodeServer.proto

package bibliotecadn

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Mensajes necesarios para implementar la red
type UploadBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre    string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3" json:"chunk_data,omitempty"`
}

func (x *UploadBookRequest) Reset() {
	*x = UploadBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bibliotecadn_DataNodeServer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadBookRequest) ProtoMessage() {}

func (x *UploadBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bibliotecadn_DataNodeServer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadBookRequest.ProtoReflect.Descriptor instead.
func (*UploadBookRequest) Descriptor() ([]byte, []int) {
	return file_bibliotecadn_DataNodeServer_proto_rawDescGZIP(), []int{0}
}

func (x *UploadBookRequest) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *UploadBookRequest) GetChunkData() []byte {
	if x != nil {
		return x.ChunkData
	}
	return nil
}

type UploadBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Respuesta string `protobuf:"bytes,1,opt,name=respuesta,proto3" json:"respuesta,omitempty"`
}

func (x *UploadBookResponse) Reset() {
	*x = UploadBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bibliotecadn_DataNodeServer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadBookResponse) ProtoMessage() {}

func (x *UploadBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bibliotecadn_DataNodeServer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadBookResponse.ProtoReflect.Descriptor instead.
func (*UploadBookResponse) Descriptor() ([]byte, []int) {
	return file_bibliotecadn_DataNodeServer_proto_rawDescGZIP(), []int{1}
}

func (x *UploadBookResponse) GetRespuesta() string {
	if x != nil {
		return x.Respuesta
	}
	return ""
}

var File_bibliotecadn_DataNodeServer_proto protoreflect.FileDescriptor

var file_bibliotecadn_DataNodeServer_proto_rawDesc = []byte{
	0x0a, 0x21, 0x62, 0x69, 0x62, 0x6c, 0x69, 0x6f, 0x74, 0x65, 0x63, 0x61, 0x64, 0x6e, 0x2f, 0x44,
	0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x32, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x70, 0x75, 0x65,
	0x73, 0x74, 0x61, 0x32, 0x4c, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x12, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x6e, 0x6f, 0x72, 0x61, 0x6d, 0x62, 0x75, 0x2f, 0x74, 0x61, 0x72, 0x65, 0x61, 0x32, 0x73,
	0x64, 0x2f, 0x62, 0x69, 0x62, 0x6c, 0x69, 0x6f, 0x74, 0x65, 0x63, 0x61, 0x64, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bibliotecadn_DataNodeServer_proto_rawDescOnce sync.Once
	file_bibliotecadn_DataNodeServer_proto_rawDescData = file_bibliotecadn_DataNodeServer_proto_rawDesc
)

func file_bibliotecadn_DataNodeServer_proto_rawDescGZIP() []byte {
	file_bibliotecadn_DataNodeServer_proto_rawDescOnce.Do(func() {
		file_bibliotecadn_DataNodeServer_proto_rawDescData = protoimpl.X.CompressGZIP(file_bibliotecadn_DataNodeServer_proto_rawDescData)
	})
	return file_bibliotecadn_DataNodeServer_proto_rawDescData
}

var file_bibliotecadn_DataNodeServer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bibliotecadn_DataNodeServer_proto_goTypes = []interface{}{
	(*UploadBookRequest)(nil),  // 0: UploadBookRequest
	(*UploadBookResponse)(nil), // 1: UploadBookResponse
}
var file_bibliotecadn_DataNodeServer_proto_depIdxs = []int32{
	0, // 0: DataNodeService.UploadBook:input_type -> UploadBookRequest
	1, // 1: DataNodeService.UploadBook:output_type -> UploadBookResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bibliotecadn_DataNodeServer_proto_init() }
func file_bibliotecadn_DataNodeServer_proto_init() {
	if File_bibliotecadn_DataNodeServer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bibliotecadn_DataNodeServer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadBookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bibliotecadn_DataNodeServer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadBookResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bibliotecadn_DataNodeServer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bibliotecadn_DataNodeServer_proto_goTypes,
		DependencyIndexes: file_bibliotecadn_DataNodeServer_proto_depIdxs,
		MessageInfos:      file_bibliotecadn_DataNodeServer_proto_msgTypes,
	}.Build()
	File_bibliotecadn_DataNodeServer_proto = out.File
	file_bibliotecadn_DataNodeServer_proto_rawDesc = nil
	file_bibliotecadn_DataNodeServer_proto_goTypes = nil
	file_bibliotecadn_DataNodeServer_proto_depIdxs = nil
}
