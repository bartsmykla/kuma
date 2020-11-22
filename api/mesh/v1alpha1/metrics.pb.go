// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mesh/v1alpha1/metrics.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// Metrics defines configuration for metrics that should be collected and
// exposed by dataplanes.
type Metrics struct {
	// Name of the enabled backend
	EnabledBackend string `protobuf:"bytes,1,opt,name=enabledBackend,proto3" json:"enabledBackend,omitempty"`
	// List of available Metrics backends
	Backends             []*MetricsBackend `protobuf:"bytes,2,rep,name=backends,proto3" json:"backends,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Metrics) Reset()         { *m = Metrics{} }
func (m *Metrics) String() string { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()    {}
func (*Metrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dd8c7f420ce268c, []int{0}
}

func (m *Metrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metrics.Unmarshal(m, b)
}
func (m *Metrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metrics.Marshal(b, m, deterministic)
}
func (m *Metrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metrics.Merge(m, src)
}
func (m *Metrics) XXX_Size() int {
	return xxx_messageInfo_Metrics.Size(m)
}
func (m *Metrics) XXX_DiscardUnknown() {
	xxx_messageInfo_Metrics.DiscardUnknown(m)
}

var xxx_messageInfo_Metrics proto.InternalMessageInfo

func (m *Metrics) GetEnabledBackend() string {
	if m != nil {
		return m.EnabledBackend
	}
	return ""
}

func (m *Metrics) GetBackends() []*MetricsBackend {
	if m != nil {
		return m.Backends
	}
	return nil
}

// MetricsBackend defines metric backends
type MetricsBackend struct {
	// Name of the backend, can be then used in Mesh.metrics.enabledBackend
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Type of the backend (Kuma ships with 'prometheus')
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Configuration of the backend
	Conf                 *_struct.Struct `protobuf:"bytes,3,opt,name=conf,proto3" json:"conf,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MetricsBackend) Reset()         { *m = MetricsBackend{} }
func (m *MetricsBackend) String() string { return proto.CompactTextString(m) }
func (*MetricsBackend) ProtoMessage()    {}
func (*MetricsBackend) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dd8c7f420ce268c, []int{1}
}

func (m *MetricsBackend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsBackend.Unmarshal(m, b)
}
func (m *MetricsBackend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsBackend.Marshal(b, m, deterministic)
}
func (m *MetricsBackend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsBackend.Merge(m, src)
}
func (m *MetricsBackend) XXX_Size() int {
	return xxx_messageInfo_MetricsBackend.Size(m)
}
func (m *MetricsBackend) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsBackend.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsBackend proto.InternalMessageInfo

func (m *MetricsBackend) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetricsBackend) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MetricsBackend) GetConf() *_struct.Struct {
	if m != nil {
		return m.Conf
	}
	return nil
}

// PrometheusMetricsBackendConfig defines configuration of Prometheus backend
type PrometheusMetricsBackendConfig struct {
	// Port on which a dataplane should expose HTTP endpoint with Prometheus
	// metrics.
	Port uint32 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	// Path on which a dataplane should expose HTTP endpoint with Prometheus
	// metrics.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// Tags associated with an application this dataplane is deployed next to,
	// e.g. service=web, version=1.0.
	// `service` tag is mandatory.
	Tags map[string]string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// If true then endpoints for scraping metrics won't require mTLS even if mTLS
	// is enabled in Mesh. If nil, then it is treated as false.
	SkipMTLS             *wrappers.BoolValue `protobuf:"bytes,4,opt,name=skipMTLS,proto3" json:"skipMTLS,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PrometheusMetricsBackendConfig) Reset()         { *m = PrometheusMetricsBackendConfig{} }
func (m *PrometheusMetricsBackendConfig) String() string { return proto.CompactTextString(m) }
func (*PrometheusMetricsBackendConfig) ProtoMessage()    {}
func (*PrometheusMetricsBackendConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dd8c7f420ce268c, []int{2}
}

func (m *PrometheusMetricsBackendConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrometheusMetricsBackendConfig.Unmarshal(m, b)
}
func (m *PrometheusMetricsBackendConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrometheusMetricsBackendConfig.Marshal(b, m, deterministic)
}
func (m *PrometheusMetricsBackendConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrometheusMetricsBackendConfig.Merge(m, src)
}
func (m *PrometheusMetricsBackendConfig) XXX_Size() int {
	return xxx_messageInfo_PrometheusMetricsBackendConfig.Size(m)
}
func (m *PrometheusMetricsBackendConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PrometheusMetricsBackendConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PrometheusMetricsBackendConfig proto.InternalMessageInfo

func (m *PrometheusMetricsBackendConfig) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *PrometheusMetricsBackendConfig) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *PrometheusMetricsBackendConfig) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *PrometheusMetricsBackendConfig) GetSkipMTLS() *wrappers.BoolValue {
	if m != nil {
		return m.SkipMTLS
	}
	return nil
}

func init() {
	proto.RegisterType((*Metrics)(nil), "kuma.mesh.v1alpha1.Metrics")
	proto.RegisterType((*MetricsBackend)(nil), "kuma.mesh.v1alpha1.MetricsBackend")
	proto.RegisterType((*PrometheusMetricsBackendConfig)(nil), "kuma.mesh.v1alpha1.PrometheusMetricsBackendConfig")
	proto.RegisterMapType((map[string]string)(nil), "kuma.mesh.v1alpha1.PrometheusMetricsBackendConfig.TagsEntry")
}

func init() { proto.RegisterFile("mesh/v1alpha1/metrics.proto", fileDescriptor_7dd8c7f420ce268c) }

var fileDescriptor_7dd8c7f420ce268c = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x4b, 0xeb, 0x40,
	0x14, 0xc5, 0xc9, 0x9f, 0xf7, 0x5e, 0x3b, 0xe5, 0x95, 0xc7, 0xf0, 0xc0, 0x10, 0xa5, 0x94, 0x2c,
	0x24, 0x28, 0x4c, 0x68, 0x05, 0x15, 0x11, 0x17, 0x15, 0x77, 0x16, 0x4a, 0x5a, 0x5c, 0xb8, 0x9b,
	0xa4, 0xd3, 0x24, 0xe4, 0xcf, 0x4c, 0x67, 0x26, 0x95, 0x7e, 0x06, 0xbf, 0xb4, 0xcc, 0x24, 0x2d,
	0xb4, 0x15, 0x57, 0x39, 0xb9, 0xf7, 0xdc, 0x7b, 0xee, 0x2f, 0x04, 0x9c, 0x97, 0x44, 0xa4, 0xc1,
	0x66, 0x84, 0x0b, 0x96, 0xe2, 0x51, 0x50, 0x12, 0xc9, 0xb3, 0x58, 0x20, 0xc6, 0xa9, 0xa4, 0x10,
	0xe6, 0x75, 0x89, 0x91, 0x72, 0xa0, 0x9d, 0xc3, 0xbd, 0x48, 0x28, 0x4d, 0x0a, 0x12, 0x68, 0x47,
	0x54, 0xaf, 0x02, 0x21, 0x79, 0x1d, 0xcb, 0x66, 0xc2, 0x1d, 0x1c, 0x77, 0x3f, 0x38, 0x66, 0x8c,
	0xf0, 0x76, 0xa3, 0xb7, 0x06, 0x7f, 0xa6, 0x4d, 0x04, 0xbc, 0x04, 0x7d, 0x52, 0xe1, 0xa8, 0x20,
	0xcb, 0x09, 0x8e, 0x73, 0x52, 0x2d, 0x1d, 0x63, 0x68, 0xf8, 0xdd, 0xf0, 0xa8, 0x0a, 0x9f, 0x40,
	0x27, 0x6a, 0xa4, 0x70, 0xcc, 0xa1, 0xe5, 0xf7, 0xc6, 0x1e, 0x3a, 0xbd, 0x0b, 0xb5, 0x6b, 0xdb,
	0xa9, 0x70, 0x3f, 0xe3, 0x11, 0xd0, 0x3f, 0xec, 0x41, 0x08, 0xec, 0x0a, 0x97, 0xa4, 0xcd, 0xd3,
	0x5a, 0xd5, 0xe4, 0x96, 0x11, 0xc7, 0x6c, 0x6a, 0x4a, 0xc3, 0x6b, 0x60, 0xc7, 0xb4, 0x5a, 0x39,
	0xd6, 0xd0, 0xf0, 0x7b, 0xe3, 0x33, 0xd4, 0xb0, 0xa1, 0x1d, 0x1b, 0x9a, 0x6b, 0xf2, 0x50, 0x9b,
	0xbc, 0x4f, 0x13, 0x0c, 0x66, 0x9c, 0x96, 0x44, 0xa6, 0xa4, 0x16, 0x87, 0x89, 0xcf, 0xb4, 0x5a,
	0x65, 0x89, 0xca, 0x60, 0x94, 0x4b, 0x9d, 0xfb, 0x37, 0xd4, 0x5a, 0xd7, 0xb0, 0x4c, 0x77, 0xb9,
	0x4a, 0xc3, 0x19, 0xb0, 0x25, 0x4e, 0x84, 0x63, 0x69, 0xda, 0xc7, 0xef, 0x68, 0x7f, 0x4e, 0x42,
	0x0b, 0x9c, 0x88, 0x97, 0x4a, 0xf2, 0x6d, 0xa8, 0x37, 0xc1, 0x5b, 0xd0, 0x11, 0x79, 0xc6, 0xa6,
	0x8b, 0xd7, 0xb9, 0x63, 0x6b, 0x1a, 0xf7, 0x84, 0x66, 0x42, 0x69, 0xf1, 0x86, 0x8b, 0x9a, 0x84,
	0x7b, 0xaf, 0x7b, 0x07, 0xba, 0xfb, 0x55, 0xf0, 0x1f, 0xb0, 0x72, 0xb2, 0x6d, 0xbf, 0x9a, 0x92,
	0xf0, 0x3f, 0xf8, 0xb5, 0x51, 0x13, 0xed, 0xf5, 0xcd, 0xcb, 0x83, 0x79, 0x6f, 0x4c, 0xae, 0xde,
	0xfd, 0x24, 0x93, 0x69, 0x1d, 0xa1, 0x98, 0x96, 0x81, 0x02, 0x48, 0xd7, 0xfa, 0x11, 0x60, 0x96,
	0x05, 0x07, 0xff, 0x5c, 0xf4, 0x5b, 0x9f, 0x70, 0xf3, 0x15, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x3f,
	0x4c, 0x82, 0x8b, 0x02, 0x00, 0x00,
}
