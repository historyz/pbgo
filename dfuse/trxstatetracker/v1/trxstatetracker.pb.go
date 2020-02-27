// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dfuse/trxstatetracker/v1/trxstatetracker.proto

package v1

import (
	context "context"
	fmt "fmt"
	deth "github.com/eoscanada/pbgo/dfuse/codecs/deth"
	proto "github.com/golang/protobuf/proto"
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

type PooledRequest_FilterField int32

const (
	PooledRequest_FROM_OR_TO PooledRequest_FilterField = 0
	PooledRequest_FROM       PooledRequest_FilterField = 1
	PooledRequest_TO         PooledRequest_FilterField = 2
)

var PooledRequest_FilterField_name = map[int32]string{
	0: "FROM_OR_TO",
	1: "FROM",
	2: "TO",
}

var PooledRequest_FilterField_value = map[string]int32{
	"FROM_OR_TO": 0,
	"FROM":       1,
	"TO":         2,
}

func (x PooledRequest_FilterField) String() string {
	return proto.EnumName(PooledRequest_FilterField_name, int32(x))
}

func (PooledRequest_FilterField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8b35a6aac9a49909, []int{0, 0}
}

type PooledRequest struct {
	FilterAddresses      [][]byte                  `protobuf:"bytes,1,rep,name=filter_addresses,json=filterAddresses,proto3" json:"filter_addresses,omitempty"`
	FilterField          PooledRequest_FilterField `protobuf:"varint,2,opt,name=filter_field,json=filterField,proto3,enum=dfuse.trxstatetracker.v1.PooledRequest_FilterField" json:"filter_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *PooledRequest) Reset()         { *m = PooledRequest{} }
func (m *PooledRequest) String() string { return proto.CompactTextString(m) }
func (*PooledRequest) ProtoMessage()    {}
func (*PooledRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b35a6aac9a49909, []int{0}
}

func (m *PooledRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PooledRequest.Unmarshal(m, b)
}
func (m *PooledRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PooledRequest.Marshal(b, m, deterministic)
}
func (m *PooledRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PooledRequest.Merge(m, src)
}
func (m *PooledRequest) XXX_Size() int {
	return xxx_messageInfo_PooledRequest.Size(m)
}
func (m *PooledRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PooledRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PooledRequest proto.InternalMessageInfo

func (m *PooledRequest) GetFilterAddresses() [][]byte {
	if m != nil {
		return m.FilterAddresses
	}
	return nil
}

func (m *PooledRequest) GetFilterField() PooledRequest_FilterField {
	if m != nil {
		return m.FilterField
	}
	return PooledRequest_FROM_OR_TO
}

type Request struct {
	TransactionId        string   `protobuf:"bytes,2,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b35a6aac9a49909, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func init() {
	proto.RegisterEnum("dfuse.trxstatetracker.v1.PooledRequest_FilterField", PooledRequest_FilterField_name, PooledRequest_FilterField_value)
	proto.RegisterType((*PooledRequest)(nil), "dfuse.trxstatetracker.v1.PooledRequest")
	proto.RegisterType((*Request)(nil), "dfuse.trxstatetracker.v1.Request")
}

func init() {
	proto.RegisterFile("dfuse/trxstatetracker/v1/trxstatetracker.proto", fileDescriptor_8b35a6aac9a49909)
}

var fileDescriptor_8b35a6aac9a49909 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x4f, 0x2a, 0x31,
	0x14, 0x7d, 0xe5, 0xbd, 0x87, 0x7a, 0xf9, 0x9a, 0x74, 0x45, 0x88, 0x31, 0x88, 0x26, 0xe2, 0xa6,
	0xe5, 0xc3, 0x3f, 0xa0, 0x0b, 0x12, 0x17, 0x66, 0x74, 0x98, 0xb8, 0x70, 0x43, 0xca, 0xf4, 0x0e,
	0x4c, 0x44, 0x8a, 0x6d, 0x21, 0xfc, 0x35, 0xd7, 0xfe, 0x31, 0x43, 0x0b, 0x41, 0x48, 0x30, 0x6e,
	0x26, 0xb9, 0xe7, 0xde, 0x7b, 0xce, 0xb9, 0x67, 0x0a, 0x4c, 0xa6, 0x73, 0x83, 0xdc, 0xea, 0xa5,
	0xb1, 0xc2, 0xa2, 0xd5, 0x22, 0x79, 0x45, 0xcd, 0x17, 0xed, 0x7d, 0x88, 0xcd, 0xb4, 0xb2, 0x8a,
	0x56, 0xdd, 0x3c, 0xdb, 0x6f, 0x2e, 0xda, 0xb5, 0x53, 0xcf, 0x94, 0x28, 0x89, 0x89, 0xe1, 0x12,
	0xed, 0xd8, 0x7d, 0xfc, 0x5e, 0xe3, 0x93, 0x40, 0xe9, 0x51, 0xa9, 0x09, 0xca, 0x08, 0xdf, 0xe7,
	0x68, 0x2c, 0xbd, 0x86, 0x20, 0xcd, 0x26, 0x16, 0xf5, 0x40, 0x48, 0xa9, 0xd1, 0x18, 0x34, 0x55,
	0x52, 0xff, 0xdb, 0x2c, 0x46, 0x15, 0x8f, 0xdf, 0x6e, 0x60, 0xfa, 0x0c, 0xc5, 0xf5, 0x68, 0x9a,
	0xe1, 0x44, 0x56, 0x73, 0x75, 0xd2, 0x2c, 0x77, 0xba, 0xec, 0x90, 0x17, 0xb6, 0xa3, 0xc4, 0x7a,
	0x6e, 0xb7, 0xb7, 0x5a, 0x8d, 0x0a, 0xe9, 0xb6, 0x68, 0x70, 0x28, 0x7c, 0xeb, 0xd1, 0x32, 0x40,
	0x2f, 0x0a, 0x1f, 0x06, 0x61, 0x34, 0x88, 0xc3, 0xe0, 0x0f, 0x3d, 0x86, 0x7f, 0xab, 0x3a, 0x20,
	0x34, 0x0f, 0xb9, 0x38, 0x0c, 0x72, 0x0d, 0x0e, 0x47, 0x1b, 0xfb, 0x97, 0x50, 0xb2, 0x5a, 0x4c,
	0x8d, 0x48, 0x6c, 0xa6, 0xa6, 0xf7, 0xde, 0xd4, 0x49, 0xb4, 0x0b, 0x76, 0x3e, 0x08, 0x54, 0x62,
	0xbd, 0xec, 0xaf, 0xfc, 0xc5, 0xde, 0x1f, 0x7d, 0x82, 0xff, 0xae, 0xa6, 0xe7, 0x87, 0x0f, 0x58,
	0xab, 0xd4, 0x2e, 0xd6, 0x23, 0x3e, 0x55, 0xe6, 0x02, 0x8d, 0xb7, 0x0a, 0x8e, 0xa7, 0x45, 0x68,
	0x1f, 0xf2, 0xfe, 0x64, 0x7a, 0xf5, 0xcb, 0x50, 0x6a, 0x67, 0x3f, 0x33, 0xb7, 0xc8, 0xdd, 0xcd,
	0x4b, 0x67, 0x94, 0xd9, 0xf1, 0x7c, 0xc8, 0x12, 0xf5, 0xc6, 0x51, 0x99, 0x44, 0x4c, 0x85, 0x14,
	0x7c, 0x36, 0x1c, 0x29, 0x7e, 0xe8, 0xd9, 0x0c, 0xf3, 0xee, 0x7f, 0x77, 0xbf, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xe8, 0x4b, 0xbb, 0x24, 0x59, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TrxStateTrackerClient is the client API for TrxStateTracker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TrxStateTrackerClient interface {
	State(ctx context.Context, in *Request, opts ...grpc.CallOption) (TrxStateTracker_StateClient, error)
	Pooled(ctx context.Context, in *PooledRequest, opts ...grpc.CallOption) (TrxStateTracker_PooledClient, error)
}

type trxStateTrackerClient struct {
	cc *grpc.ClientConn
}

func NewTrxStateTrackerClient(cc *grpc.ClientConn) TrxStateTrackerClient {
	return &trxStateTrackerClient{cc}
}

func (c *trxStateTrackerClient) State(ctx context.Context, in *Request, opts ...grpc.CallOption) (TrxStateTracker_StateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TrxStateTracker_serviceDesc.Streams[0], "/dfuse.trxstatetracker.v1.TrxStateTracker/State", opts...)
	if err != nil {
		return nil, err
	}
	x := &trxStateTrackerStateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TrxStateTracker_StateClient interface {
	Recv() (*deth.TransactionState, error)
	grpc.ClientStream
}

type trxStateTrackerStateClient struct {
	grpc.ClientStream
}

func (x *trxStateTrackerStateClient) Recv() (*deth.TransactionState, error) {
	m := new(deth.TransactionState)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *trxStateTrackerClient) Pooled(ctx context.Context, in *PooledRequest, opts ...grpc.CallOption) (TrxStateTracker_PooledClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TrxStateTracker_serviceDesc.Streams[1], "/dfuse.trxstatetracker.v1.TrxStateTracker/Pooled", opts...)
	if err != nil {
		return nil, err
	}
	x := &trxStateTrackerPooledClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TrxStateTracker_PooledClient interface {
	Recv() (*deth.Transaction, error)
	grpc.ClientStream
}

type trxStateTrackerPooledClient struct {
	grpc.ClientStream
}

func (x *trxStateTrackerPooledClient) Recv() (*deth.Transaction, error) {
	m := new(deth.Transaction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TrxStateTrackerServer is the server API for TrxStateTracker service.
type TrxStateTrackerServer interface {
	State(*Request, TrxStateTracker_StateServer) error
	Pooled(*PooledRequest, TrxStateTracker_PooledServer) error
}

// UnimplementedTrxStateTrackerServer can be embedded to have forward compatible implementations.
type UnimplementedTrxStateTrackerServer struct {
}

func (*UnimplementedTrxStateTrackerServer) State(req *Request, srv TrxStateTracker_StateServer) error {
	return status.Errorf(codes.Unimplemented, "method State not implemented")
}
func (*UnimplementedTrxStateTrackerServer) Pooled(req *PooledRequest, srv TrxStateTracker_PooledServer) error {
	return status.Errorf(codes.Unimplemented, "method Pooled not implemented")
}

func RegisterTrxStateTrackerServer(s *grpc.Server, srv TrxStateTrackerServer) {
	s.RegisterService(&_TrxStateTracker_serviceDesc, srv)
}

func _TrxStateTracker_State_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TrxStateTrackerServer).State(m, &trxStateTrackerStateServer{stream})
}

type TrxStateTracker_StateServer interface {
	Send(*deth.TransactionState) error
	grpc.ServerStream
}

type trxStateTrackerStateServer struct {
	grpc.ServerStream
}

func (x *trxStateTrackerStateServer) Send(m *deth.TransactionState) error {
	return x.ServerStream.SendMsg(m)
}

func _TrxStateTracker_Pooled_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PooledRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TrxStateTrackerServer).Pooled(m, &trxStateTrackerPooledServer{stream})
}

type TrxStateTracker_PooledServer interface {
	Send(*deth.Transaction) error
	grpc.ServerStream
}

type trxStateTrackerPooledServer struct {
	grpc.ServerStream
}

func (x *trxStateTrackerPooledServer) Send(m *deth.Transaction) error {
	return x.ServerStream.SendMsg(m)
}

var _TrxStateTracker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dfuse.trxstatetracker.v1.TrxStateTracker",
	HandlerType: (*TrxStateTrackerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "State",
			Handler:       _TrxStateTracker_State_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Pooled",
			Handler:       _TrxStateTracker_Pooled_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dfuse/trxstatetracker/v1/trxstatetracker.proto",
}
