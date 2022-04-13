// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: pb/hello.proto

package pb

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

type StreamReqData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StreamReqData) Reset() {
	*x = StreamReqData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamReqData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamReqData) ProtoMessage() {}

func (x *StreamReqData) ProtoReflect() protoreflect.Message {
	mi := &file_pb_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamReqData.ProtoReflect.Descriptor instead.
func (*StreamReqData) Descriptor() ([]byte, []int) {
	return file_pb_hello_proto_rawDescGZIP(), []int{0}
}

func (x *StreamReqData) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type StreamResData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StreamResData) Reset() {
	*x = StreamResData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResData) ProtoMessage() {}

func (x *StreamResData) ProtoReflect() protoreflect.Message {
	mi := &file_pb_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResData.ProtoReflect.Descriptor instead.
func (*StreamResData) Descriptor() ([]byte, []int) {
	return file_pb_hello_proto_rawDescGZIP(), []int{1}
}

func (x *StreamResData) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_pb_hello_proto protoreflect.FileDescriptor

var file_pb_hello_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x62, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x23, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x23, 0x0a, 0x0d, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x40,
	0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x30, 0x01,
	0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pb_hello_proto_rawDescOnce sync.Once
	file_pb_hello_proto_rawDescData = file_pb_hello_proto_rawDesc
)

func file_pb_hello_proto_rawDescGZIP() []byte {
	file_pb_hello_proto_rawDescOnce.Do(func() {
		file_pb_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_hello_proto_rawDescData)
	})
	return file_pb_hello_proto_rawDescData
}

var file_pb_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_hello_proto_goTypes = []interface{}{
	(*StreamReqData)(nil), // 0: pb.StreamReqData
	(*StreamResData)(nil), // 1: pb.StreamResData
}
var file_pb_hello_proto_depIdxs = []int32{
	0, // 0: pb.Greeter.GetStream:input_type -> pb.StreamReqData
	1, // 1: pb.Greeter.GetStream:output_type -> pb.StreamResData
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_hello_proto_init() }
func file_pb_hello_proto_init() {
	if File_pb_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamReqData); i {
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
		file_pb_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamResData); i {
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
			RawDescriptor: file_pb_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_hello_proto_goTypes,
		DependencyIndexes: file_pb_hello_proto_depIdxs,
		MessageInfos:      file_pb_hello_proto_msgTypes,
	}.Build()
	File_pb_hello_proto = out.File
	file_pb_hello_proto_rawDesc = nil
	file_pb_hello_proto_goTypes = nil
	file_pb_hello_proto_depIdxs = nil
}
