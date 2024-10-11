// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: thetan/gateway/v1/gs_immortal.proto

package thetan_gateway_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	v1 "thetan-buf/gen/go/thetan/immortal/v1"
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
	return file_thetan_gateway_v1_gs_immortal_proto_enumTypes[0].Descriptor()
}

func (GameName) Type() protoreflect.EnumType {
	return &file_thetan_gateway_v1_gs_immortal_proto_enumTypes[0]
}

func (x GameName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameName.Descriptor instead.
func (GameName) EnumDescriptor() ([]byte, []int) {
	return file_thetan_gateway_v1_gs_immortal_proto_rawDescGZIP(), []int{0}
}

type RoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameServerName string `protobuf:"bytes,1,opt,name=gameServerName,proto3" json:"gameServerName,omitempty"`
}

func (x *RoomRequest) Reset() {
	*x = RoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_gateway_v1_gs_immortal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomRequest) ProtoMessage() {}

func (x *RoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_gateway_v1_gs_immortal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomRequest.ProtoReflect.Descriptor instead.
func (*RoomRequest) Descriptor() ([]byte, []int) {
	return file_thetan_gateway_v1_gs_immortal_proto_rawDescGZIP(), []int{0}
}

func (x *RoomRequest) GetGameServerName() string {
	if x != nil {
		return x.GameServerName
	}
	return ""
}

type ImmortalRoomAllocationResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//fishnet ip+port
	ServerIP   string `protobuf:"bytes,1,opt,name=serverIP,proto3" json:"serverIP,omitempty"`
	ServerPort uint32 `protobuf:"varint,2,opt,name=serverPort,proto3" json:"serverPort,omitempty"`
	//matchid
	RoomID string `protobuf:"bytes,3,opt,name=roomID,proto3" json:"roomID,omitempty"`
}

func (x *ImmortalRoomAllocationResp) Reset() {
	*x = ImmortalRoomAllocationResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_gateway_v1_gs_immortal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImmortalRoomAllocationResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImmortalRoomAllocationResp) ProtoMessage() {}

func (x *ImmortalRoomAllocationResp) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_gateway_v1_gs_immortal_proto_msgTypes[1]
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
	return file_thetan_gateway_v1_gs_immortal_proto_rawDescGZIP(), []int{1}
}

func (x *ImmortalRoomAllocationResp) GetServerIP() string {
	if x != nil {
		return x.ServerIP
	}
	return ""
}

func (x *ImmortalRoomAllocationResp) GetServerPort() uint32 {
	if x != nil {
		return x.ServerPort
	}
	return 0
}

func (x *ImmortalRoomAllocationResp) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

// nên đặt bên proto của match-director nhưng để đây để tránh import cycle
type ImmortalMatchFoundForMultiplayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomAlloc  *ImmortalRoomAllocationResp `protobuf:"bytes,1,opt,name=roomAlloc,proto3" json:"roomAlloc,omitempty"`
	PartyIDs   []string                    `protobuf:"bytes,2,rep,name=partyIDs,proto3" json:"partyIDs,omitempty"`
	MatchFound *v1.MatchFoundResponseProto `protobuf:"bytes,3,opt,name=matchFound,proto3" json:"matchFound,omitempty"`
}

func (x *ImmortalMatchFoundForMultiplayer) Reset() {
	*x = ImmortalMatchFoundForMultiplayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_gateway_v1_gs_immortal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImmortalMatchFoundForMultiplayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImmortalMatchFoundForMultiplayer) ProtoMessage() {}

func (x *ImmortalMatchFoundForMultiplayer) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_gateway_v1_gs_immortal_proto_msgTypes[2]
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
	return file_thetan_gateway_v1_gs_immortal_proto_rawDescGZIP(), []int{2}
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

func (x *ImmortalMatchFoundForMultiplayer) GetMatchFound() *v1.MatchFoundResponseProto {
	if x != nil {
		return x.MatchFound
	}
	return nil
}

var File_thetan_gateway_v1_gs_immortal_proto protoreflect.FileDescriptor

var file_thetan_gateway_v1_gs_immortal_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x73, 0x5f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x69, 0x6d,
	0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35,
	0x0a, 0x0b, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x0e, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x70, 0x0a, 0x1a, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x50, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x50, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x22, 0xd8, 0x01, 0x0a, 0x20, 0x49, 0x6d, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x46, 0x6f,
	0x72, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x09,
	0x72, 0x6f, 0x6f, 0x6d, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x52, 0x6f, 0x6f, 0x6d,
	0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x09,
	0x72, 0x6f, 0x6f, 0x6d, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x49, 0x44, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x49, 0x44, 0x73, 0x12, 0x4b, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x46, 0x6f,
	0x75, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x74, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x2a, 0x2e, 0x0a, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x09,
	0x0a, 0x05, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x6d, 0x6d,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x72, 0x65, 0x6e, 0x61,
	0x10, 0x02, 0x32, 0xab, 0x01, 0x0a, 0x15, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x49, 0x0a, 0x0d,
	0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1e, 0x2e,
	0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0b, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x1e, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x32, 0x8b, 0x01, 0x0a, 0x15, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x12, 0x72, 0x0a, 0x12, 0x41, 0x6c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x2b, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x46, 0x6f, 0x75, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x2e,
	0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x6c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0xc5,
	0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x47, 0x73, 0x49, 0x6d, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x74, 0x68, 0x65,
	0x74, 0x61, 0x6e, 0x2d, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31,
	0x3b, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x47, 0x58, 0xaa, 0x02, 0x11, 0x54, 0x68, 0x65, 0x74, 0x61,
	0x6e, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x54,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x1d, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x13, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x3a, 0x3a, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thetan_gateway_v1_gs_immortal_proto_rawDescOnce sync.Once
	file_thetan_gateway_v1_gs_immortal_proto_rawDescData = file_thetan_gateway_v1_gs_immortal_proto_rawDesc
)

func file_thetan_gateway_v1_gs_immortal_proto_rawDescGZIP() []byte {
	file_thetan_gateway_v1_gs_immortal_proto_rawDescOnce.Do(func() {
		file_thetan_gateway_v1_gs_immortal_proto_rawDescData = protoimpl.X.CompressGZIP(file_thetan_gateway_v1_gs_immortal_proto_rawDescData)
	})
	return file_thetan_gateway_v1_gs_immortal_proto_rawDescData
}

var file_thetan_gateway_v1_gs_immortal_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_thetan_gateway_v1_gs_immortal_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_thetan_gateway_v1_gs_immortal_proto_goTypes = []interface{}{
	(GameName)(0),                            // 0: thetan.gateway.v1.GameName
	(*RoomRequest)(nil),                      // 1: thetan.gateway.v1.RoomRequest
	(*ImmortalRoomAllocationResp)(nil),       // 2: thetan.gateway.v1.ImmortalRoomAllocationResp
	(*ImmortalMatchFoundForMultiplayer)(nil), // 3: thetan.gateway.v1.ImmortalMatchFoundForMultiplayer
	(*v1.MatchFoundResponseProto)(nil),       // 4: thetan.immortal.v1.MatchFoundResponseProto
	(*emptypb.Empty)(nil),                    // 5: google.protobuf.Empty
}
var file_thetan_gateway_v1_gs_immortal_proto_depIdxs = []int32{
	2, // 0: thetan.gateway.v1.ImmortalMatchFoundForMultiplayer.roomAlloc:type_name -> thetan.gateway.v1.ImmortalRoomAllocationResp
	4, // 1: thetan.gateway.v1.ImmortalMatchFoundForMultiplayer.matchFound:type_name -> thetan.immortal.v1.MatchFoundResponseProto
	1, // 2: thetan.gateway.v1.ThetanGatewayTracking.RoomAllocated:input_type -> thetan.gateway.v1.RoomRequest
	1, // 3: thetan.gateway.v1.ThetanGatewayTracking.RoomRelease:input_type -> thetan.gateway.v1.RoomRequest
	4, // 4: thetan.gateway.v1.ThetanGatewayImmortal.AllocateGameServer:input_type -> thetan.immortal.v1.MatchFoundResponseProto
	5, // 5: thetan.gateway.v1.ThetanGatewayTracking.RoomAllocated:output_type -> google.protobuf.Empty
	5, // 6: thetan.gateway.v1.ThetanGatewayTracking.RoomRelease:output_type -> google.protobuf.Empty
	2, // 7: thetan.gateway.v1.ThetanGatewayImmortal.AllocateGameServer:output_type -> thetan.gateway.v1.ImmortalRoomAllocationResp
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_thetan_gateway_v1_gs_immortal_proto_init() }
func file_thetan_gateway_v1_gs_immortal_proto_init() {
	if File_thetan_gateway_v1_gs_immortal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thetan_gateway_v1_gs_immortal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomRequest); i {
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
		file_thetan_gateway_v1_gs_immortal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_thetan_gateway_v1_gs_immortal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_thetan_gateway_v1_gs_immortal_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_thetan_gateway_v1_gs_immortal_proto_goTypes,
		DependencyIndexes: file_thetan_gateway_v1_gs_immortal_proto_depIdxs,
		EnumInfos:         file_thetan_gateway_v1_gs_immortal_proto_enumTypes,
		MessageInfos:      file_thetan_gateway_v1_gs_immortal_proto_msgTypes,
	}.Build()
	File_thetan_gateway_v1_gs_immortal_proto = out.File
	file_thetan_gateway_v1_gs_immortal_proto_rawDesc = nil
	file_thetan_gateway_v1_gs_immortal_proto_goTypes = nil
	file_thetan_gateway_v1_gs_immortal_proto_depIdxs = nil
}
