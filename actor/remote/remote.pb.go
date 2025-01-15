// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: remote/remote.proto

package remote

import (
	actor "github.com/iakud/knoll/actor"
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

type Envelope struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Target        *actor.PID             `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Sender        *actor.PID             `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Message       []byte                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	mi := &file_remote_remote_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_remote_remote_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_remote_remote_proto_rawDescGZIP(), []int{0}
}

func (x *Envelope) GetTarget() *actor.PID {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *Envelope) GetSender() *actor.PID {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *Envelope) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_remote_remote_proto protoreflect.FileDescriptor

var file_remote_remote_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x1a, 0x0b, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x08, 0x45, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50,
	0x49, 0x44, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x22, 0x0a, 0x06, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x50, 0x49, 0x44, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x3d, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x12, 0x10, 0x2e,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x1a,
	0x10, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x61, 0x6b, 0x75, 0x64, 0x2f, 0x6b, 0x6e, 0x6f, 0x6c,
	0x6c, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_remote_remote_proto_rawDescOnce sync.Once
	file_remote_remote_proto_rawDescData = file_remote_remote_proto_rawDesc
)

func file_remote_remote_proto_rawDescGZIP() []byte {
	file_remote_remote_proto_rawDescOnce.Do(func() {
		file_remote_remote_proto_rawDescData = protoimpl.X.CompressGZIP(file_remote_remote_proto_rawDescData)
	})
	return file_remote_remote_proto_rawDescData
}

var file_remote_remote_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_remote_remote_proto_goTypes = []any{
	(*Envelope)(nil),  // 0: remote.Envelope
	(*actor.PID)(nil), // 1: actor.PID
}
var file_remote_remote_proto_depIdxs = []int32{
	1, // 0: remote.Envelope.target:type_name -> actor.PID
	1, // 1: remote.Envelope.sender:type_name -> actor.PID
	0, // 2: remote.Remote.Receive:input_type -> remote.Envelope
	0, // 3: remote.Remote.Receive:output_type -> remote.Envelope
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_remote_remote_proto_init() }
func file_remote_remote_proto_init() {
	if File_remote_remote_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_remote_remote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_remote_remote_proto_goTypes,
		DependencyIndexes: file_remote_remote_proto_depIdxs,
		MessageInfos:      file_remote_remote_proto_msgTypes,
	}.Build()
	File_remote_remote_proto = out.File
	file_remote_remote_proto_rawDesc = nil
	file_remote_remote_proto_goTypes = nil
	file_remote_remote_proto_depIdxs = nil
}