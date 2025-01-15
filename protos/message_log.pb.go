//*
// Messages to manage the event log operations.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.2
// source: message_log.proto

package xiraid_exporter

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

// *
// Collect the event log entries into a .tar.gz file.
type LogCollect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogCollect) Reset() {
	*x = LogCollect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogCollect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogCollect) ProtoMessage() {}

func (x *LogCollect) ProtoReflect() protoreflect.Message {
	mi := &file_message_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogCollect.ProtoReflect.Descriptor instead.
func (*LogCollect) Descriptor() ([]byte, []int) {
	return file_message_log_proto_rawDescGZIP(), []int{0}
}

// *
// Show the last entries in the event log.
type LogShow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// The number of entries in the event log, starting from the last entry.
	Lines *uint32 `protobuf:"varint,1,opt,name=lines,proto3,oneof" json:"lines,omitempty"`
}

func (x *LogShow) Reset() {
	*x = LogShow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogShow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogShow) ProtoMessage() {}

func (x *LogShow) ProtoReflect() protoreflect.Message {
	mi := &file_message_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogShow.ProtoReflect.Descriptor instead.
func (*LogShow) Descriptor() ([]byte, []int) {
	return file_message_log_proto_rawDescGZIP(), []int{1}
}

func (x *LogShow) GetLines() uint32 {
	if x != nil && x.Lines != nil {
		return *x.Lines
	}
	return 0
}

var File_message_log_proto protoreflect.FileDescriptor

var file_message_log_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x78, 0x72, 0x61, 0x69, 0x64, 0x2e, 0x76, 0x32, 0x22, 0x0c, 0x0a,
	0x0a, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x22, 0x2e, 0x0a, 0x07, 0x4c,
	0x6f, 0x67, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x88, 0x01,
	0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x42, 0x25, 0x5a, 0x23, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x72, 0x6f, 0x6e, 0x63, 0x75,
	0x62, 0x33, 0x2f, 0x78, 0x69, 0x72, 0x61, 0x69, 0x64, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_log_proto_rawDescOnce sync.Once
	file_message_log_proto_rawDescData = file_message_log_proto_rawDesc
)

func file_message_log_proto_rawDescGZIP() []byte {
	file_message_log_proto_rawDescOnce.Do(func() {
		file_message_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_log_proto_rawDescData)
	})
	return file_message_log_proto_rawDescData
}

var file_message_log_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_message_log_proto_goTypes = []interface{}{
	(*LogCollect)(nil), // 0: xraid.v2.LogCollect
	(*LogShow)(nil),    // 1: xraid.v2.LogShow
}
var file_message_log_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_log_proto_init() }
func file_message_log_proto_init() {
	if File_message_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogCollect); i {
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
		file_message_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogShow); i {
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
	file_message_log_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_log_proto_goTypes,
		DependencyIndexes: file_message_log_proto_depIdxs,
		MessageInfos:      file_message_log_proto_msgTypes,
	}.Build()
	File_message_log_proto = out.File
	file_message_log_proto_rawDesc = nil
	file_message_log_proto_goTypes = nil
	file_message_log_proto_depIdxs = nil
}
