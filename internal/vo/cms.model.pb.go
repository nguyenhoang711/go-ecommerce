// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: cms.model.proto

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

type CMS_TYPE_ID int32

const (
	CMS_TYPE_ID__UN_USE       CMS_TYPE_ID = 0
	CMS_TYPE_ID_REGISTER_USER CMS_TYPE_ID = 1 // dang ky tai khoan
	CMS_TYPE_ID_PING          CMS_TYPE_ID = 2 // PING TO SERVER
)

// Enum value maps for CMS_TYPE_ID.
var (
	CMS_TYPE_ID_name = map[int32]string{
		0: "_UN_USE",
		1: "REGISTER_USER",
		2: "PING",
	}
	CMS_TYPE_ID_value = map[string]int32{
		"_UN_USE":       0,
		"REGISTER_USER": 1,
		"PING":          2,
	}
)

func (x CMS_TYPE_ID) Enum() *CMS_TYPE_ID {
	p := new(CMS_TYPE_ID)
	*p = x
	return p
}

func (x CMS_TYPE_ID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CMS_TYPE_ID) Descriptor() protoreflect.EnumDescriptor {
	return file_cms_model_proto_enumTypes[0].Descriptor()
}

func (CMS_TYPE_ID) Type() protoreflect.EnumType {
	return &file_cms_model_proto_enumTypes[0]
}

func (x CMS_TYPE_ID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CMS_TYPE_ID.Descriptor instead.
func (CMS_TYPE_ID) EnumDescriptor() ([]byte, []int) {
	return file_cms_model_proto_rawDescGZIP(), []int{0}
}

type CMS_GROUP_ID int32

const (
	CMS_GROUP_ID__CMS_GROUP_UNUSE CMS_GROUP_ID = 0
	CMS_GROUP_ID_USER             CMS_GROUP_ID = 1
	CMS_GROUP_ID_SYST             CMS_GROUP_ID = 2
)

// Enum value maps for CMS_GROUP_ID.
var (
	CMS_GROUP_ID_name = map[int32]string{
		0: "_CMS_GROUP_UNUSE",
		1: "USER",
		2: "SYST",
	}
	CMS_GROUP_ID_value = map[string]int32{
		"_CMS_GROUP_UNUSE": 0,
		"USER":             1,
		"SYST":             2,
	}
)

func (x CMS_GROUP_ID) Enum() *CMS_GROUP_ID {
	p := new(CMS_GROUP_ID)
	*p = x
	return p
}

func (x CMS_GROUP_ID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CMS_GROUP_ID) Descriptor() protoreflect.EnumDescriptor {
	return file_cms_model_proto_enumTypes[1].Descriptor()
}

func (CMS_GROUP_ID) Type() protoreflect.EnumType {
	return &file_cms_model_proto_enumTypes[1]
}

func (x CMS_GROUP_ID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CMS_GROUP_ID.Descriptor instead.
func (CMS_GROUP_ID) EnumDescriptor() ([]byte, []int) {
	return file_cms_model_proto_rawDescGZIP(), []int{1}
}

type API_TYPE int32

const (
	API_TYPE_PORTAL API_TYPE = 0
	API_TYPE_APP    API_TYPE = 1
)

// Enum value maps for API_TYPE.
var (
	API_TYPE_name = map[int32]string{
		0: "PORTAL",
		1: "APP",
	}
	API_TYPE_value = map[string]int32{
		"PORTAL": 0,
		"APP":    1,
	}
)

func (x API_TYPE) Enum() *API_TYPE {
	p := new(API_TYPE)
	*p = x
	return p
}

func (x API_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (API_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_cms_model_proto_enumTypes[2].Descriptor()
}

func (API_TYPE) Type() protoreflect.EnumType {
	return &file_cms_model_proto_enumTypes[2]
}

func (x API_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use API_TYPE.Descriptor instead.
func (API_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_cms_model_proto_rawDescGZIP(), []int{2}
}

type BannerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TitleName string `protobuf:"bytes,2,opt,name=title_name,json=titleName,proto3" json:"title_name,omitempty"`
	ImageUrl  string `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Status    int32  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"` // 1 // lưu nháp 2 // chờ duyệt 3 // đã hoạt động 4 // hủy bỏ
	Position  uint32 `protobuf:"varint,8,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *BannerResponse) Reset() {
	*x = BannerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cms_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BannerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BannerResponse) ProtoMessage() {}

func (x *BannerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cms_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BannerResponse.ProtoReflect.Descriptor instead.
func (*BannerResponse) Descriptor() ([]byte, []int) {
	return file_cms_model_proto_rawDescGZIP(), []int{0}
}

func (x *BannerResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BannerResponse) GetTitleName() string {
	if x != nil {
		return x.TitleName
	}
	return ""
}

func (x *BannerResponse) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *BannerResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *BannerResponse) GetPosition() uint32 {
	if x != nil {
		return x.Position
	}
	return 0
}

var File_cms_model_proto protoreflect.FileDescriptor

var file_cms_model_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6d, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x90, 0x01, 0x0a, 0x0e, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x37, 0x0a, 0x0b, 0x43, 0x4d, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x49, 0x44, 0x12, 0x0b, 0x0a, 0x07, 0x5f, 0x55, 0x4e, 0x5f, 0x55, 0x53, 0x45, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x55, 0x53, 0x45,
	0x52, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x2a, 0x38, 0x0a,
	0x0c, 0x43, 0x4d, 0x53, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x10, 0x5f, 0x43, 0x4d, 0x53, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x55, 0x4e, 0x55, 0x53,
	0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x12, 0x08, 0x0a,
	0x04, 0x53, 0x59, 0x53, 0x54, 0x10, 0x02, 0x2a, 0x1f, 0x0a, 0x08, 0x41, 0x50, 0x49, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x4f, 0x52, 0x54, 0x41, 0x4c, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x41, 0x50, 0x50, 0x10, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cms_model_proto_rawDescOnce sync.Once
	file_cms_model_proto_rawDescData = file_cms_model_proto_rawDesc
)

func file_cms_model_proto_rawDescGZIP() []byte {
	file_cms_model_proto_rawDescOnce.Do(func() {
		file_cms_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_cms_model_proto_rawDescData)
	})
	return file_cms_model_proto_rawDescData
}

var file_cms_model_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_cms_model_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cms_model_proto_goTypes = []interface{}{
	(CMS_TYPE_ID)(0),       // 0: CMS_TYPE_ID
	(CMS_GROUP_ID)(0),      // 1: CMS_GROUP_ID
	(API_TYPE)(0),          // 2: API_TYPE
	(*BannerResponse)(nil), // 3: BannerResponse
}
var file_cms_model_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cms_model_proto_init() }
func file_cms_model_proto_init() {
	if File_cms_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cms_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BannerResponse); i {
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
			RawDescriptor: file_cms_model_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cms_model_proto_goTypes,
		DependencyIndexes: file_cms_model_proto_depIdxs,
		EnumInfos:         file_cms_model_proto_enumTypes,
		MessageInfos:      file_cms_model_proto_msgTypes,
	}.Build()
	File_cms_model_proto = out.File
	file_cms_model_proto_rawDesc = nil
	file_cms_model_proto_goTypes = nil
	file_cms_model_proto_depIdxs = nil
}
