// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/common/fault/v3/fault.proto

package envoy_extensions_filters_common_fault_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FaultDelay_FaultDelayType int32

const (
	FaultDelay_FIXED FaultDelay_FaultDelayType = 0
)

var FaultDelay_FaultDelayType_name = map[int32]string{
	0: "FIXED",
}

var FaultDelay_FaultDelayType_value = map[string]int32{
	"FIXED": 0,
}

func (x FaultDelay_FaultDelayType) String() string {
	return proto.EnumName(FaultDelay_FaultDelayType_name, int32(x))
}

func (FaultDelay_FaultDelayType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_99c2ca93a0e6a448, []int{0, 0}
}

type FaultDelay struct {
	HiddenEnvoyDeprecatedType FaultDelay_FaultDelayType `protobuf:"varint,1,opt,name=hidden_envoy_deprecated_type,json=hiddenEnvoyDeprecatedType,proto3,enum=envoy.extensions.filters.common.fault.v3.FaultDelay_FaultDelayType" json:"hidden_envoy_deprecated_type,omitempty"` // Deprecated: Do not use.
	// Types that are valid to be assigned to FaultDelaySecifier:
	//	*FaultDelay_FixedDelay
	//	*FaultDelay_HeaderDelay_
	FaultDelaySecifier   isFaultDelay_FaultDelaySecifier `protobuf_oneof:"fault_delay_secifier"`
	Percentage           *v3.FractionalPercent           `protobuf:"bytes,4,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *FaultDelay) Reset()         { *m = FaultDelay{} }
func (m *FaultDelay) String() string { return proto.CompactTextString(m) }
func (*FaultDelay) ProtoMessage()    {}
func (*FaultDelay) Descriptor() ([]byte, []int) {
	return fileDescriptor_99c2ca93a0e6a448, []int{0}
}

func (m *FaultDelay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultDelay.Unmarshal(m, b)
}
func (m *FaultDelay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultDelay.Marshal(b, m, deterministic)
}
func (m *FaultDelay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultDelay.Merge(m, src)
}
func (m *FaultDelay) XXX_Size() int {
	return xxx_messageInfo_FaultDelay.Size(m)
}
func (m *FaultDelay) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultDelay.DiscardUnknown(m)
}

var xxx_messageInfo_FaultDelay proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *FaultDelay) GetHiddenEnvoyDeprecatedType() FaultDelay_FaultDelayType {
	if m != nil {
		return m.HiddenEnvoyDeprecatedType
	}
	return FaultDelay_FIXED
}

type isFaultDelay_FaultDelaySecifier interface {
	isFaultDelay_FaultDelaySecifier()
}

type FaultDelay_FixedDelay struct {
	FixedDelay *duration.Duration `protobuf:"bytes,3,opt,name=fixed_delay,json=fixedDelay,proto3,oneof"`
}

type FaultDelay_HeaderDelay_ struct {
	HeaderDelay *FaultDelay_HeaderDelay `protobuf:"bytes,5,opt,name=header_delay,json=headerDelay,proto3,oneof"`
}

func (*FaultDelay_FixedDelay) isFaultDelay_FaultDelaySecifier() {}

func (*FaultDelay_HeaderDelay_) isFaultDelay_FaultDelaySecifier() {}

func (m *FaultDelay) GetFaultDelaySecifier() isFaultDelay_FaultDelaySecifier {
	if m != nil {
		return m.FaultDelaySecifier
	}
	return nil
}

func (m *FaultDelay) GetFixedDelay() *duration.Duration {
	if x, ok := m.GetFaultDelaySecifier().(*FaultDelay_FixedDelay); ok {
		return x.FixedDelay
	}
	return nil
}

func (m *FaultDelay) GetHeaderDelay() *FaultDelay_HeaderDelay {
	if x, ok := m.GetFaultDelaySecifier().(*FaultDelay_HeaderDelay_); ok {
		return x.HeaderDelay
	}
	return nil
}

func (m *FaultDelay) GetPercentage() *v3.FractionalPercent {
	if m != nil {
		return m.Percentage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FaultDelay) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FaultDelay_FixedDelay)(nil),
		(*FaultDelay_HeaderDelay_)(nil),
	}
}

type FaultDelay_HeaderDelay struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FaultDelay_HeaderDelay) Reset()         { *m = FaultDelay_HeaderDelay{} }
func (m *FaultDelay_HeaderDelay) String() string { return proto.CompactTextString(m) }
func (*FaultDelay_HeaderDelay) ProtoMessage()    {}
func (*FaultDelay_HeaderDelay) Descriptor() ([]byte, []int) {
	return fileDescriptor_99c2ca93a0e6a448, []int{0, 0}
}

func (m *FaultDelay_HeaderDelay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultDelay_HeaderDelay.Unmarshal(m, b)
}
func (m *FaultDelay_HeaderDelay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultDelay_HeaderDelay.Marshal(b, m, deterministic)
}
func (m *FaultDelay_HeaderDelay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultDelay_HeaderDelay.Merge(m, src)
}
func (m *FaultDelay_HeaderDelay) XXX_Size() int {
	return xxx_messageInfo_FaultDelay_HeaderDelay.Size(m)
}
func (m *FaultDelay_HeaderDelay) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultDelay_HeaderDelay.DiscardUnknown(m)
}

var xxx_messageInfo_FaultDelay_HeaderDelay proto.InternalMessageInfo

type FaultRateLimit struct {
	// Types that are valid to be assigned to LimitType:
	//	*FaultRateLimit_FixedLimit_
	//	*FaultRateLimit_HeaderLimit_
	LimitType            isFaultRateLimit_LimitType `protobuf_oneof:"limit_type"`
	Percentage           *v3.FractionalPercent      `protobuf:"bytes,2,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *FaultRateLimit) Reset()         { *m = FaultRateLimit{} }
func (m *FaultRateLimit) String() string { return proto.CompactTextString(m) }
func (*FaultRateLimit) ProtoMessage()    {}
func (*FaultRateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_99c2ca93a0e6a448, []int{1}
}

func (m *FaultRateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultRateLimit.Unmarshal(m, b)
}
func (m *FaultRateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultRateLimit.Marshal(b, m, deterministic)
}
func (m *FaultRateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultRateLimit.Merge(m, src)
}
func (m *FaultRateLimit) XXX_Size() int {
	return xxx_messageInfo_FaultRateLimit.Size(m)
}
func (m *FaultRateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultRateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_FaultRateLimit proto.InternalMessageInfo

type isFaultRateLimit_LimitType interface {
	isFaultRateLimit_LimitType()
}

type FaultRateLimit_FixedLimit_ struct {
	FixedLimit *FaultRateLimit_FixedLimit `protobuf:"bytes,1,opt,name=fixed_limit,json=fixedLimit,proto3,oneof"`
}

type FaultRateLimit_HeaderLimit_ struct {
	HeaderLimit *FaultRateLimit_HeaderLimit `protobuf:"bytes,3,opt,name=header_limit,json=headerLimit,proto3,oneof"`
}

func (*FaultRateLimit_FixedLimit_) isFaultRateLimit_LimitType() {}

func (*FaultRateLimit_HeaderLimit_) isFaultRateLimit_LimitType() {}

func (m *FaultRateLimit) GetLimitType() isFaultRateLimit_LimitType {
	if m != nil {
		return m.LimitType
	}
	return nil
}

func (m *FaultRateLimit) GetFixedLimit() *FaultRateLimit_FixedLimit {
	if x, ok := m.GetLimitType().(*FaultRateLimit_FixedLimit_); ok {
		return x.FixedLimit
	}
	return nil
}

func (m *FaultRateLimit) GetHeaderLimit() *FaultRateLimit_HeaderLimit {
	if x, ok := m.GetLimitType().(*FaultRateLimit_HeaderLimit_); ok {
		return x.HeaderLimit
	}
	return nil
}

func (m *FaultRateLimit) GetPercentage() *v3.FractionalPercent {
	if m != nil {
		return m.Percentage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FaultRateLimit) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FaultRateLimit_FixedLimit_)(nil),
		(*FaultRateLimit_HeaderLimit_)(nil),
	}
}

type FaultRateLimit_FixedLimit struct {
	LimitKbps            uint64   `protobuf:"varint,1,opt,name=limit_kbps,json=limitKbps,proto3" json:"limit_kbps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FaultRateLimit_FixedLimit) Reset()         { *m = FaultRateLimit_FixedLimit{} }
func (m *FaultRateLimit_FixedLimit) String() string { return proto.CompactTextString(m) }
func (*FaultRateLimit_FixedLimit) ProtoMessage()    {}
func (*FaultRateLimit_FixedLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_99c2ca93a0e6a448, []int{1, 0}
}

func (m *FaultRateLimit_FixedLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultRateLimit_FixedLimit.Unmarshal(m, b)
}
func (m *FaultRateLimit_FixedLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultRateLimit_FixedLimit.Marshal(b, m, deterministic)
}
func (m *FaultRateLimit_FixedLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultRateLimit_FixedLimit.Merge(m, src)
}
func (m *FaultRateLimit_FixedLimit) XXX_Size() int {
	return xxx_messageInfo_FaultRateLimit_FixedLimit.Size(m)
}
func (m *FaultRateLimit_FixedLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultRateLimit_FixedLimit.DiscardUnknown(m)
}

var xxx_messageInfo_FaultRateLimit_FixedLimit proto.InternalMessageInfo

func (m *FaultRateLimit_FixedLimit) GetLimitKbps() uint64 {
	if m != nil {
		return m.LimitKbps
	}
	return 0
}

type FaultRateLimit_HeaderLimit struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FaultRateLimit_HeaderLimit) Reset()         { *m = FaultRateLimit_HeaderLimit{} }
func (m *FaultRateLimit_HeaderLimit) String() string { return proto.CompactTextString(m) }
func (*FaultRateLimit_HeaderLimit) ProtoMessage()    {}
func (*FaultRateLimit_HeaderLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_99c2ca93a0e6a448, []int{1, 1}
}

func (m *FaultRateLimit_HeaderLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultRateLimit_HeaderLimit.Unmarshal(m, b)
}
func (m *FaultRateLimit_HeaderLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultRateLimit_HeaderLimit.Marshal(b, m, deterministic)
}
func (m *FaultRateLimit_HeaderLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultRateLimit_HeaderLimit.Merge(m, src)
}
func (m *FaultRateLimit_HeaderLimit) XXX_Size() int {
	return xxx_messageInfo_FaultRateLimit_HeaderLimit.Size(m)
}
func (m *FaultRateLimit_HeaderLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultRateLimit_HeaderLimit.DiscardUnknown(m)
}

var xxx_messageInfo_FaultRateLimit_HeaderLimit proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("envoy.extensions.filters.common.fault.v3.FaultDelay_FaultDelayType", FaultDelay_FaultDelayType_name, FaultDelay_FaultDelayType_value)
	proto.RegisterType((*FaultDelay)(nil), "envoy.extensions.filters.common.fault.v3.FaultDelay")
	proto.RegisterType((*FaultDelay_HeaderDelay)(nil), "envoy.extensions.filters.common.fault.v3.FaultDelay.HeaderDelay")
	proto.RegisterType((*FaultRateLimit)(nil), "envoy.extensions.filters.common.fault.v3.FaultRateLimit")
	proto.RegisterType((*FaultRateLimit_FixedLimit)(nil), "envoy.extensions.filters.common.fault.v3.FaultRateLimit.FixedLimit")
	proto.RegisterType((*FaultRateLimit_HeaderLimit)(nil), "envoy.extensions.filters.common.fault.v3.FaultRateLimit.HeaderLimit")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/common/fault/v3/fault.proto", fileDescriptor_99c2ca93a0e6a448)
}

var fileDescriptor_99c2ca93a0e6a448 = []byte{
	// 637 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0x5d, 0xa7, 0xb4, 0x1b, 0x54, 0x15, 0x0b, 0x89, 0x34, 0x85, 0xaa, 0x14, 0x09, 0x02,
	0x88, 0xb5, 0x94, 0x40, 0x81, 0x48, 0xa0, 0xca, 0xa4, 0x51, 0xcb, 0x97, 0x2a, 0x8b, 0x03, 0x37,
	0x6b, 0x63, 0x8f, 0xd3, 0x15, 0xae, 0xd7, 0xb2, 0x37, 0x51, 0x73, 0xe3, 0xc8, 0x89, 0x1f, 0xc0,
	0x4f, 0xe0, 0xc4, 0x11, 0x71, 0x47, 0x82, 0x23, 0xfc, 0x10, 0x0e, 0x9c, 0x50, 0x4f, 0x68, 0x3f,
	0x9c, 0xa4, 0x02, 0x44, 0x92, 0xdb, 0xae, 0x77, 0xde, 0x7b, 0x33, 0xef, 0x8d, 0x8c, 0x6e, 0x43,
	0x32, 0x60, 0x43, 0x07, 0x8e, 0x39, 0x24, 0x39, 0x65, 0x49, 0xee, 0x44, 0x34, 0xe6, 0x90, 0xe5,
	0x4e, 0xc0, 0x8e, 0x8e, 0x58, 0xe2, 0x44, 0xa4, 0x1f, 0x73, 0x67, 0xd0, 0x54, 0x07, 0x9c, 0x66,
	0x8c, 0x33, 0xbb, 0x2e, 0x51, 0x78, 0x8c, 0xc2, 0x1a, 0x85, 0x15, 0x0a, 0xab, 0xe2, 0x41, 0xb3,
	0xb6, 0xae, 0xf8, 0xf9, 0x30, 0x05, 0x41, 0x92, 0x42, 0x16, 0x40, 0xa2, 0x69, 0x6a, 0x1b, 0x3d,
	0xc6, 0x7a, 0x31, 0x38, 0xf2, 0xd6, 0xed, 0x47, 0x4e, 0xd8, 0xcf, 0x08, 0xa7, 0x2c, 0xd1, 0xef,
	0x57, 0x14, 0x98, 0x24, 0x09, 0xe3, 0xf2, 0x7b, 0xee, 0x84, 0x90, 0x66, 0x10, 0x4c, 0x16, 0x5d,
	0xea, 0x87, 0x29, 0x39, 0x55, 0x93, 0x73, 0xc2, 0xfb, 0xb9, 0x7e, 0xbe, 0xfc, 0xc7, 0xf3, 0x00,
	0x32, 0xd1, 0x33, 0x4d, 0x7a, 0xba, 0xe4, 0xc2, 0x80, 0xc4, 0x34, 0x24, 0x1c, 0x9c, 0xe2, 0xa0,
	0x1e, 0xb6, 0xbe, 0x5a, 0x08, 0x75, 0xc4, 0x24, 0x6d, 0x88, 0xc9, 0xd0, 0x7e, 0x6b, 0xa0, 0x8b,
	0x87, 0x34, 0x0c, 0x21, 0xf1, 0x65, 0x63, 0x7e, 0xd1, 0x0c, 0x84, 0xbe, 0x18, 0xb0, 0x6a, 0x6c,
	0x1a, 0xf5, 0x95, 0xc6, 0x23, 0x3c, 0xad, 0x3b, 0x78, 0x4c, 0x3e, 0x71, 0x7c, 0x31, 0x4c, 0xc1,
	0x5d, 0xfa, 0xf8, 0xe3, 0xe7, 0xf7, 0xb2, 0x51, 0x35, 0xbc, 0x35, 0x25, 0xb9, 0x2b, 0x38, 0xdb,
	0x23, 0x41, 0x51, 0x64, 0x77, 0x50, 0x25, 0xa2, 0xc7, 0x10, 0xfa, 0xa1, 0xc0, 0x55, 0x17, 0x36,
	0x8d, 0x7a, 0xa5, 0xb1, 0x86, 0x95, 0xab, 0xb8, 0x70, 0x15, 0xb7, 0xb5, 0xab, 0xee, 0xd2, 0x89,
	0x5b, 0x7e, 0x6f, 0x98, 0x37, 0x4a, 0x7b, 0x25, 0x0f, 0x49, 0xa4, 0x1a, 0x0c, 0xd0, 0xd9, 0x43,
	0x20, 0x21, 0x64, 0x9a, 0xa8, 0x2c, 0x89, 0x76, 0xe6, 0x9a, 0x63, 0x4f, 0x12, 0xc9, 0xf3, 0x5e,
	0xc9, 0xab, 0x1c, 0x8e, 0xaf, 0xf6, 0x0e, 0x42, 0x3a, 0x7f, 0xd2, 0x83, 0xaa, 0x25, 0x45, 0x36,
	0xb5, 0x88, 0xf0, 0x4f, 0x32, 0x65, 0x24, 0x10, 0xbd, 0x92, 0xf8, 0x40, 0x95, 0x7a, 0x13, 0x98,
	0xda, 0x3e, 0xaa, 0x4c, 0xf0, 0xb7, 0x5a, 0xef, 0x3e, 0xbf, 0xd9, 0xb8, 0x83, 0x9a, 0x8a, 0x22,
	0x60, 0x49, 0x44, 0x7b, 0xba, 0xc7, 0xa2, 0xb7, 0xc6, 0x3f, 0x7a, 0xdb, 0x5a, 0x47, 0x2b, 0xa7,
	0x2d, 0xb7, 0x97, 0x51, 0xb9, 0xb3, 0xff, 0x72, 0xb7, 0xbd, 0x5a, 0x6a, 0x61, 0x41, 0x7c, 0x1d,
	0x5d, 0x9b, 0x92, 0xd8, 0x5d, 0x47, 0xe7, 0xe5, 0x67, 0xe5, 0x9f, 0x9f, 0x43, 0x40, 0x23, 0x0a,
	0x99, 0xbd, 0xf0, 0xcb, 0x35, 0x1e, 0x5b, 0x4b, 0xe6, 0xea, 0xc2, 0xd6, 0x07, 0x4b, 0x0b, 0x7a,
	0x84, 0xc3, 0x53, 0x7a, 0x44, 0xb9, 0x1d, 0x15, 0xf1, 0xc5, 0xe2, 0x2a, 0xb7, 0xa7, 0x32, 0xf3,
	0xf6, 0x8c, 0xe8, 0x70, 0x47, 0x70, 0xc9, 0xe3, 0x28, 0x5e, 0xa5, 0x43, 0x47, 0xf1, 0x2a, 0x21,
	0xb5, 0x27, 0xed, 0xb9, 0x85, 0x94, 0x8d, 0x85, 0x92, 0x8e, 0x58, 0x49, 0x9d, 0x8e, 0xd8, 0x9c,
	0x23, 0xe2, 0x1c, 0xa1, 0xf1, 0x20, 0xf6, 0x55, 0x84, 0x64, 0xcf, 0xfe, 0xab, 0x6e, 0x9a, 0x4b,
	0x87, 0x2c, 0xf7, 0xcc, 0x89, 0x6b, 0x35, 0xcc, 0xba, 0xe1, 0x2d, 0xcb, 0xa7, 0x27, 0xdd, 0x34,
	0x6f, 0x3d, 0x10, 0x81, 0xdd, 0x43, 0xdb, 0xff, 0x0f, 0xec, 0x6f, 0x7e, 0xd5, 0x9e, 0x15, 0x7b,
	0x25, 0xaf, 0xad, 0x87, 0x82, 0xed, 0x3e, 0xba, 0x3b, 0x0b, 0xdb, 0x24, 0xbe, 0x21, 0xf0, 0xb7,
	0xd0, 0xcd, 0x19, 0xf0, 0xee, 0xb9, 0x62, 0x52, 0x61, 0x93, 0x5c, 0x1c, 0xf7, 0xf9, 0xa7, 0xd7,
	0x5f, 0xbe, 0x2d, 0x9a, 0xab, 0x26, 0xda, 0xa6, 0x4c, 0x99, 0x98, 0x66, 0xec, 0x78, 0x38, 0x75,
	0x70, 0xae, 0xfa, 0x7b, 0x1d, 0x88, 0x1f, 0xc1, 0x81, 0xd1, 0x5d, 0x94, 0x7f, 0x84, 0xe6, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x59, 0xf1, 0xb5, 0xb0, 0xf4, 0x05, 0x00, 0x00,
}
