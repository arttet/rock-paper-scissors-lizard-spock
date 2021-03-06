// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: rpsls/game_service/v1/game_service.proto

package game_service

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
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

type GetChoiceV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetChoiceV1Response) Reset() {
	*x = GetChoiceV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpsls_game_service_v1_game_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChoiceV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChoiceV1Response) ProtoMessage() {}

func (x *GetChoiceV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_rpsls_game_service_v1_game_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChoiceV1Response.ProtoReflect.Descriptor instead.
func (*GetChoiceV1Response) Descriptor() ([]byte, []int) {
	return file_rpsls_game_service_v1_game_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetChoiceV1Response) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetChoiceV1Response) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PlayRoundV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player int32 `protobuf:"varint,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *PlayRoundV1Request) Reset() {
	*x = PlayRoundV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpsls_game_service_v1_game_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayRoundV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayRoundV1Request) ProtoMessage() {}

func (x *PlayRoundV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_rpsls_game_service_v1_game_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayRoundV1Request.ProtoReflect.Descriptor instead.
func (*PlayRoundV1Request) Descriptor() ([]byte, []int) {
	return file_rpsls_game_service_v1_game_service_proto_rawDescGZIP(), []int{1}
}

func (x *PlayRoundV1Request) GetPlayer() int32 {
	if x != nil {
		return x.Player
	}
	return 0
}

type PlayRoundV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player   int32  `protobuf:"varint,1,opt,name=player,proto3" json:"player,omitempty"`
	Computer int32  `protobuf:"varint,2,opt,name=computer,proto3" json:"computer,omitempty"`
	Results  string `protobuf:"bytes,3,opt,name=results,proto3" json:"results,omitempty"`
}

func (x *PlayRoundV1Response) Reset() {
	*x = PlayRoundV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpsls_game_service_v1_game_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayRoundV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayRoundV1Response) ProtoMessage() {}

func (x *PlayRoundV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_rpsls_game_service_v1_game_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayRoundV1Response.ProtoReflect.Descriptor instead.
func (*PlayRoundV1Response) Descriptor() ([]byte, []int) {
	return file_rpsls_game_service_v1_game_service_proto_rawDescGZIP(), []int{2}
}

func (x *PlayRoundV1Response) GetPlayer() int32 {
	if x != nil {
		return x.Player
	}
	return 0
}

func (x *PlayRoundV1Response) GetComputer() int32 {
	if x != nil {
		return x.Computer
	}
	return 0
}

func (x *PlayRoundV1Response) GetResults() string {
	if x != nil {
		return x.Results
	}
	return ""
}

var File_rpsls_game_service_v1_game_service_proto protoreflect.FileDescriptor

var file_rpsls_game_service_v1_game_service_proto_rawDesc = []byte{
	0x0a, 0x28, 0x72, 0x70, 0x73, 0x6c, 0x73, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x72, 0x70, 0x73, 0x6c,
	0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70,
	0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x71, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x56, 0x31, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x1a, 0x04, 0x18, 0x05, 0x28, 0x01, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x3f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2b, 0xfa, 0x42, 0x28, 0x72, 0x26, 0x52, 0x04, 0x72, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x70,
	0x61, 0x70, 0x65, 0x72, 0x52, 0x08, 0x73, 0x63, 0x69, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x52, 0x06,
	0x6c, 0x69, 0x7a, 0x61, 0x72, 0x64, 0x52, 0x05, 0x73, 0x70, 0x6f, 0x63, 0x6b, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x6f, 0x75, 0x6e, 0x64,
	0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x06, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x1a, 0x04,
	0x18, 0x05, 0x28, 0x01, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x90, 0x01, 0x0a,
	0x13, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x1a, 0x04, 0x18, 0x05, 0x28, 0x01, 0x52,
	0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x1a, 0x04,
	0x18, 0x05, 0x28, 0x01, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x12, 0x2f,
	0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x15, 0xfa, 0x42, 0x12, 0x72, 0x10, 0x52, 0x03, 0x77, 0x69, 0x6e, 0x52, 0x04, 0x6c, 0x6f, 0x73,
	0x65, 0x52, 0x03, 0x74, 0x69, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x32,
	0xb9, 0x02, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x56, 0x31, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x10, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x12,
	0x62, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x56, 0x31, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2a, 0x2e, 0x72, 0x70, 0x73, 0x6c, 0x73, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x63, 0x68, 0x6f,
	0x69, 0x63, 0x65, 0x12, 0x76, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x6f, 0x75, 0x6e, 0x64,
	0x56, 0x31, 0x12, 0x29, 0x2e, 0x72, 0x70, 0x73, 0x6c, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x52,
	0x6f, 0x75, 0x6e, 0x64, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x72, 0x70, 0x73, 0x6c, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x56,
	0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0a, 0x22, 0x05, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x3a, 0x01, 0x2a, 0x42, 0xd6, 0x01, 0x5a, 0x53,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x74, 0x74, 0x65,
	0x74, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x2d, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2d, 0x73, 0x63, 0x69,
	0x73, 0x73, 0x6f, 0x72, 0x73, 0x2d, 0x6c, 0x69, 0x7a, 0x61, 0x72, 0x64, 0x2d, 0x73, 0x70, 0x6f,
	0x63, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x92, 0x41, 0x7e, 0x12, 0x7c, 0x0a, 0x20, 0x52, 0x6f, 0x63, 0x6b, 0x20, 0x50,
	0x61, 0x70, 0x65, 0x72, 0x20, 0x53, 0x63, 0x69, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x20, 0x4c, 0x69,
	0x7a, 0x61, 0x72, 0x64, 0x20, 0x53, 0x70, 0x6f, 0x63, 0x6b, 0x2a, 0x53, 0x0a, 0x03, 0x4d, 0x49,
	0x54, 0x12, 0x4c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x74, 0x74, 0x65, 0x74, 0x2f, 0x72, 0x6f, 0x63,
	0x6b, 0x2d, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2d, 0x73, 0x63, 0x69, 0x73, 0x73, 0x6f, 0x72, 0x73,
	0x2d, 0x6c, 0x69, 0x7a, 0x61, 0x72, 0x64, 0x2d, 0x73, 0x70, 0x6f, 0x63, 0x6b, 0x2f, 0x62, 0x6c,
	0x6f, 0x62, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x32,
	0x03, 0x31, 0x2e, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpsls_game_service_v1_game_service_proto_rawDescOnce sync.Once
	file_rpsls_game_service_v1_game_service_proto_rawDescData = file_rpsls_game_service_v1_game_service_proto_rawDesc
)

func file_rpsls_game_service_v1_game_service_proto_rawDescGZIP() []byte {
	file_rpsls_game_service_v1_game_service_proto_rawDescOnce.Do(func() {
		file_rpsls_game_service_v1_game_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpsls_game_service_v1_game_service_proto_rawDescData)
	})
	return file_rpsls_game_service_v1_game_service_proto_rawDescData
}

var file_rpsls_game_service_v1_game_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpsls_game_service_v1_game_service_proto_goTypes = []interface{}{
	(*GetChoiceV1Response)(nil), // 0: rpsls.game_service.v1.GetChoiceV1Response
	(*PlayRoundV1Request)(nil),  // 1: rpsls.game_service.v1.PlayRoundV1Request
	(*PlayRoundV1Response)(nil), // 2: rpsls.game_service.v1.PlayRoundV1Response
	(*emptypb.Empty)(nil),       // 3: google.protobuf.Empty
	(*httpbody.HttpBody)(nil),   // 4: google.api.HttpBody
}
var file_rpsls_game_service_v1_game_service_proto_depIdxs = []int32{
	3, // 0: rpsls.game_service.v1.GameService.GetChoicesV1:input_type -> google.protobuf.Empty
	3, // 1: rpsls.game_service.v1.GameService.GetChoiceV1:input_type -> google.protobuf.Empty
	1, // 2: rpsls.game_service.v1.GameService.PlayRoundV1:input_type -> rpsls.game_service.v1.PlayRoundV1Request
	4, // 3: rpsls.game_service.v1.GameService.GetChoicesV1:output_type -> google.api.HttpBody
	0, // 4: rpsls.game_service.v1.GameService.GetChoiceV1:output_type -> rpsls.game_service.v1.GetChoiceV1Response
	2, // 5: rpsls.game_service.v1.GameService.PlayRoundV1:output_type -> rpsls.game_service.v1.PlayRoundV1Response
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpsls_game_service_v1_game_service_proto_init() }
func file_rpsls_game_service_v1_game_service_proto_init() {
	if File_rpsls_game_service_v1_game_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpsls_game_service_v1_game_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChoiceV1Response); i {
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
		file_rpsls_game_service_v1_game_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayRoundV1Request); i {
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
		file_rpsls_game_service_v1_game_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayRoundV1Response); i {
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
			RawDescriptor: file_rpsls_game_service_v1_game_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpsls_game_service_v1_game_service_proto_goTypes,
		DependencyIndexes: file_rpsls_game_service_v1_game_service_proto_depIdxs,
		MessageInfos:      file_rpsls_game_service_v1_game_service_proto_msgTypes,
	}.Build()
	File_rpsls_game_service_v1_game_service_proto = out.File
	file_rpsls_game_service_v1_game_service_proto_rawDesc = nil
	file_rpsls_game_service_v1_game_service_proto_goTypes = nil
	file_rpsls_game_service_v1_game_service_proto_depIdxs = nil
}
