// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/tracers/lightstep/v4alpha/lightstep.proto

package envoy_extensions_tracers_lightstep_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type LightstepConfig_PropagationMode int32

const (
	LightstepConfig_ENVOY         LightstepConfig_PropagationMode = 0
	LightstepConfig_LIGHTSTEP     LightstepConfig_PropagationMode = 1
	LightstepConfig_B3            LightstepConfig_PropagationMode = 2
	LightstepConfig_TRACE_CONTEXT LightstepConfig_PropagationMode = 3
)

var LightstepConfig_PropagationMode_name = map[int32]string{
	0: "ENVOY",
	1: "LIGHTSTEP",
	2: "B3",
	3: "TRACE_CONTEXT",
}

var LightstepConfig_PropagationMode_value = map[string]int32{
	"ENVOY":         0,
	"LIGHTSTEP":     1,
	"B3":            2,
	"TRACE_CONTEXT": 3,
}

func (x LightstepConfig_PropagationMode) String() string {
	return proto.EnumName(LightstepConfig_PropagationMode_name, int32(x))
}

func (LightstepConfig_PropagationMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ce663b309afba8ca, []int{0, 0}
}

type LightstepConfig struct {
	CollectorCluster     string                            `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster,proto3" json:"collector_cluster,omitempty"`
	AccessTokenFile      string                            `protobuf:"bytes,2,opt,name=access_token_file,json=accessTokenFile,proto3" json:"access_token_file,omitempty"`
	PropagationModes     []LightstepConfig_PropagationMode `protobuf:"varint,3,rep,packed,name=propagation_modes,json=propagationModes,proto3,enum=envoy.extensions.tracers.lightstep.v4alpha.LightstepConfig_PropagationMode" json:"propagation_modes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *LightstepConfig) Reset()         { *m = LightstepConfig{} }
func (m *LightstepConfig) String() string { return proto.CompactTextString(m) }
func (*LightstepConfig) ProtoMessage()    {}
func (*LightstepConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce663b309afba8ca, []int{0}
}

func (m *LightstepConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LightstepConfig.Unmarshal(m, b)
}
func (m *LightstepConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LightstepConfig.Marshal(b, m, deterministic)
}
func (m *LightstepConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LightstepConfig.Merge(m, src)
}
func (m *LightstepConfig) XXX_Size() int {
	return xxx_messageInfo_LightstepConfig.Size(m)
}
func (m *LightstepConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LightstepConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LightstepConfig proto.InternalMessageInfo

func (m *LightstepConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *LightstepConfig) GetAccessTokenFile() string {
	if m != nil {
		return m.AccessTokenFile
	}
	return ""
}

func (m *LightstepConfig) GetPropagationModes() []LightstepConfig_PropagationMode {
	if m != nil {
		return m.PropagationModes
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.extensions.tracers.lightstep.v4alpha.LightstepConfig_PropagationMode", LightstepConfig_PropagationMode_name, LightstepConfig_PropagationMode_value)
	proto.RegisterType((*LightstepConfig)(nil), "envoy.extensions.tracers.lightstep.v4alpha.LightstepConfig")
}

func init() {
	proto.RegisterFile("envoy/extensions/tracers/lightstep/v4alpha/lightstep.proto", fileDescriptor_ce663b309afba8ca)
}

var fileDescriptor_ce663b309afba8ca = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x6b, 0xd4, 0x40,
	0x18, 0xc5, 0x9d, 0xc4, 0x6e, 0xd9, 0x81, 0xed, 0x4e, 0xe6, 0xe2, 0x52, 0x50, 0xd6, 0x05, 0xb1,
	0x88, 0x4c, 0xc0, 0xf4, 0x20, 0xbd, 0x99, 0x90, 0xaa, 0x58, 0xdb, 0xb0, 0x06, 0xd1, 0x53, 0x18,
	0x93, 0x69, 0x3a, 0x38, 0xce, 0x0c, 0x33, 0xb3, 0xa1, 0xbd, 0x89, 0x17, 0x3d, 0x7b, 0xf4, 0x4f,
	0xf1, 0x2e, 0x78, 0xf5, 0xdf, 0x29, 0x1e, 0x64, 0x93, 0x6d, 0xc3, 0xae, 0x17, 0xbd, 0x85, 0x7c,
	0xef, 0xf7, 0xbe, 0xf7, 0xe6, 0x83, 0x07, 0x4c, 0x36, 0xea, 0x22, 0x64, 0xe7, 0x8e, 0x49, 0xcb,
	0x95, 0xb4, 0xa1, 0x33, 0xb4, 0x64, 0xc6, 0x86, 0x82, 0xd7, 0x67, 0xce, 0x3a, 0xa6, 0xc3, 0x66,
	0x9f, 0x0a, 0x7d, 0x46, 0xfb, 0x3f, 0x44, 0x1b, 0xe5, 0x14, 0x7e, 0xd0, 0xb2, 0xa4, 0x67, 0xc9,
	0x8a, 0x25, 0xbd, 0x72, 0xc5, 0xee, 0xde, 0x5e, 0x54, 0x9a, 0x86, 0x54, 0x4a, 0xe5, 0xa8, 0x6b,
	0xf7, 0x58, 0x47, 0xdd, 0xc2, 0x76, 0x56, 0xbb, 0x77, 0xff, 0x1a, 0x37, 0xcc, 0x2c, 0x3d, 0xb9,
	0xac, 0x57, 0x92, 0x5b, 0x0d, 0x15, 0xbc, 0xa2, 0x8e, 0x85, 0x57, 0x1f, 0xdd, 0x60, 0xf6, 0xdb,
	0x83, 0xe3, 0xa3, 0xab, 0x85, 0x89, 0x92, 0xa7, 0xbc, 0xc6, 0xfb, 0x30, 0x28, 0x95, 0x10, 0xac,
	0x74, 0xca, 0x14, 0xa5, 0x58, 0x58, 0xc7, 0xcc, 0x04, 0x4c, 0xc1, 0xde, 0x30, 0xde, 0xbe, 0x8c,
	0x6f, 0x1a, 0x6f, 0x0a, 0xe6, 0xe8, 0x5a, 0x91, 0x74, 0x02, 0x1c, 0xc1, 0x80, 0x96, 0x25, 0xb3,
	0xb6, 0x70, 0xea, 0x3d, 0x93, 0xc5, 0x29, 0x17, 0x6c, 0xe2, 0xad, 0x53, 0xe3, 0x4e, 0x91, 0x2f,
	0x05, 0x87, 0x5c, 0x30, 0xfc, 0x19, 0xc0, 0x40, 0x1b, 0xa5, 0x69, 0xdd, 0x26, 0x2f, 0x3e, 0xa8,
	0x8a, 0xd9, 0x89, 0x3f, 0xf5, 0xf7, 0x76, 0x1e, 0xbd, 0x20, 0xff, 0xfe, 0x44, 0x64, 0xa3, 0x03,
	0xc9, 0x7a, 0xd3, 0x97, 0xaa, 0x62, 0xf1, 0xe8, 0x32, 0x86, 0x5f, 0xc1, 0xf6, 0x6c, 0xeb, 0x13,
	0xf0, 0x10, 0x98, 0x23, 0xbd, 0x3e, 0xb7, 0xb3, 0x43, 0x38, 0xde, 0x60, 0xf0, 0x10, 0x6e, 0xa5,
	0xc7, 0xaf, 0x4f, 0xde, 0xa2, 0x1b, 0x78, 0x04, 0x87, 0x47, 0xcf, 0x9f, 0x3e, 0xcb, 0x5f, 0xe5,
	0x69, 0x86, 0x00, 0x1e, 0x40, 0x2f, 0x8e, 0x90, 0x87, 0x03, 0x38, 0xca, 0xe7, 0x4f, 0x92, 0xb4,
	0x48, 0x4e, 0x8e, 0xf3, 0xf4, 0x4d, 0x8e, 0xfc, 0x83, 0x87, 0xdf, 0x7e, 0x7c, 0xb9, 0x73, 0x1f,
	0xde, 0xeb, 0xb2, 0x97, 0x5d, 0x9a, 0x36, 0x37, 0x69, 0xa2, 0xcd, 0x98, 0x71, 0xfe, 0xfd, 0xe3,
	0xcf, 0x5f, 0x03, 0x0f, 0xf9, 0xf0, 0x31, 0x57, 0x5d, 0x5f, 0x6d, 0xd4, 0xf9, 0xc5, 0x7f, 0x54,
	0x8f, 0x77, 0xae, 0x4d, 0xb3, 0xe5, 0x49, 0x33, 0xf0, 0x6e, 0xd0, 0xde, 0x36, 0xfa, 0x13, 0x00,
	0x00, 0xff, 0xff, 0x1b, 0x64, 0xdb, 0xe0, 0xa0, 0x02, 0x00, 0x00,
}
