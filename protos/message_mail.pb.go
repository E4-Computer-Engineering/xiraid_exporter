//*
// Messages to manage the mail notification.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.2
// source: message_mail.proto

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
// Add the recipient's email and notification level.
type MailAdd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Receiver's email.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// *
	// The notification level.
	Level string `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *MailAdd) Reset() {
	*x = MailAdd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_mail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailAdd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailAdd) ProtoMessage() {}

func (x *MailAdd) ProtoReflect() protoreflect.Message {
	mi := &file_message_mail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailAdd.ProtoReflect.Descriptor instead.
func (*MailAdd) Descriptor() ([]byte, []int) {
	return file_message_mail_proto_rawDescGZIP(), []int{0}
}

func (x *MailAdd) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MailAdd) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

// *
// Remove the email from the list of notification recipients.
type MailRemove struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// The email address to remove from the list of notification recipients.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *MailRemove) Reset() {
	*x = MailRemove{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_mail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailRemove) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailRemove) ProtoMessage() {}

func (x *MailRemove) ProtoReflect() protoreflect.Message {
	mi := &file_message_mail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailRemove.ProtoReflect.Descriptor instead.
func (*MailRemove) Descriptor() ([]byte, []int) {
	return file_message_mail_proto_rawDescGZIP(), []int{1}
}

func (x *MailRemove) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// *
// Show the list of email notification recipients.
type MailShow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MailShow) Reset() {
	*x = MailShow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_mail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailShow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailShow) ProtoMessage() {}

func (x *MailShow) ProtoReflect() protoreflect.Message {
	mi := &file_message_mail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailShow.ProtoReflect.Descriptor instead.
func (*MailShow) Descriptor() ([]byte, []int) {
	return file_message_mail_proto_rawDescGZIP(), []int{2}
}

var File_message_mail_proto protoreflect.FileDescriptor

var file_message_mail_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x78, 0x72, 0x61, 0x69, 0x64, 0x2e, 0x76, 0x32, 0x22, 0x39,
	0x0a, 0x07, 0x4d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x26, 0x0a, 0x0a, 0x4d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x0a, 0x0a, 0x08, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x68, 0x6f, 0x77, 0x42, 0x25, 0x5a,
	0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x72, 0x6f, 0x6e,
	0x63, 0x75, 0x62, 0x33, 0x2f, 0x78, 0x69, 0x72, 0x61, 0x69, 0x64, 0x5f, 0x65, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_mail_proto_rawDescOnce sync.Once
	file_message_mail_proto_rawDescData = file_message_mail_proto_rawDesc
)

func file_message_mail_proto_rawDescGZIP() []byte {
	file_message_mail_proto_rawDescOnce.Do(func() {
		file_message_mail_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_mail_proto_rawDescData)
	})
	return file_message_mail_proto_rawDescData
}

var file_message_mail_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_message_mail_proto_goTypes = []interface{}{
	(*MailAdd)(nil),    // 0: xraid.v2.MailAdd
	(*MailRemove)(nil), // 1: xraid.v2.MailRemove
	(*MailShow)(nil),   // 2: xraid.v2.MailShow
}
var file_message_mail_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_mail_proto_init() }
func file_message_mail_proto_init() {
	if File_message_mail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_mail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailAdd); i {
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
		file_message_mail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailRemove); i {
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
		file_message_mail_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailShow); i {
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
			RawDescriptor: file_message_mail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_mail_proto_goTypes,
		DependencyIndexes: file_message_mail_proto_depIdxs,
		MessageInfos:      file_message_mail_proto_msgTypes,
	}.Build()
	File_message_mail_proto = out.File
	file_message_mail_proto_rawDesc = nil
	file_message_mail_proto_goTypes = nil
	file_message_mail_proto_depIdxs = nil
}
