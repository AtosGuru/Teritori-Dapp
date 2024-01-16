// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: prices/v1/prices.proto

package pricespb

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

type PricesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Time  string   `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	VsIds []string `protobuf:"bytes,3,rep,name=vs_ids,json=vsIds,proto3" json:"vs_ids,omitempty"`
}

func (x *PricesRequest) Reset() {
	*x = PricesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prices_v1_prices_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PricesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PricesRequest) ProtoMessage() {}

func (x *PricesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prices_v1_prices_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PricesRequest.ProtoReflect.Descriptor instead.
func (*PricesRequest) Descriptor() ([]byte, []int) {
	return file_prices_v1_prices_proto_rawDescGZIP(), []int{0}
}

func (x *PricesRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PricesRequest) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *PricesRequest) GetVsIds() []string {
	if x != nil {
		return x.VsIds
	}
	return nil
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prices_v1_prices_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_prices_v1_prices_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_prices_v1_prices_proto_rawDescGZIP(), []int{1}
}

func (x *Price) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Price) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type PricesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prices []*Price `protobuf:"bytes,1,rep,name=prices,proto3" json:"prices,omitempty"`
}

func (x *PricesResponse) Reset() {
	*x = PricesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prices_v1_prices_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PricesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PricesResponse) ProtoMessage() {}

func (x *PricesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_prices_v1_prices_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PricesResponse.ProtoReflect.Descriptor instead.
func (*PricesResponse) Descriptor() ([]byte, []int) {
	return file_prices_v1_prices_proto_rawDescGZIP(), []int{2}
}

func (x *PricesResponse) GetPrices() []*Price {
	if x != nil {
		return x.Prices
	}
	return nil
}

var File_prices_v1_prices_proto protoreflect.FileDescriptor

var file_prices_v1_prices_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x22, 0x4a, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x76, 0x73, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x76, 0x73, 0x49, 0x64, 0x73, 0x22,
	0x2d, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3a,
	0x0a, 0x0e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x28, 0x0a, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x52, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x32, 0x4e, 0x0a, 0x0d, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prices_v1_prices_proto_rawDescOnce sync.Once
	file_prices_v1_prices_proto_rawDescData = file_prices_v1_prices_proto_rawDesc
)

func file_prices_v1_prices_proto_rawDescGZIP() []byte {
	file_prices_v1_prices_proto_rawDescOnce.Do(func() {
		file_prices_v1_prices_proto_rawDescData = protoimpl.X.CompressGZIP(file_prices_v1_prices_proto_rawDescData)
	})
	return file_prices_v1_prices_proto_rawDescData
}

var file_prices_v1_prices_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_prices_v1_prices_proto_goTypes = []interface{}{
	(*PricesRequest)(nil),  // 0: prices.v1.PricesRequest
	(*Price)(nil),          // 1: prices.v1.Price
	(*PricesResponse)(nil), // 2: prices.v1.PricesResponse
}
var file_prices_v1_prices_proto_depIdxs = []int32{
	1, // 0: prices.v1.PricesResponse.prices:type_name -> prices.v1.Price
	0, // 1: prices.v1.PricesService.Prices:input_type -> prices.v1.PricesRequest
	2, // 2: prices.v1.PricesService.Prices:output_type -> prices.v1.PricesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_prices_v1_prices_proto_init() }
func file_prices_v1_prices_proto_init() {
	if File_prices_v1_prices_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prices_v1_prices_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PricesRequest); i {
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
		file_prices_v1_prices_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
		file_prices_v1_prices_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PricesResponse); i {
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
			RawDescriptor: file_prices_v1_prices_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_prices_v1_prices_proto_goTypes,
		DependencyIndexes: file_prices_v1_prices_proto_depIdxs,
		MessageInfos:      file_prices_v1_prices_proto_msgTypes,
	}.Build()
	File_prices_v1_prices_proto = out.File
	file_prices_v1_prices_proto_rawDesc = nil
	file_prices_v1_prices_proto_goTypes = nil
	file_prices_v1_prices_proto_depIdxs = nil
}