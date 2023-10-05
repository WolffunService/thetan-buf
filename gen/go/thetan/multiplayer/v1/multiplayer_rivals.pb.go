// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: thetan/multiplayer/v1/multiplayer_rivals.proto

package thetan_multiplayer_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	v1 "thetan-buf/gen/go/thetan/shared/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RivalsPlayerInfoProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinionId       string           `protobuf:"bytes,1,opt,name=minionId,proto3" json:"minionId,omitempty"`
	Skin           int32            `protobuf:"varint,2,opt,name=skin,proto3" json:"skin,omitempty"`
	CosmeticInUsed map[string]int64 `protobuf:"bytes,3,rep,name=cosmeticInUsed,proto3" json:"cosmeticInUsed,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Customized     *v1.Customized   `protobuf:"bytes,4,opt,name=customized,proto3" json:"customized,omitempty"`
}

func (x *RivalsPlayerInfoProto) Reset() {
	*x = RivalsPlayerInfoProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_multiplayer_v1_multiplayer_rivals_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RivalsPlayerInfoProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RivalsPlayerInfoProto) ProtoMessage() {}

func (x *RivalsPlayerInfoProto) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_multiplayer_v1_multiplayer_rivals_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RivalsPlayerInfoProto.ProtoReflect.Descriptor instead.
func (*RivalsPlayerInfoProto) Descriptor() ([]byte, []int) {
	return file_thetan_multiplayer_v1_multiplayer_rivals_proto_rawDescGZIP(), []int{0}
}

func (x *RivalsPlayerInfoProto) GetMinionId() string {
	if x != nil {
		return x.MinionId
	}
	return ""
}

func (x *RivalsPlayerInfoProto) GetSkin() int32 {
	if x != nil {
		return x.Skin
	}
	return 0
}

func (x *RivalsPlayerInfoProto) GetCosmeticInUsed() map[string]int64 {
	if x != nil {
		return x.CosmeticInUsed
	}
	return nil
}

func (x *RivalsPlayerInfoProto) GetCustomized() *v1.Customized {
	if x != nil {
		return x.Customized
	}
	return nil
}

var File_thetan_multiplayer_v1_multiplayer_rivals_proto protoreflect.FileDescriptor

var file_thetan_multiplayer_v1_multiplayer_rivals_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x21, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x69, 0x7a, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x02, 0x0a, 0x15, 0x52,
	0x69, 0x76, 0x61, 0x6c, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x73, 0x6b, 0x69, 0x6e, 0x12, 0x68, 0x0a, 0x0e, 0x63, 0x6f, 0x73, 0x6d, 0x65, 0x74, 0x69, 0x63,
	0x49, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x74,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x73, 0x6d, 0x65,
	0x74, 0x69, 0x63, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e,
	0x63, 0x6f, 0x73, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x12, 0x3c,
	0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x1a, 0x41, 0x0a, 0x13,
	0x43, 0x6f, 0x73, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0xe8, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x16, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2d,
	0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x68, 0x65, 0x74, 0x61,
	0x6e, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x3b, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x4d, 0x58, 0xaa, 0x02, 0x15, 0x54,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x21, 0x54,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x17, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x3a, 0x3a, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_thetan_multiplayer_v1_multiplayer_rivals_proto_rawDescOnce sync.Once
	file_thetan_multiplayer_v1_multiplayer_rivals_proto_rawDescData = file_thetan_multiplayer_v1_multiplayer_rivals_proto_rawDesc
)

func file_thetan_multiplayer_v1_multiplayer_rivals_proto_rawDescGZIP() []byte {
	file_thetan_multiplayer_v1_multiplayer_rivals_proto_rawDescOnce.Do(func() {
		file_thetan_multiplayer_v1_multiplayer_rivals_proto_rawDescData = protoimpl.X.CompressGZIP(file_thetan_multiplayer_v1_multiplayer_rivals_proto_rawDescData)
	})
	return file_thetan_multiplayer_v1_multiplayer_rivals_proto_rawDescData
}

var file_thetan_multiplayer_v1_multiplayer_rivals_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_thetan_multiplayer_v1_multiplayer_rivals_proto_goTypes = []interface{}{
	(*RivalsPlayerInfoProto)(nil), // 0: thetan.multiplayer.v1.RivalsPlayerInfoProto
	nil,                           // 1: thetan.multiplayer.v1.RivalsPlayerInfoProto.CosmeticInUsedEntry
	(*v1.Customized)(nil),         // 2: thetan.shared.v1.Customized
}
var file_thetan_multiplayer_v1_multiplayer_rivals_proto_depIdxs = []int32{
	1, // 0: thetan.multiplayer.v1.RivalsPlayerInfoProto.cosmeticInUsed:type_name -> thetan.multiplayer.v1.RivalsPlayerInfoProto.CosmeticInUsedEntry
	2, // 1: thetan.multiplayer.v1.RivalsPlayerInfoProto.customized:type_name -> thetan.shared.v1.Customized
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_thetan_multiplayer_v1_multiplayer_rivals_proto_init() }
func file_thetan_multiplayer_v1_multiplayer_rivals_proto_init() {
	if File_thetan_multiplayer_v1_multiplayer_rivals_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thetan_multiplayer_v1_multiplayer_rivals_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RivalsPlayerInfoProto); i {
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
			RawDescriptor: file_thetan_multiplayer_v1_multiplayer_rivals_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_thetan_multiplayer_v1_multiplayer_rivals_proto_goTypes,
		DependencyIndexes: file_thetan_multiplayer_v1_multiplayer_rivals_proto_depIdxs,
		MessageInfos:      file_thetan_multiplayer_v1_multiplayer_rivals_proto_msgTypes,
	}.Build()
	File_thetan_multiplayer_v1_multiplayer_rivals_proto = out.File
	file_thetan_multiplayer_v1_multiplayer_rivals_proto_rawDesc = nil
	file_thetan_multiplayer_v1_multiplayer_rivals_proto_goTypes = nil
	file_thetan_multiplayer_v1_multiplayer_rivals_proto_depIdxs = nil
}
