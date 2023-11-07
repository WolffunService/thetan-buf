// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: thetan/rivals/v1/add_in.proto

package thetan_rivals_v1

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

type MigrateItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldItemID string `protobuf:"bytes,1,opt,name=oldItemID,proto3" json:"oldItemID,omitempty"`
	NewItemID string `protobuf:"bytes,2,opt,name=newItemID,proto3" json:"newItemID,omitempty"`
	NewUserID string `protobuf:"bytes,3,opt,name=newUserID,proto3" json:"newUserID,omitempty"`
}

func (x *MigrateItem) Reset() {
	*x = MigrateItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_rivals_v1_add_in_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigrateItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigrateItem) ProtoMessage() {}

func (x *MigrateItem) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_rivals_v1_add_in_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigrateItem.ProtoReflect.Descriptor instead.
func (*MigrateItem) Descriptor() ([]byte, []int) {
	return file_thetan_rivals_v1_add_in_proto_rawDescGZIP(), []int{0}
}

func (x *MigrateItem) GetOldItemID() string {
	if x != nil {
		return x.OldItemID
	}
	return ""
}

func (x *MigrateItem) GetNewItemID() string {
	if x != nil {
		return x.NewItemID
	}
	return ""
}

func (x *MigrateItem) GetNewUserID() string {
	if x != nil {
		return x.NewUserID
	}
	return ""
}

var File_thetan_rivals_v1_add_in_proto protoreflect.FileDescriptor

var file_thetan_rivals_v1_add_in_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x5f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x21, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x0b, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x6c, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x6c, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x49,
	0x44, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x77, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x42, 0xb9, 0x01,
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76,
	0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2d, 0x62, 0x75, 0x66,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x72,
	0x69, 0x76, 0x61, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5f,
	0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x5f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x52, 0x58, 0xaa,
	0x02, 0x10, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x10, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x52, 0x69, 0x76, 0x61,
	0x6c, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x52,
	0x69, 0x76, 0x61, 0x6c, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x3a, 0x3a, 0x52,
	0x69, 0x76, 0x61, 0x6c, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_thetan_rivals_v1_add_in_proto_rawDescOnce sync.Once
	file_thetan_rivals_v1_add_in_proto_rawDescData = file_thetan_rivals_v1_add_in_proto_rawDesc
)

func file_thetan_rivals_v1_add_in_proto_rawDescGZIP() []byte {
	file_thetan_rivals_v1_add_in_proto_rawDescOnce.Do(func() {
		file_thetan_rivals_v1_add_in_proto_rawDescData = protoimpl.X.CompressGZIP(file_thetan_rivals_v1_add_in_proto_rawDescData)
	})
	return file_thetan_rivals_v1_add_in_proto_rawDescData
}

var file_thetan_rivals_v1_add_in_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_thetan_rivals_v1_add_in_proto_goTypes = []interface{}{
	(*MigrateItem)(nil), // 0: thetan.rivals.v1.MigrateItem
}
var file_thetan_rivals_v1_add_in_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_thetan_rivals_v1_add_in_proto_init() }
func file_thetan_rivals_v1_add_in_proto_init() {
	if File_thetan_rivals_v1_add_in_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thetan_rivals_v1_add_in_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigrateItem); i {
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
			RawDescriptor: file_thetan_rivals_v1_add_in_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_thetan_rivals_v1_add_in_proto_goTypes,
		DependencyIndexes: file_thetan_rivals_v1_add_in_proto_depIdxs,
		MessageInfos:      file_thetan_rivals_v1_add_in_proto_msgTypes,
	}.Build()
	File_thetan_rivals_v1_add_in_proto = out.File
	file_thetan_rivals_v1_add_in_proto_rawDesc = nil
	file_thetan_rivals_v1_add_in_proto_goTypes = nil
	file_thetan_rivals_v1_add_in_proto_depIdxs = nil
}
