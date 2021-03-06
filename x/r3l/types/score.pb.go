// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: r3l/v1beta/score.proto

package types

import (
	fmt "fmt"
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

type Score struct {
	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BlockHeight int64  `protobuf:"varint,2,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	PRank       int64  `protobuf:"varint,3,opt,name=pRank,proto3" json:"pRank,omitempty"`
	NRank       int64  `protobuf:"varint,4,opt,name=nRank,proto3" json:"nRank,omitempty"`
}

func (m *Score) Reset()         { *m = Score{} }
func (m *Score) String() string { return proto.CompactTextString(m) }
func (*Score) ProtoMessage()    {}
func (*Score) Descriptor() ([]byte, []int) {
	return fileDescriptor_b09ac092a2adcc64, []int{0}
}
func (m *Score) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Score) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Score.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Score) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Score.Merge(m, src)
}
func (m *Score) XXX_Size() int {
	return m.Size()
}
func (m *Score) XXX_DiscardUnknown() {
	xxx_messageInfo_Score.DiscardUnknown(m)
}

var xxx_messageInfo_Score proto.InternalMessageInfo

func (m *Score) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Score) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *Score) GetPRank() int64 {
	if m != nil {
		return m.PRank
	}
	return 0
}

func (m *Score) GetNRank() int64 {
	if m != nil {
		return m.NRank
	}
	return 0
}

type Scores struct {
	Scores      []Score `protobuf:"bytes,1,rep,name=scores,proto3" json:"scores"`
	BlockHeight int64   `protobuf:"varint,2,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
}

func (m *Scores) Reset()         { *m = Scores{} }
func (m *Scores) String() string { return proto.CompactTextString(m) }
func (*Scores) ProtoMessage()    {}
func (*Scores) Descriptor() ([]byte, []int) {
	return fileDescriptor_b09ac092a2adcc64, []int{1}
}
func (m *Scores) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Scores) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Scores.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Scores) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scores.Merge(m, src)
}
func (m *Scores) XXX_Size() int {
	return m.Size()
}
func (m *Scores) XXX_DiscardUnknown() {
	xxx_messageInfo_Scores.DiscardUnknown(m)
}

var xxx_messageInfo_Scores proto.InternalMessageInfo

func (m *Scores) GetScores() []Score {
	if m != nil {
		return m.Scores
	}
	return nil
}

func (m *Scores) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*Score)(nil), "r3l.r3l.v1beta1.Score")
	proto.RegisterType((*Scores)(nil), "r3l.r3l.v1beta1.Scores")
}

func init() { proto.RegisterFile("r3l/v1beta/score.proto", fileDescriptor_b09ac092a2adcc64) }

var fileDescriptor_b09ac092a2adcc64 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x32, 0xce, 0xd1,
	0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0xd4, 0x2f, 0x4e, 0xce, 0x2f, 0x4a, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x2f, 0x32, 0xce, 0xd1, 0x03, 0x61, 0x88, 0x9c, 0xa1, 0x94, 0x48, 0x7a,
	0x7e, 0x7a, 0x3e, 0x58, 0x4e, 0x1f, 0xc4, 0x82, 0x28, 0x53, 0x4a, 0xe5, 0x62, 0x0d, 0x06, 0xe9,
	0x12, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c,
	0x11, 0x52, 0xe0, 0xe2, 0x4e, 0xca, 0xc9, 0x4f, 0xce, 0xf6, 0x48, 0xcd, 0x4c, 0xcf, 0x28, 0x91,
	0x60, 0x52, 0x60, 0xd4, 0x60, 0x0e, 0x42, 0x16, 0x12, 0x12, 0xe1, 0x62, 0x2d, 0x08, 0x4a, 0xcc,
	0xcb, 0x96, 0x60, 0x06, 0xcb, 0x41, 0x38, 0x20, 0xd1, 0x3c, 0xb0, 0x28, 0x0b, 0x44, 0x14, 0xcc,
	0x51, 0x4a, 0xe0, 0x62, 0x03, 0x5b, 0x53, 0x2c, 0x64, 0xc2, 0xc5, 0x06, 0x76, 0x66, 0xb1, 0x04,
	0xa3, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0x98, 0x1e, 0x9a, 0x43, 0xf5, 0xc0, 0x0a, 0x9d, 0x58, 0x4e,
	0xdc, 0x93, 0x67, 0x08, 0x82, 0xaa, 0x25, 0xec, 0x1a, 0x27, 0xf7, 0x13, 0x8f, 0xe4, 0x18, 0x2f,
	0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18,
	0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf,
	0xd5, 0x2f, 0x4a, 0xcd, 0x49, 0x2d, 0x4b, 0xcc, 0x2b, 0xd1, 0x4d, 0xce, 0xcf, 0xcd, 0x2d, 0xcd,
	0xcb, 0x2c, 0xa9, 0xd4, 0x07, 0x85, 0x5f, 0x05, 0x98, 0x2c, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62,
	0x03, 0x07, 0x8c, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x59, 0x73, 0xb9, 0x59, 0x01, 0x00,
	0x00,
}

func (m *Score) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Score) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Score) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NRank != 0 {
		i = encodeVarintScore(dAtA, i, uint64(m.NRank))
		i--
		dAtA[i] = 0x20
	}
	if m.PRank != 0 {
		i = encodeVarintScore(dAtA, i, uint64(m.PRank))
		i--
		dAtA[i] = 0x18
	}
	if m.BlockHeight != 0 {
		i = encodeVarintScore(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintScore(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Scores) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Scores) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Scores) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintScore(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Scores) > 0 {
		for iNdEx := len(m.Scores) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Scores[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintScore(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintScore(dAtA []byte, offset int, v uint64) int {
	offset -= sovScore(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Score) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovScore(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovScore(uint64(m.BlockHeight))
	}
	if m.PRank != 0 {
		n += 1 + sovScore(uint64(m.PRank))
	}
	if m.NRank != 0 {
		n += 1 + sovScore(uint64(m.NRank))
	}
	return n
}

func (m *Scores) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Scores) > 0 {
		for _, e := range m.Scores {
			l = e.Size()
			n += 1 + l + sovScore(uint64(l))
		}
	}
	if m.BlockHeight != 0 {
		n += 1 + sovScore(uint64(m.BlockHeight))
	}
	return n
}

func sovScore(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozScore(x uint64) (n int) {
	return sovScore(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Score) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScore
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
			return fmt.Errorf("proto: Score: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Score: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScore
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
				return ErrInvalidLengthScore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PRank", wireType)
			}
			m.PRank = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PRank |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NRank", wireType)
			}
			m.NRank = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NRank |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipScore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthScore
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthScore
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
func (m *Scores) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScore
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
			return fmt.Errorf("proto: Scores: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Scores: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scores", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScore
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
				return ErrInvalidLengthScore
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthScore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Scores = append(m.Scores, Score{})
			if err := m.Scores[len(m.Scores)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipScore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthScore
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthScore
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
func skipScore(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowScore
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
					return 0, ErrIntOverflowScore
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
					return 0, ErrIntOverflowScore
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
				return 0, ErrInvalidLengthScore
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupScore
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthScore
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthScore        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowScore          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupScore = fmt.Errorf("proto: unexpected end of group")
)
