// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: cpu.proto

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

type CpuCurrentStatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CpuCurrentStatRequest) Reset() {
	*x = CpuCurrentStatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cpu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CpuCurrentStatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpuCurrentStatRequest) ProtoMessage() {}

func (x *CpuCurrentStatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cpu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpuCurrentStatRequest.ProtoReflect.Descriptor instead.
func (*CpuCurrentStatRequest) Descriptor() ([]byte, []int) {
	return file_cpu_proto_rawDescGZIP(), []int{0}
}

type CpuCurrentStatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CpuCurrentStatResponse) Reset() {
	*x = CpuCurrentStatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cpu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CpuCurrentStatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpuCurrentStatResponse) ProtoMessage() {}

func (x *CpuCurrentStatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cpu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpuCurrentStatResponse.ProtoReflect.Descriptor instead.
func (*CpuCurrentStatResponse) Descriptor() ([]byte, []int) {
	return file_cpu_proto_rawDescGZIP(), []int{1}
}

var File_cpu_proto protoreflect.FileDescriptor

var file_cpu_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x70, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x70, 0x75,
	0x22, 0x17, 0x0a, 0x15, 0x43, 0x70, 0x75, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x43, 0x70, 0x75,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0x55, 0x0a, 0x03, 0x43, 0x70, 0x75, 0x12, 0x4e, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x43, 0x70, 0x75, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x12,
	0x1a, 0x2e, 0x63, 0x70, 0x75, 0x2e, 0x43, 0x70, 0x75, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x70,
	0x75, 0x2e, 0x43, 0x70, 0x75, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cpu_proto_rawDescOnce sync.Once
	file_cpu_proto_rawDescData = file_cpu_proto_rawDesc
)

func file_cpu_proto_rawDescGZIP() []byte {
	file_cpu_proto_rawDescOnce.Do(func() {
		file_cpu_proto_rawDescData = protoimpl.X.CompressGZIP(file_cpu_proto_rawDescData)
	})
	return file_cpu_proto_rawDescData
}

var (
	file_cpu_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
	file_cpu_proto_goTypes  = []interface{}{
		(*CpuCurrentStatRequest)(nil),  // 0: cpu.CpuCurrentStatRequest
		(*CpuCurrentStatResponse)(nil), // 1: cpu.CpuCurrentStatResponse
	}
)

var file_cpu_proto_depIdxs = []int32{
	0, // 0: cpu.Cpu.GetCpuCurrentStat:input_type -> cpu.CpuCurrentStatRequest
	1, // 1: cpu.Cpu.GetCpuCurrentStat:output_type -> cpu.CpuCurrentStatResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cpu_proto_init() }
func file_cpu_proto_init() {
	if File_cpu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cpu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CpuCurrentStatRequest); i {
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
		file_cpu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CpuCurrentStatResponse); i {
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
			RawDescriptor: file_cpu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cpu_proto_goTypes,
		DependencyIndexes: file_cpu_proto_depIdxs,
		MessageInfos:      file_cpu_proto_msgTypes,
	}.Build()
	File_cpu_proto = out.File
	file_cpu_proto_rawDesc = nil
	file_cpu_proto_goTypes = nil
	file_cpu_proto_depIdxs = nil
}
