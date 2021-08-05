// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: enqueuer.proto

package v1

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

type Status int32

const (
	Status_STOPPED Status = 0
	Status_STARTED Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "STOPPED",
		1: "STARTED",
	}
	Status_value = map[string]int32{
		"STOPPED": 0,
		"STARTED": 1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_enqueuer_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_enqueuer_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_enqueuer_proto_rawDescGZIP(), []int{0}
}

type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Action    int32  `protobuf:"varint,2,opt,name=action,proto3" json:"action,omitempty"`
	Selector  string `protobuf:"bytes,3,opt,name=selector,proto3" json:"selector,omitempty"`
	Attribute string `protobuf:"bytes,4,opt,name=attribute,proto3" json:"attribute,omitempty"`
}

func (x *Rule) Reset() {
	*x = Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enqueuer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_enqueuer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_enqueuer_proto_rawDescGZIP(), []int{0}
}

func (x *Rule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Rule) GetAction() int32 {
	if x != nil {
		return x.Action
	}
	return 0
}

func (x *Rule) GetSelector() string {
	if x != nil {
		return x.Selector
	}
	return ""
}

func (x *Rule) GetAttribute() string {
	if x != nil {
		return x.Attribute
	}
	return ""
}

type AllowedDomain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain            string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	UrlRegexValidator string `protobuf:"bytes,2,opt,name=url_regex_validator,json=urlRegexValidator,proto3" json:"url_regex_validator,omitempty"`
}

func (x *AllowedDomain) Reset() {
	*x = AllowedDomain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enqueuer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowedDomain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowedDomain) ProtoMessage() {}

func (x *AllowedDomain) ProtoReflect() protoreflect.Message {
	mi := &file_enqueuer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllowedDomain.ProtoReflect.Descriptor instead.
func (*AllowedDomain) Descriptor() ([]byte, []int) {
	return file_enqueuer_proto_rawDescGZIP(), []int{1}
}

func (x *AllowedDomain) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *AllowedDomain) GetUrlRegexValidator() string {
	if x != nil {
		return x.UrlRegexValidator
	}
	return ""
}

// Deal: A deal is the reflection of the fact of a trade operation execution based on an order that contains a trade request
type Enqueuer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceId          int32                  `protobuf:"varint,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	StartUrl          string                 `protobuf:"bytes,2,opt,name=start_url,json=startUrl,proto3" json:"start_url,omitempty"`
	AllowedDomain     []*AllowedDomain       `protobuf:"bytes,3,rep,name=allowed_domain,json=allowedDomain,proto3" json:"allowed_domain,omitempty"`
	DurationInMinutes uint32                 `protobuf:"varint,4,opt,name=duration_in_minutes,json=durationInMinutes,proto3" json:"duration_in_minutes,omitempty"`
	MaxDepth          uint32                 `protobuf:"varint,5,opt,name=max_depth,json=maxDepth,proto3" json:"max_depth,omitempty"`
	Status            Status                 `protobuf:"varint,6,opt,name=status,proto3,enum=tickerscraper.v1.Status" json:"status,omitempty"`
	Rules             []*Rule                `protobuf:"bytes,7,rep,name=rules,proto3" json:"rules,omitempty"`
	Modified          *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=modified,proto3" json:"modified,omitempty"`
}

func (x *Enqueuer) Reset() {
	*x = Enqueuer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enqueuer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Enqueuer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Enqueuer) ProtoMessage() {}

func (x *Enqueuer) ProtoReflect() protoreflect.Message {
	mi := &file_enqueuer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Enqueuer.ProtoReflect.Descriptor instead.
func (*Enqueuer) Descriptor() ([]byte, []int) {
	return file_enqueuer_proto_rawDescGZIP(), []int{2}
}

func (x *Enqueuer) GetSourceId() int32 {
	if x != nil {
		return x.SourceId
	}
	return 0
}

func (x *Enqueuer) GetStartUrl() string {
	if x != nil {
		return x.StartUrl
	}
	return ""
}

func (x *Enqueuer) GetAllowedDomain() []*AllowedDomain {
	if x != nil {
		return x.AllowedDomain
	}
	return nil
}

func (x *Enqueuer) GetDurationInMinutes() uint32 {
	if x != nil {
		return x.DurationInMinutes
	}
	return 0
}

func (x *Enqueuer) GetMaxDepth() uint32 {
	if x != nil {
		return x.MaxDepth
	}
	return 0
}

func (x *Enqueuer) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STOPPED
}

func (x *Enqueuer) GetRules() []*Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *Enqueuer) GetModified() *timestamppb.Timestamp {
	if x != nil {
		return x.Modified
	}
	return nil
}

var File_enqueuer_proto protoreflect.FileDescriptor

var file_enqueuer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x22, 0x57, 0x0a, 0x0d, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x75, 0x72,
	0x6c, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x75, 0x72, 0x6c, 0x52, 0x65, 0x67, 0x65,
	0x78, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x22, 0xf1, 0x02, 0x0a, 0x08, 0x45,
	0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x55, 0x72,
	0x6c, 0x12, 0x46, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x72, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78,
	0x5f, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x61,
	0x78, 0x44, 0x65, 0x70, 0x74, 0x68, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x73,
	0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72,
	0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52,
	0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2a, 0x22,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x4f, 0x50,
	0x50, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44,
	0x10, 0x01, 0x42, 0x13, 0x5a, 0x11, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x2d, 0x73, 0x63, 0x72,
	0x61, 0x70, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enqueuer_proto_rawDescOnce sync.Once
	file_enqueuer_proto_rawDescData = file_enqueuer_proto_rawDesc
)

func file_enqueuer_proto_rawDescGZIP() []byte {
	file_enqueuer_proto_rawDescOnce.Do(func() {
		file_enqueuer_proto_rawDescData = protoimpl.X.CompressGZIP(file_enqueuer_proto_rawDescData)
	})
	return file_enqueuer_proto_rawDescData
}

var file_enqueuer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enqueuer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_enqueuer_proto_goTypes = []interface{}{
	(Status)(0),                   // 0: tickerscraper.v1.Status
	(*Rule)(nil),                  // 1: tickerscraper.v1.Rule
	(*AllowedDomain)(nil),         // 2: tickerscraper.v1.AllowedDomain
	(*Enqueuer)(nil),              // 3: tickerscraper.v1.Enqueuer
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_enqueuer_proto_depIdxs = []int32{
	2, // 0: tickerscraper.v1.Enqueuer.allowed_domain:type_name -> tickerscraper.v1.AllowedDomain
	0, // 1: tickerscraper.v1.Enqueuer.status:type_name -> tickerscraper.v1.Status
	1, // 2: tickerscraper.v1.Enqueuer.rules:type_name -> tickerscraper.v1.Rule
	4, // 3: tickerscraper.v1.Enqueuer.modified:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_enqueuer_proto_init() }
func file_enqueuer_proto_init() {
	if File_enqueuer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enqueuer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rule); i {
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
		file_enqueuer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowedDomain); i {
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
		file_enqueuer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Enqueuer); i {
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
			RawDescriptor: file_enqueuer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enqueuer_proto_goTypes,
		DependencyIndexes: file_enqueuer_proto_depIdxs,
		EnumInfos:         file_enqueuer_proto_enumTypes,
		MessageInfos:      file_enqueuer_proto_msgTypes,
	}.Build()
	File_enqueuer_proto = out.File
	file_enqueuer_proto_rawDesc = nil
	file_enqueuer_proto_goTypes = nil
	file_enqueuer_proto_depIdxs = nil
}
