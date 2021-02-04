// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oracle/v1beta/querier.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_tendermint_tendermint_libs_bytes "github.com/tendermint/tendermint/libs/bytes"
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

// QueryParamsRequest is the request type for the Query/Params RPC method
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d18edc48ded26bce, []int{0}
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

// QueryParamsResponse is the response type for the Query/Params RPC method
type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d18edc48ded26bce, []int{1}
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

// QueryClaimRequest is the request type for the Query/Claim RPC method
type QueryClaimRequest struct {
	ClaimHash github_com_tendermint_tendermint_libs_bytes.HexBytes `protobuf:"bytes,1,opt,name=claim_hash,json=claimHash,proto3,casttype=github.com/tendermint/tendermint/libs/bytes.HexBytes" json:"claim_hash,omitempty"`
}

func (m *QueryClaimRequest) Reset()         { *m = QueryClaimRequest{} }
func (m *QueryClaimRequest) String() string { return proto.CompactTextString(m) }
func (*QueryClaimRequest) ProtoMessage()    {}
func (*QueryClaimRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d18edc48ded26bce, []int{2}
}
func (m *QueryClaimRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryClaimRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryClaimRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryClaimRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryClaimRequest.Merge(m, src)
}
func (m *QueryClaimRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryClaimRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryClaimRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryClaimRequest proto.InternalMessageInfo

func (m *QueryClaimRequest) GetClaimHash() github_com_tendermint_tendermint_libs_bytes.HexBytes {
	if m != nil {
		return m.ClaimHash
	}
	return nil
}

// QueryClaimResponse is the response type for the Query/Claim RPC method.
type QueryClaimResponse struct {
	// claim returns the requested claim.
	Claim *types.Any `protobuf:"bytes,1,opt,name=claim,proto3" json:"claim,omitempty"`
}

func (m *QueryClaimResponse) Reset()         { *m = QueryClaimResponse{} }
func (m *QueryClaimResponse) String() string { return proto.CompactTextString(m) }
func (*QueryClaimResponse) ProtoMessage()    {}
func (*QueryClaimResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d18edc48ded26bce, []int{3}
}
func (m *QueryClaimResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryClaimResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryClaimResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryClaimResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryClaimResponse.Merge(m, src)
}
func (m *QueryClaimResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryClaimResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryClaimResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryClaimResponse proto.InternalMessageInfo

func (m *QueryClaimResponse) GetClaim() *types.Any {
	if m != nil {
		return m.Claim
	}
	return nil
}

// QueryAllClaimRequest is the request type for the Query/AllClaim RPC method
type QueryAllClaimRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllClaimRequest) Reset()         { *m = QueryAllClaimRequest{} }
func (m *QueryAllClaimRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllClaimRequest) ProtoMessage()    {}
func (*QueryAllClaimRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d18edc48ded26bce, []int{4}
}
func (m *QueryAllClaimRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllClaimRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllClaimRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllClaimRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllClaimRequest.Merge(m, src)
}
func (m *QueryAllClaimRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllClaimRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllClaimRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllClaimRequest proto.InternalMessageInfo

func (m *QueryAllClaimRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryAllClaimResponse is the response type for the Query/Claim RPC method.
type QueryAllClaimResponse struct {
	Claim      []*types.Any        `protobuf:"bytes,1,rep,name=claim,proto3" json:"claim,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllClaimResponse) Reset()         { *m = QueryAllClaimResponse{} }
func (m *QueryAllClaimResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllClaimResponse) ProtoMessage()    {}
func (*QueryAllClaimResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d18edc48ded26bce, []int{5}
}
func (m *QueryAllClaimResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllClaimResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllClaimResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllClaimResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllClaimResponse.Merge(m, src)
}
func (m *QueryAllClaimResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllClaimResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllClaimResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllClaimResponse proto.InternalMessageInfo

func (m *QueryAllClaimResponse) GetClaim() []*types.Any {
	if m != nil {
		return m.Claim
	}
	return nil
}

func (m *QueryAllClaimResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "r3l.oracle.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "r3l.oracle.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QueryClaimRequest)(nil), "r3l.oracle.v1beta1.QueryClaimRequest")
	proto.RegisterType((*QueryClaimResponse)(nil), "r3l.oracle.v1beta1.QueryClaimResponse")
	proto.RegisterType((*QueryAllClaimRequest)(nil), "r3l.oracle.v1beta1.QueryAllClaimRequest")
	proto.RegisterType((*QueryAllClaimResponse)(nil), "r3l.oracle.v1beta1.QueryAllClaimResponse")
}

func init() { proto.RegisterFile("oracle/v1beta/querier.proto", fileDescriptor_d18edc48ded26bce) }

var fileDescriptor_d18edc48ded26bce = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0x8e, 0x03, 0x89, 0xe0, 0x60, 0xe1, 0x08, 0x12, 0x98, 0xca, 0x45, 0x96, 0x48, 0x43, 0x25,
	0xee, 0x48, 0xc3, 0xd0, 0x91, 0x06, 0x09, 0x2a, 0x16, 0xc0, 0x0b, 0x12, 0x03, 0xd5, 0xd9, 0x1c,
	0x8e, 0xa5, 0xf3, 0x9d, 0xeb, 0x3b, 0x57, 0xb5, 0x50, 0x17, 0x06, 0x26, 0x06, 0x24, 0xfe, 0x54,
	0xc7, 0x4a, 0x2c, 0x4c, 0x15, 0x4a, 0xf8, 0x01, 0xcc, 0x4c, 0xc8, 0x77, 0x17, 0xd5, 0x6e, 0x1a,
	0xa5, 0xdb, 0xbb, 0x7b, 0xdf, 0xfb, 0xde, 0xf7, 0xde, 0xf7, 0xc0, 0x7d, 0x91, 0x93, 0x88, 0x51,
	0x7c, 0x30, 0x0c, 0xa9, 0x22, 0x78, 0xbf, 0xa0, 0x79, 0x42, 0x73, 0x94, 0xe5, 0x42, 0x09, 0x08,
	0xf3, 0x11, 0x43, 0x06, 0x80, 0x0c, 0x60, 0xe8, 0x6e, 0x46, 0x42, 0xa6, 0x42, 0xe2, 0x90, 0x48,
	0xaa, 0xe1, 0xa5, 0xad, 0x1d, 0xe2, 0x8c, 0xc4, 0x09, 0x27, 0x2a, 0x11, 0xdc, 0xd4, 0xbb, 0xbd,
	0x58, 0xc4, 0x42, 0x87, 0xb8, 0x8a, 0xec, 0xef, 0xbd, 0x58, 0x88, 0x98, 0x51, 0xac, 0x5f, 0x61,
	0xf1, 0x09, 0x13, 0x5e, 0xda, 0xd4, 0x9a, 0x4d, 0x91, 0x2c, 0xc1, 0x84, 0x73, 0xa1, 0x34, 0x9b,
	0xb4, 0x59, 0xb7, 0xa9, 0x35, 0x23, 0x39, 0x49, 0x6d, 0xce, 0xef, 0x01, 0xf8, 0xb6, 0x12, 0xf3,
	0x46, 0x7f, 0x06, 0x74, 0xbf, 0xa0, 0x52, 0xf9, 0xaf, 0xc1, 0xed, 0xc6, 0xaf, 0xcc, 0x04, 0x97,
	0x14, 0x6e, 0x83, 0xae, 0x29, 0xbe, 0xeb, 0x3c, 0x70, 0x06, 0x37, 0xb6, 0x5c, 0xb4, 0x38, 0x28,
	0x32, 0x35, 0xe3, 0xab, 0xc7, 0xa7, 0xeb, 0xad, 0xc0, 0xe2, 0x7d, 0x06, 0x6e, 0x69, 0xc2, 0xe7,
	0x8c, 0x24, 0xa9, 0xed, 0x02, 0xdf, 0x01, 0x10, 0x55, 0xef, 0xbd, 0x09, 0x91, 0x13, 0x4d, 0x79,
	0x73, 0xbc, 0xfd, 0xef, 0x74, 0xfd, 0x69, 0x9c, 0xa8, 0x49, 0x11, 0xa2, 0x48, 0xa4, 0x58, 0x51,
	0xfe, 0x91, 0xe6, 0x69, 0xc2, 0x55, 0x3d, 0x64, 0x49, 0x28, 0x71, 0x58, 0x2a, 0x2a, 0xd1, 0x2e,
	0x3d, 0x1c, 0x57, 0x41, 0x70, 0x5d, 0x73, 0xed, 0x12, 0x39, 0xf1, 0x9f, 0xd9, 0xa1, 0x6c, 0x37,
	0xab, 0x7e, 0x13, 0x74, 0x34, 0xc4, 0x8a, 0xef, 0x21, 0xb3, 0x34, 0x34, 0xdf, 0x27, 0xda, 0xe1,
	0x65, 0x60, 0x20, 0xfe, 0x07, 0xd0, 0xd3, 0x0c, 0x3b, 0x8c, 0x35, 0x24, 0xbf, 0x00, 0xe0, 0xcc,
	0x2d, 0x4b, 0xd4, 0x47, 0xc6, 0x5a, 0x54, 0x59, 0x8b, 0xb4, 0xb5, 0xb5, 0x65, 0xc4, 0xd4, 0xd6,
	0x06, 0xb5, 0x4a, 0xff, 0x9b, 0x03, 0xee, 0x9c, 0x6b, 0xb0, 0xa8, 0xf2, 0xca, 0x0a, 0x95, 0xf0,
	0x65, 0x43, 0x4d, 0x5b, 0xab, 0xd9, 0x58, 0xa9, 0xc6, 0x34, 0xaa, 0xcb, 0xd9, 0xfa, 0xdb, 0x06,
	0x1d, 0x2d, 0x07, 0x1e, 0x81, 0xae, 0x31, 0x10, 0xf6, 0x2f, 0x32, 0x77, 0xf1, 0x56, 0xdc, 0x8d,
	0x95, 0x38, 0xd3, 0xd0, 0xf7, 0xbf, 0xfc, 0xfc, 0xf3, 0xa3, 0xbd, 0x06, 0x5d, 0x9c, 0x8f, 0x18,
	0x6e, 0xdc, 0xe4, 0xd0, 0x1e, 0x25, 0xdc, 0x03, 0xd7, 0xe6, 0x1b, 0x81, 0x83, 0xa5, 0xc4, 0xe7,
	0x5c, 0x71, 0x1f, 0x5d, 0x02, 0x69, 0xd7, 0xfb, 0xd5, 0x01, 0x1d, 0x43, 0xff, 0x70, 0x69, 0x51,
	0x83, 0xbb, 0xbf, 0x0a, 0x66, 0xa7, 0x43, 0x7a, 0xba, 0x01, 0xec, 0x5f, 0x34, 0x9d, 0xb6, 0x0b,
	0x7f, 0x3e, 0xbb, 0xf6, 0xa3, 0xf1, 0xab, 0xe3, 0xa9, 0xe7, 0x9c, 0x4c, 0x3d, 0xe7, 0xf7, 0xd4,
	0x73, 0xbe, 0xcf, 0xbc, 0xd6, 0xc9, 0xcc, 0x6b, 0xfd, 0x9a, 0x79, 0xad, 0xf7, 0x4f, 0x6a, 0xe7,
	0x9f, 0x53, 0x46, 0x0f, 0x08, 0x57, 0x8f, 0x23, 0x91, 0xa6, 0x05, 0x4f, 0x54, 0xa9, 0xe9, 0x0f,
	0xe7, 0x0d, 0x54, 0x99, 0x51, 0x19, 0x76, 0xf5, 0x71, 0x8c, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff,
	0xf4, 0x4e, 0xd7, 0x2d, 0x95, 0x04, 0x00, 0x00,
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
	// Params queries the parameters of slashing module
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// this line is used by starport scaffolding # 2
	AllClaim(ctx context.Context, in *QueryAllClaimRequest, opts ...grpc.CallOption) (*QueryAllClaimResponse, error)
	// Claim queries claims based on claim hash.
	Claim(ctx context.Context, in *QueryClaimRequest, opts ...grpc.CallOption) (*QueryClaimResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/r3l.oracle.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllClaim(ctx context.Context, in *QueryAllClaimRequest, opts ...grpc.CallOption) (*QueryAllClaimResponse, error) {
	out := new(QueryAllClaimResponse)
	err := c.cc.Invoke(ctx, "/r3l.oracle.v1beta1.Query/AllClaim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Claim(ctx context.Context, in *QueryClaimRequest, opts ...grpc.CallOption) (*QueryClaimResponse, error) {
	out := new(QueryClaimResponse)
	err := c.cc.Invoke(ctx, "/r3l.oracle.v1beta1.Query/Claim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params queries the parameters of slashing module
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// this line is used by starport scaffolding # 2
	AllClaim(context.Context, *QueryAllClaimRequest) (*QueryAllClaimResponse, error)
	// Claim queries claims based on claim hash.
	Claim(context.Context, *QueryClaimRequest) (*QueryClaimResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) AllClaim(ctx context.Context, req *QueryAllClaimRequest) (*QueryAllClaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllClaim not implemented")
}
func (*UnimplementedQueryServer) Claim(ctx context.Context, req *QueryClaimRequest) (*QueryClaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Claim not implemented")
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
		FullMethod: "/r3l.oracle.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllClaimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/r3l.oracle.v1beta1.Query/AllClaim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllClaim(ctx, req.(*QueryAllClaimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Claim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClaimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Claim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/r3l.oracle.v1beta1.Query/Claim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Claim(ctx, req.(*QueryClaimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "r3l.oracle.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "AllClaim",
			Handler:    _Query_AllClaim_Handler,
		},
		{
			MethodName: "Claim",
			Handler:    _Query_Claim_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oracle/v1beta/querier.proto",
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
		i = encodeVarintQuerier(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryClaimRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryClaimRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryClaimRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimHash) > 0 {
		i -= len(m.ClaimHash)
		copy(dAtA[i:], m.ClaimHash)
		i = encodeVarintQuerier(dAtA, i, uint64(len(m.ClaimHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryClaimResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryClaimResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryClaimResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Claim != nil {
		{
			size, err := m.Claim.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuerier(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllClaimRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllClaimRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllClaimRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintQuerier(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllClaimResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllClaimResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllClaimResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintQuerier(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Claim) > 0 {
		for iNdEx := len(m.Claim) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Claim[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuerier(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuerier(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuerier(v)
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
	n += 1 + l + sovQuerier(uint64(l))
	return n
}

func (m *QueryClaimRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClaimHash)
	if l > 0 {
		n += 1 + l + sovQuerier(uint64(l))
	}
	return n
}

func (m *QueryClaimResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Claim != nil {
		l = m.Claim.Size()
		n += 1 + l + sovQuerier(uint64(l))
	}
	return n
}

func (m *QueryAllClaimRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuerier(uint64(l))
	}
	return n
}

func (m *QueryAllClaimResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Claim) > 0 {
		for _, e := range m.Claim {
			l = e.Size()
			n += 1 + l + sovQuerier(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuerier(uint64(l))
	}
	return n
}

func sovQuerier(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuerier(x uint64) (n int) {
	return sovQuerier(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuerier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuerier
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
				return ErrIntOverflowQuerier
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
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
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
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuerier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QueryClaimRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryClaimRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryClaimRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimHash = append(m.ClaimHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ClaimHash == nil {
				m.ClaimHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuerier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QueryClaimResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryClaimResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryClaimResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Claim == nil {
				m.Claim = &types.Any{}
			}
			if err := m.Claim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuerier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QueryAllClaimRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryAllClaimRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllClaimRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
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
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuerier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QueryAllClaimResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryAllClaimResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllClaimResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Claim = append(m.Claim, &types.Any{})
			if err := m.Claim[len(m.Claim)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
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
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuerier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func skipQuerier(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuerier
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
					return 0, ErrIntOverflowQuerier
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
					return 0, ErrIntOverflowQuerier
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
				return 0, ErrInvalidLengthQuerier
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuerier
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuerier
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuerier        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuerier          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuerier = fmt.Errorf("proto: unexpected end of group")
)