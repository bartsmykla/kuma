// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mesh/v1alpha1/externalservice.proto

package v1alpha1

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	"github.com/kumahq/kuma/api/system/v1alpha1"

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

// ExternalService defines configuration of the externaly accessible service
type ExternalService struct {
	Networking *ExternalService_Networking `protobuf:"bytes,1,opt,name=networking,proto3" json:"networking,omitempty"`
	// Tags associated with the external service,
	// e.g. kuma.io/service=web, kuma.io/protocol, version=1.0.
	Tags                 map[string]string `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ExternalService) Reset()         { *m = ExternalService{} }
func (m *ExternalService) String() string { return proto.CompactTextString(m) }
func (*ExternalService) ProtoMessage()    {}
func (*ExternalService) Descriptor() ([]byte, []int) {
	return fileDescriptor_df6b95621b774a94, []int{0}
}

func (m *ExternalService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExternalService.Unmarshal(m, b)
}
func (m *ExternalService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExternalService.Marshal(b, m, deterministic)
}
func (m *ExternalService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalService.Merge(m, src)
}
func (m *ExternalService) XXX_Size() int {
	return xxx_messageInfo_ExternalService.Size(m)
}
func (m *ExternalService) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalService.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalService proto.InternalMessageInfo

func (m *ExternalService) GetNetworking() *ExternalService_Networking {
	if m != nil {
		return m.Networking
	}
	return nil
}

func (m *ExternalService) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// Networking describes the properties of the external service connectivity
type ExternalService_Networking struct {
	// Address of the external service
	Address              string                          `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Tls                  *ExternalService_Networking_TLS `protobuf:"bytes,2,opt,name=tls,proto3" json:"tls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ExternalService_Networking) Reset()         { *m = ExternalService_Networking{} }
func (m *ExternalService_Networking) String() string { return proto.CompactTextString(m) }
func (*ExternalService_Networking) ProtoMessage()    {}
func (*ExternalService_Networking) Descriptor() ([]byte, []int) {
	return fileDescriptor_df6b95621b774a94, []int{0, 0}
}

func (m *ExternalService_Networking) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExternalService_Networking.Unmarshal(m, b)
}
func (m *ExternalService_Networking) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExternalService_Networking.Marshal(b, m, deterministic)
}
func (m *ExternalService_Networking) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalService_Networking.Merge(m, src)
}
func (m *ExternalService_Networking) XXX_Size() int {
	return xxx_messageInfo_ExternalService_Networking.Size(m)
}
func (m *ExternalService_Networking) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalService_Networking.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalService_Networking proto.InternalMessageInfo

func (m *ExternalService_Networking) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ExternalService_Networking) GetTls() *ExternalService_Networking_TLS {
	if m != nil {
		return m.Tls
	}
	return nil
}

// TLS
type ExternalService_Networking_TLS struct {
	// denotes that the external service uses TLS
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Data source for the certificate of CA
	CaCert *v1alpha1.DataSource `protobuf:"bytes,2,opt,name=ca_cert,json=caCert,proto3" json:"ca_cert,omitempty"`
	// Data source for the authentication
	ClientCert *v1alpha1.DataSource `protobuf:"bytes,3,opt,name=client_cert,json=clientCert,proto3" json:"client_cert,omitempty"`
	// Data source for the authentication
	ClientKey            *v1alpha1.DataSource `protobuf:"bytes,4,opt,name=client_key,json=clientKey,proto3" json:"client_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ExternalService_Networking_TLS) Reset()         { *m = ExternalService_Networking_TLS{} }
func (m *ExternalService_Networking_TLS) String() string { return proto.CompactTextString(m) }
func (*ExternalService_Networking_TLS) ProtoMessage()    {}
func (*ExternalService_Networking_TLS) Descriptor() ([]byte, []int) {
	return fileDescriptor_df6b95621b774a94, []int{0, 0, 0}
}

func (m *ExternalService_Networking_TLS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExternalService_Networking_TLS.Unmarshal(m, b)
}
func (m *ExternalService_Networking_TLS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExternalService_Networking_TLS.Marshal(b, m, deterministic)
}
func (m *ExternalService_Networking_TLS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalService_Networking_TLS.Merge(m, src)
}
func (m *ExternalService_Networking_TLS) XXX_Size() int {
	return xxx_messageInfo_ExternalService_Networking_TLS.Size(m)
}
func (m *ExternalService_Networking_TLS) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalService_Networking_TLS.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalService_Networking_TLS proto.InternalMessageInfo

func (m *ExternalService_Networking_TLS) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *ExternalService_Networking_TLS) GetCaCert() *v1alpha1.DataSource {
	if m != nil {
		return m.CaCert
	}
	return nil
}

func (m *ExternalService_Networking_TLS) GetClientCert() *v1alpha1.DataSource {
	if m != nil {
		return m.ClientCert
	}
	return nil
}

func (m *ExternalService_Networking_TLS) GetClientKey() *v1alpha1.DataSource {
	if m != nil {
		return m.ClientKey
	}
	return nil
}

func init() {
	proto.RegisterType((*ExternalService)(nil), "kuma.mesh.v1alpha1.ExternalService")
	proto.RegisterMapType((map[string]string)(nil), "kuma.mesh.v1alpha1.ExternalService.TagsEntry")
	proto.RegisterType((*ExternalService_Networking)(nil), "kuma.mesh.v1alpha1.ExternalService.Networking")
	proto.RegisterType((*ExternalService_Networking_TLS)(nil), "kuma.mesh.v1alpha1.ExternalService.Networking.TLS")
}

func init() {
	proto.RegisterFile("mesh/v1alpha1/externalservice.proto", fileDescriptor_df6b95621b774a94)
}

var fileDescriptor_df6b95621b774a94 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbd, 0x8e, 0xda, 0x40,
	0x10, 0xc7, 0x65, 0x9b, 0x0f, 0x7b, 0x28, 0x12, 0xad, 0x22, 0xc5, 0x72, 0x65, 0x25, 0x0d, 0x4d,
	0x16, 0x41, 0x8a, 0x7c, 0x34, 0x51, 0x08, 0x54, 0x20, 0x0a, 0x9b, 0x2a, 0x0d, 0x1a, 0xec, 0x11,
	0x58, 0x18, 0x1b, 0xed, 0x2e, 0xbe, 0xf3, 0xab, 0xdc, 0x8b, 0xdc, 0x53, 0xdc, 0x2b, 0xdc, 0x83,
	0x5c, 0x75, 0xb2, 0x17, 0xc3, 0xe9, 0xae, 0xe1, 0xba, 0x1d, 0x69, 0xfe, 0xbf, 0xdf, 0xf8, 0x2f,
	0xc3, 0xd7, 0x3d, 0xc9, 0xed, 0xa0, 0x18, 0x62, 0x7a, 0xd8, 0xe2, 0x70, 0x40, 0xb7, 0x8a, 0x44,
	0x86, 0xa9, 0x24, 0x51, 0x24, 0x11, 0xf1, 0x83, 0xc8, 0x55, 0xce, 0xd8, 0xee, 0xb8, 0x47, 0x5e,
	0x6d, 0xf2, 0x66, 0xd3, 0xfb, 0x5c, 0x60, 0x9a, 0xc4, 0xa8, 0x68, 0xd0, 0x3c, 0xf4, 0xb2, 0xe7,
	0xcb, 0x52, 0x2a, 0xda, 0x5f, 0x98, 0x31, 0x2a, 0x94, 0xf9, 0x51, 0x34, 0xb8, 0x2f, 0xf7, 0x2d,
	0xf8, 0x30, 0x3d, 0x89, 0x42, 0x2d, 0x62, 0x0b, 0x80, 0x8c, 0xd4, 0x4d, 0x2e, 0x76, 0x49, 0xb6,
	0x71, 0x0d, 0xdf, 0xe8, 0xf7, 0x46, 0x9c, 0xbf, 0xf5, 0xf2, 0x57, 0x41, 0xbe, 0x38, 0xa7, 0x82,
	0x17, 0x04, 0x36, 0x83, 0x96, 0xc2, 0x8d, 0x74, 0x4d, 0xdf, 0xea, 0xf7, 0x46, 0xdf, 0xae, 0x21,
	0x2d, 0x71, 0x23, 0xa7, 0x99, 0x12, 0xe5, 0xd8, 0x7e, 0x1a, 0xb7, 0xef, 0x0c, 0xd3, 0x36, 0x82,
	0x1a, 0xe2, 0x3d, 0x98, 0x00, 0x17, 0x0f, 0x73, 0xa1, 0x8b, 0x71, 0x2c, 0x48, 0xca, 0xfa, 0x50,
	0x27, 0x68, 0x46, 0x36, 0x01, 0x4b, 0xa5, 0x95, 0xb4, 0x3a, 0x7f, 0xf4, 0xbe, 0xf3, 0xf9, 0x72,
	0x1e, 0x06, 0x55, 0xdc, 0x7b, 0x34, 0xc0, 0x5a, 0xce, 0xc3, 0xca, 0x43, 0x19, 0xae, 0x53, 0x8a,
	0x6b, 0x8f, 0x1d, 0x34, 0x23, 0xfb, 0x05, 0xdd, 0x08, 0x57, 0x11, 0x09, 0x75, 0x72, 0xf9, 0xda,
	0xa5, 0xab, 0xbf, 0xd8, 0x26, 0xa8, 0x30, 0xac, 0xab, 0x0f, 0x3a, 0x11, 0xfe, 0x23, 0xa1, 0xd8,
	0x5f, 0xe8, 0x45, 0x69, 0x42, 0x99, 0xd2, 0x71, 0xeb, 0xca, 0x38, 0xe8, 0x50, 0x8d, 0xf8, 0x03,
	0xa7, 0x69, 0xb5, 0xa3, 0xd2, 0x6d, 0x5d, 0x49, 0x70, 0x74, 0x66, 0x46, 0xa5, 0xf7, 0x03, 0x9c,
	0x73, 0xd9, 0xec, 0x23, 0x58, 0x15, 0x46, 0x37, 0x59, 0x3d, 0xd9, 0x27, 0x68, 0x17, 0x98, 0x1e,
	0xa9, 0xfe, 0x36, 0x27, 0xd0, 0xc3, 0x6f, 0xf3, 0xa7, 0x31, 0x86, 0xff, 0x76, 0x83, 0x5e, 0x77,
	0xea, 0x9f, 0xe9, 0xfb, 0x73, 0x00, 0x00, 0x00, 0xff, 0xff, 0x35, 0xf9, 0x68, 0xc8, 0xc2, 0x02,
	0x00, 0x00,
}
