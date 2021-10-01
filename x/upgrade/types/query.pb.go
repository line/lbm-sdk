// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lbm/upgrade/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/line/lbm-sdk/codec/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryCurrentPlanRequest is the request type for the Query/CurrentPlan RPC
// method.
type QueryCurrentPlanRequest struct {
}

func (m *QueryCurrentPlanRequest) Reset()         { *m = QueryCurrentPlanRequest{} }
func (m *QueryCurrentPlanRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCurrentPlanRequest) ProtoMessage()    {}
func (*QueryCurrentPlanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be137e92226ff1e7, []int{0}
}
func (m *QueryCurrentPlanRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCurrentPlanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCurrentPlanRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCurrentPlanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCurrentPlanRequest.Merge(m, src)
}
func (m *QueryCurrentPlanRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCurrentPlanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCurrentPlanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCurrentPlanRequest proto.InternalMessageInfo

// QueryCurrentPlanResponse is the response type for the Query/CurrentPlan RPC
// method.
type QueryCurrentPlanResponse struct {
	// plan is the current upgrade plan.
	Plan *Plan `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
}

func (m *QueryCurrentPlanResponse) Reset()         { *m = QueryCurrentPlanResponse{} }
func (m *QueryCurrentPlanResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCurrentPlanResponse) ProtoMessage()    {}
func (*QueryCurrentPlanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_be137e92226ff1e7, []int{1}
}
func (m *QueryCurrentPlanResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCurrentPlanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCurrentPlanResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCurrentPlanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCurrentPlanResponse.Merge(m, src)
}
func (m *QueryCurrentPlanResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCurrentPlanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCurrentPlanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCurrentPlanResponse proto.InternalMessageInfo

func (m *QueryCurrentPlanResponse) GetPlan() *Plan {
	if m != nil {
		return m.Plan
	}
	return nil
}

// QueryCurrentPlanRequest is the request type for the Query/AppliedPlan RPC
// method.
type QueryAppliedPlanRequest struct {
	// name is the name of the applied plan to query for.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *QueryAppliedPlanRequest) Reset()         { *m = QueryAppliedPlanRequest{} }
func (m *QueryAppliedPlanRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAppliedPlanRequest) ProtoMessage()    {}
func (*QueryAppliedPlanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be137e92226ff1e7, []int{2}
}
func (m *QueryAppliedPlanRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAppliedPlanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAppliedPlanRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAppliedPlanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAppliedPlanRequest.Merge(m, src)
}
func (m *QueryAppliedPlanRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAppliedPlanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAppliedPlanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAppliedPlanRequest proto.InternalMessageInfo

func (m *QueryAppliedPlanRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// QueryAppliedPlanResponse is the response type for the Query/AppliedPlan RPC
// method.
type QueryAppliedPlanResponse struct {
	// height is the block height at which the plan was applied.
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *QueryAppliedPlanResponse) Reset()         { *m = QueryAppliedPlanResponse{} }
func (m *QueryAppliedPlanResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAppliedPlanResponse) ProtoMessage()    {}
func (*QueryAppliedPlanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_be137e92226ff1e7, []int{3}
}
func (m *QueryAppliedPlanResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAppliedPlanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAppliedPlanResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAppliedPlanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAppliedPlanResponse.Merge(m, src)
}
func (m *QueryAppliedPlanResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAppliedPlanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAppliedPlanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAppliedPlanResponse proto.InternalMessageInfo

func (m *QueryAppliedPlanResponse) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

// QueryUpgradedConsensusStateRequest is the request type for the Query/UpgradedConsensusState
// RPC method.
type QueryUpgradedConsensusStateRequest struct {
	// last height of the current chain must be sent in request
	// as this is the height under which next consensus state is stored
	LastHeight int64 `protobuf:"varint,1,opt,name=last_height,json=lastHeight,proto3" json:"last_height,omitempty"`
}

func (m *QueryUpgradedConsensusStateRequest) Reset()         { *m = QueryUpgradedConsensusStateRequest{} }
func (m *QueryUpgradedConsensusStateRequest) String() string { return proto.CompactTextString(m) }
func (*QueryUpgradedConsensusStateRequest) ProtoMessage()    {}
func (*QueryUpgradedConsensusStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be137e92226ff1e7, []int{4}
}
func (m *QueryUpgradedConsensusStateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUpgradedConsensusStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUpgradedConsensusStateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUpgradedConsensusStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUpgradedConsensusStateRequest.Merge(m, src)
}
func (m *QueryUpgradedConsensusStateRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryUpgradedConsensusStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUpgradedConsensusStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUpgradedConsensusStateRequest proto.InternalMessageInfo

func (m *QueryUpgradedConsensusStateRequest) GetLastHeight() int64 {
	if m != nil {
		return m.LastHeight
	}
	return 0
}

// QueryUpgradedConsensusStateResponse is the response type for the Query/UpgradedConsensusState
// RPC method.
type QueryUpgradedConsensusStateResponse struct {
	UpgradedConsensusState *types.Any `protobuf:"bytes,1,opt,name=upgraded_consensus_state,json=upgradedConsensusState,proto3" json:"upgraded_consensus_state,omitempty"`
}

func (m *QueryUpgradedConsensusStateResponse) Reset()         { *m = QueryUpgradedConsensusStateResponse{} }
func (m *QueryUpgradedConsensusStateResponse) String() string { return proto.CompactTextString(m) }
func (*QueryUpgradedConsensusStateResponse) ProtoMessage()    {}
func (*QueryUpgradedConsensusStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_be137e92226ff1e7, []int{5}
}
func (m *QueryUpgradedConsensusStateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUpgradedConsensusStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUpgradedConsensusStateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUpgradedConsensusStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUpgradedConsensusStateResponse.Merge(m, src)
}
func (m *QueryUpgradedConsensusStateResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryUpgradedConsensusStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUpgradedConsensusStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUpgradedConsensusStateResponse proto.InternalMessageInfo

func (m *QueryUpgradedConsensusStateResponse) GetUpgradedConsensusState() *types.Any {
	if m != nil {
		return m.UpgradedConsensusState
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryCurrentPlanRequest)(nil), "lbm.upgrade.v1.QueryCurrentPlanRequest")
	proto.RegisterType((*QueryCurrentPlanResponse)(nil), "lbm.upgrade.v1.QueryCurrentPlanResponse")
	proto.RegisterType((*QueryAppliedPlanRequest)(nil), "lbm.upgrade.v1.QueryAppliedPlanRequest")
	proto.RegisterType((*QueryAppliedPlanResponse)(nil), "lbm.upgrade.v1.QueryAppliedPlanResponse")
	proto.RegisterType((*QueryUpgradedConsensusStateRequest)(nil), "lbm.upgrade.v1.QueryUpgradedConsensusStateRequest")
	proto.RegisterType((*QueryUpgradedConsensusStateResponse)(nil), "lbm.upgrade.v1.QueryUpgradedConsensusStateResponse")
}

func init() { proto.RegisterFile("lbm/upgrade/v1/query.proto", fileDescriptor_be137e92226ff1e7) }

var fileDescriptor_be137e92226ff1e7 = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6b, 0x13, 0x41,
	0x00, 0xcd, 0xda, 0x5a, 0x70, 0x02, 0x1e, 0x06, 0x89, 0xe9, 0x52, 0x56, 0x99, 0x2a, 0x0d, 0x48,
	0x67, 0x68, 0x0a, 0xe2, 0x49, 0x88, 0x55, 0xf0, 0x24, 0x1a, 0xf1, 0xe2, 0x25, 0xcc, 0x66, 0xc7,
	0xcd, 0xe2, 0x64, 0x66, 0xba, 0x33, 0x53, 0x0c, 0xa5, 0x17, 0x0f, 0xde, 0x04, 0xc1, 0x5f, 0xe4,
	0xcd, 0x63, 0xc1, 0x8b, 0x47, 0x49, 0xfa, 0x43, 0x64, 0x67, 0x67, 0x4b, 0x3e, 0x36, 0x11, 0x6f,
	0xc9, 0xbe, 0x37, 0xef, 0xbd, 0x9d, 0xf7, 0x16, 0x84, 0x3c, 0x1e, 0x13, 0xab, 0xd2, 0x9c, 0x26,
	0x8c, 0x9c, 0x1d, 0x91, 0x53, 0xcb, 0xf2, 0x09, 0x56, 0xb9, 0x34, 0x12, 0xde, 0xe6, 0xf1, 0x18,
	0x7b, 0x0c, 0x9f, 0x1d, 0x85, 0xbb, 0xa9, 0x94, 0x29, 0x67, 0xc4, 0xa1, 0xb1, 0xfd, 0x40, 0xa8,
	0xf0, 0xd4, 0x70, 0xcf, 0x43, 0x54, 0x65, 0x84, 0x0a, 0x21, 0x0d, 0x35, 0x99, 0x14, 0xba, 0x42,
	0x97, 0x4c, 0x2a, 0x4d, 0x87, 0xa2, 0x5d, 0x70, 0xf7, 0x4d, 0xe1, 0x7a, 0x62, 0xf3, 0x9c, 0x09,
	0xf3, 0x9a, 0x53, 0xd1, 0x67, 0xa7, 0x96, 0x69, 0x83, 0x9e, 0x83, 0xf6, 0x2a, 0xa4, 0x95, 0x14,
	0x9a, 0xc1, 0x0e, 0xd8, 0x56, 0x9c, 0x8a, 0x76, 0x70, 0x3f, 0xe8, 0x34, 0xbb, 0x77, 0xf0, 0x62,
	0x58, 0xec, 0xb8, 0x8e, 0x81, 0x0e, 0xbd, 0x41, 0x4f, 0x29, 0x9e, 0xb1, 0x64, 0xce, 0x00, 0x42,
	0xb0, 0x2d, 0xe8, 0x98, 0x39, 0x91, 0x5b, 0x7d, 0xf7, 0x1b, 0x75, 0xbd, 0xe9, 0x02, 0xdd, 0x9b,
	0xb6, 0xc0, 0xce, 0x88, 0x65, 0xe9, 0xc8, 0xb8, 0x13, 0x5b, 0x7d, 0xff, 0x0f, 0xbd, 0x00, 0xc8,
	0x9d, 0x79, 0x57, 0x06, 0x48, 0x4e, 0x0a, 0xb6, 0xd0, 0x56, 0xbf, 0x35, 0xd4, 0xb0, 0xca, 0xed,
	0x1e, 0x68, 0x72, 0xaa, 0xcd, 0x60, 0x41, 0x02, 0x14, 0x8f, 0x5e, 0x96, 0x32, 0x16, 0xec, 0x6f,
	0x94, 0xf1, 0x29, 0x5e, 0x81, 0xb6, 0x7f, 0xd3, 0x64, 0x30, 0xac, 0x28, 0x03, 0x5d, 0x70, 0xae,
	0xaf, 0xa3, 0x2c, 0x04, 0x57, 0x5d, 0xe1, 0x9e, 0x98, 0xf4, 0x5b, 0xb6, 0x56, 0xb7, 0x7b, 0xb5,
	0x05, 0x6e, 0x3a, 0x5f, 0xf8, 0x25, 0x00, 0xcd, 0xb9, 0xcb, 0x86, 0x07, 0xcb, 0xd7, 0xba, 0xa6,
	0xa9, 0xb0, 0xf3, 0x6f, 0x62, 0x19, 0x1e, 0x3d, 0xf8, 0xfc, 0xeb, 0xea, 0xfb, 0x8d, 0x08, 0xee,
	0x91, 0xa5, 0x55, 0x0c, 0x4b, 0xf2, 0xa0, 0xe8, 0x0c, 0x7e, 0x0d, 0x40, 0x73, 0xae, 0x80, 0x35,
	0x41, 0x56, 0x1b, 0x5d, 0x13, 0xa4, 0xa6, 0x4b, 0xf4, 0xc8, 0x05, 0x79, 0x08, 0xf7, 0x97, 0x83,
	0xd0, 0x92, 0xec, 0x82, 0x90, 0xf3, 0x62, 0x13, 0x17, 0xf0, 0x47, 0x00, 0x5a, 0xf5, 0xad, 0xc0,
	0x6e, 0xad, 0xe3, 0xc6, 0x25, 0x84, 0xc7, 0xff, 0x75, 0xc6, 0x07, 0x7e, 0xea, 0x02, 0x3f, 0x81,
	0x8f, 0x49, 0xfd, 0xf7, 0xb4, 0x32, 0x06, 0x72, 0x3e, 0x37, 0xb7, 0x8b, 0x67, 0xbd, 0x9f, 0xd3,
	0x28, 0xb8, 0x9c, 0x46, 0xc1, 0x9f, 0x69, 0x14, 0x7c, 0x9b, 0x45, 0x8d, 0xcb, 0x59, 0xd4, 0xf8,
	0x3d, 0x8b, 0x1a, 0xef, 0x0f, 0xd2, 0xcc, 0x8c, 0x6c, 0x8c, 0x87, 0x72, 0x4c, 0x78, 0x26, 0x58,
	0x61, 0x70, 0xa8, 0x93, 0x8f, 0xe4, 0xd3, 0xb5, 0x8d, 0x99, 0x28, 0xa6, 0xe3, 0x1d, 0xb7, 0xa7,
	0xe3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x51, 0xe2, 0x6a, 0x37, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// CurrentPlan queries the current upgrade plan.
	CurrentPlan(ctx context.Context, in *QueryCurrentPlanRequest, opts ...grpc.CallOption) (*QueryCurrentPlanResponse, error)
	// AppliedPlan queries a previously applied upgrade plan by its name.
	AppliedPlan(ctx context.Context, in *QueryAppliedPlanRequest, opts ...grpc.CallOption) (*QueryAppliedPlanResponse, error)
	// UpgradedConsensusState queries the consensus state that will serve
	// as a trusted kernel for the next version of this chain. It will only be
	// stored at the last height of this chain.
	// UpgradedConsensusState RPC not supported with legacy querier
	UpgradedConsensusState(ctx context.Context, in *QueryUpgradedConsensusStateRequest, opts ...grpc.CallOption) (*QueryUpgradedConsensusStateResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) CurrentPlan(ctx context.Context, in *QueryCurrentPlanRequest, opts ...grpc.CallOption) (*QueryCurrentPlanResponse, error) {
	out := new(QueryCurrentPlanResponse)
	err := c.cc.Invoke(ctx, "/lbm.upgrade.v1.Query/CurrentPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AppliedPlan(ctx context.Context, in *QueryAppliedPlanRequest, opts ...grpc.CallOption) (*QueryAppliedPlanResponse, error) {
	out := new(QueryAppliedPlanResponse)
	err := c.cc.Invoke(ctx, "/lbm.upgrade.v1.Query/AppliedPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) UpgradedConsensusState(ctx context.Context, in *QueryUpgradedConsensusStateRequest, opts ...grpc.CallOption) (*QueryUpgradedConsensusStateResponse, error) {
	out := new(QueryUpgradedConsensusStateResponse)
	err := c.cc.Invoke(ctx, "/lbm.upgrade.v1.Query/UpgradedConsensusState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// CurrentPlan queries the current upgrade plan.
	CurrentPlan(context.Context, *QueryCurrentPlanRequest) (*QueryCurrentPlanResponse, error)
	// AppliedPlan queries a previously applied upgrade plan by its name.
	AppliedPlan(context.Context, *QueryAppliedPlanRequest) (*QueryAppliedPlanResponse, error)
	// UpgradedConsensusState queries the consensus state that will serve
	// as a trusted kernel for the next version of this chain. It will only be
	// stored at the last height of this chain.
	// UpgradedConsensusState RPC not supported with legacy querier
	UpgradedConsensusState(context.Context, *QueryUpgradedConsensusStateRequest) (*QueryUpgradedConsensusStateResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) CurrentPlan(ctx context.Context, req *QueryCurrentPlanRequest) (*QueryCurrentPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentPlan not implemented")
}
func (*UnimplementedQueryServer) AppliedPlan(ctx context.Context, req *QueryAppliedPlanRequest) (*QueryAppliedPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppliedPlan not implemented")
}
func (*UnimplementedQueryServer) UpgradedConsensusState(ctx context.Context, req *QueryUpgradedConsensusStateRequest) (*QueryUpgradedConsensusStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpgradedConsensusState not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_CurrentPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCurrentPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CurrentPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lbm.upgrade.v1.Query/CurrentPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CurrentPlan(ctx, req.(*QueryCurrentPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AppliedPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAppliedPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AppliedPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lbm.upgrade.v1.Query/AppliedPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AppliedPlan(ctx, req.(*QueryAppliedPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_UpgradedConsensusState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUpgradedConsensusStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).UpgradedConsensusState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lbm.upgrade.v1.Query/UpgradedConsensusState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).UpgradedConsensusState(ctx, req.(*QueryUpgradedConsensusStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lbm.upgrade.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CurrentPlan",
			Handler:    _Query_CurrentPlan_Handler,
		},
		{
			MethodName: "AppliedPlan",
			Handler:    _Query_AppliedPlan_Handler,
		},
		{
			MethodName: "UpgradedConsensusState",
			Handler:    _Query_UpgradedConsensusState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lbm/upgrade/v1/query.proto",
}

func (m *QueryCurrentPlanRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCurrentPlanRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCurrentPlanRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryCurrentPlanResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCurrentPlanResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCurrentPlanResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Plan != nil {
		{
			size, err := m.Plan.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAppliedPlanRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAppliedPlanRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAppliedPlanRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAppliedPlanResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAppliedPlanResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAppliedPlanResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryUpgradedConsensusStateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUpgradedConsensusStateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUpgradedConsensusStateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.LastHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryUpgradedConsensusStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUpgradedConsensusStateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUpgradedConsensusStateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UpgradedConsensusState != nil {
		{
			size, err := m.UpgradedConsensusState.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryCurrentPlanRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryCurrentPlanResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Plan != nil {
		l = m.Plan.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAppliedPlanRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAppliedPlanResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovQuery(uint64(m.Height))
	}
	return n
}

func (m *QueryUpgradedConsensusStateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LastHeight != 0 {
		n += 1 + sovQuery(uint64(m.LastHeight))
	}
	return n
}

func (m *QueryUpgradedConsensusStateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UpgradedConsensusState != nil {
		l = m.UpgradedConsensusState.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryCurrentPlanRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCurrentPlanRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCurrentPlanRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCurrentPlanResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCurrentPlanResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCurrentPlanResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Plan", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Plan == nil {
				m.Plan = &Plan{}
			}
			if err := m.Plan.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAppliedPlanRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAppliedPlanRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAppliedPlanRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAppliedPlanResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAppliedPlanResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAppliedPlanResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryUpgradedConsensusStateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryUpgradedConsensusStateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUpgradedConsensusStateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastHeight", wireType)
			}
			m.LastHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryUpgradedConsensusStateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryUpgradedConsensusStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUpgradedConsensusStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpgradedConsensusState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UpgradedConsensusState == nil {
				m.UpgradedConsensusState = &types.Any{}
			}
			if err := m.UpgradedConsensusState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
