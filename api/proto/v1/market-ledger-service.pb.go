// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.3
// source: api/proto/v1/market-ledger-service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_market_ledger_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_market_ledger_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_market_ledger_service_proto_rawDescGZIP(), []int{0}
}

type NewInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique ID of the Issuer
	IssuerId uint64 `protobuf:"varint,1,opt,name=issuerId,proto3" json:"issuerId,omitempty"`
	// Unique invoice name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Full price of the Invoice
	FaceValue float64 `protobuf:"fixed64,3,opt,name=faceValue,proto3" json:"faceValue,omitempty"`
	// Price that needs to be covered, usually smaller but not greater than the rawValue
	NeededValue float64 `protobuf:"fixed64,4,opt,name=neededValue,proto3" json:"neededValue,omitempty"`
}

func (x *NewInvoiceRequest) Reset() {
	*x = NewInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_market_ledger_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewInvoiceRequest) ProtoMessage() {}

func (x *NewInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_market_ledger_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewInvoiceRequest.ProtoReflect.Descriptor instead.
func (*NewInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_market_ledger_service_proto_rawDescGZIP(), []int{1}
}

func (x *NewInvoiceRequest) GetIssuerId() uint64 {
	if x != nil {
		return x.IssuerId
	}
	return 0
}

func (x *NewInvoiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewInvoiceRequest) GetFaceValue() float64 {
	if x != nil {
		return x.FaceValue
	}
	return 0
}

func (x *NewInvoiceRequest) GetNeededValue() float64 {
	if x != nil {
		return x.NeededValue
	}
	return 0
}

type NewInvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the created Invoice
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NewInvoiceResponse) Reset() {
	*x = NewInvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_market_ledger_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewInvoiceResponse) ProtoMessage() {}

func (x *NewInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_market_ledger_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewInvoiceResponse.ProtoReflect.Descriptor instead.
func (*NewInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_market_ledger_service_proto_rawDescGZIP(), []int{2}
}

func (x *NewInvoiceResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type NewSellOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvoiceId uint64 `protobuf:"varint,1,opt,name=invoiceId,proto3" json:"invoiceId,omitempty"`
}

func (x *NewSellOrderRequest) Reset() {
	*x = NewSellOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_market_ledger_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewSellOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewSellOrderRequest) ProtoMessage() {}

func (x *NewSellOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_market_ledger_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewSellOrderRequest.ProtoReflect.Descriptor instead.
func (*NewSellOrderRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_market_ledger_service_proto_rawDescGZIP(), []int{3}
}

func (x *NewSellOrderRequest) GetInvoiceId() uint64 {
	if x != nil {
		return x.InvoiceId
	}
	return 0
}

type NewSellOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SellOrderId uint64 `protobuf:"varint,1,opt,name=sellOrderId,proto3" json:"sellOrderId,omitempty"`
}

func (x *NewSellOrderResponse) Reset() {
	*x = NewSellOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_market_ledger_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewSellOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewSellOrderResponse) ProtoMessage() {}

func (x *NewSellOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_market_ledger_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewSellOrderResponse.ProtoReflect.Descriptor instead.
func (*NewSellOrderResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_market_ledger_service_proto_rawDescGZIP(), []int{4}
}

func (x *NewSellOrderResponse) GetSellOrderId() uint64 {
	if x != nil {
		return x.SellOrderId
	}
	return 0
}

type NewBidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvestorId    uint64  `protobuf:"varint,1,opt,name=investorId,proto3" json:"investorId,omitempty"`
	SellOrderId   uint64  `protobuf:"varint,2,opt,name=sellOrderId,proto3" json:"sellOrderId,omitempty"`
	InvestedValue float64 `protobuf:"fixed64,3,opt,name=investedValue,proto3" json:"investedValue,omitempty"`
	Discount      float64 `protobuf:"fixed64,4,opt,name=discount,proto3" json:"discount,omitempty"`
}

func (x *NewBidRequest) Reset() {
	*x = NewBidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_market_ledger_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewBidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewBidRequest) ProtoMessage() {}

func (x *NewBidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_market_ledger_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewBidRequest.ProtoReflect.Descriptor instead.
func (*NewBidRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_market_ledger_service_proto_rawDescGZIP(), []int{5}
}

func (x *NewBidRequest) GetInvestorId() uint64 {
	if x != nil {
		return x.InvestorId
	}
	return 0
}

func (x *NewBidRequest) GetSellOrderId() uint64 {
	if x != nil {
		return x.SellOrderId
	}
	return 0
}

func (x *NewBidRequest) GetInvestedValue() float64 {
	if x != nil {
		return x.InvestedValue
	}
	return 0
}

func (x *NewBidRequest) GetDiscount() float64 {
	if x != nil {
		return x.Discount
	}
	return 0
}

type NewBidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position       uint64  `protobuf:"varint,1,opt,name=position,proto3" json:"position,omitempty"`
	ReservedValue  float64 `protobuf:"fixed64,2,opt,name=reservedValue,proto3" json:"reservedValue,omitempty"`
	ExpectedProfit float64 `protobuf:"fixed64,3,opt,name=expectedProfit,proto3" json:"expectedProfit,omitempty"`
}

func (x *NewBidResponse) Reset() {
	*x = NewBidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_market_ledger_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewBidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewBidResponse) ProtoMessage() {}

func (x *NewBidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_market_ledger_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewBidResponse.ProtoReflect.Descriptor instead.
func (*NewBidResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_market_ledger_service_proto_rawDescGZIP(), []int{6}
}

func (x *NewBidResponse) GetPosition() uint64 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *NewBidResponse) GetReservedValue() float64 {
	if x != nil {
		return x.ReservedValue
	}
	return 0
}

func (x *NewBidResponse) GetExpectedProfit() float64 {
	if x != nil {
		return x.ExpectedProfit
	}
	return 0
}

type MatchingAlgorithmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MatchingAlgorithmRequest) Reset() {
	*x = MatchingAlgorithmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_market_ledger_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchingAlgorithmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchingAlgorithmRequest) ProtoMessage() {}

func (x *MatchingAlgorithmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_market_ledger_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchingAlgorithmRequest.ProtoReflect.Descriptor instead.
func (*MatchingAlgorithmRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_market_ledger_service_proto_rawDescGZIP(), []int{7}
}

type MatchingAlgorithmResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MatchingAlgorithmResponse) Reset() {
	*x = MatchingAlgorithmResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_market_ledger_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchingAlgorithmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchingAlgorithmResponse) ProtoMessage() {}

func (x *MatchingAlgorithmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_market_ledger_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchingAlgorithmResponse.ProtoReflect.Descriptor instead.
func (*MatchingAlgorithmResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_market_ledger_service_proto_rawDescGZIP(), []int{8}
}

type LedgerEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvestorName    string  `protobuf:"bytes,1,opt,name=investorName,proto3" json:"investorName,omitempty"`
	InvoiceName     string  `protobuf:"bytes,2,opt,name=invoiceName,proto3" json:"invoiceName,omitempty"`
	InvestedBalance float64 `protobuf:"fixed64,3,opt,name=investedBalance,proto3" json:"investedBalance,omitempty"`
	ReservedBalance float64 `protobuf:"fixed64,4,opt,name=reservedBalance,proto3" json:"reservedBalance,omitempty"`
	ExpectedProfit  float64 `protobuf:"fixed64,5,opt,name=expectedProfit,proto3" json:"expectedProfit,omitempty"`
}

func (x *LedgerEntry) Reset() {
	*x = LedgerEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_market_ledger_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LedgerEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LedgerEntry) ProtoMessage() {}

func (x *LedgerEntry) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_market_ledger_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LedgerEntry.ProtoReflect.Descriptor instead.
func (*LedgerEntry) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_market_ledger_service_proto_rawDescGZIP(), []int{9}
}

func (x *LedgerEntry) GetInvestorName() string {
	if x != nil {
		return x.InvestorName
	}
	return ""
}

func (x *LedgerEntry) GetInvoiceName() string {
	if x != nil {
		return x.InvoiceName
	}
	return ""
}

func (x *LedgerEntry) GetInvestedBalance() float64 {
	if x != nil {
		return x.InvestedBalance
	}
	return 0
}

func (x *LedgerEntry) GetReservedBalance() float64 {
	if x != nil {
		return x.ReservedBalance
	}
	return 0
}

func (x *LedgerEntry) GetExpectedProfit() float64 {
	if x != nil {
		return x.ExpectedProfit
	}
	return 0
}

type LedgerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*LedgerEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *LedgerResponse) Reset() {
	*x = LedgerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_market_ledger_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LedgerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LedgerResponse) ProtoMessage() {}

func (x *LedgerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_market_ledger_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LedgerResponse.ProtoReflect.Descriptor instead.
func (*LedgerResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_market_ledger_service_proto_rawDescGZIP(), []int{10}
}

func (x *LedgerResponse) GetEntries() []*LedgerEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

var File_api_proto_v1_market_ledger_service_proto protoreflect.FileDescriptor

var file_api_proto_v1_market_ledger_service_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x83, 0x01, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66,
	0x61, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09,
	0x66, 0x61, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x65,
	0x64, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b,
	0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x4e,
	0x65, 0x77, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x33, 0x0a, 0x13, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x69, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x14, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x6c,
	0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x73, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x93, 0x01, 0x0a, 0x0d, 0x4e, 0x65, 0x77, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x69, 0x6e, 0x76,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7a, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x42, 0x69, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x78,
	0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x74, 0x22, 0x1a, 0x0a, 0x18, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x41, 0x6c,
	0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1b,
	0x0a, 0x19, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xcf, 0x01, 0x0a, 0x0b,
	0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x69,
	0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x65, 0x64, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x69, 0x6e, 0x76, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x65,
	0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x22, 0x3b, 0x0a,
	0x0e, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x29, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x32, 0xbb, 0x03, 0x0a, 0x0c, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x04, 0x53,
	0x65, 0x6c, 0x6c, 0x12, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x65, 0x77, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x1a, 0x08, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x6c, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x60, 0x0a, 0x0c, 0x4e, 0x65, 0x77, 0x53, 0x65,
	0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77,
	0x53, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x17, 0x1a, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x6c, 0x6c, 0x2d, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2f, 0x6e, 0x65, 0x77, 0x3a, 0x01, 0x2a, 0x12, 0x48, 0x0a, 0x06, 0x4e, 0x65, 0x77,
	0x42, 0x69, 0x64, 0x12, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77, 0x42, 0x69, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77, 0x42,
	0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x11, 0x1a, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x64, 0x73, 0x2f, 0x6e, 0x65, 0x77,
	0x3a, 0x01, 0x2a, 0x12, 0x73, 0x0a, 0x11, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x41,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x1a, 0x16, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2d, 0x61, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x3a, 0x01, 0x2a, 0x12, 0x3e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x42, 0x0e, 0x5a, 0x0c, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_v1_market_ledger_service_proto_rawDescOnce sync.Once
	file_api_proto_v1_market_ledger_service_proto_rawDescData = file_api_proto_v1_market_ledger_service_proto_rawDesc
)

func file_api_proto_v1_market_ledger_service_proto_rawDescGZIP() []byte {
	file_api_proto_v1_market_ledger_service_proto_rawDescOnce.Do(func() {
		file_api_proto_v1_market_ledger_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_v1_market_ledger_service_proto_rawDescData)
	})
	return file_api_proto_v1_market_ledger_service_proto_rawDescData
}

var file_api_proto_v1_market_ledger_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_proto_v1_market_ledger_service_proto_goTypes = []interface{}{
	(*Empty)(nil),                     // 0: v1.Empty
	(*NewInvoiceRequest)(nil),         // 1: v1.NewInvoiceRequest
	(*NewInvoiceResponse)(nil),        // 2: v1.NewInvoiceResponse
	(*NewSellOrderRequest)(nil),       // 3: v1.NewSellOrderRequest
	(*NewSellOrderResponse)(nil),      // 4: v1.NewSellOrderResponse
	(*NewBidRequest)(nil),             // 5: v1.NewBidRequest
	(*NewBidResponse)(nil),            // 6: v1.NewBidResponse
	(*MatchingAlgorithmRequest)(nil),  // 7: v1.MatchingAlgorithmRequest
	(*MatchingAlgorithmResponse)(nil), // 8: v1.MatchingAlgorithmResponse
	(*LedgerEntry)(nil),               // 9: v1.LedgerEntry
	(*LedgerResponse)(nil),            // 10: v1.LedgerResponse
}
var file_api_proto_v1_market_ledger_service_proto_depIdxs = []int32{
	9,  // 0: v1.LedgerResponse.entries:type_name -> v1.LedgerEntry
	1,  // 1: v1.MarketLedger.Sell:input_type -> v1.NewInvoiceRequest
	3,  // 2: v1.MarketLedger.NewSellOrder:input_type -> v1.NewSellOrderRequest
	5,  // 3: v1.MarketLedger.NewBid:input_type -> v1.NewBidRequest
	7,  // 4: v1.MarketLedger.MatchingAlgorithm:input_type -> v1.MatchingAlgorithmRequest
	0,  // 5: v1.MarketLedger.GetLedger:input_type -> v1.Empty
	2,  // 6: v1.MarketLedger.Sell:output_type -> v1.NewInvoiceResponse
	4,  // 7: v1.MarketLedger.NewSellOrder:output_type -> v1.NewSellOrderResponse
	6,  // 8: v1.MarketLedger.NewBid:output_type -> v1.NewBidResponse
	8,  // 9: v1.MarketLedger.MatchingAlgorithm:output_type -> v1.MatchingAlgorithmResponse
	10, // 10: v1.MarketLedger.GetLedger:output_type -> v1.LedgerResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_v1_market_ledger_service_proto_init() }
func file_api_proto_v1_market_ledger_service_proto_init() {
	if File_api_proto_v1_market_ledger_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_v1_market_ledger_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_api_proto_v1_market_ledger_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewInvoiceRequest); i {
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
		file_api_proto_v1_market_ledger_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewInvoiceResponse); i {
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
		file_api_proto_v1_market_ledger_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewSellOrderRequest); i {
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
		file_api_proto_v1_market_ledger_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewSellOrderResponse); i {
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
		file_api_proto_v1_market_ledger_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewBidRequest); i {
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
		file_api_proto_v1_market_ledger_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewBidResponse); i {
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
		file_api_proto_v1_market_ledger_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchingAlgorithmRequest); i {
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
		file_api_proto_v1_market_ledger_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchingAlgorithmResponse); i {
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
		file_api_proto_v1_market_ledger_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LedgerEntry); i {
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
		file_api_proto_v1_market_ledger_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LedgerResponse); i {
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
			RawDescriptor: file_api_proto_v1_market_ledger_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_v1_market_ledger_service_proto_goTypes,
		DependencyIndexes: file_api_proto_v1_market_ledger_service_proto_depIdxs,
		MessageInfos:      file_api_proto_v1_market_ledger_service_proto_msgTypes,
	}.Build()
	File_api_proto_v1_market_ledger_service_proto = out.File
	file_api_proto_v1_market_ledger_service_proto_rawDesc = nil
	file_api_proto_v1_market_ledger_service_proto_goTypes = nil
	file_api_proto_v1_market_ledger_service_proto_depIdxs = nil
}
