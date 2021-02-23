// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: proto/kafka.proto

package bridge

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ConsumeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *ConsumeRequest) Reset() {
	*x = ConsumeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kafka_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeRequest) ProtoMessage() {}

func (x *ConsumeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kafka_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeRequest.ProtoReflect.Descriptor instead.
func (*ConsumeRequest) Descriptor() ([]byte, []int) {
	return file_proto_kafka_proto_rawDescGZIP(), []int{0}
}

func (x *ConsumeRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

// Used to publish requests to the `Subscribe` and `Produce` rpc
type PublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	// Types that are assignable to OptionalContent:
	//	*PublishRequest_Content
	OptionalContent isPublishRequest_OptionalContent `protobuf_oneof:"optional_content"`
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kafka_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kafka_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_proto_kafka_proto_rawDescGZIP(), []int{1}
}

func (x *PublishRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (m *PublishRequest) GetOptionalContent() isPublishRequest_OptionalContent {
	if m != nil {
		return m.OptionalContent
	}
	return nil
}

func (x *PublishRequest) GetContent() []byte {
	if x, ok := x.GetOptionalContent().(*PublishRequest_Content); ok {
		return x.Content
	}
	return nil
}

type isPublishRequest_OptionalContent interface {
	isPublishRequest_OptionalContent()
}

type PublishRequest_Content struct {
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3,oneof"`
}

func (*PublishRequest_Content) isPublishRequest_OptionalContent() {}

// Used to respond to `Subscribe` and `Consume` rpcs
// - Optional content to indicate data from a kafka broker
type KafkaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// Types that are assignable to OptionalContent:
	//	*KafkaResponse_Content
	OptionalContent isKafkaResponse_OptionalContent `protobuf_oneof:"optional_content"`
}

func (x *KafkaResponse) Reset() {
	*x = KafkaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kafka_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KafkaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KafkaResponse) ProtoMessage() {}

func (x *KafkaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kafka_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KafkaResponse.ProtoReflect.Descriptor instead.
func (*KafkaResponse) Descriptor() ([]byte, []int) {
	return file_proto_kafka_proto_rawDescGZIP(), []int{2}
}

func (x *KafkaResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (m *KafkaResponse) GetOptionalContent() isKafkaResponse_OptionalContent {
	if m != nil {
		return m.OptionalContent
	}
	return nil
}

func (x *KafkaResponse) GetContent() []byte {
	if x, ok := x.GetOptionalContent().(*KafkaResponse_Content); ok {
		return x.Content
	}
	return nil
}

type isKafkaResponse_OptionalContent interface {
	isKafkaResponse_OptionalContent()
}

type KafkaResponse_Content struct {
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3,oneof"`
}

func (*KafkaResponse_Content) isKafkaResponse_OptionalContent() {}

// Used to respond to `Produce` rpcs
// - Optional content to indicate error text
type ProduceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// Types that are assignable to Message:
	//	*ProduceResponse_Content
	Message isProduceResponse_Message `protobuf_oneof:"message"`
}

func (x *ProduceResponse) Reset() {
	*x = ProduceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kafka_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProduceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProduceResponse) ProtoMessage() {}

func (x *ProduceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kafka_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProduceResponse.ProtoReflect.Descriptor instead.
func (*ProduceResponse) Descriptor() ([]byte, []int) {
	return file_proto_kafka_proto_rawDescGZIP(), []int{3}
}

func (x *ProduceResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (m *ProduceResponse) GetMessage() isProduceResponse_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *ProduceResponse) GetContent() string {
	if x, ok := x.GetMessage().(*ProduceResponse_Content); ok {
		return x.Content
	}
	return ""
}

type isProduceResponse_Message interface {
	isProduceResponse_Message()
}

type ProduceResponse_Content struct {
	Content string `protobuf:"bytes,2,opt,name=content,proto3,oneof"`
}

func (*ProduceResponse_Content) isProduceResponse_Message() {}

var File_proto_kafka_proto protoreflect.FileDescriptor

var file_proto_kafka_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x22, 0x26, 0x0a, 0x0e, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x22, 0x56, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1a, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x12, 0x0a, 0x10, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x59, 0x0a, 0x0d, 0x4b,
	0x61, 0x66, 0x6b, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x42, 0x12, 0x0a, 0x10, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x52, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42,
	0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xcf, 0x01, 0x0a, 0x0b, 0x4b,
	0x61, 0x66, 0x6b, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x40, 0x0a, 0x09, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x16, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65,
	0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x3c, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65,
	0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x40, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x65, 0x12, 0x16, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_kafka_proto_rawDescOnce sync.Once
	file_proto_kafka_proto_rawDescData = file_proto_kafka_proto_rawDesc
)

func file_proto_kafka_proto_rawDescGZIP() []byte {
	file_proto_kafka_proto_rawDescOnce.Do(func() {
		file_proto_kafka_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_kafka_proto_rawDescData)
	})
	return file_proto_kafka_proto_rawDescData
}

var file_proto_kafka_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_kafka_proto_goTypes = []interface{}{
	(*ConsumeRequest)(nil),  // 0: bridge.ConsumeRequest
	(*PublishRequest)(nil),  // 1: bridge.PublishRequest
	(*KafkaResponse)(nil),   // 2: bridge.KafkaResponse
	(*ProduceResponse)(nil), // 3: bridge.ProduceResponse
}
var file_proto_kafka_proto_depIdxs = []int32{
	1, // 0: bridge.KafkaStream.Subscribe:input_type -> bridge.PublishRequest
	0, // 1: bridge.KafkaStream.Consume:input_type -> bridge.ConsumeRequest
	1, // 2: bridge.KafkaStream.Produce:input_type -> bridge.PublishRequest
	2, // 3: bridge.KafkaStream.Subscribe:output_type -> bridge.KafkaResponse
	2, // 4: bridge.KafkaStream.Consume:output_type -> bridge.KafkaResponse
	3, // 5: bridge.KafkaStream.Produce:output_type -> bridge.ProduceResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_kafka_proto_init() }
func file_proto_kafka_proto_init() {
	if File_proto_kafka_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_kafka_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumeRequest); i {
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
		file_proto_kafka_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRequest); i {
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
		file_proto_kafka_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KafkaResponse); i {
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
		file_proto_kafka_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProduceResponse); i {
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
	file_proto_kafka_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*PublishRequest_Content)(nil),
	}
	file_proto_kafka_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*KafkaResponse_Content)(nil),
	}
	file_proto_kafka_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ProduceResponse_Content)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_kafka_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_kafka_proto_goTypes,
		DependencyIndexes: file_proto_kafka_proto_depIdxs,
		MessageInfos:      file_proto_kafka_proto_msgTypes,
	}.Build()
	File_proto_kafka_proto = out.File
	file_proto_kafka_proto_rawDesc = nil
	file_proto_kafka_proto_goTypes = nil
	file_proto_kafka_proto_depIdxs = nil
}
