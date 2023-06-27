// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: thetan/gateway/v1/gs_sync.proto

package thetan_gateway_v1

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

type GameName int32

const (
	GameName_Rival    GameName = 0
	GameName_Immortal GameName = 1
	GameName_Arena    GameName = 2 // Thêm các giá trị enum tương ứng với các game khác
)

// Enum value maps for GameName.
var (
	GameName_name = map[int32]string{
		0: "Rival",
		1: "Immortal",
		2: "Arena",
	}
	GameName_value = map[string]int32{
		"Rival":    0,
		"Immortal": 1,
		"Arena":    2,
	}
)

func (x GameName) Enum() *GameName {
	p := new(GameName)
	*p = x
	return p
}

func (x GameName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameName) Descriptor() protoreflect.EnumDescriptor {
	return file_thetan_gateway_v1_gs_sync_proto_enumTypes[0].Descriptor()
}

func (GameName) Type() protoreflect.EnumType {
	return &file_thetan_gateway_v1_gs_sync_proto_enumTypes[0]
}

func (x GameName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameName.Descriptor instead.
func (GameName) EnumDescriptor() ([]byte, []int) {
	return file_thetan_gateway_v1_gs_sync_proto_rawDescGZIP(), []int{0}
}

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//Bo qua field nay, 1 project 1 game thoi
	GameName       GameName    `protobuf:"varint,1,opt,name=game_name,json=gameName,proto3,enum=thetan.gateway.v1.GameName" json:"game_name,omitempty"`
	GameServerName string      `protobuf:"bytes,2,opt,name=game_server_name,json=gameServerName,proto3" json:"game_server_name,omitempty"`
	Host           string      `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Port           int32       `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	Region         string      `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"` //AS, NA, EU
	Rooms          []*RoomInfo `protobuf:"bytes,6,rep,name=rooms,proto3" json:"rooms,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_gateway_v1_gs_sync_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_gateway_v1_gs_sync_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_thetan_gateway_v1_gs_sync_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetGameName() GameName {
	if x != nil {
		return x.GameName
	}
	return GameName_Rival
}

func (x *PingRequest) GetGameServerName() string {
	if x != nil {
		return x.GameServerName
	}
	return ""
}

func (x *PingRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *PingRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *PingRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *PingRequest) GetRooms() []*RoomInfo {
	if x != nil {
		return x.Rooms
	}
	return nil
}

type PlayerConnectedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameName       GameName `protobuf:"varint,1,opt,name=game_name,json=gameName,proto3,enum=thetan.gateway.v1.GameName" json:"game_name,omitempty"`
	GameServerName string   `protobuf:"bytes,2,opt,name=game_server_name,json=gameServerName,proto3" json:"game_server_name,omitempty"`
	PlayerId       string   `protobuf:"bytes,3,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	PlayerCount    int32    `protobuf:"varint,4,opt,name=player_count,json=playerCount,proto3" json:"player_count,omitempty"`
}

func (x *PlayerConnectedRequest) Reset() {
	*x = PlayerConnectedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_gateway_v1_gs_sync_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerConnectedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerConnectedRequest) ProtoMessage() {}

func (x *PlayerConnectedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_gateway_v1_gs_sync_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerConnectedRequest.ProtoReflect.Descriptor instead.
func (*PlayerConnectedRequest) Descriptor() ([]byte, []int) {
	return file_thetan_gateway_v1_gs_sync_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerConnectedRequest) GetGameName() GameName {
	if x != nil {
		return x.GameName
	}
	return GameName_Rival
}

func (x *PlayerConnectedRequest) GetGameServerName() string {
	if x != nil {
		return x.GameServerName
	}
	return ""
}

func (x *PlayerConnectedRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerConnectedRequest) GetPlayerCount() int32 {
	if x != nil {
		return x.PlayerCount
	}
	return 0
}

type PlayerDisconnectedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameName       GameName `protobuf:"varint,1,opt,name=game_name,json=gameName,proto3,enum=thetan.gateway.v1.GameName" json:"game_name,omitempty"`
	GameServerName string   `protobuf:"bytes,2,opt,name=game_server_name,json=gameServerName,proto3" json:"game_server_name,omitempty"`
	PlayerId       string   `protobuf:"bytes,3,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	PlayerCount    int32    `protobuf:"varint,4,opt,name=player_count,json=playerCount,proto3" json:"player_count,omitempty"`
}

func (x *PlayerDisconnectedRequest) Reset() {
	*x = PlayerDisconnectedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_gateway_v1_gs_sync_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerDisconnectedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerDisconnectedRequest) ProtoMessage() {}

func (x *PlayerDisconnectedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_gateway_v1_gs_sync_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerDisconnectedRequest.ProtoReflect.Descriptor instead.
func (*PlayerDisconnectedRequest) Descriptor() ([]byte, []int) {
	return file_thetan_gateway_v1_gs_sync_proto_rawDescGZIP(), []int{2}
}

func (x *PlayerDisconnectedRequest) GetGameName() GameName {
	if x != nil {
		return x.GameName
	}
	return GameName_Rival
}

func (x *PlayerDisconnectedRequest) GetGameServerName() string {
	if x != nil {
		return x.GameServerName
	}
	return ""
}

func (x *PlayerDisconnectedRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerDisconnectedRequest) GetPlayerCount() int32 {
	if x != nil {
		return x.PlayerCount
	}
	return 0
}

type RoomDestroyedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameName       GameName `protobuf:"varint,1,opt,name=game_name,json=gameName,proto3,enum=thetan.gateway.v1.GameName" json:"game_name,omitempty"`
	GameServerName string   `protobuf:"bytes,2,opt,name=game_server_name,json=gameServerName,proto3" json:"game_server_name,omitempty"`
	RoomId         string   `protobuf:"bytes,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *RoomDestroyedRequest) Reset() {
	*x = RoomDestroyedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_gateway_v1_gs_sync_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomDestroyedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomDestroyedRequest) ProtoMessage() {}

func (x *RoomDestroyedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_gateway_v1_gs_sync_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomDestroyedRequest.ProtoReflect.Descriptor instead.
func (*RoomDestroyedRequest) Descriptor() ([]byte, []int) {
	return file_thetan_gateway_v1_gs_sync_proto_rawDescGZIP(), []int{3}
}

func (x *RoomDestroyedRequest) GetGameName() GameName {
	if x != nil {
		return x.GameName
	}
	return GameName_Rival
}

func (x *RoomDestroyedRequest) GetGameServerName() string {
	if x != nil {
		return x.GameServerName
	}
	return ""
}

func (x *RoomDestroyedRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_gateway_v1_gs_sync_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_gateway_v1_gs_sync_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_thetan_gateway_v1_gs_sync_proto_rawDescGZIP(), []int{4}
}

func (x *PingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PlayerStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlayerStatusResponse) Reset() {
	*x = PlayerStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_gateway_v1_gs_sync_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerStatusResponse) ProtoMessage() {}

func (x *PlayerStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_gateway_v1_gs_sync_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerStatusResponse.ProtoReflect.Descriptor instead.
func (*PlayerStatusResponse) Descriptor() ([]byte, []int) {
	return file_thetan_gateway_v1_gs_sync_proto_rawDescGZIP(), []int{5}
}

type RoomInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId           string   `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	RoomNumber       int32    `protobuf:"varint,2,opt,name=room_number,json=roomNumber,proto3" json:"room_number,omitempty"`
	TownId           string   `protobuf:"bytes,3,opt,name=town_id,json=townId,proto3" json:"town_id,omitempty"`
	ConnectedPlayers []string `protobuf:"bytes,4,rep,name=connected_players,json=connectedPlayers,proto3" json:"connected_players,omitempty"`
}

func (x *RoomInfo) Reset() {
	*x = RoomInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_gateway_v1_gs_sync_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomInfo) ProtoMessage() {}

func (x *RoomInfo) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_gateway_v1_gs_sync_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomInfo.ProtoReflect.Descriptor instead.
func (*RoomInfo) Descriptor() ([]byte, []int) {
	return file_thetan_gateway_v1_gs_sync_proto_rawDescGZIP(), []int{6}
}

func (x *RoomInfo) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *RoomInfo) GetRoomNumber() int32 {
	if x != nil {
		return x.RoomNumber
	}
	return 0
}

func (x *RoomInfo) GetTownId() string {
	if x != nil {
		return x.TownId
	}
	return ""
}

func (x *RoomInfo) GetConnectedPlayers() []string {
	if x != nil {
		return x.ConnectedPlayers
	}
	return nil
}

type RoomDestroyedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RoomDestroyedResponse) Reset() {
	*x = RoomDestroyedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_gateway_v1_gs_sync_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomDestroyedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomDestroyedResponse) ProtoMessage() {}

func (x *RoomDestroyedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_gateway_v1_gs_sync_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomDestroyedResponse.ProtoReflect.Descriptor instead.
func (*RoomDestroyedResponse) Descriptor() ([]byte, []int) {
	return file_thetan_gateway_v1_gs_sync_proto_rawDescGZIP(), []int{7}
}

type ImmortalRoomAllocationResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//fishnet ip
	ServerIp string `protobuf:"bytes,1,opt,name=server_ip,json=serverIp,proto3" json:"server_ip,omitempty"`
	//matchid
	RoomId string `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *ImmortalRoomAllocationResp) Reset() {
	*x = ImmortalRoomAllocationResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_gateway_v1_gs_sync_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImmortalRoomAllocationResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImmortalRoomAllocationResp) ProtoMessage() {}

func (x *ImmortalRoomAllocationResp) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_gateway_v1_gs_sync_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImmortalRoomAllocationResp.ProtoReflect.Descriptor instead.
func (*ImmortalRoomAllocationResp) Descriptor() ([]byte, []int) {
	return file_thetan_gateway_v1_gs_sync_proto_rawDescGZIP(), []int{8}
}

func (x *ImmortalRoomAllocationResp) GetServerIp() string {
	if x != nil {
		return x.ServerIp
	}
	return ""
}

func (x *ImmortalRoomAllocationResp) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

// nên đặt bên proto của match-director nhưng để đây để tránh import cycle
type ImmortalMatchFoundForMultiplayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomAlloc *ImmortalRoomAllocationResp `protobuf:"bytes,1,opt,name=roomAlloc,proto3" json:"roomAlloc,omitempty"`
	PartyIDs  []string                    `protobuf:"bytes,2,rep,name=partyIDs,proto3" json:"partyIDs,omitempty"`
}

func (x *ImmortalMatchFoundForMultiplayer) Reset() {
	*x = ImmortalMatchFoundForMultiplayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_gateway_v1_gs_sync_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImmortalMatchFoundForMultiplayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImmortalMatchFoundForMultiplayer) ProtoMessage() {}

func (x *ImmortalMatchFoundForMultiplayer) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_gateway_v1_gs_sync_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImmortalMatchFoundForMultiplayer.ProtoReflect.Descriptor instead.
func (*ImmortalMatchFoundForMultiplayer) Descriptor() ([]byte, []int) {
	return file_thetan_gateway_v1_gs_sync_proto_rawDescGZIP(), []int{9}
}

func (x *ImmortalMatchFoundForMultiplayer) GetRoomAlloc() *ImmortalRoomAllocationResp {
	if x != nil {
		return x.RoomAlloc
	}
	return nil
}

func (x *ImmortalMatchFoundForMultiplayer) GetPartyIDs() []string {
	if x != nil {
		return x.PartyIDs
	}
	return nil
}

var File_thetan_gateway_v1_gs_sync_proto protoreflect.FileDescriptor

var file_thetan_gateway_v1_gs_sync_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x73, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x76, 0x31, 0x22, 0xe4, 0x01, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28,
	0x0a, 0x10, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x22, 0xbc, 0x01, 0x0a, 0x16,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x74, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61,
	0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x28, 0x0a, 0x10, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x67, 0x61, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xbf, 0x01, 0x0a, 0x19, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x74, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x61, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x67, 0x61,
	0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x93, 0x01, 0x0a,
	0x14, 0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61,
	0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x28, 0x0a, 0x10, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x67, 0x61, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x22, 0x28, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x16, 0x0a, 0x14,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8a, 0x01, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x74,
	0x6f, 0x77, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f,
	0x77, 0x6e, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x22, 0x17, 0x0a, 0x15, 0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x52, 0x0a, 0x1a, 0x49, 0x6d,
	0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x8b,
	0x01, 0x0a, 0x20, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x46, 0x6f, 0x75, 0x6e, 0x64, 0x46, 0x6f, 0x72, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x41, 0x6c, 0x6c, 0x6f, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x6d, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x41, 0x6c, 0x6c, 0x6f, 0x63,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x44, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x44, 0x73, 0x2a, 0x2e, 0x0a, 0x08,
	0x47, 0x61, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x69, 0x76, 0x61,
	0x6c, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x10,
	0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x72, 0x65, 0x6e, 0x61, 0x10, 0x02, 0x32, 0x98, 0x03, 0x0a,
	0x0d, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x49,
	0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x0f, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x29, 0x2e, 0x74,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x6d, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x2c, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61,
	0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x64, 0x0a, 0x0d, 0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79,
	0x65, 0x64, 0x12, 0x27, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x65, 0x73, 0x74, 0x72,
	0x6f, 0x79, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xc1, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e,
	0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76,
	0x31, 0x42, 0x0b, 0x47, 0x73, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x35, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2d, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x5f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x47, 0x58, 0xaa, 0x02, 0x11,
	0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x11, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1d, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x3a, 0x3a,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_thetan_gateway_v1_gs_sync_proto_rawDescOnce sync.Once
	file_thetan_gateway_v1_gs_sync_proto_rawDescData = file_thetan_gateway_v1_gs_sync_proto_rawDesc
)

func file_thetan_gateway_v1_gs_sync_proto_rawDescGZIP() []byte {
	file_thetan_gateway_v1_gs_sync_proto_rawDescOnce.Do(func() {
		file_thetan_gateway_v1_gs_sync_proto_rawDescData = protoimpl.X.CompressGZIP(file_thetan_gateway_v1_gs_sync_proto_rawDescData)
	})
	return file_thetan_gateway_v1_gs_sync_proto_rawDescData
}

var file_thetan_gateway_v1_gs_sync_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_thetan_gateway_v1_gs_sync_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_thetan_gateway_v1_gs_sync_proto_goTypes = []interface{}{
	(GameName)(0),                            // 0: thetan.gateway.v1.GameName
	(*PingRequest)(nil),                      // 1: thetan.gateway.v1.PingRequest
	(*PlayerConnectedRequest)(nil),           // 2: thetan.gateway.v1.PlayerConnectedRequest
	(*PlayerDisconnectedRequest)(nil),        // 3: thetan.gateway.v1.PlayerDisconnectedRequest
	(*RoomDestroyedRequest)(nil),             // 4: thetan.gateway.v1.RoomDestroyedRequest
	(*PingResponse)(nil),                     // 5: thetan.gateway.v1.PingResponse
	(*PlayerStatusResponse)(nil),             // 6: thetan.gateway.v1.PlayerStatusResponse
	(*RoomInfo)(nil),                         // 7: thetan.gateway.v1.RoomInfo
	(*RoomDestroyedResponse)(nil),            // 8: thetan.gateway.v1.RoomDestroyedResponse
	(*ImmortalRoomAllocationResp)(nil),       // 9: thetan.gateway.v1.ImmortalRoomAllocationResp
	(*ImmortalMatchFoundForMultiplayer)(nil), // 10: thetan.gateway.v1.ImmortalMatchFoundForMultiplayer
}
var file_thetan_gateway_v1_gs_sync_proto_depIdxs = []int32{
	0,  // 0: thetan.gateway.v1.PingRequest.game_name:type_name -> thetan.gateway.v1.GameName
	7,  // 1: thetan.gateway.v1.PingRequest.rooms:type_name -> thetan.gateway.v1.RoomInfo
	0,  // 2: thetan.gateway.v1.PlayerConnectedRequest.game_name:type_name -> thetan.gateway.v1.GameName
	0,  // 3: thetan.gateway.v1.PlayerDisconnectedRequest.game_name:type_name -> thetan.gateway.v1.GameName
	0,  // 4: thetan.gateway.v1.RoomDestroyedRequest.game_name:type_name -> thetan.gateway.v1.GameName
	9,  // 5: thetan.gateway.v1.ImmortalMatchFoundForMultiplayer.roomAlloc:type_name -> thetan.gateway.v1.ImmortalRoomAllocationResp
	1,  // 6: thetan.gateway.v1.ThetanGateway.Ping:input_type -> thetan.gateway.v1.PingRequest
	2,  // 7: thetan.gateway.v1.ThetanGateway.PlayerConnected:input_type -> thetan.gateway.v1.PlayerConnectedRequest
	3,  // 8: thetan.gateway.v1.ThetanGateway.PlayerDisconnected:input_type -> thetan.gateway.v1.PlayerDisconnectedRequest
	4,  // 9: thetan.gateway.v1.ThetanGateway.RoomDestroyed:input_type -> thetan.gateway.v1.RoomDestroyedRequest
	5,  // 10: thetan.gateway.v1.ThetanGateway.Ping:output_type -> thetan.gateway.v1.PingResponse
	6,  // 11: thetan.gateway.v1.ThetanGateway.PlayerConnected:output_type -> thetan.gateway.v1.PlayerStatusResponse
	6,  // 12: thetan.gateway.v1.ThetanGateway.PlayerDisconnected:output_type -> thetan.gateway.v1.PlayerStatusResponse
	8,  // 13: thetan.gateway.v1.ThetanGateway.RoomDestroyed:output_type -> thetan.gateway.v1.RoomDestroyedResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_thetan_gateway_v1_gs_sync_proto_init() }
func file_thetan_gateway_v1_gs_sync_proto_init() {
	if File_thetan_gateway_v1_gs_sync_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thetan_gateway_v1_gs_sync_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_thetan_gateway_v1_gs_sync_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerConnectedRequest); i {
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
		file_thetan_gateway_v1_gs_sync_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerDisconnectedRequest); i {
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
		file_thetan_gateway_v1_gs_sync_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomDestroyedRequest); i {
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
		file_thetan_gateway_v1_gs_sync_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_thetan_gateway_v1_gs_sync_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerStatusResponse); i {
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
		file_thetan_gateway_v1_gs_sync_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomInfo); i {
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
		file_thetan_gateway_v1_gs_sync_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomDestroyedResponse); i {
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
		file_thetan_gateway_v1_gs_sync_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImmortalRoomAllocationResp); i {
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
		file_thetan_gateway_v1_gs_sync_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImmortalMatchFoundForMultiplayer); i {
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
			RawDescriptor: file_thetan_gateway_v1_gs_sync_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thetan_gateway_v1_gs_sync_proto_goTypes,
		DependencyIndexes: file_thetan_gateway_v1_gs_sync_proto_depIdxs,
		EnumInfos:         file_thetan_gateway_v1_gs_sync_proto_enumTypes,
		MessageInfos:      file_thetan_gateway_v1_gs_sync_proto_msgTypes,
	}.Build()
	File_thetan_gateway_v1_gs_sync_proto = out.File
	file_thetan_gateway_v1_gs_sync_proto_rawDesc = nil
	file_thetan_gateway_v1_gs_sync_proto_goTypes = nil
	file_thetan_gateway_v1_gs_sync_proto_depIdxs = nil
}
