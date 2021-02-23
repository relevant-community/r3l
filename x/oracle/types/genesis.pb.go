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
	Params            Params                             `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Claims            []*types.Any                       `protobuf:"bytes,2,rep,name=claims,proto3" json:"claims,omitempty"`
	Rounds            []Round                            `protobuf:"bytes,3,rep,name=rounds,proto3" json:"rounds"`
	Pending           map[string]GenesisState_ListOfUint `protobuf:"bytes,4,rep,name=pending,proto3" json:"pending" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FeederDelegations []MsgDelegateFeedConsent           `protobuf:"bytes,5,rep,name=feeder_delegations,json=feederDelegations,proto3" json:"feeder_delegations"`
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

type GenesisState_ListOfUint struct {
	Pending []uint64 `protobuf:"varint,1,rep,packed,name=pending,proto3" json:"pending,omitempty"`
}

func (m *GenesisState_ListOfUint) Reset()         { *m = GenesisState_ListOfUint{} }
func (m *GenesisState_ListOfUint) String() string { return proto.CompactTextString(m) }
func (*GenesisState_ListOfUint) ProtoMessage()    {}
func (*GenesisState_ListOfUint) Descriptor() ([]byte, []int) {
	return fileDescriptor_e69d96eeb50dae05, []int{0, 0}
}
func (m *GenesisState_ListOfUint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState_ListOfUint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState_ListOfUint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState_ListOfUint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState_ListOfUint.Merge(m, src)
}
func (m *GenesisState_ListOfUint) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState_ListOfUint) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState_ListOfUint.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState_ListOfUint proto.InternalMessageInfo

func (m *GenesisState_ListOfUint) GetPending() []uint64 {
	if m != nil {
		return m.Pending
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "r3l.oracle.v1beta1.GenesisState")
	proto.RegisterMapType((map[string]GenesisState_ListOfUint)(nil), "r3l.oracle.v1beta1.GenesisState.PendingEntry")
	proto.RegisterType((*GenesisState_ListOfUint)(nil), "r3l.oracle.v1beta1.GenesisState.ListOfUint")
}

func init() { proto.RegisterFile("oracle/v1beta/genesis.proto", fileDescriptor_e69d96eeb50dae05) }

var fileDescriptor_e69d96eeb50dae05 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xb1, 0x6f, 0x13, 0x31,
	0x14, 0xc6, 0xef, 0x92, 0x34, 0x80, 0xdb, 0x01, 0xac, 0x0a, 0x5d, 0x0f, 0xe9, 0x12, 0x31, 0x45,
	0xd0, 0xda, 0xb4, 0x19, 0xa8, 0xba, 0xb5, 0x14, 0x90, 0x10, 0x15, 0xd5, 0x21, 0x16, 0x96, 0xca,
	0xc9, 0xbd, 0x98, 0x13, 0x77, 0x76, 0x64, 0xfb, 0xa2, 0xde, 0xc6, 0xc8, 0xc8, 0x9f, 0xd0, 0x3f,
	0xa7, 0x63, 0x47, 0x26, 0x84, 0x92, 0x85, 0x3f, 0x03, 0xc5, 0x76, 0x94, 0x06, 0x22, 0x75, 0x73,
	0xf2, 0xfd, 0xde, 0xf7, 0x7d, 0x4f, 0xef, 0xd0, 0x13, 0xa9, 0xd8, 0xb0, 0x00, 0x3a, 0xd9, 0x1f,
	0x80, 0x61, 0x94, 0x83, 0x00, 0x9d, 0x6b, 0x32, 0x56, 0xd2, 0x48, 0x8c, 0x55, 0xbf, 0x20, 0x0e,
	0x20, 0x0e, 0xd8, 0x8f, 0xb7, 0xb9, 0xe4, 0xd2, 0xca, 0x74, 0xfe, 0x72, 0x64, 0xbc, 0xc3, 0xa5,
	0xe4, 0x05, 0x50, 0xfb, 0x6b, 0x50, 0x8d, 0x28, 0x13, 0xb5, 0x97, 0xe2, 0xd5, 0x84, 0x31, 0x53,
	0xac, 0xf4, 0x01, 0x71, 0xb4, 0xaa, 0x4d, 0xa4, 0x01, 0xaf, 0x3c, 0x5e, 0x55, 0xcc, 0xa5, 0xfb,
	0xff, 0xe9, 0xb7, 0x16, 0xda, 0x7a, 0xeb, 0x4a, 0x7e, 0x34, 0xcc, 0x00, 0x3e, 0x44, 0x6d, 0x67,
	0x19, 0x85, 0xdd, 0xb0, 0xb7, 0x79, 0x10, 0x93, 0xff, 0x4b, 0x93, 0x73, 0x4b, 0x9c, 0xb4, 0xae,
	0x7f, 0x75, 0x82, 0xd4, 0xf3, 0x78, 0x17, 0xb5, 0x87, 0x05, 0xcb, 0x4b, 0x1d, 0x35, 0xba, 0xcd,
	0xde, 0xe6, 0xc1, 0x36, 0x71, 0x4b, 0x90, 0xc5, 0x12, 0xe4, 0x58, 0xd4, 0xa9, 0x67, 0xf0, 0x4b,
	0xd4, 0x56, 0xb2, 0x12, 0x99, 0x8e, 0x9a, 0x96, 0xde, 0x59, 0x97, 0x93, 0xce, 0x89, 0x45, 0x8c,
	0xc3, 0xf1, 0x19, 0xba, 0x37, 0x06, 0x91, 0xe5, 0x82, 0x47, 0x2d, 0x3b, 0xb9, 0xb7, 0x6e, 0xf2,
	0xf6, 0x4e, 0xe4, 0xdc, 0xf1, 0xaf, 0x85, 0x51, 0xb5, 0x77, 0x5b, 0x78, 0xe0, 0x0b, 0x84, 0x47,
	0x00, 0x19, 0xa8, 0x8b, 0x0c, 0x0a, 0xe0, 0xcc, 0xe4, 0x52, 0xe8, 0x68, 0xc3, 0x3a, 0x3f, 0x5b,
	0xe7, 0x7c, 0xa6, 0xf9, 0xa9, 0x23, 0xe1, 0x0d, 0x40, 0xf6, 0x4a, 0x0a, 0x0d, 0xc2, 0x78, 0xdb,
	0x47, 0xce, 0xeb, 0x74, 0x69, 0x15, 0xef, 0x22, 0xf4, 0x3e, 0xd7, 0xe6, 0xc3, 0xe8, 0x53, 0x2e,
	0x0c, 0x4e, 0x96, 0xed, 0xc3, 0x6e, 0xb3, 0xd7, 0xfa, 0xa7, 0x4e, 0xcc, 0xd1, 0xd6, 0xed, 0xb6,
	0xf8, 0x21, 0x6a, 0x7e, 0x85, 0xda, 0xde, 0xe2, 0x41, 0x3a, 0x7f, 0xe2, 0x63, 0xb4, 0x31, 0x61,
	0x45, 0x05, 0x51, 0xc3, 0xde, 0xe7, 0xf9, 0x9d, 0xdb, 0x2f, 0xd3, 0x53, 0x37, 0x79, 0xd4, 0x38,
	0x0c, 0x8f, 0xee, 0x7f, 0xbf, 0xea, 0x04, 0x7f, 0xae, 0x3a, 0xc1, 0xc9, 0xbb, 0xeb, 0x69, 0x12,
	0xde, 0x4c, 0x93, 0xf0, 0xf7, 0x34, 0x09, 0x7f, 0xcc, 0x92, 0xe0, 0x66, 0x96, 0x04, 0x3f, 0x67,
	0x49, 0xf0, 0xf9, 0x05, 0xcf, 0xcd, 0x97, 0x6a, 0x40, 0x86, 0xb2, 0xa4, 0x0a, 0x0a, 0x98, 0x30,
	0x61, 0xf6, 0x86, 0xb2, 0x2c, 0x2b, 0x91, 0x9b, 0x9a, 0xaa, 0x7e, 0x41, 0x2f, 0xa9, 0xff, 0xb0,
	0x4c, 0x3d, 0x06, 0x3d, 0x68, 0xdb, 0x5b, 0xf7, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x18,
	0xff, 0x2e, 0x07, 0x03, 0x00, 0x00,
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
	if len(m.FeederDelegations) > 0 {
		for iNdEx := len(m.FeederDelegations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeederDelegations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Pending) > 0 {
		for k := range m.Pending {
			v := m.Pending[k]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintGenesis(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenesis(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Rounds) > 0 {
		for iNdEx := len(m.Rounds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rounds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *GenesisState_ListOfUint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState_ListOfUint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState_ListOfUint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Pending) > 0 {
		dAtA4 := make([]byte, len(m.Pending)*10)
		var j3 int
		for _, num := range m.Pending {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintGenesis(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0xa
	}
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
	if len(m.Claims) > 0 {
		for _, e := range m.Claims {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Rounds) > 0 {
		for _, e := range m.Rounds {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Pending) > 0 {
		for k, v := range m.Pending {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovGenesis(uint64(len(k))) + 1 + l + sovGenesis(uint64(l))
			n += mapEntrySize + 1 + sovGenesis(uint64(mapEntrySize))
		}
	}
	if len(m.FeederDelegations) > 0 {
		for _, e := range m.FeederDelegations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisState_ListOfUint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Pending) > 0 {
		l = 0
		for _, e := range m.Pending {
			l += sovGenesis(uint64(e))
		}
		n += 1 + sovGenesis(uint64(l)) + l
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rounds", wireType)
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
			m.Rounds = append(m.Rounds, Round{})
			if err := m.Rounds[len(m.Rounds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pending", wireType)
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
			if m.Pending == nil {
				m.Pending = make(map[string]GenesisState_ListOfUint)
			}
			var mapkey string
			mapvalue := &GenesisState_ListOfUint{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthGenesis
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthGenesis
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &GenesisState_ListOfUint{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenesis(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthGenesis
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Pending[mapkey] = *mapvalue
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeederDelegations", wireType)
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
			m.FeederDelegations = append(m.FeederDelegations, MsgDelegateFeedConsent{})
			if err := m.FeederDelegations[len(m.FeederDelegations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *GenesisState_ListOfUint) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ListOfUint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListOfUint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Pending = append(m.Pending, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGenesis
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGenesis
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Pending) == 0 {
					m.Pending = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Pending = append(m.Pending, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Pending", wireType)
			}
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
