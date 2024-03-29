// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: thetan/world/service_adapter.proto

package thetan_world_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AvailableItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*AvailableItem_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *AvailableItem) Reset() {
	*x = AvailableItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_world_service_adapter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvailableItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailableItem) ProtoMessage() {}

func (x *AvailableItem) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_world_service_adapter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailableItem.ProtoReflect.Descriptor instead.
func (*AvailableItem) Descriptor() ([]byte, []int) {
	return file_thetan_world_service_adapter_proto_rawDescGZIP(), []int{0}
}

func (x *AvailableItem) GetItems() []*AvailableItem_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type SimpleItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind   int32   `protobuf:"varint,1,opt,name=Kind,proto3" json:"Kind,omitempty"`
	Type   int32   `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	Amount float32 `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *SimpleItem) Reset() {
	*x = SimpleItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_world_service_adapter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleItem) ProtoMessage() {}

func (x *SimpleItem) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_world_service_adapter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleItem.ProtoReflect.Descriptor instead.
func (*SimpleItem) Descriptor() ([]byte, []int) {
	return file_thetan_world_service_adapter_proto_rawDescGZIP(), []int{1}
}

func (x *SimpleItem) GetKind() int32 {
	if x != nil {
		return x.Kind
	}
	return 0
}

func (x *SimpleItem) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *SimpleItem) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type CheckItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data bool `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CheckItemResponse) Reset() {
	*x = CheckItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_world_service_adapter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckItemResponse) ProtoMessage() {}

func (x *CheckItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_world_service_adapter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckItemResponse.ProtoReflect.Descriptor instead.
func (*CheckItemResponse) Descriptor() ([]byte, []int) {
	return file_thetan_world_service_adapter_proto_rawDescGZIP(), []int{2}
}

func (x *CheckItemResponse) GetData() bool {
	if x != nil {
		return x.Data
	}
	return false
}

type SendItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string        `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Items  []*SimpleItem `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (x *SendItemsRequest) Reset() {
	*x = SendItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_world_service_adapter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendItemsRequest) ProtoMessage() {}

func (x *SendItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_world_service_adapter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendItemsRequest.ProtoReflect.Descriptor instead.
func (*SendItemsRequest) Descriptor() ([]byte, []int) {
	return file_thetan_world_service_adapter_proto_rawDescGZIP(), []int{3}
}

func (x *SendItemsRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *SendItemsRequest) GetItems() []*SimpleItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type AvailableItem_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind     int32  `protobuf:"varint,1,opt,name=Kind,proto3" json:"Kind,omitempty"`
	Type     int32  `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	ItemType string `protobuf:"bytes,3,opt,name=ItemType,proto3" json:"ItemType,omitempty"`
	ItenName string `protobuf:"bytes,4,opt,name=ItenName,proto3" json:"ItenName,omitempty"`
}

func (x *AvailableItem_Item) Reset() {
	*x = AvailableItem_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_world_service_adapter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvailableItem_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailableItem_Item) ProtoMessage() {}

func (x *AvailableItem_Item) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_world_service_adapter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailableItem_Item.ProtoReflect.Descriptor instead.
func (*AvailableItem_Item) Descriptor() ([]byte, []int) {
	return file_thetan_world_service_adapter_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AvailableItem_Item) GetKind() int32 {
	if x != nil {
		return x.Kind
	}
	return 0
}

func (x *AvailableItem_Item) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *AvailableItem_Item) GetItemType() string {
	if x != nil {
		return x.ItemType
	}
	return ""
}

func (x *AvailableItem_Item) GetItenName() string {
	if x != nil {
		return x.ItenName
	}
	return ""
}

var File_thetan_world_service_adapter_proto protoreflect.FileDescriptor

var file_thetan_world_service_adapter_proto_rawDesc = []byte{
	0x0a, 0x22, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x0d, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x39, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a,
	0x66, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49,
	0x74, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49,
	0x74, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4c, 0x0a, 0x0a, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x27, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5d,
	0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x31, 0x0a, 0x05, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xff, 0x01,
	0x0a, 0x19, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x41, 0x64, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61,
	0x6e, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x4e, 0x0a, 0x0b, 0x49, 0x73, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1b, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x1a, 0x22, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x21, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42,
	0xbb, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31,
	0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2d, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76,
	0x31, 0x3b, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x54, 0x57, 0x58, 0xaa, 0x02, 0x0f, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x54, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x5c, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x54, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x5c, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x54, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x3a, 0x3a, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thetan_world_service_adapter_proto_rawDescOnce sync.Once
	file_thetan_world_service_adapter_proto_rawDescData = file_thetan_world_service_adapter_proto_rawDesc
)

func file_thetan_world_service_adapter_proto_rawDescGZIP() []byte {
	file_thetan_world_service_adapter_proto_rawDescOnce.Do(func() {
		file_thetan_world_service_adapter_proto_rawDescData = protoimpl.X.CompressGZIP(file_thetan_world_service_adapter_proto_rawDescData)
	})
	return file_thetan_world_service_adapter_proto_rawDescData
}

var file_thetan_world_service_adapter_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_thetan_world_service_adapter_proto_goTypes = []interface{}{
	(*AvailableItem)(nil),      // 0: thetan.world.v1.AvailableItem
	(*SimpleItem)(nil),         // 1: thetan.world.v1.SimpleItem
	(*CheckItemResponse)(nil),  // 2: thetan.world.v1.CheckItemResponse
	(*SendItemsRequest)(nil),   // 3: thetan.world.v1.SendItemsRequest
	(*AvailableItem_Item)(nil), // 4: thetan.world.v1.AvailableItem.Item
	(*emptypb.Empty)(nil),      // 5: google.protobuf.Empty
}
var file_thetan_world_service_adapter_proto_depIdxs = []int32{
	4, // 0: thetan.world.v1.AvailableItem.items:type_name -> thetan.world.v1.AvailableItem.Item
	1, // 1: thetan.world.v1.SendItemsRequest.Items:type_name -> thetan.world.v1.SimpleItem
	5, // 2: thetan.world.v1.ThetanWorldAdapterService.GetAvailableItems:input_type -> google.protobuf.Empty
	1, // 3: thetan.world.v1.ThetanWorldAdapterService.IsValidItem:input_type -> thetan.world.v1.SimpleItem
	3, // 4: thetan.world.v1.ThetanWorldAdapterService.SendItem:input_type -> thetan.world.v1.SendItemsRequest
	0, // 5: thetan.world.v1.ThetanWorldAdapterService.GetAvailableItems:output_type -> thetan.world.v1.AvailableItem
	2, // 6: thetan.world.v1.ThetanWorldAdapterService.IsValidItem:output_type -> thetan.world.v1.CheckItemResponse
	5, // 7: thetan.world.v1.ThetanWorldAdapterService.SendItem:output_type -> google.protobuf.Empty
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_thetan_world_service_adapter_proto_init() }
func file_thetan_world_service_adapter_proto_init() {
	if File_thetan_world_service_adapter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thetan_world_service_adapter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvailableItem); i {
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
		file_thetan_world_service_adapter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleItem); i {
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
		file_thetan_world_service_adapter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckItemResponse); i {
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
		file_thetan_world_service_adapter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendItemsRequest); i {
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
		file_thetan_world_service_adapter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvailableItem_Item); i {
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
			RawDescriptor: file_thetan_world_service_adapter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thetan_world_service_adapter_proto_goTypes,
		DependencyIndexes: file_thetan_world_service_adapter_proto_depIdxs,
		MessageInfos:      file_thetan_world_service_adapter_proto_msgTypes,
	}.Build()
	File_thetan_world_service_adapter_proto = out.File
	file_thetan_world_service_adapter_proto_rawDesc = nil
	file_thetan_world_service_adapter_proto_goTypes = nil
	file_thetan_world_service_adapter_proto_depIdxs = nil
}
