// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sequencer/sequencers_by_rollapp.proto

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

// SequencersByRollapp defines an map between rollappId to a list of
// all sequencers that belongs to it.
type SequencersByRollapp struct {
	// rollappId is the unique identifier of the rollapp chain.
	// The rollappId follows the same standard as cosmos chain_id.
	RollappId string `protobuf:"bytes,1,opt,name=rollappId,proto3" json:"rollappId,omitempty"`
	// list of sequencers' account address
	Sequencers Sequencers `protobuf:"bytes,2,opt,name=sequencers,proto3" json:"sequencers"`
}

func (m *SequencersByRollapp) Reset()         { *m = SequencersByRollapp{} }
func (m *SequencersByRollapp) String() string { return proto.CompactTextString(m) }
func (*SequencersByRollapp) ProtoMessage()    {}
func (*SequencersByRollapp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e660bdc020803b1f, []int{0}
}
func (m *SequencersByRollapp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SequencersByRollapp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SequencersByRollapp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SequencersByRollapp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SequencersByRollapp.Merge(m, src)
}
func (m *SequencersByRollapp) XXX_Size() int {
	return m.Size()
}
func (m *SequencersByRollapp) XXX_DiscardUnknown() {
	xxx_messageInfo_SequencersByRollapp.DiscardUnknown(m)
}

var xxx_messageInfo_SequencersByRollapp proto.InternalMessageInfo

func (m *SequencersByRollapp) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *SequencersByRollapp) GetSequencers() Sequencers {
	if m != nil {
		return m.Sequencers
	}
	return Sequencers{}
}

func init() {
	proto.RegisterType((*SequencersByRollapp)(nil), "dymensionxyz.dymension.sequencer.SequencersByRollapp")
}

func init() {
	proto.RegisterFile("sequencer/sequencers_by_rollapp.proto", fileDescriptor_e660bdc020803b1f)
}

var fileDescriptor_e660bdc020803b1f = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x4e, 0x2d, 0x2c,
	0x4d, 0xcd, 0x4b, 0x4e, 0x2d, 0xd2, 0x87, 0xb3, 0x8a, 0xe3, 0x93, 0x2a, 0xe3, 0x8b, 0xf2, 0x73,
	0x72, 0x12, 0x0b, 0x0a, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x14, 0x52, 0x2a, 0x73, 0x53,
	0xf3, 0x8a, 0x33, 0xf3, 0xf3, 0x2a, 0x2a, 0xab, 0xf4, 0xe0, 0x1c, 0x3d, 0xb8, 0x1e, 0x29, 0x29,
	0x6c, 0x06, 0x41, 0x74, 0x4b, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x99, 0xfa, 0x20, 0x16, 0x44,
	0x54, 0xa9, 0x9d, 0x91, 0x4b, 0x38, 0x18, 0xae, 0xd4, 0xa9, 0x32, 0x08, 0x62, 0xa3, 0x90, 0x0c,
	0x17, 0x27, 0xd4, 0x72, 0xcf, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x84, 0x80, 0x50,
	0x10, 0x17, 0x17, 0xc2, 0x7c, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x1d, 0x3d, 0x42, 0xce,
	0xd3, 0x43, 0xb2, 0x88, 0xe5, 0xc4, 0x3d, 0x79, 0x86, 0x20, 0x24, 0x53, 0x9c, 0x7c, 0x4f, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e,
	0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x38, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49,
	0x2f, 0x39, 0x3f, 0x57, 0x1f, 0xd9, 0x0e, 0x04, 0x47, 0xbf, 0x02, 0xe1, 0x5f, 0xfd, 0x92, 0xca,
	0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xff, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x3b,
	0xf3, 0x33, 0x5c, 0x01, 0x00, 0x00,
}

func (m *SequencersByRollapp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SequencersByRollapp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SequencersByRollapp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Sequencers.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSequencersByRollapp(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintSequencersByRollapp(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSequencersByRollapp(dAtA []byte, offset int, v uint64) int {
	offset -= sovSequencersByRollapp(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SequencersByRollapp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovSequencersByRollapp(uint64(l))
	}
	l = m.Sequencers.Size()
	n += 1 + l + sovSequencersByRollapp(uint64(l))
	return n
}

func sovSequencersByRollapp(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSequencersByRollapp(x uint64) (n int) {
	return sovSequencersByRollapp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SequencersByRollapp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSequencersByRollapp
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
			return fmt.Errorf("proto: SequencersByRollapp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SequencersByRollapp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencersByRollapp
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
				return ErrInvalidLengthSequencersByRollapp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSequencersByRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequencers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencersByRollapp
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
				return ErrInvalidLengthSequencersByRollapp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSequencersByRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Sequencers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSequencersByRollapp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSequencersByRollapp
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
func skipSequencersByRollapp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSequencersByRollapp
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
					return 0, ErrIntOverflowSequencersByRollapp
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
					return 0, ErrIntOverflowSequencersByRollapp
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
				return 0, ErrInvalidLengthSequencersByRollapp
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSequencersByRollapp
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSequencersByRollapp
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSequencersByRollapp        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSequencersByRollapp          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSequencersByRollapp = fmt.Errorf("proto: unexpected end of group")
)
