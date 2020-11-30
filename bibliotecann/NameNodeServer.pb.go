// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: bibliotecann/NameNodeServer.proto

package bibliotecann

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

type Propuesta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunksmaquina1 int64 `protobuf:"varint,1,opt,name=chunksmaquina1,proto3" json:"chunksmaquina1,omitempty"`
	Chunksmaquina2 int64 `protobuf:"varint,2,opt,name=chunksmaquina2,proto3" json:"chunksmaquina2,omitempty"`
	Chunksmaquina3 int64 `protobuf:"varint,3,opt,name=chunksmaquina3,proto3" json:"chunksmaquina3,omitempty"`
}

func (x *Propuesta) Reset() {
	*x = Propuesta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bibliotecann_NameNodeServer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Propuesta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Propuesta) ProtoMessage() {}

func (x *Propuesta) ProtoReflect() protoreflect.Message {
	mi := &file_bibliotecann_NameNodeServer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Propuesta.ProtoReflect.Descriptor instead.
func (*Propuesta) Descriptor() ([]byte, []int) {
	return file_bibliotecann_NameNodeServer_proto_rawDescGZIP(), []int{0}
}

func (x *Propuesta) GetChunksmaquina1() int64 {
	if x != nil {
		return x.Chunksmaquina1
	}
	return 0
}

func (x *Propuesta) GetChunksmaquina2() int64 {
	if x != nil {
		return x.Chunksmaquina2
	}
	return 0
}

func (x *Propuesta) GetChunksmaquina3() int64 {
	if x != nil {
		return x.Chunksmaquina3
	}
	return 0
}

//Logchunk es un mensaje que se pasara desde el DataNode que desee escribir
//en el log, hacia el NameNode
type Logchunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre    string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Ipmaquina string `protobuf:"bytes,2,opt,name=ipmaquina,proto3" json:"ipmaquina,omitempty"`
}

func (x *Logchunk) Reset() {
	*x = Logchunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bibliotecann_NameNodeServer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logchunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logchunk) ProtoMessage() {}

func (x *Logchunk) ProtoReflect() protoreflect.Message {
	mi := &file_bibliotecann_NameNodeServer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logchunk.ProtoReflect.Descriptor instead.
func (*Logchunk) Descriptor() ([]byte, []int) {
	return file_bibliotecann_NameNodeServer_proto_rawDescGZIP(), []int{1}
}

func (x *Logchunk) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Logchunk) GetIpmaquina() string {
	if x != nil {
		return x.Ipmaquina
	}
	return ""
}

type Confirmacion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mensaje string `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
}

func (x *Confirmacion) Reset() {
	*x = Confirmacion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bibliotecann_NameNodeServer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Confirmacion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Confirmacion) ProtoMessage() {}

func (x *Confirmacion) ProtoReflect() protoreflect.Message {
	mi := &file_bibliotecann_NameNodeServer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Confirmacion.ProtoReflect.Descriptor instead.
func (*Confirmacion) Descriptor() ([]byte, []int) {
	return file_bibliotecann_NameNodeServer_proto_rawDescGZIP(), []int{2}
}

func (x *Confirmacion) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

type Consultalista struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Listadelibros string `protobuf:"bytes,1,opt,name=listadelibros,proto3" json:"listadelibros,omitempty"`
}

func (x *Consultalista) Reset() {
	*x = Consultalista{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bibliotecann_NameNodeServer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Consultalista) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Consultalista) ProtoMessage() {}

func (x *Consultalista) ProtoReflect() protoreflect.Message {
	mi := &file_bibliotecann_NameNodeServer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Consultalista.ProtoReflect.Descriptor instead.
func (*Consultalista) Descriptor() ([]byte, []int) {
	return file_bibliotecann_NameNodeServer_proto_rawDescGZIP(), []int{3}
}

func (x *Consultalista) GetListadelibros() string {
	if x != nil {
		return x.Listadelibros
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bibliotecann_NameNodeServer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_bibliotecann_NameNodeServer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_bibliotecann_NameNodeServer_proto_rawDescGZIP(), []int{4}
}

var File_bibliotecann_NameNodeServer_proto protoreflect.FileDescriptor

var file_bibliotecann_NameNodeServer_proto_rawDesc = []byte{
	0x0a, 0x21, 0x62, 0x69, 0x62, 0x6c, 0x69, 0x6f, 0x74, 0x65, 0x63, 0x61, 0x6e, 0x6e, 0x2f, 0x4e,
	0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74,
	0x61, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x6d, 0x61, 0x71, 0x75, 0x69,
	0x6e, 0x61, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x73, 0x6d, 0x61, 0x71, 0x75, 0x69, 0x6e, 0x61, 0x31, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x73, 0x6d, 0x61, 0x71, 0x75, 0x69, 0x6e, 0x61, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0e, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x6d, 0x61, 0x71, 0x75, 0x69, 0x6e, 0x61,
	0x32, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x6d, 0x61, 0x71, 0x75, 0x69,
	0x6e, 0x61, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x73, 0x6d, 0x61, 0x71, 0x75, 0x69, 0x6e, 0x61, 0x33, 0x22, 0x40, 0x0a, 0x08, 0x4c, 0x6f, 0x67,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x69, 0x70, 0x6d, 0x61, 0x71, 0x75, 0x69, 0x6e, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x70, 0x6d, 0x61, 0x71, 0x75, 0x69, 0x6e, 0x61, 0x22, 0x28, 0x0a, 0x0c, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x6e, 0x73, 0x61, 0x6a, 0x65, 0x22, 0x35, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74,
	0x61, 0x6c, 0x69, 0x73, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x69, 0x73, 0x74, 0x61, 0x64,
	0x65, 0x6c, 0x69, 0x62, 0x72, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c,
	0x69, 0x73, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x69, 0x62, 0x72, 0x6f, 0x73, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x95, 0x01, 0x0a, 0x0f, 0x4e, 0x61, 0x6d, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x0d, 0x53, 0x65, 0x6e,
	0x64, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x12, 0x0a, 0x2e, 0x50, 0x72, 0x6f,
	0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x1a, 0x0a, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73,
	0x74, 0x61, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0d, 0x45, 0x73, 0x63, 0x72, 0x69, 0x62, 0x69, 0x72,
	0x65, 0x6e, 0x4c, 0x6f, 0x67, 0x12, 0x09, 0x2e, 0x4c, 0x6f, 0x67, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x1a, 0x0d, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x22,
	0x00, 0x28, 0x01, 0x12, 0x28, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x6c, 0x69, 0x62, 0x72, 0x6f, 0x73,
	0x68, 0x61, 0x79, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x6c, 0x69, 0x73, 0x74, 0x61, 0x22, 0x00, 0x42, 0x2b, 0x5a,
	0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6e, 0x6f, 0x72,
	0x61, 0x6d, 0x62, 0x75, 0x2f, 0x74, 0x61, 0x72, 0x65, 0x61, 0x32, 0x73, 0x64, 0x2f, 0x62, 0x69,
	0x62, 0x6c, 0x69, 0x6f, 0x74, 0x65, 0x63, 0x61, 0x6e, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_bibliotecann_NameNodeServer_proto_rawDescOnce sync.Once
	file_bibliotecann_NameNodeServer_proto_rawDescData = file_bibliotecann_NameNodeServer_proto_rawDesc
)

func file_bibliotecann_NameNodeServer_proto_rawDescGZIP() []byte {
	file_bibliotecann_NameNodeServer_proto_rawDescOnce.Do(func() {
		file_bibliotecann_NameNodeServer_proto_rawDescData = protoimpl.X.CompressGZIP(file_bibliotecann_NameNodeServer_proto_rawDescData)
	})
	return file_bibliotecann_NameNodeServer_proto_rawDescData
}

var file_bibliotecann_NameNodeServer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bibliotecann_NameNodeServer_proto_goTypes = []interface{}{
	(*Propuesta)(nil),     // 0: Propuesta
	(*Logchunk)(nil),      // 1: Logchunk
	(*Confirmacion)(nil),  // 2: Confirmacion
	(*Consultalista)(nil), // 3: Consultalista
	(*Empty)(nil),         // 4: Empty
}
var file_bibliotecann_NameNodeServer_proto_depIdxs = []int32{
	0, // 0: NameNodeService.SendPropuesta:input_type -> Propuesta
	1, // 1: NameNodeService.EscribirenLog:input_type -> Logchunk
	4, // 2: NameNodeService.Quelibroshay:input_type -> Empty
	0, // 3: NameNodeService.SendPropuesta:output_type -> Propuesta
	2, // 4: NameNodeService.EscribirenLog:output_type -> Confirmacion
	3, // 5: NameNodeService.Quelibroshay:output_type -> Consultalista
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bibliotecann_NameNodeServer_proto_init() }
func file_bibliotecann_NameNodeServer_proto_init() {
	if File_bibliotecann_NameNodeServer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bibliotecann_NameNodeServer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Propuesta); i {
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
		file_bibliotecann_NameNodeServer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logchunk); i {
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
		file_bibliotecann_NameNodeServer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Confirmacion); i {
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
		file_bibliotecann_NameNodeServer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Consultalista); i {
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
		file_bibliotecann_NameNodeServer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_bibliotecann_NameNodeServer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bibliotecann_NameNodeServer_proto_goTypes,
		DependencyIndexes: file_bibliotecann_NameNodeServer_proto_depIdxs,
		MessageInfos:      file_bibliotecann_NameNodeServer_proto_msgTypes,
	}.Build()
	File_bibliotecann_NameNodeServer_proto = out.File
	file_bibliotecann_NameNodeServer_proto_rawDesc = nil
	file_bibliotecann_NameNodeServer_proto_goTypes = nil
	file_bibliotecann_NameNodeServer_proto_depIdxs = nil
}
