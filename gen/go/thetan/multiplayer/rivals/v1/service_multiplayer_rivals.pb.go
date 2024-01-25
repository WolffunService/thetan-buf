// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: thetan/multiplayer/rivals/v1/service_multiplayer_rivals.proto

package thetan_multiplayer_rivals_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
	v1 "thetan-buf/gen/go/thetan/multiplayer/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetOnlineStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIDs          []string `protobuf:"bytes,1,rep,name=userIDs,proto3" json:"userIDs,omitempty"`
	SkipEntryOnError bool     `protobuf:"varint,2,opt,name=skipEntryOnError,proto3" json:"skipEntryOnError,omitempty"`
}

func (x *GetOnlineStatusRequest) Reset() {
	*x = GetOnlineStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOnlineStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOnlineStatusRequest) ProtoMessage() {}

func (x *GetOnlineStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOnlineStatusRequest.ProtoReflect.Descriptor instead.
func (*GetOnlineStatusRequest) Descriptor() ([]byte, []int) {
	return file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_rawDescGZIP(), []int{0}
}

func (x *GetOnlineStatusRequest) GetUserIDs() []string {
	if x != nil {
		return x.UserIDs
	}
	return nil
}

func (x *GetOnlineStatusRequest) GetSkipEntryOnError() bool {
	if x != nil {
		return x.SkipEntryOnError
	}
	return false
}

type GetOnlineStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OnlineStatuses []*OnlineStatus `protobuf:"bytes,1,rep,name=OnlineStatuses,proto3" json:"OnlineStatuses,omitempty"`
}

func (x *GetOnlineStatusResponse) Reset() {
	*x = GetOnlineStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOnlineStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOnlineStatusResponse) ProtoMessage() {}

func (x *GetOnlineStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOnlineStatusResponse.ProtoReflect.Descriptor instead.
func (*GetOnlineStatusResponse) Descriptor() ([]byte, []int) {
	return file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_rawDescGZIP(), []int{1}
}

func (x *GetOnlineStatusResponse) GetOnlineStatuses() []*OnlineStatus {
	if x != nil {
		return x.OnlineStatuses
	}
	return nil
}

type OnlineStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID     string          `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	LastOnline int64           `protobuf:"varint,2,opt,name=lastOnline,proto3" json:"lastOnline,omitempty"`
	Status     v1.PlayerStatus `protobuf:"varint,3,opt,name=status,proto3,enum=thetan.multiplayer.v1.PlayerStatus" json:"status,omitempty"`
}

func (x *OnlineStatus) Reset() {
	*x = OnlineStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlineStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlineStatus) ProtoMessage() {}

func (x *OnlineStatus) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlineStatus.ProtoReflect.Descriptor instead.
func (*OnlineStatus) Descriptor() ([]byte, []int) {
	return file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_rawDescGZIP(), []int{2}
}

func (x *OnlineStatus) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *OnlineStatus) GetLastOnline() int64 {
	if x != nil {
		return x.LastOnline
	}
	return 0
}

func (x *OnlineStatus) GetStatus() v1.PlayerStatus {
	if x != nil {
		return x.Status
	}
	return v1.PlayerStatus(0)
}

type NotifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIDs []string   `protobuf:"bytes,1,rep,name=userIDs,proto3" json:"userIDs,omitempty"`
	Event   string     `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	MsgBody *anypb.Any `protobuf:"bytes,3,opt,name=msgBody,proto3" json:"msgBody,omitempty"`
}

func (x *NotifyRequest) Reset() {
	*x = NotifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyRequest) ProtoMessage() {}

func (x *NotifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyRequest.ProtoReflect.Descriptor instead.
func (*NotifyRequest) Descriptor() ([]byte, []int) {
	return file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_rawDescGZIP(), []int{3}
}

func (x *NotifyRequest) GetUserIDs() []string {
	if x != nil {
		return x.UserIDs
	}
	return nil
}

func (x *NotifyRequest) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *NotifyRequest) GetMsgBody() *anypb.Any {
	if x != nil {
		return x.MsgBody
	}
	return nil
}

type NotifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyResponse) Reset() {
	*x = NotifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyResponse) ProtoMessage() {}

func (x *NotifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyResponse.ProtoReflect.Descriptor instead.
func (*NotifyResponse) Descriptor() ([]byte, []int) {
	return file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_rawDescGZIP(), []int{4}
}

type GetTownOnlineUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TownID string `protobuf:"bytes,1,opt,name=townID,proto3" json:"townID,omitempty"`
}

func (x *GetTownOnlineUsersRequest) Reset() {
	*x = GetTownOnlineUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTownOnlineUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTownOnlineUsersRequest) ProtoMessage() {}

func (x *GetTownOnlineUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTownOnlineUsersRequest.ProtoReflect.Descriptor instead.
func (*GetTownOnlineUsersRequest) Descriptor() ([]byte, []int) {
	return file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_rawDescGZIP(), []int{5}
}

func (x *GetTownOnlineUsersRequest) GetTownID() string {
	if x != nil {
		return x.TownID
	}
	return ""
}

type GetTownOnlineUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIDs []string `protobuf:"bytes,1,rep,name=userIDs,proto3" json:"userIDs,omitempty"`
}

func (x *GetTownOnlineUsersResponse) Reset() {
	*x = GetTownOnlineUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTownOnlineUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTownOnlineUsersResponse) ProtoMessage() {}

func (x *GetTownOnlineUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTownOnlineUsersResponse.ProtoReflect.Descriptor instead.
func (*GetTownOnlineUsersResponse) Descriptor() ([]byte, []int) {
	return file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_rawDescGZIP(), []int{6}
}

func (x *GetTownOnlineUsersResponse) GetUserIDs() []string {
	if x != nil {
		return x.UserIDs
	}
	return nil
}

var File_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto protoreflect.FileDescriptor

var file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2f, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1c, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73,
	0x12, 0x2a, 0x0a, 0x10, 0x73, 0x6b, 0x69, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4f, 0x6e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x73, 0x6b, 0x69, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x4f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x6d, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0e, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x0c,
	0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x6f, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x42, 0x6f,
	0x64, 0x79, 0x22, 0x10, 0x0a, 0x0e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x77, 0x6e, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x77, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x6f, 0x77, 0x6e, 0x49, 0x44, 0x22, 0x36, 0x0a, 0x1a, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x77, 0x6e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x73, 0x32, 0x90, 0x03, 0x0a, 0x18, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x80,
	0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x34, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61,
	0x6e, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x72, 0x69,
	0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x65, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2b, 0x2e, 0x74, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61,
	0x6e, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x72, 0x69,
	0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x89, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x77, 0x6e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x37, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x77, 0x6e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61,
	0x6e, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x72, 0x69,
	0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x77, 0x6e, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0xa1, 0x02, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x68, 0x65,
	0x74, 0x61, 0x6e, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x1d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x69, 0x76,
	0x61, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x74, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x2d, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x2f, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x68, 0x65, 0x74, 0x61,
	0x6e, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x72, 0x69,
	0x76, 0x61, 0x6c, 0x73, 0x5f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x4d, 0x52, 0xaa, 0x02, 0x1c,
	0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x2e, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1c, 0x54,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x5c, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x28, 0x54, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x5c, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5c, 0x52, 0x69, 0x76, 0x61, 0x6c, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1f, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x3a,
	0x3a, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x3a, 0x3a, 0x52, 0x69,
	0x76, 0x61, 0x6c, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_rawDescOnce sync.Once
	file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_rawDescData = file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_rawDesc
)

func file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_rawDescGZIP() []byte {
	file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_rawDescOnce.Do(func() {
		file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_rawDescData = protoimpl.X.CompressGZIP(file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_rawDescData)
	})
	return file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_rawDescData
}

var file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_goTypes = []interface{}{
	(*GetOnlineStatusRequest)(nil),     // 0: thetan.multiplayer.rivals.v1.GetOnlineStatusRequest
	(*GetOnlineStatusResponse)(nil),    // 1: thetan.multiplayer.rivals.v1.GetOnlineStatusResponse
	(*OnlineStatus)(nil),               // 2: thetan.multiplayer.rivals.v1.OnlineStatus
	(*NotifyRequest)(nil),              // 3: thetan.multiplayer.rivals.v1.NotifyRequest
	(*NotifyResponse)(nil),             // 4: thetan.multiplayer.rivals.v1.NotifyResponse
	(*GetTownOnlineUsersRequest)(nil),  // 5: thetan.multiplayer.rivals.v1.GetTownOnlineUsersRequest
	(*GetTownOnlineUsersResponse)(nil), // 6: thetan.multiplayer.rivals.v1.GetTownOnlineUsersResponse
	(v1.PlayerStatus)(0),               // 7: thetan.multiplayer.v1.PlayerStatus
	(*anypb.Any)(nil),                  // 8: google.protobuf.Any
}
var file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_depIdxs = []int32{
	2, // 0: thetan.multiplayer.rivals.v1.GetOnlineStatusResponse.OnlineStatuses:type_name -> thetan.multiplayer.rivals.v1.OnlineStatus
	7, // 1: thetan.multiplayer.rivals.v1.OnlineStatus.status:type_name -> thetan.multiplayer.v1.PlayerStatus
	8, // 2: thetan.multiplayer.rivals.v1.NotifyRequest.msgBody:type_name -> google.protobuf.Any
	0, // 3: thetan.multiplayer.rivals.v1.RivalsMultiplayerService.GetOnlineStatus:input_type -> thetan.multiplayer.rivals.v1.GetOnlineStatusRequest
	3, // 4: thetan.multiplayer.rivals.v1.RivalsMultiplayerService.Notify:input_type -> thetan.multiplayer.rivals.v1.NotifyRequest
	5, // 5: thetan.multiplayer.rivals.v1.RivalsMultiplayerService.GetTownOnlineUsers:input_type -> thetan.multiplayer.rivals.v1.GetTownOnlineUsersRequest
	1, // 6: thetan.multiplayer.rivals.v1.RivalsMultiplayerService.GetOnlineStatus:output_type -> thetan.multiplayer.rivals.v1.GetOnlineStatusResponse
	4, // 7: thetan.multiplayer.rivals.v1.RivalsMultiplayerService.Notify:output_type -> thetan.multiplayer.rivals.v1.NotifyResponse
	6, // 8: thetan.multiplayer.rivals.v1.RivalsMultiplayerService.GetTownOnlineUsers:output_type -> thetan.multiplayer.rivals.v1.GetTownOnlineUsersResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_init() }
func file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_init() {
	if File_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOnlineStatusRequest); i {
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
		file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOnlineStatusResponse); i {
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
		file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlineStatus); i {
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
		file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyRequest); i {
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
		file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyResponse); i {
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
		file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTownOnlineUsersRequest); i {
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
		file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTownOnlineUsersResponse); i {
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
			RawDescriptor: file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_goTypes,
		DependencyIndexes: file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_depIdxs,
		MessageInfos:      file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_msgTypes,
	}.Build()
	File_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto = out.File
	file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_rawDesc = nil
	file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_goTypes = nil
	file_thetan_multiplayer_rivals_v1_service_multiplayer_rivals_proto_depIdxs = nil
}
