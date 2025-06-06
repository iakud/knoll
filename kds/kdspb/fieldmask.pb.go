// Code generated by kds. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: fieldmask.proto

package kdspb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FieldMask struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Number        int32                  `protobuf:"varint,1,opt,name=Number,proto3" json:"Number,omitempty"`
	MapMask       *MapMask               `protobuf:"bytes,2,opt,name=MapMask,proto3" json:"MapMask,omitempty"`
	FieldMasks    []*FieldMask           `protobuf:"bytes,10,rep,name=FieldMasks,proto3" json:"FieldMasks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FieldMask) Reset() {
	*x = FieldMask{}
	mi := &file_fieldmask_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FieldMask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldMask) ProtoMessage() {}

func (x *FieldMask) ProtoReflect() protoreflect.Message {
	mi := &file_fieldmask_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldMask.ProtoReflect.Descriptor instead.
func (*FieldMask) Descriptor() ([]byte, []int) {
	return file_fieldmask_proto_rawDescGZIP(), []int{0}
}

func (x *FieldMask) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *FieldMask) GetMapMask() *MapMask {
	if x != nil {
		return x.MapMask
	}
	return nil
}

func (x *FieldMask) GetFieldMasks() []*FieldMask {
	if x != nil {
		return x.FieldMasks
	}
	return nil
}

type Int32Array struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Values        []int32                `protobuf:"varint,1,rep,packed,name=Values,proto3" json:"Values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Int32Array) Reset() {
	*x = Int32Array{}
	mi := &file_fieldmask_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Int32Array) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int32Array) ProtoMessage() {}

func (x *Int32Array) ProtoReflect() protoreflect.Message {
	mi := &file_fieldmask_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int32Array.ProtoReflect.Descriptor instead.
func (*Int32Array) Descriptor() ([]byte, []int) {
	return file_fieldmask_proto_rawDescGZIP(), []int{1}
}

func (x *Int32Array) GetValues() []int32 {
	if x != nil {
		return x.Values
	}
	return nil
}

type Int64Array struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Values        []int64                `protobuf:"varint,1,rep,packed,name=Values,proto3" json:"Values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Int64Array) Reset() {
	*x = Int64Array{}
	mi := &file_fieldmask_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Int64Array) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int64Array) ProtoMessage() {}

func (x *Int64Array) ProtoReflect() protoreflect.Message {
	mi := &file_fieldmask_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int64Array.ProtoReflect.Descriptor instead.
func (*Int64Array) Descriptor() ([]byte, []int) {
	return file_fieldmask_proto_rawDescGZIP(), []int{2}
}

func (x *Int64Array) GetValues() []int64 {
	if x != nil {
		return x.Values
	}
	return nil
}

type UInt32Array struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Values        []uint32               `protobuf:"varint,1,rep,packed,name=Values,proto3" json:"Values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UInt32Array) Reset() {
	*x = UInt32Array{}
	mi := &file_fieldmask_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UInt32Array) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UInt32Array) ProtoMessage() {}

func (x *UInt32Array) ProtoReflect() protoreflect.Message {
	mi := &file_fieldmask_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UInt32Array.ProtoReflect.Descriptor instead.
func (*UInt32Array) Descriptor() ([]byte, []int) {
	return file_fieldmask_proto_rawDescGZIP(), []int{3}
}

func (x *UInt32Array) GetValues() []uint32 {
	if x != nil {
		return x.Values
	}
	return nil
}

type Uint64Array struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Values        []uint64               `protobuf:"varint,1,rep,packed,name=Values,proto3" json:"Values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Uint64Array) Reset() {
	*x = Uint64Array{}
	mi := &file_fieldmask_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Uint64Array) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uint64Array) ProtoMessage() {}

func (x *Uint64Array) ProtoReflect() protoreflect.Message {
	mi := &file_fieldmask_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Uint64Array.ProtoReflect.Descriptor instead.
func (*Uint64Array) Descriptor() ([]byte, []int) {
	return file_fieldmask_proto_rawDescGZIP(), []int{4}
}

func (x *Uint64Array) GetValues() []uint64 {
	if x != nil {
		return x.Values
	}
	return nil
}

type BoolArray struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Values        []bool                 `protobuf:"varint,1,rep,packed,name=Values,proto3" json:"Values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BoolArray) Reset() {
	*x = BoolArray{}
	mi := &file_fieldmask_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BoolArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolArray) ProtoMessage() {}

func (x *BoolArray) ProtoReflect() protoreflect.Message {
	mi := &file_fieldmask_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolArray.ProtoReflect.Descriptor instead.
func (*BoolArray) Descriptor() ([]byte, []int) {
	return file_fieldmask_proto_rawDescGZIP(), []int{5}
}

func (x *BoolArray) GetValues() []bool {
	if x != nil {
		return x.Values
	}
	return nil
}

type StringArray struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Values        []string               `protobuf:"bytes,1,rep,name=Values,proto3" json:"Values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StringArray) Reset() {
	*x = StringArray{}
	mi := &file_fieldmask_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StringArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringArray) ProtoMessage() {}

func (x *StringArray) ProtoReflect() protoreflect.Message {
	mi := &file_fieldmask_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringArray.ProtoReflect.Descriptor instead.
func (*StringArray) Descriptor() ([]byte, []int) {
	return file_fieldmask_proto_rawDescGZIP(), []int{6}
}

func (x *StringArray) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type MapMask struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Clear bool                   `protobuf:"varint,1,opt,name=Clear,proto3" json:"Clear,omitempty"`
	// Types that are valid to be assigned to DeleteKeys:
	//
	//	*MapMask_Int32DeleteKeys
	//	*MapMask_Int64DeleteKeys
	//	*MapMask_Uint32DeleteKeys
	//	*MapMask_Uint64DeleteKeys
	//	*MapMask_BoolDeleteKeys
	//	*MapMask_StringDeleteKeys
	DeleteKeys    isMapMask_DeleteKeys `protobuf_oneof:"DeleteKeys"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MapMask) Reset() {
	*x = MapMask{}
	mi := &file_fieldmask_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapMask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapMask) ProtoMessage() {}

func (x *MapMask) ProtoReflect() protoreflect.Message {
	mi := &file_fieldmask_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapMask.ProtoReflect.Descriptor instead.
func (*MapMask) Descriptor() ([]byte, []int) {
	return file_fieldmask_proto_rawDescGZIP(), []int{7}
}

func (x *MapMask) GetClear() bool {
	if x != nil {
		return x.Clear
	}
	return false
}

func (x *MapMask) GetDeleteKeys() isMapMask_DeleteKeys {
	if x != nil {
		return x.DeleteKeys
	}
	return nil
}

func (x *MapMask) GetInt32DeleteKeys() *Int32Array {
	if x != nil {
		if x, ok := x.DeleteKeys.(*MapMask_Int32DeleteKeys); ok {
			return x.Int32DeleteKeys
		}
	}
	return nil
}

func (x *MapMask) GetInt64DeleteKeys() *Int64Array {
	if x != nil {
		if x, ok := x.DeleteKeys.(*MapMask_Int64DeleteKeys); ok {
			return x.Int64DeleteKeys
		}
	}
	return nil
}

func (x *MapMask) GetUint32DeleteKeys() *UInt32Array {
	if x != nil {
		if x, ok := x.DeleteKeys.(*MapMask_Uint32DeleteKeys); ok {
			return x.Uint32DeleteKeys
		}
	}
	return nil
}

func (x *MapMask) GetUint64DeleteKeys() *Uint64Array {
	if x != nil {
		if x, ok := x.DeleteKeys.(*MapMask_Uint64DeleteKeys); ok {
			return x.Uint64DeleteKeys
		}
	}
	return nil
}

func (x *MapMask) GetBoolDeleteKeys() *BoolArray {
	if x != nil {
		if x, ok := x.DeleteKeys.(*MapMask_BoolDeleteKeys); ok {
			return x.BoolDeleteKeys
		}
	}
	return nil
}

func (x *MapMask) GetStringDeleteKeys() *StringArray {
	if x != nil {
		if x, ok := x.DeleteKeys.(*MapMask_StringDeleteKeys); ok {
			return x.StringDeleteKeys
		}
	}
	return nil
}

type isMapMask_DeleteKeys interface {
	isMapMask_DeleteKeys()
}

type MapMask_Int32DeleteKeys struct {
	Int32DeleteKeys *Int32Array `protobuf:"bytes,2,opt,name=Int32DeleteKeys,proto3,oneof"`
}

type MapMask_Int64DeleteKeys struct {
	Int64DeleteKeys *Int64Array `protobuf:"bytes,3,opt,name=Int64DeleteKeys,proto3,oneof"`
}

type MapMask_Uint32DeleteKeys struct {
	Uint32DeleteKeys *UInt32Array `protobuf:"bytes,4,opt,name=Uint32DeleteKeys,proto3,oneof"`
}

type MapMask_Uint64DeleteKeys struct {
	Uint64DeleteKeys *Uint64Array `protobuf:"bytes,5,opt,name=Uint64DeleteKeys,proto3,oneof"`
}

type MapMask_BoolDeleteKeys struct {
	BoolDeleteKeys *BoolArray `protobuf:"bytes,6,opt,name=BoolDeleteKeys,proto3,oneof"`
}

type MapMask_StringDeleteKeys struct {
	StringDeleteKeys *StringArray `protobuf:"bytes,7,opt,name=StringDeleteKeys,proto3,oneof"`
}

func (*MapMask_Int32DeleteKeys) isMapMask_DeleteKeys() {}

func (*MapMask_Int64DeleteKeys) isMapMask_DeleteKeys() {}

func (*MapMask_Uint32DeleteKeys) isMapMask_DeleteKeys() {}

func (*MapMask_Uint64DeleteKeys) isMapMask_DeleteKeys() {}

func (*MapMask_BoolDeleteKeys) isMapMask_DeleteKeys() {}

func (*MapMask_StringDeleteKeys) isMapMask_DeleteKeys() {}

var File_fieldmask_proto protoreflect.FileDescriptor

const file_fieldmask_proto_rawDesc = "" +
	"\n" +
	"\x0ffieldmask.proto\x12\x05kdspb\"\x7f\n" +
	"\tFieldMask\x12\x16\n" +
	"\x06Number\x18\x01 \x01(\x05R\x06Number\x12(\n" +
	"\aMapMask\x18\x02 \x01(\v2\x0e.kdspb.MapMaskR\aMapMask\x120\n" +
	"\n" +
	"FieldMasks\x18\n" +
	" \x03(\v2\x10.kdspb.FieldMaskR\n" +
	"FieldMasks\"$\n" +
	"\n" +
	"Int32Array\x12\x16\n" +
	"\x06Values\x18\x01 \x03(\x05R\x06Values\"$\n" +
	"\n" +
	"Int64Array\x12\x16\n" +
	"\x06Values\x18\x01 \x03(\x03R\x06Values\"%\n" +
	"\vUInt32Array\x12\x16\n" +
	"\x06Values\x18\x01 \x03(\rR\x06Values\"%\n" +
	"\vUint64Array\x12\x16\n" +
	"\x06Values\x18\x01 \x03(\x04R\x06Values\"#\n" +
	"\tBoolArray\x12\x16\n" +
	"\x06Values\x18\x01 \x03(\bR\x06Values\"%\n" +
	"\vStringArray\x12\x16\n" +
	"\x06Values\x18\x01 \x03(\tR\x06Values\"\xad\x03\n" +
	"\aMapMask\x12\x14\n" +
	"\x05Clear\x18\x01 \x01(\bR\x05Clear\x12=\n" +
	"\x0fInt32DeleteKeys\x18\x02 \x01(\v2\x11.kdspb.Int32ArrayH\x00R\x0fInt32DeleteKeys\x12=\n" +
	"\x0fInt64DeleteKeys\x18\x03 \x01(\v2\x11.kdspb.Int64ArrayH\x00R\x0fInt64DeleteKeys\x12@\n" +
	"\x10Uint32DeleteKeys\x18\x04 \x01(\v2\x12.kdspb.UInt32ArrayH\x00R\x10Uint32DeleteKeys\x12@\n" +
	"\x10Uint64DeleteKeys\x18\x05 \x01(\v2\x12.kdspb.Uint64ArrayH\x00R\x10Uint64DeleteKeys\x12:\n" +
	"\x0eBoolDeleteKeys\x18\x06 \x01(\v2\x10.kdspb.BoolArrayH\x00R\x0eBoolDeleteKeys\x12@\n" +
	"\x10StringDeleteKeys\x18\a \x01(\v2\x12.kdspb.StringArrayH\x00R\x10StringDeleteKeysB\f\n" +
	"\n" +
	"DeleteKeysB\x1cZ\x1agithub.com/iakud/kds/kdspbb\x06proto3"

var (
	file_fieldmask_proto_rawDescOnce sync.Once
	file_fieldmask_proto_rawDescData []byte
)

func file_fieldmask_proto_rawDescGZIP() []byte {
	file_fieldmask_proto_rawDescOnce.Do(func() {
		file_fieldmask_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_fieldmask_proto_rawDesc), len(file_fieldmask_proto_rawDesc)))
	})
	return file_fieldmask_proto_rawDescData
}

var file_fieldmask_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_fieldmask_proto_goTypes = []any{
	(*FieldMask)(nil),   // 0: kdspb.FieldMask
	(*Int32Array)(nil),  // 1: kdspb.Int32Array
	(*Int64Array)(nil),  // 2: kdspb.Int64Array
	(*UInt32Array)(nil), // 3: kdspb.UInt32Array
	(*Uint64Array)(nil), // 4: kdspb.Uint64Array
	(*BoolArray)(nil),   // 5: kdspb.BoolArray
	(*StringArray)(nil), // 6: kdspb.StringArray
	(*MapMask)(nil),     // 7: kdspb.MapMask
}
var file_fieldmask_proto_depIdxs = []int32{
	7, // 0: kdspb.FieldMask.MapMask:type_name -> kdspb.MapMask
	0, // 1: kdspb.FieldMask.FieldMasks:type_name -> kdspb.FieldMask
	1, // 2: kdspb.MapMask.Int32DeleteKeys:type_name -> kdspb.Int32Array
	2, // 3: kdspb.MapMask.Int64DeleteKeys:type_name -> kdspb.Int64Array
	3, // 4: kdspb.MapMask.Uint32DeleteKeys:type_name -> kdspb.UInt32Array
	4, // 5: kdspb.MapMask.Uint64DeleteKeys:type_name -> kdspb.Uint64Array
	5, // 6: kdspb.MapMask.BoolDeleteKeys:type_name -> kdspb.BoolArray
	6, // 7: kdspb.MapMask.StringDeleteKeys:type_name -> kdspb.StringArray
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_fieldmask_proto_init() }
func file_fieldmask_proto_init() {
	if File_fieldmask_proto != nil {
		return
	}
	file_fieldmask_proto_msgTypes[7].OneofWrappers = []any{
		(*MapMask_Int32DeleteKeys)(nil),
		(*MapMask_Int64DeleteKeys)(nil),
		(*MapMask_Uint32DeleteKeys)(nil),
		(*MapMask_Uint64DeleteKeys)(nil),
		(*MapMask_BoolDeleteKeys)(nil),
		(*MapMask_StringDeleteKeys)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_fieldmask_proto_rawDesc), len(file_fieldmask_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fieldmask_proto_goTypes,
		DependencyIndexes: file_fieldmask_proto_depIdxs,
		MessageInfos:      file_fieldmask_proto_msgTypes,
	}.Build()
	File_fieldmask_proto = out.File
	file_fieldmask_proto_goTypes = nil
	file_fieldmask_proto_depIdxs = nil
}
