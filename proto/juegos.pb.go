// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: juegos.proto

package proto

import (
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

type JuegoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NombreJuego string `protobuf:"bytes,1,opt,name=nombreJuego,proto3" json:"nombreJuego,omitempty"`
	NoJugadores int64  `protobuf:"varint,2,opt,name=noJugadores,proto3" json:"noJugadores,omitempty"`
}

func (x *JuegoReq) Reset() {
	*x = JuegoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_juegos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JuegoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JuegoReq) ProtoMessage() {}

func (x *JuegoReq) ProtoReflect() protoreflect.Message {
	mi := &file_juegos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JuegoReq.ProtoReflect.Descriptor instead.
func (*JuegoReq) Descriptor() ([]byte, []int) {
	return file_juegos_proto_rawDescGZIP(), []int{0}
}

func (x *JuegoReq) GetNombreJuego() string {
	if x != nil {
		return x.NombreJuego
	}
	return ""
}

func (x *JuegoReq) GetNoJugadores() int64 {
	if x != nil {
		return x.NoJugadores
	}
	return 0
}

type JuegoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Estado  int64  `protobuf:"varint,2,opt,name=estado,proto3" json:"estado,omitempty"`
}

func (x *JuegoResp) Reset() {
	*x = JuegoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_juegos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JuegoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JuegoResp) ProtoMessage() {}

func (x *JuegoResp) ProtoReflect() protoreflect.Message {
	mi := &file_juegos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JuegoResp.ProtoReflect.Descriptor instead.
func (*JuegoResp) Descriptor() ([]byte, []int) {
	return file_juegos_proto_rawDescGZIP(), []int{1}
}

func (x *JuegoResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *JuegoResp) GetEstado() int64 {
	if x != nil {
		return x.Estado
	}
	return 0
}

var File_juegos_proto protoreflect.FileDescriptor

var file_juegos_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6a, 0x75, 0x65, 0x67, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x67, 0x72, 0x70, 0x63, 0x22, 0x4e, 0x0a, 0x08, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x52, 0x65, 0x71,
	0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x4a, 0x75, 0x65,
	0x67, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x6f, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6e, 0x6f, 0x4a, 0x75, 0x67, 0x61, 0x64,
	0x6f, 0x72, 0x65, 0x73, 0x22, 0x3d, 0x0a, 0x09, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x73, 0x74,
	0x61, 0x64, 0x6f, 0x32, 0x93, 0x01, 0x0a, 0x10, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x4a, 0x75, 0x65, 0x67,
	0x6f, 0x31, 0x12, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x52,
	0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x29, 0x0a, 0x06, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x32, 0x12, 0x0e, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x29,
	0x0a, 0x06, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x33, 0x12, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x4a, 0x75, 0x65, 0x67, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x4a, 0x75, 0x65, 0x67, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x67, 0x63, 0x65, 0x6c, 0x6f, 0x2f,
	0x53, 0x49, 0x4f, 0x50, 0x31, 0x5f, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_juegos_proto_rawDescOnce sync.Once
	file_juegos_proto_rawDescData = file_juegos_proto_rawDesc
)

func file_juegos_proto_rawDescGZIP() []byte {
	file_juegos_proto_rawDescOnce.Do(func() {
		file_juegos_proto_rawDescData = protoimpl.X.CompressGZIP(file_juegos_proto_rawDescData)
	})
	return file_juegos_proto_rawDescData
}

var file_juegos_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_juegos_proto_goTypes = []interface{}{
	(*JuegoReq)(nil),  // 0: grpc.JuegoReq
	(*JuegoResp)(nil), // 1: grpc.JuegoResp
}
var file_juegos_proto_depIdxs = []int32{
	0, // 0: grpc.JuegoListService.Juego1:input_type -> grpc.JuegoReq
	0, // 1: grpc.JuegoListService.Juego2:input_type -> grpc.JuegoReq
	0, // 2: grpc.JuegoListService.Juego3:input_type -> grpc.JuegoReq
	1, // 3: grpc.JuegoListService.Juego1:output_type -> grpc.JuegoResp
	1, // 4: grpc.JuegoListService.Juego2:output_type -> grpc.JuegoResp
	1, // 5: grpc.JuegoListService.Juego3:output_type -> grpc.JuegoResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_juegos_proto_init() }
func file_juegos_proto_init() {
	if File_juegos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_juegos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JuegoReq); i {
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
		file_juegos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JuegoResp); i {
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
			RawDescriptor: file_juegos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_juegos_proto_goTypes,
		DependencyIndexes: file_juegos_proto_depIdxs,
		MessageInfos:      file_juegos_proto_msgTypes,
	}.Build()
	File_juegos_proto = out.File
	file_juegos_proto_rawDesc = nil
	file_juegos_proto_goTypes = nil
	file_juegos_proto_depIdxs = nil
}
