// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: thetan/world/v1/service_tw.proto

package thetan_world_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InMarketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *InMarketRequest) Reset() {
	*x = InMarketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_world_v1_service_tw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InMarketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InMarketRequest) ProtoMessage() {}

func (x *InMarketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_world_v1_service_tw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InMarketRequest.ProtoReflect.Descriptor instead.
func (*InMarketRequest) Descriptor() ([]byte, []int) {
	return file_thetan_world_v1_service_tw_proto_rawDescGZIP(), []int{0}
}

func (x *InMarketRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type InMarketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InMarket bool `protobuf:"varint,1,opt,name=inMarket,proto3" json:"inMarket,omitempty"`
}

func (x *InMarketResponse) Reset() {
	*x = InMarketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_world_v1_service_tw_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InMarketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InMarketResponse) ProtoMessage() {}

func (x *InMarketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_world_v1_service_tw_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InMarketResponse.ProtoReflect.Descriptor instead.
func (*InMarketResponse) Descriptor() ([]byte, []int) {
	return file_thetan_world_v1_service_tw_proto_rawDescGZIP(), []int{1}
}

func (x *InMarketResponse) GetInMarket() bool {
	if x != nil {
		return x.InMarket
	}
	return false
}

var File_thetan_world_v1_service_tw_proto protoreflect.FileDescriptor

var file_thetan_world_v1_service_tw_proto_rawDesc = []byte{
	0x0a, 0x20, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x77, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x21, 0x0a, 0x0f, 0x49, 0x6e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x10, 0x49, 0x6e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x6e, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x32, 0x67, 0x0a, 0x12, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x57, 0x6f, 0x72,
	0x6c, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x49, 0x73, 0x49,
	0x6e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x20, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xb6, 0x01, 0x0a,
	0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x77, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2d, 0x62,
	0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x2f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x57, 0x58, 0xaa,
	0x02, 0x0f, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x0f, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x57, 0x6f, 0x72, 0x6c, 0x64,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x57, 0x6f, 0x72,
	0x6c, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x11, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x3a, 0x3a, 0x57, 0x6f, 0x72, 0x6c,
	0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thetan_world_v1_service_tw_proto_rawDescOnce sync.Once
	file_thetan_world_v1_service_tw_proto_rawDescData = file_thetan_world_v1_service_tw_proto_rawDesc
)

func file_thetan_world_v1_service_tw_proto_rawDescGZIP() []byte {
	file_thetan_world_v1_service_tw_proto_rawDescOnce.Do(func() {
		file_thetan_world_v1_service_tw_proto_rawDescData = protoimpl.X.CompressGZIP(file_thetan_world_v1_service_tw_proto_rawDescData)
	})
	return file_thetan_world_v1_service_tw_proto_rawDescData
}

var file_thetan_world_v1_service_tw_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_thetan_world_v1_service_tw_proto_goTypes = []interface{}{
	(*InMarketRequest)(nil),  // 0: thetan.world.v1.InMarketRequest
	(*InMarketResponse)(nil), // 1: thetan.world.v1.InMarketResponse
}
var file_thetan_world_v1_service_tw_proto_depIdxs = []int32{
	0, // 0: thetan.world.v1.ThetanWorldService.IsInMarket:input_type -> thetan.world.v1.InMarketRequest
	1, // 1: thetan.world.v1.ThetanWorldService.IsInMarket:output_type -> thetan.world.v1.InMarketResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_thetan_world_v1_service_tw_proto_init() }
func file_thetan_world_v1_service_tw_proto_init() {
	if File_thetan_world_v1_service_tw_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thetan_world_v1_service_tw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InMarketRequest); i {
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
		file_thetan_world_v1_service_tw_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InMarketResponse); i {
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
			RawDescriptor: file_thetan_world_v1_service_tw_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thetan_world_v1_service_tw_proto_goTypes,
		DependencyIndexes: file_thetan_world_v1_service_tw_proto_depIdxs,
		MessageInfos:      file_thetan_world_v1_service_tw_proto_msgTypes,
	}.Build()
	File_thetan_world_v1_service_tw_proto = out.File
	file_thetan_world_v1_service_tw_proto_rawDesc = nil
	file_thetan_world_v1_service_tw_proto_goTypes = nil
	file_thetan_world_v1_service_tw_proto_depIdxs = nil
}
