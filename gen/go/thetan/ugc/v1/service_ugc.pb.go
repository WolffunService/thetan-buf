// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: thetan/ugc/v1/service_ugc.proto

package thetan_ugc_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
	_ "thetan-buf/gen/go/thetan/shared/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserColoringRequest_ColoringFilter int32

const (
	UserColoringRequest_ALL      UserColoringRequest_ColoringFilter = 0
	UserColoringRequest_THENION  UserColoringRequest_ColoringFilter = 1
	UserColoringRequest_COSMETIC UserColoringRequest_ColoringFilter = 2
)

// Enum value maps for UserColoringRequest_ColoringFilter.
var (
	UserColoringRequest_ColoringFilter_name = map[int32]string{
		0: "ALL",
		1: "THENION",
		2: "COSMETIC",
	}
	UserColoringRequest_ColoringFilter_value = map[string]int32{
		"ALL":      0,
		"THENION":  1,
		"COSMETIC": 2,
	}
)

func (x UserColoringRequest_ColoringFilter) Enum() *UserColoringRequest_ColoringFilter {
	p := new(UserColoringRequest_ColoringFilter)
	*p = x
	return p
}

func (x UserColoringRequest_ColoringFilter) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserColoringRequest_ColoringFilter) Descriptor() protoreflect.EnumDescriptor {
	return file_thetan_ugc_v1_service_ugc_proto_enumTypes[0].Descriptor()
}

func (UserColoringRequest_ColoringFilter) Type() protoreflect.EnumType {
	return &file_thetan_ugc_v1_service_ugc_proto_enumTypes[0]
}

func (x UserColoringRequest_ColoringFilter) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserColoringRequest_ColoringFilter.Descriptor instead.
func (UserColoringRequest_ColoringFilter) EnumDescriptor() ([]byte, []int) {
	return file_thetan_ugc_v1_service_ugc_proto_rawDescGZIP(), []int{0, 0}
}

type UserColoringRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string                             `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Filter UserColoringRequest_ColoringFilter `protobuf:"varint,2,opt,name=filter,proto3,enum=thetan.ugc.v1.UserColoringRequest_ColoringFilter" json:"filter,omitempty"`
}

func (x *UserColoringRequest) Reset() {
	*x = UserColoringRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_ugc_v1_service_ugc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserColoringRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserColoringRequest) ProtoMessage() {}

func (x *UserColoringRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_ugc_v1_service_ugc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserColoringRequest.ProtoReflect.Descriptor instead.
func (*UserColoringRequest) Descriptor() ([]byte, []int) {
	return file_thetan_ugc_v1_service_ugc_proto_rawDescGZIP(), []int{0}
}

func (x *UserColoringRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserColoringRequest) GetFilter() UserColoringRequest_ColoringFilter {
	if x != nil {
		return x.Filter
	}
	return UserColoringRequest_ALL
}

type UserOneColoringRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ColoringId string `protobuf:"bytes,1,opt,name=coloring_id,json=coloringId,proto3" json:"coloring_id,omitempty"`
}

func (x *UserOneColoringRequest) Reset() {
	*x = UserOneColoringRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_ugc_v1_service_ugc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOneColoringRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOneColoringRequest) ProtoMessage() {}

func (x *UserOneColoringRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_ugc_v1_service_ugc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserOneColoringRequest.ProtoReflect.Descriptor instead.
func (*UserOneColoringRequest) Descriptor() ([]byte, []int) {
	return file_thetan_ugc_v1_service_ugc_proto_rawDescGZIP(), []int{1}
}

func (x *UserOneColoringRequest) GetColoringId() string {
	if x != nil {
		return x.ColoringId
	}
	return ""
}

type UserColoringResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Colorings []*UserColoring `protobuf:"bytes,1,rep,name=colorings,proto3" json:"colorings,omitempty"`
}

func (x *UserColoringResponse) Reset() {
	*x = UserColoringResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_ugc_v1_service_ugc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserColoringResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserColoringResponse) ProtoMessage() {}

func (x *UserColoringResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_ugc_v1_service_ugc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserColoringResponse.ProtoReflect.Descriptor instead.
func (*UserColoringResponse) Descriptor() ([]byte, []int) {
	return file_thetan_ugc_v1_service_ugc_proto_rawDescGZIP(), []int{2}
}

func (x *UserColoringResponse) GetColorings() []*UserColoring {
	if x != nil {
		return x.Colorings
	}
	return nil
}

type FileAttrs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Bucket   string `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Size     int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Md5      string `protobuf:"bytes,4,opt,name=md5,proto3" json:"md5,omitempty"`
	Url      string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	BlurHash string `protobuf:"bytes,6,opt,name=blur_hash,json=blurHash,proto3" json:"blur_hash,omitempty"`
}

func (x *FileAttrs) Reset() {
	*x = FileAttrs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_ugc_v1_service_ugc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileAttrs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileAttrs) ProtoMessage() {}

func (x *FileAttrs) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_ugc_v1_service_ugc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileAttrs.ProtoReflect.Descriptor instead.
func (*FileAttrs) Descriptor() ([]byte, []int) {
	return file_thetan_ugc_v1_service_ugc_proto_rawDescGZIP(), []int{3}
}

func (x *FileAttrs) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileAttrs) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *FileAttrs) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileAttrs) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *FileAttrs) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *FileAttrs) GetBlurHash() string {
	if x != nil {
		return x.BlurHash
	}
	return ""
}

type ShaderAttrs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShaderId int32            `protobuf:"varint,1,opt,name=shader_id,json=shaderId,proto3" json:"shader_id,omitempty"`
	Params   *structpb.Struct `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *ShaderAttrs) Reset() {
	*x = ShaderAttrs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_ugc_v1_service_ugc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShaderAttrs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShaderAttrs) ProtoMessage() {}

func (x *ShaderAttrs) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_ugc_v1_service_ugc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShaderAttrs.ProtoReflect.Descriptor instead.
func (*ShaderAttrs) Descriptor() ([]byte, []int) {
	return file_thetan_ugc_v1_service_ugc_proto_rawDescGZIP(), []int{4}
}

func (x *ShaderAttrs) GetShaderId() int32 {
	if x != nil {
		return x.ShaderId
	}
	return 0
}

func (x *ShaderAttrs) GetParams() *structpb.Struct {
	if x != nil {
		return x.Params
	}
	return nil
}

type UserColoring struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId                 string           `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GameId                 int32            `protobuf:"varint,3,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	KindId                 int32            `protobuf:"varint,4,opt,name=kind_id,json=kindId,proto3" json:"kind_id,omitempty"`
	TypeId                 int32            `protobuf:"varint,5,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	PublishCount           int32            `protobuf:"varint,6,opt,name=publish_count,json=publishCount,proto3" json:"publish_count,omitempty"`
	DisplayName            string           `protobuf:"bytes,7,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Revision               int32            `protobuf:"varint,8,opt,name=revision,proto3" json:"revision,omitempty"`
	TextureLastPublished   *FileAttrs       `protobuf:"bytes,9,opt,name=textureLastPublished,proto3" json:"textureLastPublished,omitempty"`
	TextureLastSaved       *FileAttrs       `protobuf:"bytes,10,opt,name=textureLastSaved,proto3" json:"textureLastSaved,omitempty"`
	ThumbnailLastPublished *FileAttrs       `protobuf:"bytes,11,opt,name=thumbnailLastPublished,proto3" json:"thumbnailLastPublished,omitempty"`
	ThumbnailLastSaved     *FileAttrs       `protobuf:"bytes,12,opt,name=thumbnailLastSaved,proto3" json:"thumbnailLastSaved,omitempty"`
	ShaderLastPublished    *ShaderAttrs     `protobuf:"bytes,13,opt,name=shaderLastPublished,proto3" json:"shaderLastPublished,omitempty"`
	ShaderLastSaved        *ShaderAttrs     `protobuf:"bytes,14,opt,name=shaderLastSaved,proto3" json:"shaderLastSaved,omitempty"`
	Props                  *structpb.Struct `protobuf:"bytes,15,opt,name=props,proto3" json:"props,omitempty"`
	UpdatedAt              int64            `protobuf:"varint,16,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedAt              int64            `protobuf:"varint,17,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *UserColoring) Reset() {
	*x = UserColoring{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_ugc_v1_service_ugc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserColoring) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserColoring) ProtoMessage() {}

func (x *UserColoring) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_ugc_v1_service_ugc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserColoring.ProtoReflect.Descriptor instead.
func (*UserColoring) Descriptor() ([]byte, []int) {
	return file_thetan_ugc_v1_service_ugc_proto_rawDescGZIP(), []int{5}
}

func (x *UserColoring) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserColoring) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserColoring) GetGameId() int32 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *UserColoring) GetKindId() int32 {
	if x != nil {
		return x.KindId
	}
	return 0
}

func (x *UserColoring) GetTypeId() int32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *UserColoring) GetPublishCount() int32 {
	if x != nil {
		return x.PublishCount
	}
	return 0
}

func (x *UserColoring) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *UserColoring) GetRevision() int32 {
	if x != nil {
		return x.Revision
	}
	return 0
}

func (x *UserColoring) GetTextureLastPublished() *FileAttrs {
	if x != nil {
		return x.TextureLastPublished
	}
	return nil
}

func (x *UserColoring) GetTextureLastSaved() *FileAttrs {
	if x != nil {
		return x.TextureLastSaved
	}
	return nil
}

func (x *UserColoring) GetThumbnailLastPublished() *FileAttrs {
	if x != nil {
		return x.ThumbnailLastPublished
	}
	return nil
}

func (x *UserColoring) GetThumbnailLastSaved() *FileAttrs {
	if x != nil {
		return x.ThumbnailLastSaved
	}
	return nil
}

func (x *UserColoring) GetShaderLastPublished() *ShaderAttrs {
	if x != nil {
		return x.ShaderLastPublished
	}
	return nil
}

func (x *UserColoring) GetShaderLastSaved() *ShaderAttrs {
	if x != nil {
		return x.ShaderLastSaved
	}
	return nil
}

func (x *UserColoring) GetProps() *structpb.Struct {
	if x != nil {
		return x.Props
	}
	return nil
}

func (x *UserColoring) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *UserColoring) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

var File_thetan_ugc_v1_service_ugc_proto protoreflect.FileDescriptor

var file_thetan_ugc_v1_service_ugc_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x75, 0x67, 0x63, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x67, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x75, 0x67, 0x63, 0x2e, 0x76, 0x31,
	0x1a, 0x21, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x49, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x31, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x75, 0x67, 0x63, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x34, 0x0a,
	0x0e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x48, 0x45, 0x4e,
	0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x53, 0x4d, 0x45, 0x54, 0x49,
	0x43, 0x10, 0x02, 0x22, 0x39, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x6e, 0x65, 0x43, 0x6f,
	0x6c, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x51,
	0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x2e, 0x75, 0x67, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f,
	0x6c, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x09, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x73, 0x22, 0x8c, 0x01, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x64, 0x35, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x64,
	0x35, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6c, 0x75, 0x72, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6c, 0x75, 0x72, 0x48, 0x61, 0x73, 0x68,
	0x22, 0x5b, 0x0a, 0x0b, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x68, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x73, 0x68, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x97, 0x06,
	0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x6b, 0x69, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6b, 0x69, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x14, 0x74, 0x65, 0x78, 0x74, 0x75, 0x72,
	0x65, 0x4c, 0x61, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x75, 0x67,
	0x63, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x73, 0x52, 0x14,
	0x74, 0x65, 0x78, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x12, 0x44, 0x0a, 0x10, 0x74, 0x65, 0x78, 0x74, 0x75, 0x72, 0x65, 0x4c,
	0x61, 0x73, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x75, 0x67, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x73, 0x52, 0x10, 0x74, 0x65, 0x78, 0x74, 0x75, 0x72,
	0x65, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x12, 0x50, 0x0a, 0x16, 0x74, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x4c, 0x61, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x68, 0x65,
	0x74, 0x61, 0x6e, 0x2e, 0x75, 0x67, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x41,
	0x74, 0x74, 0x72, 0x73, 0x52, 0x16, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x4c,
	0x61, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x48, 0x0a, 0x12,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x61, 0x76,
	0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61,
	0x6e, 0x2e, 0x75, 0x67, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x41, 0x74, 0x74,
	0x72, 0x73, 0x52, 0x12, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x4c, 0x61, 0x73,
	0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x12, 0x4c, 0x0a, 0x13, 0x73, 0x68, 0x61, 0x64, 0x65, 0x72,
	0x4c, 0x61, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x75, 0x67, 0x63,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x73, 0x52,
	0x13, 0x73, 0x68, 0x61, 0x64, 0x65, 0x72, 0x4c, 0x61, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x12, 0x44, 0x0a, 0x0f, 0x73, 0x68, 0x61, 0x64, 0x65, 0x72, 0x4c, 0x61,
	0x73, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x75, 0x67, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68,
	0x61, 0x64, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x73, 0x52, 0x0f, 0x73, 0x68, 0x61, 0x64, 0x65,
	0x72, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x12, 0x2d, 0x0a, 0x05, 0x70, 0x72,
	0x6f, 0x70, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xcd, 0x01, 0x0a, 0x10, 0x54, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x55, 0x47, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x22, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x75, 0x67, 0x63, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x75, 0x67,
	0x63, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x4f, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x25, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x75, 0x67, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x6e, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61,
	0x6e, 0x2e, 0x75, 0x67, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x42, 0xa9, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e,
	0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x75, 0x67, 0x63, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x67, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x2d, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2d, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x75, 0x67, 0x63, 0x2f, 0x76,
	0x31, 0x3b, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5f, 0x75, 0x67, 0x63, 0x5f, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x54, 0x55, 0x58, 0xaa, 0x02, 0x0d, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x55,
	0x67, 0x63, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x55,
	0x67, 0x63, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x55,
	0x67, 0x63, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0f, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x3a, 0x3a, 0x55, 0x67, 0x63, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thetan_ugc_v1_service_ugc_proto_rawDescOnce sync.Once
	file_thetan_ugc_v1_service_ugc_proto_rawDescData = file_thetan_ugc_v1_service_ugc_proto_rawDesc
)

func file_thetan_ugc_v1_service_ugc_proto_rawDescGZIP() []byte {
	file_thetan_ugc_v1_service_ugc_proto_rawDescOnce.Do(func() {
		file_thetan_ugc_v1_service_ugc_proto_rawDescData = protoimpl.X.CompressGZIP(file_thetan_ugc_v1_service_ugc_proto_rawDescData)
	})
	return file_thetan_ugc_v1_service_ugc_proto_rawDescData
}

var file_thetan_ugc_v1_service_ugc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_thetan_ugc_v1_service_ugc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_thetan_ugc_v1_service_ugc_proto_goTypes = []interface{}{
	(UserColoringRequest_ColoringFilter)(0), // 0: thetan.ugc.v1.UserColoringRequest.ColoringFilter
	(*UserColoringRequest)(nil),             // 1: thetan.ugc.v1.UserColoringRequest
	(*UserOneColoringRequest)(nil),          // 2: thetan.ugc.v1.UserOneColoringRequest
	(*UserColoringResponse)(nil),            // 3: thetan.ugc.v1.UserColoringResponse
	(*FileAttrs)(nil),                       // 4: thetan.ugc.v1.FileAttrs
	(*ShaderAttrs)(nil),                     // 5: thetan.ugc.v1.ShaderAttrs
	(*UserColoring)(nil),                    // 6: thetan.ugc.v1.UserColoring
	(*structpb.Struct)(nil),                 // 7: google.protobuf.Struct
}
var file_thetan_ugc_v1_service_ugc_proto_depIdxs = []int32{
	0,  // 0: thetan.ugc.v1.UserColoringRequest.filter:type_name -> thetan.ugc.v1.UserColoringRequest.ColoringFilter
	6,  // 1: thetan.ugc.v1.UserColoringResponse.colorings:type_name -> thetan.ugc.v1.UserColoring
	7,  // 2: thetan.ugc.v1.ShaderAttrs.params:type_name -> google.protobuf.Struct
	4,  // 3: thetan.ugc.v1.UserColoring.textureLastPublished:type_name -> thetan.ugc.v1.FileAttrs
	4,  // 4: thetan.ugc.v1.UserColoring.textureLastSaved:type_name -> thetan.ugc.v1.FileAttrs
	4,  // 5: thetan.ugc.v1.UserColoring.thumbnailLastPublished:type_name -> thetan.ugc.v1.FileAttrs
	4,  // 6: thetan.ugc.v1.UserColoring.thumbnailLastSaved:type_name -> thetan.ugc.v1.FileAttrs
	5,  // 7: thetan.ugc.v1.UserColoring.shaderLastPublished:type_name -> thetan.ugc.v1.ShaderAttrs
	5,  // 8: thetan.ugc.v1.UserColoring.shaderLastSaved:type_name -> thetan.ugc.v1.ShaderAttrs
	7,  // 9: thetan.ugc.v1.UserColoring.props:type_name -> google.protobuf.Struct
	1,  // 10: thetan.ugc.v1.ThetanUGCService.GetUserColorings:input_type -> thetan.ugc.v1.UserColoringRequest
	2,  // 11: thetan.ugc.v1.ThetanUGCService.GetOneUserColoring:input_type -> thetan.ugc.v1.UserOneColoringRequest
	3,  // 12: thetan.ugc.v1.ThetanUGCService.GetUserColorings:output_type -> thetan.ugc.v1.UserColoringResponse
	6,  // 13: thetan.ugc.v1.ThetanUGCService.GetOneUserColoring:output_type -> thetan.ugc.v1.UserColoring
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_thetan_ugc_v1_service_ugc_proto_init() }
func file_thetan_ugc_v1_service_ugc_proto_init() {
	if File_thetan_ugc_v1_service_ugc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thetan_ugc_v1_service_ugc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserColoringRequest); i {
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
		file_thetan_ugc_v1_service_ugc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserOneColoringRequest); i {
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
		file_thetan_ugc_v1_service_ugc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserColoringResponse); i {
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
		file_thetan_ugc_v1_service_ugc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileAttrs); i {
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
		file_thetan_ugc_v1_service_ugc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShaderAttrs); i {
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
		file_thetan_ugc_v1_service_ugc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserColoring); i {
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
			RawDescriptor: file_thetan_ugc_v1_service_ugc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thetan_ugc_v1_service_ugc_proto_goTypes,
		DependencyIndexes: file_thetan_ugc_v1_service_ugc_proto_depIdxs,
		EnumInfos:         file_thetan_ugc_v1_service_ugc_proto_enumTypes,
		MessageInfos:      file_thetan_ugc_v1_service_ugc_proto_msgTypes,
	}.Build()
	File_thetan_ugc_v1_service_ugc_proto = out.File
	file_thetan_ugc_v1_service_ugc_proto_rawDesc = nil
	file_thetan_ugc_v1_service_ugc_proto_goTypes = nil
	file_thetan_ugc_v1_service_ugc_proto_depIdxs = nil
}
