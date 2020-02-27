// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dfuse/headinfo/v1/headinfo.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HeadInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeadInfoRequest) Reset()         { *m = HeadInfoRequest{} }
func (m *HeadInfoRequest) String() string { return proto.CompactTextString(m) }
func (*HeadInfoRequest) ProtoMessage()    {}
func (*HeadInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62a316d938b72ee6, []int{0}
}

func (m *HeadInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeadInfoRequest.Unmarshal(m, b)
}
func (m *HeadInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeadInfoRequest.Marshal(b, m, deterministic)
}
func (m *HeadInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeadInfoRequest.Merge(m, src)
}
func (m *HeadInfoRequest) XXX_Size() int {
	return xxx_messageInfo_HeadInfoRequest.Size(m)
}
func (m *HeadInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HeadInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HeadInfoRequest proto.InternalMessageInfo

type HeadInfoResponse struct {
	LibNum               uint64               `protobuf:"varint,1,opt,name=libNum,proto3" json:"libNum,omitempty"`
	LibID                string               `protobuf:"bytes,2,opt,name=libID,proto3" json:"libID,omitempty"`
	HeadNum              uint64               `protobuf:"varint,10,opt,name=headNum,proto3" json:"headNum,omitempty"`
	HeadID               string               `protobuf:"bytes,11,opt,name=headID,proto3" json:"headID,omitempty"`
	HeadTime             *timestamp.Timestamp `protobuf:"bytes,12,opt,name=headTime,proto3" json:"headTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *HeadInfoResponse) Reset()         { *m = HeadInfoResponse{} }
func (m *HeadInfoResponse) String() string { return proto.CompactTextString(m) }
func (*HeadInfoResponse) ProtoMessage()    {}
func (*HeadInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62a316d938b72ee6, []int{1}
}

func (m *HeadInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeadInfoResponse.Unmarshal(m, b)
}
func (m *HeadInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeadInfoResponse.Marshal(b, m, deterministic)
}
func (m *HeadInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeadInfoResponse.Merge(m, src)
}
func (m *HeadInfoResponse) XXX_Size() int {
	return xxx_messageInfo_HeadInfoResponse.Size(m)
}
func (m *HeadInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HeadInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HeadInfoResponse proto.InternalMessageInfo

func (m *HeadInfoResponse) GetLibNum() uint64 {
	if m != nil {
		return m.LibNum
	}
	return 0
}

func (m *HeadInfoResponse) GetLibID() string {
	if m != nil {
		return m.LibID
	}
	return ""
}

func (m *HeadInfoResponse) GetHeadNum() uint64 {
	if m != nil {
		return m.HeadNum
	}
	return 0
}

func (m *HeadInfoResponse) GetHeadID() string {
	if m != nil {
		return m.HeadID
	}
	return ""
}

func (m *HeadInfoResponse) GetHeadTime() *timestamp.Timestamp {
	if m != nil {
		return m.HeadTime
	}
	return nil
}

func init() {
	proto.RegisterType((*HeadInfoRequest)(nil), "dfuse.headinfo.v1.HeadInfoRequest")
	proto.RegisterType((*HeadInfoResponse)(nil), "dfuse.headinfo.v1.HeadInfoResponse")
}

func init() { proto.RegisterFile("dfuse/headinfo/v1/headinfo.proto", fileDescriptor_62a316d938b72ee6) }

var fileDescriptor_62a316d938b72ee6 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x59, 0xad, 0x35, 0x4e, 0xfd, 0xd3, 0x2e, 0x45, 0x42, 0x2f, 0x86, 0x7a, 0x29, 0x88,
	0xbb, 0xb6, 0x82, 0x1f, 0x40, 0x0a, 0xda, 0x1e, 0x3c, 0xc4, 0xe2, 0x41, 0x4f, 0xbb, 0xcd, 0x24,
	0x5d, 0xe8, 0x66, 0x63, 0x77, 0xd3, 0x8f, 0xe6, 0xe7, 0x93, 0x24, 0x4d, 0x04, 0x0b, 0x9e, 0xbc,
	0xcd, 0x7b, 0xfc, 0x66, 0xf2, 0xf2, 0x58, 0x08, 0xa2, 0x38, 0xb7, 0xc8, 0x57, 0x28, 0x22, 0x95,
	0xc6, 0x86, 0x6f, 0xc7, 0xcd, 0xcc, 0xb2, 0x8d, 0x71, 0x86, 0xf6, 0x4a, 0x82, 0x35, 0xee, 0x76,
	0x3c, 0xb8, 0x4a, 0x8c, 0x49, 0xd6, 0xc8, 0x4b, 0x40, 0xe6, 0x31, 0x77, 0x4a, 0xa3, 0x75, 0x42,
	0x67, 0xd5, 0xce, 0xb0, 0x07, 0x17, 0xcf, 0x28, 0xa2, 0x59, 0x1a, 0x9b, 0x10, 0x3f, 0x73, 0xb4,
	0x6e, 0xf8, 0x45, 0xa0, 0xfb, 0xe3, 0xd9, 0xcc, 0xa4, 0x16, 0xe9, 0x25, 0xb4, 0xd7, 0x4a, 0xbe,
	0xe4, 0xda, 0x27, 0x01, 0x19, 0xb5, 0xc2, 0x9d, 0xa2, 0x7d, 0x38, 0x5a, 0x2b, 0x39, 0x9b, 0xfa,
	0x07, 0x01, 0x19, 0x9d, 0x84, 0x95, 0xa0, 0x3e, 0x1c, 0x17, 0x29, 0x0a, 0x1c, 0x4a, 0xbc, 0x96,
	0xc5, 0x9d, 0x62, 0x9c, 0x4d, 0xfd, 0x4e, 0xb9, 0xb0, 0x53, 0xf4, 0x01, 0xbc, 0x62, 0x5a, 0x28,
	0x8d, 0xfe, 0x69, 0x40, 0x46, 0x9d, 0xc9, 0x80, 0x55, 0xd9, 0x59, 0x9d, 0x9d, 0x2d, 0xea, 0xec,
	0x61, 0xc3, 0xce, 0x5b, 0xde, 0x61, 0x17, 0xe6, 0x2d, 0xef, 0xac, 0xdb, 0x9f, 0x48, 0xf0, 0xea,
	0xdc, 0xf4, 0x0d, 0x3a, 0x4f, 0xe8, 0x1a, 0x39, 0x64, 0x7b, 0xdd, 0xb0, 0x5f, 0xff, 0x3d, 0xb8,
	0xfe, 0x93, 0xa9, 0x7a, 0x98, 0x64, 0xd0, 0x7b, 0x75, 0x1b, 0x14, 0x5a, 0xa5, 0x49, 0x73, 0xfd,
	0x03, 0xce, 0x2b, 0xf3, 0xdf, 0xbf, 0x77, 0x47, 0x1e, 0x6f, 0xdf, 0x6f, 0x12, 0xe5, 0x56, 0xb9,
	0x64, 0x4b, 0xa3, 0x39, 0x1a, 0xbb, 0x14, 0xa9, 0x88, 0x04, 0xcf, 0x64, 0x62, 0xf8, 0xde, 0x9b,
	0x90, 0xed, 0xb2, 0xae, 0xfb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x61, 0xc5, 0xf4, 0x57, 0x2f,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HeadInfoClient is the client API for HeadInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HeadInfoClient interface {
	GetHeadInfo(ctx context.Context, in *HeadInfoRequest, opts ...grpc.CallOption) (*HeadInfoResponse, error)
}

type headInfoClient struct {
	cc *grpc.ClientConn
}

func NewHeadInfoClient(cc *grpc.ClientConn) HeadInfoClient {
	return &headInfoClient{cc}
}

func (c *headInfoClient) GetHeadInfo(ctx context.Context, in *HeadInfoRequest, opts ...grpc.CallOption) (*HeadInfoResponse, error) {
	out := new(HeadInfoResponse)
	err := c.cc.Invoke(ctx, "/dfuse.headinfo.v1.HeadInfo/GetHeadInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeadInfoServer is the server API for HeadInfo service.
type HeadInfoServer interface {
	GetHeadInfo(context.Context, *HeadInfoRequest) (*HeadInfoResponse, error)
}

// UnimplementedHeadInfoServer can be embedded to have forward compatible implementations.
type UnimplementedHeadInfoServer struct {
}

func (*UnimplementedHeadInfoServer) GetHeadInfo(ctx context.Context, req *HeadInfoRequest) (*HeadInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHeadInfo not implemented")
}

func RegisterHeadInfoServer(s *grpc.Server, srv HeadInfoServer) {
	s.RegisterService(&_HeadInfo_serviceDesc, srv)
}

func _HeadInfo_GetHeadInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeadInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadInfoServer).GetHeadInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfuse.headinfo.v1.HeadInfo/GetHeadInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadInfoServer).GetHeadInfo(ctx, req.(*HeadInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HeadInfo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dfuse.headinfo.v1.HeadInfo",
	HandlerType: (*HeadInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHeadInfo",
			Handler:    _HeadInfo_GetHeadInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dfuse/headinfo/v1/headinfo.proto",
}

// StreamingHeadInfoClient is the client API for StreamingHeadInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamingHeadInfoClient interface {
	StreamHeadInfo(ctx context.Context, in *HeadInfoRequest, opts ...grpc.CallOption) (StreamingHeadInfo_StreamHeadInfoClient, error)
}

type streamingHeadInfoClient struct {
	cc *grpc.ClientConn
}

func NewStreamingHeadInfoClient(cc *grpc.ClientConn) StreamingHeadInfoClient {
	return &streamingHeadInfoClient{cc}
}

func (c *streamingHeadInfoClient) StreamHeadInfo(ctx context.Context, in *HeadInfoRequest, opts ...grpc.CallOption) (StreamingHeadInfo_StreamHeadInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamingHeadInfo_serviceDesc.Streams[0], "/dfuse.headinfo.v1.StreamingHeadInfo/StreamHeadInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingHeadInfoStreamHeadInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamingHeadInfo_StreamHeadInfoClient interface {
	Recv() (*HeadInfoResponse, error)
	grpc.ClientStream
}

type streamingHeadInfoStreamHeadInfoClient struct {
	grpc.ClientStream
}

func (x *streamingHeadInfoStreamHeadInfoClient) Recv() (*HeadInfoResponse, error) {
	m := new(HeadInfoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamingHeadInfoServer is the server API for StreamingHeadInfo service.
type StreamingHeadInfoServer interface {
	StreamHeadInfo(*HeadInfoRequest, StreamingHeadInfo_StreamHeadInfoServer) error
}

// UnimplementedStreamingHeadInfoServer can be embedded to have forward compatible implementations.
type UnimplementedStreamingHeadInfoServer struct {
}

func (*UnimplementedStreamingHeadInfoServer) StreamHeadInfo(req *HeadInfoRequest, srv StreamingHeadInfo_StreamHeadInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamHeadInfo not implemented")
}

func RegisterStreamingHeadInfoServer(s *grpc.Server, srv StreamingHeadInfoServer) {
	s.RegisterService(&_StreamingHeadInfo_serviceDesc, srv)
}

func _StreamingHeadInfo_StreamHeadInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HeadInfoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamingHeadInfoServer).StreamHeadInfo(m, &streamingHeadInfoStreamHeadInfoServer{stream})
}

type StreamingHeadInfo_StreamHeadInfoServer interface {
	Send(*HeadInfoResponse) error
	grpc.ServerStream
}

type streamingHeadInfoStreamHeadInfoServer struct {
	grpc.ServerStream
}

func (x *streamingHeadInfoStreamHeadInfoServer) Send(m *HeadInfoResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _StreamingHeadInfo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dfuse.headinfo.v1.StreamingHeadInfo",
	HandlerType: (*StreamingHeadInfoServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamHeadInfo",
			Handler:       _StreamingHeadInfo_StreamHeadInfo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dfuse/headinfo/v1/headinfo.proto",
}
