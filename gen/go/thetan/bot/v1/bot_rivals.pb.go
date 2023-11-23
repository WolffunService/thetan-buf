// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: thetan/bot/v1/bot_rivals.proto

package thetan_bot_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	v1 "thetan-buf/gen/go/thetan/rivals/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SpenderType int32

const (
	SpenderType_SPENDER_TYPE_NONE    SpenderType = 0
	SpenderType_SPENDER_TYPE_SPENDER SpenderType = 1
)

// Enum value maps for SpenderType.
var (
	SpenderType_name = map[int32]string{
		0: "SPENDER_TYPE_NONE",
		1: "SPENDER_TYPE_SPENDER",
	}
	SpenderType_value = map[string]int32{
		"SPENDER_TYPE_NONE":    0,
		"SPENDER_TYPE_SPENDER": 1,
	}
)

func (x SpenderType) Enum() *SpenderType {
	p := new(SpenderType)
	*p = x
	return p
}

func (x SpenderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SpenderType) Descriptor() protoreflect.EnumDescriptor {
	return file_thetan_bot_v1_bot_rivals_proto_enumTypes[0].Descriptor()
}

func (SpenderType) Type() protoreflect.EnumType {
	return &file_thetan_bot_v1_bot_rivals_proto_enumTypes[0]
}

func (x SpenderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SpenderType.Descriptor instead.
func (SpenderType) EnumDescriptor() ([]byte, []int) {
	return file_thetan_bot_v1_bot_rivals_proto_rawDescGZIP(), []int{0}
}

type BotActionType int32

const (
	BotActionType_None             BotActionType = 0
	BotActionType_CheckClass       BotActionType = 1
	BotActionType_CheckMines       BotActionType = 2
	BotActionType_CheckShop        BotActionType = 3
	BotActionType_CheckContest     BotActionType = 4
	BotActionType_HugZone          BotActionType = 5
	BotActionType_TalkingBench     BotActionType = 6
	BotActionType_RunAround        BotActionType = 7
	BotActionType_PushBack         BotActionType = 8
	BotActionType_DancePlay        BotActionType = 9
	BotActionType_EmotePlay        BotActionType = 10
	BotActionType_VehicleRunAround BotActionType = 11
	BotActionType_VehiclePushBack  BotActionType = 12
	BotActionType_ReactProfile     BotActionType = 13
	BotActionType_ReactGuildbook   BotActionType = 14
	// WriteGuestbook= 00;
	// Email= 00;
	BotActionType_AddFriend       BotActionType = 15
	BotActionType_UseVehicle      BotActionType = 16
	BotActionType_UseFlycraft     BotActionType = 17
	BotActionType_UseDance        BotActionType = 18
	BotActionType_UseEmoji        BotActionType = 19
	BotActionType_ChangeThenion   BotActionType = 20
	BotActionType_ChangeBackbling BotActionType = 21
	BotActionType_ChangeFootprint BotActionType = 22
	BotActionType_ChangeGlow      BotActionType = 23
	BotActionType_ChangeVoice     BotActionType = 24
	BotActionType_ChangeFlycraft  BotActionType = 25
	BotActionType_ChangeVehicle   BotActionType = 26
	BotActionType_CopyThenion     BotActionType = 27
	BotActionType_CopyVehicle     BotActionType = 28
	BotActionType_CopyFlycraft    BotActionType = 29
	BotActionType_Findmatch       BotActionType = 30
	BotActionType_InviteParty     BotActionType = 31
	BotActionType_ShareParty      BotActionType = 32
	BotActionType_StayIdle        BotActionType = 33
)

// Enum value maps for BotActionType.
var (
	BotActionType_name = map[int32]string{
		0:  "None",
		1:  "CheckClass",
		2:  "CheckMines",
		3:  "CheckShop",
		4:  "CheckContest",
		5:  "HugZone",
		6:  "TalkingBench",
		7:  "RunAround",
		8:  "PushBack",
		9:  "DancePlay",
		10: "EmotePlay",
		11: "VehicleRunAround",
		12: "VehiclePushBack",
		13: "ReactProfile",
		14: "ReactGuildbook",
		15: "AddFriend",
		16: "UseVehicle",
		17: "UseFlycraft",
		18: "UseDance",
		19: "UseEmoji",
		20: "ChangeThenion",
		21: "ChangeBackbling",
		22: "ChangeFootprint",
		23: "ChangeGlow",
		24: "ChangeVoice",
		25: "ChangeFlycraft",
		26: "ChangeVehicle",
		27: "CopyThenion",
		28: "CopyVehicle",
		29: "CopyFlycraft",
		30: "Findmatch",
		31: "InviteParty",
		32: "ShareParty",
		33: "StayIdle",
	}
	BotActionType_value = map[string]int32{
		"None":             0,
		"CheckClass":       1,
		"CheckMines":       2,
		"CheckShop":        3,
		"CheckContest":     4,
		"HugZone":          5,
		"TalkingBench":     6,
		"RunAround":        7,
		"PushBack":         8,
		"DancePlay":        9,
		"EmotePlay":        10,
		"VehicleRunAround": 11,
		"VehiclePushBack":  12,
		"ReactProfile":     13,
		"ReactGuildbook":   14,
		"AddFriend":        15,
		"UseVehicle":       16,
		"UseFlycraft":      17,
		"UseDance":         18,
		"UseEmoji":         19,
		"ChangeThenion":    20,
		"ChangeBackbling":  21,
		"ChangeFootprint":  22,
		"ChangeGlow":       23,
		"ChangeVoice":      24,
		"ChangeFlycraft":   25,
		"ChangeVehicle":    26,
		"CopyThenion":      27,
		"CopyVehicle":      28,
		"CopyFlycraft":     29,
		"Findmatch":        30,
		"InviteParty":      31,
		"ShareParty":       32,
		"StayIdle":         33,
	}
)

func (x BotActionType) Enum() *BotActionType {
	p := new(BotActionType)
	*p = x
	return p
}

func (x BotActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BotActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_thetan_bot_v1_bot_rivals_proto_enumTypes[1].Descriptor()
}

func (BotActionType) Type() protoreflect.EnumType {
	return &file_thetan_bot_v1_bot_rivals_proto_enumTypes[1]
}

func (x BotActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BotActionType.Descriptor instead.
func (BotActionType) EnumDescriptor() ([]byte, []int) {
	return file_thetan_bot_v1_bot_rivals_proto_rawDescGZIP(), []int{1}
}

type FetchLobbyBotsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TownID string `protobuf:"bytes,1,opt,name=townID,proto3" json:"townID,omitempty"`
}

func (x *FetchLobbyBotsRequest) Reset() {
	*x = FetchLobbyBotsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_bot_v1_bot_rivals_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchLobbyBotsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchLobbyBotsRequest) ProtoMessage() {}

func (x *FetchLobbyBotsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_bot_v1_bot_rivals_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchLobbyBotsRequest.ProtoReflect.Descriptor instead.
func (*FetchLobbyBotsRequest) Descriptor() ([]byte, []int) {
	return file_thetan_bot_v1_bot_rivals_proto_rawDescGZIP(), []int{0}
}

func (x *FetchLobbyBotsRequest) GetTownID() string {
	if x != nil {
		return x.TownID
	}
	return ""
}

type FetchLobbyBotsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bots           []*LobbyBotInfo `protobuf:"bytes,1,rep,name=bots,proto3" json:"bots,omitempty"`
	RandomRateSolo float32         `protobuf:"fixed32,2,opt,name=randomRateSolo,proto3" json:"randomRateSolo,omitempty"`
	RandomRatePair float32         `protobuf:"fixed32,3,opt,name=randomRatePair,proto3" json:"randomRatePair,omitempty"`
}

func (x *FetchLobbyBotsResponse) Reset() {
	*x = FetchLobbyBotsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_bot_v1_bot_rivals_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchLobbyBotsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchLobbyBotsResponse) ProtoMessage() {}

func (x *FetchLobbyBotsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_bot_v1_bot_rivals_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchLobbyBotsResponse.ProtoReflect.Descriptor instead.
func (*FetchLobbyBotsResponse) Descriptor() ([]byte, []int) {
	return file_thetan_bot_v1_bot_rivals_proto_rawDescGZIP(), []int{1}
}

func (x *FetchLobbyBotsResponse) GetBots() []*LobbyBotInfo {
	if x != nil {
		return x.Bots
	}
	return nil
}

func (x *FetchLobbyBotsResponse) GetRandomRateSolo() float32 {
	if x != nil {
		return x.RandomRateSolo
	}
	return 0
}

func (x *FetchLobbyBotsResponse) GetRandomRatePair() float32 {
	if x != nil {
		return x.RandomRatePair
	}
	return 0
}

type LobbyBotInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile        *v1.UserProfileResponse        `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	SelectedMinion *v1.UserSelectedMinionResponse `protobuf:"bytes,2,opt,name=selectedMinion,proto3" json:"selectedMinion,omitempty"`
	ExpiresAt      int64                          `protobuf:"varint,3,opt,name=expiresAt,proto3" json:"expiresAt,omitempty"`
	IsSpender      bool                           `protobuf:"varint,4,opt,name=isSpender,proto3" json:"isSpender,omitempty"`
}

func (x *LobbyBotInfo) Reset() {
	*x = LobbyBotInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_bot_v1_bot_rivals_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LobbyBotInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LobbyBotInfo) ProtoMessage() {}

func (x *LobbyBotInfo) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_bot_v1_bot_rivals_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LobbyBotInfo.ProtoReflect.Descriptor instead.
func (*LobbyBotInfo) Descriptor() ([]byte, []int) {
	return file_thetan_bot_v1_bot_rivals_proto_rawDescGZIP(), []int{2}
}

func (x *LobbyBotInfo) GetProfile() *v1.UserProfileResponse {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *LobbyBotInfo) GetSelectedMinion() *v1.UserSelectedMinionResponse {
	if x != nil {
		return x.SelectedMinion
	}
	return nil
}

func (x *LobbyBotInfo) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *LobbyBotInfo) GetIsSpender() bool {
	if x != nil {
		return x.IsSpender
	}
	return false
}

type BotInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BotID       string      `protobuf:"bytes,1,opt,name=botID,proto3" json:"botID,omitempty"`
	ExpiresAt   int64       `protobuf:"varint,2,opt,name=expiresAt,proto3" json:"expiresAt,omitempty"`
	SpenderType SpenderType `protobuf:"varint,3,opt,name=spenderType,proto3,enum=thetan.rivals.v1.SpenderType" json:"spenderType,omitempty"`
}

func (x *BotInfo) Reset() {
	*x = BotInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_bot_v1_bot_rivals_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BotInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotInfo) ProtoMessage() {}

func (x *BotInfo) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_bot_v1_bot_rivals_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotInfo.ProtoReflect.Descriptor instead.
func (*BotInfo) Descriptor() ([]byte, []int) {
	return file_thetan_bot_v1_bot_rivals_proto_rawDescGZIP(), []int{3}
}

func (x *BotInfo) GetBotID() string {
	if x != nil {
		return x.BotID
	}
	return ""
}

func (x *BotInfo) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *BotInfo) GetSpenderType() SpenderType {
	if x != nil {
		return x.SpenderType
	}
	return SpenderType_SPENDER_TYPE_NONE
}

type LobbyBotActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BotId    string        `protobuf:"bytes,1,opt,name=botId,proto3" json:"botId,omitempty"`
	TargetId string        `protobuf:"bytes,2,opt,name=targetId,proto3" json:"targetId,omitempty"`
	Action   BotActionType `protobuf:"varint,3,opt,name=action,proto3,enum=thetan.rivals.v1.BotActionType" json:"action,omitempty"` //more
}

func (x *LobbyBotActionRequest) Reset() {
	*x = LobbyBotActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_bot_v1_bot_rivals_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LobbyBotActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LobbyBotActionRequest) ProtoMessage() {}

func (x *LobbyBotActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_bot_v1_bot_rivals_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LobbyBotActionRequest.ProtoReflect.Descriptor instead.
func (*LobbyBotActionRequest) Descriptor() ([]byte, []int) {
	return file_thetan_bot_v1_bot_rivals_proto_rawDescGZIP(), []int{4}
}

func (x *LobbyBotActionRequest) GetBotId() string {
	if x != nil {
		return x.BotId
	}
	return ""
}

func (x *LobbyBotActionRequest) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *LobbyBotActionRequest) GetAction() BotActionType {
	if x != nil {
		return x.Action
	}
	return BotActionType_None
}

type LobbyBotActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*LobbyBotActionResponse_Id
	//	*LobbyBotActionResponse_Minion
	Data isLobbyBotActionResponse_Data `protobuf_oneof:"data"`
}

func (x *LobbyBotActionResponse) Reset() {
	*x = LobbyBotActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_bot_v1_bot_rivals_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LobbyBotActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LobbyBotActionResponse) ProtoMessage() {}

func (x *LobbyBotActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_bot_v1_bot_rivals_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LobbyBotActionResponse.ProtoReflect.Descriptor instead.
func (*LobbyBotActionResponse) Descriptor() ([]byte, []int) {
	return file_thetan_bot_v1_bot_rivals_proto_rawDescGZIP(), []int{5}
}

func (m *LobbyBotActionResponse) GetData() isLobbyBotActionResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *LobbyBotActionResponse) GetId() int64 {
	if x, ok := x.GetData().(*LobbyBotActionResponse_Id); ok {
		return x.Id
	}
	return 0
}

func (x *LobbyBotActionResponse) GetMinion() *v1.UserSelectedMinionResponse {
	if x, ok := x.GetData().(*LobbyBotActionResponse_Minion); ok {
		return x.Minion
	}
	return nil
}

type isLobbyBotActionResponse_Data interface {
	isLobbyBotActionResponse_Data()
}

type LobbyBotActionResponse_Id struct {
	Id int64 `protobuf:"varint,1,opt,name=id,proto3,oneof"`
}

type LobbyBotActionResponse_Minion struct {
	Minion *v1.UserSelectedMinionResponse `protobuf:"bytes,2,opt,name=minion,proto3,oneof"`
}

func (*LobbyBotActionResponse_Id) isLobbyBotActionResponse_Data() {}

func (*LobbyBotActionResponse_Minion) isLobbyBotActionResponse_Data() {}

var File_thetan_bot_v1_bot_rivals_proto protoreflect.FileDescriptor

var file_thetan_bot_v1_bot_rivals_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x62, 0x6f, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x6f, 0x74, 0x5f, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x25, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x72, 0x69, 0x76, 0x61, 0x6c,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x69, 0x76,
	0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x15, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x42, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x77, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x77, 0x6e, 0x49, 0x44, 0x22, 0x9c, 0x01, 0x0a, 0x16, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x42, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x62, 0x6f, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76,
	0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x42, 0x6f, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x62, 0x6f, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x61, 0x6e,
	0x64, 0x6f, 0x6d, 0x52, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6c, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0e, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x52, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6c,
	0x6f, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x52, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x69, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x72, 0x61, 0x6e, 0x64, 0x6f,
	0x6d, 0x52, 0x61, 0x74, 0x65, 0x50, 0x61, 0x69, 0x72, 0x22, 0xe1, 0x01, 0x0a, 0x0c, 0x4c, 0x6f,
	0x62, 0x62, 0x79, 0x42, 0x6f, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3f, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x54, 0x0a, 0x0e, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76,
	0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x4d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x0e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4d, 0x69, 0x6e, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x73, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x7e, 0x0a,
	0x07, 0x42, 0x6f, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6f, 0x74, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x74, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x3f, 0x0a, 0x0b,
	0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1d, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0b, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x82, 0x01,
	0x0a, 0x15, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x42, 0x6f, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6f, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x74, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x7a, 0x0a, 0x16, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x42, 0x6f, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x46,
	0x0a, 0x06, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4d, 0x69,
	0x6e, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x06,
	0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x3e,
	0x0a, 0x0b, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a,
	0x11, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x10, 0x01, 0x2a, 0xc8,
	0x04, 0x0a, 0x0d, 0x42, 0x6f, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x4d, 0x69, 0x6e, 0x65, 0x73, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x53, 0x68, 0x6f, 0x70, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x48,
	0x75, 0x67, 0x5a, 0x6f, 0x6e, 0x65, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x61, 0x6c, 0x6b,
	0x69, 0x6e, 0x67, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x75,
	0x6e, 0x41, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x75, 0x73,
	0x68, 0x42, 0x61, 0x63, 0x6b, 0x10, 0x08, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x61, 0x6e, 0x63, 0x65,
	0x50, 0x6c, 0x61, 0x79, 0x10, 0x09, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x10, 0x0a, 0x12, 0x14, 0x0a, 0x10, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x52, 0x75, 0x6e, 0x41, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x0b, 0x12, 0x13, 0x0a, 0x0f, 0x56,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x75, 0x73, 0x68, 0x42, 0x61, 0x63, 0x6b, 0x10, 0x0c,
	0x12, 0x10, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x10, 0x0d, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x65, 0x61, 0x63, 0x74, 0x47, 0x75, 0x69, 0x6c, 0x64,
	0x62, 0x6f, 0x6f, 0x6b, 0x10, 0x0e, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x10, 0x0f, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x56, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x10, 0x10, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x46, 0x6c, 0x79, 0x63,
	0x72, 0x61, 0x66, 0x74, 0x10, 0x11, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x44, 0x61, 0x6e,
	0x63, 0x65, 0x10, 0x12, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x45, 0x6d, 0x6f, 0x6a, 0x69,
	0x10, 0x13, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x68, 0x65, 0x6e,
	0x69, 0x6f, 0x6e, 0x10, 0x14, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x42,
	0x61, 0x63, 0x6b, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x10, 0x15, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x46, 0x6f, 0x6f, 0x74, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x10, 0x16, 0x12,
	0x0e, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x47, 0x6c, 0x6f, 0x77, 0x10, 0x17, 0x12,
	0x0f, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x10, 0x18,
	0x12, 0x12, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x46, 0x6c, 0x79, 0x63, 0x72, 0x61,
	0x66, 0x74, 0x10, 0x19, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x56, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x10, 0x1a, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x6f, 0x70, 0x79, 0x54,
	0x68, 0x65, 0x6e, 0x69, 0x6f, 0x6e, 0x10, 0x1b, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x6f, 0x70, 0x79,
	0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x10, 0x1c, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x6f, 0x70,
	0x79, 0x46, 0x6c, 0x79, 0x63, 0x72, 0x61, 0x66, 0x74, 0x10, 0x1d, 0x12, 0x0d, 0x0a, 0x09, 0x46,
	0x69, 0x6e, 0x64, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x10, 0x1e, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x10, 0x1f, 0x12, 0x0e, 0x0a, 0x0a, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x10, 0x20, 0x12, 0x0c, 0x0a, 0x08, 0x53,
	0x74, 0x61, 0x79, 0x49, 0x64, 0x6c, 0x65, 0x10, 0x21, 0x32, 0xe0, 0x01, 0x0a, 0x10, 0x42, 0x6f,
	0x74, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65,
	0x0a, 0x0e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x42, 0x6f, 0x74, 0x73,
	0x12, 0x27, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x42, 0x6f,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x42, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x0e, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x42, 0x6f,
	0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x62, 0x62, 0x79,
	0x42, 0x6f, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x42, 0x6f, 0x74, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xb7, 0x01, 0x0a,
	0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61,
	0x6c, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x42, 0x6f, 0x74, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2d, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2d,
	0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x68, 0x65, 0x74, 0x61,
	0x6e, 0x2f, 0x62, 0x6f, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5f,
	0x62, 0x6f, 0x74, 0x5f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x52, 0x58, 0xaa, 0x02, 0x10, 0x54,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x10, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x1c, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x52, 0x69, 0x76, 0x61,
	0x6c, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x12, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x3a, 0x3a, 0x52, 0x69, 0x76, 0x61,
	0x6c, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thetan_bot_v1_bot_rivals_proto_rawDescOnce sync.Once
	file_thetan_bot_v1_bot_rivals_proto_rawDescData = file_thetan_bot_v1_bot_rivals_proto_rawDesc
)

func file_thetan_bot_v1_bot_rivals_proto_rawDescGZIP() []byte {
	file_thetan_bot_v1_bot_rivals_proto_rawDescOnce.Do(func() {
		file_thetan_bot_v1_bot_rivals_proto_rawDescData = protoimpl.X.CompressGZIP(file_thetan_bot_v1_bot_rivals_proto_rawDescData)
	})
	return file_thetan_bot_v1_bot_rivals_proto_rawDescData
}

var file_thetan_bot_v1_bot_rivals_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_thetan_bot_v1_bot_rivals_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_thetan_bot_v1_bot_rivals_proto_goTypes = []interface{}{
	(SpenderType)(0),                      // 0: thetan.rivals.v1.SpenderType
	(BotActionType)(0),                    // 1: thetan.rivals.v1.BotActionType
	(*FetchLobbyBotsRequest)(nil),         // 2: thetan.rivals.v1.FetchLobbyBotsRequest
	(*FetchLobbyBotsResponse)(nil),        // 3: thetan.rivals.v1.FetchLobbyBotsResponse
	(*LobbyBotInfo)(nil),                  // 4: thetan.rivals.v1.LobbyBotInfo
	(*BotInfo)(nil),                       // 5: thetan.rivals.v1.BotInfo
	(*LobbyBotActionRequest)(nil),         // 6: thetan.rivals.v1.LobbyBotActionRequest
	(*LobbyBotActionResponse)(nil),        // 7: thetan.rivals.v1.LobbyBotActionResponse
	(*v1.UserProfileResponse)(nil),        // 8: thetan.rivals.v1.UserProfileResponse
	(*v1.UserSelectedMinionResponse)(nil), // 9: thetan.rivals.v1.UserSelectedMinionResponse
}
var file_thetan_bot_v1_bot_rivals_proto_depIdxs = []int32{
	4, // 0: thetan.rivals.v1.FetchLobbyBotsResponse.bots:type_name -> thetan.rivals.v1.LobbyBotInfo
	8, // 1: thetan.rivals.v1.LobbyBotInfo.profile:type_name -> thetan.rivals.v1.UserProfileResponse
	9, // 2: thetan.rivals.v1.LobbyBotInfo.selectedMinion:type_name -> thetan.rivals.v1.UserSelectedMinionResponse
	0, // 3: thetan.rivals.v1.BotInfo.spenderType:type_name -> thetan.rivals.v1.SpenderType
	1, // 4: thetan.rivals.v1.LobbyBotActionRequest.action:type_name -> thetan.rivals.v1.BotActionType
	9, // 5: thetan.rivals.v1.LobbyBotActionResponse.minion:type_name -> thetan.rivals.v1.UserSelectedMinionResponse
	2, // 6: thetan.rivals.v1.BotRivalsService.FetchLobbyBots:input_type -> thetan.rivals.v1.FetchLobbyBotsRequest
	6, // 7: thetan.rivals.v1.BotRivalsService.LobbyBotAction:input_type -> thetan.rivals.v1.LobbyBotActionRequest
	3, // 8: thetan.rivals.v1.BotRivalsService.FetchLobbyBots:output_type -> thetan.rivals.v1.FetchLobbyBotsResponse
	7, // 9: thetan.rivals.v1.BotRivalsService.LobbyBotAction:output_type -> thetan.rivals.v1.LobbyBotActionResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_thetan_bot_v1_bot_rivals_proto_init() }
func file_thetan_bot_v1_bot_rivals_proto_init() {
	if File_thetan_bot_v1_bot_rivals_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thetan_bot_v1_bot_rivals_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchLobbyBotsRequest); i {
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
		file_thetan_bot_v1_bot_rivals_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchLobbyBotsResponse); i {
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
		file_thetan_bot_v1_bot_rivals_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LobbyBotInfo); i {
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
		file_thetan_bot_v1_bot_rivals_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BotInfo); i {
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
		file_thetan_bot_v1_bot_rivals_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LobbyBotActionRequest); i {
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
		file_thetan_bot_v1_bot_rivals_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LobbyBotActionResponse); i {
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
	file_thetan_bot_v1_bot_rivals_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*LobbyBotActionResponse_Id)(nil),
		(*LobbyBotActionResponse_Minion)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_thetan_bot_v1_bot_rivals_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thetan_bot_v1_bot_rivals_proto_goTypes,
		DependencyIndexes: file_thetan_bot_v1_bot_rivals_proto_depIdxs,
		EnumInfos:         file_thetan_bot_v1_bot_rivals_proto_enumTypes,
		MessageInfos:      file_thetan_bot_v1_bot_rivals_proto_msgTypes,
	}.Build()
	File_thetan_bot_v1_bot_rivals_proto = out.File
	file_thetan_bot_v1_bot_rivals_proto_rawDesc = nil
	file_thetan_bot_v1_bot_rivals_proto_goTypes = nil
	file_thetan_bot_v1_bot_rivals_proto_depIdxs = nil
}
