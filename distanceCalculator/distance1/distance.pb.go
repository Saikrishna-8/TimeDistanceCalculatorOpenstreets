// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: distance.proto

package distance1

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

type DistanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAddress string `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	ToAddress   string `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
}

func (x *DistanceRequest) Reset() {
	*x = DistanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_distance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistanceRequest) ProtoMessage() {}

func (x *DistanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_distance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistanceRequest.ProtoReflect.Descriptor instead.
func (*DistanceRequest) Descriptor() ([]byte, []int) {
	return file_distance_proto_rawDescGZIP(), []int{0}
}

func (x *DistanceRequest) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *DistanceRequest) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

type DistanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Distance float64 `protobuf:"fixed64,1,opt,name=distance,proto3" json:"distance,omitempty"`
}

func (x *DistanceResponse) Reset() {
	*x = DistanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_distance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistanceResponse) ProtoMessage() {}

func (x *DistanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_distance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistanceResponse.ProtoReflect.Descriptor instead.
func (*DistanceResponse) Descriptor() ([]byte, []int) {
	return file_distance_proto_rawDescGZIP(), []int{1}
}

func (x *DistanceResponse) GetDistance() float64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

var File_distance_proto protoreflect.FileDescriptor

var file_distance_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x31, 0x22, 0x53, 0x0a, 0x0f, 0x44,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x2e, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x32, 0x65, 0x0a, 0x19, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x2e, 0x64,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x64, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x64, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_distance_proto_rawDescOnce sync.Once
	file_distance_proto_rawDescData = file_distance_proto_rawDesc
)

func file_distance_proto_rawDescGZIP() []byte {
	file_distance_proto_rawDescOnce.Do(func() {
		file_distance_proto_rawDescData = protoimpl.X.CompressGZIP(file_distance_proto_rawDescData)
	})
	return file_distance_proto_rawDescData
}

var file_distance_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_distance_proto_goTypes = []any{
	(*DistanceRequest)(nil),  // 0: distance1.DistanceRequest
	(*DistanceResponse)(nil), // 1: distance1.DistanceResponse
}
var file_distance_proto_depIdxs = []int32{
	0, // 0: distance1.DistanceCalculatorService.GetDistance:input_type -> distance1.DistanceRequest
	1, // 1: distance1.DistanceCalculatorService.GetDistance:output_type -> distance1.DistanceResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_distance_proto_init() }
func file_distance_proto_init() {
	if File_distance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_distance_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DistanceRequest); i {
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
		file_distance_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DistanceResponse); i {
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
			RawDescriptor: file_distance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_distance_proto_goTypes,
		DependencyIndexes: file_distance_proto_depIdxs,
		MessageInfos:      file_distance_proto_msgTypes,
	}.Build()
	File_distance_proto = out.File
	file_distance_proto_rawDesc = nil
	file_distance_proto_goTypes = nil
	file_distance_proto_depIdxs = nil
}
