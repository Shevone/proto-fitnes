// тут описан некий контракт того как должно выглядит наше gRPC приложение

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: accounts/accounts.proto

package accounts

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

type Gender int32

const (
	Gender_Men   Gender = 0
	Gender_Women Gender = 1
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "Men",
		1: "Women",
	}
	Gender_value = map[string]int32{
		"Men":   0,
		"Women": 1,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_accounts_accounts_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_accounts_accounts_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_accounts_accounts_proto_rawDescGZIP(), []int{0}
}

type UserModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid       string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Surname    string `protobuf:"bytes,4,opt,name=surname,proto3" json:"surname,omitempty"`
	Patronymic string `protobuf:"bytes,5,opt,name=patronymic,proto3" json:"patronymic,omitempty"`
	Gender     Gender `protobuf:"varint,6,opt,name=gender,proto3,enum=accounts.Gender" json:"gender,omitempty"`
	Age        int32  `protobuf:"varint,7,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *UserModel) Reset() {
	*x = UserModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_accounts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserModel) ProtoMessage() {}

func (x *UserModel) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_accounts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserModel.ProtoReflect.Descriptor instead.
func (*UserModel) Descriptor() ([]byte, []int) {
	return file_accounts_accounts_proto_rawDescGZIP(), []int{0}
}

func (x *UserModel) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *UserModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserModel) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *UserModel) GetPatronymic() string {
	if x != nil {
		return x.Patronymic
	}
	return ""
}

func (x *UserModel) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_Men
}

func (x *UserModel) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

// =============================================================================
type EditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUUID   string `protobuf:"bytes,2,opt,name=userUUID,proto3" json:"userUUID,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Surname    string `protobuf:"bytes,4,opt,name=surname,proto3" json:"surname,omitempty"`
	Patronymic string `protobuf:"bytes,5,opt,name=patronymic,proto3" json:"patronymic,omitempty"`
}

func (x *EditRequest) Reset() {
	*x = EditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_accounts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditRequest) ProtoMessage() {}

func (x *EditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_accounts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditRequest.ProtoReflect.Descriptor instead.
func (*EditRequest) Descriptor() ([]byte, []int) {
	return file_accounts_accounts_proto_rawDescGZIP(), []int{1}
}

func (x *EditRequest) GetUserUUID() string {
	if x != nil {
		return x.UserUUID
	}
	return ""
}

func (x *EditRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EditRequest) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *EditRequest) GetPatronymic() string {
	if x != nil {
		return x.Patronymic
	}
	return ""
}

type EditResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EditResponse) Reset() {
	*x = EditResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_accounts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditResponse) ProtoMessage() {}

func (x *EditResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_accounts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditResponse.ProtoReflect.Descriptor instead.
func (*EditResponse) Descriptor() ([]byte, []int) {
	return file_accounts_accounts_proto_rawDescGZIP(), []int{2}
}

func (x *EditResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// ============================================================================
type GetUserRequests struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUUUID string `protobuf:"bytes,1,opt,name=userUUUID,proto3" json:"userUUUID,omitempty"`
}

func (x *GetUserRequests) Reset() {
	*x = GetUserRequests{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_accounts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequests) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequests) ProtoMessage() {}

func (x *GetUserRequests) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_accounts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequests.ProtoReflect.Descriptor instead.
func (*GetUserRequests) Descriptor() ([]byte, []int) {
	return file_accounts_accounts_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserRequests) GetUserUUUID() string {
	if x != nil {
		return x.UserUUUID
	}
	return ""
}

type GetDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *UserModel `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetDataResponse) Reset() {
	*x = GetDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_accounts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataResponse) ProtoMessage() {}

func (x *GetDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_accounts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataResponse.ProtoReflect.Descriptor instead.
func (*GetDataResponse) Descriptor() ([]byte, []int) {
	return file_accounts_accounts_proto_rawDescGZIP(), []int{4}
}

func (x *GetDataResponse) GetUser() *UserModel {
	if x != nil {
		return x.User
	}
	return nil
}

// ============================================================================
type GetUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetUsersRequest) Reset() {
	*x = GetUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_accounts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersRequest) ProtoMessage() {}

func (x *GetUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_accounts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersRequest.ProtoReflect.Descriptor instead.
func (*GetUsersRequest) Descriptor() ([]byte, []int) {
	return file_accounts_accounts_proto_rawDescGZIP(), []int{5}
}

func (x *GetUsersRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetUsersRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserModel `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetUsersResponse) Reset() {
	*x = GetUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_accounts_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersResponse) ProtoMessage() {}

func (x *GetUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_accounts_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersResponse.ProtoReflect.Descriptor instead.
func (*GetUsersResponse) Descriptor() ([]byte, []int) {
	return file_accounts_accounts_proto_rawDescGZIP(), []int{6}
}

func (x *GetUsersResponse) GetUsers() []*UserModel {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_accounts_accounts_proto protoreflect.FileDescriptor

var file_accounts_accounts_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x22, 0xa9, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x69,
	0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x74, 0x72, 0x6f, 0x6e, 0x79,
	0x6d, 0x69, 0x63, 0x12, 0x28, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x47,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22,
	0x77, 0x0a, 0x0b, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x72,
	0x6f, 0x6e, 0x79, 0x6d, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61,
	0x74, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x69, 0x63, 0x22, 0x28, 0x0a, 0x0c, 0x45, 0x64, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x2f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x55, 0x55, 0x55,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x55, 0x55,
	0x55, 0x49, 0x44, 0x22, 0x3a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x3b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x3d, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x29, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2a, 0x1c, 0x0a, 0x06, 0x47,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x65, 0x6e, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x57, 0x6f, 0x6d, 0x65, 0x6e, 0x10, 0x01, 0x32, 0xd0, 0x01, 0x0a, 0x08, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x3c, 0x0a, 0x0b, 0x45, 0x64, 0x69, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x2e, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x19, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x1a, 0x19,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x19, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x10,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accounts_accounts_proto_rawDescOnce sync.Once
	file_accounts_accounts_proto_rawDescData = file_accounts_accounts_proto_rawDesc
)

func file_accounts_accounts_proto_rawDescGZIP() []byte {
	file_accounts_accounts_proto_rawDescOnce.Do(func() {
		file_accounts_accounts_proto_rawDescData = protoimpl.X.CompressGZIP(file_accounts_accounts_proto_rawDescData)
	})
	return file_accounts_accounts_proto_rawDescData
}

var file_accounts_accounts_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_accounts_accounts_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_accounts_accounts_proto_goTypes = []interface{}{
	(Gender)(0),              // 0: accounts.Gender
	(*UserModel)(nil),        // 1: accounts.UserModel
	(*EditRequest)(nil),      // 2: accounts.EditRequest
	(*EditResponse)(nil),     // 3: accounts.EditResponse
	(*GetUserRequests)(nil),  // 4: accounts.GetUserRequests
	(*GetDataResponse)(nil),  // 5: accounts.GetDataResponse
	(*GetUsersRequest)(nil),  // 6: accounts.GetUsersRequest
	(*GetUsersResponse)(nil), // 7: accounts.GetUsersResponse
}
var file_accounts_accounts_proto_depIdxs = []int32{
	0, // 0: accounts.UserModel.gender:type_name -> accounts.Gender
	1, // 1: accounts.GetDataResponse.user:type_name -> accounts.UserModel
	1, // 2: accounts.GetUsersResponse.users:type_name -> accounts.UserModel
	2, // 3: accounts.Accounts.EditProfile:input_type -> accounts.EditRequest
	4, // 4: accounts.Accounts.GetUserData:input_type -> accounts.GetUserRequests
	6, // 5: accounts.Accounts.GetUsers:input_type -> accounts.GetUsersRequest
	3, // 6: accounts.Accounts.EditProfile:output_type -> accounts.EditResponse
	5, // 7: accounts.Accounts.GetUserData:output_type -> accounts.GetDataResponse
	7, // 8: accounts.Accounts.GetUsers:output_type -> accounts.GetUsersResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_accounts_accounts_proto_init() }
func file_accounts_accounts_proto_init() {
	if File_accounts_accounts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_accounts_accounts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserModel); i {
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
		file_accounts_accounts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditRequest); i {
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
		file_accounts_accounts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditResponse); i {
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
		file_accounts_accounts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequests); i {
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
		file_accounts_accounts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataResponse); i {
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
		file_accounts_accounts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersRequest); i {
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
		file_accounts_accounts_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersResponse); i {
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
			RawDescriptor: file_accounts_accounts_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_accounts_accounts_proto_goTypes,
		DependencyIndexes: file_accounts_accounts_proto_depIdxs,
		EnumInfos:         file_accounts_accounts_proto_enumTypes,
		MessageInfos:      file_accounts_accounts_proto_msgTypes,
	}.Build()
	File_accounts_accounts_proto = out.File
	file_accounts_accounts_proto_rawDesc = nil
	file_accounts_accounts_proto_goTypes = nil
	file_accounts_accounts_proto_depIdxs = nil
}
