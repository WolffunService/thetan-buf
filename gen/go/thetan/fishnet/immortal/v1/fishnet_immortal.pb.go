// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: thetan/fishnet/immortal/v1/fishnet_immortal.proto

package thetan_fishnet_immortal_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	v1 "thetan-buf/gen/go/thetan/immortal/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RoomAllocationResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool `protobuf:"varint,1,opt,name=isSuccess,proto3" json:"isSuccess,omitempty"`
}

func (x *RoomAllocationResp) Reset() {
	*x = RoomAllocationResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomAllocationResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomAllocationResp) ProtoMessage() {}

func (x *RoomAllocationResp) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomAllocationResp.ProtoReflect.Descriptor instead.
func (*RoomAllocationResp) Descriptor() ([]byte, []int) {
	return file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_rawDescGZIP(), []int{0}
}

func (x *RoomAllocationResp) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

var File_thetan_fishnet_immortal_v1_fishnet_immortal_proto protoreflect.FileDescriptor

var file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_rawDesc = []byte{
	0x0a, 0x31, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x66, 0x69, 0x73, 0x68, 0x6e, 0x65, 0x74,
	0x2f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x73,
	0x68, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x66, 0x69, 0x73, 0x68,
	0x6e, 0x65, 0x74, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x1a,
	0x27, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x12, 0x52, 0x6f, 0x6f, 0x6d,
	0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c,
	0x0a, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x8e, 0x01, 0x0a,
	0x15, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x46, 0x69, 0x73, 0x68, 0x4e, 0x65, 0x74, 0x49, 0x6d,
	0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x12, 0x75, 0x0a, 0x0e, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x6c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61,
	0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d,
	0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x46, 0x6f, 0x75, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x2e,
	0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x66, 0x69, 0x73, 0x68, 0x6e, 0x65, 0x74, 0x2e, 0x69,
	0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x41,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x8a, 0x02,
	0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x66, 0x69, 0x73,
	0x68, 0x6e, 0x65, 0x74, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31,
	0x42, 0x14, 0x46, 0x69, 0x73, 0x68, 0x6e, 0x65, 0x74, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x2d, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x2f, 0x66, 0x69, 0x73, 0x68, 0x6e, 0x65, 0x74, 0x2f, 0x69, 0x6d, 0x6d, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5f, 0x66, 0x69,
	0x73, 0x68, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5f, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x54, 0x46, 0x49, 0xaa, 0x02, 0x1a, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x2e, 0x46, 0x69, 0x73, 0x68, 0x6e, 0x65, 0x74, 0x2e, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x46, 0x69,
	0x73, 0x68, 0x6e, 0x65, 0x74, 0x5c, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x26, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x46, 0x69, 0x73, 0x68, 0x6e,
	0x65, 0x74, 0x5c, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1d, 0x54, 0x68, 0x65,
	0x74, 0x61, 0x6e, 0x3a, 0x3a, 0x46, 0x69, 0x73, 0x68, 0x6e, 0x65, 0x74, 0x3a, 0x3a, 0x49, 0x6d,
	0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_rawDescOnce sync.Once
	file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_rawDescData = file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_rawDesc
)

func file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_rawDescGZIP() []byte {
	file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_rawDescOnce.Do(func() {
		file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_rawDescData = protoimpl.X.CompressGZIP(file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_rawDescData)
	})
	return file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_rawDescData
}

var file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_goTypes = []interface{}{
	(*RoomAllocationResp)(nil),                 // 0: thetan.fishnet.immortal.v1.RoomAllocationResp
	(*v1.ImmortalMatchFoundResponseProto)(nil), // 1: thetan.immortal.v1.ImmortalMatchFoundResponseProto
}
var file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_depIdxs = []int32{
	1, // 0: thetan.fishnet.immortal.v1.ThetanFishNetImmortal.RoomAllocation:input_type -> thetan.immortal.v1.ImmortalMatchFoundResponseProto
	0, // 1: thetan.fishnet.immortal.v1.ThetanFishNetImmortal.RoomAllocation:output_type -> thetan.fishnet.immortal.v1.RoomAllocationResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_init() }
func file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_init() {
	if File_thetan_fishnet_immortal_v1_fishnet_immortal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomAllocationResp); i {
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
			RawDescriptor: file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_goTypes,
		DependencyIndexes: file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_depIdxs,
		MessageInfos:      file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_msgTypes,
	}.Build()
	File_thetan_fishnet_immortal_v1_fishnet_immortal_proto = out.File
	file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_rawDesc = nil
	file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_goTypes = nil
	file_thetan_fishnet_immortal_v1_fishnet_immortal_proto_depIdxs = nil
}
