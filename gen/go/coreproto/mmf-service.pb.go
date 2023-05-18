// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0-devel
// 	protoc        v3.21.12
// source: mmf-service.proto

package coreproto

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

type MmfTicketData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountPlayer     int32   `protobuf:"varint,1,opt,name=countPlayer,proto3" json:"countPlayer,omitempty"`
	BehaviorPoint   int32   `protobuf:"varint,2,opt,name=behaviorPoint,proto3" json:"behaviorPoint,omitempty"`
	IsHeroNFT       bool    `protobuf:"varint,3,opt,name=isHeroNFT,proto3" json:"isHeroNFT,omitempty"`
	TrophiesRank    int32   `protobuf:"varint,4,opt,name=trophiesRank,proto3" json:"trophiesRank,omitempty"`
	Regions         []int32 `protobuf:"varint,5,rep,packed,name=regions,proto3" json:"regions,omitempty"`
	MatchSearchType int32   `protobuf:"varint,6,opt,name=matchSearchType,proto3" json:"matchSearchType,omitempty"`
}

func (x *MmfTicketData) Reset() {
	*x = MmfTicketData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmf_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MmfTicketData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MmfTicketData) ProtoMessage() {}

func (x *MmfTicketData) ProtoReflect() protoreflect.Message {
	mi := &file_mmf_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MmfTicketData.ProtoReflect.Descriptor instead.
func (*MmfTicketData) Descriptor() ([]byte, []int) {
	return file_mmf_service_proto_rawDescGZIP(), []int{0}
}

func (x *MmfTicketData) GetCountPlayer() int32 {
	if x != nil {
		return x.CountPlayer
	}
	return 0
}

func (x *MmfTicketData) GetBehaviorPoint() int32 {
	if x != nil {
		return x.BehaviorPoint
	}
	return 0
}

func (x *MmfTicketData) GetIsHeroNFT() bool {
	if x != nil {
		return x.IsHeroNFT
	}
	return false
}

func (x *MmfTicketData) GetTrophiesRank() int32 {
	if x != nil {
		return x.TrophiesRank
	}
	return 0
}

func (x *MmfTicketData) GetRegions() []int32 {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *MmfTicketData) GetMatchSearchType() int32 {
	if x != nil {
		return x.MatchSearchType
	}
	return 0
}

var File_mmf_service_proto protoreflect.FileDescriptor

var file_mmf_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x6d, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x6d, 0x66, 0x22, 0xdd, 0x01, 0x0a, 0x0d, 0x4d, 0x6d, 0x66,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x48, 0x65, 0x72, 0x6f, 0x4e, 0x46, 0x54, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x48, 0x65, 0x72, 0x6f, 0x4e, 0x46, 0x54,
	0x12, 0x22, 0x0a, 0x0c, 0x74, 0x72, 0x6f, 0x70, 0x68, 0x69, 0x65, 0x73, 0x52, 0x61, 0x6e, 0x6b,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x72, 0x6f, 0x70, 0x68, 0x69, 0x65, 0x73,
	0x52, 0x61, 0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28,
	0x0a, 0x0f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x42, 0x2d, 0x0a, 0x0a, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x11, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x01, 0x5a, 0x0a, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmf_service_proto_rawDescOnce sync.Once
	file_mmf_service_proto_rawDescData = file_mmf_service_proto_rawDesc
)

func file_mmf_service_proto_rawDescGZIP() []byte {
	file_mmf_service_proto_rawDescOnce.Do(func() {
		file_mmf_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmf_service_proto_rawDescData)
	})
	return file_mmf_service_proto_rawDescData
}

var file_mmf_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mmf_service_proto_goTypes = []interface{}{
	(*MmfTicketData)(nil), // 0: mmf.MmfTicketData
}
var file_mmf_service_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mmf_service_proto_init() }
func file_mmf_service_proto_init() {
	if File_mmf_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmf_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MmfTicketData); i {
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
			RawDescriptor: file_mmf_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmf_service_proto_goTypes,
		DependencyIndexes: file_mmf_service_proto_depIdxs,
		MessageInfos:      file_mmf_service_proto_msgTypes,
	}.Build()
	File_mmf_service_proto = out.File
	file_mmf_service_proto_rawDesc = nil
	file_mmf_service_proto_goTypes = nil
	file_mmf_service_proto_depIdxs = nil
}
