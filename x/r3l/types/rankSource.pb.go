// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: r3l/v1beta/rankSource.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type MsgRankSource struct {
	Creator github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty"`
	Account string                                        `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (m *MsgRankSource) Reset()         { *m = MsgRankSource{} }
func (m *MsgRankSource) String() string { return proto.CompactTextString(m) }
func (*MsgRankSource) ProtoMessage()    {}
func (*MsgRankSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ebc70fa092f0b52, []int{0}
}
func (m *MsgRankSource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRankSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRankSource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRankSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRankSource.Merge(m, src)
}
func (m *MsgRankSource) XXX_Size() int {
	return m.Size()
}
func (m *MsgRankSource) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRankSource.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRankSource proto.InternalMessageInfo

func (m *MsgRankSource) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *MsgRankSource) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgRankSource)(nil), "r3l.r3l.v1beta1.MsgRankSource")
}

func init() { proto.RegisterFile("r3l/v1beta/rankSource.proto", fileDescriptor_1ebc70fa092f0b52) }

var fileDescriptor_1ebc70fa092f0b52 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x32, 0xce, 0xd1,
	0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0xd4, 0x2f, 0x4a, 0xcc, 0xcb, 0x0e, 0xce, 0x2f, 0x2d, 0x4a,
	0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2f, 0x32, 0xce, 0xd1, 0x03, 0x61, 0x88,
	0x02, 0x43, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0, 0x9c, 0x3e, 0x88, 0x05, 0x51, 0xa6, 0x54,
	0xc6, 0xc5, 0xeb, 0x5b, 0x9c, 0x1e, 0x04, 0xd7, 0x2d, 0xe4, 0xcd, 0xc5, 0x9e, 0x5c, 0x94, 0x9a,
	0x58, 0x92, 0x5f, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe3, 0x64, 0xf8, 0xeb, 0x9e, 0xbc, 0x6e,
	0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x72, 0x7e, 0x71, 0x6e, 0x7e,
	0x31, 0x94, 0xd2, 0x2d, 0x4e, 0xc9, 0xd6, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x73, 0x4c, 0x4e,
	0x76, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x0e, 0x82, 0x99, 0x20, 0x24, 0xc1, 0xc5, 0x9e, 0x98,
	0x9c, 0x9c, 0x5f, 0x9a, 0x57, 0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe3, 0x3a, 0xb9,
	0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb,
	0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x14, 0xb2, 0x5d, 0x45, 0xa9, 0x39,
	0xa9, 0x65, 0x89, 0x79, 0x25, 0xba, 0xc9, 0xf9, 0xb9, 0xb9, 0xa5, 0x79, 0x99, 0x25, 0x95, 0xfa,
	0x20, 0x3f, 0x57, 0x80, 0x49, 0xb0, 0xb5, 0x49, 0x6c, 0x60, 0x7f, 0x18, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x41, 0x83, 0x44, 0x1e, 0x0d, 0x01, 0x00, 0x00,
}

func (m *MsgRankSource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRankSource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRankSource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintRankSource(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintRankSource(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRankSource(dAtA []byte, offset int, v uint64) int {
	offset -= sovRankSource(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRankSource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovRankSource(uint64(l))
	}
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovRankSource(uint64(l))
	}
	return n
}

func sovRankSource(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRankSource(x uint64) (n int) {
	return sovRankSource(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRankSource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRankSource
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
			return fmt.Errorf("proto: MsgRankSource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRankSource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRankSource
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
				return ErrInvalidLengthRankSource
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRankSource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRankSource
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
				return ErrInvalidLengthRankSource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRankSource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRankSource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRankSource
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRankSource
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
func skipRankSource(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRankSource
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
					return 0, ErrIntOverflowRankSource
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
					return 0, ErrIntOverflowRankSource
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
				return 0, ErrInvalidLengthRankSource
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRankSource
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRankSource
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRankSource        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRankSource          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRankSource = fmt.Errorf("proto: unexpected end of group")
)
