// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetacore/node_account.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	common "github.com/zeta-chain/zetacore/common"
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

type NodeStatus int32

const (
	NodeStatus_Unknown     NodeStatus = 0
	NodeStatus_Whitelisted NodeStatus = 1
	NodeStatus_Standby     NodeStatus = 2
	NodeStatus_Ready       NodeStatus = 3
	NodeStatus_Active      NodeStatus = 4
	NodeStatus_Disabled    NodeStatus = 5
)

var NodeStatus_name = map[int32]string{
	0: "Unknown",
	1: "Whitelisted",
	2: "Standby",
	3: "Ready",
	4: "Active",
	5: "Disabled",
}

var NodeStatus_value = map[string]int32{
	"Unknown":     0,
	"Whitelisted": 1,
	"Standby":     2,
	"Ready":       3,
	"Active":      4,
	"Disabled":    5,
}

func (x NodeStatus) String() string {
	return proto.EnumName(NodeStatus_name, int32(x))
}

func (NodeStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_62f33b2230c4851a, []int{0}
}

type NodeAccount struct {
	Creator     string                                        `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index       string                                        `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	NodeAddress github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,3,opt,name=nodeAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"nodeAddress,omitempty"`
	PubkeySet   *common.PubKeySet                             `protobuf:"bytes,4,opt,name=pubkeySet,proto3" json:"pubkeySet,omitempty"`
	NodeStatus  NodeStatus                                    `protobuf:"varint,5,opt,name=nodeStatus,proto3,enum=zetachain.zetacore.zetacore.NodeStatus" json:"nodeStatus,omitempty"`
}

func (m *NodeAccount) Reset()         { *m = NodeAccount{} }
func (m *NodeAccount) String() string { return proto.CompactTextString(m) }
func (*NodeAccount) ProtoMessage()    {}
func (*NodeAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_62f33b2230c4851a, []int{0}
}
func (m *NodeAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodeAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NodeAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NodeAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeAccount.Merge(m, src)
}
func (m *NodeAccount) XXX_Size() int {
	return m.Size()
}
func (m *NodeAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeAccount.DiscardUnknown(m)
}

var xxx_messageInfo_NodeAccount proto.InternalMessageInfo

func (m *NodeAccount) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *NodeAccount) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *NodeAccount) GetNodeAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.NodeAddress
	}
	return nil
}

func (m *NodeAccount) GetPubkeySet() *common.PubKeySet {
	if m != nil {
		return m.PubkeySet
	}
	return nil
}

func (m *NodeAccount) GetNodeStatus() NodeStatus {
	if m != nil {
		return m.NodeStatus
	}
	return NodeStatus_Unknown
}

func init() {
	proto.RegisterEnum("zetachain.zetacore.zetacore.NodeStatus", NodeStatus_name, NodeStatus_value)
	proto.RegisterType((*NodeAccount)(nil), "zetachain.zetacore.zetacore.NodeAccount")
}

func init() { proto.RegisterFile("zetacore/node_account.proto", fileDescriptor_62f33b2230c4851a) }

var fileDescriptor_62f33b2230c4851a = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0xaa, 0xd3, 0x40,
	0x18, 0xcd, 0xf4, 0x36, 0xf7, 0xda, 0x2f, 0x17, 0x8d, 0x63, 0x17, 0xa1, 0x42, 0x0c, 0x6e, 0x0c,
	0x42, 0x33, 0x58, 0x9f, 0x20, 0x45, 0x10, 0x11, 0x8a, 0x24, 0x88, 0xe0, 0x46, 0x26, 0x33, 0x43,
	0x1b, 0xda, 0xcc, 0x94, 0xcc, 0x44, 0x1b, 0x9f, 0xc1, 0x85, 0x0f, 0xe1, 0xc2, 0x47, 0x71, 0xd9,
	0xa5, 0x2b, 0x91, 0xf6, 0x2d, 0x5c, 0x49, 0x92, 0xfe, 0xad, 0xee, 0xea, 0xfb, 0x39, 0xe7, 0x7c,
	0x1c, 0xce, 0x0c, 0x3c, 0xfe, 0x2a, 0x0c, 0x65, 0xaa, 0x14, 0x44, 0x2a, 0x2e, 0x3e, 0x51, 0xc6,
	0x54, 0x25, 0x4d, 0xb4, 0x2e, 0x95, 0x51, 0xb8, 0x03, 0x17, 0x34, 0x97, 0xd1, 0x91, 0x76, 0x6a,
	0x46, 0xc3, 0xb9, 0x9a, 0xab, 0x96, 0x47, 0x9a, 0xae, 0x93, 0x8c, 0x1e, 0x31, 0x55, 0x14, 0x4a,
	0x92, 0xae, 0x74, 0xcb, 0xa7, 0xdf, 0x7a, 0xe0, 0xcc, 0x14, 0x17, 0x71, 0x77, 0x1d, 0x7b, 0x70,
	0xc3, 0x4a, 0x41, 0x8d, 0x2a, 0x3d, 0x14, 0xa0, 0x70, 0x90, 0x1c, 0x47, 0x3c, 0x04, 0x3b, 0x97,
	0x5c, 0x6c, 0xbc, 0x5e, 0xbb, 0xef, 0x06, 0x9c, 0x82, 0xd3, 0xb8, 0x8b, 0x39, 0x2f, 0x85, 0xd6,
	0xde, 0x55, 0x80, 0xc2, 0xdb, 0xe9, 0x8b, 0x7f, 0x7f, 0x9e, 0x8c, 0xe7, 0xb9, 0x59, 0x54, 0x59,
	0xc4, 0x54, 0x41, 0x98, 0xd2, 0x85, 0xd2, 0x87, 0x32, 0xd6, 0x7c, 0x49, 0x4c, 0xbd, 0x16, 0x3a,
	0x8a, 0x19, 0x3b, 0x08, 0x93, 0xcb, 0x2b, 0x98, 0xc0, 0x60, 0x5d, 0x65, 0x4b, 0x51, 0xa7, 0xc2,
	0x78, 0xfd, 0x00, 0x85, 0xce, 0xe4, 0x61, 0x74, 0xb0, 0xfd, 0xae, 0xca, 0xde, 0xb6, 0x40, 0x72,
	0xe6, 0xe0, 0xd7, 0x00, 0x8d, 0x3e, 0x35, 0xd4, 0x54, 0xda, 0xb3, 0x03, 0x14, 0xde, 0x9f, 0x3c,
	0x8b, 0xee, 0x88, 0x28, 0x9a, 0x9d, 0xe8, 0xc9, 0x85, 0xf4, 0x79, 0x06, 0x70, 0x46, 0xb0, 0x03,
	0x37, 0xef, 0xe5, 0x52, 0xaa, 0x2f, 0xd2, 0xb5, 0xf0, 0x03, 0x70, 0x3e, 0x2c, 0x72, 0x23, 0x56,
	0xb9, 0x36, 0x82, 0xbb, 0xa8, 0x41, 0x53, 0x43, 0x25, 0xcf, 0x6a, 0xb7, 0x87, 0x07, 0x60, 0x27,
	0x82, 0xf2, 0xda, 0xbd, 0xc2, 0x00, 0xd7, 0x31, 0x33, 0xf9, 0x67, 0xe1, 0xf6, 0xf1, 0x2d, 0xdc,
	0x7b, 0x95, 0x6b, 0x9a, 0xad, 0x04, 0x77, 0xed, 0x51, 0xff, 0xe7, 0x0f, 0x1f, 0x4d, 0xdf, 0xfc,
	0xda, 0xf9, 0x68, 0xbb, 0xf3, 0xd1, 0xdf, 0x9d, 0x8f, 0xbe, 0xef, 0x7d, 0x6b, 0xbb, 0xf7, 0xad,
	0xdf, 0x7b, 0xdf, 0xfa, 0x48, 0x2e, 0x32, 0x6b, 0x9c, 0x8e, 0x5b, 0xf7, 0xe4, 0xf4, 0x0f, 0x36,
	0xe7, 0xb6, 0x0d, 0x30, 0xbb, 0x6e, 0x1f, 0xf1, 0xe5, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x12,
	0xf8, 0x61, 0xb1, 0x2b, 0x02, 0x00, 0x00,
}

func (m *NodeAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NodeAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NodeStatus != 0 {
		i = encodeVarintNodeAccount(dAtA, i, uint64(m.NodeStatus))
		i--
		dAtA[i] = 0x28
	}
	if m.PubkeySet != nil {
		{
			size, err := m.PubkeySet.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNodeAccount(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.NodeAddress) > 0 {
		i -= len(m.NodeAddress)
		copy(dAtA[i:], m.NodeAddress)
		i = encodeVarintNodeAccount(dAtA, i, uint64(len(m.NodeAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintNodeAccount(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintNodeAccount(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNodeAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovNodeAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NodeAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovNodeAccount(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovNodeAccount(uint64(l))
	}
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovNodeAccount(uint64(l))
	}
	if m.PubkeySet != nil {
		l = m.PubkeySet.Size()
		n += 1 + l + sovNodeAccount(uint64(l))
	}
	if m.NodeStatus != 0 {
		n += 1 + sovNodeAccount(uint64(m.NodeStatus))
	}
	return n
}

func sovNodeAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNodeAccount(x uint64) (n int) {
	return sovNodeAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NodeAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNodeAccount
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
			return fmt.Errorf("proto: NodeAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeAccount
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
				return ErrInvalidLengthNodeAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNodeAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeAccount
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
				return ErrInvalidLengthNodeAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNodeAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeAccount
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
				return ErrInvalidLengthNodeAccount
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNodeAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAddress = append(m.NodeAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.NodeAddress == nil {
				m.NodeAddress = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubkeySet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeAccount
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
				return ErrInvalidLengthNodeAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNodeAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PubkeySet == nil {
				m.PubkeySet = &common.PubKeySet{}
			}
			if err := m.PubkeySet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeStatus", wireType)
			}
			m.NodeStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeStatus |= NodeStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNodeAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNodeAccount
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
func skipNodeAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNodeAccount
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
					return 0, ErrIntOverflowNodeAccount
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
					return 0, ErrIntOverflowNodeAccount
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
				return 0, ErrInvalidLengthNodeAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNodeAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNodeAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNodeAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNodeAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNodeAccount = fmt.Errorf("proto: unexpected end of group")
)