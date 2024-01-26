// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: thetan/rivals/v1/add_in.proto

package thetan_rivals_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

// Deprecated
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

type ChangeOwner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldOwner string     `protobuf:"bytes,1,opt,name=oldOwner,proto3" json:"oldOwner,omitempty"`
	NewOwner string     `protobuf:"bytes,2,opt,name=newOwner,proto3" json:"newOwner,omitempty"`
	ItemID   string     `protobuf:"bytes,3,opt,name=itemID,proto3" json:"itemID,omitempty"`
	Game     int32      `protobuf:"varint,4,opt,name=game,proto3" json:"game,omitempty"`
	Kind     int32      `protobuf:"varint,5,opt,name=kind,proto3" json:"kind,omitempty"`
	Type     int32      `protobuf:"varint,6,opt,name=type,proto3" json:"type,omitempty"`
	Reason   string     `protobuf:"bytes,7,opt,name=reason,proto3" json:"reason,omitempty"`
	Data     *anypb.Any `protobuf:"bytes,8,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ChangeOwner) Reset() {
	*x = ChangeOwner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_rivals_v1_add_in_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeOwner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeOwner) ProtoMessage() {}

func (x *ChangeOwner) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_rivals_v1_add_in_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeOwner.ProtoReflect.Descriptor instead.
func (*ChangeOwner) Descriptor() ([]byte, []int) {
	return file_thetan_rivals_v1_add_in_proto_rawDescGZIP(), []int{1}
}

func (x *ChangeOwner) GetOldOwner() string {
	if x != nil {
		return x.OldOwner
	}
	return ""
}

func (x *ChangeOwner) GetNewOwner() string {
	if x != nil {
		return x.NewOwner
	}
	return ""
}

func (x *ChangeOwner) GetItemID() string {
	if x != nil {
		return x.ItemID
	}
	return ""
}

func (x *ChangeOwner) GetGame() int32 {
	if x != nil {
		return x.Game
	}
	return 0
}

func (x *ChangeOwner) GetKind() int32 {
	if x != nil {
		return x.Kind
	}
	return 0
}

func (x *ChangeOwner) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ChangeOwner) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *ChangeOwner) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_thetan_rivals_v1_add_in_proto protoreflect.FileDescriptor

var file_thetan_rivals_v1_add_in_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x5f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x74, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x67, 0x0a, 0x0b, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c,
	0x0a, 0x09, 0x6f, 0x6c, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6f, 0x6c, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x65, 0x77, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x65, 0x77, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x65,
	0x77, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0xdb, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x6c, 0x64, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0xb9, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x74,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x42,
	0x0a, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x74,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x2d, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5f, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x5f,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x52, 0x58, 0xaa, 0x02, 0x10, 0x54, 0x68, 0x65, 0x74, 0x61,
	0x6e, 0x2e, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x54, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x5c, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x1c, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12,
	0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x3a, 0x3a, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_thetan_rivals_v1_add_in_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_thetan_rivals_v1_add_in_proto_goTypes = []interface{}{
	(*MigrateItem)(nil), // 0: thetan.rivals.v1.MigrateItem
	(*ChangeOwner)(nil), // 1: thetan.rivals.v1.ChangeOwner
	(*anypb.Any)(nil),   // 2: google.protobuf.Any
}
var file_thetan_rivals_v1_add_in_proto_depIdxs = []int32{
	2, // 0: thetan.rivals.v1.ChangeOwner.data:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
		file_thetan_rivals_v1_add_in_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeOwner); i {
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
			NumMessages:   2,
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
