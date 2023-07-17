// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: thetan/immortal/v1/immortal.proto

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

type SearchPlayerInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerIDs []string `protobuf:"bytes,1,rep,name=playerIDs,proto3" json:"playerIDs,omitempty"`
}

func (x *SearchPlayerInfoRequest) Reset() {
	*x = SearchPlayerInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_immortal_v1_immortal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchPlayerInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPlayerInfoRequest) ProtoMessage() {}

func (x *SearchPlayerInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_immortal_v1_immortal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPlayerInfoRequest.ProtoReflect.Descriptor instead.
func (*SearchPlayerInfoRequest) Descriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_proto_rawDescGZIP(), []int{0}
}

func (x *SearchPlayerInfoRequest) GetPlayerIDs() []string {
	if x != nil {
		return x.PlayerIDs
	}
	return nil
}

type SearchPlayerInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*ImmortalPlayerInfoMatchProto `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *SearchPlayerInfoResponse) Reset() {
	*x = SearchPlayerInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_immortal_v1_immortal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchPlayerInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPlayerInfoResponse) ProtoMessage() {}

func (x *SearchPlayerInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_immortal_v1_immortal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPlayerInfoResponse.ProtoReflect.Descriptor instead.
func (*SearchPlayerInfoResponse) Descriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_proto_rawDescGZIP(), []int{1}
}

func (x *SearchPlayerInfoResponse) GetPlayers() []*ImmortalPlayerInfoMatchProto {
	if x != nil {
		return x.Players
	}
	return nil
}

type GetUserProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *GetUserProfileRequest) Reset() {
	*x = GetUserProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_immortal_v1_immortal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserProfileRequest) ProtoMessage() {}

func (x *GetUserProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_immortal_v1_immortal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserProfileRequest.ProtoReflect.Descriptor instead.
func (*GetUserProfileRequest) Descriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserProfileRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type GetUserProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID      string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	UserName    string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Country     string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	AvatarID    int32  `protobuf:"varint,4,opt,name=avatarID,proto3" json:"avatarID,omitempty"`
	FrameID     int32  `protobuf:"varint,5,opt,name=frameID,proto3" json:"frameID,omitempty"`
	NameColorID int32  `protobuf:"varint,6,opt,name=nameColorID,proto3" json:"nameColorID,omitempty"`
	Rank        int32  `protobuf:"varint,7,opt,name=rank,proto3" json:"rank,omitempty"`
	Trophy      int32  `protobuf:"varint,8,opt,name=trophy,proto3" json:"trophy,omitempty"`
}

func (x *GetUserProfileResponse) Reset() {
	*x = GetUserProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_immortal_v1_immortal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserProfileResponse) ProtoMessage() {}

func (x *GetUserProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_immortal_v1_immortal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserProfileResponse.ProtoReflect.Descriptor instead.
func (*GetUserProfileResponse) Descriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserProfileResponse) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetUserProfileResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *GetUserProfileResponse) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *GetUserProfileResponse) GetAvatarID() int32 {
	if x != nil {
		return x.AvatarID
	}
	return 0
}

func (x *GetUserProfileResponse) GetFrameID() int32 {
	if x != nil {
		return x.FrameID
	}
	return 0
}

func (x *GetUserProfileResponse) GetNameColorID() int32 {
	if x != nil {
		return x.NameColorID
	}
	return 0
}

func (x *GetUserProfileResponse) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *GetUserProfileResponse) GetTrophy() int32 {
	if x != nil {
		return x.Trophy
	}
	return 0
}

var File_thetan_immortal_v1_immortal_proto protoreflect.FileDescriptor

var file_thetan_immortal_v1_immortal_proto_rawDesc = []byte{
	0x0a, 0x21, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x27, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f,
	0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x37, 0x0a, 0x17, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x73, 0x22, 0x66, 0x0a, 0x18, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e,
	0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x22, 0x2f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x22, 0xea, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x49,
	0x44, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x49, 0x44,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6c, 0x6f,
	0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x72, 0x6f, 0x70, 0x68,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x72, 0x6f, 0x70, 0x68, 0x79, 0x32,
	0xed, 0x01, 0x0a, 0x0f, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69, 0x6d,
	0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x29, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e,
	0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2a, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0xca, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69,
	0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x49, 0x6d, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x74, 0x68, 0x65,
	0x74, 0x61, 0x6e, 0x2d, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x76,
	0x31, 0x3b, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x5f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x49, 0x58, 0xaa, 0x02, 0x12, 0x54, 0x68, 0x65,
	0x74, 0x61, 0x6e, 0x2e, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x12, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x49, 0x6d,
	0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x3a, 0x3a,
	0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thetan_immortal_v1_immortal_proto_rawDescOnce sync.Once
	file_thetan_immortal_v1_immortal_proto_rawDescData = file_thetan_immortal_v1_immortal_proto_rawDesc
)

func file_thetan_immortal_v1_immortal_proto_rawDescGZIP() []byte {
	file_thetan_immortal_v1_immortal_proto_rawDescOnce.Do(func() {
		file_thetan_immortal_v1_immortal_proto_rawDescData = protoimpl.X.CompressGZIP(file_thetan_immortal_v1_immortal_proto_rawDescData)
	})
	return file_thetan_immortal_v1_immortal_proto_rawDescData
}

var file_thetan_immortal_v1_immortal_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_thetan_immortal_v1_immortal_proto_goTypes = []interface{}{
	(*SearchPlayerInfoRequest)(nil),      // 0: thetan.immortal.v1.SearchPlayerInfoRequest
	(*SearchPlayerInfoResponse)(nil),     // 1: thetan.immortal.v1.SearchPlayerInfoResponse
	(*GetUserProfileRequest)(nil),        // 2: thetan.immortal.v1.GetUserProfileRequest
	(*GetUserProfileResponse)(nil),       // 3: thetan.immortal.v1.GetUserProfileResponse
	(*ImmortalPlayerInfoMatchProto)(nil), // 4: thetan.immortal.v1.ImmortalPlayerInfoMatchProto
}
var file_thetan_immortal_v1_immortal_proto_depIdxs = []int32{
	4, // 0: thetan.immortal.v1.SearchPlayerInfoResponse.players:type_name -> thetan.immortal.v1.ImmortalPlayerInfoMatchProto
	0, // 1: thetan.immortal.v1.ImmortalService.SearchPlayerInfo:input_type -> thetan.immortal.v1.SearchPlayerInfoRequest
	2, // 2: thetan.immortal.v1.ImmortalService.GetUserProfile:input_type -> thetan.immortal.v1.GetUserProfileRequest
	1, // 3: thetan.immortal.v1.ImmortalService.SearchPlayerInfo:output_type -> thetan.immortal.v1.SearchPlayerInfoResponse
	3, // 4: thetan.immortal.v1.ImmortalService.GetUserProfile:output_type -> thetan.immortal.v1.GetUserProfileResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_thetan_immortal_v1_immortal_proto_init() }
func file_thetan_immortal_v1_immortal_proto_init() {
	if File_thetan_immortal_v1_immortal_proto != nil {
		return
	}
	file_thetan_immortal_v1_immortal_match_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_thetan_immortal_v1_immortal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchPlayerInfoRequest); i {
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
		file_thetan_immortal_v1_immortal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchPlayerInfoResponse); i {
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
		file_thetan_immortal_v1_immortal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserProfileRequest); i {
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
		file_thetan_immortal_v1_immortal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserProfileResponse); i {
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
			RawDescriptor: file_thetan_immortal_v1_immortal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thetan_immortal_v1_immortal_proto_goTypes,
		DependencyIndexes: file_thetan_immortal_v1_immortal_proto_depIdxs,
		MessageInfos:      file_thetan_immortal_v1_immortal_proto_msgTypes,
	}.Build()
	File_thetan_immortal_v1_immortal_proto = out.File
	file_thetan_immortal_v1_immortal_proto_rawDesc = nil
	file_thetan_immortal_v1_immortal_proto_goTypes = nil
	file_thetan_immortal_v1_immortal_proto_depIdxs = nil
}
