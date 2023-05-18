// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0-devel
// 	protoc        v3.21.12
// source: service_firebase.proto

package srvfirebase

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

type TrackPlayerStatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	StatName  string `protobuf:"bytes,2,opt,name=statName,proto3" json:"statName,omitempty"`
	StatValue string `protobuf:"bytes,3,opt,name=statValue,proto3" json:"statValue,omitempty"`
}

func (x *TrackPlayerStatRequest) Reset() {
	*x = TrackPlayerStatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_firebase_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackPlayerStatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackPlayerStatRequest) ProtoMessage() {}

func (x *TrackPlayerStatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_firebase_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackPlayerStatRequest.ProtoReflect.Descriptor instead.
func (*TrackPlayerStatRequest) Descriptor() ([]byte, []int) {
	return file_service_firebase_proto_rawDescGZIP(), []int{0}
}

func (x *TrackPlayerStatRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TrackPlayerStatRequest) GetStatName() string {
	if x != nil {
		return x.StatName
	}
	return ""
}

func (x *TrackPlayerStatRequest) GetStatValue() string {
	if x != nil {
		return x.StatValue
	}
	return ""
}

var File_service_firebase_proto protoreflect.FileDescriptor

var file_service_firebase_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6a, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x74, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x60, 0x0a, 0x0f, 0x46,
	0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d,
	0x0a, 0x0f, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x12, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0f, 0x5a,
	0x0d, 0x2e, 0x2f, 0x73, 0x72, 0x76, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_firebase_proto_rawDescOnce sync.Once
	file_service_firebase_proto_rawDescData = file_service_firebase_proto_rawDesc
)

func file_service_firebase_proto_rawDescGZIP() []byte {
	file_service_firebase_proto_rawDescOnce.Do(func() {
		file_service_firebase_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_firebase_proto_rawDescData)
	})
	return file_service_firebase_proto_rawDescData
}

var file_service_firebase_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_service_firebase_proto_goTypes = []interface{}{
	(*TrackPlayerStatRequest)(nil), // 0: services.TrackPlayerStatRequest
	(*emptypb.Empty)(nil),          // 1: google.protobuf.Empty
}
var file_service_firebase_proto_depIdxs = []int32{
	0, // 0: services.FirebaseService.TrackPlayerStat:input_type -> services.TrackPlayerStatRequest
	1, // 1: services.FirebaseService.TrackPlayerStat:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_firebase_proto_init() }
func file_service_firebase_proto_init() {
	if File_service_firebase_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_firebase_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackPlayerStatRequest); i {
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
			RawDescriptor: file_service_firebase_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_firebase_proto_goTypes,
		DependencyIndexes: file_service_firebase_proto_depIdxs,
		MessageInfos:      file_service_firebase_proto_msgTypes,
	}.Build()
	File_service_firebase_proto = out.File
	file_service_firebase_proto_rawDesc = nil
	file_service_firebase_proto_goTypes = nil
	file_service_firebase_proto_depIdxs = nil
}
