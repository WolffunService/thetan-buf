// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: game_info.proto

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

type GameMode int32

const (
	GameMode_RANKED             GameMode = 0
	GameMode_SPECIAL_EVENT      GameMode = 1
	GameMode_CUSTOM_MODE        GameMode = 2
	GameMode_THETAN_RIVALS_MODE GameMode = 3
)

// Enum value maps for GameMode.
var (
	GameMode_name = map[int32]string{
		0: "RANKED",
		1: "SPECIAL_EVENT",
		2: "CUSTOM_MODE",
		3: "THETAN_RIVALS_MODE",
	}
	GameMode_value = map[string]int32{
		"RANKED":             0,
		"SPECIAL_EVENT":      1,
		"CUSTOM_MODE":        2,
		"THETAN_RIVALS_MODE": 3,
	}
)

func (x GameMode) Enum() *GameMode {
	p := new(GameMode)
	*p = x
	return p
}

func (x GameMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameMode) Descriptor() protoreflect.EnumDescriptor {
	return file_game_info_proto_enumTypes[0].Descriptor()
}

func (GameMode) Type() protoreflect.EnumType {
	return &file_game_info_proto_enumTypes[0]
}

func (x GameMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameMode.Descriptor instead.
func (GameMode) EnumDescriptor() ([]byte, []int) {
	return file_game_info_proto_rawDescGZIP(), []int{0}
}

type InGameMode int32

const (
	InGameMode_TEAM_COLLECT_STAR        InGameMode = 0
	InGameMode_SOLO_SURVIVAL            InGameMode = 1
	InGameMode_DUAL_SURVIVAL            InGameMode = 2
	InGameMode_TEAM_COLLECT_STAR_4_VS_4 InGameMode = 3
	InGameMode_KING                     InGameMode = 5
	InGameMode_DEATH_MATCH              InGameMode = 6
	InGameMode_DEATH_MATCH_3_VS_3       InGameMode = 7
	InGameMode_FLAG                     InGameMode = 8
	InGameMode_TOWER                    InGameMode = 9
	InGameMode_BATTLE_ROYALE            InGameMode = 12
	InGameMode_SQUAD_BATTLE_ROYALE      InGameMode = 13
	InGameMode_DUO_BATTLE_ROYALE        InGameMode = 14
	InGameMode_TRIO_BATTLE_ROAYLE       InGameMode = 15
	InGameMode_THETAN_RIVALS            InGameMode = 20
	InGameMode_NONE_MODE                InGameMode = -1
)

// Enum value maps for InGameMode.
var (
	InGameMode_name = map[int32]string{
		0:  "TEAM_COLLECT_STAR",
		1:  "SOLO_SURVIVAL",
		2:  "DUAL_SURVIVAL",
		3:  "TEAM_COLLECT_STAR_4_VS_4",
		5:  "KING",
		6:  "DEATH_MATCH",
		7:  "DEATH_MATCH_3_VS_3",
		8:  "FLAG",
		9:  "TOWER",
		12: "BATTLE_ROYALE",
		13: "SQUAD_BATTLE_ROYALE",
		14: "DUO_BATTLE_ROYALE",
		15: "TRIO_BATTLE_ROAYLE",
		20: "THETAN_RIVALS",
		-1: "NONE_MODE",
	}
	InGameMode_value = map[string]int32{
		"TEAM_COLLECT_STAR":        0,
		"SOLO_SURVIVAL":            1,
		"DUAL_SURVIVAL":            2,
		"TEAM_COLLECT_STAR_4_VS_4": 3,
		"KING":                     5,
		"DEATH_MATCH":              6,
		"DEATH_MATCH_3_VS_3":       7,
		"FLAG":                     8,
		"TOWER":                    9,
		"BATTLE_ROYALE":            12,
		"SQUAD_BATTLE_ROYALE":      13,
		"DUO_BATTLE_ROYALE":        14,
		"TRIO_BATTLE_ROAYLE":       15,
		"THETAN_RIVALS":            20,
		"NONE_MODE":                -1,
	}
)

func (x InGameMode) Enum() *InGameMode {
	p := new(InGameMode)
	*p = x
	return p
}

func (x InGameMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InGameMode) Descriptor() protoreflect.EnumDescriptor {
	return file_game_info_proto_enumTypes[1].Descriptor()
}

func (InGameMode) Type() protoreflect.EnumType {
	return &file_game_info_proto_enumTypes[1]
}

func (x InGameMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InGameMode.Descriptor instead.
func (InGameMode) EnumDescriptor() ([]byte, []int) {
	return file_game_info_proto_rawDescGZIP(), []int{1}
}

var File_game_info_proto protoreflect.FileDescriptor

var file_game_info_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x52, 0x0a,
	0x08, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x41, 0x4e,
	0x4b, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x50, 0x45, 0x43, 0x49, 0x41, 0x4c,
	0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x55, 0x53, 0x54,
	0x4f, 0x4d, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x48, 0x45,
	0x54, 0x41, 0x4e, 0x5f, 0x52, 0x49, 0x56, 0x41, 0x4c, 0x53, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x10,
	0x03, 0x2a, 0xb5, 0x02, 0x0a, 0x0a, 0x49, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x15, 0x0a, 0x11, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54,
	0x5f, 0x53, 0x54, 0x41, 0x52, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x4f, 0x4c, 0x4f, 0x5f,
	0x53, 0x55, 0x52, 0x56, 0x49, 0x56, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x55,
	0x41, 0x4c, 0x5f, 0x53, 0x55, 0x52, 0x56, 0x49, 0x56, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x1c, 0x0a,
	0x18, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x54,
	0x41, 0x52, 0x5f, 0x34, 0x5f, 0x56, 0x53, 0x5f, 0x34, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x4b,
	0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x45, 0x41, 0x54, 0x48, 0x5f, 0x4d,
	0x41, 0x54, 0x43, 0x48, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x45, 0x41, 0x54, 0x48, 0x5f,
	0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x33, 0x5f, 0x56, 0x53, 0x5f, 0x33, 0x10, 0x07, 0x12, 0x08,
	0x0a, 0x04, 0x46, 0x4c, 0x41, 0x47, 0x10, 0x08, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x4f, 0x57, 0x45,
	0x52, 0x10, 0x09, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x52, 0x4f,
	0x59, 0x41, 0x4c, 0x45, 0x10, 0x0c, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x51, 0x55, 0x41, 0x44, 0x5f,
	0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x52, 0x4f, 0x59, 0x41, 0x4c, 0x45, 0x10, 0x0d, 0x12,
	0x15, 0x0a, 0x11, 0x44, 0x55, 0x4f, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x52, 0x4f,
	0x59, 0x41, 0x4c, 0x45, 0x10, 0x0e, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x52, 0x49, 0x4f, 0x5f, 0x42,
	0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x52, 0x4f, 0x41, 0x59, 0x4c, 0x45, 0x10, 0x0f, 0x12, 0x11,
	0x0a, 0x0d, 0x54, 0x48, 0x45, 0x54, 0x41, 0x4e, 0x5f, 0x52, 0x49, 0x56, 0x41, 0x4c, 0x53, 0x10,
	0x14, 0x12, 0x16, 0x0a, 0x09, 0x4e, 0x4f, 0x4e, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x10, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x42, 0x75, 0x0a, 0x0e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0d, 0x47, 0x61, 0x6d,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0b, 0x2e, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x43, 0x50, 0x58, 0xaa,
	0x02, 0x0a, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xca, 0x02, 0x0a, 0x43,
	0x6f, 0x72, 0x65, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xe2, 0x02, 0x16, 0x43, 0x6f, 0x72, 0x65,
	0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0b, 0x43, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_info_proto_rawDescOnce sync.Once
	file_game_info_proto_rawDescData = file_game_info_proto_rawDesc
)

func file_game_info_proto_rawDescGZIP() []byte {
	file_game_info_proto_rawDescOnce.Do(func() {
		file_game_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_info_proto_rawDescData)
	})
	return file_game_info_proto_rawDescData
}

var file_game_info_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_game_info_proto_goTypes = []interface{}{
	(GameMode)(0),   // 0: core.proto.GameMode
	(InGameMode)(0), // 1: core.proto.InGameMode
}
var file_game_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_game_info_proto_init() }
func file_game_info_proto_init() {
	if File_game_info_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_info_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_info_proto_goTypes,
		DependencyIndexes: file_game_info_proto_depIdxs,
		EnumInfos:         file_game_info_proto_enumTypes,
	}.Build()
	File_game_info_proto = out.File
	file_game_info_proto_rawDesc = nil
	file_game_info_proto_goTypes = nil
	file_game_info_proto_depIdxs = nil
}
