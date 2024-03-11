// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0--rc1
// source: lessons/lessons.proto

package lessonsFitnesv1

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

type Difficulty int32

const (
	Difficulty_EASY   Difficulty = 0
	Difficulty_MEDIUM Difficulty = 1
	Difficulty_HARD   Difficulty = 2
)

// Enum value maps for Difficulty.
var (
	Difficulty_name = map[int32]string{
		0: "EASY",
		1: "MEDIUM",
		2: "HARD",
	}
	Difficulty_value = map[string]int32{
		"EASY":   0,
		"MEDIUM": 1,
		"HARD":   2,
	}
)

func (x Difficulty) Enum() *Difficulty {
	p := new(Difficulty)
	*p = x
	return p
}

func (x Difficulty) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Difficulty) Descriptor() protoreflect.EnumDescriptor {
	return file_lessons_lessons_proto_enumTypes[0].Descriptor()
}

func (Difficulty) Type() protoreflect.EnumType {
	return &file_lessons_lessons_proto_enumTypes[0]
}

func (x Difficulty) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Difficulty.Descriptor instead.
func (Difficulty) EnumDescriptor() ([]byte, []int) {
	return file_lessons_lessons_proto_rawDescGZIP(), []int{0}
}

type Lesson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LessonId       int64      `protobuf:"varint,1,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`
	Title          string     `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"` // Название занятия
	Time           string     `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	TrainerId      int64      `protobuf:"varint,4,opt,name=trainer_id,json=trainerId,proto3" json:"trainer_id,omitempty"`
	AvailableSeats int32      `protobuf:"varint,5,opt,name=available_seats,json=availableSeats,proto3" json:"available_seats,omitempty"`
	Description    string     `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Difficulty     Difficulty `protobuf:"varint,7,opt,name=difficulty,proto3,enum=lessons.Difficulty" json:"difficulty,omitempty"`
	IsComplete     int32      `protobuf:"varint,8,opt,name=is_complete,json=isComplete,proto3" json:"is_complete,omitempty"`
}

func (x *Lesson) Reset() {
	*x = Lesson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lessons_lessons_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lesson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lesson) ProtoMessage() {}

func (x *Lesson) ProtoReflect() protoreflect.Message {
	mi := &file_lessons_lessons_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lesson.ProtoReflect.Descriptor instead.
func (*Lesson) Descriptor() ([]byte, []int) {
	return file_lessons_lessons_proto_rawDescGZIP(), []int{0}
}

func (x *Lesson) GetLessonId() int64 {
	if x != nil {
		return x.LessonId
	}
	return 0
}

func (x *Lesson) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Lesson) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *Lesson) GetTrainerId() int64 {
	if x != nil {
		return x.TrainerId
	}
	return 0
}

func (x *Lesson) GetAvailableSeats() int32 {
	if x != nil {
		return x.AvailableSeats
	}
	return 0
}

func (x *Lesson) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Lesson) GetDifficulty() Difficulty {
	if x != nil {
		return x.Difficulty
	}
	return Difficulty_EASY
}

func (x *Lesson) GetIsComplete() int32 {
	if x != nil {
		return x.IsComplete
	}
	return 0
}

// ======================================================================================
type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title          string     `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"` // Название занятия
	Time           string     `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	TrainerId      int64      `protobuf:"varint,3,opt,name=trainer_id,json=trainerId,proto3" json:"trainer_id,omitempty"`
	AvailableSeats int32      `protobuf:"varint,4,opt,name=available_seats,json=availableSeats,proto3" json:"available_seats,omitempty"`
	Description    string     `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Difficulty     Difficulty `protobuf:"varint,6,opt,name=difficulty,proto3,enum=lessons.Difficulty" json:"difficulty,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lessons_lessons_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lessons_lessons_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_lessons_lessons_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateRequest) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *CreateRequest) GetTrainerId() int64 {
	if x != nil {
		return x.TrainerId
	}
	return 0
}

func (x *CreateRequest) GetAvailableSeats() int32 {
	if x != nil {
		return x.AvailableSeats
	}
	return 0
}

func (x *CreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateRequest) GetDifficulty() Difficulty {
	if x != nil {
		return x.Difficulty
	}
	return Difficulty_EASY
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LessonId int64 `protobuf:"varint,1,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lessons_lessons_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lessons_lessons_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_lessons_lessons_proto_rawDescGZIP(), []int{2}
}

func (x *CreateResponse) GetLessonId() int64 {
	if x != nil {
		return x.LessonId
	}
	return 0
}

// ======================================================================================
type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LessonId int64 `protobuf:"varint,1,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lessons_lessons_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lessons_lessons_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_lessons_lessons_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRequest) GetLessonId() int64 {
	if x != nil {
		return x.LessonId
	}
	return 0
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lessons_lessons_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lessons_lessons_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_lessons_lessons_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

// ======================================================================================
type EditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lesson *Lesson `protobuf:"bytes,1,opt,name=lesson,proto3" json:"lesson,omitempty"`
}

func (x *EditRequest) Reset() {
	*x = EditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lessons_lessons_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditRequest) ProtoMessage() {}

func (x *EditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lessons_lessons_proto_msgTypes[5]
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
	return file_lessons_lessons_proto_rawDescGZIP(), []int{5}
}

func (x *EditRequest) GetLesson() *Lesson {
	if x != nil {
		return x.Lesson
	}
	return nil
}

type EditResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LessonId int64 `protobuf:"varint,1,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`
}

func (x *EditResponse) Reset() {
	*x = EditResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lessons_lessons_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditResponse) ProtoMessage() {}

func (x *EditResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lessons_lessons_proto_msgTypes[6]
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
	return file_lessons_lessons_proto_rawDescGZIP(), []int{6}
}

func (x *EditResponse) GetLessonId() int64 {
	if x != nil {
		return x.LessonId
	}
	return 0
}

// ======================================================================================
type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GetId int64 `protobuf:"varint,1,opt,name=get_id,json=getId,proto3" json:"get_id,omitempty"`
	Page  int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lessons_lessons_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lessons_lessons_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_lessons_lessons_proto_rawDescGZIP(), []int{7}
}

func (x *GetRequest) GetGetId() int64 {
	if x != nil {
		return x.GetId
	}
	return 0
}

func (x *GetRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lessons []*Lesson `protobuf:"bytes,1,rep,name=lessons,proto3" json:"lessons,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lessons_lessons_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lessons_lessons_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_lessons_lessons_proto_rawDescGZIP(), []int{8}
}

func (x *GetResponse) GetLessons() []*Lesson {
	if x != nil {
		return x.Lessons
	}
	return nil
}

// ======================================================================================
type CloseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrainerId int64 `protobuf:"varint,1,opt,name=trainer_id,json=trainerId,proto3" json:"trainer_id,omitempty"`
}

func (x *CloseRequest) Reset() {
	*x = CloseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lessons_lessons_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseRequest) ProtoMessage() {}

func (x *CloseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lessons_lessons_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseRequest.ProtoReflect.Descriptor instead.
func (*CloseRequest) Descriptor() ([]byte, []int) {
	return file_lessons_lessons_proto_rawDescGZIP(), []int{9}
}

func (x *CloseRequest) GetTrainerId() int64 {
	if x != nil {
		return x.TrainerId
	}
	return 0
}

type CloseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CloseResponse) Reset() {
	*x = CloseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lessons_lessons_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseResponse) ProtoMessage() {}

func (x *CloseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lessons_lessons_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseResponse.ProtoReflect.Descriptor instead.
func (*CloseResponse) Descriptor() ([]byte, []int) {
	return file_lessons_lessons_proto_rawDescGZIP(), []int{10}
}

func (x *CloseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// ======================================================================================
type SignUpOrCancelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	LessonId int64 `protobuf:"varint,2,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`
}

func (x *SignUpOrCancelRequest) Reset() {
	*x = SignUpOrCancelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lessons_lessons_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpOrCancelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpOrCancelRequest) ProtoMessage() {}

func (x *SignUpOrCancelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lessons_lessons_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpOrCancelRequest.ProtoReflect.Descriptor instead.
func (*SignUpOrCancelRequest) Descriptor() ([]byte, []int) {
	return file_lessons_lessons_proto_rawDescGZIP(), []int{11}
}

func (x *SignUpOrCancelRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SignUpOrCancelRequest) GetLessonId() int64 {
	if x != nil {
		return x.LessonId
	}
	return 0
}

type SignUpOrCancelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SignUpOrCancelResponse) Reset() {
	*x = SignUpOrCancelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lessons_lessons_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpOrCancelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpOrCancelResponse) ProtoMessage() {}

func (x *SignUpOrCancelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lessons_lessons_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpOrCancelResponse.ProtoReflect.Descriptor instead.
func (*SignUpOrCancelResponse) Descriptor() ([]byte, []int) {
	return file_lessons_lessons_proto_rawDescGZIP(), []int{12}
}

func (x *SignUpOrCancelResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_lessons_lessons_proto protoreflect.FileDescriptor

var file_lessons_lessons_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73,
	0x22, 0x8f, 0x02, 0x0a, 0x06, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73,
	0x65, 0x61, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x61, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x0a,
	0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x13, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x69, 0x66, 0x66, 0x69,
	0x63, 0x75, 0x6c, 0x74, 0x79, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74,
	0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x22, 0xd8, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a,
	0x0f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x53, 0x65, 0x61, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66,
	0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6c,
	0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74,
	0x79, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x22, 0x2d, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x0e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x36, 0x0a, 0x0b, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x65,
	0x73, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x22, 0x2b, 0x0a, 0x0c,
	0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x67, 0x65, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x22, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x29, 0x0a, 0x07, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x52, 0x07, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x22, 0x2d, 0x0a, 0x0c,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x0d, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4d, 0x0a, 0x15, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70,
	0x4f, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x16, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x4f,
	0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x2c, 0x0a, 0x0a, 0x44, 0x69, 0x66,
	0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x41, 0x53, 0x59, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x01, 0x12, 0x08, 0x0a,
	0x04, 0x48, 0x41, 0x52, 0x44, 0x10, 0x02, 0x32, 0xa4, 0x04, 0x0a, 0x0e, 0x4c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x6c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x6c, 0x65,
	0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a,
	0x45, 0x64, 0x69, 0x74, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x6c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x65,
	0x73, 0x73, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x13,
	0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x2e,
	0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x13, 0x2e, 0x6c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x4c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5a, 0x0a, 0x17, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x46, 0x6f, 0x72, 0x4c,
	0x65, 0x73, 0x73, 0x6f, 0x6e, 0x4f, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x1e, 0x2e,
	0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x4f, 0x72,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x4f, 0x72,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2b,
	0x5a, 0x29, 0x6e, 0x69, 0x63, 0x6b, 0x6f, 0x6c, 0x61, 0x79, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x73, 0x46, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x3b, 0x6c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x73, 0x46, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_lessons_lessons_proto_rawDescOnce sync.Once
	file_lessons_lessons_proto_rawDescData = file_lessons_lessons_proto_rawDesc
)

func file_lessons_lessons_proto_rawDescGZIP() []byte {
	file_lessons_lessons_proto_rawDescOnce.Do(func() {
		file_lessons_lessons_proto_rawDescData = protoimpl.X.CompressGZIP(file_lessons_lessons_proto_rawDescData)
	})
	return file_lessons_lessons_proto_rawDescData
}

var file_lessons_lessons_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_lessons_lessons_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_lessons_lessons_proto_goTypes = []interface{}{
	(Difficulty)(0),                // 0: lessons.Difficulty
	(*Lesson)(nil),                 // 1: lessons.Lesson
	(*CreateRequest)(nil),          // 2: lessons.CreateRequest
	(*CreateResponse)(nil),         // 3: lessons.CreateResponse
	(*DeleteRequest)(nil),          // 4: lessons.DeleteRequest
	(*DeleteResponse)(nil),         // 5: lessons.DeleteResponse
	(*EditRequest)(nil),            // 6: lessons.EditRequest
	(*EditResponse)(nil),           // 7: lessons.EditResponse
	(*GetRequest)(nil),             // 8: lessons.GetRequest
	(*GetResponse)(nil),            // 9: lessons.GetResponse
	(*CloseRequest)(nil),           // 10: lessons.CloseRequest
	(*CloseResponse)(nil),          // 11: lessons.CloseResponse
	(*SignUpOrCancelRequest)(nil),  // 12: lessons.SignUpOrCancelRequest
	(*SignUpOrCancelResponse)(nil), // 13: lessons.SignUpOrCancelResponse
}
var file_lessons_lessons_proto_depIdxs = []int32{
	0,  // 0: lessons.Lesson.difficulty:type_name -> lessons.Difficulty
	0,  // 1: lessons.CreateRequest.difficulty:type_name -> lessons.Difficulty
	1,  // 2: lessons.EditRequest.lesson:type_name -> lessons.Lesson
	1,  // 3: lessons.GetResponse.lessons:type_name -> lessons.Lesson
	2,  // 4: lessons.LessonsService.CreateLesson:input_type -> lessons.CreateRequest
	4,  // 5: lessons.LessonsService.DeleteLesson:input_type -> lessons.DeleteRequest
	6,  // 6: lessons.LessonsService.EditLesson:input_type -> lessons.EditRequest
	8,  // 7: lessons.LessonsService.GetLessonsByTrainer:input_type -> lessons.GetRequest
	8,  // 8: lessons.LessonsService.GetLessonsByUser:input_type -> lessons.GetRequest
	8,  // 9: lessons.LessonsService.GetAllLessons:input_type -> lessons.GetRequest
	10, // 10: lessons.LessonsService.CloseLesson:input_type -> lessons.CloseRequest
	12, // 11: lessons.LessonsService.SignUpForLessonOrCancel:input_type -> lessons.SignUpOrCancelRequest
	3,  // 12: lessons.LessonsService.CreateLesson:output_type -> lessons.CreateResponse
	5,  // 13: lessons.LessonsService.DeleteLesson:output_type -> lessons.DeleteResponse
	7,  // 14: lessons.LessonsService.EditLesson:output_type -> lessons.EditResponse
	9,  // 15: lessons.LessonsService.GetLessonsByTrainer:output_type -> lessons.GetResponse
	9,  // 16: lessons.LessonsService.GetLessonsByUser:output_type -> lessons.GetResponse
	9,  // 17: lessons.LessonsService.GetAllLessons:output_type -> lessons.GetResponse
	11, // 18: lessons.LessonsService.CloseLesson:output_type -> lessons.CloseResponse
	13, // 19: lessons.LessonsService.SignUpForLessonOrCancel:output_type -> lessons.SignUpOrCancelResponse
	12, // [12:20] is the sub-list for method output_type
	4,  // [4:12] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_lessons_lessons_proto_init() }
func file_lessons_lessons_proto_init() {
	if File_lessons_lessons_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lessons_lessons_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lesson); i {
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
		file_lessons_lessons_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_lessons_lessons_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_lessons_lessons_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_lessons_lessons_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_lessons_lessons_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_lessons_lessons_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_lessons_lessons_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_lessons_lessons_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_lessons_lessons_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseRequest); i {
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
		file_lessons_lessons_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseResponse); i {
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
		file_lessons_lessons_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpOrCancelRequest); i {
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
		file_lessons_lessons_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpOrCancelResponse); i {
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
			RawDescriptor: file_lessons_lessons_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lessons_lessons_proto_goTypes,
		DependencyIndexes: file_lessons_lessons_proto_depIdxs,
		EnumInfos:         file_lessons_lessons_proto_enumTypes,
		MessageInfos:      file_lessons_lessons_proto_msgTypes,
	}.Build()
	File_lessons_lessons_proto = out.File
	file_lessons_lessons_proto_rawDesc = nil
	file_lessons_lessons_proto_goTypes = nil
	file_lessons_lessons_proto_depIdxs = nil
}
