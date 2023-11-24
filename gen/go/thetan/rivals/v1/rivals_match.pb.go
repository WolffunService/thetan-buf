// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: thetan/rivals/v1/rivals_match.proto

package thetan_rivals_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	v1 "thetan-buf/gen/go/thetan/shared/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RivalCancelTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID string `protobuf:"bytes,1,opt,name=playerID,proto3" json:"playerID,omitempty"`
	PartyID  string `protobuf:"bytes,2,opt,name=partyID,proto3" json:"partyID,omitempty"`
	TicketID string `protobuf:"bytes,3,opt,name=ticketID,proto3" json:"ticketID,omitempty"`
}

func (x *RivalCancelTicketRequest) Reset() {
	*x = RivalCancelTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_rivals_v1_rivals_match_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RivalCancelTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RivalCancelTicketRequest) ProtoMessage() {}

func (x *RivalCancelTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_rivals_v1_rivals_match_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RivalCancelTicketRequest.ProtoReflect.Descriptor instead.
func (*RivalCancelTicketRequest) Descriptor() ([]byte, []int) {
	return file_thetan_rivals_v1_rivals_match_proto_rawDescGZIP(), []int{0}
}

func (x *RivalCancelTicketRequest) GetPlayerID() string {
	if x != nil {
		return x.PlayerID
	}
	return ""
}

func (x *RivalCancelTicketRequest) GetPartyID() string {
	if x != nil {
		return x.PartyID
	}
	return ""
}

func (x *RivalCancelTicketRequest) GetTicketID() string {
	if x != nil {
		return x.TicketID
	}
	return ""
}

type RivalCancelTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RivalCancelTicketResponse) Reset() {
	*x = RivalCancelTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_rivals_v1_rivals_match_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RivalCancelTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RivalCancelTicketResponse) ProtoMessage() {}

func (x *RivalCancelTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_rivals_v1_rivals_match_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RivalCancelTicketResponse.ProtoReflect.Descriptor instead.
func (*RivalCancelTicketResponse) Descriptor() ([]byte, []int) {
	return file_thetan_rivals_v1_rivals_match_proto_rawDescGZIP(), []int{1}
}

type CreateMatchNonMatchingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InGameMode       v1.InGameMode              `protobuf:"varint,1,opt,name=inGameMode,proto3,enum=thetan.shared.v1.InGameMode" json:"inGameMode,omitempty"`
	Players          []*v1.PlayerInfoMatchProto `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
	Regions          []int32                    `protobuf:"varint,3,rep,packed,name=regions,proto3" json:"regions,omitempty"`
	FindMatchVersion int32                      `protobuf:"varint,4,opt,name=findMatchVersion,proto3" json:"findMatchVersion,omitempty"`
}

func (x *CreateMatchNonMatchingRequest) Reset() {
	*x = CreateMatchNonMatchingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_rivals_v1_rivals_match_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMatchNonMatchingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMatchNonMatchingRequest) ProtoMessage() {}

func (x *CreateMatchNonMatchingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_rivals_v1_rivals_match_proto_msgTypes[2]
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
	return file_thetan_rivals_v1_rivals_match_proto_rawDescGZIP(), []int{2}
}

func (x *CreateMatchNonMatchingRequest) GetInGameMode() v1.InGameMode {
	if x != nil {
		return x.InGameMode
	}
	return v1.InGameMode(0)
}

func (x *CreateMatchNonMatchingRequest) GetPlayers() []*v1.PlayerInfoMatchProto {
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

func (x *CreateMatchNonMatchingRequest) GetFindMatchVersion() int32 {
	if x != nil {
		return x.FindMatchVersion
	}
	return 0
}

type CreateMatchNonMatchingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateMatchNonMatchingResponse) Reset() {
	*x = CreateMatchNonMatchingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_rivals_v1_rivals_match_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMatchNonMatchingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMatchNonMatchingResponse) ProtoMessage() {}

func (x *CreateMatchNonMatchingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_rivals_v1_rivals_match_proto_msgTypes[3]
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
	return file_thetan_rivals_v1_rivals_match_proto_rawDescGZIP(), []int{3}
}

var File_thetan_rivals_v1_rivals_match_proto protoreflect.FileDescriptor

var file_thetan_rivals_v1_rivals_match_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69,
	0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x25, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f,
	0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x74, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c,
	0x0a, 0x18, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x22, 0x1b, 0x0a, 0x19,
	0x52, 0x69, 0x76, 0x61, 0x6c, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xe5, 0x01, 0x0a, 0x1d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0a, 0x69,
	0x6e, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1c, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x69,
	0x6e, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x68, 0x65,
	0x74, 0x61, 0x6e, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x66, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x10, 0x66, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x20, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x4e, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xef, 0x02, 0x0a, 0x19, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x69, 0x0a, 0x0c, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x2a, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x6e, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x12, 0x25, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61,
	0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x74, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x00, 0x12, 0x7d, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67,
	0x12, 0x2f, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4e,
	0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x30, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x4e, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xbf, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x10,
	0x52, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x33, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2d, 0x62, 0x75, 0x66, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x72, 0x69, 0x76,
	0x61, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5f, 0x72, 0x69,
	0x76, 0x61, 0x6c, 0x73, 0x5f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x52, 0x58, 0xaa, 0x02, 0x10,
	0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x10, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x73,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x52, 0x69, 0x76,
	0x61, 0x6c, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x12, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x3a, 0x3a, 0x52, 0x69, 0x76,
	0x61, 0x6c, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thetan_rivals_v1_rivals_match_proto_rawDescOnce sync.Once
	file_thetan_rivals_v1_rivals_match_proto_rawDescData = file_thetan_rivals_v1_rivals_match_proto_rawDesc
)

func file_thetan_rivals_v1_rivals_match_proto_rawDescGZIP() []byte {
	file_thetan_rivals_v1_rivals_match_proto_rawDescOnce.Do(func() {
		file_thetan_rivals_v1_rivals_match_proto_rawDescData = protoimpl.X.CompressGZIP(file_thetan_rivals_v1_rivals_match_proto_rawDescData)
	})
	return file_thetan_rivals_v1_rivals_match_proto_rawDescData
}

var file_thetan_rivals_v1_rivals_match_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_thetan_rivals_v1_rivals_match_proto_goTypes = []interface{}{
	(*RivalCancelTicketRequest)(nil),       // 0: thetan.rivals.v1.RivalCancelTicketRequest
	(*RivalCancelTicketResponse)(nil),      // 1: thetan.rivals.v1.RivalCancelTicketResponse
	(*CreateMatchNonMatchingRequest)(nil),  // 2: thetan.rivals.v1.CreateMatchNonMatchingRequest
	(*CreateMatchNonMatchingResponse)(nil), // 3: thetan.rivals.v1.CreateMatchNonMatchingResponse
	(v1.InGameMode)(0),                     // 4: thetan.shared.v1.InGameMode
	(*v1.PlayerInfoMatchProto)(nil),        // 5: thetan.shared.v1.PlayerInfoMatchProto
	(*GetMatchInfoRequest)(nil),            // 6: thetan.rivals.v1.GetMatchInfoRequest
	(*v1.MatchFoundResponseProto)(nil),     // 7: thetan.shared.v1.MatchFoundResponseProto
}
var file_thetan_rivals_v1_rivals_match_proto_depIdxs = []int32{
	4, // 0: thetan.rivals.v1.CreateMatchNonMatchingRequest.inGameMode:type_name -> thetan.shared.v1.InGameMode
	5, // 1: thetan.rivals.v1.CreateMatchNonMatchingRequest.players:type_name -> thetan.shared.v1.PlayerInfoMatchProto
	0, // 2: thetan.rivals.v1.RivalMatchDirectorService.CancelTicket:input_type -> thetan.rivals.v1.RivalCancelTicketRequest
	6, // 3: thetan.rivals.v1.RivalMatchDirectorService.CreateMatchOnboard:input_type -> thetan.rivals.v1.GetMatchInfoRequest
	2, // 4: thetan.rivals.v1.RivalMatchDirectorService.CreateMatchNonMatching:input_type -> thetan.rivals.v1.CreateMatchNonMatchingRequest
	1, // 5: thetan.rivals.v1.RivalMatchDirectorService.CancelTicket:output_type -> thetan.rivals.v1.RivalCancelTicketResponse
	7, // 6: thetan.rivals.v1.RivalMatchDirectorService.CreateMatchOnboard:output_type -> thetan.shared.v1.MatchFoundResponseProto
	3, // 7: thetan.rivals.v1.RivalMatchDirectorService.CreateMatchNonMatching:output_type -> thetan.rivals.v1.CreateMatchNonMatchingResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_thetan_rivals_v1_rivals_match_proto_init() }
func file_thetan_rivals_v1_rivals_match_proto_init() {
	if File_thetan_rivals_v1_rivals_match_proto != nil {
		return
	}
	file_thetan_rivals_v1_service_rivals_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_thetan_rivals_v1_rivals_match_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RivalCancelTicketRequest); i {
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
		file_thetan_rivals_v1_rivals_match_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RivalCancelTicketResponse); i {
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
		file_thetan_rivals_v1_rivals_match_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_thetan_rivals_v1_rivals_match_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_thetan_rivals_v1_rivals_match_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thetan_rivals_v1_rivals_match_proto_goTypes,
		DependencyIndexes: file_thetan_rivals_v1_rivals_match_proto_depIdxs,
		MessageInfos:      file_thetan_rivals_v1_rivals_match_proto_msgTypes,
	}.Build()
	File_thetan_rivals_v1_rivals_match_proto = out.File
	file_thetan_rivals_v1_rivals_match_proto_rawDesc = nil
	file_thetan_rivals_v1_rivals_match_proto_goTypes = nil
	file_thetan_rivals_v1_rivals_match_proto_depIdxs = nil
}
