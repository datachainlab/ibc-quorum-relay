// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/lightclients/qbft/v1/qbft.proto

package module

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

type ClientState struct {
	TrustLevelNumerator   uint32 `protobuf:"varint,1,opt,name=trust_level_numerator,json=trustLevelNumerator,proto3" json:"trust_level_numerator,omitempty"`
	TrustLevelDenominator uint32 `protobuf:"varint,2,opt,name=trust_level_denominator,json=trustLevelDenominator,proto3" json:"trust_level_denominator,omitempty"`
	TrustingPeriod        uint64 `protobuf:"varint,3,opt,name=trusting_period,json=trustingPeriod,proto3" json:"trusting_period,omitempty"`
	ChainId               int32  `protobuf:"varint,4,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	LatestHeight          int32  `protobuf:"varint,5,opt,name=latest_height,json=latestHeight,proto3" json:"latest_height,omitempty"`
	Frozen                int32  `protobuf:"varint,6,opt,name=frozen,proto3" json:"frozen,omitempty"`
	IbcStoreAddress       []byte `protobuf:"bytes,7,opt,name=ibc_store_address,json=ibcStoreAddress,proto3" json:"ibc_store_address,omitempty"`
}

func (m *ClientState) Reset()         { *m = ClientState{} }
func (m *ClientState) String() string { return proto.CompactTextString(m) }
func (*ClientState) ProtoMessage()    {}
func (*ClientState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2e4ed46cb60dd4a, []int{0}
}
func (m *ClientState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientState.Merge(m, src)
}
func (m *ClientState) XXX_Size() int {
	return m.Size()
}
func (m *ClientState) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientState.DiscardUnknown(m)
}

var xxx_messageInfo_ClientState proto.InternalMessageInfo

type ConsensusState struct {
	Timestamp    uint64   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Root         []byte   `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
	ValidatorSet [][]byte `protobuf:"bytes,3,rep,name=validator_set,json=validatorSet,proto3" json:"validator_set,omitempty"`
}

func (m *ConsensusState) Reset()         { *m = ConsensusState{} }
func (m *ConsensusState) String() string { return proto.CompactTextString(m) }
func (*ConsensusState) ProtoMessage()    {}
func (*ConsensusState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2e4ed46cb60dd4a, []int{1}
}
func (m *ConsensusState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConsensusState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConsensusState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConsensusState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusState.Merge(m, src)
}
func (m *ConsensusState) XXX_Size() int {
	return m.Size()
}
func (m *ConsensusState) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusState.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusState proto.InternalMessageInfo

func (m *ConsensusState) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ConsensusState) GetRoot() []byte {
	if m != nil {
		return m.Root
	}
	return nil
}

func (m *ConsensusState) GetValidatorSet() [][]byte {
	if m != nil {
		return m.ValidatorSet
	}
	return nil
}

type Header struct {
	TrustedHeight     int32  `protobuf:"varint,1,opt,name=trusted_height,json=trustedHeight,proto3" json:"trusted_height,omitempty"`
	AccountProof      []byte `protobuf:"bytes,2,opt,name=account_proof,json=accountProof,proto3" json:"account_proof,omitempty"`
	GoQuorumHeaderRlp []byte `protobuf:"bytes,3,opt,name=go_quorum_header_rlp,json=goQuorumHeaderRlp,proto3" json:"go_quorum_header_rlp,omitempty"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2e4ed46cb60dd4a, []int{2}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Header.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return m.Size()
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetTrustedHeight() int32 {
	if m != nil {
		return m.TrustedHeight
	}
	return 0
}

func (m *Header) GetAccountProof() []byte {
	if m != nil {
		return m.AccountProof
	}
	return nil
}

func (m *Header) GetGoQuorumHeaderRlp() []byte {
	if m != nil {
		return m.GoQuorumHeaderRlp
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientState)(nil), "ibc.lightclients.qbft.v1.ClientState")
	proto.RegisterType((*ConsensusState)(nil), "ibc.lightclients.qbft.v1.ConsensusState")
	proto.RegisterType((*Header)(nil), "ibc.lightclients.qbft.v1.Header")
}

func init() {
	proto.RegisterFile("ibc/lightclients/qbft/v1/qbft.proto", fileDescriptor_b2e4ed46cb60dd4a)
}

var fileDescriptor_b2e4ed46cb60dd4a = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x9b, 0xad, 0xeb, 0xc0, 0xa4, 0x9b, 0x66, 0x36, 0x08, 0x08, 0x85, 0xaa, 0x15, 0xa2,
	0x42, 0x5a, 0xa3, 0x81, 0xc4, 0x81, 0x1b, 0x8c, 0xc3, 0x26, 0x21, 0x34, 0xd2, 0x1b, 0x17, 0xcb,
	0x89, 0x5f, 0x53, 0x4b, 0x8e, 0x9d, 0xd9, 0x4e, 0x25, 0xb8, 0x23, 0x71, 0xe4, 0x23, 0xf0, 0x01,
	0xf8, 0x20, 0x1c, 0x77, 0xe4, 0x88, 0xda, 0x2f, 0x82, 0xf2, 0xd2, 0xd2, 0x9d, 0x6a, 0xff, 0x7e,
	0xef, 0xf5, 0xe5, 0xfd, 0x13, 0x32, 0x92, 0x59, 0x9e, 0x28, 0x59, 0xcc, 0x7d, 0xae, 0x24, 0x68,
	0xef, 0x92, 0xeb, 0x6c, 0xe6, 0x93, 0xc5, 0x19, 0xfe, 0x4e, 0x2a, 0x6b, 0xbc, 0xa1, 0x91, 0xcc,
	0xf2, 0xc9, 0xed, 0xa2, 0x09, 0xca, 0xc5, 0xd9, 0xe3, 0xe3, 0xc2, 0x14, 0x06, 0x8b, 0x92, 0xe6,
	0xd4, 0xd6, 0x0f, 0x7f, 0xed, 0x90, 0x7b, 0xe7, 0x58, 0x39, 0xf5, 0xdc, 0x03, 0x7d, 0x49, 0x4e,
	0xbc, 0xad, 0x9d, 0x67, 0x0a, 0x16, 0xa0, 0x98, 0xae, 0x4b, 0xb0, 0xdc, 0x1b, 0x1b, 0x05, 0x83,
	0x60, 0xdc, 0x4f, 0xef, 0xa3, 0xfc, 0xd0, 0xb8, 0x8f, 0x1b, 0x45, 0x5f, 0x93, 0x87, 0xb7, 0x7b,
	0x04, 0x68, 0x53, 0x4a, 0x8d, 0x5d, 0x3b, 0xd8, 0x75, 0xb2, 0xed, 0x7a, 0xbf, 0x95, 0xf4, 0x39,
	0x39, 0x44, 0x21, 0x75, 0xc1, 0x2a, 0xb0, 0xd2, 0x88, 0x68, 0x77, 0x10, 0x8c, 0xbb, 0xe9, 0xc1,
	0x06, 0x5f, 0x21, 0xa5, 0x8f, 0xc8, 0x9d, 0x7c, 0xce, 0xa5, 0x66, 0x52, 0x44, 0xdd, 0x41, 0x30,
	0xde, 0x4b, 0xf7, 0xf1, 0x7e, 0x29, 0xe8, 0x88, 0xf4, 0x15, 0xf7, 0xe0, 0x3c, 0x9b, 0x43, 0xb3,
	0x75, 0xb4, 0x87, 0x3e, 0x6c, 0xe1, 0x05, 0x32, 0xfa, 0x80, 0xf4, 0x66, 0xd6, 0x7c, 0x05, 0x1d,
	0xf5, 0xd0, 0xae, 0x6f, 0xf4, 0x05, 0x39, 0x92, 0x59, 0xce, 0x9c, 0x37, 0x16, 0x18, 0x17, 0xc2,
	0x82, 0x73, 0xd1, 0xfe, 0x20, 0x18, 0x87, 0xe9, 0xa1, 0xcc, 0xf2, 0x69, 0xc3, 0xdf, 0xb6, 0xf8,
	0x4d, 0xf7, 0xfb, 0xcf, 0xa7, 0x9d, 0x61, 0x41, 0x0e, 0xce, 0x8d, 0x76, 0xa0, 0x5d, 0xed, 0xda,
	0xc0, 0x9e, 0x90, 0xbb, 0x5e, 0x96, 0xe0, 0x3c, 0x2f, 0x2b, 0x0c, 0xa9, 0x9b, 0x6e, 0x01, 0xa5,
	0xa4, 0x6b, 0x8d, 0xf1, 0x98, 0x43, 0x98, 0xe2, 0xb9, 0x79, 0xe4, 0x05, 0x57, 0x52, 0x34, 0x19,
	0x30, 0x07, 0x3e, 0xda, 0x1d, 0xec, 0x8e, 0xc3, 0x34, 0xfc, 0x0f, 0xa7, 0xe0, 0x87, 0xdf, 0x02,
	0xd2, 0xbb, 0x00, 0x2e, 0xc0, 0xd2, 0x67, 0xa4, 0xcd, 0x03, 0xc4, 0x66, 0xc7, 0x00, 0xb7, 0xe8,
	0xaf, 0xe9, 0x7a, 0xc9, 0x11, 0xe9, 0xf3, 0x3c, 0x37, 0xb5, 0xf6, 0xac, 0xb2, 0xc6, 0xcc, 0xd6,
	0x33, 0xc3, 0x35, 0xbc, 0x6a, 0x18, 0x4d, 0xc8, 0x71, 0x61, 0xd8, 0x75, 0x6d, 0x6c, 0x5d, 0xb2,
	0x39, 0xfe, 0x3f, 0xb3, 0xaa, 0xc2, 0xdc, 0xc3, 0xf4, 0xa8, 0x30, 0x9f, 0x50, 0xb5, 0x93, 0x53,
	0x55, 0xbd, 0xbb, 0xfc, 0xbd, 0x8c, 0x83, 0x9b, 0x65, 0x1c, 0xfc, 0x5d, 0xc6, 0xc1, 0x8f, 0x55,
	0xdc, 0xb9, 0x59, 0xc5, 0x9d, 0x3f, 0xab, 0xb8, 0xf3, 0x39, 0x29, 0xa4, 0x9f, 0xd7, 0xd9, 0x24,
	0x37, 0x65, 0x22, 0xb8, 0xe7, 0xf8, 0x46, 0x14, 0xcf, 0x12, 0x99, 0xe5, 0xa7, 0xed, 0x84, 0x53,
	0x0b, 0x8a, 0x7f, 0x49, 0x4a, 0x23, 0x6a, 0x05, 0x59, 0x0f, 0xbf, 0xb8, 0x57, 0xff, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x98, 0xea, 0x84, 0x52, 0xc8, 0x02, 0x00, 0x00,
}

func (m *ClientState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IbcStoreAddress) > 0 {
		i -= len(m.IbcStoreAddress)
		copy(dAtA[i:], m.IbcStoreAddress)
		i = encodeVarintQbft(dAtA, i, uint64(len(m.IbcStoreAddress)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Frozen != 0 {
		i = encodeVarintQbft(dAtA, i, uint64(m.Frozen))
		i--
		dAtA[i] = 0x30
	}
	if m.LatestHeight != 0 {
		i = encodeVarintQbft(dAtA, i, uint64(m.LatestHeight))
		i--
		dAtA[i] = 0x28
	}
	if m.ChainId != 0 {
		i = encodeVarintQbft(dAtA, i, uint64(m.ChainId))
		i--
		dAtA[i] = 0x20
	}
	if m.TrustingPeriod != 0 {
		i = encodeVarintQbft(dAtA, i, uint64(m.TrustingPeriod))
		i--
		dAtA[i] = 0x18
	}
	if m.TrustLevelDenominator != 0 {
		i = encodeVarintQbft(dAtA, i, uint64(m.TrustLevelDenominator))
		i--
		dAtA[i] = 0x10
	}
	if m.TrustLevelNumerator != 0 {
		i = encodeVarintQbft(dAtA, i, uint64(m.TrustLevelNumerator))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ConsensusState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConsensusState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConsensusState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorSet) > 0 {
		for iNdEx := len(m.ValidatorSet) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ValidatorSet[iNdEx])
			copy(dAtA[i:], m.ValidatorSet[iNdEx])
			i = encodeVarintQbft(dAtA, i, uint64(len(m.ValidatorSet[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Root) > 0 {
		i -= len(m.Root)
		copy(dAtA[i:], m.Root)
		i = encodeVarintQbft(dAtA, i, uint64(len(m.Root)))
		i--
		dAtA[i] = 0x12
	}
	if m.Timestamp != 0 {
		i = encodeVarintQbft(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Header) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Header) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Header) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GoQuorumHeaderRlp) > 0 {
		i -= len(m.GoQuorumHeaderRlp)
		copy(dAtA[i:], m.GoQuorumHeaderRlp)
		i = encodeVarintQbft(dAtA, i, uint64(len(m.GoQuorumHeaderRlp)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AccountProof) > 0 {
		i -= len(m.AccountProof)
		copy(dAtA[i:], m.AccountProof)
		i = encodeVarintQbft(dAtA, i, uint64(len(m.AccountProof)))
		i--
		dAtA[i] = 0x12
	}
	if m.TrustedHeight != 0 {
		i = encodeVarintQbft(dAtA, i, uint64(m.TrustedHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQbft(dAtA []byte, offset int, v uint64) int {
	offset -= sovQbft(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClientState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TrustLevelNumerator != 0 {
		n += 1 + sovQbft(uint64(m.TrustLevelNumerator))
	}
	if m.TrustLevelDenominator != 0 {
		n += 1 + sovQbft(uint64(m.TrustLevelDenominator))
	}
	if m.TrustingPeriod != 0 {
		n += 1 + sovQbft(uint64(m.TrustingPeriod))
	}
	if m.ChainId != 0 {
		n += 1 + sovQbft(uint64(m.ChainId))
	}
	if m.LatestHeight != 0 {
		n += 1 + sovQbft(uint64(m.LatestHeight))
	}
	if m.Frozen != 0 {
		n += 1 + sovQbft(uint64(m.Frozen))
	}
	l = len(m.IbcStoreAddress)
	if l > 0 {
		n += 1 + l + sovQbft(uint64(l))
	}
	return n
}

func (m *ConsensusState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != 0 {
		n += 1 + sovQbft(uint64(m.Timestamp))
	}
	l = len(m.Root)
	if l > 0 {
		n += 1 + l + sovQbft(uint64(l))
	}
	if len(m.ValidatorSet) > 0 {
		for _, b := range m.ValidatorSet {
			l = len(b)
			n += 1 + l + sovQbft(uint64(l))
		}
	}
	return n
}

func (m *Header) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TrustedHeight != 0 {
		n += 1 + sovQbft(uint64(m.TrustedHeight))
	}
	l = len(m.AccountProof)
	if l > 0 {
		n += 1 + l + sovQbft(uint64(l))
	}
	l = len(m.GoQuorumHeaderRlp)
	if l > 0 {
		n += 1 + l + sovQbft(uint64(l))
	}
	return n
}

func sovQbft(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQbft(x uint64) (n int) {
	return sovQbft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQbft
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
			return fmt.Errorf("proto: ClientState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustLevelNumerator", wireType)
			}
			m.TrustLevelNumerator = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQbft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TrustLevelNumerator |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustLevelDenominator", wireType)
			}
			m.TrustLevelDenominator = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQbft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TrustLevelDenominator |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustingPeriod", wireType)
			}
			m.TrustingPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQbft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TrustingPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			m.ChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQbft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestHeight", wireType)
			}
			m.LatestHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQbft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestHeight |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Frozen", wireType)
			}
			m.Frozen = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQbft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Frozen |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcStoreAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQbft
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
				return ErrInvalidLengthQbft
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQbft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcStoreAddress = append(m.IbcStoreAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.IbcStoreAddress == nil {
				m.IbcStoreAddress = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQbft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQbft
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
func (m *ConsensusState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQbft
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
			return fmt.Errorf("proto: ConsensusState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConsensusState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQbft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQbft
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
				return ErrInvalidLengthQbft
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQbft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Root = append(m.Root[:0], dAtA[iNdEx:postIndex]...)
			if m.Root == nil {
				m.Root = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorSet", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQbft
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
				return ErrInvalidLengthQbft
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQbft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorSet = append(m.ValidatorSet, make([]byte, postIndex-iNdEx))
			copy(m.ValidatorSet[len(m.ValidatorSet)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQbft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQbft
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
func (m *Header) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQbft
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
			return fmt.Errorf("proto: Header: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Header: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustedHeight", wireType)
			}
			m.TrustedHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQbft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TrustedHeight |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountProof", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQbft
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
				return ErrInvalidLengthQbft
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQbft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountProof = append(m.AccountProof[:0], dAtA[iNdEx:postIndex]...)
			if m.AccountProof == nil {
				m.AccountProof = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoQuorumHeaderRlp", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQbft
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
				return ErrInvalidLengthQbft
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQbft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GoQuorumHeaderRlp = append(m.GoQuorumHeaderRlp[:0], dAtA[iNdEx:postIndex]...)
			if m.GoQuorumHeaderRlp == nil {
				m.GoQuorumHeaderRlp = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQbft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQbft
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
func skipQbft(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQbft
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
					return 0, ErrIntOverflowQbft
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
					return 0, ErrIntOverflowQbft
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
				return 0, ErrInvalidLengthQbft
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQbft
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQbft
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQbft        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQbft          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQbft = fmt.Errorf("proto: unexpected end of group")
)
