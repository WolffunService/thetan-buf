// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: thetan/shared/v1/customized.proto

package thetan_shared_v1

import (
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CustomizeType int32

const (
	CustomizeType_TEXTURE   CustomizeType = 0
	CustomizeType_ANIMATION CustomizeType = 1
)

// Enum value maps for CustomizeType.
var (
	CustomizeType_name = map[int32]string{
		0: "TEXTURE",
		1: "ANIMATION",
	}
	CustomizeType_value = map[string]int32{
		"TEXTURE":   0,
		"ANIMATION": 1,
	}
)

func (x CustomizeType) Enum() *CustomizeType {
	p := new(CustomizeType)
	*p = x
	return p
}

func (x CustomizeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CustomizeType) Descriptor() protoreflect.EnumDescriptor {
	return file_thetan_shared_v1_customized_proto_enumTypes[0].Descriptor()
}

func (CustomizeType) Type() protoreflect.EnumType {
	return &file_thetan_shared_v1_customized_proto_enumTypes[0]
}

func (x CustomizeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CustomizeType.Descriptor instead.
func (CustomizeType) EnumDescriptor() ([]byte, []int) {
	return file_thetan_shared_v1_customized_proto_rawDescGZIP(), []int{0}
}

// Customized dùng để gắn vào các item để cho phép app Thetan Creator có thể lưu được bản vẽ.
type Customized struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type CustomizeType `protobuf:"varint,6,opt,name=type,proto3,enum=thetan.shared.v1.CustomizeType" json:"type,omitempty" bson:"customType"`
	// Name cho phép user tự đặt tên cho item của mình ở Thetan Creator
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" bson:"name"`
	// ===== TEXTURE ========
	// TextureURL đường dẫn tuyệt đối đến file texture lưu trên cloud
	TextureURL string `protobuf:"bytes,2,opt,name=textureURL,proto3" json:"textureURL,omitempty" bson:"textureURL"`
	// ThumbnailURL đường dẫn tuyệt đối đến file thumbnail lưu trên cloud
	ThumbnailURL string `protobuf:"bytes,5,opt,name=thumbnailURL,proto3" json:"thumbnailURL,omitempty" bson:"thumbnailURL"`
	// CustomShader shader và các cấu hình của shader đó
	Shader *CustomizedShader `protobuf:"bytes,4,opt,name=shader,proto3,oneof" json:"shader,omitempty" bson:"shader"`
	// ===== ANIMATION ========
	AnimationURL *string `protobuf:"bytes,7,opt,name=animationURL,proto3,oneof" json:"animationURL,omitempty" bson:"animationURL"`
	// PublishAt thời điểm publish item customized này
	PublishAt int64 `protobuf:"varint,3,opt,name=publishAt,proto3" json:"publishAt,omitempty" bson:"publishAt"`
}

func (x *Customized) Reset() {
	*x = Customized{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_shared_v1_customized_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customized) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customized) ProtoMessage() {}

func (x *Customized) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_shared_v1_customized_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customized.ProtoReflect.Descriptor instead.
func (*Customized) Descriptor() ([]byte, []int) {
	return file_thetan_shared_v1_customized_proto_rawDescGZIP(), []int{0}
}

func (x *Customized) GetType() CustomizeType {
	if x != nil {
		return x.Type
	}
	return CustomizeType_TEXTURE
}

func (x *Customized) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Customized) GetTextureURL() string {
	if x != nil {
		return x.TextureURL
	}
	return ""
}

func (x *Customized) GetThumbnailURL() string {
	if x != nil {
		return x.ThumbnailURL
	}
	return ""
}

func (x *Customized) GetShader() *CustomizedShader {
	if x != nil {
		return x.Shader
	}
	return nil
}

func (x *Customized) GetAnimationURL() string {
	if x != nil && x.AnimationURL != nil {
		return *x.AnimationURL
	}
	return ""
}

func (x *Customized) GetPublishAt() int64 {
	if x != nil {
		return x.PublishAt
	}
	return 0
}

type CustomizedShader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShaderID int32            `protobuf:"varint,1,opt,name=shaderID,proto3" json:"shaderID,omitempty" bson:"shaderID"`
	Params   *structpb.Struct `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty" bson:"params"`
}

func (x *CustomizedShader) Reset() {
	*x = CustomizedShader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_shared_v1_customized_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomizedShader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomizedShader) ProtoMessage() {}

func (x *CustomizedShader) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_shared_v1_customized_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomizedShader.ProtoReflect.Descriptor instead.
func (*CustomizedShader) Descriptor() ([]byte, []int) {
	return file_thetan_shared_v1_customized_proto_rawDescGZIP(), []int{1}
}

func (x *CustomizedShader) GetShaderID() int32 {
	if x != nil {
		return x.ShaderID
	}
	return 0
}

func (x *CustomizedShader) GetParams() *structpb.Struct {
	if x != nil {
		return x.Params
	}
	return nil
}

type ItemCustomized struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       int32       `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty" bson:"type"`
	Kind       int32       `protobuf:"varint,2,opt,name=kind,proto3" json:"kind,omitempty" bson:"kind"`
	Id         string      `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty" bson:"id"`
	Customized *Customized `protobuf:"bytes,4,opt,name=customized,proto3" json:"customized,omitempty" bson:"customized"`
}

func (x *ItemCustomized) Reset() {
	*x = ItemCustomized{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_shared_v1_customized_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemCustomized) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemCustomized) ProtoMessage() {}

func (x *ItemCustomized) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_shared_v1_customized_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemCustomized.ProtoReflect.Descriptor instead.
func (*ItemCustomized) Descriptor() ([]byte, []int) {
	return file_thetan_shared_v1_customized_proto_rawDescGZIP(), []int{2}
}

func (x *ItemCustomized) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ItemCustomized) GetKind() int32 {
	if x != nil {
		return x.Kind
	}
	return 0
}

func (x *ItemCustomized) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ItemCustomized) GetCustomized() *Customized {
	if x != nil {
		return x.Customized
	}
	return nil
}

var File_thetan_shared_v1_customized_proto protoreflect.FileDescriptor

var file_thetan_shared_v1_customized_proto_rawDesc = []byte{
	0x0a, 0x21, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x03, 0x0a, 0x0a, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x4b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69,
	0x7a, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x16, 0x9a, 0x84, 0x9e, 0x03, 0x11, 0x62, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x22, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x74, 0x65,
	0x78, 0x74, 0x75, 0x72, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16,
	0x9a, 0x84, 0x9e, 0x03, 0x11, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x65, 0x78, 0x74, 0x75,
	0x72, 0x65, 0x55, 0x52, 0x4c, 0x22, 0x52, 0x0a, 0x74, 0x65, 0x78, 0x74, 0x75, 0x72, 0x65, 0x55,
	0x52, 0x4c, 0x12, 0x3c, 0x0a, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55,
	0x52, 0x4c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0x9a, 0x84, 0x9e, 0x03, 0x13, 0x62,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x52,
	0x4c, 0x22, 0x52, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x52, 0x4c,
	0x12, 0x53, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x53, 0x68,
	0x61, 0x64, 0x65, 0x72, 0x42, 0x12, 0x9a, 0x84, 0x9e, 0x03, 0x0d, 0x62, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x73, 0x68, 0x61, 0x64, 0x65, 0x72, 0x22, 0x48, 0x00, 0x52, 0x06, 0x73, 0x68, 0x61, 0x64,
	0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x41, 0x0a, 0x0c, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x55, 0x52, 0x4c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0x9a, 0x84, 0x9e,
	0x03, 0x13, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x52, 0x4c, 0x22, 0x48, 0x01, 0x52, 0x0c, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x55, 0x52, 0x4c, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x15, 0x9a, 0x84, 0x9e,
	0x03, 0x10, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x41,
	0x74, 0x22, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x41, 0x74, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x73, 0x68, 0x61, 0x64, 0x65, 0x72, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x61, 0x6e, 0x69,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x52, 0x4c, 0x22, 0x8f, 0x01, 0x0a, 0x10, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x12, 0x30,
	0x0a, 0x08, 0x73, 0x68, 0x61, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x14, 0x9a, 0x84, 0x9e, 0x03, 0x0f, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x68, 0x61,
	0x64, 0x65, 0x72, 0x49, 0x44, 0x22, 0x52, 0x08, 0x73, 0x68, 0x61, 0x64, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x43, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x12, 0x9a, 0x84, 0x9e, 0x03, 0x0d,
	0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x52, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0xd2, 0x01, 0x0a, 0x0e,
	0x49, 0x74, 0x65, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x24,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x10, 0x9a, 0x84,
	0x9e, 0x03, 0x0b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6b,
	0x69, 0x6e, 0x64, 0x22, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0x9a, 0x84, 0x9e, 0x03, 0x09, 0x62, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x54, 0x0a, 0x0a, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x42, 0x16, 0x9a, 0x84,
	0x9e, 0x03, 0x11, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69,
	0x7a, 0x65, 0x64, 0x22, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64,
	0x2a, 0x2b, 0x0a, 0x0d, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x45, 0x58, 0x54, 0x55, 0x52, 0x45, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x41, 0x4e, 0x49, 0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x42, 0xbe, 0x01,
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a,
	0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x74, 0x68, 0x65, 0x74, 0x61,
	0x6e, 0x2d, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x68, 0x65,
	0x74, 0x61, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x54, 0x53, 0x58, 0xaa, 0x02, 0x10, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e,
	0x5c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x54, 0x68, 0x65,
	0x74, 0x61, 0x6e, 0x5c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x54, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x3a, 0x3a, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thetan_shared_v1_customized_proto_rawDescOnce sync.Once
	file_thetan_shared_v1_customized_proto_rawDescData = file_thetan_shared_v1_customized_proto_rawDesc
)

func file_thetan_shared_v1_customized_proto_rawDescGZIP() []byte {
	file_thetan_shared_v1_customized_proto_rawDescOnce.Do(func() {
		file_thetan_shared_v1_customized_proto_rawDescData = protoimpl.X.CompressGZIP(file_thetan_shared_v1_customized_proto_rawDescData)
	})
	return file_thetan_shared_v1_customized_proto_rawDescData
}

var file_thetan_shared_v1_customized_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_thetan_shared_v1_customized_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_thetan_shared_v1_customized_proto_goTypes = []interface{}{
	(CustomizeType)(0),       // 0: thetan.shared.v1.CustomizeType
	(*Customized)(nil),       // 1: thetan.shared.v1.Customized
	(*CustomizedShader)(nil), // 2: thetan.shared.v1.CustomizedShader
	(*ItemCustomized)(nil),   // 3: thetan.shared.v1.ItemCustomized
	(*structpb.Struct)(nil),  // 4: google.protobuf.Struct
}
var file_thetan_shared_v1_customized_proto_depIdxs = []int32{
	0, // 0: thetan.shared.v1.Customized.type:type_name -> thetan.shared.v1.CustomizeType
	2, // 1: thetan.shared.v1.Customized.shader:type_name -> thetan.shared.v1.CustomizedShader
	4, // 2: thetan.shared.v1.CustomizedShader.params:type_name -> google.protobuf.Struct
	1, // 3: thetan.shared.v1.ItemCustomized.customized:type_name -> thetan.shared.v1.Customized
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_thetan_shared_v1_customized_proto_init() }
func file_thetan_shared_v1_customized_proto_init() {
	if File_thetan_shared_v1_customized_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thetan_shared_v1_customized_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customized); i {
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
		file_thetan_shared_v1_customized_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomizedShader); i {
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
		file_thetan_shared_v1_customized_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemCustomized); i {
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
	file_thetan_shared_v1_customized_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_thetan_shared_v1_customized_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_thetan_shared_v1_customized_proto_goTypes,
		DependencyIndexes: file_thetan_shared_v1_customized_proto_depIdxs,
		EnumInfos:         file_thetan_shared_v1_customized_proto_enumTypes,
		MessageInfos:      file_thetan_shared_v1_customized_proto_msgTypes,
	}.Build()
	File_thetan_shared_v1_customized_proto = out.File
	file_thetan_shared_v1_customized_proto_rawDesc = nil
	file_thetan_shared_v1_customized_proto_goTypes = nil
	file_thetan_shared_v1_customized_proto_depIdxs = nil
}
