// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: thetan/immortal/v1/immortal_match.proto

package thetan_immortal_v1

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

type MatchFoundResponseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchID    string                  `protobuf:"bytes,1,opt,name=matchID,proto3" json:"matchID,omitempty"`
	Players    []*PlayerInfoMatchProto `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
	BestRegion int32                   `protobuf:"varint,3,opt,name=bestRegion,proto3" json:"bestRegion,omitempty"`
	GameMode   GameMode                `protobuf:"varint,4,opt,name=gameMode,proto3,enum=thetan.immortal.v1.GameMode" json:"gameMode,omitempty"`
	InGameMode InGameMode              `protobuf:"varint,5,opt,name=inGameMode,proto3,enum=thetan.immortal.v1.InGameMode" json:"inGameMode,omitempty"`
}

func (x *MatchFoundResponseProto) Reset() {
	*x = MatchFoundResponseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_immortal_v1_immortal_match_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchFoundResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchFoundResponseProto) ProtoMessage() {}

func (x *MatchFoundResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_immortal_v1_immortal_match_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchFoundResponseProto.ProtoReflect.Descriptor instead.
func (*MatchFoundResponseProto) Descriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_match_proto_rawDescGZIP(), []int{0}
}

func (x *MatchFoundResponseProto) GetMatchID() string {
	if x != nil {
		return x.MatchID
	}
	return ""
}

func (x *MatchFoundResponseProto) GetPlayers() []*PlayerInfoMatchProto {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *MatchFoundResponseProto) GetBestRegion() int32 {
	if x != nil {
		return x.BestRegion
	}
	return 0
}

func (x *MatchFoundResponseProto) GetGameMode() GameMode {
	if x != nil {
		return x.GameMode
	}
	return GameMode_NONE_MODE
}

func (x *MatchFoundResponseProto) GetInGameMode() InGameMode {
	if x != nil {
		return x.InGameMode
	}
	return InGameMode_SOLO
}

// ImmortalTicketData is used to create ticket
type TicketData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players    []*PlayerInfoMatchProto `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	Regions    []int32                 `protobuf:"varint,2,rep,packed,name=regions,proto3" json:"regions,omitempty"`
	Country    string                  `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"` //using for get bot
	GameMode   GameMode                `protobuf:"varint,4,opt,name=gameMode,proto3,enum=thetan.immortal.v1.GameMode" json:"gameMode,omitempty"`
	InGameMode InGameMode              `protobuf:"varint,5,opt,name=inGameMode,proto3,enum=thetan.immortal.v1.InGameMode" json:"inGameMode,omitempty"`
	PartyID    string                  `protobuf:"bytes,6,opt,name=partyID,proto3" json:"partyID,omitempty"`
}

func (x *TicketData) Reset() {
	*x = TicketData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_immortal_v1_immortal_match_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketData) ProtoMessage() {}

func (x *TicketData) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_immortal_v1_immortal_match_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketData.ProtoReflect.Descriptor instead.
func (*TicketData) Descriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_match_proto_rawDescGZIP(), []int{1}
}

func (x *TicketData) GetPlayers() []*PlayerInfoMatchProto {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *TicketData) GetRegions() []int32 {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *TicketData) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *TicketData) GetGameMode() GameMode {
	if x != nil {
		return x.GameMode
	}
	return GameMode_NONE_MODE
}

func (x *TicketData) GetInGameMode() InGameMode {
	if x != nil {
		return x.InGameMode
	}
	return InGameMode_SOLO
}

func (x *TicketData) GetPartyID() string {
	if x != nil {
		return x.PartyID
	}
	return ""
}

type CancelTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID string `protobuf:"bytes,1,opt,name=playerID,proto3" json:"playerID,omitempty"`
	PartyID  string `protobuf:"bytes,2,opt,name=partyID,proto3" json:"partyID,omitempty"`
	TicketID string `protobuf:"bytes,3,opt,name=ticketID,proto3" json:"ticketID,omitempty"`
}

func (x *CancelTicketRequest) Reset() {
	*x = CancelTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_immortal_v1_immortal_match_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelTicketRequest) ProtoMessage() {}

func (x *CancelTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_immortal_v1_immortal_match_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelTicketRequest.ProtoReflect.Descriptor instead.
func (*CancelTicketRequest) Descriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_match_proto_rawDescGZIP(), []int{2}
}

func (x *CancelTicketRequest) GetPlayerID() string {
	if x != nil {
		return x.PlayerID
	}
	return ""
}

func (x *CancelTicketRequest) GetPartyID() string {
	if x != nil {
		return x.PartyID
	}
	return ""
}

func (x *CancelTicketRequest) GetTicketID() string {
	if x != nil {
		return x.TicketID
	}
	return ""
}

type CancelTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelTicketResponse) Reset() {
	*x = CancelTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_immortal_v1_immortal_match_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelTicketResponse) ProtoMessage() {}

func (x *CancelTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_immortal_v1_immortal_match_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelTicketResponse.ProtoReflect.Descriptor instead.
func (*CancelTicketResponse) Descriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_match_proto_rawDescGZIP(), []int{3}
}

// CancelTicketSuccess is used to push pub/sub when the ticket is canceled successfully
type CancelTicketSuccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID string `protobuf:"bytes,1,opt,name=playerID,proto3" json:"playerID,omitempty"`
	PartyID  string `protobuf:"bytes,2,opt,name=partyID,proto3" json:"partyID,omitempty"`
	TicketID string `protobuf:"bytes,3,opt,name=ticketID,proto3" json:"ticketID,omitempty"`
}

func (x *CancelTicketSuccess) Reset() {
	*x = CancelTicketSuccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_immortal_v1_immortal_match_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelTicketSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelTicketSuccess) ProtoMessage() {}

func (x *CancelTicketSuccess) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_immortal_v1_immortal_match_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelTicketSuccess.ProtoReflect.Descriptor instead.
func (*CancelTicketSuccess) Descriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_match_proto_rawDescGZIP(), []int{4}
}

func (x *CancelTicketSuccess) GetPlayerID() string {
	if x != nil {
		return x.PlayerID
	}
	return ""
}

func (x *CancelTicketSuccess) GetPartyID() string {
	if x != nil {
		return x.PartyID
	}
	return ""
}

func (x *CancelTicketSuccess) GetTicketID() string {
	if x != nil {
		return x.TicketID
	}
	return ""
}

type CreateMatchNonMatchingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players    []*PlayerInfoMatchProto `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	Regions    []int32                 `protobuf:"varint,2,rep,packed,name=regions,proto3" json:"regions,omitempty"`
	Country    string                  `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	GameMode   GameMode                `protobuf:"varint,4,opt,name=gameMode,proto3,enum=thetan.immortal.v1.GameMode" json:"gameMode,omitempty"`
	InGameMode InGameMode              `protobuf:"varint,5,opt,name=inGameMode,proto3,enum=thetan.immortal.v1.InGameMode" json:"inGameMode,omitempty"`
}

func (x *CreateMatchNonMatchingRequest) Reset() {
	*x = CreateMatchNonMatchingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_immortal_v1_immortal_match_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMatchNonMatchingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMatchNonMatchingRequest) ProtoMessage() {}

func (x *CreateMatchNonMatchingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_immortal_v1_immortal_match_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMatchNonMatchingRequest.ProtoReflect.Descriptor instead.
func (*CreateMatchNonMatchingRequest) Descriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_match_proto_rawDescGZIP(), []int{5}
}

func (x *CreateMatchNonMatchingRequest) GetPlayers() []*PlayerInfoMatchProto {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *CreateMatchNonMatchingRequest) GetRegions() []int32 {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *CreateMatchNonMatchingRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *CreateMatchNonMatchingRequest) GetGameMode() GameMode {
	if x != nil {
		return x.GameMode
	}
	return GameMode_NONE_MODE
}

func (x *CreateMatchNonMatchingRequest) GetInGameMode() InGameMode {
	if x != nil {
		return x.InGameMode
	}
	return InGameMode_SOLO
}

type CreateMatchNonMatchingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateMatchNonMatchingResponse) Reset() {
	*x = CreateMatchNonMatchingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_immortal_v1_immortal_match_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMatchNonMatchingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMatchNonMatchingResponse) ProtoMessage() {}

func (x *CreateMatchNonMatchingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_immortal_v1_immortal_match_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMatchNonMatchingResponse.ProtoReflect.Descriptor instead.
func (*CreateMatchNonMatchingResponse) Descriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_match_proto_rawDescGZIP(), []int{6}
}

var File_thetan_immortal_v1_immortal_match_proto protoreflect.FileDescriptor

var file_thetan_immortal_v1_immortal_match_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x68, 0x65, 0x74, 0x61,
	0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x28, 0x74,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f,
	0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x02, 0x0a, 0x17, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x46, 0x6f,
	0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x44, 0x12, 0x42, 0x0a, 0x07, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x74, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x62, 0x65, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x62, 0x65, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x38,
	0x0a, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1c, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x08,
	0x67, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x69, 0x6e, 0x47, 0x61,
	0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x74,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x69, 0x6e,
	0x47, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0x98, 0x02, 0x0a, 0x0a, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x42, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61,
	0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x38, 0x0a, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1c, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52,
	0x08, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x69, 0x6e, 0x47,
	0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e,
	0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x69,
	0x6e, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x49, 0x44, 0x22, 0x67, 0x0a, 0x13, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x22, 0x16, 0x0a, 0x14,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x67, 0x0a, 0x13, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x22, 0x91, 0x02,
	0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x6f, 0x6e,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x42, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x38, 0x0a, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x74, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69,
	0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x47, 0x61, 0x6d,
	0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x22, 0x20, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x4e, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xff, 0x01, 0x0a, 0x14, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x0c,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x27, 0x2e, 0x74,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69,
	0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x81, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x4e, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x31, 0x2e, 0x74,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x6f, 0x6e,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x32, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x4e, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xcf, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31,
	0x42, 0x12, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2d, 0x62,
	0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x2f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x68, 0x65,
	0x74, 0x61, 0x6e, 0x5f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5f, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x54, 0x49, 0x58, 0xaa, 0x02, 0x12, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x49,
	0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x54, 0x68, 0x65,
	0x74, 0x61, 0x6e, 0x5c, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x1e, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x14, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x3a, 0x3a, 0x49, 0x6d, 0x6d, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thetan_immortal_v1_immortal_match_proto_rawDescOnce sync.Once
	file_thetan_immortal_v1_immortal_match_proto_rawDescData = file_thetan_immortal_v1_immortal_match_proto_rawDesc
)

func file_thetan_immortal_v1_immortal_match_proto_rawDescGZIP() []byte {
	file_thetan_immortal_v1_immortal_match_proto_rawDescOnce.Do(func() {
		file_thetan_immortal_v1_immortal_match_proto_rawDescData = protoimpl.X.CompressGZIP(file_thetan_immortal_v1_immortal_match_proto_rawDescData)
	})
	return file_thetan_immortal_v1_immortal_match_proto_rawDescData
}

var file_thetan_immortal_v1_immortal_match_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_thetan_immortal_v1_immortal_match_proto_goTypes = []interface{}{
	(*MatchFoundResponseProto)(nil),        // 0: thetan.immortal.v1.MatchFoundResponseProto
	(*TicketData)(nil),                     // 1: thetan.immortal.v1.TicketData
	(*CancelTicketRequest)(nil),            // 2: thetan.immortal.v1.CancelTicketRequest
	(*CancelTicketResponse)(nil),           // 3: thetan.immortal.v1.CancelTicketResponse
	(*CancelTicketSuccess)(nil),            // 4: thetan.immortal.v1.CancelTicketSuccess
	(*CreateMatchNonMatchingRequest)(nil),  // 5: thetan.immortal.v1.CreateMatchNonMatchingRequest
	(*CreateMatchNonMatchingResponse)(nil), // 6: thetan.immortal.v1.CreateMatchNonMatchingResponse
	(*PlayerInfoMatchProto)(nil),           // 7: thetan.immortal.v1.PlayerInfoMatchProto
	(GameMode)(0),                          // 8: thetan.immortal.v1.GameMode
	(InGameMode)(0),                        // 9: thetan.immortal.v1.InGameMode
}
var file_thetan_immortal_v1_immortal_match_proto_depIdxs = []int32{
	7,  // 0: thetan.immortal.v1.MatchFoundResponseProto.players:type_name -> thetan.immortal.v1.PlayerInfoMatchProto
	8,  // 1: thetan.immortal.v1.MatchFoundResponseProto.gameMode:type_name -> thetan.immortal.v1.GameMode
	9,  // 2: thetan.immortal.v1.MatchFoundResponseProto.inGameMode:type_name -> thetan.immortal.v1.InGameMode
	7,  // 3: thetan.immortal.v1.TicketData.players:type_name -> thetan.immortal.v1.PlayerInfoMatchProto
	8,  // 4: thetan.immortal.v1.TicketData.gameMode:type_name -> thetan.immortal.v1.GameMode
	9,  // 5: thetan.immortal.v1.TicketData.inGameMode:type_name -> thetan.immortal.v1.InGameMode
	7,  // 6: thetan.immortal.v1.CreateMatchNonMatchingRequest.players:type_name -> thetan.immortal.v1.PlayerInfoMatchProto
	8,  // 7: thetan.immortal.v1.CreateMatchNonMatchingRequest.gameMode:type_name -> thetan.immortal.v1.GameMode
	9,  // 8: thetan.immortal.v1.CreateMatchNonMatchingRequest.inGameMode:type_name -> thetan.immortal.v1.InGameMode
	2,  // 9: thetan.immortal.v1.MatchDirectorService.CancelTicket:input_type -> thetan.immortal.v1.CancelTicketRequest
	5,  // 10: thetan.immortal.v1.MatchDirectorService.CreateMatchNonMatching:input_type -> thetan.immortal.v1.CreateMatchNonMatchingRequest
	3,  // 11: thetan.immortal.v1.MatchDirectorService.CancelTicket:output_type -> thetan.immortal.v1.CancelTicketResponse
	6,  // 12: thetan.immortal.v1.MatchDirectorService.CreateMatchNonMatching:output_type -> thetan.immortal.v1.CreateMatchNonMatchingResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_thetan_immortal_v1_immortal_match_proto_init() }
func file_thetan_immortal_v1_immortal_match_proto_init() {
	if File_thetan_immortal_v1_immortal_match_proto != nil {
		return
	}
	file_thetan_immortal_v1_immortal_shared_proto_init()
	file_thetan_immortal_v1_immortal_game_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_thetan_immortal_v1_immortal_match_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchFoundResponseProto); i {
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
		file_thetan_immortal_v1_immortal_match_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketData); i {
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
		file_thetan_immortal_v1_immortal_match_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelTicketRequest); i {
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
		file_thetan_immortal_v1_immortal_match_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelTicketResponse); i {
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
		file_thetan_immortal_v1_immortal_match_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelTicketSuccess); i {
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
		file_thetan_immortal_v1_immortal_match_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMatchNonMatchingRequest); i {
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
		file_thetan_immortal_v1_immortal_match_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMatchNonMatchingResponse); i {
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
			RawDescriptor: file_thetan_immortal_v1_immortal_match_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thetan_immortal_v1_immortal_match_proto_goTypes,
		DependencyIndexes: file_thetan_immortal_v1_immortal_match_proto_depIdxs,
		MessageInfos:      file_thetan_immortal_v1_immortal_match_proto_msgTypes,
	}.Build()
	File_thetan_immortal_v1_immortal_match_proto = out.File
	file_thetan_immortal_v1_immortal_match_proto_rawDesc = nil
	file_thetan_immortal_v1_immortal_match_proto_goTypes = nil
	file_thetan_immortal_v1_immortal_match_proto_depIdxs = nil
}
