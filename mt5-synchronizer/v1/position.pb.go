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
// source: position.proto

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

// PositionType: Position Type
type PositionType int32

const (
	PositionType_POSITION_TYPE_BUY  PositionType = 0 //Buy
	PositionType_POSITION_TYPE_SELL PositionType = 1 //Sell
)

// Enum value maps for PositionType.
var (
	PositionType_name = map[int32]string{
		0: "POSITION_TYPE_BUY",
		1: "POSITION_TYPE_SELL",
	}
	PositionType_value = map[string]int32{
		"POSITION_TYPE_BUY":  0,
		"POSITION_TYPE_SELL": 1,
	}
)

func (x PositionType) Enum() *PositionType {
	p := new(PositionType)
	*p = x
	return p
}

func (x PositionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PositionType) Descriptor() protoreflect.EnumDescriptor {
	return file_position_proto_enumTypes[0].Descriptor()
}

func (PositionType) Type() protoreflect.EnumType {
	return &file_position_proto_enumTypes[0]
}

func (x PositionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PositionType.Descriptor instead.
func (PositionType) EnumDescriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{0}
}

// Position: Execution of trade operations results in the opening of a position
type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId        int32                  `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	AccountId      int32                  `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	PositionNumber int32                  `protobuf:"varint,3,opt,name=position_number,json=positionNumber,proto3" json:"position_number,omitempty"`
	Ticket         int32                  `protobuf:"varint,4,opt,name=ticket,proto3" json:"ticket,omitempty"`
	Symbol         string                 `protobuf:"bytes,5,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Time           *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=time,proto3" json:"time,omitempty"`
	PositionType   PositionType           `protobuf:"varint,7,opt,name=position_type,json=positionType,proto3,enum=mql5_background.v1.PositionType" json:"position_type,omitempty"`
	Volume         float64                `protobuf:"fixed64,8,opt,name=volume,proto3" json:"volume,omitempty"`
	PriceOpen      float64                `protobuf:"fixed64,9,opt,name=price_open,json=priceOpen,proto3" json:"price_open,omitempty"`
	StopLoss       float64                `protobuf:"fixed64,10,opt,name=stop_loss,json=stopLoss,proto3" json:"stop_loss,omitempty"`
	TakeProfit     float64                `protobuf:"fixed64,11,opt,name=take_profit,json=takeProfit,proto3" json:"take_profit,omitempty"`
	PriceCurrent   float64                `protobuf:"fixed64,12,opt,name=price_current,json=priceCurrent,proto3" json:"price_current,omitempty"`
	Commission     float64                `protobuf:"fixed64,13,opt,name=commission,proto3" json:"commission,omitempty"`
	Swap           float64                `protobuf:"fixed64,14,opt,name=swap,proto3" json:"swap,omitempty"`
	Profit         float64                `protobuf:"fixed64,15,opt,name=profit,proto3" json:"profit,omitempty"`
	Identifier     int32                  `protobuf:"varint,16,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Comment        string                 `protobuf:"bytes,17,opt,name=comment,proto3" json:"comment,omitempty"`
	Created        *timestamppb.Timestamp `protobuf:"bytes,18,opt,name=created,proto3" json:"created,omitempty"`
	Modified       *timestamppb.Timestamp `protobuf:"bytes,19,opt,name=modified,proto3" json:"modified,omitempty"`
	Deleted        *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_position_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{0}
}

func (x *Position) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *Position) GetAccountId() int32 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *Position) GetPositionNumber() int32 {
	if x != nil {
		return x.PositionNumber
	}
	return 0
}

func (x *Position) GetTicket() int32 {
	if x != nil {
		return x.Ticket
	}
	return 0
}

func (x *Position) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Position) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Position) GetPositionType() PositionType {
	if x != nil {
		return x.PositionType
	}
	return PositionType_POSITION_TYPE_BUY
}

func (x *Position) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *Position) GetPriceOpen() float64 {
	if x != nil {
		return x.PriceOpen
	}
	return 0
}

func (x *Position) GetStopLoss() float64 {
	if x != nil {
		return x.StopLoss
	}
	return 0
}

func (x *Position) GetTakeProfit() float64 {
	if x != nil {
		return x.TakeProfit
	}
	return 0
}

func (x *Position) GetPriceCurrent() float64 {
	if x != nil {
		return x.PriceCurrent
	}
	return 0
}

func (x *Position) GetCommission() float64 {
	if x != nil {
		return x.Commission
	}
	return 0
}

func (x *Position) GetSwap() float64 {
	if x != nil {
		return x.Swap
	}
	return 0
}

func (x *Position) GetProfit() float64 {
	if x != nil {
		return x.Profit
	}
	return 0
}

func (x *Position) GetIdentifier() int32 {
	if x != nil {
		return x.Identifier
	}
	return 0
}

func (x *Position) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Position) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Position) GetModified() *timestamppb.Timestamp {
	if x != nil {
		return x.Modified
	}
	return nil
}

func (x *Position) GetDeleted() *timestamppb.Timestamp {
	if x != nil {
		return x.Deleted
	}
	return nil
}

var File_position_proto protoreflect.FileDescriptor

var file_position_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x6d, 0x71, 0x6c, 0x35, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x05, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x0d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x6d,
	0x71, 0x6c, 0x35, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x6f, 0x70,
	0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x4f,
	0x70, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x6c, 0x6f, 0x73, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x70, 0x4c, 0x6f, 0x73, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x6b, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x61, 0x6b, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x63, 0x65, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x77, 0x61, 0x70, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x73, 0x77, 0x61, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x36, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x2a, 0x3d, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x15, 0x0a, 0x11, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x42, 0x55, 0x59, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x4f, 0x53, 0x49, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x45, 0x4c, 0x4c, 0x10, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_position_proto_rawDescOnce sync.Once
	file_position_proto_rawDescData = file_position_proto_rawDesc
)

func file_position_proto_rawDescGZIP() []byte {
	file_position_proto_rawDescOnce.Do(func() {
		file_position_proto_rawDescData = protoimpl.X.CompressGZIP(file_position_proto_rawDescData)
	})
	return file_position_proto_rawDescData
}

var file_position_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_position_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_position_proto_goTypes = []interface{}{
	(PositionType)(0),             // 0: mql5_background.v1.PositionType
	(*Position)(nil),              // 1: mql5_background.v1.Position
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_position_proto_depIdxs = []int32{
	2, // 0: mql5_background.v1.Position.time:type_name -> google.protobuf.Timestamp
	0, // 1: mql5_background.v1.Position.position_type:type_name -> mql5_background.v1.PositionType
	2, // 2: mql5_background.v1.Position.created:type_name -> google.protobuf.Timestamp
	2, // 3: mql5_background.v1.Position.modified:type_name -> google.protobuf.Timestamp
	2, // 4: mql5_background.v1.Position.deleted:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_position_proto_init() }
func file_position_proto_init() {
	if File_position_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_position_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
			RawDescriptor: file_position_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_position_proto_goTypes,
		DependencyIndexes: file_position_proto_depIdxs,
		EnumInfos:         file_position_proto_enumTypes,
		MessageInfos:      file_position_proto_msgTypes,
	}.Build()
	File_position_proto = out.File
	file_position_proto_rawDesc = nil
	file_position_proto_goTypes = nil
	file_position_proto_depIdxs = nil
}
