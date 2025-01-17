// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: proto/incexp.proto

package incexpb

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

type RegisterIncomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount     int64  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency   string `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	CategoryId string `protobuf:"bytes,4,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Date       string `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *RegisterIncomeRequest) Reset() {
	*x = RegisterIncomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_incexp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterIncomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterIncomeRequest) ProtoMessage() {}

func (x *RegisterIncomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_incexp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterIncomeRequest.ProtoReflect.Descriptor instead.
func (*RegisterIncomeRequest) Descriptor() ([]byte, []int) {
	return file_proto_incexp_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterIncomeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RegisterIncomeRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *RegisterIncomeRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *RegisterIncomeRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *RegisterIncomeRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type RegisterIncomeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message       string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	TransactionId string `protobuf:"bytes,2,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *RegisterIncomeResponse) Reset() {
	*x = RegisterIncomeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_incexp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterIncomeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterIncomeResponse) ProtoMessage() {}

func (x *RegisterIncomeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_incexp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterIncomeResponse.ProtoReflect.Descriptor instead.
func (*RegisterIncomeResponse) Descriptor() ([]byte, []int) {
	return file_proto_incexp_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterIncomeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RegisterIncomeResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

type RegisterExpenseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserEmail  string `protobuf:"bytes,2,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	Amount     int64  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency   string `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
	CategoryId string `protobuf:"bytes,5,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Date       string `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *RegisterExpenseRequest) Reset() {
	*x = RegisterExpenseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_incexp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterExpenseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterExpenseRequest) ProtoMessage() {}

func (x *RegisterExpenseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_incexp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterExpenseRequest.ProtoReflect.Descriptor instead.
func (*RegisterExpenseRequest) Descriptor() ([]byte, []int) {
	return file_proto_incexp_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterExpenseRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RegisterExpenseRequest) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *RegisterExpenseRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *RegisterExpenseRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *RegisterExpenseRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *RegisterExpenseRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type RegisterExpenseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message       string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	TransactionId string `protobuf:"bytes,2,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *RegisterExpenseResponse) Reset() {
	*x = RegisterExpenseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_incexp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterExpenseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterExpenseResponse) ProtoMessage() {}

func (x *RegisterExpenseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_incexp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterExpenseResponse.ProtoReflect.Descriptor instead.
func (*RegisterExpenseResponse) Descriptor() ([]byte, []int) {
	return file_proto_incexp_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterExpenseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RegisterExpenseResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

type IncomeVSExpense struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	Type          string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Amount        int64  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency      string `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
	Category      string `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	Date          string `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *IncomeVSExpense) Reset() {
	*x = IncomeVSExpense{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_incexp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncomeVSExpense) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncomeVSExpense) ProtoMessage() {}

func (x *IncomeVSExpense) ProtoReflect() protoreflect.Message {
	mi := &file_proto_incexp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncomeVSExpense.ProtoReflect.Descriptor instead.
func (*IncomeVSExpense) Descriptor() ([]byte, []int) {
	return file_proto_incexp_proto_rawDescGZIP(), []int{4}
}

func (x *IncomeVSExpense) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *IncomeVSExpense) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *IncomeVSExpense) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *IncomeVSExpense) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *IncomeVSExpense) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *IncomeVSExpense) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type GetListIncomeVSExpenseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetListIncomeVSExpenseRequest) Reset() {
	*x = GetListIncomeVSExpenseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_incexp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListIncomeVSExpenseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListIncomeVSExpenseRequest) ProtoMessage() {}

func (x *GetListIncomeVSExpenseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_incexp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListIncomeVSExpenseRequest.ProtoReflect.Descriptor instead.
func (*GetListIncomeVSExpenseRequest) Descriptor() ([]byte, []int) {
	return file_proto_incexp_proto_rawDescGZIP(), []int{5}
}

func (x *GetListIncomeVSExpenseRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetListIncomeVSExpenseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*IncomeVSExpense `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *GetListIncomeVSExpenseResponse) Reset() {
	*x = GetListIncomeVSExpenseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_incexp_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListIncomeVSExpenseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListIncomeVSExpenseResponse) ProtoMessage() {}

func (x *GetListIncomeVSExpenseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_incexp_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListIncomeVSExpenseResponse.ProtoReflect.Descriptor instead.
func (*GetListIncomeVSExpenseResponse) Descriptor() ([]byte, []int) {
	return file_proto_incexp_proto_rawDescGZIP(), []int{6}
}

func (x *GetListIncomeVSExpenseResponse) GetList() []*IncomeVSExpense {
	if x != nil {
		return x.List
	}
	return nil
}

var File_proto_incexp_proto protoreflect.FileDescriptor

var file_proto_incexp_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x63, 0x65, 0x78, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x22, 0x59, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x63, 0x6f,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xb9, 0x01, 0x0a, 0x16,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x5a, 0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0xb0, 0x01, 0x0a, 0x0f, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x56, 0x53,
	0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x38, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x56, 0x53, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x46, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x63, 0x6f, 0x6d,
	0x65, 0x56, 0x53, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x56, 0x53, 0x45, 0x78, 0x70, 0x65, 0x6e,
	0x73, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xf3, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x63,
	0x45, 0x78, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a,
	0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65,
	0x12, 0x17, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x65, 0x6e,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e,
	0x63, 0x6f, 0x6d, 0x65, 0x56, 0x53, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x56, 0x53, 0x45,
	0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x56, 0x53, 0x45,
	0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x14,
	0x5a, 0x12, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x63,
	0x65, 0x78, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_incexp_proto_rawDescOnce sync.Once
	file_proto_incexp_proto_rawDescData = file_proto_incexp_proto_rawDesc
)

func file_proto_incexp_proto_rawDescGZIP() []byte {
	file_proto_incexp_proto_rawDescOnce.Do(func() {
		file_proto_incexp_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_incexp_proto_rawDescData)
	})
	return file_proto_incexp_proto_rawDescData
}

var file_proto_incexp_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_incexp_proto_goTypes = []any{
	(*RegisterIncomeRequest)(nil),          // 0: RegisterIncomeRequest
	(*RegisterIncomeResponse)(nil),         // 1: RegisterIncomeResponse
	(*RegisterExpenseRequest)(nil),         // 2: RegisterExpenseRequest
	(*RegisterExpenseResponse)(nil),        // 3: RegisterExpenseResponse
	(*IncomeVSExpense)(nil),                // 4: IncomeVSExpense
	(*GetListIncomeVSExpenseRequest)(nil),  // 5: GetListIncomeVSExpenseRequest
	(*GetListIncomeVSExpenseResponse)(nil), // 6: GetListIncomeVSExpenseResponse
}
var file_proto_incexp_proto_depIdxs = []int32{
	4, // 0: GetListIncomeVSExpenseResponse.list:type_name -> IncomeVSExpense
	0, // 1: IncExpService.RegisterIncome:input_type -> RegisterIncomeRequest
	2, // 2: IncExpService.RegisterExpense:input_type -> RegisterExpenseRequest
	5, // 3: IncExpService.GetListIncomeVSExpense:input_type -> GetListIncomeVSExpenseRequest
	1, // 4: IncExpService.RegisterIncome:output_type -> RegisterIncomeResponse
	3, // 5: IncExpService.RegisterExpense:output_type -> RegisterExpenseResponse
	6, // 6: IncExpService.GetListIncomeVSExpense:output_type -> GetListIncomeVSExpenseResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_incexp_proto_init() }
func file_proto_incexp_proto_init() {
	if File_proto_incexp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_incexp_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterIncomeRequest); i {
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
		file_proto_incexp_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterIncomeResponse); i {
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
		file_proto_incexp_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterExpenseRequest); i {
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
		file_proto_incexp_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterExpenseResponse); i {
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
		file_proto_incexp_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*IncomeVSExpense); i {
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
		file_proto_incexp_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetListIncomeVSExpenseRequest); i {
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
		file_proto_incexp_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetListIncomeVSExpenseResponse); i {
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
			RawDescriptor: file_proto_incexp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_incexp_proto_goTypes,
		DependencyIndexes: file_proto_incexp_proto_depIdxs,
		MessageInfos:      file_proto_incexp_proto_msgTypes,
	}.Build()
	File_proto_incexp_proto = out.File
	file_proto_incexp_proto_rawDesc = nil
	file_proto_incexp_proto_goTypes = nil
	file_proto_incexp_proto_depIdxs = nil
}
