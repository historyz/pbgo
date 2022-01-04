// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: sf/merger/v1/merger.proto

package pbmerger

import (
	context "context"
	v1 "github.com/streamingfast/pbgo/sf/bstream/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LowBlockNum uint64 `protobuf:"varint,1,opt,name=lowBlockNum,proto3" json:"lowBlockNum,omitempty"`
	HighBlockID string `protobuf:"bytes,2,opt,name=highBlockID,proto3" json:"highBlockID,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_merger_v1_merger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_sf_merger_v1_merger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_sf_merger_v1_merger_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetLowBlockNum() uint64 {
	if x != nil {
		return x.LowBlockNum
	}
	return 0
}

func (x *Request) GetHighBlockID() string {
	if x != nil {
		return x.HighBlockID
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Found bool      `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
	Block *v1.Block `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_merger_v1_merger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_sf_merger_v1_merger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_sf_merger_v1_merger_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

func (x *Response) GetBlock() *v1.Block {
	if x != nil {
		return x.Block
	}
	return nil
}

var File_sf_merger_v1_merger_proto protoreflect.FileDescriptor

var file_sf_merger_v1_merger_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x66, 0x2f, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x72, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x66, 0x2e,
	0x6d, 0x65, 0x72, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x73, 0x66, 0x2f, 0x62, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x6f, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6c, 0x6f, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x4e, 0x75, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x68, 0x69, 0x67, 0x68, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x69, 0x67, 0x68, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x49, 0x44, 0x22, 0x4c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x2a, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x66, 0x2e, 0x62, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x32, 0x4c, 0x0a, 0x06, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x72, 0x12, 0x42, 0x0a,
	0x0f, 0x50, 0x72, 0x65, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x12, 0x15, 0x2e, 0x73, 0x66, 0x2e, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x66, 0x2e, 0x6d, 0x65, 0x72,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x70, 0x62,
	0x67, 0x6f, 0x2f, 0x73, 0x66, 0x2f, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b,
	0x70, 0x62, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sf_merger_v1_merger_proto_rawDescOnce sync.Once
	file_sf_merger_v1_merger_proto_rawDescData = file_sf_merger_v1_merger_proto_rawDesc
)

func file_sf_merger_v1_merger_proto_rawDescGZIP() []byte {
	file_sf_merger_v1_merger_proto_rawDescOnce.Do(func() {
		file_sf_merger_v1_merger_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_merger_v1_merger_proto_rawDescData)
	})
	return file_sf_merger_v1_merger_proto_rawDescData
}

var file_sf_merger_v1_merger_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sf_merger_v1_merger_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: sf.merger.v1.Request
	(*Response)(nil), // 1: sf.merger.v1.Response
	(*v1.Block)(nil), // 2: sf.bstream.v1.Block
}
var file_sf_merger_v1_merger_proto_depIdxs = []int32{
	2, // 0: sf.merger.v1.Response.block:type_name -> sf.bstream.v1.Block
	0, // 1: sf.merger.v1.Merger.PreMergedBlocks:input_type -> sf.merger.v1.Request
	1, // 2: sf.merger.v1.Merger.PreMergedBlocks:output_type -> sf.merger.v1.Response
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sf_merger_v1_merger_proto_init() }
func file_sf_merger_v1_merger_proto_init() {
	if File_sf_merger_v1_merger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sf_merger_v1_merger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_sf_merger_v1_merger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_sf_merger_v1_merger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sf_merger_v1_merger_proto_goTypes,
		DependencyIndexes: file_sf_merger_v1_merger_proto_depIdxs,
		MessageInfos:      file_sf_merger_v1_merger_proto_msgTypes,
	}.Build()
	File_sf_merger_v1_merger_proto = out.File
	file_sf_merger_v1_merger_proto_rawDesc = nil
	file_sf_merger_v1_merger_proto_goTypes = nil
	file_sf_merger_v1_merger_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MergerClient is the client API for Merger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MergerClient interface {
	PreMergedBlocks(ctx context.Context, in *Request, opts ...grpc.CallOption) (Merger_PreMergedBlocksClient, error)
}

type mergerClient struct {
	cc grpc.ClientConnInterface
}

func NewMergerClient(cc grpc.ClientConnInterface) MergerClient {
	return &mergerClient{cc}
}

func (c *mergerClient) PreMergedBlocks(ctx context.Context, in *Request, opts ...grpc.CallOption) (Merger_PreMergedBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Merger_serviceDesc.Streams[0], "/sf.merger.v1.Merger/PreMergedBlocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &mergerPreMergedBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Merger_PreMergedBlocksClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type mergerPreMergedBlocksClient struct {
	grpc.ClientStream
}

func (x *mergerPreMergedBlocksClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MergerServer is the server API for Merger service.
type MergerServer interface {
	PreMergedBlocks(*Request, Merger_PreMergedBlocksServer) error
}

// UnimplementedMergerServer can be embedded to have forward compatible implementations.
type UnimplementedMergerServer struct {
}

func (*UnimplementedMergerServer) PreMergedBlocks(*Request, Merger_PreMergedBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method PreMergedBlocks not implemented")
}

func RegisterMergerServer(s *grpc.Server, srv MergerServer) {
	s.RegisterService(&_Merger_serviceDesc, srv)
}

func _Merger_PreMergedBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MergerServer).PreMergedBlocks(m, &mergerPreMergedBlocksServer{stream})
}

type Merger_PreMergedBlocksServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type mergerPreMergedBlocksServer struct {
	grpc.ServerStream
}

func (x *mergerPreMergedBlocksServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

var _Merger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sf.merger.v1.Merger",
	HandlerType: (*MergerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PreMergedBlocks",
			Handler:       _Merger_PreMergedBlocks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sf/merger/v1/merger.proto",
}
