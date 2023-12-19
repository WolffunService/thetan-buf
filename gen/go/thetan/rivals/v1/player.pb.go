// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: thetan/rivals/v1/player.proto

package thetan_rivals_v1

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

type CreatePlayersInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListPlayerInfo []*CreatePlayerInfo `protobuf:"bytes,1,rep,name=listPlayerInfo,proto3" json:"listPlayerInfo,omitempty"`
}

func (x *CreatePlayersInfoRequest) Reset() {
	*x = CreatePlayersInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_rivals_v1_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlayersInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlayersInfoRequest) ProtoMessage() {}

func (x *CreatePlayersInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_rivals_v1_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlayersInfoRequest.ProtoReflect.Descriptor instead.
func (*CreatePlayersInfoRequest) Descriptor() ([]byte, []int) {
	return file_thetan_rivals_v1_player_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePlayersInfoRequest) GetListPlayerInfo() []*CreatePlayerInfo {
	if x != nil {
		return x.ListPlayerInfo
	}
	return nil
}

type CreatePlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID string `protobuf:"bytes,1,opt,name=playerID,proto3" json:"playerID,omitempty"`
	Trophy   int32  `protobuf:"varint,2,opt,name=trophy,proto3" json:"trophy,omitempty"`
}

func (x *CreatePlayerInfo) Reset() {
	*x = CreatePlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_rivals_v1_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlayerInfo) ProtoMessage() {}

func (x *CreatePlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_rivals_v1_player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlayerInfo.ProtoReflect.Descriptor instead.
func (*CreatePlayerInfo) Descriptor() ([]byte, []int) {
	return file_thetan_rivals_v1_player_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePlayerInfo) GetPlayerID() string {
	if x != nil {
		return x.PlayerID
	}
	return ""
}

func (x *CreatePlayerInfo) GetTrophy() int32 {
	if x != nil {
		return x.Trophy
	}
	return 0
}

type CreatePlayersInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreatePlayersInfoResponse) Reset() {
	*x = CreatePlayersInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_rivals_v1_player_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlayersInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlayersInfoResponse) ProtoMessage() {}

func (x *CreatePlayersInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_rivals_v1_player_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlayersInfoResponse.ProtoReflect.Descriptor instead.
func (*CreatePlayersInfoResponse) Descriptor() ([]byte, []int) {
	return file_thetan_rivals_v1_player_proto_rawDescGZIP(), []int{2}
}

type UpdatePlayersInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListPlayerInfo []*UpdatePlayerInfo `protobuf:"bytes,1,rep,name=listPlayerInfo,proto3" json:"listPlayerInfo,omitempty"`
}

func (x *UpdatePlayersInfoRequest) Reset() {
	*x = UpdatePlayersInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_rivals_v1_player_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePlayersInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePlayersInfoRequest) ProtoMessage() {}

func (x *UpdatePlayersInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_rivals_v1_player_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePlayersInfoRequest.ProtoReflect.Descriptor instead.
func (*UpdatePlayersInfoRequest) Descriptor() ([]byte, []int) {
	return file_thetan_rivals_v1_player_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePlayersInfoRequest) GetListPlayerInfo() []*UpdatePlayerInfo {
	if x != nil {
		return x.ListPlayerInfo
	}
	return nil
}

type UpdatePlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID string `protobuf:"bytes,1,opt,name=playerID,proto3" json:"playerID,omitempty"`
	Trophy   int32  `protobuf:"varint,2,opt,name=trophy,proto3" json:"trophy,omitempty"`
}

func (x *UpdatePlayerInfo) Reset() {
	*x = UpdatePlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_rivals_v1_player_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePlayerInfo) ProtoMessage() {}

func (x *UpdatePlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_rivals_v1_player_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePlayerInfo.ProtoReflect.Descriptor instead.
func (*UpdatePlayerInfo) Descriptor() ([]byte, []int) {
	return file_thetan_rivals_v1_player_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePlayerInfo) GetPlayerID() string {
	if x != nil {
		return x.PlayerID
	}
	return ""
}

func (x *UpdatePlayerInfo) GetTrophy() int32 {
	if x != nil {
		return x.Trophy
	}
	return 0
}

type UpdatePlayersInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePlayersInfoResponse) Reset() {
	*x = UpdatePlayersInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_rivals_v1_player_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePlayersInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePlayersInfoResponse) ProtoMessage() {}

func (x *UpdatePlayersInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_rivals_v1_player_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePlayersInfoResponse.ProtoReflect.Descriptor instead.
func (*UpdatePlayersInfoResponse) Descriptor() ([]byte, []int) {
	return file_thetan_rivals_v1_player_proto_rawDescGZIP(), []int{5}
}

type FindEligibleGuildForPlayerJoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TownID         string  `protobuf:"bytes,1,opt,name=townID,proto3" json:"townID,omitempty"`
	BotID          string  `protobuf:"bytes,2,opt,name=botID,proto3" json:"botID,omitempty"`
	BotBalanceRate float32 `protobuf:"fixed32,3,opt,name=botBalanceRate,proto3" json:"botBalanceRate,omitempty"`
}

func (x *FindEligibleGuildForPlayerJoinRequest) Reset() {
	*x = FindEligibleGuildForPlayerJoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_rivals_v1_player_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEligibleGuildForPlayerJoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEligibleGuildForPlayerJoinRequest) ProtoMessage() {}

func (x *FindEligibleGuildForPlayerJoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_rivals_v1_player_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEligibleGuildForPlayerJoinRequest.ProtoReflect.Descriptor instead.
func (*FindEligibleGuildForPlayerJoinRequest) Descriptor() ([]byte, []int) {
	return file_thetan_rivals_v1_player_proto_rawDescGZIP(), []int{6}
}

func (x *FindEligibleGuildForPlayerJoinRequest) GetTownID() string {
	if x != nil {
		return x.TownID
	}
	return ""
}

func (x *FindEligibleGuildForPlayerJoinRequest) GetBotID() string {
	if x != nil {
		return x.BotID
	}
	return ""
}

func (x *FindEligibleGuildForPlayerJoinRequest) GetBotBalanceRate() float32 {
	if x != nil {
		return x.BotBalanceRate
	}
	return 0
}

type FindEligibleGuildForPlayerJoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuildID string `protobuf:"bytes,1,opt,name=guildID,proto3" json:"guildID,omitempty"`
}

func (x *FindEligibleGuildForPlayerJoinResponse) Reset() {
	*x = FindEligibleGuildForPlayerJoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_rivals_v1_player_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEligibleGuildForPlayerJoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEligibleGuildForPlayerJoinResponse) ProtoMessage() {}

func (x *FindEligibleGuildForPlayerJoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_rivals_v1_player_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEligibleGuildForPlayerJoinResponse.ProtoReflect.Descriptor instead.
func (*FindEligibleGuildForPlayerJoinResponse) Descriptor() ([]byte, []int) {
	return file_thetan_rivals_v1_player_proto_rawDescGZIP(), []int{7}
}

func (x *FindEligibleGuildForPlayerJoinResponse) GetGuildID() string {
	if x != nil {
		return x.GuildID
	}
	return ""
}

var File_thetan_rivals_v1_player_proto protoreflect.FileDescriptor

var file_thetan_rivals_v1_player_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76,
	0x31, 0x22, 0x66, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a,
	0x0e, 0x6c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72,
	0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x6c, 0x69, 0x73, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x46, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x72, 0x6f,
	0x70, 0x68, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x72, 0x6f, 0x70, 0x68,
	0x79, 0x22, 0x1b, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x66,
	0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x0e, 0x6c, 0x69,
	0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61,
	0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x6c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x46, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x72, 0x6f, 0x70, 0x68, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x72, 0x6f, 0x70, 0x68, 0x79, 0x22, 0x1b,
	0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7d, 0x0a, 0x25, 0x46,
	0x69, 0x6e, 0x64, 0x45, 0x6c, 0x69, 0x67, 0x69, 0x62, 0x6c, 0x65, 0x47, 0x75, 0x69, 0x6c, 0x64,
	0x46, 0x6f, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x77, 0x6e, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x77, 0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x62, 0x6f, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x74,
	0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x62, 0x6f, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x62, 0x6f, 0x74, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x61, 0x74, 0x65, 0x22, 0x42, 0x0a, 0x26, 0x46, 0x69,
	0x6e, 0x64, 0x45, 0x6c, 0x69, 0x67, 0x69, 0x62, 0x6c, 0x65, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x46,
	0x6f, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x44, 0x32, 0x93,
	0x03, 0x0a, 0x19, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x2a, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x11,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x2a, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x95, 0x01, 0x0a,
	0x1e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6c, 0x69, 0x67, 0x69, 0x62, 0x6c, 0x65, 0x47, 0x75, 0x69,
	0x6c, 0x64, 0x46, 0x6f, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e, 0x12,
	0x37, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6c, 0x69, 0x67, 0x69, 0x62, 0x6c, 0x65, 0x47,
	0x75, 0x69, 0x6c, 0x64, 0x46, 0x6f, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4a, 0x6f, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61,
	0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x45, 0x6c, 0x69, 0x67, 0x69, 0x62, 0x6c, 0x65, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x46, 0x6f, 0x72,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0xba, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x68, 0x65,
	0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x74, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x2d, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5f, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x5f, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x54, 0x52, 0x58, 0xaa, 0x02, 0x10, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x2e, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x54, 0x68, 0x65,
	0x74, 0x61, 0x6e, 0x5c, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c,
	0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x54,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x3a, 0x3a, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thetan_rivals_v1_player_proto_rawDescOnce sync.Once
	file_thetan_rivals_v1_player_proto_rawDescData = file_thetan_rivals_v1_player_proto_rawDesc
)

func file_thetan_rivals_v1_player_proto_rawDescGZIP() []byte {
	file_thetan_rivals_v1_player_proto_rawDescOnce.Do(func() {
		file_thetan_rivals_v1_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_thetan_rivals_v1_player_proto_rawDescData)
	})
	return file_thetan_rivals_v1_player_proto_rawDescData
}

var file_thetan_rivals_v1_player_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_thetan_rivals_v1_player_proto_goTypes = []interface{}{
	(*CreatePlayersInfoRequest)(nil),               // 0: thetan.rivals.v1.CreatePlayersInfoRequest
	(*CreatePlayerInfo)(nil),                       // 1: thetan.rivals.v1.CreatePlayerInfo
	(*CreatePlayersInfoResponse)(nil),              // 2: thetan.rivals.v1.CreatePlayersInfoResponse
	(*UpdatePlayersInfoRequest)(nil),               // 3: thetan.rivals.v1.UpdatePlayersInfoRequest
	(*UpdatePlayerInfo)(nil),                       // 4: thetan.rivals.v1.UpdatePlayerInfo
	(*UpdatePlayersInfoResponse)(nil),              // 5: thetan.rivals.v1.UpdatePlayersInfoResponse
	(*FindEligibleGuildForPlayerJoinRequest)(nil),  // 6: thetan.rivals.v1.FindEligibleGuildForPlayerJoinRequest
	(*FindEligibleGuildForPlayerJoinResponse)(nil), // 7: thetan.rivals.v1.FindEligibleGuildForPlayerJoinResponse
}
var file_thetan_rivals_v1_player_proto_depIdxs = []int32{
	1, // 0: thetan.rivals.v1.CreatePlayersInfoRequest.listPlayerInfo:type_name -> thetan.rivals.v1.CreatePlayerInfo
	4, // 1: thetan.rivals.v1.UpdatePlayersInfoRequest.listPlayerInfo:type_name -> thetan.rivals.v1.UpdatePlayerInfo
	0, // 2: thetan.rivals.v1.ThetanRivalsPlayerService.CreatePlayersInfo:input_type -> thetan.rivals.v1.CreatePlayersInfoRequest
	3, // 3: thetan.rivals.v1.ThetanRivalsPlayerService.UpdatePlayersInfo:input_type -> thetan.rivals.v1.UpdatePlayersInfoRequest
	6, // 4: thetan.rivals.v1.ThetanRivalsPlayerService.FindEligibleGuildForPlayerJoin:input_type -> thetan.rivals.v1.FindEligibleGuildForPlayerJoinRequest
	2, // 5: thetan.rivals.v1.ThetanRivalsPlayerService.CreatePlayersInfo:output_type -> thetan.rivals.v1.CreatePlayersInfoResponse
	5, // 6: thetan.rivals.v1.ThetanRivalsPlayerService.UpdatePlayersInfo:output_type -> thetan.rivals.v1.UpdatePlayersInfoResponse
	7, // 7: thetan.rivals.v1.ThetanRivalsPlayerService.FindEligibleGuildForPlayerJoin:output_type -> thetan.rivals.v1.FindEligibleGuildForPlayerJoinResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_thetan_rivals_v1_player_proto_init() }
func file_thetan_rivals_v1_player_proto_init() {
	if File_thetan_rivals_v1_player_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thetan_rivals_v1_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlayersInfoRequest); i {
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
		file_thetan_rivals_v1_player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlayerInfo); i {
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
		file_thetan_rivals_v1_player_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlayersInfoResponse); i {
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
		file_thetan_rivals_v1_player_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePlayersInfoRequest); i {
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
		file_thetan_rivals_v1_player_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePlayerInfo); i {
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
		file_thetan_rivals_v1_player_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePlayersInfoResponse); i {
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
		file_thetan_rivals_v1_player_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEligibleGuildForPlayerJoinRequest); i {
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
		file_thetan_rivals_v1_player_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEligibleGuildForPlayerJoinResponse); i {
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
			RawDescriptor: file_thetan_rivals_v1_player_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thetan_rivals_v1_player_proto_goTypes,
		DependencyIndexes: file_thetan_rivals_v1_player_proto_depIdxs,
		MessageInfos:      file_thetan_rivals_v1_player_proto_msgTypes,
	}.Build()
	File_thetan_rivals_v1_player_proto = out.File
	file_thetan_rivals_v1_player_proto_rawDesc = nil
	file_thetan_rivals_v1_player_proto_goTypes = nil
	file_thetan_rivals_v1_player_proto_depIdxs = nil
}
