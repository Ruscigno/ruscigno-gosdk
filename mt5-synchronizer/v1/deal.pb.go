// Copyright 2021 Sander Ruscigno
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: deal.proto

package mql5_background_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// DealType: Deal type
type DealType int32

const (
	DealType_DEAL_TYPE_BUY                      DealType = 0  //Buy
	DealType_DEAL_TYPE_SELL                     DealType = 1  //Sell
	DealType_DEAL_TYPE_BALANCE                  DealType = 2  //Balance
	DealType_DEAL_TYPE_CREDIT                   DealType = 3  //Credit
	DealType_DEAL_TYPE_CHARGE                   DealType = 4  //Additional charge
	DealType_DEAL_TYPE_CORRECTION               DealType = 5  //Correction
	DealType_DEAL_TYPE_BONUS                    DealType = 6  //Bonus
	DealType_DEAL_TYPE_COMMISSION               DealType = 7  //Additional commission
	DealType_DEAL_TYPE_COMMISSION_DAILY         DealType = 8  //Daily commission
	DealType_DEAL_TYPE_COMMISSION_MONTHLY       DealType = 9  //Monthly commission
	DealType_DEAL_TYPE_COMMISSION_AGENT_DAILY   DealType = 10 //Daily agent commission
	DealType_DEAL_TYPE_COMMISSION_AGENT_MONTHLY DealType = 11 //Monthly agent commission
	DealType_DEAL_TYPE_INTEREST                 DealType = 12 //Interest rate
	DealType_DEAL_TYPE_BUY_CANCELED             DealType = 13 //Canceled buy deal.
	DealType_DEAL_TYPE_SELL_CANCELED            DealType = 14 //Canceled sell deal.
	DealType_DEAL_DIVIDEND                      DealType = 15 //Dividend operations
	DealType_DEAL_DIVIDEND_FRANKED              DealType = 16 //Franked (non-taxable) dividend operations
	DealType_DEAL_TAX                           DealType = 17 //Tax charges
)

// Enum value maps for DealType.
var (
	DealType_name = map[int32]string{
		0:  "DEAL_TYPE_BUY",
		1:  "DEAL_TYPE_SELL",
		2:  "DEAL_TYPE_BALANCE",
		3:  "DEAL_TYPE_CREDIT",
		4:  "DEAL_TYPE_CHARGE",
		5:  "DEAL_TYPE_CORRECTION",
		6:  "DEAL_TYPE_BONUS",
		7:  "DEAL_TYPE_COMMISSION",
		8:  "DEAL_TYPE_COMMISSION_DAILY",
		9:  "DEAL_TYPE_COMMISSION_MONTHLY",
		10: "DEAL_TYPE_COMMISSION_AGENT_DAILY",
		11: "DEAL_TYPE_COMMISSION_AGENT_MONTHLY",
		12: "DEAL_TYPE_INTEREST",
		13: "DEAL_TYPE_BUY_CANCELED",
		14: "DEAL_TYPE_SELL_CANCELED",
		15: "DEAL_DIVIDEND",
		16: "DEAL_DIVIDEND_FRANKED",
		17: "DEAL_TAX",
	}
	DealType_value = map[string]int32{
		"DEAL_TYPE_BUY":                      0,
		"DEAL_TYPE_SELL":                     1,
		"DEAL_TYPE_BALANCE":                  2,
		"DEAL_TYPE_CREDIT":                   3,
		"DEAL_TYPE_CHARGE":                   4,
		"DEAL_TYPE_CORRECTION":               5,
		"DEAL_TYPE_BONUS":                    6,
		"DEAL_TYPE_COMMISSION":               7,
		"DEAL_TYPE_COMMISSION_DAILY":         8,
		"DEAL_TYPE_COMMISSION_MONTHLY":       9,
		"DEAL_TYPE_COMMISSION_AGENT_DAILY":   10,
		"DEAL_TYPE_COMMISSION_AGENT_MONTHLY": 11,
		"DEAL_TYPE_INTEREST":                 12,
		"DEAL_TYPE_BUY_CANCELED":             13,
		"DEAL_TYPE_SELL_CANCELED":            14,
		"DEAL_DIVIDEND":                      15,
		"DEAL_DIVIDEND_FRANKED":              16,
		"DEAL_TAX":                           17,
	}
)

func (x DealType) Enum() *DealType {
	p := new(DealType)
	*p = x
	return p
}

func (x DealType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DealType) Descriptor() protoreflect.EnumDescriptor {
	return file_deal_proto_enumTypes[0].Descriptor()
}

func (DealType) Type() protoreflect.EnumType {
	return &file_deal_proto_enumTypes[0]
}

func (x DealType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DealType.Descriptor instead.
func (DealType) EnumDescriptor() ([]byte, []int) {
	return file_deal_proto_rawDescGZIP(), []int{0}
}

// DealEntry: Deal entry - entry in, entry out, reverse
type DealEntry int32

const (
	DealEntry_DEAL_ENTRY_IN     DealEntry = 0 //Entry in
	DealEntry_DEAL_ENTRY_OUT    DealEntry = 1 //Entry out
	DealEntry_DEAL_ENTRY_INOUT  DealEntry = 2 //Reverse
	DealEntry_DEAL_ENTRY_OUT_BY DealEntry = 3 //Close a position by an opposite one
)

// Enum value maps for DealEntry.
var (
	DealEntry_name = map[int32]string{
		0: "DEAL_ENTRY_IN",
		1: "DEAL_ENTRY_OUT",
		2: "DEAL_ENTRY_INOUT",
		3: "DEAL_ENTRY_OUT_BY",
	}
	DealEntry_value = map[string]int32{
		"DEAL_ENTRY_IN":     0,
		"DEAL_ENTRY_OUT":    1,
		"DEAL_ENTRY_INOUT":  2,
		"DEAL_ENTRY_OUT_BY": 3,
	}
)

func (x DealEntry) Enum() *DealEntry {
	p := new(DealEntry)
	*p = x
	return p
}

func (x DealEntry) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DealEntry) Descriptor() protoreflect.EnumDescriptor {
	return file_deal_proto_enumTypes[1].Descriptor()
}

func (DealEntry) Type() protoreflect.EnumType {
	return &file_deal_proto_enumTypes[1]
}

func (x DealEntry) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DealEntry.Descriptor instead.
func (DealEntry) EnumDescriptor() ([]byte, []int) {
	return file_deal_proto_rawDescGZIP(), []int{1}
}

// Deal: A deal is the reflection of the fact of a trade operation execution based on an order that contains a trade request
type Deal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DealId     int32                  `protobuf:"varint,1,opt,name=deal_id,json=dealId,proto3" json:"deal_id,omitempty"`
	AccountId  int32                  `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Ticket     int32                  `protobuf:"varint,3,opt,name=ticket,proto3" json:"ticket,omitempty"`
	Magic      int32                  `protobuf:"varint,4,opt,name=magic,proto3" json:"magic,omitempty"`
	Order      int32                  `protobuf:"varint,5,opt,name=order,proto3" json:"order,omitempty"`
	Symbol     string                 `protobuf:"bytes,6,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Time       int64                  `protobuf:"varint,7,opt,name=time,proto3" json:"time,omitempty"`
	DealType   DealType               `protobuf:"varint,8,opt,name=deal_type,json=dealType,proto3,enum=mql5_background.v1.DealType" json:"deal_type,omitempty"`
	Entry      DealEntry              `protobuf:"varint,9,opt,name=entry,proto3,enum=mql5_background.v1.DealEntry" json:"entry,omitempty"`
	PositionId int32                  `protobuf:"varint,10,opt,name=position_id,json=positionId,proto3" json:"position_id,omitempty"`
	Volume     float64                `protobuf:"fixed64,11,opt,name=volume,proto3" json:"volume,omitempty"`
	Price      float64                `protobuf:"fixed64,12,opt,name=price,proto3" json:"price,omitempty"`
	Commission float64                `protobuf:"fixed64,13,opt,name=commission,proto3" json:"commission,omitempty"`
	Swap       float64                `protobuf:"fixed64,14,opt,name=swap,proto3" json:"swap,omitempty"`
	Profit     float64                `protobuf:"fixed64,15,opt,name=profit,proto3" json:"profit,omitempty"`
	Comment    string                 `protobuf:"bytes,16,opt,name=comment,proto3" json:"comment,omitempty"`
	ExternalId string                 `protobuf:"bytes,17,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	Created    *timestamppb.Timestamp `protobuf:"bytes,18,opt,name=created,proto3" json:"created,omitempty"`
	Updated    *timestamppb.Timestamp `protobuf:"bytes,19,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted    *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *Deal) Reset() {
	*x = Deal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deal) ProtoMessage() {}

func (x *Deal) ProtoReflect() protoreflect.Message {
	mi := &file_deal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deal.ProtoReflect.Descriptor instead.
func (*Deal) Descriptor() ([]byte, []int) {
	return file_deal_proto_rawDescGZIP(), []int{0}
}

func (x *Deal) GetDealId() int32 {
	if x != nil {
		return x.DealId
	}
	return 0
}

func (x *Deal) GetAccountId() int32 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *Deal) GetTicket() int32 {
	if x != nil {
		return x.Ticket
	}
	return 0
}

func (x *Deal) GetMagic() int32 {
	if x != nil {
		return x.Magic
	}
	return 0
}

func (x *Deal) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *Deal) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Deal) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Deal) GetDealType() DealType {
	if x != nil {
		return x.DealType
	}
	return DealType_DEAL_TYPE_BUY
}

func (x *Deal) GetEntry() DealEntry {
	if x != nil {
		return x.Entry
	}
	return DealEntry_DEAL_ENTRY_IN
}

func (x *Deal) GetPositionId() int32 {
	if x != nil {
		return x.PositionId
	}
	return 0
}

func (x *Deal) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *Deal) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Deal) GetCommission() float64 {
	if x != nil {
		return x.Commission
	}
	return 0
}

func (x *Deal) GetSwap() float64 {
	if x != nil {
		return x.Swap
	}
	return 0
}

func (x *Deal) GetProfit() float64 {
	if x != nil {
		return x.Profit
	}
	return 0
}

func (x *Deal) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Deal) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *Deal) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Deal) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *Deal) GetDeleted() *timestamppb.Timestamp {
	if x != nil {
		return x.Deleted
	}
	return nil
}

var File_deal_proto protoreflect.FileDescriptor

var file_deal_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x65, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6d, 0x71,
	0x6c, 0x35, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x96, 0x05, 0x0a, 0x04, 0x44, 0x65, 0x61, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x65,
	0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x64, 0x65, 0x61,
	0x6c, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61,
	0x67, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x61, 0x67, 0x69, 0x63,
	0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x64, 0x65, 0x61, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6d, 0x71, 0x6c, 0x35, 0x5f, 0x62, 0x61, 0x63,
	0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x61, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x08, 0x64, 0x65, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x33, 0x0a,
	0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6d,
	0x71, 0x6c, 0x35, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x77, 0x61, 0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x04, 0x73, 0x77, 0x61, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34,
	0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x2a, 0xda, 0x03, 0x0a, 0x08, 0x44,
	0x65, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x45, 0x41, 0x4c, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x55, 0x59, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x45,
	0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x45, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x15,
	0x0a, 0x11, 0x44, 0x45, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x41, 0x4c, 0x41,
	0x4e, 0x43, 0x45, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x45, 0x41, 0x4c, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x44, 0x49, 0x54, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x44,
	0x45, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x47, 0x45, 0x10,
	0x04, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x45, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43,
	0x4f, 0x52, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x44,
	0x45, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4e, 0x55, 0x53, 0x10, 0x06,
	0x12, 0x18, 0x0a, 0x14, 0x44, 0x45, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f,
	0x4d, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x12, 0x1e, 0x0a, 0x1a, 0x44, 0x45,
	0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x44, 0x41, 0x49, 0x4c, 0x59, 0x10, 0x08, 0x12, 0x20, 0x0a, 0x1c, 0x44, 0x45,
	0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x4c, 0x59, 0x10, 0x09, 0x12, 0x24, 0x0a, 0x20,
	0x44, 0x45, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x41, 0x49, 0x4c, 0x59,
	0x10, 0x0a, 0x12, 0x26, 0x0a, 0x22, 0x44, 0x45, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54,
	0x5f, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x4c, 0x59, 0x10, 0x0b, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x45,
	0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x45, 0x53, 0x54,
	0x10, 0x0c, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x45, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x42, 0x55, 0x59, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x0d, 0x12, 0x1b,
	0x0a, 0x17, 0x44, 0x45, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x45, 0x4c, 0x4c,
	0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x0e, 0x12, 0x11, 0x0a, 0x0d, 0x44,
	0x45, 0x41, 0x4c, 0x5f, 0x44, 0x49, 0x56, 0x49, 0x44, 0x45, 0x4e, 0x44, 0x10, 0x0f, 0x12, 0x19,
	0x0a, 0x15, 0x44, 0x45, 0x41, 0x4c, 0x5f, 0x44, 0x49, 0x56, 0x49, 0x44, 0x45, 0x4e, 0x44, 0x5f,
	0x46, 0x52, 0x41, 0x4e, 0x4b, 0x45, 0x44, 0x10, 0x10, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x41,
	0x4c, 0x5f, 0x54, 0x41, 0x58, 0x10, 0x11, 0x2a, 0x5f, 0x0a, 0x09, 0x44, 0x65, 0x61, 0x6c, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x45, 0x41, 0x4c, 0x5f, 0x45, 0x4e, 0x54,
	0x52, 0x59, 0x5f, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x45, 0x41, 0x4c, 0x5f,
	0x45, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x4f, 0x55, 0x54, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x44,
	0x45, 0x41, 0x4c, 0x5f, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x49, 0x4e, 0x4f, 0x55, 0x54, 0x10,
	0x02, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x45, 0x41, 0x4c, 0x5f, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x5f,
	0x4f, 0x55, 0x54, 0x5f, 0x42, 0x59, 0x10, 0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deal_proto_rawDescOnce sync.Once
	file_deal_proto_rawDescData = file_deal_proto_rawDesc
)

func file_deal_proto_rawDescGZIP() []byte {
	file_deal_proto_rawDescOnce.Do(func() {
		file_deal_proto_rawDescData = protoimpl.X.CompressGZIP(file_deal_proto_rawDescData)
	})
	return file_deal_proto_rawDescData
}

var file_deal_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_deal_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_deal_proto_goTypes = []interface{}{
	(DealType)(0),                 // 0: mql5_background.v1.DealType
	(DealEntry)(0),                // 1: mql5_background.v1.DealEntry
	(*Deal)(nil),                  // 2: mql5_background.v1.Deal
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_deal_proto_depIdxs = []int32{
	0, // 0: mql5_background.v1.Deal.deal_type:type_name -> mql5_background.v1.DealType
	1, // 1: mql5_background.v1.Deal.entry:type_name -> mql5_background.v1.DealEntry
	3, // 2: mql5_background.v1.Deal.created:type_name -> google.protobuf.Timestamp
	3, // 3: mql5_background.v1.Deal.updated:type_name -> google.protobuf.Timestamp
	3, // 4: mql5_background.v1.Deal.deleted:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_deal_proto_init() }
func file_deal_proto_init() {
	if File_deal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deal); i {
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
			RawDescriptor: file_deal_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_deal_proto_goTypes,
		DependencyIndexes: file_deal_proto_depIdxs,
		EnumInfos:         file_deal_proto_enumTypes,
		MessageInfos:      file_deal_proto_msgTypes,
	}.Build()
	File_deal_proto = out.File
	file_deal_proto_rawDesc = nil
	file_deal_proto_goTypes = nil
	file_deal_proto_depIdxs = nil
}
