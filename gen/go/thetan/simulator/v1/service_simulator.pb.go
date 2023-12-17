// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: thetan/simulator/v1/service_simulator.proto

package thetan_simulator_v1

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

type SimulateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchId    string `protobuf:"bytes,1,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
	GameInputs []byte `protobuf:"bytes,2,opt,name=game_inputs,json=gameInputs,proto3" json:"game_inputs,omitempty"`
}

func (x *SimulateRequest) Reset() {
	*x = SimulateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_simulator_v1_service_simulator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulateRequest) ProtoMessage() {}

func (x *SimulateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_simulator_v1_service_simulator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulateRequest.ProtoReflect.Descriptor instead.
func (*SimulateRequest) Descriptor() ([]byte, []int) {
	return file_thetan_simulator_v1_service_simulator_proto_rawDescGZIP(), []int{0}
}

func (x *SimulateRequest) GetMatchId() string {
	if x != nil {
		return x.MatchId
	}
	return ""
}

func (x *SimulateRequest) GetGameInputs() []byte {
	if x != nil {
		return x.GameInputs
	}
	return nil
}

type SimulateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchId string            `protobuf:"bytes,1,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
	Players []*SimulatePlayer `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *SimulateResponse) Reset() {
	*x = SimulateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_simulator_v1_service_simulator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulateResponse) ProtoMessage() {}

func (x *SimulateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_simulator_v1_service_simulator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulateResponse.ProtoReflect.Descriptor instead.
func (*SimulateResponse) Descriptor() ([]byte, []int) {
	return file_thetan_simulator_v1_service_simulator_proto_rawDescGZIP(), []int{1}
}

func (x *SimulateResponse) GetMatchId() string {
	if x != nil {
		return x.MatchId
	}
	return ""
}

func (x *SimulateResponse) GetPlayers() []*SimulatePlayer {
	if x != nil {
		return x.Players
	}
	return nil
}

type SimulatePlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId   string    `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	BattleRank int32     `protobuf:"varint,2,opt,name=battle_rank,json=battleRank,proto3" json:"battle_rank,omitempty"`
	RoundTimes []float32 `protobuf:"fixed32,3,rep,packed,name=round_times,json=roundTimes,proto3" json:"round_times,omitempty"`
}

func (x *SimulatePlayer) Reset() {
	*x = SimulatePlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_simulator_v1_service_simulator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulatePlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulatePlayer) ProtoMessage() {}

func (x *SimulatePlayer) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_simulator_v1_service_simulator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulatePlayer.ProtoReflect.Descriptor instead.
func (*SimulatePlayer) Descriptor() ([]byte, []int) {
	return file_thetan_simulator_v1_service_simulator_proto_rawDescGZIP(), []int{2}
}

func (x *SimulatePlayer) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *SimulatePlayer) GetBattleRank() int32 {
	if x != nil {
		return x.BattleRank
	}
	return 0
}

func (x *SimulatePlayer) GetRoundTimes() []float32 {
	if x != nil {
		return x.RoundTimes
	}
	return nil
}

var File_thetan_simulator_v1_service_simulator_proto protoreflect.FileDescriptor

var file_thetan_simulator_v1_service_simulator_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x69,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x74,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x76, 0x31, 0x22, 0x4d, 0x0a, 0x0f, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x73, 0x22, 0x6c, 0x0a, 0x10, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64,
	0x12, 0x3d, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22,
	0x6f, 0x0a, 0x0e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x02, 0x52, 0x0a, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x32, 0x73, 0x0a, 0x16, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x08, 0x53, 0x69,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e,
	0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xd9, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76,
	0x31, 0x42, 0x15, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x74, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x2d, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x76,
	0x31, 0x3b, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x53, 0x58, 0xaa, 0x02, 0x13, 0x54, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x13, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x53, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1f, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x5c, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x54, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x3a, 0x3a, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thetan_simulator_v1_service_simulator_proto_rawDescOnce sync.Once
	file_thetan_simulator_v1_service_simulator_proto_rawDescData = file_thetan_simulator_v1_service_simulator_proto_rawDesc
)

func file_thetan_simulator_v1_service_simulator_proto_rawDescGZIP() []byte {
	file_thetan_simulator_v1_service_simulator_proto_rawDescOnce.Do(func() {
		file_thetan_simulator_v1_service_simulator_proto_rawDescData = protoimpl.X.CompressGZIP(file_thetan_simulator_v1_service_simulator_proto_rawDescData)
	})
	return file_thetan_simulator_v1_service_simulator_proto_rawDescData
}

var file_thetan_simulator_v1_service_simulator_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_thetan_simulator_v1_service_simulator_proto_goTypes = []interface{}{
	(*SimulateRequest)(nil),  // 0: thetan.simulator.v1.SimulateRequest
	(*SimulateResponse)(nil), // 1: thetan.simulator.v1.SimulateResponse
	(*SimulatePlayer)(nil),   // 2: thetan.simulator.v1.SimulatePlayer
}
var file_thetan_simulator_v1_service_simulator_proto_depIdxs = []int32{
	2, // 0: thetan.simulator.v1.SimulateResponse.players:type_name -> thetan.simulator.v1.SimulatePlayer
	0, // 1: thetan.simulator.v1.ThetanSimulatorService.Simulate:input_type -> thetan.simulator.v1.SimulateRequest
	1, // 2: thetan.simulator.v1.ThetanSimulatorService.Simulate:output_type -> thetan.simulator.v1.SimulateResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_thetan_simulator_v1_service_simulator_proto_init() }
func file_thetan_simulator_v1_service_simulator_proto_init() {
	if File_thetan_simulator_v1_service_simulator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thetan_simulator_v1_service_simulator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulateRequest); i {
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
		file_thetan_simulator_v1_service_simulator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulateResponse); i {
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
		file_thetan_simulator_v1_service_simulator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulatePlayer); i {
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
			RawDescriptor: file_thetan_simulator_v1_service_simulator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thetan_simulator_v1_service_simulator_proto_goTypes,
		DependencyIndexes: file_thetan_simulator_v1_service_simulator_proto_depIdxs,
		MessageInfos:      file_thetan_simulator_v1_service_simulator_proto_msgTypes,
	}.Build()
	File_thetan_simulator_v1_service_simulator_proto = out.File
	file_thetan_simulator_v1_service_simulator_proto_rawDesc = nil
	file_thetan_simulator_v1_service_simulator_proto_goTypes = nil
	file_thetan_simulator_v1_service_simulator_proto_depIdxs = nil
}
