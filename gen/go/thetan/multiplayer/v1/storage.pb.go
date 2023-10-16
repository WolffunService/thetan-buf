// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: thetan/multiplayer/v1/storage.proto

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

type PlayerStorageMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerInfo *PlayerInfoProto `protobuf:"bytes,1,opt,name=playerInfo,proto3" json:"playerInfo,omitempty"`
	Version    int32            `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	MatchInfo  *MatchInfo       `protobuf:"bytes,3,opt,name=matchInfo,proto3" json:"matchInfo,omitempty"`
}

func (x *PlayerStorageMessage) Reset() {
	*x = PlayerStorageMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_multiplayer_v1_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerStorageMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerStorageMessage) ProtoMessage() {}

func (x *PlayerStorageMessage) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_multiplayer_v1_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerStorageMessage.ProtoReflect.Descriptor instead.
func (*PlayerStorageMessage) Descriptor() ([]byte, []int) {
	return file_thetan_multiplayer_v1_storage_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerStorageMessage) GetPlayerInfo() *PlayerInfoProto {
	if x != nil {
		return x.PlayerInfo
	}
	return nil
}

func (x *PlayerStorageMessage) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *PlayerStorageMessage) GetMatchInfo() *MatchInfo {
	if x != nil {
		return x.MatchInfo
	}
	return nil
}

type PartyStorageMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartyInfo              *PartyInfoProto     `protobuf:"bytes,1,opt,name=partyInfo,proto3" json:"partyInfo,omitempty"`
	PlayersReadyForRematch map[string]bool     `protobuf:"bytes,2,rep,name=playersReadyForRematch,proto3" json:"playersReadyForRematch,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	PlayersRegions         map[string]*Regions `protobuf:"bytes,3,rep,name=playersRegions,proto3" json:"playersRegions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PartyStorageMessage) Reset() {
	*x = PartyStorageMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_multiplayer_v1_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyStorageMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyStorageMessage) ProtoMessage() {}

func (x *PartyStorageMessage) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_multiplayer_v1_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyStorageMessage.ProtoReflect.Descriptor instead.
func (*PartyStorageMessage) Descriptor() ([]byte, []int) {
	return file_thetan_multiplayer_v1_storage_proto_rawDescGZIP(), []int{1}
}

func (x *PartyStorageMessage) GetPartyInfo() *PartyInfoProto {
	if x != nil {
		return x.PartyInfo
	}
	return nil
}

func (x *PartyStorageMessage) GetPlayersReadyForRematch() map[string]bool {
	if x != nil {
		return x.PlayersReadyForRematch
	}
	return nil
}

func (x *PartyStorageMessage) GetPlayersRegions() map[string]*Regions {
	if x != nil {
		return x.PlayersRegions
	}
	return nil
}

type Regions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regions []int32 `protobuf:"varint,1,rep,packed,name=regions,proto3" json:"regions,omitempty"`
}

func (x *Regions) Reset() {
	*x = Regions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_multiplayer_v1_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Regions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Regions) ProtoMessage() {}

func (x *Regions) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_multiplayer_v1_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Regions.ProtoReflect.Descriptor instead.
func (*Regions) Descriptor() ([]byte, []int) {
	return file_thetan_multiplayer_v1_storage_proto_rawDescGZIP(), []int{2}
}

func (x *Regions) GetRegions() []int32 {
	if x != nil {
		return x.Regions
	}
	return nil
}

type MatchInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketID string `protobuf:"bytes,1,opt,name=ticketID,proto3" json:"ticketID,omitempty"`
	MatchID  string `protobuf:"bytes,2,opt,name=matchID,proto3" json:"matchID,omitempty"`
}

func (x *MatchInfo) Reset() {
	*x = MatchInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_multiplayer_v1_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchInfo) ProtoMessage() {}

func (x *MatchInfo) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_multiplayer_v1_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchInfo.ProtoReflect.Descriptor instead.
func (*MatchInfo) Descriptor() ([]byte, []int) {
	return file_thetan_multiplayer_v1_storage_proto_rawDescGZIP(), []int{3}
}

func (x *MatchInfo) GetTicketID() string {
	if x != nil {
		return x.TicketID
	}
	return ""
}

func (x *MatchInfo) GetMatchID() string {
	if x != nil {
		return x.MatchID
	}
	return ""
}

var File_thetan_multiplayer_v1_storage_proto protoreflect.FileDescriptor

var file_thetan_multiplayer_v1_storage_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x22, 0x74, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb8, 0x01, 0x0a, 0x14, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x09, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xf0, 0x03, 0x0a, 0x13,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x09, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x7e, 0x0a, 0x16, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x61, 0x64, 0x79, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61,
	0x6e, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x61, 0x64,
	0x79, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x16, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x61, 0x64, 0x79, 0x46, 0x6f,
	0x72, 0x52, 0x65, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x66, 0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3e, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73,
	0x1a, 0x49, 0x0a, 0x1b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x61, 0x64, 0x79,
	0x46, 0x6f, 0x72, 0x52, 0x65, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x61, 0x0a, 0x13, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x23,
	0x0a, 0x07, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x41, 0x0a, 0x09, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x49, 0x44, 0x42, 0xde, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x74,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2d, 0x62, 0x75, 0x66,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x4d, 0x58, 0xaa, 0x02, 0x15, 0x54, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x15, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x21, 0x54, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x5c, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17,
	0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x3a, 0x3a, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thetan_multiplayer_v1_storage_proto_rawDescOnce sync.Once
	file_thetan_multiplayer_v1_storage_proto_rawDescData = file_thetan_multiplayer_v1_storage_proto_rawDesc
)

func file_thetan_multiplayer_v1_storage_proto_rawDescGZIP() []byte {
	file_thetan_multiplayer_v1_storage_proto_rawDescOnce.Do(func() {
		file_thetan_multiplayer_v1_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_thetan_multiplayer_v1_storage_proto_rawDescData)
	})
	return file_thetan_multiplayer_v1_storage_proto_rawDescData
}

var file_thetan_multiplayer_v1_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_thetan_multiplayer_v1_storage_proto_goTypes = []interface{}{
	(*PlayerStorageMessage)(nil), // 0: thetan.multiplayer.v1.PlayerStorageMessage
	(*PartyStorageMessage)(nil),  // 1: thetan.multiplayer.v1.PartyStorageMessage
	(*Regions)(nil),              // 2: thetan.multiplayer.v1.Regions
	(*MatchInfo)(nil),            // 3: thetan.multiplayer.v1.MatchInfo
	nil,                          // 4: thetan.multiplayer.v1.PartyStorageMessage.PlayersReadyForRematchEntry
	nil,                          // 5: thetan.multiplayer.v1.PartyStorageMessage.PlayersRegionsEntry
	(*PlayerInfoProto)(nil),      // 6: thetan.multiplayer.v1.PlayerInfoProto
	(*PartyInfoProto)(nil),       // 7: thetan.multiplayer.v1.PartyInfoProto
}
var file_thetan_multiplayer_v1_storage_proto_depIdxs = []int32{
	6, // 0: thetan.multiplayer.v1.PlayerStorageMessage.playerInfo:type_name -> thetan.multiplayer.v1.PlayerInfoProto
	3, // 1: thetan.multiplayer.v1.PlayerStorageMessage.matchInfo:type_name -> thetan.multiplayer.v1.MatchInfo
	7, // 2: thetan.multiplayer.v1.PartyStorageMessage.partyInfo:type_name -> thetan.multiplayer.v1.PartyInfoProto
	4, // 3: thetan.multiplayer.v1.PartyStorageMessage.playersReadyForRematch:type_name -> thetan.multiplayer.v1.PartyStorageMessage.PlayersReadyForRematchEntry
	5, // 4: thetan.multiplayer.v1.PartyStorageMessage.playersRegions:type_name -> thetan.multiplayer.v1.PartyStorageMessage.PlayersRegionsEntry
	2, // 5: thetan.multiplayer.v1.PartyStorageMessage.PlayersRegionsEntry.value:type_name -> thetan.multiplayer.v1.Regions
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_thetan_multiplayer_v1_storage_proto_init() }
func file_thetan_multiplayer_v1_storage_proto_init() {
	if File_thetan_multiplayer_v1_storage_proto != nil {
		return
	}
	file_thetan_multiplayer_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_thetan_multiplayer_v1_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerStorageMessage); i {
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
		file_thetan_multiplayer_v1_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyStorageMessage); i {
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
		file_thetan_multiplayer_v1_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Regions); i {
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
		file_thetan_multiplayer_v1_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchInfo); i {
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
			RawDescriptor: file_thetan_multiplayer_v1_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_thetan_multiplayer_v1_storage_proto_goTypes,
		DependencyIndexes: file_thetan_multiplayer_v1_storage_proto_depIdxs,
		MessageInfos:      file_thetan_multiplayer_v1_storage_proto_msgTypes,
	}.Build()
	File_thetan_multiplayer_v1_storage_proto = out.File
	file_thetan_multiplayer_v1_storage_proto_rawDesc = nil
	file_thetan_multiplayer_v1_storage_proto_goTypes = nil
	file_thetan_multiplayer_v1_storage_proto_depIdxs = nil
}
