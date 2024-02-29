// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: matches/matches.proto

package matches

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

// Requests
type CreateMatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CompetitionId int32  `protobuf:"varint,2,opt,name=competition_id,json=competitionId,proto3" json:"competition_id,omitempty"`
}

func (x *CreateMatchRequest) Reset() {
	*x = CreateMatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matches_matches_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMatchRequest) ProtoMessage() {}

func (x *CreateMatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_matches_matches_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMatchRequest.ProtoReflect.Descriptor instead.
func (*CreateMatchRequest) Descriptor() ([]byte, []int) {
	return file_matches_matches_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMatchRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateMatchRequest) GetCompetitionId() int32 {
	if x != nil {
		return x.CompetitionId
	}
	return 0
}

type UpdateMatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	CompetitionId *int32  `protobuf:"varint,3,opt,name=competition_id,json=competitionId,proto3,oneof" json:"competition_id,omitempty"`
}

func (x *UpdateMatchRequest) Reset() {
	*x = UpdateMatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matches_matches_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMatchRequest) ProtoMessage() {}

func (x *UpdateMatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_matches_matches_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMatchRequest.ProtoReflect.Descriptor instead.
func (*UpdateMatchRequest) Descriptor() ([]byte, []int) {
	return file_matches_matches_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateMatchRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateMatchRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateMatchRequest) GetCompetitionId() int32 {
	if x != nil && x.CompetitionId != nil {
		return *x.CompetitionId
	}
	return 0
}

// Responses
type CreateMatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateMatchResponse) Reset() {
	*x = CreateMatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matches_matches_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMatchResponse) ProtoMessage() {}

func (x *CreateMatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_matches_matches_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMatchResponse.ProtoReflect.Descriptor instead.
func (*CreateMatchResponse) Descriptor() ([]byte, []int) {
	return file_matches_matches_proto_rawDescGZIP(), []int{2}
}

func (x *CreateMatchResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_matches_matches_proto protoreflect.FileDescriptor

var file_matches_matches_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x85, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52,
	0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x22, 0x25, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x32, 0x83, 0x01, 0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73,
	0x12, 0x3a, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x13, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x6d, 0x61,
	0x69, 0x6e, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_matches_matches_proto_rawDescOnce sync.Once
	file_matches_matches_proto_rawDescData = file_matches_matches_proto_rawDesc
)

func file_matches_matches_proto_rawDescGZIP() []byte {
	file_matches_matches_proto_rawDescOnce.Do(func() {
		file_matches_matches_proto_rawDescData = protoimpl.X.CompressGZIP(file_matches_matches_proto_rawDescData)
	})
	return file_matches_matches_proto_rawDescData
}

var file_matches_matches_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_matches_matches_proto_goTypes = []interface{}{
	(*CreateMatchRequest)(nil),  // 0: CreateMatchRequest
	(*UpdateMatchRequest)(nil),  // 1: UpdateMatchRequest
	(*CreateMatchResponse)(nil), // 2: CreateMatchResponse
	(*empty.Empty)(nil),         // 3: google.protobuf.Empty
}
var file_matches_matches_proto_depIdxs = []int32{
	0, // 0: Matches.CreateMatch:input_type -> CreateMatchRequest
	1, // 1: Matches.UpdateMatch:input_type -> UpdateMatchRequest
	2, // 2: Matches.CreateMatch:output_type -> CreateMatchResponse
	3, // 3: Matches.UpdateMatch:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_matches_matches_proto_init() }
func file_matches_matches_proto_init() {
	if File_matches_matches_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_matches_matches_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMatchRequest); i {
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
		file_matches_matches_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMatchRequest); i {
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
		file_matches_matches_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMatchResponse); i {
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
	file_matches_matches_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_matches_matches_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_matches_matches_proto_goTypes,
		DependencyIndexes: file_matches_matches_proto_depIdxs,
		MessageInfos:      file_matches_matches_proto_msgTypes,
	}.Build()
	File_matches_matches_proto = out.File
	file_matches_matches_proto_rawDesc = nil
	file_matches_matches_proto_goTypes = nil
	file_matches_matches_proto_depIdxs = nil
}
