// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/trace/v2/zipkin.proto

package envoy_config_trace_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type ZipkinConfig_CollectorEndpointVersion int32

const (
	ZipkinConfig_HTTP_JSON_V1 ZipkinConfig_CollectorEndpointVersion = 0 // Deprecated: Do not use.
	ZipkinConfig_HTTP_JSON    ZipkinConfig_CollectorEndpointVersion = 1
	ZipkinConfig_HTTP_PROTO   ZipkinConfig_CollectorEndpointVersion = 2
	ZipkinConfig_GRPC         ZipkinConfig_CollectorEndpointVersion = 3
)

var ZipkinConfig_CollectorEndpointVersion_name = map[int32]string{
	0: "HTTP_JSON_V1",
	1: "HTTP_JSON",
	2: "HTTP_PROTO",
	3: "GRPC",
}

var ZipkinConfig_CollectorEndpointVersion_value = map[string]int32{
	"HTTP_JSON_V1": 0,
	"HTTP_JSON":    1,
	"HTTP_PROTO":   2,
	"GRPC":         3,
}

func (x ZipkinConfig_CollectorEndpointVersion) String() string {
	return proto.EnumName(ZipkinConfig_CollectorEndpointVersion_name, int32(x))
}

func (ZipkinConfig_CollectorEndpointVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_02b591b5d30f8add, []int{0, 0}
}

type ZipkinConfig struct {
	CollectorCluster         string                                `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster,proto3" json:"collector_cluster,omitempty"`
	CollectorEndpoint        string                                `protobuf:"bytes,2,opt,name=collector_endpoint,json=collectorEndpoint,proto3" json:"collector_endpoint,omitempty"`
	TraceId_128Bit           bool                                  `protobuf:"varint,3,opt,name=trace_id_128bit,json=traceId128bit,proto3" json:"trace_id_128bit,omitempty"`
	SharedSpanContext        *wrappers.BoolValue                   `protobuf:"bytes,4,opt,name=shared_span_context,json=sharedSpanContext,proto3" json:"shared_span_context,omitempty"`
	CollectorEndpointVersion ZipkinConfig_CollectorEndpointVersion `protobuf:"varint,5,opt,name=collector_endpoint_version,json=collectorEndpointVersion,proto3,enum=envoy.config.trace.v2.ZipkinConfig_CollectorEndpointVersion" json:"collector_endpoint_version,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                              `json:"-"`
	XXX_unrecognized         []byte                                `json:"-"`
	XXX_sizecache            int32                                 `json:"-"`
}

func (m *ZipkinConfig) Reset()         { *m = ZipkinConfig{} }
func (m *ZipkinConfig) String() string { return proto.CompactTextString(m) }
func (*ZipkinConfig) ProtoMessage()    {}
func (*ZipkinConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_02b591b5d30f8add, []int{0}
}

func (m *ZipkinConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZipkinConfig.Unmarshal(m, b)
}
func (m *ZipkinConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZipkinConfig.Marshal(b, m, deterministic)
}
func (m *ZipkinConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZipkinConfig.Merge(m, src)
}
func (m *ZipkinConfig) XXX_Size() int {
	return xxx_messageInfo_ZipkinConfig.Size(m)
}
func (m *ZipkinConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ZipkinConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ZipkinConfig proto.InternalMessageInfo

func (m *ZipkinConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *ZipkinConfig) GetCollectorEndpoint() string {
	if m != nil {
		return m.CollectorEndpoint
	}
	return ""
}

func (m *ZipkinConfig) GetTraceId_128Bit() bool {
	if m != nil {
		return m.TraceId_128Bit
	}
	return false
}

func (m *ZipkinConfig) GetSharedSpanContext() *wrappers.BoolValue {
	if m != nil {
		return m.SharedSpanContext
	}
	return nil
}

func (m *ZipkinConfig) GetCollectorEndpointVersion() ZipkinConfig_CollectorEndpointVersion {
	if m != nil {
		return m.CollectorEndpointVersion
	}
	return ZipkinConfig_HTTP_JSON_V1
}

func init() {
	proto.RegisterEnum("envoy.config.trace.v2.ZipkinConfig_CollectorEndpointVersion", ZipkinConfig_CollectorEndpointVersion_name, ZipkinConfig_CollectorEndpointVersion_value)
	proto.RegisterType((*ZipkinConfig)(nil), "envoy.config.trace.v2.ZipkinConfig")
}

func init() { proto.RegisterFile("envoy/config/trace/v2/zipkin.proto", fileDescriptor_02b591b5d30f8add) }

var fileDescriptor_02b591b5d30f8add = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x71, 0x5a, 0x46, 0xe7, 0xfd, 0x21, 0x33, 0x42, 0x44, 0x91, 0x40, 0x51, 0x27, 0xa1,
	0x9e, 0x1c, 0x2d, 0x20, 0xc4, 0x61, 0xa7, 0x54, 0x08, 0xd8, 0x81, 0x46, 0x59, 0xd5, 0x03, 0x97,
	0xc8, 0x4d, 0xbc, 0x62, 0x11, 0xf9, 0xb5, 0x1c, 0x27, 0x6c, 0x3b, 0x71, 0xe7, 0x8b, 0xf0, 0x01,
	0x38, 0xf1, 0x09, 0xb8, 0x70, 0xe0, 0xd3, 0x20, 0x71, 0x42, 0xb5, 0xd7, 0x8d, 0xb1, 0xf5, 0x16,
	0xbf, 0xef, 0xef, 0x79, 0xf4, 0x3e, 0x4f, 0xf0, 0x90, 0xcb, 0x0e, 0xce, 0xe2, 0x12, 0xe4, 0x89,
	0x58, 0xc4, 0x46, 0xb3, 0x92, 0xc7, 0x5d, 0x12, 0x9f, 0x0b, 0xf5, 0x51, 0x48, 0xaa, 0x34, 0x18,
	0x20, 0x0f, 0x2d, 0x43, 0x1d, 0x43, 0x2d, 0x43, 0xbb, 0x24, 0x7c, 0xb2, 0x00, 0x58, 0xd4, 0x3c,
	0xb6, 0xd0, 0xbc, 0x3d, 0x89, 0x3f, 0x69, 0xa6, 0x14, 0xd7, 0x8d, 0x93, 0x85, 0xfb, 0xce, 0x9a,
	0x49, 0x09, 0x86, 0x19, 0x01, 0xb2, 0x89, 0x2b, 0xae, 0x34, 0x2f, 0xed, 0xe3, 0x02, 0x7a, 0xdc,
	0x56, 0x8a, 0x5d, 0x63, 0x1a, 0xc3, 0x4c, 0xbb, 0xf2, 0x78, 0xd4, 0xb1, 0x5a, 0x54, 0xcc, 0xf0,
	0x78, 0xf5, 0xe1, 0x16, 0xc3, 0x9f, 0x3d, 0xbc, 0xfd, 0xde, 0x1e, 0x39, 0xb6, 0x67, 0x91, 0xe7,
	0x78, 0xaf, 0x84, 0xba, 0xe6, 0xa5, 0x01, 0x5d, 0x94, 0x75, 0xdb, 0x18, 0xae, 0x03, 0x14, 0xa1,
	0xd1, 0x66, 0x7a, 0xef, 0x4f, 0xda, 0xd7, 0x5e, 0x84, 0x72, 0xff, 0x92, 0x18, 0x3b, 0x80, 0xbc,
	0xc0, 0xe4, 0x4a, 0xc5, 0x65, 0xa5, 0x40, 0x48, 0x13, 0x78, 0xd7, 0x65, 0x57, 0xc6, 0xaf, 0x2e,
	0x08, 0xf2, 0x14, 0xdf, 0xb7, 0x3d, 0x14, 0xa2, 0x2a, 0x0e, 0x92, 0x97, 0x73, 0x61, 0x82, 0x5e,
	0x84, 0x46, 0x83, 0x7c, 0xc7, 0x8e, 0xdf, 0x56, 0x6e, 0x48, 0x8e, 0xf0, 0x83, 0xe6, 0x03, 0xd3,
	0xbc, 0x2a, 0x1a, 0xc5, 0x64, 0x51, 0x82, 0x34, 0xfc, 0xd4, 0x04, 0xfd, 0x08, 0x8d, 0xb6, 0x92,
	0x90, 0xba, 0x06, 0xe9, 0xaa, 0x41, 0x9a, 0x02, 0xd4, 0x33, 0x56, 0xb7, 0x3c, 0xdf, 0x73, 0xb2,
	0x63, 0xc5, 0x96, 0x01, 0x97, 0x22, 0x72, 0x8e, 0xc3, 0x9b, 0xb7, 0x16, 0x1d, 0xd7, 0x8d, 0x00,
	0x19, 0xdc, 0x8d, 0xd0, 0x68, 0x37, 0x39, 0xa4, 0xb7, 0xfe, 0x2b, 0xfa, 0x6f, 0x55, 0x74, 0xfc,
	0x7f, 0x9c, 0x99, 0xf3, 0xc8, 0x83, 0x72, 0xcd, 0x66, 0x58, 0xe0, 0x60, 0x9d, 0x8a, 0x84, 0x78,
	0xfb, 0xcd, 0x74, 0x9a, 0x15, 0x47, 0xc7, 0x93, 0x77, 0xc5, 0xec, 0xc0, 0xbf, 0x13, 0x0e, 0xbe,
	0xfe, 0xfe, 0xf6, 0xc5, 0x43, 0x03, 0x44, 0x76, 0xf0, 0xe6, 0xe5, 0xce, 0x47, 0x64, 0x17, 0x63,
	0xfb, 0xcc, 0xf2, 0xc9, 0x74, 0xe2, 0x7b, 0x64, 0x80, 0xfb, 0xaf, 0xf3, 0x6c, 0xec, 0xf7, 0xd2,
	0xc3, 0xef, 0x9f, 0x7f, 0xfc, 0xda, 0xf0, 0x7c, 0x84, 0xf7, 0x05, 0xb8, 0x10, 0x4a, 0xc3, 0xe9,
	0xd9, 0xed, 0x79, 0xd2, 0x2d, 0x17, 0x28, 0x5b, 0x16, 0x97, 0xa1, 0xf9, 0x86, 0x6d, 0xf0, 0xd9,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x19, 0x41, 0x06, 0x63, 0xce, 0x02, 0x00, 0x00,
}
