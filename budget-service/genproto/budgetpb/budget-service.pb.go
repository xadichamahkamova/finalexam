// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: proto/budget-service.proto

package budgetpb

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

type CreateBudgetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CategoryId string `protobuf:"bytes,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Amount     int64  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency   string `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *CreateBudgetRequest) Reset() {
	*x = CreateBudgetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_budget_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBudgetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBudgetRequest) ProtoMessage() {}

func (x *CreateBudgetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_budget_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBudgetRequest.ProtoReflect.Descriptor instead.
func (*CreateBudgetRequest) Descriptor() ([]byte, []int) {
	return file_proto_budget_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBudgetRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateBudgetRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *CreateBudgetRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateBudgetRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type CreateBudgetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message  string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	BudgetId string `protobuf:"bytes,2,opt,name=budget_id,json=budgetId,proto3" json:"budget_id,omitempty"`
}

func (x *CreateBudgetResponse) Reset() {
	*x = CreateBudgetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_budget_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBudgetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBudgetResponse) ProtoMessage() {}

func (x *CreateBudgetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_budget_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBudgetResponse.ProtoReflect.Descriptor instead.
func (*CreateBudgetResponse) Descriptor() ([]byte, []int) {
	return file_proto_budget_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBudgetResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateBudgetResponse) GetBudgetId() string {
	if x != nil {
		return x.BudgetId
	}
	return ""
}

type GetListBudgetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetListBudgetRequest) Reset() {
	*x = GetListBudgetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_budget_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListBudgetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListBudgetRequest) ProtoMessage() {}

func (x *GetListBudgetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_budget_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListBudgetRequest.ProtoReflect.Descriptor instead.
func (*GetListBudgetRequest) Descriptor() ([]byte, []int) {
	return file_proto_budget_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetListBudgetRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type Budget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BudgetId string `protobuf:"bytes,1,opt,name=budget_id,json=budgetId,proto3" json:"budget_id,omitempty"`
	Category string `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Amount   int64  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Spent    int64  `protobuf:"varint,4,opt,name=spent,proto3" json:"spent,omitempty"`
	Currency string `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *Budget) Reset() {
	*x = Budget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_budget_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Budget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Budget) ProtoMessage() {}

func (x *Budget) ProtoReflect() protoreflect.Message {
	mi := &file_proto_budget_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Budget.ProtoReflect.Descriptor instead.
func (*Budget) Descriptor() ([]byte, []int) {
	return file_proto_budget_service_proto_rawDescGZIP(), []int{3}
}

func (x *Budget) GetBudgetId() string {
	if x != nil {
		return x.BudgetId
	}
	return ""
}

func (x *Budget) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Budget) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Budget) GetSpent() int64 {
	if x != nil {
		return x.Spent
	}
	return 0
}

func (x *Budget) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type GetListBudgetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Budget `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *GetListBudgetResponse) Reset() {
	*x = GetListBudgetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_budget_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListBudgetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListBudgetResponse) ProtoMessage() {}

func (x *GetListBudgetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_budget_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListBudgetResponse.ProtoReflect.Descriptor instead.
func (*GetListBudgetResponse) Descriptor() ([]byte, []int) {
	return file_proto_budget_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetListBudgetResponse) GetList() []*Budget {
	if x != nil {
		return x.List
	}
	return nil
}

type UpdateBudgetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BudgetId string `protobuf:"bytes,2,opt,name=budget_id,json=budgetId,proto3" json:"budget_id,omitempty"`
	Amount   int64  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *UpdateBudgetRequest) Reset() {
	*x = UpdateBudgetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_budget_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBudgetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBudgetRequest) ProtoMessage() {}

func (x *UpdateBudgetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_budget_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBudgetRequest.ProtoReflect.Descriptor instead.
func (*UpdateBudgetRequest) Descriptor() ([]byte, []int) {
	return file_proto_budget_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBudgetRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateBudgetRequest) GetBudgetId() string {
	if x != nil {
		return x.BudgetId
	}
	return ""
}

func (x *UpdateBudgetRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type UpdateBudgetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateBudgetResponse) Reset() {
	*x = UpdateBudgetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_budget_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBudgetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBudgetResponse) ProtoMessage() {}

func (x *UpdateBudgetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_budget_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBudgetResponse.ProtoReflect.Descriptor instead.
func (*UpdateBudgetResponse) Descriptor() ([]byte, []int) {
	return file_proto_budget_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateBudgetResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_budget_service_proto protoreflect.FileDescriptor

var file_proto_budget_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x22, 0x4d, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x49,
	0x64, 0x22, 0x2f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x06, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x70, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73,
	0x70, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x22, 0x34, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x63, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x30, 0x0a, 0x14, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xc9, 0x01,
	0x0a, 0x0d, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12,
	0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x15, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_budget_service_proto_rawDescOnce sync.Once
	file_proto_budget_service_proto_rawDescData = file_proto_budget_service_proto_rawDesc
)

func file_proto_budget_service_proto_rawDescGZIP() []byte {
	file_proto_budget_service_proto_rawDescOnce.Do(func() {
		file_proto_budget_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_budget_service_proto_rawDescData)
	})
	return file_proto_budget_service_proto_rawDescData
}

var file_proto_budget_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_budget_service_proto_goTypes = []any{
	(*CreateBudgetRequest)(nil),   // 0: CreateBudgetRequest
	(*CreateBudgetResponse)(nil),  // 1: CreateBudgetResponse
	(*GetListBudgetRequest)(nil),  // 2: GetListBudgetRequest
	(*Budget)(nil),                // 3: Budget
	(*GetListBudgetResponse)(nil), // 4: GetListBudgetResponse
	(*UpdateBudgetRequest)(nil),   // 5: UpdateBudgetRequest
	(*UpdateBudgetResponse)(nil),  // 6: UpdateBudgetResponse
}
var file_proto_budget_service_proto_depIdxs = []int32{
	3, // 0: GetListBudgetResponse.list:type_name -> Budget
	0, // 1: BudgetService.CreateBudget:input_type -> CreateBudgetRequest
	2, // 2: BudgetService.GetListBudget:input_type -> GetListBudgetRequest
	5, // 3: BudgetService.UpdateBudget:input_type -> UpdateBudgetRequest
	1, // 4: BudgetService.CreateBudget:output_type -> CreateBudgetResponse
	4, // 5: BudgetService.GetListBudget:output_type -> GetListBudgetResponse
	6, // 6: BudgetService.UpdateBudget:output_type -> UpdateBudgetResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_budget_service_proto_init() }
func file_proto_budget_service_proto_init() {
	if File_proto_budget_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_budget_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateBudgetRequest); i {
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
		file_proto_budget_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateBudgetResponse); i {
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
		file_proto_budget_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetListBudgetRequest); i {
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
		file_proto_budget_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Budget); i {
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
		file_proto_budget_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetListBudgetResponse); i {
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
		file_proto_budget_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateBudgetRequest); i {
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
		file_proto_budget_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateBudgetResponse); i {
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
			RawDescriptor: file_proto_budget_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_budget_service_proto_goTypes,
		DependencyIndexes: file_proto_budget_service_proto_depIdxs,
		MessageInfos:      file_proto_budget_service_proto_msgTypes,
	}.Build()
	File_proto_budget_service_proto = out.File
	file_proto_budget_service_proto_rawDesc = nil
	file_proto_budget_service_proto_goTypes = nil
	file_proto_budget_service_proto_depIdxs = nil
}
