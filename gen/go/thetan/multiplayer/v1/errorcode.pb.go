// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: thetan/multiplayer/v1/errorcode.proto

package thetan_multiplayer_v1

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

type ChatErrorCode int32

const (
	ChatErrorCode_None      ChatErrorCode = 0
	ChatErrorCode_TimeLimit ChatErrorCode = 1
)

// Enum value maps for ChatErrorCode.
var (
	ChatErrorCode_name = map[int32]string{
		0: "None",
		1: "TimeLimit",
	}
	ChatErrorCode_value = map[string]int32{
		"None":      0,
		"TimeLimit": 1,
	}
)

func (x ChatErrorCode) Enum() *ChatErrorCode {
	p := new(ChatErrorCode)
	*p = x
	return p
}

func (x ChatErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChatErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_thetan_multiplayer_v1_errorcode_proto_enumTypes[0].Descriptor()
}

func (ChatErrorCode) Type() protoreflect.EnumType {
	return &file_thetan_multiplayer_v1_errorcode_proto_enumTypes[0]
}

func (x ChatErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChatErrorCode.Descriptor instead.
func (ChatErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_thetan_multiplayer_v1_errorcode_proto_rawDescGZIP(), []int{0}
}

type PartyErrorCode int32

const (
	PartyErrorCode_Unknown                PartyErrorCode = 0
	PartyErrorCode_InternalError          PartyErrorCode = -99
	PartyErrorCode_PartyIsFull            PartyErrorCode = 1
	PartyErrorCode_Dnd                    PartyErrorCode = 2
	PartyErrorCode_BlockInvite            PartyErrorCode = 3
	PartyErrorCode_PartyAlreadyStarted    PartyErrorCode = 4 // Join party when party is finding match
	PartyErrorCode_DifferentVersionClient PartyErrorCode = 5 //findmatch version
	PartyErrorCode_PartyIDInvalid         PartyErrorCode = 6
)

// Enum value maps for PartyErrorCode.
var (
	PartyErrorCode_name = map[int32]string{
		0:   "Unknown",
		-99: "InternalError",
		1:   "PartyIsFull",
		2:   "Dnd",
		3:   "BlockInvite",
		4:   "PartyAlreadyStarted",
		5:   "DifferentVersionClient",
		6:   "PartyIDInvalid",
	}
	PartyErrorCode_value = map[string]int32{
		"Unknown":                0,
		"InternalError":          -99,
		"PartyIsFull":            1,
		"Dnd":                    2,
		"BlockInvite":            3,
		"PartyAlreadyStarted":    4,
		"DifferentVersionClient": 5,
		"PartyIDInvalid":         6,
	}
)

func (x PartyErrorCode) Enum() *PartyErrorCode {
	p := new(PartyErrorCode)
	*p = x
	return p
}

func (x PartyErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PartyErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_thetan_multiplayer_v1_errorcode_proto_enumTypes[1].Descriptor()
}

func (PartyErrorCode) Type() protoreflect.EnumType {
	return &file_thetan_multiplayer_v1_errorcode_proto_enumTypes[1]
}

func (x PartyErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PartyErrorCode.Descriptor instead.
func (PartyErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_thetan_multiplayer_v1_errorcode_proto_rawDescGZIP(), []int{1}
}

var File_thetan_multiplayer_v1_errorcode_proto protoreflect.FileDescriptor

var file_thetan_multiplayer_v1_errorcode_proto_rawDesc = []byte{
	0x0a, 0x25, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2a, 0x28,
	0x0a, 0x0d, 0x43, 0x68, 0x61, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x69, 0x6d,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x10, 0x01, 0x2a, 0xad, 0x01, 0x0a, 0x0e, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x9d, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x74, 0x79, 0x49, 0x73, 0x46,
	0x75, 0x6c, 0x6c, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x6e, 0x64, 0x10, 0x02, 0x12, 0x0f,
	0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x10, 0x03, 0x12,
	0x17, 0x0a, 0x13, 0x50, 0x61, 0x72, 0x74, 0x79, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x69, 0x66, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x49, 0x44, 0x49,
	0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x06, 0x42, 0xe0, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d,
	0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x2d, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x3b, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x4d, 0x58, 0xaa, 0x02, 0x15,
	0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x21,
	0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x17, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x3a, 0x3a, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_thetan_multiplayer_v1_errorcode_proto_rawDescOnce sync.Once
	file_thetan_multiplayer_v1_errorcode_proto_rawDescData = file_thetan_multiplayer_v1_errorcode_proto_rawDesc
)

func file_thetan_multiplayer_v1_errorcode_proto_rawDescGZIP() []byte {
	file_thetan_multiplayer_v1_errorcode_proto_rawDescOnce.Do(func() {
		file_thetan_multiplayer_v1_errorcode_proto_rawDescData = protoimpl.X.CompressGZIP(file_thetan_multiplayer_v1_errorcode_proto_rawDescData)
	})
	return file_thetan_multiplayer_v1_errorcode_proto_rawDescData
}

var file_thetan_multiplayer_v1_errorcode_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_thetan_multiplayer_v1_errorcode_proto_goTypes = []interface{}{
	(ChatErrorCode)(0),  // 0: thetan.multiplayer.v1.ChatErrorCode
	(PartyErrorCode)(0), // 1: thetan.multiplayer.v1.PartyErrorCode
}
var file_thetan_multiplayer_v1_errorcode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_thetan_multiplayer_v1_errorcode_proto_init() }
func file_thetan_multiplayer_v1_errorcode_proto_init() {
	if File_thetan_multiplayer_v1_errorcode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_thetan_multiplayer_v1_errorcode_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_thetan_multiplayer_v1_errorcode_proto_goTypes,
		DependencyIndexes: file_thetan_multiplayer_v1_errorcode_proto_depIdxs,
		EnumInfos:         file_thetan_multiplayer_v1_errorcode_proto_enumTypes,
	}.Build()
	File_thetan_multiplayer_v1_errorcode_proto = out.File
	file_thetan_multiplayer_v1_errorcode_proto_rawDesc = nil
	file_thetan_multiplayer_v1_errorcode_proto_goTypes = nil
	file_thetan_multiplayer_v1_errorcode_proto_depIdxs = nil
}
