// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: r3l/v1beta/querier.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("r3l/v1beta/querier.proto", fileDescriptor_bc3e27884096409a) }

var fileDescriptor_bc3e27884096409a = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x32, 0xce, 0xd1,
	0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0xd4, 0x2f, 0x2c, 0x4d, 0x2d, 0xca, 0x4c, 0x2d, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2f, 0x32, 0xce, 0xd1, 0x03, 0x61, 0x88, 0xac, 0xa1, 0x94,
	0x56, 0x72, 0x7e, 0x71, 0x6e, 0x7e, 0xb1, 0x7e, 0x52, 0x62, 0x71, 0x2a, 0x58, 0x6d, 0x25, 0x54,
	0xa3, 0xa1, 0x7e, 0x41, 0x62, 0x7a, 0x66, 0x5e, 0x62, 0x49, 0x66, 0x7e, 0x1e, 0x44, 0xb3, 0x11,
	0x3b, 0x17, 0x6b, 0x20, 0x48, 0x85, 0x93, 0xfb, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31,
	0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb,
	0x31, 0x44, 0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x17, 0xa5,
	0xe6, 0xa4, 0x96, 0x25, 0xe6, 0x95, 0xe8, 0x26, 0xe7, 0xe7, 0xe6, 0x96, 0xe6, 0x65, 0x96, 0x54,
	0xea, 0x83, 0xdc, 0x55, 0x01, 0x26, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x06, 0x1b,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x14, 0x40, 0x0a, 0x59, 0xb1, 0x00, 0x00, 0x00,
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
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "r3l.r3l.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "r3l/v1beta/querier.proto",
}
