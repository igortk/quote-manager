// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: message/balance_component.proto

package proto

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

type EmmitBalanceByUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId   string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Currency string  `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount   float64 `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *EmmitBalanceByUserIdRequest) Reset() {
	*x = EmmitBalanceByUserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_balance_component_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmmitBalanceByUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmmitBalanceByUserIdRequest) ProtoMessage() {}

func (x *EmmitBalanceByUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_balance_component_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmmitBalanceByUserIdRequest.ProtoReflect.Descriptor instead.
func (*EmmitBalanceByUserIdRequest) Descriptor() ([]byte, []int) {
	return file_message_balance_component_proto_rawDescGZIP(), []int{0}
}

func (x *EmmitBalanceByUserIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EmmitBalanceByUserIdRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *EmmitBalanceByUserIdRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *EmmitBalanceByUserIdRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type GetBalanceByUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetBalanceByUserIdRequest) Reset() {
	*x = GetBalanceByUserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_balance_component_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceByUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceByUserIdRequest) ProtoMessage() {}

func (x *GetBalanceByUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_balance_component_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceByUserIdRequest.ProtoReflect.Descriptor instead.
func (*GetBalanceByUserIdRequest) Descriptor() ([]byte, []int) {
	return file_message_balance_component_proto_rawDescGZIP(), []int{1}
}

func (x *GetBalanceByUserIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetBalanceByUserIdRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetAllPairRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAllPairRequest) Reset() {
	*x = GetAllPairRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_balance_component_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPairRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPairRequest) ProtoMessage() {}

func (x *GetAllPairRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_balance_component_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPairRequest.ProtoReflect.Descriptor instead.
func (*GetAllPairRequest) Descriptor() ([]byte, []int) {
	return file_message_balance_component_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllPairRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAllPairResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Pairs []string `protobuf:"bytes,2,rep,name=pairs,proto3" json:"pairs,omitempty"`
}

func (x *GetAllPairResponse) Reset() {
	*x = GetAllPairResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_balance_component_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPairResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPairResponse) ProtoMessage() {}

func (x *GetAllPairResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_balance_component_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPairResponse.ProtoReflect.Descriptor instead.
func (*GetAllPairResponse) Descriptor() ([]byte, []int) {
	return file_message_balance_component_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllPairResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetAllPairResponse) GetPairs() []string {
	if x != nil {
		return x.Pairs
	}
	return nil
}

type GetBalanceByUserIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserBalance *UserBalance `protobuf:"bytes,2,opt,name=user_balance,json=userBalance,proto3" json:"user_balance,omitempty"`
	Error       *Error       `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetBalanceByUserIdResponse) Reset() {
	*x = GetBalanceByUserIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_balance_component_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceByUserIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceByUserIdResponse) ProtoMessage() {}

func (x *GetBalanceByUserIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_balance_component_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceByUserIdResponse.ProtoReflect.Descriptor instead.
func (*GetBalanceByUserIdResponse) Descriptor() ([]byte, []int) {
	return file_message_balance_component_proto_rawDescGZIP(), []int{4}
}

func (x *GetBalanceByUserIdResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetBalanceByUserIdResponse) GetUserBalance() *UserBalance {
	if x != nil {
		return x.UserBalance
	}
	return nil
}

func (x *GetBalanceByUserIdResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type UserBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string     `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Balances []*Balance `protobuf:"bytes,2,rep,name=balances,proto3" json:"balances,omitempty"`
}

func (x *UserBalance) Reset() {
	*x = UserBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_balance_component_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBalance) ProtoMessage() {}

func (x *UserBalance) ProtoReflect() protoreflect.Message {
	mi := &file_message_balance_component_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBalance.ProtoReflect.Descriptor instead.
func (*UserBalance) Descriptor() ([]byte, []int) {
	return file_message_balance_component_proto_rawDescGZIP(), []int{5}
}

func (x *UserBalance) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserBalance) GetBalances() []*Balance {
	if x != nil {
		return x.Balances
	}
	return nil
}

type Balance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currency      string  `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	Balance       float64 `protobuf:"fixed64,2,opt,name=balance,proto3" json:"balance,omitempty"`
	LockedBalance float64 `protobuf:"fixed64,3,opt,name=locked_balance,json=lockedBalance,proto3" json:"locked_balance,omitempty"`
	UpdatedDate   int64   `protobuf:"varint,4,opt,name=updated_date,json=updatedDate,proto3" json:"updated_date,omitempty"`
}

func (x *Balance) Reset() {
	*x = Balance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_balance_component_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Balance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Balance) ProtoMessage() {}

func (x *Balance) ProtoReflect() protoreflect.Message {
	mi := &file_message_balance_component_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Balance.ProtoReflect.Descriptor instead.
func (*Balance) Descriptor() ([]byte, []int) {
	return file_message_balance_component_proto_rawDescGZIP(), []int{6}
}

func (x *Balance) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Balance) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *Balance) GetLockedBalance() float64 {
	if x != nil {
		return x.LockedBalance
	}
	return 0
}

func (x *Balance) GetUpdatedDate() int64 {
	if x != nil {
		return x.UpdatedDate
	}
	return 0
}

var File_message_balance_component_proto protoreflect.FileDescriptor

var file_message_balance_component_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x62, 0x73, 0x1a, 0x13, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x1b, 0x45, 0x6d,
	0x6d, 0x69, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x44, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x23, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x69, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x3a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x69, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x69, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x69, 0x72, 0x73, 0x22, 0x84, 0x01,
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x22, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x4f, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x08,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x62, 0x73, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0d, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_balance_component_proto_rawDescOnce sync.Once
	file_message_balance_component_proto_rawDescData = file_message_balance_component_proto_rawDesc
)

func file_message_balance_component_proto_rawDescGZIP() []byte {
	file_message_balance_component_proto_rawDescOnce.Do(func() {
		file_message_balance_component_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_balance_component_proto_rawDescData)
	})
	return file_message_balance_component_proto_rawDescData
}

var file_message_balance_component_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_message_balance_component_proto_goTypes = []interface{}{
	(*EmmitBalanceByUserIdRequest)(nil), // 0: bs.EmmitBalanceByUserIdRequest
	(*GetBalanceByUserIdRequest)(nil),   // 1: bs.GetBalanceByUserIdRequest
	(*GetAllPairRequest)(nil),           // 2: bs.GetAllPairRequest
	(*GetAllPairResponse)(nil),          // 3: bs.GetAllPairResponse
	(*GetBalanceByUserIdResponse)(nil),  // 4: bs.GetBalanceByUserIdResponse
	(*UserBalance)(nil),                 // 5: bs.UserBalance
	(*Balance)(nil),                     // 6: bs.Balance
	(*Error)(nil),                       // 7: error.Error
}
var file_message_balance_component_proto_depIdxs = []int32{
	5, // 0: bs.GetBalanceByUserIdResponse.user_balance:type_name -> bs.UserBalance
	7, // 1: bs.GetBalanceByUserIdResponse.error:type_name -> error.Error
	6, // 2: bs.UserBalance.balances:type_name -> bs.Balance
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_message_balance_component_proto_init() }
func file_message_balance_component_proto_init() {
	if File_message_balance_component_proto != nil {
		return
	}
	file_message_error_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_message_balance_component_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmmitBalanceByUserIdRequest); i {
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
		file_message_balance_component_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceByUserIdRequest); i {
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
		file_message_balance_component_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPairRequest); i {
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
		file_message_balance_component_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPairResponse); i {
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
		file_message_balance_component_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceByUserIdResponse); i {
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
		file_message_balance_component_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBalance); i {
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
		file_message_balance_component_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Balance); i {
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
			RawDescriptor: file_message_balance_component_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_balance_component_proto_goTypes,
		DependencyIndexes: file_message_balance_component_proto_depIdxs,
		MessageInfos:      file_message_balance_component_proto_msgTypes,
	}.Build()
	File_message_balance_component_proto = out.File
	file_message_balance_component_proto_rawDesc = nil
	file_message_balance_component_proto_goTypes = nil
	file_message_balance_component_proto_depIdxs = nil
}
