// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: r3l/v1beta/vote.proto

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

type MsgVote struct {
	Id      string                                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty"`
	To      string                                        `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Amount  int32                                         `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *MsgVote) Reset()         { *m = MsgVote{} }
func (m *MsgVote) String() string { return proto.CompactTextString(m) }
func (*MsgVote) ProtoMessage()    {}
func (*MsgVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1837f1c345dbfcd, []int{0}
}
func (m *MsgVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVote.Merge(m, src)
}
func (m *MsgVote) XXX_Size() int {
	return m.Size()
}
func (m *MsgVote) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgVote.DiscardUnknown(m)
}

var xxx_messageInfo_MsgVote proto.InternalMessageInfo

func (m *MsgVote) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgVote) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *MsgVote) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *MsgVote) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgVote)(nil), "r3l.r3l.v1beta1.MsgVote")
}

func init() { proto.RegisterFile("r3l/v1beta/vote.proto", fileDescriptor_d1837f1c345dbfcd) }

var fileDescriptor_d1837f1c345dbfcd = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x32, 0xce, 0xd1,
	0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0xd4, 0x2f, 0xcb, 0x2f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x2f, 0x32, 0xce, 0xd1, 0x03, 0x61, 0x88, 0x94, 0xa1, 0x94, 0x48, 0x7a, 0x7e,
	0x7a, 0x3e, 0x58, 0x4e, 0x1f, 0xc4, 0x82, 0x28, 0x53, 0xea, 0x63, 0xe4, 0x62, 0xf7, 0x2d, 0x4e,
	0x0f, 0xcb, 0x2f, 0x49, 0x15, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x62, 0xca, 0x4c, 0x11, 0xf2, 0xe6, 0x62, 0x4f, 0x2e, 0x4a, 0x4d, 0x2c, 0xc9, 0x2f, 0x92,
	0x60, 0x52, 0x60, 0xd4, 0xe0, 0x71, 0x32, 0xfc, 0x75, 0x4f, 0x5e, 0x37, 0x3d, 0xb3, 0x24, 0xa3,
	0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x39, 0xbf, 0x38, 0x37, 0xbf, 0x18, 0x4a, 0xe9, 0x16,
	0xa7, 0x64, 0xeb, 0x97, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0x39, 0x26, 0x27, 0x3b, 0xa6, 0xa4, 0x14,
	0xa5, 0x16, 0x17, 0x07, 0xc1, 0x4c, 0x00, 0x19, 0x5e, 0x92, 0x2f, 0xc1, 0x0c, 0x31, 0xbc, 0x24,
	0x5f, 0x48, 0x8c, 0x8b, 0x2d, 0x31, 0x37, 0xbf, 0x34, 0xaf, 0x44, 0x82, 0x45, 0x81, 0x51, 0x83,
	0x35, 0x08, 0xca, 0x73, 0x72, 0x3f, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f,
	0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28,
	0x64, 0x9b, 0x8b, 0x52, 0x73, 0x52, 0xcb, 0x12, 0xf3, 0x4a, 0x74, 0x93, 0xf3, 0x73, 0x73, 0x4b,
	0xf3, 0x32, 0x4b, 0x2a, 0xf5, 0x41, 0xc1, 0x50, 0x01, 0x26, 0xc1, 0x8e, 0x48, 0x62, 0x03, 0x7b,
	0xd0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x53, 0xb0, 0x1a, 0x20, 0x01, 0x00, 0x00,
}

func (m *MsgVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintVote(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintVote(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintVote(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintVote(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVote(dAtA []byte, offset int, v uint64) int {
	offset -= sovVote(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovVote(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovVote(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovVote(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovVote(uint64(m.Amount))
	}
	return n
}

func sovVote(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVote(x uint64) (n int) {
	return sovVote(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVote
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
			return fmt.Errorf("proto: MsgVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVote
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
				return ErrInvalidLengthVote
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVote
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
				return ErrInvalidLengthVote
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVote
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
				return ErrInvalidLengthVote
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVote
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVote
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
func skipVote(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVote
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
					return 0, ErrIntOverflowVote
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
					return 0, ErrIntOverflowVote
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
				return 0, ErrInvalidLengthVote
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVote
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVote
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVote        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVote          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVote = fmt.Errorf("proto: unexpected end of group")
)
