// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: thetan/multiplayer/v1/multiplayer_immortal.proto

package thetan_multiplayer_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	_ "thetan-buf/gen/go/thetan/shared/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ImmortalPlayerInfoProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkinId   int32   `protobuf:"varint,1,opt,name=skinId,proto3" json:"skinId,omitempty"` //todo update neu cho nhieu hon 1 hero, hien thi dev chi nghi toi case choi 1 nen code 1
	SkillIds []int32 `protobuf:"varint,2,rep,packed,name=skillIds,proto3" json:"skillIds,omitempty"`
}

func (x *ImmortalPlayerInfoProto) Reset() {
	*x = ImmortalPlayerInfoProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_multiplayer_v1_multiplayer_immortal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImmortalPlayerInfoProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImmortalPlayerInfoProto) ProtoMessage() {}

func (x *ImmortalPlayerInfoProto) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_multiplayer_v1_multiplayer_immortal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImmortalPlayerInfoProto.ProtoReflect.Descriptor instead.
func (*ImmortalPlayerInfoProto) Descriptor() ([]byte, []int) {
	return file_thetan_multiplayer_v1_multiplayer_immortal_proto_rawDescGZIP(), []int{0}
}

func (x *ImmortalPlayerInfoProto) GetSkinId() int32 {
	if x != nil {
		return x.SkinId
	}
	return 0
}

func (x *ImmortalPlayerInfoProto) GetSkillIds() []int32 {
	if x != nil {
		return x.SkillIds
	}
	return nil
}

var File_thetan_multiplayer_v1_multiplayer_immortal_proto protoreflect.FileDescriptor

var file_thetan_multiplayer_v1_multiplayer_immortal_proto_rawDesc = []byte{
	0x0a, 0x30, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x21, 0x74, 0x68, 0x65, 0x74, 0x61,
	0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x17,
	0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6b, 0x69, 0x6e, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x6b, 0x69, 0x6e, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x73, 0x42, 0xea, 0x01, 0x0a, 0x19,
	0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x18, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2d, 0x62, 0x75,
	0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x74,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x5f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x4d, 0x58, 0xaa, 0x02, 0x15, 0x54, 0x68, 0x65,
	0x74, 0x61, 0x6e, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x15, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x21, 0x54, 0x68, 0x65,
	0x74, 0x61, 0x6e, 0x5c, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x17, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x3a, 0x3a, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thetan_multiplayer_v1_multiplayer_immortal_proto_rawDescOnce sync.Once
	file_thetan_multiplayer_v1_multiplayer_immortal_proto_rawDescData = file_thetan_multiplayer_v1_multiplayer_immortal_proto_rawDesc
)

func file_thetan_multiplayer_v1_multiplayer_immortal_proto_rawDescGZIP() []byte {
	file_thetan_multiplayer_v1_multiplayer_immortal_proto_rawDescOnce.Do(func() {
		file_thetan_multiplayer_v1_multiplayer_immortal_proto_rawDescData = protoimpl.X.CompressGZIP(file_thetan_multiplayer_v1_multiplayer_immortal_proto_rawDescData)
	})
	return file_thetan_multiplayer_v1_multiplayer_immortal_proto_rawDescData
}

var file_thetan_multiplayer_v1_multiplayer_immortal_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_thetan_multiplayer_v1_multiplayer_immortal_proto_goTypes = []interface{}{
	(*ImmortalPlayerInfoProto)(nil), // 0: thetan.multiplayer.v1.ImmortalPlayerInfoProto
}
var file_thetan_multiplayer_v1_multiplayer_immortal_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_thetan_multiplayer_v1_multiplayer_immortal_proto_init() }
func file_thetan_multiplayer_v1_multiplayer_immortal_proto_init() {
	if File_thetan_multiplayer_v1_multiplayer_immortal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thetan_multiplayer_v1_multiplayer_immortal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImmortalPlayerInfoProto); i {
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
			RawDescriptor: file_thetan_multiplayer_v1_multiplayer_immortal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_thetan_multiplayer_v1_multiplayer_immortal_proto_goTypes,
		DependencyIndexes: file_thetan_multiplayer_v1_multiplayer_immortal_proto_depIdxs,
		MessageInfos:      file_thetan_multiplayer_v1_multiplayer_immortal_proto_msgTypes,
	}.Build()
	File_thetan_multiplayer_v1_multiplayer_immortal_proto = out.File
	file_thetan_multiplayer_v1_multiplayer_immortal_proto_rawDesc = nil
	file_thetan_multiplayer_v1_multiplayer_immortal_proto_goTypes = nil
	file_thetan_multiplayer_v1_multiplayer_immortal_proto_depIdxs = nil
}
