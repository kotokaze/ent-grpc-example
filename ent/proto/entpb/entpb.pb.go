// Code generated by entproto. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: entpb/entpb.proto

package entpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetUserRequest_View int32

const (
	GetUserRequest_VIEW_UNSPECIFIED GetUserRequest_View = 0
	GetUserRequest_BASIC            GetUserRequest_View = 1
	GetUserRequest_WITH_EDGE_IDS    GetUserRequest_View = 2
)

// Enum value maps for GetUserRequest_View.
var (
	GetUserRequest_View_name = map[int32]string{
		0: "VIEW_UNSPECIFIED",
		1: "BASIC",
		2: "WITH_EDGE_IDS",
	}
	GetUserRequest_View_value = map[string]int32{
		"VIEW_UNSPECIFIED": 0,
		"BASIC":            1,
		"WITH_EDGE_IDS":    2,
	}
)

func (x GetUserRequest_View) Enum() *GetUserRequest_View {
	p := new(GetUserRequest_View)
	*p = x
	return p
}

func (x GetUserRequest_View) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetUserRequest_View) Descriptor() protoreflect.EnumDescriptor {
	return file_entpb_entpb_proto_enumTypes[0].Descriptor()
}

func (GetUserRequest_View) Type() protoreflect.EnumType {
	return &file_entpb_entpb_proto_enumTypes[0]
}

func (x GetUserRequest_View) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetUserRequest_View.Descriptor instead.
func (GetUserRequest_View) EnumDescriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{3, 0}
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Admin *User  `protobuf:"bytes,3,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{0}
}

func (x *Category) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Category) GetAdmin() *User {
	if x != nil {
		return x.Admin
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	EmailAddress string                  `protobuf:"bytes,3,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	Alias        *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=alias,proto3" json:"alias,omitempty"`
	Administered []*Category             `protobuf:"bytes,5,rep,name=administered,proto3" json:"administered,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *User) GetAlias() *wrapperspb.StringValue {
	if x != nil {
		return x.Alias
	}
	return nil
}

func (x *User) GetAdministered() []*Category {
	if x != nil {
		return x.Administered
	}
	return nil
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	View GetUserRequest_View `protobuf:"varint,2,opt,name=view,proto3,enum=entpb.GetUserRequest_View" json:"view,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetUserRequest) GetView() GetUserRequest_View {
	if x != nil {
		return x.View
	}
	return GetUserRequest_VIEW_UNSPECIFIED
}

var File_entpb_entpb_proto protoreflect.FileDescriptor

var file_entpb_entpb_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0xb8, 0x01, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x32, 0x0a, 0x05,
	0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x12, 0x33, 0x0a, 0x0c, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0c, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x65, 0x64, 0x22, 0x34, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x8c, 0x01, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e,
	0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x22, 0x3a,
	0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x14, 0x0a, 0x10, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x42, 0x41, 0x53, 0x49, 0x43, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x57, 0x49, 0x54, 0x48, 0x5f,
	0x45, 0x44, 0x47, 0x45, 0x5f, 0x49, 0x44, 0x53, 0x10, 0x02, 0x32, 0x69, 0x0a, 0x0b, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x15, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x74, 0x6f, 0x6b, 0x61, 0x7a, 0x65, 0x2f, 0x65, 0x6e, 0x74,
	0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x6e,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entpb_entpb_proto_rawDescOnce sync.Once
	file_entpb_entpb_proto_rawDescData = file_entpb_entpb_proto_rawDesc
)

func file_entpb_entpb_proto_rawDescGZIP() []byte {
	file_entpb_entpb_proto_rawDescOnce.Do(func() {
		file_entpb_entpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_entpb_entpb_proto_rawDescData)
	})
	return file_entpb_entpb_proto_rawDescData
}

var file_entpb_entpb_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_entpb_entpb_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_entpb_entpb_proto_goTypes = []interface{}{
	(GetUserRequest_View)(0),       // 0: entpb.GetUserRequest.View
	(*Category)(nil),               // 1: entpb.Category
	(*User)(nil),                   // 2: entpb.User
	(*CreateUserRequest)(nil),      // 3: entpb.CreateUserRequest
	(*GetUserRequest)(nil),         // 4: entpb.GetUserRequest
	(*wrapperspb.StringValue)(nil), // 5: google.protobuf.StringValue
}
var file_entpb_entpb_proto_depIdxs = []int32{
	2, // 0: entpb.Category.admin:type_name -> entpb.User
	5, // 1: entpb.User.alias:type_name -> google.protobuf.StringValue
	1, // 2: entpb.User.administered:type_name -> entpb.Category
	2, // 3: entpb.CreateUserRequest.user:type_name -> entpb.User
	0, // 4: entpb.GetUserRequest.view:type_name -> entpb.GetUserRequest.View
	3, // 5: entpb.UserService.Create:input_type -> entpb.CreateUserRequest
	4, // 6: entpb.UserService.Get:input_type -> entpb.GetUserRequest
	2, // 7: entpb.UserService.Create:output_type -> entpb.User
	2, // 8: entpb.UserService.Get:output_type -> entpb.User
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_entpb_entpb_proto_init() }
func file_entpb_entpb_proto_init() {
	if File_entpb_entpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entpb_entpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
		file_entpb_entpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_entpb_entpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_entpb_entpb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
			RawDescriptor: file_entpb_entpb_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_entpb_entpb_proto_goTypes,
		DependencyIndexes: file_entpb_entpb_proto_depIdxs,
		EnumInfos:         file_entpb_entpb_proto_enumTypes,
		MessageInfos:      file_entpb_entpb_proto_msgTypes,
	}.Build()
	File_entpb_entpb_proto = out.File
	file_entpb_entpb_proto_rawDesc = nil
	file_entpb_entpb_proto_goTypes = nil
	file_entpb_entpb_proto_depIdxs = nil
}