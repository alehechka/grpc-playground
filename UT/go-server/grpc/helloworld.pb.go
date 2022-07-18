// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: go-server/grpc/helloworld.proto

package helloworld

import (
	_ "github.com/ysugimoto/grpc-graphql-gateway/graphql"
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

// The request message containing the user's name.
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age    int64  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Status bool   `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_server_grpc_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_server_grpc_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_go_server_grpc_helloworld_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloRequest) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *HelloRequest) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

// The response message containing the greetings
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_server_grpc_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_go_server_grpc_helloworld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_go_server_grpc_helloworld_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GoodbyeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GoodbyeRequest) Reset() {
	*x = GoodbyeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_server_grpc_helloworld_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodbyeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodbyeRequest) ProtoMessage() {}

func (x *GoodbyeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_server_grpc_helloworld_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodbyeRequest.ProtoReflect.Descriptor instead.
func (*GoodbyeRequest) Descriptor() ([]byte, []int) {
	return file_go_server_grpc_helloworld_proto_rawDescGZIP(), []int{2}
}

func (x *GoodbyeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GoodbyeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GoodbyeReply) Reset() {
	*x = GoodbyeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_server_grpc_helloworld_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodbyeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodbyeReply) ProtoMessage() {}

func (x *GoodbyeReply) ProtoReflect() protoreflect.Message {
	mi := &file_go_server_grpc_helloworld_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodbyeReply.ProtoReflect.Descriptor instead.
func (*GoodbyeReply) Descriptor() ([]byte, []int) {
	return file_go_server_grpc_helloworld_proto_rawDescGZIP(), []int{3}
}

func (x *GoodbyeReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_go_server_grpc_helloworld_proto protoreflect.FileDescriptor

var file_go_server_grpc_helloworld_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x67, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x1a, 0x1c, 0x67,
	0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x0c, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0xba, 0x43, 0x02, 0x08, 0x01,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x03, 0xba, 0x43, 0x00, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xba,
	0x43, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x2b, 0x0a, 0x0e, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x05, 0xba, 0x43, 0x02, 0x08, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x28, 0x0a, 0x0c, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xba, 0x01, 0x0a, 0x07, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x12, 0x18, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x0a, 0xba, 0x43, 0x07, 0x12, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x12,
	0x50, 0x0a, 0x0a, 0x53, 0x61, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x12, 0x1a, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x62,
	0x79, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x0c, 0xba, 0x43, 0x09, 0x12, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x79,
	0x65, 0x1a, 0x13, 0xba, 0x43, 0x10, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73,
	0x74, 0x3a, 0x38, 0x30, 0x10, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_server_grpc_helloworld_proto_rawDescOnce sync.Once
	file_go_server_grpc_helloworld_proto_rawDescData = file_go_server_grpc_helloworld_proto_rawDesc
)

func file_go_server_grpc_helloworld_proto_rawDescGZIP() []byte {
	file_go_server_grpc_helloworld_proto_rawDescOnce.Do(func() {
		file_go_server_grpc_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_server_grpc_helloworld_proto_rawDescData)
	})
	return file_go_server_grpc_helloworld_proto_rawDescData
}

var file_go_server_grpc_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_go_server_grpc_helloworld_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),   // 0: helloworld.HelloRequest
	(*HelloReply)(nil),     // 1: helloworld.HelloReply
	(*GoodbyeRequest)(nil), // 2: helloworld.GoodbyeRequest
	(*GoodbyeReply)(nil),   // 3: helloworld.GoodbyeReply
}
var file_go_server_grpc_helloworld_proto_depIdxs = []int32{
	0, // 0: helloworld.Greeter.SayHello:input_type -> helloworld.HelloRequest
	2, // 1: helloworld.Greeter.SayGoodbye:input_type -> helloworld.GoodbyeRequest
	1, // 2: helloworld.Greeter.SayHello:output_type -> helloworld.HelloReply
	3, // 3: helloworld.Greeter.SayGoodbye:output_type -> helloworld.GoodbyeReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_server_grpc_helloworld_proto_init() }
func file_go_server_grpc_helloworld_proto_init() {
	if File_go_server_grpc_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_server_grpc_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_go_server_grpc_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
		file_go_server_grpc_helloworld_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodbyeRequest); i {
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
		file_go_server_grpc_helloworld_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodbyeReply); i {
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
			RawDescriptor: file_go_server_grpc_helloworld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_server_grpc_helloworld_proto_goTypes,
		DependencyIndexes: file_go_server_grpc_helloworld_proto_depIdxs,
		MessageInfos:      file_go_server_grpc_helloworld_proto_msgTypes,
	}.Build()
	File_go_server_grpc_helloworld_proto = out.File
	file_go_server_grpc_helloworld_proto_rawDesc = nil
	file_go_server_grpc_helloworld_proto_goTypes = nil
	file_go_server_grpc_helloworld_proto_depIdxs = nil
}