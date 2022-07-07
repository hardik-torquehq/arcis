// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: arcis/feesplit/v1/genesis.proto

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

// GenesisState defines the module's genesis state.
type GenesisState struct {
	// module parameters
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// active registered contracts for fee distribution
	FeeSplits []FeeSplit `protobuf:"bytes,2,rep,name=fee_splits,json=feeSplits,proto3" json:"fee_splits"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_908469c9903cc008, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetFeeSplits() []FeeSplit {
	if m != nil {
		return m.FeeSplits
	}
	return nil
}

// Params defines the feesplit module params
type Params struct {
	// enable_fee_split defines a parameter to enable the feesplit module
	EnableFeeSplit bool `protobuf:"varint,1,opt,name=enable_fee_split,json=enableFeeSplit,proto3" json:"enable_fee_split,omitempty"`
	// developer_shares defines the proportion of the transaction fees to be
	// distributed to the registered contract owner
	DeveloperShares github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=developer_shares,json=developerShares,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"developer_shares"`
	// addr_derivation_cost_create defines the cost of address derivation for
	// verifying the contract deployer at fee registration
	AddrDerivationCostCreate uint64 `protobuf:"varint,3,opt,name=addr_derivation_cost_create,json=addrDerivationCostCreate,proto3" json:"addr_derivation_cost_create,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_908469c9903cc008, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetEnableFeeSplit() bool {
	if m != nil {
		return m.EnableFeeSplit
	}
	return false
}

func (m *Params) GetAddrDerivationCostCreate() uint64 {
	if m != nil {
		return m.AddrDerivationCostCreate
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "arcis.feesplit.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "arcis.feesplit.v1.Params")
}

func init() { proto.RegisterFile("arcis/feesplit/v1/genesis.proto", fileDescriptor_908469c9903cc008) }

var fileDescriptor_908469c9903cc008 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x6b, 0xea, 0x40,
	0x10, 0xc7, 0xb3, 0x4f, 0x91, 0xe7, 0xfa, 0x78, 0xcf, 0x17, 0x7a, 0x48, 0x15, 0x62, 0xf0, 0x50,
	0x42, 0xa1, 0x1b, 0xb4, 0xd0, 0x9e, 0x0a, 0x45, 0xa5, 0xbd, 0x96, 0x78, 0x6a, 0x2f, 0x61, 0x4d,
	0xc6, 0x18, 0xaa, 0x6e, 0xd8, 0xdd, 0x86, 0xf6, 0x23, 0xf4, 0xd6, 0x8f, 0xe5, 0xa1, 0x07, 0x8f,
	0xa5, 0x07, 0x29, 0xfa, 0x45, 0x4a, 0x76, 0x35, 0x2d, 0xd8, 0xcb, 0xee, 0xec, 0xcc, 0xff, 0xff,
	0x9b, 0x61, 0x07, 0xb7, 0x20, 0x9b, 0x31, 0xe1, 0x8d, 0x01, 0x44, 0x3a, 0x4d, 0xa4, 0x97, 0x75,
	0xbc, 0x18, 0xe6, 0x20, 0x12, 0x41, 0x52, 0xce, 0x24, 0x33, 0xff, 0x2b, 0x01, 0xd9, 0x09, 0x48,
	0xd6, 0x69, 0x38, 0xfb, 0x9e, 0xa2, 0xac, 0x4c, 0x8d, 0x83, 0x98, 0xc5, 0x4c, 0x85, 0x5e, 0x1e,
	0xe9, 0x6c, 0xfb, 0x19, 0xe1, 0x3f, 0xd7, 0x1a, 0x3e, 0x94, 0x54, 0x82, 0x79, 0x8e, 0x2b, 0x29,
	0xe5, 0x74, 0x26, 0x2c, 0xe4, 0x20, 0xb7, 0xd6, 0x3d, 0x24, 0x7b, 0xcd, 0xc8, 0x8d, 0x12, 0xf4,
	0xca, 0x8b, 0x55, 0xcb, 0xf0, 0xb7, 0x72, 0xf3, 0x12, 0xe3, 0x31, 0x40, 0xa0, 0x44, 0xc2, 0xfa,
	0xe5, 0x94, 0xdc, 0x5a, 0xb7, 0xf9, 0x83, 0xf9, 0x0a, 0x60, 0x98, 0xc7, 0x5b, 0x7b, 0x75, 0xbc,
	0x7d, 0x8b, 0xf6, 0x2b, 0xc2, 0x15, 0x8d, 0x36, 0x5d, 0x5c, 0x87, 0x39, 0x1d, 0x4d, 0x21, 0x28,
	0x98, 0x6a, 0x9e, 0xdf, 0xfe, 0x5f, 0x9d, 0xdf, 0x51, 0xcc, 0x5b, 0x5c, 0x8f, 0x20, 0x83, 0x29,
	0x4b, 0x81, 0x07, 0x62, 0x42, 0x39, 0xe4, 0xcd, 0x91, 0x5b, 0xed, 0x91, 0x9c, 0xff, 0xbe, 0x6a,
	0x1d, 0xc5, 0x89, 0x9c, 0x3c, 0x8c, 0x48, 0xc8, 0x66, 0x5e, 0xc8, 0x44, 0xfe, 0x4d, 0xfa, 0x3a,
	0x11, 0xd1, 0xbd, 0x27, 0x9f, 0x52, 0x10, 0x64, 0x00, 0xa1, 0xff, 0xaf, 0xe0, 0x0c, 0x15, 0xc6,
	0xbc, 0xc0, 0x4d, 0x1a, 0x45, 0x3c, 0x88, 0x80, 0x27, 0x19, 0x95, 0x09, 0x9b, 0x07, 0x21, 0x13,
	0x32, 0x08, 0x39, 0x50, 0x09, 0x56, 0xc9, 0x41, 0x6e, 0xd9, 0xb7, 0x72, 0xc9, 0xa0, 0x50, 0xf4,
	0x99, 0x90, 0x7d, 0x55, 0xef, 0x0d, 0x16, 0x6b, 0x1b, 0x2d, 0xd7, 0x36, 0xfa, 0x58, 0xdb, 0xe8,
	0x65, 0x63, 0x1b, 0xcb, 0x8d, 0x6d, 0xbc, 0x6d, 0x6c, 0xe3, 0xee, 0xf8, 0xdb, 0x44, 0x7a, 0x6f,
	0xfa, 0xcc, 0xce, 0xbc, 0xc7, 0xaf, 0x15, 0xaa, 0xc9, 0x46, 0x15, 0xb5, 0xa7, 0xd3, 0xcf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x28, 0x26, 0xd8, 0x74, 0x15, 0x02, 0x00, 0x00,
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
	if len(m.FeeSplits) > 0 {
		for iNdEx := len(m.FeeSplits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeSplits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AddrDerivationCostCreate != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AddrDerivationCostCreate))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.DeveloperShares.Size()
		i -= size
		if _, err := m.DeveloperShares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.EnableFeeSplit {
		i--
		if m.EnableFeeSplit {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
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
	if len(m.FeeSplits) > 0 {
		for _, e := range m.FeeSplits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EnableFeeSplit {
		n += 2
	}
	l = m.DeveloperShares.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.AddrDerivationCostCreate != 0 {
		n += 1 + sovGenesis(uint64(m.AddrDerivationCostCreate))
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
				return fmt.Errorf("proto: wrong wireType = %d for field FeeSplits", wireType)
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
			m.FeeSplits = append(m.FeeSplits, FeeSplit{})
			if err := m.FeeSplits[len(m.FeeSplits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *Params) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableFeeSplit", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableFeeSplit = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeveloperShares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DeveloperShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddrDerivationCostCreate", wireType)
			}
			m.AddrDerivationCostCreate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AddrDerivationCostCreate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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