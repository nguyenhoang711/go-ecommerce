// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: portal.user.proto

package vo

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

type RegisterUser_ERROR int32

const (
	RegisterUser_NO_ERROR        RegisterUser_ERROR = 0
	RegisterUser_USER_EXISTED    RegisterUser_ERROR = 1
	RegisterUser_INVALID_OTP     RegisterUser_ERROR = 2
	RegisterUser_OTP_NOT_EXISTS  RegisterUser_ERROR = 3
	RegisterUser_SEND_OTP_ERRROR RegisterUser_ERROR = 4
)

// Enum value maps for RegisterUser_ERROR.
var (
	RegisterUser_ERROR_name = map[int32]string{
		0: "NO_ERROR",
		1: "USER_EXISTED",
		2: "INVALID_OTP",
		3: "OTP_NOT_EXISTS",
		4: "SEND_OTP_ERRROR",
	}
	RegisterUser_ERROR_value = map[string]int32{
		"NO_ERROR":        0,
		"USER_EXISTED":    1,
		"INVALID_OTP":     2,
		"OTP_NOT_EXISTS":  3,
		"SEND_OTP_ERRROR": 4,
	}
)

func (x RegisterUser_ERROR) Enum() *RegisterUser_ERROR {
	p := new(RegisterUser_ERROR)
	*p = x
	return p
}

func (x RegisterUser_ERROR) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegisterUser_ERROR) Descriptor() protoreflect.EnumDescriptor {
	return file_portal_user_proto_enumTypes[0].Descriptor()
}

func (RegisterUser_ERROR) Type() protoreflect.EnumType {
	return &file_portal_user_proto_enumTypes[0]
}

func (x RegisterUser_ERROR) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegisterUser_ERROR.Descriptor instead.
func (RegisterUser_ERROR) EnumDescriptor() ([]byte, []int) {
	return file_portal_user_proto_rawDescGZIP(), []int{0, 0}
}

type RegisterUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterUser) Reset() {
	*x = RegisterUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUser) ProtoMessage() {}

func (x *RegisterUser) ProtoReflect() protoreflect.Message {
	mi := &file_portal_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUser.ProtoReflect.Descriptor instead.
func (*RegisterUser) Descriptor() ([]byte, []int) {
	return file_portal_user_proto_rawDescGZIP(), []int{0}
}

type RegisterUser_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VerifyKey     string `protobuf:"bytes,1,opt,name=verify_key,json=verifyKey,proto3" json:"verify_key,omitempty"`             // @gotags: form:"verify_key"
	VerifyType    int32  `protobuf:"varint,2,opt,name=verify_type,json=verifyType,proto3" json:"verify_type,omitempty"`         // @gotags: form:"verify_type"
	VerifyPurpose string `protobuf:"bytes,3,opt,name=verify_purpose,json=verifyPurpose,proto3" json:"verify_purpose,omitempty"` // @gotags: form:"verify_purpose"
}

func (x *RegisterUser_Request) Reset() {
	*x = RegisterUser_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUser_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUser_Request) ProtoMessage() {}

func (x *RegisterUser_Request) ProtoReflect() protoreflect.Message {
	mi := &file_portal_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUser_Request.ProtoReflect.Descriptor instead.
func (*RegisterUser_Request) Descriptor() ([]byte, []int) {
	return file_portal_user_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RegisterUser_Request) GetVerifyKey() string {
	if x != nil {
		return x.VerifyKey
	}
	return ""
}

func (x *RegisterUser_Request) GetVerifyType() int32 {
	if x != nil {
		return x.VerifyType
	}
	return 0
}

func (x *RegisterUser_Request) GetVerifyPurpose() string {
	if x != nil {
		return x.VerifyPurpose
	}
	return ""
}

type RegisterUser_Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterUser_Reply) Reset() {
	*x = RegisterUser_Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUser_Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUser_Reply) ProtoMessage() {}

func (x *RegisterUser_Reply) ProtoReflect() protoreflect.Message {
	mi := &file_portal_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUser_Reply.ProtoReflect.Descriptor instead.
func (*RegisterUser_Reply) Descriptor() ([]byte, []int) {
	return file_portal_user_proto_rawDescGZIP(), []int{0, 1}
}

var File_portal_user_proto protoreflect.FileDescriptor

var file_portal_user_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xec, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x1a, 0x70, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x1f,
	0x0a, 0x0b, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50,
	0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x1a, 0x07, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x61, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x45,
	0x58, 0x49, 0x53, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x5f, 0x4f, 0x54, 0x50, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x54, 0x50,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x03, 0x12, 0x13, 0x0a,
	0x0f, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x4f, 0x54, 0x50, 0x5f, 0x45, 0x52, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x04, 0x42, 0x0d, 0x5a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_portal_user_proto_rawDescOnce sync.Once
	file_portal_user_proto_rawDescData = file_portal_user_proto_rawDesc
)

func file_portal_user_proto_rawDescGZIP() []byte {
	file_portal_user_proto_rawDescOnce.Do(func() {
		file_portal_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_portal_user_proto_rawDescData)
	})
	return file_portal_user_proto_rawDescData
}

var file_portal_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_portal_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_portal_user_proto_goTypes = []interface{}{
	(RegisterUser_ERROR)(0),      // 0: RegisterUser.ERROR
	(*RegisterUser)(nil),         // 1: RegisterUser
	(*RegisterUser_Request)(nil), // 2: RegisterUser.Request
	(*RegisterUser_Reply)(nil),   // 3: RegisterUser.Reply
}
var file_portal_user_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_portal_user_proto_init() }
func file_portal_user_proto_init() {
	if File_portal_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_portal_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterUser); i {
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
		file_portal_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterUser_Request); i {
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
		file_portal_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterUser_Reply); i {
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
			RawDescriptor: file_portal_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_portal_user_proto_goTypes,
		DependencyIndexes: file_portal_user_proto_depIdxs,
		EnumInfos:         file_portal_user_proto_enumTypes,
		MessageInfos:      file_portal_user_proto_msgTypes,
	}.Build()
	File_portal_user_proto = out.File
	file_portal_user_proto_rawDesc = nil
	file_portal_user_proto_goTypes = nil
	file_portal_user_proto_depIdxs = nil
}
