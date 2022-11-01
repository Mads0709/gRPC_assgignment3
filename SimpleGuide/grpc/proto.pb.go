// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.7
// source: grpc/proto.proto

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Port int64 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Request) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

type ResponsMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Respond     string `protobuf:"bytes,1,opt,name=respond,proto3" json:"respond,omitempty"`
	Id          int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Vectorclock int64  `protobuf:"varint,3,opt,name=vectorclock,proto3" json:"vectorclock,omitempty"`
}

func (x *ResponsMessage) Reset() {
	*x = ResponsMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponsMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponsMessage) ProtoMessage() {}

func (x *ResponsMessage) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponsMessage.ProtoReflect.Descriptor instead.
func (*ResponsMessage) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{1}
}

func (x *ResponsMessage) GetRespond() string {
	if x != nil {
		return x.Respond
	}
	return ""
}

func (x *ResponsMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ResponsMessage) GetVectorclock() int64 {
	if x != nil {
		return x.Vectorclock
	}
	return 0
}

type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message     string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Vectorclock int64  `protobuf:"varint,3,opt,name=vectorclock,proto3" json:"vectorclock,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{2}
}

func (x *ChatMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChatMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ChatMessage) GetVectorclock() int64 {
	if x != nil {
		return x.Vectorclock
	}
	return 0
}

type ErrorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorMessage) Reset() {
	*x = ErrorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorMessage) ProtoMessage() {}

func (x *ErrorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorMessage.ProtoReflect.Descriptor instead.
func (*ErrorMessage) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{3}
}

func (x *ErrorMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_grpc_proto_proto protoreflect.FileDescriptor

var file_grpc_proto_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x22,
	0x2d, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x5c,
	0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x59, 0x0a, 0x0b,
	0x63, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x63,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x28, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0xa5, 0x01, 0x0a, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x10, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x12, 0x4a, 0x0a,
	0x13, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x67, 0x75, 0x69,
	0x64, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x19,
	0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_proto_rawDescOnce sync.Once
	file_grpc_proto_proto_rawDescData = file_grpc_proto_proto_rawDesc
)

func file_grpc_proto_proto_rawDescGZIP() []byte {
	file_grpc_proto_proto_rawDescOnce.Do(func() {
		file_grpc_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_proto_rawDescData)
	})
	return file_grpc_proto_proto_rawDescData
}

var file_grpc_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_proto_proto_goTypes = []interface{}{
	(*Request)(nil),        // 0: simpleguide.request
	(*ResponsMessage)(nil), // 1: simpleguide.responsMessage
	(*ChatMessage)(nil),    // 2: simpleguide.chatMessage
	(*ErrorMessage)(nil),   // 3: simpleguide.errorMessage
}
var file_grpc_proto_proto_depIdxs = []int32{
	0, // 0: simpleguide.registerClient.registerToServer:input_type -> simpleguide.request
	2, // 1: simpleguide.registerClient.populateChatMessage:input_type -> simpleguide.chatMessage
	1, // 2: simpleguide.registerClient.registerToServer:output_type -> simpleguide.responsMessage
	3, // 3: simpleguide.registerClient.populateChatMessage:output_type -> simpleguide.errorMessage
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_proto_proto_init() }
func file_grpc_proto_proto_init() {
	if File_grpc_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_grpc_proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponsMessage); i {
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
		file_grpc_proto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessage); i {
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
		file_grpc_proto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorMessage); i {
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
			RawDescriptor: file_grpc_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_proto_goTypes,
		DependencyIndexes: file_grpc_proto_proto_depIdxs,
		MessageInfos:      file_grpc_proto_proto_msgTypes,
	}.Build()
	File_grpc_proto_proto = out.File
	file_grpc_proto_proto_rawDesc = nil
	file_grpc_proto_proto_goTypes = nil
	file_grpc_proto_proto_depIdxs = nil
}
