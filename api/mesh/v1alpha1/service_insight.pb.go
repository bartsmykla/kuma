// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mesh/v1alpha1/service_insight.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ServiceInsight struct {
	LastSync             *timestamp.Timestamp                     `protobuf:"bytes,1,opt,name=last_sync,json=lastSync,proto3" json:"last_sync,omitempty"`
	Services             map[string]*ServiceInsight_DataplaneStat `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *ServiceInsight) Reset()         { *m = ServiceInsight{} }
func (m *ServiceInsight) String() string { return proto.CompactTextString(m) }
func (*ServiceInsight) ProtoMessage()    {}
func (*ServiceInsight) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6edf9b69ef9de99, []int{0}
}

func (m *ServiceInsight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceInsight.Unmarshal(m, b)
}
func (m *ServiceInsight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceInsight.Marshal(b, m, deterministic)
}
func (m *ServiceInsight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceInsight.Merge(m, src)
}
func (m *ServiceInsight) XXX_Size() int {
	return xxx_messageInfo_ServiceInsight.Size(m)
}
func (m *ServiceInsight) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceInsight.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceInsight proto.InternalMessageInfo

func (m *ServiceInsight) GetLastSync() *timestamp.Timestamp {
	if m != nil {
		return m.LastSync
	}
	return nil
}

func (m *ServiceInsight) GetServices() map[string]*ServiceInsight_DataplaneStat {
	if m != nil {
		return m.Services
	}
	return nil
}

type ServiceInsight_DataplaneStat struct {
	Total                uint32   `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Online               uint32   `protobuf:"varint,2,opt,name=online,proto3" json:"online,omitempty"`
	Offline              uint32   `protobuf:"varint,3,opt,name=offline,proto3" json:"offline,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceInsight_DataplaneStat) Reset()         { *m = ServiceInsight_DataplaneStat{} }
func (m *ServiceInsight_DataplaneStat) String() string { return proto.CompactTextString(m) }
func (*ServiceInsight_DataplaneStat) ProtoMessage()    {}
func (*ServiceInsight_DataplaneStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6edf9b69ef9de99, []int{0, 0}
}

func (m *ServiceInsight_DataplaneStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceInsight_DataplaneStat.Unmarshal(m, b)
}
func (m *ServiceInsight_DataplaneStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceInsight_DataplaneStat.Marshal(b, m, deterministic)
}
func (m *ServiceInsight_DataplaneStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceInsight_DataplaneStat.Merge(m, src)
}
func (m *ServiceInsight_DataplaneStat) XXX_Size() int {
	return xxx_messageInfo_ServiceInsight_DataplaneStat.Size(m)
}
func (m *ServiceInsight_DataplaneStat) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceInsight_DataplaneStat.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceInsight_DataplaneStat proto.InternalMessageInfo

func (m *ServiceInsight_DataplaneStat) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ServiceInsight_DataplaneStat) GetOnline() uint32 {
	if m != nil {
		return m.Online
	}
	return 0
}

func (m *ServiceInsight_DataplaneStat) GetOffline() uint32 {
	if m != nil {
		return m.Offline
	}
	return 0
}

func init() {
	proto.RegisterType((*ServiceInsight)(nil), "kuma.mesh.v1alpha1.ServiceInsight")
	proto.RegisterMapType((map[string]*ServiceInsight_DataplaneStat)(nil), "kuma.mesh.v1alpha1.ServiceInsight.ServicesEntry")
	proto.RegisterType((*ServiceInsight_DataplaneStat)(nil), "kuma.mesh.v1alpha1.ServiceInsight.DataplaneStat")
}

func init() {
	proto.RegisterFile("mesh/v1alpha1/service_insight.proto", fileDescriptor_a6edf9b69ef9de99)
}

var fileDescriptor_a6edf9b69ef9de99 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x69, 0xcb, 0xe6, 0x96, 0x51, 0x91, 0x20, 0x52, 0x7a, 0x71, 0xe8, 0x65, 0xa7, 0xd4,
	0xcd, 0x83, 0xe2, 0x59, 0x05, 0xc1, 0x53, 0x2b, 0x78, 0x1c, 0xef, 0x4a, 0xfa, 0x87, 0xa5, 0x49,
	0x69, 0xde, 0x16, 0xfa, 0x69, 0xfd, 0x2a, 0xd2, 0x64, 0x15, 0x8a, 0x97, 0xdd, 0xf2, 0xe4, 0x7d,
	0xde, 0xdf, 0xf3, 0x24, 0xe4, 0xbe, 0xe2, 0xba, 0x88, 0xba, 0x2d, 0x88, 0xba, 0x80, 0x6d, 0xa4,
	0x79, 0xd3, 0x95, 0x29, 0xdf, 0x97, 0x52, 0x97, 0x79, 0x81, 0xac, 0x6e, 0x14, 0x2a, 0x4a, 0x8f,
	0x6d, 0x05, 0x6c, 0x70, 0xb2, 0xd1, 0x19, 0xde, 0xe6, 0x4a, 0xe5, 0x82, 0x47, 0xc6, 0x71, 0x68,
	0xb3, 0x08, 0xcb, 0x8a, 0x6b, 0x84, 0xaa, 0xb6, 0x4b, 0x77, 0x3f, 0x2e, 0xb9, 0x4c, 0x2c, 0xee,
	0xc3, 0xd2, 0xe8, 0x13, 0x59, 0x0a, 0xd0, 0xb8, 0xd7, 0xbd, 0x4c, 0x03, 0x67, 0xed, 0x6c, 0x56,
	0xbb, 0x90, 0x59, 0x0e, 0x1b, 0x39, 0xec, 0x6b, 0xe4, 0xc4, 0x8b, 0xc1, 0x9c, 0xf4, 0x32, 0xa5,
	0x9f, 0x64, 0x71, 0x6a, 0xa6, 0x03, 0x77, 0xed, 0x6d, 0x56, 0xbb, 0x07, 0xf6, 0xbf, 0x13, 0x9b,
	0xc6, 0x8d, 0x52, 0xbf, 0x49, 0x6c, 0xfa, 0xf8, 0x8f, 0x10, 0x7e, 0x13, 0xff, 0x15, 0x10, 0x6a,
	0x01, 0x92, 0x27, 0x08, 0x48, 0xaf, 0xc9, 0x0c, 0x15, 0x82, 0x30, 0x9d, 0xfc, 0xd8, 0x0a, 0x7a,
	0x43, 0xe6, 0x4a, 0x8a, 0x52, 0xf2, 0xc0, 0x35, 0xd7, 0x27, 0x45, 0x03, 0x72, 0xa1, 0xb2, 0xcc,
	0x0c, 0x3c, 0x33, 0x18, 0x65, 0x58, 0x11, 0x7f, 0x92, 0x49, 0xaf, 0x88, 0x77, 0xe4, 0xbd, 0xc1,
	0x2e, 0xe3, 0xe1, 0x48, 0xdf, 0xc9, 0xac, 0x03, 0xd1, 0x5a, 0xe6, 0x79, 0xcf, 0x98, 0x74, 0x8d,
	0xed, 0xfa, 0x8b, 0xfb, 0xec, 0x1c, 0xe6, 0xe6, 0xcf, 0x1e, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x28, 0xe4, 0x48, 0x54, 0xc4, 0x01, 0x00, 0x00,
}
