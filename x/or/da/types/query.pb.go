// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: finschia/or/da/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/Finschia/finschia-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b019ae279835818, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b019ae279835818, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryCCBatchesRequest struct {
	RollupName string             `protobuf:"bytes,1,opt,name=rollup_name,json=rollupName,proto3" json:"rollup_name,omitempty"`
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryCCBatchesRequest) Reset()         { *m = QueryCCBatchesRequest{} }
func (m *QueryCCBatchesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCCBatchesRequest) ProtoMessage()    {}
func (*QueryCCBatchesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b019ae279835818, []int{2}
}
func (m *QueryCCBatchesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCCBatchesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCCBatchesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCCBatchesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCCBatchesRequest.Merge(m, src)
}
func (m *QueryCCBatchesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCCBatchesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCCBatchesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCCBatchesRequest proto.InternalMessageInfo

func (m *QueryCCBatchesRequest) GetRollupName() string {
	if m != nil {
		return m.RollupName
	}
	return ""
}

func (m *QueryCCBatchesRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryCCBatchesResponse struct {
	BatchesLoc []*CCRef            `protobuf:"bytes,1,rep,name=batches_loc,json=batchesLoc,proto3" json:"batches_loc,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryCCBatchesResponse) Reset()         { *m = QueryCCBatchesResponse{} }
func (m *QueryCCBatchesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCCBatchesResponse) ProtoMessage()    {}
func (*QueryCCBatchesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b019ae279835818, []int{3}
}
func (m *QueryCCBatchesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCCBatchesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCCBatchesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCCBatchesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCCBatchesResponse.Merge(m, src)
}
func (m *QueryCCBatchesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCCBatchesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCCBatchesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCCBatchesResponse proto.InternalMessageInfo

func (m *QueryCCBatchesResponse) GetBatchesLoc() []*CCRef {
	if m != nil {
		return m.BatchesLoc
	}
	return nil
}

func (m *QueryCCBatchesResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryCCBatchRequest struct {
	RollupName  string `protobuf:"bytes,1,opt,name=rollup_name,json=rollupName,proto3" json:"rollup_name,omitempty"`
	BatchHeight uint64 `protobuf:"varint,2,opt,name=batch_height,json=batchHeight,proto3" json:"batch_height,omitempty"`
}

func (m *QueryCCBatchRequest) Reset()         { *m = QueryCCBatchRequest{} }
func (m *QueryCCBatchRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCCBatchRequest) ProtoMessage()    {}
func (*QueryCCBatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b019ae279835818, []int{4}
}
func (m *QueryCCBatchRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCCBatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCCBatchRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCCBatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCCBatchRequest.Merge(m, src)
}
func (m *QueryCCBatchRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCCBatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCCBatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCCBatchRequest proto.InternalMessageInfo

func (m *QueryCCBatchRequest) GetRollupName() string {
	if m != nil {
		return m.RollupName
	}
	return ""
}

func (m *QueryCCBatchRequest) GetBatchHeight() uint64 {
	if m != nil {
		return m.BatchHeight
	}
	return 0
}

type QueryCCBatchResponse struct {
	BatchLoc *CCRef `protobuf:"bytes,1,opt,name=batch_loc,json=batchLoc,proto3" json:"batch_loc,omitempty"`
}

func (m *QueryCCBatchResponse) Reset()         { *m = QueryCCBatchResponse{} }
func (m *QueryCCBatchResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCCBatchResponse) ProtoMessage()    {}
func (*QueryCCBatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b019ae279835818, []int{5}
}
func (m *QueryCCBatchResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCCBatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCCBatchResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCCBatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCCBatchResponse.Merge(m, src)
}
func (m *QueryCCBatchResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCCBatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCCBatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCCBatchResponse proto.InternalMessageInfo

func (m *QueryCCBatchResponse) GetBatchLoc() *CCRef {
	if m != nil {
		return m.BatchLoc
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "finschia.or.da.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "finschia.or.da.v1.QueryParamsResponse")
	proto.RegisterType((*QueryCCBatchesRequest)(nil), "finschia.or.da.v1.QueryCCBatchesRequest")
	proto.RegisterType((*QueryCCBatchesResponse)(nil), "finschia.or.da.v1.QueryCCBatchesResponse")
	proto.RegisterType((*QueryCCBatchRequest)(nil), "finschia.or.da.v1.QueryCCBatchRequest")
	proto.RegisterType((*QueryCCBatchResponse)(nil), "finschia.or.da.v1.QueryCCBatchResponse")
}

func init() { proto.RegisterFile("finschia/or/da/v1/query.proto", fileDescriptor_6b019ae279835818) }

var fileDescriptor_6b019ae279835818 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xc1, 0x6f, 0x12, 0x4f,
	0x14, 0xc7, 0x99, 0xb6, 0x3f, 0x7e, 0xf2, 0xf0, 0xe2, 0x88, 0x86, 0xa2, 0x5d, 0xda, 0x4d, 0x4a,
	0xd1, 0xe8, 0x4c, 0x40, 0x8d, 0xf1, 0x0a, 0x49, 0xed, 0xa1, 0x36, 0x75, 0x6f, 0x7a, 0x21, 0xc3,
	0x32, 0xdd, 0xdd, 0xc8, 0xee, 0x6c, 0x77, 0x16, 0x62, 0x35, 0x26, 0xc6, 0x8b, 0xd7, 0x26, 0x5e,
	0x3d, 0xf8, 0xe7, 0xf4, 0xd8, 0xc4, 0x8b, 0x27, 0x63, 0xc0, 0x3f, 0xc4, 0x30, 0x33, 0xab, 0xa5,
	0xd0, 0xd2, 0x1b, 0x99, 0xf7, 0x7d, 0xdf, 0xf7, 0xf9, 0xbe, 0x7d, 0x01, 0xd6, 0x0e, 0x82, 0x48,
	0xba, 0x7e, 0xc0, 0xa8, 0x48, 0x68, 0x8f, 0xd1, 0x61, 0x83, 0x1e, 0x0e, 0x78, 0x72, 0x44, 0xe2,
	0x44, 0xa4, 0x02, 0xdf, 0xc8, 0xca, 0x44, 0x24, 0xa4, 0xc7, 0xc8, 0xb0, 0x51, 0x29, 0x79, 0xc2,
	0x13, 0xaa, 0x4a, 0x27, 0xbf, 0xb4, 0xb0, 0x72, 0xd7, 0x13, 0xc2, 0xeb, 0x73, 0xca, 0xe2, 0x80,
	0xb2, 0x28, 0x12, 0x29, 0x4b, 0x03, 0x11, 0x49, 0x53, 0xbd, 0xef, 0x0a, 0x19, 0x0a, 0x49, 0xbb,
	0x4c, 0x72, 0xed, 0x4f, 0x87, 0x8d, 0x2e, 0x4f, 0x59, 0x83, 0xc6, 0xcc, 0x0b, 0x22, 0x25, 0x36,
	0x5a, 0x6b, 0x96, 0x28, 0x66, 0x09, 0x0b, 0x33, 0xaf, 0xca, 0x6c, 0xbd, 0xc7, 0x74, 0xcd, 0x2e,
	0x01, 0x7e, 0x39, 0x71, 0xdf, 0x57, 0x0d, 0x0e, 0x3f, 0x1c, 0x70, 0x99, 0xda, 0x7b, 0x70, 0x73,
	0xea, 0x55, 0xc6, 0x22, 0x92, 0x1c, 0x3f, 0x85, 0xbc, 0x36, 0x2e, 0xa3, 0x75, 0x54, 0x2f, 0x36,
	0x57, 0xc9, 0x4c, 0x58, 0xa2, 0x5b, 0x5a, 0x2b, 0x27, 0x3f, 0xab, 0x39, 0xc7, 0xc8, 0xed, 0x8f,
	0x08, 0x6e, 0x29, 0xc3, 0x76, 0xbb, 0xc5, 0x52, 0xd7, 0xe7, 0xd9, 0x24, 0x5c, 0x85, 0x62, 0x22,
	0xfa, 0xfd, 0x41, 0xdc, 0x89, 0x58, 0xc8, 0x95, 0x6f, 0xc1, 0x01, 0xfd, 0xb4, 0xc7, 0x42, 0x8e,
	0xb7, 0x01, 0xfe, 0x05, 0x2e, 0x2f, 0xa9, 0xb9, 0x35, 0xa2, 0xb7, 0x43, 0x26, 0xdb, 0x21, 0x7a,
	0xfb, 0x66, 0x3b, 0x64, 0x9f, 0x79, 0xdc, 0x98, 0x3b, 0x67, 0x3a, 0xed, 0xaf, 0x08, 0x6e, 0x9f,
	0x47, 0x30, 0xb1, 0x9e, 0x41, 0xb1, 0xab, 0x9f, 0x3a, 0x7d, 0xe1, 0x96, 0xd1, 0xfa, 0x72, 0xbd,
	0xd8, 0x2c, 0xcf, 0xc9, 0xd6, 0x6e, 0x3b, 0xfc, 0xc0, 0x01, 0x23, 0xde, 0x15, 0x2e, 0x7e, 0x3e,
	0x87, 0x6e, 0x6b, 0x21, 0x9d, 0x9e, 0x3b, 0x85, 0xf7, 0xca, 0x6c, 0xdc, 0xd0, 0x5d, 0x79, 0x3d,
	0x1b, 0x70, 0x5d, 0xe1, 0x74, 0x7c, 0x1e, 0x78, 0x7e, 0xaa, 0x10, 0x56, 0x1c, 0x9d, 0x67, 0x47,
	0x3d, 0xd9, 0x2f, 0xa0, 0x34, 0x6d, 0x6d, 0x62, 0x3f, 0x81, 0x82, 0x6e, 0xd5, 0xa1, 0xd1, 0xa5,
	0xa1, 0xaf, 0x29, 0xe9, 0xae, 0x70, 0x9b, 0xdf, 0x96, 0xe1, 0x3f, 0xe5, 0x87, 0xdf, 0x41, 0x5e,
	0x7f, 0x6d, 0xbc, 0x39, 0xa7, 0x6f, 0xf6, 0xac, 0x2a, 0xb5, 0x45, 0x32, 0x4d, 0x66, 0x6f, 0x7c,
	0xfa, 0xfe, 0xfb, 0xcb, 0xd2, 0x1d, 0xbc, 0x4a, 0x2f, 0xba, 0x6c, 0xfc, 0x19, 0x41, 0xe1, 0xef,
	0x97, 0xc4, 0xf5, 0x8b, 0x8c, 0xcf, 0xdf, 0x5b, 0xe5, 0xde, 0x15, 0x94, 0x86, 0x62, 0x53, 0x51,
	0x54, 0xf1, 0xda, 0x1c, 0x0a, 0x73, 0x02, 0xd4, 0x75, 0xf1, 0x31, 0x82, 0xff, 0x4d, 0x33, 0xae,
	0x2d, 0x70, 0xcf, 0x28, 0xb6, 0x16, 0xea, 0x0c, 0xc3, 0x63, 0xc5, 0x40, 0xf0, 0x83, 0x4b, 0x19,
	0xe8, 0xfb, 0xb3, 0x37, 0xf0, 0xa1, 0xb5, 0x73, 0x32, 0xb2, 0xd0, 0xe9, 0xc8, 0x42, 0xbf, 0x46,
	0x16, 0x3a, 0x1e, 0x5b, 0xb9, 0xd3, 0xb1, 0x95, 0xfb, 0x31, 0xb6, 0x72, 0xaf, 0x89, 0x17, 0xa4,
	0xfe, 0xa0, 0x4b, 0x5c, 0x11, 0xd2, 0xed, 0xcc, 0x31, 0xb3, 0x7e, 0x28, 0x7b, 0x6f, 0xe8, 0x5b,
	0x33, 0x20, 0x3d, 0x8a, 0xb9, 0xec, 0xe6, 0xd5, 0xbf, 0xc4, 0xa3, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x90, 0xee, 0x76, 0x11, 0xf5, 0x04, 0x00, 0x00,
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
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	CCBatches(ctx context.Context, in *QueryCCBatchesRequest, opts ...grpc.CallOption) (*QueryCCBatchesResponse, error)
	CCBatch(ctx context.Context, in *QueryCCBatchRequest, opts ...grpc.CallOption) (*QueryCCBatchResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/finschia.or.da.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CCBatches(ctx context.Context, in *QueryCCBatchesRequest, opts ...grpc.CallOption) (*QueryCCBatchesResponse, error) {
	out := new(QueryCCBatchesResponse)
	err := c.cc.Invoke(ctx, "/finschia.or.da.v1.Query/CCBatches", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CCBatch(ctx context.Context, in *QueryCCBatchRequest, opts ...grpc.CallOption) (*QueryCCBatchResponse, error) {
	out := new(QueryCCBatchResponse)
	err := c.cc.Invoke(ctx, "/finschia.or.da.v1.Query/CCBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	CCBatches(context.Context, *QueryCCBatchesRequest) (*QueryCCBatchesResponse, error)
	CCBatch(context.Context, *QueryCCBatchRequest) (*QueryCCBatchResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) CCBatches(ctx context.Context, req *QueryCCBatchesRequest) (*QueryCCBatchesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CCBatches not implemented")
}
func (*UnimplementedQueryServer) CCBatch(ctx context.Context, req *QueryCCBatchRequest) (*QueryCCBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CCBatch not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/finschia.or.da.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CCBatches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCCBatchesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CCBatches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/finschia.or.da.v1.Query/CCBatches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CCBatches(ctx, req.(*QueryCCBatchesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CCBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCCBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CCBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/finschia.or.da.v1.Query/CCBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CCBatch(ctx, req.(*QueryCCBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "finschia.or.da.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "CCBatches",
			Handler:    _Query_CCBatches_Handler,
		},
		{
			MethodName: "CCBatch",
			Handler:    _Query_CCBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "finschia/or/da/v1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryCCBatchesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCCBatchesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCCBatchesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.RollupName) > 0 {
		i -= len(m.RollupName)
		copy(dAtA[i:], m.RollupName)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.RollupName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCCBatchesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCCBatchesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCCBatchesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.BatchesLoc) > 0 {
		for iNdEx := len(m.BatchesLoc) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BatchesLoc[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryCCBatchRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCCBatchRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCCBatchRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BatchHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.BatchHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.RollupName) > 0 {
		i -= len(m.RollupName)
		copy(dAtA[i:], m.RollupName)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.RollupName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCCBatchResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCCBatchResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCCBatchResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BatchLoc != nil {
		{
			size, err := m.BatchLoc.MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryCCBatchesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RollupName)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCCBatchesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.BatchesLoc) > 0 {
		for _, e := range m.BatchesLoc {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCCBatchRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RollupName)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.BatchHeight != 0 {
		n += 1 + sovQuery(uint64(m.BatchHeight))
	}
	return n
}

func (m *QueryCCBatchResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BatchLoc != nil {
		l = m.BatchLoc.Size()
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
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryCCBatchesRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCCBatchesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCCBatchesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollupName", wireType)
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
			m.RollupName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryCCBatchesResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCCBatchesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCCBatchesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchesLoc", wireType)
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
			m.BatchesLoc = append(m.BatchesLoc, &CCRef{})
			if err := m.BatchesLoc[len(m.BatchesLoc)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryCCBatchRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCCBatchRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCCBatchRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollupName", wireType)
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
			m.RollupName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchHeight", wireType)
			}
			m.BatchHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BatchHeight |= uint64(b&0x7F) << shift
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
func (m *QueryCCBatchResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCCBatchResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCCBatchResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchLoc", wireType)
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
			if m.BatchLoc == nil {
				m.BatchLoc = &CCRef{}
			}
			if err := m.BatchLoc.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
