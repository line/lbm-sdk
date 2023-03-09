// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: messages.proto

package testutil

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/line/lbm-sdk/codec/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgCounter struct {
	Counter       int64 `protobuf:"varint,1,opt,name=counter,proto3" json:"counter,omitempty"`
	FailOnHandler bool  `protobuf:"varint,2,opt,name=fail_on_handler,json=failOnHandler,proto3" json:"fail_on_handler,omitempty"`
}

func (m *MsgCounter) Reset()         { *m = MsgCounter{} }
func (m *MsgCounter) String() string { return proto.CompactTextString(m) }
func (*MsgCounter) ProtoMessage()    {}
func (*MsgCounter) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}
func (m *MsgCounter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCounter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCounter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCounter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCounter.Merge(m, src)
}
func (m *MsgCounter) XXX_Size() int {
	return m.Size()
}
func (m *MsgCounter) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCounter.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCounter proto.InternalMessageInfo

func (m *MsgCounter) GetCounter() int64 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func (m *MsgCounter) GetFailOnHandler() bool {
	if m != nil {
		return m.FailOnHandler
	}
	return false
}

type MsgCounter2 struct {
	Counter       int64 `protobuf:"varint,1,opt,name=counter,proto3" json:"counter,omitempty"`
	FailOnHandler bool  `protobuf:"varint,2,opt,name=fail_on_handler,json=failOnHandler,proto3" json:"fail_on_handler,omitempty"`
}

func (m *MsgCounter2) Reset()         { *m = MsgCounter2{} }
func (m *MsgCounter2) String() string { return proto.CompactTextString(m) }
func (*MsgCounter2) ProtoMessage()    {}
func (*MsgCounter2) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}
func (m *MsgCounter2) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCounter2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCounter2.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCounter2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCounter2.Merge(m, src)
}
func (m *MsgCounter2) XXX_Size() int {
	return m.Size()
}
func (m *MsgCounter2) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCounter2.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCounter2 proto.InternalMessageInfo

func (m *MsgCounter2) GetCounter() int64 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func (m *MsgCounter2) GetFailOnHandler() bool {
	if m != nil {
		return m.FailOnHandler
	}
	return false
}

type MsgCreateCounterResponse struct {
}

func (m *MsgCreateCounterResponse) Reset()         { *m = MsgCreateCounterResponse{} }
func (m *MsgCreateCounterResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateCounterResponse) ProtoMessage()    {}
func (*MsgCreateCounterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{2}
}
func (m *MsgCreateCounterResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateCounterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateCounterResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateCounterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateCounterResponse.Merge(m, src)
}
func (m *MsgCreateCounterResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateCounterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateCounterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateCounterResponse proto.InternalMessageInfo

type MsgKeyValue struct {
	Key    []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value  []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Signer string `protobuf:"bytes,3,opt,name=signer,proto3" json:"signer,omitempty"`
}

func (m *MsgKeyValue) Reset()         { *m = MsgKeyValue{} }
func (m *MsgKeyValue) String() string { return proto.CompactTextString(m) }
func (*MsgKeyValue) ProtoMessage()    {}
func (*MsgKeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{3}
}
func (m *MsgKeyValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgKeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgKeyValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgKeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgKeyValue.Merge(m, src)
}
func (m *MsgKeyValue) XXX_Size() int {
	return m.Size()
}
func (m *MsgKeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgKeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_MsgKeyValue proto.InternalMessageInfo

func (m *MsgKeyValue) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *MsgKeyValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MsgKeyValue) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

type MsgCreateKeyValueResponse struct {
}

func (m *MsgCreateKeyValueResponse) Reset()         { *m = MsgCreateKeyValueResponse{} }
func (m *MsgCreateKeyValueResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateKeyValueResponse) ProtoMessage()    {}
func (*MsgCreateKeyValueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{4}
}
func (m *MsgCreateKeyValueResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateKeyValueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateKeyValueResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateKeyValueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateKeyValueResponse.Merge(m, src)
}
func (m *MsgCreateKeyValueResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateKeyValueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateKeyValueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateKeyValueResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCounter)(nil), "testdata.MsgCounter")
	proto.RegisterType((*MsgCounter2)(nil), "testdata.MsgCounter2")
	proto.RegisterType((*MsgCreateCounterResponse)(nil), "testdata.MsgCreateCounterResponse")
	proto.RegisterType((*MsgKeyValue)(nil), "testdata.MsgKeyValue")
	proto.RegisterType((*MsgCreateKeyValueResponse)(nil), "testdata.MsgCreateKeyValueResponse")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x8a, 0x9b, 0x50,
	0x14, 0x86, 0x63, 0xa5, 0x89, 0x3d, 0x4d, 0xdb, 0x20, 0x69, 0x31, 0x16, 0x24, 0x58, 0x28, 0xd9,
	0x44, 0xc1, 0x3e, 0x41, 0xdb, 0x45, 0x5b, 0x5a, 0x1b, 0xb0, 0xd0, 0x61, 0x66, 0x13, 0xae, 0xe6,
	0xe4, 0x46, 0xa2, 0xf7, 0x8a, 0xf7, 0x3a, 0x90, 0xb7, 0x98, 0xc7, 0x9a, 0x65, 0x96, 0xb3, 0x1c,
	0x92, 0x17, 0x19, 0xd4, 0x98, 0x30, 0xc1, 0xc5, 0x2c, 0x66, 0xe5, 0x39, 0xff, 0x0f, 0xdf, 0xcf,
	0xf9, 0xbd, 0xf0, 0x36, 0x45, 0x21, 0x08, 0x45, 0xe1, 0x64, 0x39, 0x97, 0x5c, 0xd7, 0x24, 0x0a,
	0xb9, 0x20, 0x92, 0x98, 0x43, 0xca, 0x29, 0xaf, 0x44, 0xb7, 0x9c, 0x6a, 0xdf, 0x1c, 0x51, 0xce,
	0x69, 0x82, 0x6e, 0xb5, 0x85, 0xc5, 0xd2, 0x25, 0x6c, 0x53, 0x5b, 0xf6, 0x5f, 0x00, 0x5f, 0xd0,
	0xef, 0xbc, 0x60, 0x12, 0x73, 0xdd, 0x80, 0x5e, 0x54, 0x8f, 0x86, 0x32, 0x56, 0x26, 0x6a, 0xd0,
	0xac, 0xfa, 0x67, 0x78, 0xb7, 0x24, 0x71, 0x32, 0xe7, 0x6c, 0xbe, 0x22, 0x6c, 0x91, 0x60, 0x6e,
	0xbc, 0x18, 0x2b, 0x13, 0x2d, 0x78, 0x53, 0xca, 0x33, 0xf6, 0xb3, 0x16, 0xed, 0x19, 0xbc, 0x3e,
	0xf1, 0xbc, 0x67, 0x00, 0x9a, 0x60, 0x94, 0xc0, 0x1c, 0x89, 0xc4, 0x03, 0x36, 0x40, 0x91, 0x71,
	0x26, 0xd0, 0xf6, 0xab, 0xb0, 0xdf, 0xb8, 0xf9, 0x4f, 0x92, 0x02, 0xf5, 0x01, 0xa8, 0x6b, 0xdc,
	0x54, 0x41, 0xfd, 0xa0, 0x1c, 0xf5, 0x21, 0xbc, 0xbc, 0x2e, 0xad, 0x0a, 0xdd, 0x0f, 0xea, 0x45,
	0xff, 0x00, 0x5d, 0x11, 0x53, 0x86, 0xb9, 0xa1, 0x8e, 0x95, 0xc9, 0xab, 0xe0, 0xb0, 0xd9, 0x1f,
	0x61, 0x74, 0x8c, 0x6a, 0xa0, 0x4d, 0x96, 0x77, 0x01, 0xbd, 0xa6, 0xa5, 0x3f, 0x30, 0xf8, 0xc5,
	0xa2, 0x1c, 0x53, 0x64, 0xb2, 0xd1, 0x86, 0x4e, 0xf3, 0x0f, 0x9c, 0xd3, 0xfd, 0xa6, 0xfd, 0x58,
	0x6d, 0x3b, 0xc2, 0xbb, 0x04, 0xed, 0x58, 0x97, 0xdf, 0x42, 0x7e, 0xdf, 0x46, 0xf6, 0x9e, 0x84,
	0xf6, 0x41, 0x3b, 0x96, 0xf3, 0x15, 0xd4, 0x7f, 0x28, 0xcf, 0x68, 0x8d, 0x6b, 0x7e, 0x6a, 0xa1,
	0x9d, 0x57, 0xf0, 0xed, 0xc7, 0xed, 0xce, 0x52, 0xb6, 0x3b, 0x4b, 0xb9, 0xdf, 0x59, 0xca, 0xcd,
	0xde, 0xea, 0x6c, 0xf7, 0x56, 0xe7, 0x6e, 0x6f, 0x75, 0xae, 0xa6, 0x34, 0x96, 0xab, 0x22, 0x74,
	0x22, 0x9e, 0xba, 0x11, 0x17, 0x29, 0x17, 0x87, 0xcf, 0x54, 0x2c, 0xd6, 0x6e, 0x48, 0x04, 0x92,
	0x2c, 0x73, 0xcb, 0x88, 0x42, 0xc6, 0x49, 0xd8, 0xad, 0xde, 0xde, 0x97, 0x87, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x63, 0x31, 0xab, 0xcc, 0xc8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CounterClient is the client API for Counter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CounterClient interface {
	IncrementCounter(ctx context.Context, in *MsgCounter, opts ...grpc.CallOption) (*MsgCreateCounterResponse, error)
}

type counterClient struct {
	cc grpc1.ClientConn
}

func NewCounterClient(cc grpc1.ClientConn) CounterClient {
	return &counterClient{cc}
}

func (c *counterClient) IncrementCounter(ctx context.Context, in *MsgCounter, opts ...grpc.CallOption) (*MsgCreateCounterResponse, error) {
	out := new(MsgCreateCounterResponse)
	err := c.cc.Invoke(ctx, "/testdata.Counter/IncrementCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CounterServer is the server API for Counter service.
type CounterServer interface {
	IncrementCounter(context.Context, *MsgCounter) (*MsgCreateCounterResponse, error)
}

// UnimplementedCounterServer can be embedded to have forward compatible implementations.
type UnimplementedCounterServer struct {
}

func (*UnimplementedCounterServer) IncrementCounter(ctx context.Context, req *MsgCounter) (*MsgCreateCounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrementCounter not implemented")
}

func RegisterCounterServer(s grpc1.Server, srv CounterServer) {
	s.RegisterService(&_Counter_serviceDesc, srv)
}

func _Counter_IncrementCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCounter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServer).IncrementCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testdata.Counter/IncrementCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServer).IncrementCounter(ctx, req.(*MsgCounter))
	}
	return interceptor(ctx, in, info, handler)
}

var _Counter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testdata.Counter",
	HandlerType: (*CounterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IncrementCounter",
			Handler:    _Counter_IncrementCounter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messages.proto",
}

// Counter2Client is the client API for Counter2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Counter2Client interface {
	IncrementCounter(ctx context.Context, in *MsgCounter2, opts ...grpc.CallOption) (*MsgCreateCounterResponse, error)
}

type counter2Client struct {
	cc grpc1.ClientConn
}

func NewCounter2Client(cc grpc1.ClientConn) Counter2Client {
	return &counter2Client{cc}
}

func (c *counter2Client) IncrementCounter(ctx context.Context, in *MsgCounter2, opts ...grpc.CallOption) (*MsgCreateCounterResponse, error) {
	out := new(MsgCreateCounterResponse)
	err := c.cc.Invoke(ctx, "/testdata.Counter2/IncrementCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Counter2Server is the server API for Counter2 service.
type Counter2Server interface {
	IncrementCounter(context.Context, *MsgCounter2) (*MsgCreateCounterResponse, error)
}

// UnimplementedCounter2Server can be embedded to have forward compatible implementations.
type UnimplementedCounter2Server struct {
}

func (*UnimplementedCounter2Server) IncrementCounter(ctx context.Context, req *MsgCounter2) (*MsgCreateCounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrementCounter not implemented")
}

func RegisterCounter2Server(s grpc1.Server, srv Counter2Server) {
	s.RegisterService(&_Counter2_serviceDesc, srv)
}

func _Counter2_IncrementCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCounter2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Counter2Server).IncrementCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testdata.Counter2/IncrementCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Counter2Server).IncrementCounter(ctx, req.(*MsgCounter2))
	}
	return interceptor(ctx, in, info, handler)
}

var _Counter2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testdata.Counter2",
	HandlerType: (*Counter2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IncrementCounter",
			Handler:    _Counter2_IncrementCounter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messages.proto",
}

// KeyValueClient is the client API for KeyValue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeyValueClient interface {
	Set(ctx context.Context, in *MsgKeyValue, opts ...grpc.CallOption) (*MsgCreateKeyValueResponse, error)
}

type keyValueClient struct {
	cc grpc1.ClientConn
}

func NewKeyValueClient(cc grpc1.ClientConn) KeyValueClient {
	return &keyValueClient{cc}
}

func (c *keyValueClient) Set(ctx context.Context, in *MsgKeyValue, opts ...grpc.CallOption) (*MsgCreateKeyValueResponse, error) {
	out := new(MsgCreateKeyValueResponse)
	err := c.cc.Invoke(ctx, "/testdata.KeyValue/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyValueServer is the server API for KeyValue service.
type KeyValueServer interface {
	Set(context.Context, *MsgKeyValue) (*MsgCreateKeyValueResponse, error)
}

// UnimplementedKeyValueServer can be embedded to have forward compatible implementations.
type UnimplementedKeyValueServer struct {
}

func (*UnimplementedKeyValueServer) Set(ctx context.Context, req *MsgKeyValue) (*MsgCreateKeyValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}

func RegisterKeyValueServer(s grpc1.Server, srv KeyValueServer) {
	s.RegisterService(&_KeyValue_serviceDesc, srv)
}

func _KeyValue_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgKeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testdata.KeyValue/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServer).Set(ctx, req.(*MsgKeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyValue_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testdata.KeyValue",
	HandlerType: (*KeyValueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _KeyValue_Set_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messages.proto",
}

func (m *MsgCounter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCounter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCounter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FailOnHandler {
		i--
		if m.FailOnHandler {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Counter != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.Counter))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgCounter2) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCounter2) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCounter2) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FailOnHandler {
		i--
		if m.FailOnHandler {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Counter != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.Counter))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateCounterResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateCounterResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateCounterResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgKeyValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgKeyValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgKeyValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateKeyValueResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateKeyValueResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateKeyValueResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMessages(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessages(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCounter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Counter != 0 {
		n += 1 + sovMessages(uint64(m.Counter))
	}
	if m.FailOnHandler {
		n += 2
	}
	return n
}

func (m *MsgCounter2) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Counter != 0 {
		n += 1 + sovMessages(uint64(m.Counter))
	}
	if m.FailOnHandler {
		n += 2
	}
	return n
}

func (m *MsgCreateCounterResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgKeyValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *MsgCreateKeyValueResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMessages(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessages(x uint64) (n int) {
	return sovMessages(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCounter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCounter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCounter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counter", wireType)
			}
			m.Counter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Counter |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailOnHandler", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FailOnHandler = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCounter2) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCounter2: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCounter2: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counter", wireType)
			}
			m.Counter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Counter |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailOnHandler", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FailOnHandler = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateCounterResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateCounterResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateCounterResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgKeyValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgKeyValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgKeyValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateKeyValueResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateKeyValueResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateKeyValueResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMessages(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessages
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMessages
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessages
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessages
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessages        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessages          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessages = fmt.Errorf("proto: unexpected end of group")
)
