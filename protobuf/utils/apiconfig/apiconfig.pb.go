// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.20.1
// source: lemon/protobuf/utils/apiconfig/apiconfig.proto

package apiconfig

import (
	annotations "google.golang.org/genproto/googleapis/api/annotations"
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

// GrpcAPIService represents a stripped down version of google.api.Service .
// Compare to https://github.com/googleapis/googleapis/blob/master/google/api/service.proto
// The original imports 23 other protobuf files we are not interested in. If a significant
// subset (>50%) of these start being reproduced in this file we should swap to using the
// full generated version instead.
//
// For the purposes of the gateway generator we only consider a small subset of all
// available features google supports in their service descriptions. Thanks to backwards
// compatibility guarantees by protobuf it is safe for us to remove the other fields.
type GrpcAPIService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Http Rule.
	Http *annotations.Http `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
}

func (x *GrpcAPIService) Reset() {
	*x = GrpcAPIService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lemon_protobuf_utils_apiconfig_apiconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcAPIService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcAPIService) ProtoMessage() {}

func (x *GrpcAPIService) ProtoReflect() protoreflect.Message {
	mi := &file_lemon_protobuf_utils_apiconfig_apiconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcAPIService.ProtoReflect.Descriptor instead.
func (*GrpcAPIService) Descriptor() ([]byte, []int) {
	return file_lemon_protobuf_utils_apiconfig_apiconfig_proto_rawDescGZIP(), []int{0}
}

func (x *GrpcAPIService) GetHttp() *annotations.Http {
	if x != nil {
		return x.Http
	}
	return nil
}

var File_lemon_protobuf_utils_apiconfig_apiconfig_proto protoreflect.FileDescriptor

var file_lemon_protobuf_utils_apiconfig_apiconfig_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x6c, 0x65, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x61, 0x70, 0x69, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x15, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x36, 0x0a, 0x0e, 0x47, 0x72, 0x70, 0x63, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x48, 0x74, 0x74, 0x70, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x42, 0x5c, 0x0a, 0x28, 0x78, 0x79,
	0x7a, 0x2e, 0x68, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x6c, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x70, 0x65, 0x69, 0x6f, 0x2f, 0x6c, 0x65, 0x6d, 0x6f, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lemon_protobuf_utils_apiconfig_apiconfig_proto_rawDescOnce sync.Once
	file_lemon_protobuf_utils_apiconfig_apiconfig_proto_rawDescData = file_lemon_protobuf_utils_apiconfig_apiconfig_proto_rawDesc
)

func file_lemon_protobuf_utils_apiconfig_apiconfig_proto_rawDescGZIP() []byte {
	file_lemon_protobuf_utils_apiconfig_apiconfig_proto_rawDescOnce.Do(func() {
		file_lemon_protobuf_utils_apiconfig_apiconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_lemon_protobuf_utils_apiconfig_apiconfig_proto_rawDescData)
	})
	return file_lemon_protobuf_utils_apiconfig_apiconfig_proto_rawDescData
}

var file_lemon_protobuf_utils_apiconfig_apiconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_lemon_protobuf_utils_apiconfig_apiconfig_proto_goTypes = []interface{}{
	(*GrpcAPIService)(nil),   // 0: apiconfig.GrpcAPIService
	(*annotations.Http)(nil), // 1: google.api.Http
}
var file_lemon_protobuf_utils_apiconfig_apiconfig_proto_depIdxs = []int32{
	1, // 0: apiconfig.GrpcAPIService.http:type_name -> google.api.Http
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_lemon_protobuf_utils_apiconfig_apiconfig_proto_init() }
func file_lemon_protobuf_utils_apiconfig_apiconfig_proto_init() {
	if File_lemon_protobuf_utils_apiconfig_apiconfig_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lemon_protobuf_utils_apiconfig_apiconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcAPIService); i {
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
			RawDescriptor: file_lemon_protobuf_utils_apiconfig_apiconfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lemon_protobuf_utils_apiconfig_apiconfig_proto_goTypes,
		DependencyIndexes: file_lemon_protobuf_utils_apiconfig_apiconfig_proto_depIdxs,
		MessageInfos:      file_lemon_protobuf_utils_apiconfig_apiconfig_proto_msgTypes,
	}.Build()
	File_lemon_protobuf_utils_apiconfig_apiconfig_proto = out.File
	file_lemon_protobuf_utils_apiconfig_apiconfig_proto_rawDesc = nil
	file_lemon_protobuf_utils_apiconfig_apiconfig_proto_goTypes = nil
	file_lemon_protobuf_utils_apiconfig_apiconfig_proto_depIdxs = nil
}
