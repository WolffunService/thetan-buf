// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: thetan/immortal/v1/immortal_player.proto

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

type SearchPlayersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Countries []string `protobuf:"bytes,1,rep,name=countries,proto3" json:"countries,omitempty"`
	Amount    int32    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *SearchPlayersRequest) Reset() {
	*x = SearchPlayersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_immortal_v1_immortal_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchPlayersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPlayersRequest) ProtoMessage() {}

func (x *SearchPlayersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_immortal_v1_immortal_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPlayersRequest.ProtoReflect.Descriptor instead.
func (*SearchPlayersRequest) Descriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_player_proto_rawDescGZIP(), []int{0}
}

func (x *SearchPlayersRequest) GetCountries() []string {
	if x != nil {
		return x.Countries
	}
	return nil
}

func (x *SearchPlayersRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type SearchPlayersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*Player `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *SearchPlayersResponse) Reset() {
	*x = SearchPlayersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_immortal_v1_immortal_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchPlayersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPlayersResponse) ProtoMessage() {}

func (x *SearchPlayersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_immortal_v1_immortal_player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPlayersResponse.ProtoReflect.Descriptor instead.
func (*SearchPlayersResponse) Descriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_player_proto_rawDescGZIP(), []int{1}
}

func (x *SearchPlayersResponse) GetPlayers() []*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Country string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_immortal_v1_immortal_player_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_immortal_v1_immortal_player_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_player_proto_rawDescGZIP(), []int{2}
}

func (x *Player) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Player) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type GetBotNamesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount    int32    `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Countries []string `protobuf:"bytes,2,rep,name=countries,proto3" json:"countries,omitempty"`
}

func (x *GetBotNamesRequest) Reset() {
	*x = GetBotNamesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_immortal_v1_immortal_player_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBotNamesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBotNamesRequest) ProtoMessage() {}

func (x *GetBotNamesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_immortal_v1_immortal_player_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBotNamesRequest.ProtoReflect.Descriptor instead.
func (*GetBotNamesRequest) Descriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_player_proto_rawDescGZIP(), []int{3}
}

func (x *GetBotNamesRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *GetBotNamesRequest) GetCountries() []string {
	if x != nil {
		return x.Countries
	}
	return nil
}

type GetBotNamesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BotNames []string `protobuf:"bytes,1,rep,name=botNames,proto3" json:"botNames,omitempty"`
}

func (x *GetBotNamesResponse) Reset() {
	*x = GetBotNamesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_immortal_v1_immortal_player_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBotNamesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBotNamesResponse) ProtoMessage() {}

func (x *GetBotNamesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_immortal_v1_immortal_player_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBotNamesResponse.ProtoReflect.Descriptor instead.
func (*GetBotNamesResponse) Descriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_player_proto_rawDescGZIP(), []int{4}
}

func (x *GetBotNamesResponse) GetBotNames() []string {
	if x != nil {
		return x.BotNames
	}
	return nil
}

var File_thetan_immortal_v1_immortal_player_proto protoreflect.FileDescriptor

var file_thetan_immortal_v1_immortal_player_proto_rawDesc = []byte{
	0x0a, 0x28, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0x4c,
	0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4d, 0x0a, 0x15,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e,
	0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0x46, 0x0a, 0x06, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x22, 0x4a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22,
	0x31, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6f, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x32, 0xe1, 0x01, 0x0a, 0x15, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x0d,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x28, 0x2e,
	0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x74, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xd0, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x74,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76,
	0x31, 0x42, 0x13, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x2d, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x2f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x74,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x5f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5f, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x54, 0x49, 0x58, 0xaa, 0x02, 0x12, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x2e, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x54,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x1e, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x49, 0x6d, 0x6d, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x14, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x3a, 0x3a, 0x49, 0x6d, 0x6d,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_thetan_immortal_v1_immortal_player_proto_rawDescOnce sync.Once
	file_thetan_immortal_v1_immortal_player_proto_rawDescData = file_thetan_immortal_v1_immortal_player_proto_rawDesc
)

func file_thetan_immortal_v1_immortal_player_proto_rawDescGZIP() []byte {
	file_thetan_immortal_v1_immortal_player_proto_rawDescOnce.Do(func() {
		file_thetan_immortal_v1_immortal_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_thetan_immortal_v1_immortal_player_proto_rawDescData)
	})
	return file_thetan_immortal_v1_immortal_player_proto_rawDescData
}

var file_thetan_immortal_v1_immortal_player_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_thetan_immortal_v1_immortal_player_proto_goTypes = []interface{}{
	(*SearchPlayersRequest)(nil),  // 0: thetan.immortal.v1.SearchPlayersRequest
	(*SearchPlayersResponse)(nil), // 1: thetan.immortal.v1.SearchPlayersResponse
	(*Player)(nil),                // 2: thetan.immortal.v1.Player
	(*GetBotNamesRequest)(nil),    // 3: thetan.immortal.v1.GetBotNamesRequest
	(*GetBotNamesResponse)(nil),   // 4: thetan.immortal.v1.GetBotNamesResponse
}
var file_thetan_immortal_v1_immortal_player_proto_depIdxs = []int32{
	2, // 0: thetan.immortal.v1.SearchPlayersResponse.players:type_name -> thetan.immortal.v1.Player
	0, // 1: thetan.immortal.v1.ImmortalPlayerService.SearchPlayers:input_type -> thetan.immortal.v1.SearchPlayersRequest
	3, // 2: thetan.immortal.v1.ImmortalPlayerService.GetBotNames:input_type -> thetan.immortal.v1.GetBotNamesRequest
	1, // 3: thetan.immortal.v1.ImmortalPlayerService.SearchPlayers:output_type -> thetan.immortal.v1.SearchPlayersResponse
	4, // 4: thetan.immortal.v1.ImmortalPlayerService.GetBotNames:output_type -> thetan.immortal.v1.GetBotNamesResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_thetan_immortal_v1_immortal_player_proto_init() }
func file_thetan_immortal_v1_immortal_player_proto_init() {
	if File_thetan_immortal_v1_immortal_player_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thetan_immortal_v1_immortal_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchPlayersRequest); i {
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
		file_thetan_immortal_v1_immortal_player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchPlayersResponse); i {
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
		file_thetan_immortal_v1_immortal_player_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_thetan_immortal_v1_immortal_player_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBotNamesRequest); i {
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
		file_thetan_immortal_v1_immortal_player_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBotNamesResponse); i {
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
			RawDescriptor: file_thetan_immortal_v1_immortal_player_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thetan_immortal_v1_immortal_player_proto_goTypes,
		DependencyIndexes: file_thetan_immortal_v1_immortal_player_proto_depIdxs,
		MessageInfos:      file_thetan_immortal_v1_immortal_player_proto_msgTypes,
	}.Build()
	File_thetan_immortal_v1_immortal_player_proto = out.File
	file_thetan_immortal_v1_immortal_player_proto_rawDesc = nil
	file_thetan_immortal_v1_immortal_player_proto_goTypes = nil
	file_thetan_immortal_v1_immortal_player_proto_depIdxs = nil
}
