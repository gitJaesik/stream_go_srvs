// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: stream_common/v1/stream_common.proto

package stream_commonv1

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

type ClientInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceType  string `protobuf:"bytes,1,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	DeviceToken string `protobuf:"bytes,2,opt,name=device_token,json=deviceToken,proto3" json:"device_token,omitempty"`
	Location    string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *ClientInfo) Reset() {
	*x = ClientInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_common_v1_stream_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInfo) ProtoMessage() {}

func (x *ClientInfo) ProtoReflect() protoreflect.Message {
	mi := &file_stream_common_v1_stream_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInfo.ProtoReflect.Descriptor instead.
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return file_stream_common_v1_stream_common_proto_rawDescGZIP(), []int{0}
}

func (x *ClientInfo) GetDeviceType() string {
	if x != nil {
		return x.DeviceType
	}
	return ""
}

func (x *ClientInfo) GetDeviceToken() string {
	if x != nil {
		return x.DeviceToken
	}
	return ""
}

func (x *ClientInfo) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type AudioMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sample_rate * bit_depth * frame_duration / frame_number
	AudioData []byte `protobuf:"bytes,1,opt,name=audio_data,json=audioData,proto3" json:"audio_data,omitempty"`
	// pcm, wav, mp3
	AudioFormat string `protobuf:"bytes,2,opt,name=audio_format,json=audioFormat,proto3" json:"audio_format,omitempty"`
	// 16000
	SampleRate int32 `protobuf:"varint,3,opt,name=sample_rate,json=sampleRate,proto3" json:"sample_rate,omitempty"`
	// 16
	BitDepth int32 `protobuf:"varint,4,opt,name=bit_depth,json=bitDepth,proto3" json:"bit_depth,omitempty"`
	// 10ms, 20ms, 30ms
	FrameDuration int32 `protobuf:"varint,5,opt,name=frame_duration,json=frameDuration,proto3" json:"frame_duration,omitempty"`
	// audio_data / sample_rate / bit_depth / frame_duration
	FrameNumber int32 `protobuf:"varint,6,opt,name=frame_number,json=frameNumber,proto3" json:"frame_number,omitempty"`
}

func (x *AudioMessage) Reset() {
	*x = AudioMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_common_v1_stream_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioMessage) ProtoMessage() {}

func (x *AudioMessage) ProtoReflect() protoreflect.Message {
	mi := &file_stream_common_v1_stream_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioMessage.ProtoReflect.Descriptor instead.
func (*AudioMessage) Descriptor() ([]byte, []int) {
	return file_stream_common_v1_stream_common_proto_rawDescGZIP(), []int{1}
}

func (x *AudioMessage) GetAudioData() []byte {
	if x != nil {
		return x.AudioData
	}
	return nil
}

func (x *AudioMessage) GetAudioFormat() string {
	if x != nil {
		return x.AudioFormat
	}
	return ""
}

func (x *AudioMessage) GetSampleRate() int32 {
	if x != nil {
		return x.SampleRate
	}
	return 0
}

func (x *AudioMessage) GetBitDepth() int32 {
	if x != nil {
		return x.BitDepth
	}
	return 0
}

func (x *AudioMessage) GetFrameDuration() int32 {
	if x != nil {
		return x.FrameDuration
	}
	return 0
}

func (x *AudioMessage) GetFrameNumber() int32 {
	if x != nil {
		return x.FrameNumber
	}
	return 0
}

var File_stream_common_v1_stream_common_proto protoreflect.FileDescriptor

var file_stream_common_v1_stream_common_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x6c, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd8, 0x01, 0x0a, 0x0c, 0x41, 0x75, 0x64, 0x69, 0x6f,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x6f,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x61, 0x75, 0x64,
	0x69, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x75,
	0x64, 0x69, 0x6f, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69,
	0x74, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x62,
	0x69, 0x74, 0x44, 0x65, 0x70, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21,
	0x0a, 0x0c, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x42, 0xe5, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x5d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x4a,
	0x61, 0x65, 0x73, 0x69, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x67, 0x6f, 0x5f,
	0x73, 0x72, 0x76, 0x73, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x67, 0x6f, 0x6c, 0x69, 0x62,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x0f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_stream_common_v1_stream_common_proto_rawDescOnce sync.Once
	file_stream_common_v1_stream_common_proto_rawDescData = file_stream_common_v1_stream_common_proto_rawDesc
)

func file_stream_common_v1_stream_common_proto_rawDescGZIP() []byte {
	file_stream_common_v1_stream_common_proto_rawDescOnce.Do(func() {
		file_stream_common_v1_stream_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_stream_common_v1_stream_common_proto_rawDescData)
	})
	return file_stream_common_v1_stream_common_proto_rawDescData
}

var file_stream_common_v1_stream_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stream_common_v1_stream_common_proto_goTypes = []interface{}{
	(*ClientInfo)(nil),   // 0: stream_common.v1.ClientInfo
	(*AudioMessage)(nil), // 1: stream_common.v1.AudioMessage
}
var file_stream_common_v1_stream_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stream_common_v1_stream_common_proto_init() }
func file_stream_common_v1_stream_common_proto_init() {
	if File_stream_common_v1_stream_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stream_common_v1_stream_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientInfo); i {
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
		file_stream_common_v1_stream_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudioMessage); i {
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
			RawDescriptor: file_stream_common_v1_stream_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stream_common_v1_stream_common_proto_goTypes,
		DependencyIndexes: file_stream_common_v1_stream_common_proto_depIdxs,
		MessageInfos:      file_stream_common_v1_stream_common_proto_msgTypes,
	}.Build()
	File_stream_common_v1_stream_common_proto = out.File
	file_stream_common_v1_stream_common_proto_rawDesc = nil
	file_stream_common_v1_stream_common_proto_goTypes = nil
	file_stream_common_v1_stream_common_proto_depIdxs = nil
}
