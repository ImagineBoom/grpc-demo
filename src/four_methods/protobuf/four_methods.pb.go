// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: protobuf/four_methods.proto

package four_methods_pb

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

type DemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DemoRequest) Reset() {
	*x = DemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_four_methods_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoRequest) ProtoMessage() {}

func (x *DemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_four_methods_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoRequest.ProtoReflect.Descriptor instead.
func (*DemoRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_four_methods_proto_rawDescGZIP(), []int{0}
}

func (x *DemoRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type DemoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DemoResponse) Reset() {
	*x = DemoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_four_methods_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoResponse) ProtoMessage() {}

func (x *DemoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_four_methods_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoResponse.ProtoReflect.Descriptor instead.
func (*DemoResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_four_methods_proto_rawDescGZIP(), []int{1}
}

func (x *DemoResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_protobuf_four_methods_proto protoreflect.FileDescriptor

var file_protobuf_four_methods_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x6f, 0x75, 0x72, 0x5f,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x21, 0x0a, 0x0b, 0x44, 0x65, 0x6d, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x22, 0x0a, 0x0c, 0x44, 0x65,
	0x6d, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x93,
	0x02, 0x0a, 0x04, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x3b, 0x0a, 0x08, 0x55, 0x6e, 0x61, 0x72, 0x79,
	0x52, 0x50, 0x43, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x50, 0x43, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x44, 0x0a, 0x0f, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x50, 0x43, 0x12, 0x15, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x12, 0x42, 0x0a, 0x0b, 0x42, 0x69, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x50, 0x43, 0x12,
	0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x66, 0x6f, 0x75, 0x72, 0x5f, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protobuf_four_methods_proto_rawDescOnce sync.Once
	file_protobuf_four_methods_proto_rawDescData = file_protobuf_four_methods_proto_rawDesc
)

func file_protobuf_four_methods_proto_rawDescGZIP() []byte {
	file_protobuf_four_methods_proto_rawDescOnce.Do(func() {
		file_protobuf_four_methods_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_four_methods_proto_rawDescData)
	})
	return file_protobuf_four_methods_proto_rawDescData
}

var file_protobuf_four_methods_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protobuf_four_methods_proto_goTypes = []interface{}{
	(*DemoRequest)(nil),  // 0: protobuf.DemoRequest
	(*DemoResponse)(nil), // 1: protobuf.DemoResponse
}
var file_protobuf_four_methods_proto_depIdxs = []int32{
	0, // 0: protobuf.Demo.UnaryRPC:input_type -> protobuf.DemoRequest
	0, // 1: protobuf.Demo.ServerStreamRPC:input_type -> protobuf.DemoRequest
	0, // 2: protobuf.Demo.ClientStreamRPC:input_type -> protobuf.DemoRequest
	0, // 3: protobuf.Demo.BiStreamRPC:input_type -> protobuf.DemoRequest
	1, // 4: protobuf.Demo.UnaryRPC:output_type -> protobuf.DemoResponse
	1, // 5: protobuf.Demo.ServerStreamRPC:output_type -> protobuf.DemoResponse
	1, // 6: protobuf.Demo.ClientStreamRPC:output_type -> protobuf.DemoResponse
	1, // 7: protobuf.Demo.BiStreamRPC:output_type -> protobuf.DemoResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_four_methods_proto_init() }
func file_protobuf_four_methods_proto_init() {
	if File_protobuf_four_methods_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_four_methods_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoRequest); i {
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
		file_protobuf_four_methods_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoResponse); i {
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
			RawDescriptor: file_protobuf_four_methods_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_four_methods_proto_goTypes,
		DependencyIndexes: file_protobuf_four_methods_proto_depIdxs,
		MessageInfos:      file_protobuf_four_methods_proto_msgTypes,
	}.Build()
	File_protobuf_four_methods_proto = out.File
	file_protobuf_four_methods_proto_rawDesc = nil
	file_protobuf_four_methods_proto_goTypes = nil
	file_protobuf_four_methods_proto_depIdxs = nil
}