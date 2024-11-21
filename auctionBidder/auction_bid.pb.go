// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: auctionBidder/auction_bid.proto

package auctionBidder

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_auctionBidder_auction_bid_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_auctionBidder_auction_bid_proto_msgTypes[0]
	if x != nil {
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
	return file_auctionBidder_auction_bid_proto_rawDescGZIP(), []int{0}
}

// Bidder information for the Auction
type FromBidder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount    int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	ID        int64 `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *FromBidder) Reset() {
	*x = FromBidder{}
	mi := &file_auctionBidder_auction_bid_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FromBidder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FromBidder) ProtoMessage() {}

func (x *FromBidder) ProtoReflect() protoreflect.Message {
	mi := &file_auctionBidder_auction_bid_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FromBidder.ProtoReflect.Descriptor instead.
func (*FromBidder) Descriptor() ([]byte, []int) {
	return file_auctionBidder_auction_bid_proto_rawDescGZIP(), []int{1}
}

func (x *FromBidder) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *FromBidder) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *FromBidder) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// Auction responds
type FromAuction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Acknowledgment bool   `protobuf:"varint,1,opt,name=acknowledgment,proto3" json:"acknowledgment,omitempty"`
	Comment        string `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *FromAuction) Reset() {
	*x = FromAuction{}
	mi := &file_auctionBidder_auction_bid_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FromAuction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FromAuction) ProtoMessage() {}

func (x *FromAuction) ProtoReflect() protoreflect.Message {
	mi := &file_auctionBidder_auction_bid_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FromAuction.ProtoReflect.Descriptor instead.
func (*FromAuction) Descriptor() ([]byte, []int) {
	return file_auctionBidder_auction_bid_proto_rawDescGZIP(), []int{2}
}

func (x *FromAuction) GetAcknowledgment() bool {
	if x != nil {
		return x.Acknowledgment
	}
	return false
}

func (x *FromAuction) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuctionActive bool   `protobuf:"varint,1,opt,name=AuctionActive,proto3" json:"AuctionActive,omitempty"`
	Comment       string `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	ID            int64  `protobuf:"varint,3,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	mi := &file_auctionBidder_auction_bid_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_auctionBidder_auction_bid_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_auctionBidder_auction_bid_proto_rawDescGZIP(), []int{3}
}

func (x *Result) GetAuctionActive() bool {
	if x != nil {
		return x.AuctionActive
	}
	return false
}

func (x *Result) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Result) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

var File_auctionBidder_auction_bid_proto protoreflect.FileDescriptor

var file_auctionBidder_auction_bid_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x2f,
	0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x0a, 0x0a, 0x46, 0x72, 0x6f,
	0x6d, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x4f, 0x0a,
	0x0b, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e,
	0x61, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x61, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x58,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x41, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x32, 0x88, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x03, 0x62, 0x69,
	0x64, 0x12, 0x19, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x64, 0x64, 0x65,
	0x72, 0x2e, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x1a, 0x1a, 0x2e, 0x61,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x2e, 0x46, 0x72, 0x6f,
	0x6d, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69,
	0x64, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x56, 0x69, 0x6b, 0x74, 0x6f, 0x72, 0x45, 0x6d, 0x69, 0x6c, 0x32, 0x30, 0x30, 0x30,
	0x2f, 0x44, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2d, 0x43, 0x61, 0x74, 0x61, 0x6e, 0x45, 0x6e,
	0x6a, 0x6f, 0x79, 0x65, 0x72, 0x73, 0x2d, 0x48, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x35, 0x2f, 0x61,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auctionBidder_auction_bid_proto_rawDescOnce sync.Once
	file_auctionBidder_auction_bid_proto_rawDescData = file_auctionBidder_auction_bid_proto_rawDesc
)

func file_auctionBidder_auction_bid_proto_rawDescGZIP() []byte {
	file_auctionBidder_auction_bid_proto_rawDescOnce.Do(func() {
		file_auctionBidder_auction_bid_proto_rawDescData = protoimpl.X.CompressGZIP(file_auctionBidder_auction_bid_proto_rawDescData)
	})
	return file_auctionBidder_auction_bid_proto_rawDescData
}

var file_auctionBidder_auction_bid_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_auctionBidder_auction_bid_proto_goTypes = []any{
	(*Empty)(nil),       // 0: auctionBidder.Empty
	(*FromBidder)(nil),  // 1: auctionBidder.FromBidder
	(*FromAuction)(nil), // 2: auctionBidder.FromAuction
	(*Result)(nil),      // 3: auctionBidder.Result
}
var file_auctionBidder_auction_bid_proto_depIdxs = []int32{
	1, // 0: auctionBidder.Communication.bid:input_type -> auctionBidder.FromBidder
	0, // 1: auctionBidder.Communication.result:input_type -> auctionBidder.Empty
	2, // 2: auctionBidder.Communication.bid:output_type -> auctionBidder.FromAuction
	3, // 3: auctionBidder.Communication.result:output_type -> auctionBidder.Result
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auctionBidder_auction_bid_proto_init() }
func file_auctionBidder_auction_bid_proto_init() {
	if File_auctionBidder_auction_bid_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auctionBidder_auction_bid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auctionBidder_auction_bid_proto_goTypes,
		DependencyIndexes: file_auctionBidder_auction_bid_proto_depIdxs,
		MessageInfos:      file_auctionBidder_auction_bid_proto_msgTypes,
	}.Build()
	File_auctionBidder_auction_bid_proto = out.File
	file_auctionBidder_auction_bid_proto_rawDesc = nil
	file_auctionBidder_auction_bid_proto_goTypes = nil
	file_auctionBidder_auction_bid_proto_depIdxs = nil
}
