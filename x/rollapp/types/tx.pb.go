// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rollapp/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/dymensionxyz/dymension/x/sequencer/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgCreateRollapp struct {
	// creator is the bech32-encoded address of the rollapp creator.
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// rollappId is the unique identifier of the rollapp chain.
	// The rollappId follows the same standard as cosmos chain_id.
	RollappId string `protobuf:"bytes,2,opt,name=rollappId,proto3" json:"rollappId,omitempty"`
	// genesisPath is the description of the genesis file location on the DA.
	CodeStamp string `protobuf:"bytes,3,opt,name=codeStamp,proto3" json:"codeStamp,omitempty"`
	// genesisPath is the description of the genesis file location on the DA.
	GenesisPath string `protobuf:"bytes,4,opt,name=genesisPath,proto3" json:"genesisPath,omitempty"`
	// maxWithholdingBlocks is the maximum number of blocks for
	// an active sequencer to send a state update (MsgUpdateState).
	MaxWithholdingBlocks uint64 `protobuf:"varint,5,opt,name=maxWithholdingBlocks,proto3" json:"maxWithholdingBlocks,omitempty"`
	// maxSequencers is the maximum number of sequencers.
	MaxSequencers uint64 `protobuf:"varint,6,opt,name=maxSequencers,proto3" json:"maxSequencers,omitempty"`
	// permissionedAddresses is a bech32-encoded address list of the
	// sequencers that are allowed to serve this rollappId.
	// In the case of an empty list, the rollapp is considered permissionless.
	PermissionedAddresses types.Sequencers `protobuf:"bytes,7,opt,name=permissionedAddresses,proto3" json:"permissionedAddresses"`
}

func (m *MsgCreateRollapp) Reset()         { *m = MsgCreateRollapp{} }
func (m *MsgCreateRollapp) String() string { return proto.CompactTextString(m) }
func (*MsgCreateRollapp) ProtoMessage()    {}
func (*MsgCreateRollapp) Descriptor() ([]byte, []int) {
	return fileDescriptor_28d664456664c256, []int{0}
}
func (m *MsgCreateRollapp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateRollapp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateRollapp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateRollapp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateRollapp.Merge(m, src)
}
func (m *MsgCreateRollapp) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateRollapp) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateRollapp.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateRollapp proto.InternalMessageInfo

func (m *MsgCreateRollapp) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateRollapp) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *MsgCreateRollapp) GetCodeStamp() string {
	if m != nil {
		return m.CodeStamp
	}
	return ""
}

func (m *MsgCreateRollapp) GetGenesisPath() string {
	if m != nil {
		return m.GenesisPath
	}
	return ""
}

func (m *MsgCreateRollapp) GetMaxWithholdingBlocks() uint64 {
	if m != nil {
		return m.MaxWithholdingBlocks
	}
	return 0
}

func (m *MsgCreateRollapp) GetMaxSequencers() uint64 {
	if m != nil {
		return m.MaxSequencers
	}
	return 0
}

func (m *MsgCreateRollapp) GetPermissionedAddresses() types.Sequencers {
	if m != nil {
		return m.PermissionedAddresses
	}
	return types.Sequencers{}
}

type MsgCreateRollappResponse struct {
}

func (m *MsgCreateRollappResponse) Reset()         { *m = MsgCreateRollappResponse{} }
func (m *MsgCreateRollappResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateRollappResponse) ProtoMessage()    {}
func (*MsgCreateRollappResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_28d664456664c256, []int{1}
}
func (m *MsgCreateRollappResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateRollappResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateRollappResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateRollappResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateRollappResponse.Merge(m, src)
}
func (m *MsgCreateRollappResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateRollappResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateRollappResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateRollappResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateRollapp)(nil), "dymensionxyz.dymension.rollapp.MsgCreateRollapp")
	proto.RegisterType((*MsgCreateRollappResponse)(nil), "dymensionxyz.dymension.rollapp.MsgCreateRollappResponse")
}

func init() { proto.RegisterFile("rollapp/tx.proto", fileDescriptor_28d664456664c256) }

var fileDescriptor_28d664456664c256 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xd5, 0xda, 0xae, 0x8d, 0xd7, 0x18, 0xcc, 0xe2, 0xc2, 0x22, 0x8a, 0x2a, 0x4c, 0x0f, 0x3e,
	0x94, 0x55, 0xeb, 0x5e, 0x7a, 0xad, 0x7b, 0x2a, 0xc5, 0x10, 0xe4, 0x43, 0x20, 0x37, 0x59, 0x1a,
	0x24, 0x11, 0x49, 0xab, 0x68, 0x64, 0x90, 0x93, 0x5b, 0x7e, 0x41, 0x7e, 0x96, 0x4f, 0xc1, 0xc7,
	0x9c, 0x42, 0xb0, 0xff, 0x48, 0x90, 0x64, 0xcb, 0x1f, 0x38, 0x81, 0xdc, 0xe6, 0xbd, 0x37, 0xf3,
	0x76, 0xf7, 0xed, 0xd0, 0x5e, 0x22, 0x83, 0xc0, 0x8a, 0x63, 0x23, 0xcd, 0x44, 0x9c, 0xc8, 0x54,
	0x32, 0xcd, 0x59, 0x84, 0x10, 0xa1, 0x2f, 0xa3, 0x6c, 0x71, 0x2b, 0x2a, 0x20, 0xb6, 0x8d, 0xaa,
	0x8a, 0x70, 0x33, 0x87, 0xc8, 0x86, 0xc4, 0xa8, 0x2a, 0x2c, 0x67, 0xd5, 0xbe, 0x2b, 0x5d, 0x59,
	0x94, 0x46, 0x5e, 0x95, 0xec, 0xe0, 0xb1, 0x46, 0x7b, 0x13, 0x74, 0xff, 0x26, 0x60, 0xa5, 0x60,
	0x96, 0x36, 0x8c, 0xd3, 0x96, 0x9d, 0x13, 0x32, 0xe1, 0x44, 0x27, 0xc3, 0xb6, 0xb9, 0x83, 0xec,
	0x0b, 0x6d, 0x6f, 0xcf, 0xfa, 0xe7, 0xf0, 0x5a, 0xa1, 0xed, 0x89, 0x5c, 0xb5, 0xa5, 0x03, 0xd3,
	0xd4, 0x0a, 0x63, 0x5e, 0x2f, 0xd5, 0x8a, 0x60, 0x3a, 0xed, 0xb8, 0x10, 0x01, 0xfa, 0x78, 0x61,
	0xa5, 0x1e, 0x6f, 0x14, 0xfa, 0x21, 0xc5, 0x46, 0xb4, 0x1f, 0x5a, 0xd9, 0xa5, 0x9f, 0x7a, 0x9e,
	0x0c, 0x1c, 0x3f, 0x72, 0xc7, 0x81, 0xb4, 0xaf, 0x91, 0x7f, 0xd2, 0xc9, 0xb0, 0x61, 0x9e, 0xd5,
	0xd8, 0x37, 0xda, 0x0d, 0xad, 0x6c, 0x5a, 0xbd, 0x96, 0x37, 0x8b, 0xe6, 0x63, 0x92, 0x79, 0xf4,
	0x73, 0x0c, 0x49, 0xe8, 0x63, 0x1e, 0x17, 0x38, 0x7f, 0x1c, 0x27, 0x01, 0x44, 0x40, 0xde, 0xd2,
	0xc9, 0xb0, 0x33, 0xfa, 0x2e, 0xde, 0x08, 0xb6, 0x4a, 0x51, 0xec, 0xcd, 0xc6, 0x8d, 0xe5, 0xf3,
	0x57, 0xc5, 0x3c, 0x6f, 0x38, 0x50, 0x29, 0x3f, 0xcd, 0xd3, 0x04, 0x8c, 0x65, 0x84, 0x30, 0xba,
	0x27, 0xb4, 0x3e, 0x41, 0x97, 0xdd, 0xd1, 0xee, 0x71, 0xe0, 0x3f, 0xc4, 0xfb, 0x1f, 0x2b, 0x4e,
	0x2d, 0xd5, 0xdf, 0x1f, 0x9d, 0xd8, 0x5d, 0x62, 0xfc, 0x7f, 0xb9, 0xd6, 0xc8, 0x6a, 0xad, 0x91,
	0x97, 0xb5, 0x46, 0x1e, 0x36, 0x9a, 0xb2, 0xda, 0x68, 0xca, 0xd3, 0x46, 0x53, 0xae, 0x7e, 0xba,
	0x7e, 0xea, 0xcd, 0x67, 0xc2, 0x96, 0xa1, 0x71, 0xe8, 0xbe, 0x07, 0x46, 0x66, 0x54, 0x3b, 0xb9,
	0x88, 0x01, 0x67, 0xcd, 0x62, 0x8b, 0x7e, 0xbd, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xab, 0xc7,
	0x25, 0xab, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateRollapp(ctx context.Context, in *MsgCreateRollapp, opts ...grpc.CallOption) (*MsgCreateRollappResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateRollapp(ctx context.Context, in *MsgCreateRollapp, opts ...grpc.CallOption) (*MsgCreateRollappResponse, error) {
	out := new(MsgCreateRollappResponse)
	err := c.cc.Invoke(ctx, "/dymensionxyz.dymension.rollapp.Msg/CreateRollapp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateRollapp(context.Context, *MsgCreateRollapp) (*MsgCreateRollappResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateRollapp(ctx context.Context, req *MsgCreateRollapp) (*MsgCreateRollappResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRollapp not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateRollapp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateRollapp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateRollapp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dymensionxyz.dymension.rollapp.Msg/CreateRollapp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateRollapp(ctx, req.(*MsgCreateRollapp))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dymensionxyz.dymension.rollapp.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRollapp",
			Handler:    _Msg_CreateRollapp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rollapp/tx.proto",
}

func (m *MsgCreateRollapp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateRollapp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateRollapp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.PermissionedAddresses.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.MaxSequencers != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.MaxSequencers))
		i--
		dAtA[i] = 0x30
	}
	if m.MaxWithholdingBlocks != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.MaxWithholdingBlocks))
		i--
		dAtA[i] = 0x28
	}
	if len(m.GenesisPath) > 0 {
		i -= len(m.GenesisPath)
		copy(dAtA[i:], m.GenesisPath)
		i = encodeVarintTx(dAtA, i, uint64(len(m.GenesisPath)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CodeStamp) > 0 {
		i -= len(m.CodeStamp)
		copy(dAtA[i:], m.CodeStamp)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CodeStamp)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateRollappResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateRollappResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateRollappResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateRollapp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.CodeStamp)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.GenesisPath)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.MaxWithholdingBlocks != 0 {
		n += 1 + sovTx(uint64(m.MaxWithholdingBlocks))
	}
	if m.MaxSequencers != 0 {
		n += 1 + sovTx(uint64(m.MaxSequencers))
	}
	l = m.PermissionedAddresses.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgCreateRollappResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateRollapp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateRollapp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateRollapp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeStamp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodeStamp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenesisPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxWithholdingBlocks", wireType)
			}
			m.MaxWithholdingBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxWithholdingBlocks |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSequencers", wireType)
			}
			m.MaxSequencers = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSequencers |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PermissionedAddresses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PermissionedAddresses.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateRollappResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateRollappResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateRollappResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
