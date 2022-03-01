// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: py_vad/v1/py_vad.proto

package py_vadv1

import (
	v1 "github.com/gitJaesik/stream_go_srvs/streamgolib/gen/proto/go/stream_common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoRequest) Reset() {
	*x = EchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_py_vad_v1_py_vad_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequest) ProtoMessage() {}

func (x *EchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_py_vad_v1_py_vad_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRequest.ProtoReflect.Descriptor instead.
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return file_py_vad_v1_py_vad_proto_rawDescGZIP(), []int{0}
}

func (x *EchoRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoResponse) Reset() {
	*x = EchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_py_vad_v1_py_vad_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResponse) ProtoMessage() {}

func (x *EchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_py_vad_v1_py_vad_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResponse.ProtoReflect.Descriptor instead.
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return file_py_vad_v1_py_vad_proto_rawDescGZIP(), []int{1}
}

func (x *EchoResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CheckVadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AudioMessage *v1.AudioMessage `protobuf:"bytes,1,opt,name=audio_message,json=audioMessage,proto3" json:"audio_message,omitempty"`
}

func (x *CheckVadRequest) Reset() {
	*x = CheckVadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_py_vad_v1_py_vad_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckVadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckVadRequest) ProtoMessage() {}

func (x *CheckVadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_py_vad_v1_py_vad_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckVadRequest.ProtoReflect.Descriptor instead.
func (*CheckVadRequest) Descriptor() ([]byte, []int) {
	return file_py_vad_v1_py_vad_proto_rawDescGZIP(), []int{2}
}

func (x *CheckVadRequest) GetAudioMessage() *v1.AudioMessage {
	if x != nil {
		return x.AudioMessage
	}
	return nil
}

type CheckVadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSpeech bool `protobuf:"varint,1,opt,name=is_speech,json=isSpeech,proto3" json:"is_speech,omitempty"`
}

func (x *CheckVadResponse) Reset() {
	*x = CheckVadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_py_vad_v1_py_vad_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckVadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckVadResponse) ProtoMessage() {}

func (x *CheckVadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_py_vad_v1_py_vad_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckVadResponse.ProtoReflect.Descriptor instead.
func (*CheckVadResponse) Descriptor() ([]byte, []int) {
	return file_py_vad_v1_py_vad_proto_rawDescGZIP(), []int{3}
}

func (x *CheckVadResponse) GetIsSpeech() bool {
	if x != nil {
		return x.IsSpeech
	}
	return false
}

var File_py_vad_v1_py_vad_proto protoreflect.FileDescriptor

var file_py_vad_v1_py_vad_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x79, 0x5f, 0x76, 0x61, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x79, 0x5f, 0x76,
	0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x79, 0x5f, 0x76, 0x61, 0x64,
	0x2e, 0x76, 0x31, 0x1a, 0x24, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x0b, 0x45, 0x63,
	0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x56, 0x0a,
	0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x56, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x43, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0c, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2f, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x56, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f,
	0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73,
	0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x32, 0x90, 0x01, 0x0a, 0x0c, 0x50, 0x79, 0x56, 0x61, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12,
	0x16, 0x2e, 0x70, 0x79, 0x5f, 0x76, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x79, 0x5f, 0x76, 0x61, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x45, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x56, 0x61, 0x64, 0x12, 0x1a,
	0x2e, 0x70, 0x79, 0x5f, 0x76, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x56, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x79, 0x5f,
	0x76, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x56, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xad, 0x01, 0x0a, 0x0d, 0x63, 0x6f,
	0x6d, 0x2e, 0x70, 0x79, 0x5f, 0x76, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x50, 0x79, 0x56,
	0x61, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x4a, 0x61, 0x65, 0x73, 0x69, 0x6b, 0x2f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x67, 0x6f, 0x5f, 0x73, 0x72, 0x76, 0x73, 0x2f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x67, 0x6f, 0x6c, 0x69, 0x62, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x79, 0x5f, 0x76, 0x61, 0x64, 0x2f, 0x76,
	0x31, 0x3b, 0x70, 0x79, 0x5f, 0x76, 0x61, 0x64, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58,
	0xaa, 0x02, 0x08, 0x50, 0x79, 0x56, 0x61, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x50, 0x79,
	0x56, 0x61, 0x64, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x50, 0x79, 0x56, 0x61, 0x64, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09,
	0x50, 0x79, 0x56, 0x61, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_py_vad_v1_py_vad_proto_rawDescOnce sync.Once
	file_py_vad_v1_py_vad_proto_rawDescData = file_py_vad_v1_py_vad_proto_rawDesc
)

func file_py_vad_v1_py_vad_proto_rawDescGZIP() []byte {
	file_py_vad_v1_py_vad_proto_rawDescOnce.Do(func() {
		file_py_vad_v1_py_vad_proto_rawDescData = protoimpl.X.CompressGZIP(file_py_vad_v1_py_vad_proto_rawDescData)
	})
	return file_py_vad_v1_py_vad_proto_rawDescData
}

var file_py_vad_v1_py_vad_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_py_vad_v1_py_vad_proto_goTypes = []interface{}{
	(*EchoRequest)(nil),      // 0: py_vad.v1.EchoRequest
	(*EchoResponse)(nil),     // 1: py_vad.v1.EchoResponse
	(*CheckVadRequest)(nil),  // 2: py_vad.v1.CheckVadRequest
	(*CheckVadResponse)(nil), // 3: py_vad.v1.CheckVadResponse
	(*v1.AudioMessage)(nil),  // 4: stream_common.v1.AudioMessage
}
var file_py_vad_v1_py_vad_proto_depIdxs = []int32{
	4, // 0: py_vad.v1.CheckVadRequest.audio_message:type_name -> stream_common.v1.AudioMessage
	0, // 1: py_vad.v1.PyVadService.Echo:input_type -> py_vad.v1.EchoRequest
	2, // 2: py_vad.v1.PyVadService.CheckVad:input_type -> py_vad.v1.CheckVadRequest
	1, // 3: py_vad.v1.PyVadService.Echo:output_type -> py_vad.v1.EchoResponse
	3, // 4: py_vad.v1.PyVadService.CheckVad:output_type -> py_vad.v1.CheckVadResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_py_vad_v1_py_vad_proto_init() }
func file_py_vad_v1_py_vad_proto_init() {
	if File_py_vad_v1_py_vad_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_py_vad_v1_py_vad_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoRequest); i {
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
		file_py_vad_v1_py_vad_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoResponse); i {
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
		file_py_vad_v1_py_vad_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckVadRequest); i {
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
		file_py_vad_v1_py_vad_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckVadResponse); i {
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
			RawDescriptor: file_py_vad_v1_py_vad_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_py_vad_v1_py_vad_proto_goTypes,
		DependencyIndexes: file_py_vad_v1_py_vad_proto_depIdxs,
		MessageInfos:      file_py_vad_v1_py_vad_proto_msgTypes,
	}.Build()
	File_py_vad_v1_py_vad_proto = out.File
	file_py_vad_v1_py_vad_proto_rawDesc = nil
	file_py_vad_v1_py_vad_proto_goTypes = nil
	file_py_vad_v1_py_vad_proto_depIdxs = nil
}