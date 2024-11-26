// Code generated by kds. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: example2.proto

package examplepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

type City struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId        int64            `protobuf:"varint,1,opt,name=PlayerId,proto3" json:"PlayerId,omitempty"`
	PlayerBasicInfo *PlayerBasicInfo `protobuf:"bytes,2,opt,name=PlayerBasicInfo,proto3" json:"PlayerBasicInfo,omitempty"`
	CityInfo        *CityBaseInfo    `protobuf:"bytes,3,opt,name=CityInfo,proto3" json:"CityInfo,omitempty"`
}

func (x *City) Reset() {
	*x = City{}
	mi := &file_example2_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *City) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*City) ProtoMessage() {}

func (x *City) ProtoReflect() protoreflect.Message {
	mi := &file_example2_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use City.ProtoReflect.Descriptor instead.
func (*City) Descriptor() ([]byte, []int) {
	return file_example2_proto_rawDescGZIP(), []int{0}
}

func (x *City) GetPlayerId() int64 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *City) GetPlayerBasicInfo() *PlayerBasicInfo {
	if x != nil {
		return x.PlayerBasicInfo
	}
	return nil
}

func (x *City) GetCityInfo() *CityBaseInfo {
	if x != nil {
		return x.CityInfo
	}
	return nil
}

type CityBaseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position *Vector `protobuf:"bytes,1,opt,name=Position,proto3" json:"Position,omitempty"`
}

func (x *CityBaseInfo) Reset() {
	*x = CityBaseInfo{}
	mi := &file_example2_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CityBaseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityBaseInfo) ProtoMessage() {}

func (x *CityBaseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_example2_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityBaseInfo.ProtoReflect.Descriptor instead.
func (*CityBaseInfo) Descriptor() ([]byte, []int) {
	return file_example2_proto_rawDescGZIP(), []int{1}
}

func (x *CityBaseInfo) GetPosition() *Vector {
	if x != nil {
		return x.Position
	}
	return nil
}

type Vector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (x *Vector) Reset() {
	*x = Vector{}
	mi := &file_example2_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Vector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vector) ProtoMessage() {}

func (x *Vector) ProtoReflect() protoreflect.Message {
	mi := &file_example2_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vector.ProtoReflect.Descriptor instead.
func (*Vector) Descriptor() ([]byte, []int) {
	return file_example2_proto_rawDescGZIP(), []int{2}
}

func (x *Vector) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vector) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

var File_example2_proto protoreflect.FileDescriptor

var file_example2_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x6b, 0x64, 0x73, 0x1a, 0x0e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x01,
	0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x45, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x6b, 0x64, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42,
	0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a, 0x08, 0x43, 0x69, 0x74,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x6b, 0x64, 0x73, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x42, 0x61, 0x73,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x43, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x3e, 0x0a, 0x0c, 0x43, 0x69, 0x74, 0x79, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x2e, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x6b, 0x64, 0x73, 0x2e, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x24, 0x0a, 0x06, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x59, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x61, 0x6b, 0x75, 0x64, 0x2f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x2f, 0x6b, 0x64, 0x73, 0x2f, 0x6b, 0x64, 0x73, 0x63, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_example2_proto_rawDescOnce sync.Once
	file_example2_proto_rawDescData = file_example2_proto_rawDesc
)

func file_example2_proto_rawDescGZIP() []byte {
	file_example2_proto_rawDescOnce.Do(func() {
		file_example2_proto_rawDescData = protoimpl.X.CompressGZIP(file_example2_proto_rawDescData)
	})
	return file_example2_proto_rawDescData
}

var file_example2_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_example2_proto_goTypes = []any{
	(*City)(nil),            // 0: examplekds.City
	(*CityBaseInfo)(nil),    // 1: examplekds.CityBaseInfo
	(*Vector)(nil),          // 2: examplekds.Vector
	(*PlayerBasicInfo)(nil), // 3: examplekds.PlayerBasicInfo
}
var file_example2_proto_depIdxs = []int32{
	3, // 0: examplekds.City.PlayerBasicInfo:type_name -> examplekds.PlayerBasicInfo
	1, // 1: examplekds.City.CityInfo:type_name -> examplekds.CityBaseInfo
	2, // 2: examplekds.CityBaseInfo.Position:type_name -> examplekds.Vector
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_example2_proto_init() }
func file_example2_proto_init() {
	if File_example2_proto != nil {
		return
	}
	file_example1_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_example2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example2_proto_goTypes,
		DependencyIndexes: file_example2_proto_depIdxs,
		MessageInfos:      file_example2_proto_msgTypes,
	}.Build()
	File_example2_proto = out.File
	file_example2_proto_rawDesc = nil
	file_example2_proto_goTypes = nil
	file_example2_proto_depIdxs = nil
}
