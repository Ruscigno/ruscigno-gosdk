// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: signal.proto

package v1

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

// SignalType: type of the signal
type SignalType int32

const (
	SignalType_SIGINAL_TYPE_GET_ALL       SignalType = 0 // Return all types of signal
	SignalType_SIGINAL_TYPE_ORDER         SignalType = 1 // Return all unsynced orders
	SignalType_SIGINAL_TYPE_POSITION      SignalType = 2 // Return all unsynced positions
	SignalType_SIGINAL_TYPE_TRADE_REQUEST SignalType = 3 // Return all pending trades
	SignalType_SIGINAL_TYPE_DEAL          SignalType = 4 // Return all unsynced deals
	SignalType_SIGINAL_TYPE_STOP_TAKE     SignalType = 5 // Return all Stop Loss or Take Profit changes
)

// Enum value maps for SignalType.
var (
	SignalType_name = map[int32]string{
		0: "SIGINAL_TYPE_GET_ALL",
		1: "SIGINAL_TYPE_ORDER",
		2: "SIGINAL_TYPE_POSITION",
		3: "SIGINAL_TYPE_TRADE_REQUEST",
		4: "SIGINAL_TYPE_DEAL",
		5: "SIGINAL_TYPE_STOP_TAKE",
	}
	SignalType_value = map[string]int32{
		"SIGINAL_TYPE_GET_ALL":       0,
		"SIGINAL_TYPE_ORDER":         1,
		"SIGINAL_TYPE_POSITION":      2,
		"SIGINAL_TYPE_TRADE_REQUEST": 3,
		"SIGINAL_TYPE_DEAL":          4,
		"SIGINAL_TYPE_STOP_TAKE":     5,
	}
)

func (x SignalType) Enum() *SignalType {
	p := new(SignalType)
	*p = x
	return p
}

func (x SignalType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SignalType) Descriptor() protoreflect.EnumDescriptor {
	return file_signal_proto_enumTypes[0].Descriptor()
}

func (SignalType) Type() protoreflect.EnumType {
	return &file_signal_proto_enumTypes[0]
}

func (x SignalType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SignalType.Descriptor instead.
func (SignalType) EnumDescriptor() ([]byte, []int) {
	return file_signal_proto_rawDescGZIP(), []int{0}
}

type Signal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignalId             int64 `protobuf:"varint,1,opt,name=SignalId,proto3" json:"SignalId,omitempty"`
	SourceAccountId      int64 `protobuf:"varint,2,opt,name=SourceAccountId,proto3" json:"SourceAccountId,omitempty"`
	DestinationAccountId int64 `protobuf:"varint,3,opt,name=DestinationAccountId,proto3" json:"DestinationAccountId,omitempty"`
	Active               bool  `protobuf:"varint,4,opt,name=Active,proto3" json:"Active,omitempty"`
	MaxDeposit           int32 `protobuf:"varint,5,opt,name=MaxDeposit,proto3" json:"MaxDeposit,omitempty"`
	StopIfLessThan       int32 `protobuf:"varint,6,opt,name=StopIfLessThan,proto3" json:"StopIfLessThan,omitempty"`
	Deviation            int64 `protobuf:"varint,7,opt,name=Deviation,proto3" json:"Deviation,omitempty"`
	MinutesToExpire      int32 `protobuf:"varint,8,opt,name=MinutesToExpire,proto3" json:"MinutesToExpire,omitempty"`
}

func (x *Signal) Reset() {
	*x = Signal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signal) ProtoMessage() {}

func (x *Signal) ProtoReflect() protoreflect.Message {
	mi := &file_signal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signal.ProtoReflect.Descriptor instead.
func (*Signal) Descriptor() ([]byte, []int) {
	return file_signal_proto_rawDescGZIP(), []int{0}
}

func (x *Signal) GetSignalId() int64 {
	if x != nil {
		return x.SignalId
	}
	return 0
}

func (x *Signal) GetSourceAccountId() int64 {
	if x != nil {
		return x.SourceAccountId
	}
	return 0
}

func (x *Signal) GetDestinationAccountId() int64 {
	if x != nil {
		return x.DestinationAccountId
	}
	return 0
}

func (x *Signal) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Signal) GetMaxDeposit() int32 {
	if x != nil {
		return x.MaxDeposit
	}
	return 0
}

func (x *Signal) GetStopIfLessThan() int32 {
	if x != nil {
		return x.StopIfLessThan
	}
	return 0
}

func (x *Signal) GetDeviation() int64 {
	if x != nil {
		return x.Deviation
	}
	return 0
}

func (x *Signal) GetMinutesToExpire() int32 {
	if x != nil {
		return x.MinutesToExpire
	}
	return 0
}

type SignalResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignalResultId       int64 `protobuf:"varint,1,opt,name=SignalResultId,proto3" json:"SignalResultId,omitempty"`
	SourceAccountId      int64 `protobuf:"varint,2,opt,name=SourceAccountId,proto3" json:"SourceAccountId,omitempty"`
	DestinationAccountId int64 `protobuf:"varint,3,opt,name=DestinationAccountId,proto3" json:"DestinationAccountId,omitempty"`
	SourceBeatsId        int64 `protobuf:"varint,4,opt,name=SourceBeatsId,proto3" json:"SourceBeatsId,omitempty"`
	DestinationBeatsId   int64 `protobuf:"varint,5,opt,name=DestinationBeatsId,proto3" json:"DestinationBeatsId,omitempty"`
	ExternalId           int64 `protobuf:"varint,6,opt,name=ExternalId,proto3" json:"ExternalId,omitempty"`
	SentTime             int64 `protobuf:"varint,7,opt,name=SentTime,proto3" json:"SentTime,omitempty"`
	ConfirmationTime     int64 `protobuf:"varint,8,opt,name=ConfirmationTime,proto3" json:"ConfirmationTime,omitempty"`
	Created              int64 `protobuf:"varint,9,opt,name=Created,proto3" json:"Created,omitempty"`
	Updated              int64 `protobuf:"varint,10,opt,name=Updated,proto3" json:"Updated,omitempty"`
}

func (x *SignalResult) Reset() {
	*x = SignalResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignalResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignalResult) ProtoMessage() {}

func (x *SignalResult) ProtoReflect() protoreflect.Message {
	mi := &file_signal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignalResult.ProtoReflect.Descriptor instead.
func (*SignalResult) Descriptor() ([]byte, []int) {
	return file_signal_proto_rawDescGZIP(), []int{1}
}

func (x *SignalResult) GetSignalResultId() int64 {
	if x != nil {
		return x.SignalResultId
	}
	return 0
}

func (x *SignalResult) GetSourceAccountId() int64 {
	if x != nil {
		return x.SourceAccountId
	}
	return 0
}

func (x *SignalResult) GetDestinationAccountId() int64 {
	if x != nil {
		return x.DestinationAccountId
	}
	return 0
}

func (x *SignalResult) GetSourceBeatsId() int64 {
	if x != nil {
		return x.SourceBeatsId
	}
	return 0
}

func (x *SignalResult) GetDestinationBeatsId() int64 {
	if x != nil {
		return x.DestinationBeatsId
	}
	return 0
}

func (x *SignalResult) GetExternalId() int64 {
	if x != nil {
		return x.ExternalId
	}
	return 0
}

func (x *SignalResult) GetSentTime() int64 {
	if x != nil {
		return x.SentTime
	}
	return 0
}

func (x *SignalResult) GetConfirmationTime() int64 {
	if x != nil {
		return x.ConfirmationTime
	}
	return 0
}

func (x *SignalResult) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *SignalResult) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

var File_signal_proto protoreflect.FileDescriptor

var file_signal_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x22, 0xaa,
	0x02, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x32, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4d,
	0x61, 0x78, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x4d, 0x61, 0x78, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x53,
	0x74, 0x6f, 0x70, 0x49, 0x66, 0x4c, 0x65, 0x73, 0x73, 0x54, 0x68, 0x61, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x53, 0x74, 0x6f, 0x70, 0x49, 0x66, 0x4c, 0x65, 0x73, 0x73, 0x54,
	0x68, 0x61, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x44, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x54, 0x6f, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x4d, 0x69, 0x6e, 0x75,
	0x74, 0x65, 0x73, 0x54, 0x6f, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x22, 0x86, 0x03, 0x0a, 0x0c,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x26, 0x0a, 0x0e,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x32,
	0x0a, 0x14, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x65, 0x61, 0x74,
	0x73, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x42, 0x65, 0x61, 0x74, 0x73, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x65, 0x61, 0x74, 0x73, 0x49, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x65, 0x61, 0x74, 0x73, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x45, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65, 0x6e, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x2a, 0xac, 0x01, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x49, 0x47, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x16, 0x0a,
	0x12, 0x53, 0x49, 0x47, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x52,
	0x44, 0x45, 0x52, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x49, 0x47, 0x49, 0x4e, 0x41, 0x4c,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02,
	0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x49, 0x47, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x54, 0x52, 0x41, 0x44, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x03,
	0x12, 0x15, 0x0a, 0x11, 0x53, 0x49, 0x47, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x44, 0x45, 0x41, 0x4c, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x49, 0x47, 0x49, 0x4e,
	0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x54, 0x41, 0x4b,
	0x45, 0x10, 0x05, 0x42, 0x11, 0x5a, 0x0f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x2d, 0x62, 0x65,
	0x61, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_signal_proto_rawDescOnce sync.Once
	file_signal_proto_rawDescData = file_signal_proto_rawDesc
)

func file_signal_proto_rawDescGZIP() []byte {
	file_signal_proto_rawDescOnce.Do(func() {
		file_signal_proto_rawDescData = protoimpl.X.CompressGZIP(file_signal_proto_rawDescData)
	})
	return file_signal_proto_rawDescData
}

var file_signal_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_signal_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_signal_proto_goTypes = []interface{}{
	(SignalType)(0),      // 0: tickerbeats.v1.SignalType
	(*Signal)(nil),       // 1: tickerbeats.v1.Signal
	(*SignalResult)(nil), // 2: tickerbeats.v1.SignalResult
}
var file_signal_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_signal_proto_init() }
func file_signal_proto_init() {
	if File_signal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_signal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signal); i {
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
		file_signal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignalResult); i {
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
			RawDescriptor: file_signal_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_signal_proto_goTypes,
		DependencyIndexes: file_signal_proto_depIdxs,
		EnumInfos:         file_signal_proto_enumTypes,
		MessageInfos:      file_signal_proto_msgTypes,
	}.Build()
	File_signal_proto = out.File
	file_signal_proto_rawDesc = nil
	file_signal_proto_goTypes = nil
	file_signal_proto_depIdxs = nil
}
