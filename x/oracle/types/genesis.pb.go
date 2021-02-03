// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oracle/v1beta/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
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

// GenesisState defines the oracle module's genesis state.
type GenesisState struct {
	// params defines all the paramaters of related to deposit.
	Params     Params       `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	RoundVotes []RoundVotes `protobuf:"bytes,2,rep,name=roundVotes,proto3" json:"roundVotes"`
	Claims     []*types.Any `protobuf:"bytes,3,rep,name=claims,proto3" json:"claims,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e69d96eeb50dae05, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "r3l.oracle.v1beta1.GenesisState")
}

func init() { proto.RegisterFile("oracle/v1beta/genesis.proto", fileDescriptor_e69d96eeb50dae05) }

var fileDescriptor_e69d96eeb50dae05 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4b, 0x03, 0x31,
	0x18, 0x86, 0x2f, 0x56, 0x8a, 0xa4, 0x4e, 0x47, 0x87, 0xf3, 0x84, 0xb4, 0x38, 0x75, 0xd0, 0xc4,
	0xb6, 0x8b, 0xb8, 0x59, 0x04, 0xc1, 0x49, 0x2a, 0x38, 0xb8, 0xe5, 0xce, 0x18, 0x0f, 0x72, 0xf9,
	0x4a, 0x2e, 0x57, 0xbc, 0x7f, 0xe0, 0xe8, 0x4f, 0xe8, 0xaf, 0x91, 0x8e, 0x1d, 0x9d, 0x44, 0x7a,
	0x8b, 0x3f, 0x43, 0xbc, 0x44, 0xb1, 0xe8, 0x96, 0xf0, 0x3e, 0xef, 0xf7, 0x7e, 0xef, 0x87, 0xf7,
	0xc1, 0xf0, 0x54, 0x09, 0x36, 0x1f, 0x26, 0xc2, 0x72, 0x26, 0x85, 0x16, 0x45, 0x56, 0xd0, 0x99,
	0x01, 0x0b, 0x61, 0x68, 0xc6, 0x8a, 0x3a, 0x80, 0x3a, 0x60, 0x18, 0x77, 0x25, 0x48, 0x68, 0x64,
	0xf6, 0xf5, 0x72, 0x64, 0xbc, 0x27, 0x01, 0xa4, 0x12, 0xac, 0xf9, 0x25, 0xe5, 0x3d, 0xe3, 0xba,
	0xf2, 0x52, 0xbc, 0x99, 0x30, 0xe3, 0x86, 0xe7, 0x3e, 0x20, 0x8e, 0x36, 0xb5, 0x39, 0x58, 0xe1,
	0x94, 0x83, 0x17, 0x84, 0x77, 0x2f, 0xdc, 0x32, 0xd7, 0x96, 0x5b, 0x11, 0x9e, 0xe0, 0xb6, 0xb3,
	0x46, 0xa8, 0x8f, 0x06, 0x9d, 0x51, 0x4c, 0xff, 0x2e, 0x47, 0xaf, 0x1a, 0x62, 0xb2, 0xbd, 0x7c,
	0xeb, 0x05, 0x53, 0xcf, 0x87, 0xe7, 0x18, 0x1b, 0x28, 0xf5, 0xdd, 0x0d, 0x58, 0x51, 0x44, 0x5b,
	0xfd, 0xd6, 0xa0, 0x33, 0x22, 0xff, 0xb9, 0xa7, 0x3f, 0x94, 0x9f, 0xf0, 0xcb, 0x17, 0x1e, 0xe2,
	0x76, 0xaa, 0x78, 0x96, 0x17, 0x51, 0xab, 0x99, 0xd0, 0xa5, 0xae, 0x32, 0xfd, 0xae, 0x4c, 0xcf,
	0x74, 0x35, 0xf5, 0xcc, 0xe9, 0xce, 0xd3, 0xa2, 0x17, 0x7c, 0x2c, 0x7a, 0xc1, 0xe4, 0x72, 0xb9,
	0x26, 0x68, 0xb5, 0x26, 0xe8, 0x7d, 0x4d, 0xd0, 0x73, 0x4d, 0x82, 0x55, 0x4d, 0x82, 0xd7, 0x9a,
	0x04, 0xb7, 0xc7, 0x32, 0xb3, 0x0f, 0x65, 0x42, 0x53, 0xc8, 0x99, 0x11, 0x4a, 0xcc, 0xb9, 0xb6,
	0x47, 0x29, 0xe4, 0x79, 0xa9, 0x33, 0x5b, 0x31, 0x33, 0x56, 0xec, 0x91, 0xf9, 0x03, 0xd9, 0x6a,
	0x26, 0x8a, 0xa4, 0xdd, 0x64, 0x8d, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xba, 0xe0, 0x4e, 0x21,
	0xb5, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Claims) > 0 {
		for iNdEx := len(m.Claims) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Claims[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.RoundVotes) > 0 {
		for iNdEx := len(m.RoundVotes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RoundVotes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.RoundVotes) > 0 {
		for _, e := range m.RoundVotes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Claims) > 0 {
		for _, e := range m.Claims {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoundVotes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoundVotes = append(m.RoundVotes, RoundVotes{})
			if err := m.RoundVotes[len(m.RoundVotes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claims", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Claims = append(m.Claims, &types.Any{})
			if err := m.Claims[len(m.Claims)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
