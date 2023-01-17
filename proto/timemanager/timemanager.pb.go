// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: timemanager/timemanager.proto

package timemanager

import (
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

type TimeStatus int32

const (
	TimeStatus_success TimeStatus = 0
	TimeStatus_fail    TimeStatus = 1
)

// Enum value maps for TimeStatus.
var (
	TimeStatus_name = map[int32]string{
		0: "success",
		1: "fail",
	}
	TimeStatus_value = map[string]int32{
		"success": 0,
		"fail":    1,
	}
)

func (x TimeStatus) Enum() *TimeStatus {
	p := new(TimeStatus)
	*p = x
	return p
}

func (x TimeStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TimeStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_timemanager_timemanager_proto_enumTypes[0].Descriptor()
}

func (TimeStatus) Type() protoreflect.EnumType {
	return &file_timemanager_timemanager_proto_enumTypes[0]
}

func (x TimeStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TimeStatus.Descriptor instead.
func (TimeStatus) EnumDescriptor() ([]byte, []int) {
	return file_timemanager_timemanager_proto_rawDescGZIP(), []int{0}
}

type GetServerTimeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset  int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Accrate int32 `protobuf:"varint,2,opt,name=accrate,proto3" json:"accrate,omitempty"`
}

func (x *GetServerTimeRes) Reset() {
	*x = GetServerTimeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timemanager_timemanager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServerTimeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServerTimeRes) ProtoMessage() {}

func (x *GetServerTimeRes) ProtoReflect() protoreflect.Message {
	mi := &file_timemanager_timemanager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServerTimeRes.ProtoReflect.Descriptor instead.
func (*GetServerTimeRes) Descriptor() ([]byte, []int) {
	return file_timemanager_timemanager_proto_rawDescGZIP(), []int{0}
}

func (x *GetServerTimeRes) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetServerTimeRes) GetAccrate() int32 {
	if x != nil {
		return x.Accrate
	}
	return 0
}

type SetServerTimeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset  int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Accrate int32 `protobuf:"varint,2,opt,name=accrate,proto3" json:"accrate,omitempty"`
}

func (x *SetServerTimeReq) Reset() {
	*x = SetServerTimeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timemanager_timemanager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetServerTimeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetServerTimeReq) ProtoMessage() {}

func (x *SetServerTimeReq) ProtoReflect() protoreflect.Message {
	mi := &file_timemanager_timemanager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetServerTimeReq.ProtoReflect.Descriptor instead.
func (*SetServerTimeReq) Descriptor() ([]byte, []int) {
	return file_timemanager_timemanager_proto_rawDescGZIP(), []int{1}
}

func (x *SetServerTimeReq) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *SetServerTimeReq) GetAccrate() int32 {
	if x != nil {
		return x.Accrate
	}
	return 0
}

type SetServerTimeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status TimeStatus `protobuf:"varint,1,opt,name=status,proto3,enum=timemanager.TimeStatus" json:"status,omitempty"`
}

func (x *SetServerTimeRes) Reset() {
	*x = SetServerTimeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timemanager_timemanager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetServerTimeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetServerTimeRes) ProtoMessage() {}

func (x *SetServerTimeRes) ProtoReflect() protoreflect.Message {
	mi := &file_timemanager_timemanager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetServerTimeRes.ProtoReflect.Descriptor instead.
func (*SetServerTimeRes) Descriptor() ([]byte, []int) {
	return file_timemanager_timemanager_proto_rawDescGZIP(), []int{2}
}

func (x *SetServerTimeRes) GetStatus() TimeStatus {
	if x != nil {
		return x.Status
	}
	return TimeStatus_success
}

type ClearServerTimeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status TimeStatus `protobuf:"varint,1,opt,name=status,proto3,enum=timemanager.TimeStatus" json:"status,omitempty"`
}

func (x *ClearServerTimeRes) Reset() {
	*x = ClearServerTimeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timemanager_timemanager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearServerTimeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearServerTimeRes) ProtoMessage() {}

func (x *ClearServerTimeRes) ProtoReflect() protoreflect.Message {
	mi := &file_timemanager_timemanager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearServerTimeRes.ProtoReflect.Descriptor instead.
func (*ClearServerTimeRes) Descriptor() ([]byte, []int) {
	return file_timemanager_timemanager_proto_rawDescGZIP(), []int{3}
}

func (x *ClearServerTimeRes) GetStatus() TimeStatus {
	if x != nil {
		return x.Status
	}
	return TimeStatus_success
}

var File_timemanager_timemanager_proto protoreflect.FileDescriptor

var file_timemanager_timemanager_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x74, 0x69, 0x6d, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x74, 0x69, 0x6d, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x72, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x63, 0x63, 0x72, 0x61, 0x74, 0x65, 0x22,
	0x44, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x63, 0x63, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x63,
	0x63, 0x72, 0x61, 0x74, 0x65, 0x22, 0x43, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x74, 0x69, 0x6d, 0x65,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x45, 0x0a, 0x12, 0x43, 0x6c,
	0x65, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x17, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2a, 0x23, 0x0a, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0b, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x66, 0x61, 0x69, 0x6c, 0x10, 0x01, 0x32, 0xf3, 0x01, 0x0a, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x47,
	0x52, 0x50, 0x43, 0x12, 0x48, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x74,
	0x69, 0x6d, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x4f, 0x0a,
	0x0d, 0x53, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d,
	0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e,
	0x74, 0x69, 0x6d, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x4c,
	0x0a, 0x0f, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x74, 0x69, 0x6d, 0x65,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x6e, 0x67, 0x79,
	0x75, 0x63, 0x68, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_timemanager_timemanager_proto_rawDescOnce sync.Once
	file_timemanager_timemanager_proto_rawDescData = file_timemanager_timemanager_proto_rawDesc
)

func file_timemanager_timemanager_proto_rawDescGZIP() []byte {
	file_timemanager_timemanager_proto_rawDescOnce.Do(func() {
		file_timemanager_timemanager_proto_rawDescData = protoimpl.X.CompressGZIP(file_timemanager_timemanager_proto_rawDescData)
	})
	return file_timemanager_timemanager_proto_rawDescData
}

var file_timemanager_timemanager_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_timemanager_timemanager_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_timemanager_timemanager_proto_goTypes = []interface{}{
	(TimeStatus)(0),            // 0: timemanager.TimeStatus
	(*GetServerTimeRes)(nil),   // 1: timemanager.GetServerTimeRes
	(*SetServerTimeReq)(nil),   // 2: timemanager.SetServerTimeReq
	(*SetServerTimeRes)(nil),   // 3: timemanager.SetServerTimeRes
	(*ClearServerTimeRes)(nil), // 4: timemanager.ClearServerTimeRes
	(*emptypb.Empty)(nil),      // 5: google.protobuf.Empty
}
var file_timemanager_timemanager_proto_depIdxs = []int32{
	0, // 0: timemanager.SetServerTimeRes.status:type_name -> timemanager.TimeStatus
	0, // 1: timemanager.ClearServerTimeRes.status:type_name -> timemanager.TimeStatus
	5, // 2: timemanager.TimeGRPC.GetServerTime:input_type -> google.protobuf.Empty
	2, // 3: timemanager.TimeGRPC.SetServerTime:input_type -> timemanager.SetServerTimeReq
	5, // 4: timemanager.TimeGRPC.ClearServerTime:input_type -> google.protobuf.Empty
	1, // 5: timemanager.TimeGRPC.GetServerTime:output_type -> timemanager.GetServerTimeRes
	3, // 6: timemanager.TimeGRPC.SetServerTime:output_type -> timemanager.SetServerTimeRes
	4, // 7: timemanager.TimeGRPC.ClearServerTime:output_type -> timemanager.ClearServerTimeRes
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_timemanager_timemanager_proto_init() }
func file_timemanager_timemanager_proto_init() {
	if File_timemanager_timemanager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_timemanager_timemanager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServerTimeRes); i {
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
		file_timemanager_timemanager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetServerTimeReq); i {
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
		file_timemanager_timemanager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetServerTimeRes); i {
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
		file_timemanager_timemanager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearServerTimeRes); i {
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
			RawDescriptor: file_timemanager_timemanager_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_timemanager_timemanager_proto_goTypes,
		DependencyIndexes: file_timemanager_timemanager_proto_depIdxs,
		EnumInfos:         file_timemanager_timemanager_proto_enumTypes,
		MessageInfos:      file_timemanager_timemanager_proto_msgTypes,
	}.Build()
	File_timemanager_timemanager_proto = out.File
	file_timemanager_timemanager_proto_rawDesc = nil
	file_timemanager_timemanager_proto_goTypes = nil
	file_timemanager_timemanager_proto_depIdxs = nil
}