// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lbm/fbridge/v1/fbridge.proto

package types

import (
	fmt "fmt"
	github_com_Finschia_finschia_sdk_types "github.com/Finschia/finschia-sdk/types"
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

// Role defines the role of the operator, guardian, and judge.
type Role int32

const (
	Role_UNSPECIFIED Role = 0
	Role_GUARDIAN    Role = 1
	Role_OPERATOR    Role = 2
	Role_JUDGE       Role = 3
)

var Role_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "GUARDIAN",
	2: "OPERATOR",
	3: "JUDGE",
}

var Role_value = map[string]int32{
	"UNSPECIFIED": 0,
	"GUARDIAN":    1,
	"OPERATOR":    2,
	"JUDGE":       3,
}

func (x Role) String() string {
	return proto.EnumName(Role_name, int32(x))
}

func (Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_62374d75fc6aa1ba, []int{0}
}

type Params struct {
	// ratio of how many operators' confirmations are needed to be valid.
	OperatorTrustLevel *Fraction `protobuf:"bytes,1,opt,name=operator_trust_level,json=operatorTrustLevel,proto3" json:"operator_trust_level,omitempty"`
	// ratio of how many guardians' confirmations are needed to be valid.
	GuardianTrustLevel *Fraction `protobuf:"bytes,2,opt,name=guardian_trust_level,json=guardianTrustLevel,proto3" json:"guardian_trust_level,omitempty"`
	// ratio of how many judges' confirmations are needed to be valid.
	JudgeTrustLevel *Fraction `protobuf:"bytes,3,opt,name=judge_trust_level,json=judgeTrustLevel,proto3" json:"judge_trust_level,omitempty"`
	// default timelock period for each provision (unix timestamp)
	TimelockPeriod uint64 `protobuf:"varint,4,opt,name=timelock_period,json=timelockPeriod,proto3" json:"timelock_period,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_62374d75fc6aa1ba, []int{0}
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

func (m *Params) GetOperatorTrustLevel() *Fraction {
	if m != nil {
		return m.OperatorTrustLevel
	}
	return nil
}

func (m *Params) GetGuardianTrustLevel() *Fraction {
	if m != nil {
		return m.GuardianTrustLevel
	}
	return nil
}

func (m *Params) GetJudgeTrustLevel() *Fraction {
	if m != nil {
		return m.JudgeTrustLevel
	}
	return nil
}

func (m *Params) GetTimelockPeriod() uint64 {
	if m != nil {
		return m.TimelockPeriod
	}
	return 0
}

// Provision is a struct that represents a provision internally.
type ProvisionData struct {
	// the sequence number of the bridge request
	Seq uint64 `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	// the amount of token to be claimed
	Amount github_com_Finschia_finschia_sdk_types.Int `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/Finschia/finschia-sdk/types.Int" json:"amount"`
	// the sender address on the source chain
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	// the recipient address on the destination chain
	Receiver string `protobuf:"bytes,4,opt,name=receiver,proto3" json:"receiver,omitempty"`
}

func (m *ProvisionData) Reset()         { *m = ProvisionData{} }
func (m *ProvisionData) String() string { return proto.CompactTextString(m) }
func (*ProvisionData) ProtoMessage()    {}
func (*ProvisionData) Descriptor() ([]byte, []int) {
	return fileDescriptor_62374d75fc6aa1ba, []int{1}
}
func (m *ProvisionData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProvisionData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProvisionData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProvisionData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProvisionData.Merge(m, src)
}
func (m *ProvisionData) XXX_Size() int {
	return m.Size()
}
func (m *ProvisionData) XXX_DiscardUnknown() {
	xxx_messageInfo_ProvisionData.DiscardUnknown(m)
}

var xxx_messageInfo_ProvisionData proto.InternalMessageInfo

func (m *ProvisionData) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *ProvisionData) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *ProvisionData) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

// ProvisionStatus is a struct that represents the status of a provision.
// To optimize computational cost, we have collected frequently changing values from provision.
type ProvisionStatus struct {
	// the unix timestamp the provision will be able to be claimed (unix timestamp)
	TimelockEnd uint64 `protobuf:"varint,1,opt,name=timelock_end,json=timelockEnd,proto3" json:"timelock_end,omitempty"`
	// a value that tells how many operators have submitted this provision
	ConfirmCounts int32 `protobuf:"varint,2,opt,name=confirm_counts,json=confirmCounts,proto3" json:"confirm_counts,omitempty"`
	// whether the provision has been claimed
	IsClaimed bool `protobuf:"varint,3,opt,name=is_claimed,json=isClaimed,proto3" json:"is_claimed,omitempty"`
}

func (m *ProvisionStatus) Reset()         { *m = ProvisionStatus{} }
func (m *ProvisionStatus) String() string { return proto.CompactTextString(m) }
func (*ProvisionStatus) ProtoMessage()    {}
func (*ProvisionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_62374d75fc6aa1ba, []int{2}
}
func (m *ProvisionStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProvisionStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProvisionStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProvisionStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProvisionStatus.Merge(m, src)
}
func (m *ProvisionStatus) XXX_Size() int {
	return m.Size()
}
func (m *ProvisionStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ProvisionStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ProvisionStatus proto.InternalMessageInfo

func (m *ProvisionStatus) GetTimelockEnd() uint64 {
	if m != nil {
		return m.TimelockEnd
	}
	return 0
}

func (m *ProvisionStatus) GetConfirmCounts() int32 {
	if m != nil {
		return m.ConfirmCounts
	}
	return 0
}

func (m *ProvisionStatus) GetIsClaimed() bool {
	if m != nil {
		return m.IsClaimed
	}
	return false
}

// Fraction defines the protobuf message type for tmmath.Fraction that only
// supports positive values.
type Fraction struct {
	Numerator   uint64 `protobuf:"varint,1,opt,name=numerator,proto3" json:"numerator,omitempty"`
	Denominator uint64 `protobuf:"varint,2,opt,name=denominator,proto3" json:"denominator,omitempty"`
}

func (m *Fraction) Reset()         { *m = Fraction{} }
func (m *Fraction) String() string { return proto.CompactTextString(m) }
func (*Fraction) ProtoMessage()    {}
func (*Fraction) Descriptor() ([]byte, []int) {
	return fileDescriptor_62374d75fc6aa1ba, []int{3}
}
func (m *Fraction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Fraction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fraction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fraction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fraction.Merge(m, src)
}
func (m *Fraction) XXX_Size() int {
	return m.Size()
}
func (m *Fraction) XXX_DiscardUnknown() {
	xxx_messageInfo_Fraction.DiscardUnknown(m)
}

var xxx_messageInfo_Fraction proto.InternalMessageInfo

func (m *Fraction) GetNumerator() uint64 {
	if m != nil {
		return m.Numerator
	}
	return 0
}

func (m *Fraction) GetDenominator() uint64 {
	if m != nil {
		return m.Denominator
	}
	return 0
}

func init() {
	proto.RegisterEnum("lbm.fbridge.v1.Role", Role_name, Role_value)
	proto.RegisterType((*Params)(nil), "lbm.fbridge.v1.Params")
	proto.RegisterType((*ProvisionData)(nil), "lbm.fbridge.v1.ProvisionData")
	proto.RegisterType((*ProvisionStatus)(nil), "lbm.fbridge.v1.ProvisionStatus")
	proto.RegisterType((*Fraction)(nil), "lbm.fbridge.v1.Fraction")
}

func init() { proto.RegisterFile("lbm/fbridge/v1/fbridge.proto", fileDescriptor_62374d75fc6aa1ba) }

var fileDescriptor_62374d75fc6aa1ba = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xdf, 0x6e, 0xda, 0x3c,
	0x18, 0xc6, 0x09, 0x50, 0x44, 0x4c, 0x0b, 0x7c, 0x56, 0xf5, 0x09, 0x55, 0x5d, 0xca, 0x90, 0xa6,
	0x55, 0x95, 0x96, 0xac, 0xdd, 0xf9, 0x24, 0xca, 0x9f, 0x0a, 0x34, 0xb5, 0xc8, 0x2d, 0x27, 0x3b,
	0x41, 0x26, 0x31, 0xa9, 0xd7, 0xc4, 0x66, 0xb6, 0x13, 0x6d, 0xbb, 0x89, 0xed, 0x16, 0x76, 0x37,
	0x3d, 0xec, 0xe1, 0xb4, 0x83, 0x6a, 0x82, 0x1b, 0x99, 0x62, 0x12, 0x06, 0x47, 0x9c, 0xbd, 0xef,
	0x93, 0x3c, 0x3f, 0x3f, 0xaf, 0x5f, 0x19, 0x1c, 0x07, 0xd3, 0xd0, 0x99, 0x4d, 0x05, 0xf5, 0x7c,
	0xe2, 0xc4, 0xe7, 0x59, 0x69, 0xcf, 0x05, 0x57, 0x1c, 0x56, 0x83, 0x69, 0x68, 0x67, 0x52, 0x7c,
	0x7e, 0x74, 0xe8, 0x73, 0x9f, 0xeb, 0x4f, 0x4e, 0x52, 0xad, 0xfe, 0x6a, 0x7d, 0xcf, 0x83, 0xd2,
	0x08, 0x0b, 0x1c, 0x4a, 0x38, 0x04, 0x87, 0x7c, 0x4e, 0x04, 0x56, 0x5c, 0x4c, 0x94, 0x88, 0xa4,
	0x9a, 0x04, 0x24, 0x26, 0x41, 0xc3, 0x68, 0x1a, 0xa7, 0x95, 0x8b, 0x86, 0xbd, 0xcd, 0xb3, 0xfb,
	0x02, 0xbb, 0x8a, 0x72, 0x86, 0x60, 0xe6, 0xba, 0x4b, 0x4c, 0x1f, 0x12, 0x4f, 0xc2, 0xf2, 0x23,
	0x2c, 0x3c, 0x8a, 0xd9, 0x16, 0x2b, 0xbf, 0x8b, 0x95, 0xb9, 0x36, 0x58, 0x5d, 0xf0, 0xdf, 0xa7,
	0xc8, 0xf3, 0xc9, 0x16, 0xa8, 0xb0, 0x03, 0x54, 0xd3, 0x96, 0x0d, 0xca, 0x6b, 0x50, 0x53, 0x34,
	0x24, 0x01, 0x77, 0x1f, 0x26, 0x73, 0x22, 0x28, 0xf7, 0x1a, 0xc5, 0xa6, 0x71, 0x5a, 0x44, 0xd5,
	0x4c, 0x1e, 0x69, 0xb5, 0xf5, 0xd3, 0x00, 0x07, 0x23, 0xc1, 0x63, 0x2a, 0x29, 0x67, 0x5d, 0xac,
	0x30, 0xac, 0x83, 0x82, 0x24, 0x9f, 0xf5, 0x3d, 0x14, 0x51, 0x52, 0xc2, 0x21, 0x28, 0xe1, 0x90,
	0x47, 0x4c, 0xe9, 0x81, 0xcc, 0xcb, 0x8b, 0xc7, 0xe7, 0x93, 0xdc, 0xef, 0xe7, 0x93, 0x33, 0x9f,
	0xaa, 0xfb, 0x68, 0x6a, 0xbb, 0x3c, 0x74, 0xfa, 0x94, 0x49, 0xf7, 0x9e, 0x62, 0x67, 0x96, 0x16,
	0x6f, 0xa4, 0xf7, 0xe0, 0xa8, 0xaf, 0x73, 0x22, 0xed, 0x01, 0x53, 0x28, 0x25, 0xc0, 0xff, 0x41,
	0x49, 0x12, 0xe6, 0x11, 0xa1, 0x67, 0x32, 0x51, 0xda, 0xc1, 0x23, 0x50, 0x16, 0xc4, 0x25, 0x34,
	0x26, 0x42, 0x27, 0x35, 0xd1, 0xba, 0x6f, 0x7d, 0x03, 0xb5, 0x75, 0xc4, 0x5b, 0x85, 0x55, 0x24,
	0xe1, 0x4b, 0xb0, 0xbf, 0x9e, 0x8f, 0x30, 0x2f, 0x4d, 0x5b, 0xc9, 0xb4, 0x1e, 0xf3, 0xe0, 0x2b,
	0x50, 0x75, 0x39, 0x9b, 0x51, 0x11, 0x4e, 0xdc, 0xe4, 0x68, 0xa9, 0xd3, 0xef, 0xa1, 0x83, 0x54,
	0xed, 0x68, 0x11, 0xbe, 0x00, 0x80, 0xca, 0x89, 0x1b, 0x60, 0x1a, 0x12, 0x4f, 0x87, 0x2a, 0x23,
	0x93, 0xca, 0xce, 0x4a, 0x68, 0x0d, 0x41, 0x39, 0xbb, 0x65, 0x78, 0x0c, 0x4c, 0x16, 0x85, 0xab,
	0xed, 0xa7, 0x27, 0xfe, 0x13, 0x60, 0x13, 0x54, 0x3c, 0xc2, 0x78, 0x48, 0x99, 0xfe, 0x9e, 0x5f,
	0x25, 0xda, 0x90, 0xce, 0xde, 0x83, 0x22, 0xe2, 0x01, 0x81, 0x35, 0x50, 0x19, 0x5f, 0xdf, 0x8e,
	0x7a, 0x9d, 0x41, 0x7f, 0xd0, 0xeb, 0xd6, 0x73, 0x70, 0x1f, 0x94, 0xaf, 0xc6, 0x6d, 0xd4, 0x1d,
	0xb4, 0xaf, 0xeb, 0x46, 0xd2, 0xdd, 0x8c, 0x7a, 0xa8, 0x7d, 0x77, 0x83, 0xea, 0x79, 0x68, 0x82,
	0xbd, 0xe1, 0xb8, 0x7b, 0xd5, 0xab, 0x17, 0x2e, 0x87, 0x8f, 0x0b, 0xcb, 0x78, 0x5a, 0x58, 0xc6,
	0x9f, 0x85, 0x65, 0xfc, 0x58, 0x5a, 0xb9, 0xa7, 0xa5, 0x95, 0xfb, 0xb5, 0xb4, 0x72, 0x1f, 0xdf,
	0xee, 0xdc, 0xc4, 0x97, 0xf5, 0xd3, 0xd1, 0x3b, 0x99, 0x96, 0xf4, 0x83, 0x78, 0xf7, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x47, 0x3c, 0x5e, 0x9b, 0x56, 0x03, 0x00, 0x00,
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
	if m.TimelockPeriod != 0 {
		i = encodeVarintFbridge(dAtA, i, uint64(m.TimelockPeriod))
		i--
		dAtA[i] = 0x20
	}
	if m.JudgeTrustLevel != nil {
		{
			size, err := m.JudgeTrustLevel.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFbridge(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.GuardianTrustLevel != nil {
		{
			size, err := m.GuardianTrustLevel.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFbridge(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.OperatorTrustLevel != nil {
		{
			size, err := m.OperatorTrustLevel.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFbridge(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProvisionData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProvisionData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProvisionData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintFbridge(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintFbridge(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintFbridge(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Seq != 0 {
		i = encodeVarintFbridge(dAtA, i, uint64(m.Seq))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ProvisionStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProvisionStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProvisionStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsClaimed {
		i--
		if m.IsClaimed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.ConfirmCounts != 0 {
		i = encodeVarintFbridge(dAtA, i, uint64(m.ConfirmCounts))
		i--
		dAtA[i] = 0x10
	}
	if m.TimelockEnd != 0 {
		i = encodeVarintFbridge(dAtA, i, uint64(m.TimelockEnd))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Fraction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fraction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Fraction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Denominator != 0 {
		i = encodeVarintFbridge(dAtA, i, uint64(m.Denominator))
		i--
		dAtA[i] = 0x10
	}
	if m.Numerator != 0 {
		i = encodeVarintFbridge(dAtA, i, uint64(m.Numerator))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintFbridge(dAtA []byte, offset int, v uint64) int {
	offset -= sovFbridge(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OperatorTrustLevel != nil {
		l = m.OperatorTrustLevel.Size()
		n += 1 + l + sovFbridge(uint64(l))
	}
	if m.GuardianTrustLevel != nil {
		l = m.GuardianTrustLevel.Size()
		n += 1 + l + sovFbridge(uint64(l))
	}
	if m.JudgeTrustLevel != nil {
		l = m.JudgeTrustLevel.Size()
		n += 1 + l + sovFbridge(uint64(l))
	}
	if m.TimelockPeriod != 0 {
		n += 1 + sovFbridge(uint64(m.TimelockPeriod))
	}
	return n
}

func (m *ProvisionData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Seq != 0 {
		n += 1 + sovFbridge(uint64(m.Seq))
	}
	l = m.Amount.Size()
	n += 1 + l + sovFbridge(uint64(l))
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovFbridge(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovFbridge(uint64(l))
	}
	return n
}

func (m *ProvisionStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TimelockEnd != 0 {
		n += 1 + sovFbridge(uint64(m.TimelockEnd))
	}
	if m.ConfirmCounts != 0 {
		n += 1 + sovFbridge(uint64(m.ConfirmCounts))
	}
	if m.IsClaimed {
		n += 2
	}
	return n
}

func (m *Fraction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Numerator != 0 {
		n += 1 + sovFbridge(uint64(m.Numerator))
	}
	if m.Denominator != 0 {
		n += 1 + sovFbridge(uint64(m.Denominator))
	}
	return n
}

func sovFbridge(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFbridge(x uint64) (n int) {
	return sovFbridge(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFbridge
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorTrustLevel", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFbridge
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
				return ErrInvalidLengthFbridge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFbridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OperatorTrustLevel == nil {
				m.OperatorTrustLevel = &Fraction{}
			}
			if err := m.OperatorTrustLevel.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GuardianTrustLevel", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFbridge
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
				return ErrInvalidLengthFbridge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFbridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GuardianTrustLevel == nil {
				m.GuardianTrustLevel = &Fraction{}
			}
			if err := m.GuardianTrustLevel.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JudgeTrustLevel", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFbridge
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
				return ErrInvalidLengthFbridge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFbridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.JudgeTrustLevel == nil {
				m.JudgeTrustLevel = &Fraction{}
			}
			if err := m.JudgeTrustLevel.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimelockPeriod", wireType)
			}
			m.TimelockPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFbridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimelockPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFbridge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFbridge
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
func (m *ProvisionData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFbridge
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
			return fmt.Errorf("proto: ProvisionData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProvisionData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seq", wireType)
			}
			m.Seq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFbridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFbridge
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
				return ErrInvalidLengthFbridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFbridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFbridge
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
				return ErrInvalidLengthFbridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFbridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFbridge
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
				return ErrInvalidLengthFbridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFbridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFbridge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFbridge
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
func (m *ProvisionStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFbridge
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
			return fmt.Errorf("proto: ProvisionStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProvisionStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimelockEnd", wireType)
			}
			m.TimelockEnd = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFbridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimelockEnd |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfirmCounts", wireType)
			}
			m.ConfirmCounts = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFbridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConfirmCounts |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsClaimed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFbridge
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
			m.IsClaimed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipFbridge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFbridge
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
func (m *Fraction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFbridge
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
			return fmt.Errorf("proto: Fraction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fraction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Numerator", wireType)
			}
			m.Numerator = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFbridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Numerator |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denominator", wireType)
			}
			m.Denominator = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFbridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Denominator |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFbridge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFbridge
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
func skipFbridge(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFbridge
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
					return 0, ErrIntOverflowFbridge
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
					return 0, ErrIntOverflowFbridge
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
				return 0, ErrInvalidLengthFbridge
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFbridge
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFbridge
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFbridge        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFbridge          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFbridge = fmt.Errorf("proto: unexpected end of group")
)
