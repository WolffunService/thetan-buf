// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: thetan/world/v1/service_adapter.proto

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

type ItemResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind     int32   `protobuf:"varint,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Type     int32   `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	ItemType string  `protobuf:"bytes,3,opt,name=itemType,proto3" json:"itemType,omitempty"`
	Name     string  `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Amount   float32 `protobuf:"fixed32,5,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ItemResp) Reset() {
	*x = ItemResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_world_v1_service_adapter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemResp) ProtoMessage() {}

func (x *ItemResp) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_world_v1_service_adapter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemResp.ProtoReflect.Descriptor instead.
func (*ItemResp) Descriptor() ([]byte, []int) {
	return file_thetan_world_v1_service_adapter_proto_rawDescGZIP(), []int{0}
}

func (x *ItemResp) GetKind() int32 {
	if x != nil {
		return x.Kind
	}
	return 0
}

func (x *ItemResp) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ItemResp) GetItemType() string {
	if x != nil {
		return x.ItemType
	}
	return ""
}

func (x *ItemResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ItemResp) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type AvailableItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ItemResp `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *AvailableItem) Reset() {
	*x = AvailableItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_world_v1_service_adapter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvailableItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailableItem) ProtoMessage() {}

func (x *AvailableItem) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_world_v1_service_adapter_proto_msgTypes[1]
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
	return file_thetan_world_v1_service_adapter_proto_rawDescGZIP(), []int{1}
}

func (x *AvailableItem) GetItems() []*ItemResp {
	if x != nil {
		return x.Items
	}
	return nil
}

type SimpleItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind   int32   `protobuf:"varint,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Type   int32   `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Amount float32 `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *SimpleItem) Reset() {
	*x = SimpleItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_world_v1_service_adapter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleItem) ProtoMessage() {}

func (x *SimpleItem) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_world_v1_service_adapter_proto_msgTypes[2]
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
	return file_thetan_world_v1_service_adapter_proto_rawDescGZIP(), []int{2}
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

type IsValidItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*SimpleItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *IsValidItemsRequest) Reset() {
	*x = IsValidItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_world_v1_service_adapter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsValidItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsValidItemsRequest) ProtoMessage() {}

func (x *IsValidItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_world_v1_service_adapter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsValidItemsRequest.ProtoReflect.Descriptor instead.
func (*IsValidItemsRequest) Descriptor() ([]byte, []int) {
	return file_thetan_world_v1_service_adapter_proto_rawDescGZIP(), []int{3}
}

func (x *IsValidItemsRequest) GetItems() []*SimpleItem {
	if x != nil {
		return x.Items
	}
	return nil
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
		mi := &file_thetan_world_v1_service_adapter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckItemResponse) ProtoMessage() {}

func (x *CheckItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_world_v1_service_adapter_proto_msgTypes[4]
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
	return file_thetan_world_v1_service_adapter_proto_rawDescGZIP(), []int{4}
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
	Items  []*SimpleItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *SendItemsRequest) Reset() {
	*x = SendItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_world_v1_service_adapter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendItemsRequest) ProtoMessage() {}

func (x *SendItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_world_v1_service_adapter_proto_msgTypes[5]
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
	return file_thetan_world_v1_service_adapter_proto_rawDescGZIP(), []int{5}
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

type SendItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ItemResp `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *SendItemsResponse) Reset() {
	*x = SendItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_world_v1_service_adapter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendItemsResponse) ProtoMessage() {}

func (x *SendItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_world_v1_service_adapter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendItemsResponse.ProtoReflect.Descriptor instead.
func (*SendItemsResponse) Descriptor() ([]byte, []int) {
	return file_thetan_world_v1_service_adapter_proto_rawDescGZIP(), []int{6}
}

func (x *SendItemsResponse) GetItems() []*ItemResp {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_thetan_world_v1_service_adapter_proto protoreflect.FileDescriptor

var file_thetan_world_v1_service_adapter_proto_rawDesc = []byte{
	0x0a, 0x25, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x74, 0x65,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x74, 0x65,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x40, 0x0a, 0x0d, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x4c, 0x0a, 0x0a, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x48, 0x0a, 0x13, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x27, 0x0a, 0x11, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x5d, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x31, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x44, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x96, 0x02, 0x0a, 0x19, 0x54, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x58, 0x0a, 0x0c, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x24, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x74, 0x68, 0x65,
	0x74, 0x61, 0x6e, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52,
	0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x21, 0x2e, 0x74, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0xbb, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61,
	0x6e, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x31, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2d, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x5f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x57, 0x58, 0xaa, 0x02, 0x0f, 0x54, 0x68, 0x65,
	0x74, 0x61, 0x6e, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x54,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x1b, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x54,
	0x68, 0x65, 0x74, 0x61, 0x6e, 0x3a, 0x3a, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thetan_world_v1_service_adapter_proto_rawDescOnce sync.Once
	file_thetan_world_v1_service_adapter_proto_rawDescData = file_thetan_world_v1_service_adapter_proto_rawDesc
)

func file_thetan_world_v1_service_adapter_proto_rawDescGZIP() []byte {
	file_thetan_world_v1_service_adapter_proto_rawDescOnce.Do(func() {
		file_thetan_world_v1_service_adapter_proto_rawDescData = protoimpl.X.CompressGZIP(file_thetan_world_v1_service_adapter_proto_rawDescData)
	})
	return file_thetan_world_v1_service_adapter_proto_rawDescData
}

var file_thetan_world_v1_service_adapter_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_thetan_world_v1_service_adapter_proto_goTypes = []interface{}{
	(*ItemResp)(nil),            // 0: thetan.world.v1.ItemResp
	(*AvailableItem)(nil),       // 1: thetan.world.v1.AvailableItem
	(*SimpleItem)(nil),          // 2: thetan.world.v1.SimpleItem
	(*IsValidItemsRequest)(nil), // 3: thetan.world.v1.IsValidItemsRequest
	(*CheckItemResponse)(nil),   // 4: thetan.world.v1.CheckItemResponse
	(*SendItemsRequest)(nil),    // 5: thetan.world.v1.SendItemsRequest
	(*SendItemsResponse)(nil),   // 6: thetan.world.v1.SendItemsResponse
	(*emptypb.Empty)(nil),       // 7: google.protobuf.Empty
}
var file_thetan_world_v1_service_adapter_proto_depIdxs = []int32{
	0, // 0: thetan.world.v1.AvailableItem.items:type_name -> thetan.world.v1.ItemResp
	2, // 1: thetan.world.v1.IsValidItemsRequest.items:type_name -> thetan.world.v1.SimpleItem
	2, // 2: thetan.world.v1.SendItemsRequest.items:type_name -> thetan.world.v1.SimpleItem
	0, // 3: thetan.world.v1.SendItemsResponse.items:type_name -> thetan.world.v1.ItemResp
	7, // 4: thetan.world.v1.ThetanWorldAdapterService.GetAvailableItems:input_type -> google.protobuf.Empty
	3, // 5: thetan.world.v1.ThetanWorldAdapterService.IsValidItems:input_type -> thetan.world.v1.IsValidItemsRequest
	5, // 6: thetan.world.v1.ThetanWorldAdapterService.SendItems:input_type -> thetan.world.v1.SendItemsRequest
	1, // 7: thetan.world.v1.ThetanWorldAdapterService.GetAvailableItems:output_type -> thetan.world.v1.AvailableItem
	4, // 8: thetan.world.v1.ThetanWorldAdapterService.IsValidItems:output_type -> thetan.world.v1.CheckItemResponse
	6, // 9: thetan.world.v1.ThetanWorldAdapterService.SendItems:output_type -> thetan.world.v1.SendItemsResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_thetan_world_v1_service_adapter_proto_init() }
func file_thetan_world_v1_service_adapter_proto_init() {
	if File_thetan_world_v1_service_adapter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thetan_world_v1_service_adapter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemResp); i {
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
		file_thetan_world_v1_service_adapter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_thetan_world_v1_service_adapter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_thetan_world_v1_service_adapter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsValidItemsRequest); i {
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
		file_thetan_world_v1_service_adapter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_thetan_world_v1_service_adapter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_thetan_world_v1_service_adapter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendItemsResponse); i {
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
			RawDescriptor: file_thetan_world_v1_service_adapter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thetan_world_v1_service_adapter_proto_goTypes,
		DependencyIndexes: file_thetan_world_v1_service_adapter_proto_depIdxs,
		MessageInfos:      file_thetan_world_v1_service_adapter_proto_msgTypes,
	}.Build()
	File_thetan_world_v1_service_adapter_proto = out.File
	file_thetan_world_v1_service_adapter_proto_rawDesc = nil
	file_thetan_world_v1_service_adapter_proto_goTypes = nil
	file_thetan_world_v1_service_adapter_proto_depIdxs = nil
}
