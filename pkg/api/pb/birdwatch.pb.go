// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: birdwatch.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SendingInterval   int32 `protobuf:"varint,1,opt,name=SendingInterval,proto3" json:"SendingInterval,omitempty"`
	AveragingInterval int32 `protobuf:"varint,2,opt,name=AveragingInterval,proto3" json:"AveragingInterval,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_birdwatch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_birdwatch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_birdwatch_proto_rawDescGZIP(), []int{0}
}

func (x *Query) GetSendingInterval() int32 {
	if x != nil {
		return x.SendingInterval
	}
	return 0
}

func (x *Query) GetAveragingInterval() int32 {
	if x != nil {
		return x.AveragingInterval
	}
	return 0
}

var File_birdwatch_proto protoreflect.FileDescriptor

var file_birdwatch_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x69, 0x72, 0x64, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x62, 0x69, 0x72, 0x64, 0x77, 0x61, 0x74, 0x63, 0x68, 0x22, 0x5f, 0x0a, 0x05,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f,
	0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12,
	0x2c, 0x0a, 0x11, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x41, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x07, 0x5a,
	0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_birdwatch_proto_rawDescOnce sync.Once
	file_birdwatch_proto_rawDescData = file_birdwatch_proto_rawDesc
)

func file_birdwatch_proto_rawDescGZIP() []byte {
	file_birdwatch_proto_rawDescOnce.Do(func() {
		file_birdwatch_proto_rawDescData = protoimpl.X.CompressGZIP(file_birdwatch_proto_rawDescData)
	})
	return file_birdwatch_proto_rawDescData
}

var (
	file_birdwatch_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
	file_birdwatch_proto_goTypes  = []interface{}{
		(*Query)(nil), // 0: birdwatch.Query
	}
)

var file_birdwatch_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_birdwatch_proto_init() }
func file_birdwatch_proto_init() {
	if File_birdwatch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_birdwatch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
			RawDescriptor: file_birdwatch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_birdwatch_proto_goTypes,
		DependencyIndexes: file_birdwatch_proto_depIdxs,
		MessageInfos:      file_birdwatch_proto_msgTypes,
	}.Build()
	File_birdwatch_proto = out.File
	file_birdwatch_proto_rawDesc = nil
	file_birdwatch_proto_goTypes = nil
	file_birdwatch_proto_depIdxs = nil
}
