// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: pkg/plugins/policies/donothingpolicy/api/v1alpha1/donothingpolicy.proto

package v1alpha1

import (
	v1alpha1 "github.com/kumahq/kuma/api/common/v1alpha1"
	_ "github.com/kumahq/kuma/api/mesh"
	_ "github.com/kumahq/protoc-gen-kumadoc/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// DoNothingPolicy a policy that does nothing
type DoNothingPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TargetRef is a reference to the resource the policy takes an effect on.
	// The resource could be either a real store object or virtual resource
	// defined inplace.
	TargetRef *v1alpha1.TargetRef     `protobuf:"bytes,1,opt,name=targetRef,proto3" json:"targetRef,omitempty"`
	// +optional
	// +nullable
	To        []*DoNothingPolicy_To   `protobuf:"bytes,2,rep,name=to,proto3" json:"to"`
	// +optional
	// +nullable
	From      []*DoNothingPolicy_From `protobuf:"bytes,3,rep,name=from,proto3" json:"from"`
}

func (x *DoNothingPolicy) Reset() {
	*x = DoNothingPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoNothingPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoNothingPolicy) ProtoMessage() {}

func (x *DoNothingPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoNothingPolicy.ProtoReflect.Descriptor instead.
func (*DoNothingPolicy) Descriptor() ([]byte, []int) {
	return file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_rawDescGZIP(), []int{0}
}

func (x *DoNothingPolicy) GetTargetRef() *v1alpha1.TargetRef {
	if x != nil {
		return x.TargetRef
	}
	return nil
}

func (x *DoNothingPolicy) GetTo() []*DoNothingPolicy_To {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *DoNothingPolicy) GetFrom() []*DoNothingPolicy_From {
	if x != nil {
		return x.From
	}
	return nil
}

type DoNothingPolicy_Conf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User defined fields
	// Set true in case of doing nothing
	EnableDoNothing bool `protobuf:"varint,1,opt,name=enableDoNothing,proto3" json:"enableDoNothing,omitempty"`
}

func (x *DoNothingPolicy_Conf) Reset() {
	*x = DoNothingPolicy_Conf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoNothingPolicy_Conf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoNothingPolicy_Conf) ProtoMessage() {}

func (x *DoNothingPolicy_Conf) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoNothingPolicy_Conf.ProtoReflect.Descriptor instead.
func (*DoNothingPolicy_Conf) Descriptor() ([]byte, []int) {
	return file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DoNothingPolicy_Conf) GetEnableDoNothing() bool {
	if x != nil {
		return x.EnableDoNothing
	}
	return false
}

type DoNothingPolicy_To struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TargetRef is a reference to the resource that represents a group of
	// destinations.
	TargetRef *v1alpha1.TargetRef `protobuf:"bytes,1,opt,name=targetRef,proto3" json:"targetRef,omitempty"`
	// Default is a configuration specific to the group of destinations
	// referenced in 'targetRef'
	Default *DoNothingPolicy_Conf `protobuf:"bytes,2,opt,name=default,proto3" json:"default,omitempty"`
}

func (x *DoNothingPolicy_To) Reset() {
	*x = DoNothingPolicy_To{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoNothingPolicy_To) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoNothingPolicy_To) ProtoMessage() {}

func (x *DoNothingPolicy_To) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoNothingPolicy_To.ProtoReflect.Descriptor instead.
func (*DoNothingPolicy_To) Descriptor() ([]byte, []int) {
	return file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_rawDescGZIP(), []int{0, 1}
}

func (x *DoNothingPolicy_To) GetTargetRef() *v1alpha1.TargetRef {
	if x != nil {
		return x.TargetRef
	}
	return nil
}

func (x *DoNothingPolicy_To) GetDefault() *DoNothingPolicy_Conf {
	if x != nil {
		return x.Default
	}
	return nil
}

type DoNothingPolicy_From struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TargetRef is a reference to the resource that represents a group of
	// clients.
	TargetRef *v1alpha1.TargetRef `protobuf:"bytes,1,opt,name=targetRef,proto3" json:"targetRef,omitempty"`
	// Default is a configuration specific to the group of clients referenced in
	// 'targetRef'
	Default *DoNothingPolicy_Conf `protobuf:"bytes,2,opt,name=default,proto3" json:"default,omitempty"`
}

func (x *DoNothingPolicy_From) Reset() {
	*x = DoNothingPolicy_From{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoNothingPolicy_From) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoNothingPolicy_From) ProtoMessage() {}

func (x *DoNothingPolicy_From) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoNothingPolicy_From.ProtoReflect.Descriptor instead.
func (*DoNothingPolicy_From) Descriptor() ([]byte, []int) {
	return file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_rawDescGZIP(), []int{0, 2}
}

func (x *DoNothingPolicy_From) GetTargetRef() *v1alpha1.TargetRef {
	if x != nil {
		return x.TargetRef
	}
	return nil
}

func (x *DoNothingPolicy_From) GetDefault() *DoNothingPolicy_Conf {
	if x != nil {
		return x.Default
	}
	return nil
}

var File_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto protoreflect.FileDescriptor

var file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_rawDesc = []byte{
	0x0a, 0x47, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x64, 0x6f, 0x6e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x64, 0x6f, 0x6e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x2e, 0x64, 0x6f, 0x6e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x12, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x72, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x6b, 0x75, 0x6d, 0x61, 0x2d, 0x64, 0x6f, 0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x05, 0x0a, 0x0f, 0x44, 0x6f, 0x4e, 0x6f, 0x74, 0x68,
	0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x3d, 0x0a, 0x09, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x52, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6b,
	0x75, 0x6d, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x66, 0x52, 0x09, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x66, 0x12, 0x52, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x64, 0x6f, 0x6e,
	0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x6f, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x54, 0x6f, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x58, 0x0a, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x6b, 0x75, 0x6d,
	0x61, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x2e, 0x64, 0x6f, 0x6e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x6f, 0x4e, 0x6f,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x46, 0x72, 0x6f, 0x6d,
	0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x1a, 0x30, 0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x28,
	0x0a, 0x0f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x6f, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x44,
	0x6f, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x1a, 0xaf, 0x01, 0x0a, 0x02, 0x54, 0x6f, 0x12,
	0x43, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x52, 0x65, 0x66, 0x42, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x52, 0x65, 0x66, 0x12, 0x64, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x64, 0x6f,
	0x6e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x6f, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x42, 0x04, 0x88, 0xb5, 0x18,
	0x01, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x1a, 0xb1, 0x01, 0x0a, 0x04, 0x46,
	0x72, 0x6f, 0x6d, 0x12, 0x43, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x66,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x66, 0x42, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x52, 0x09, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x66, 0x12, 0x64, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x6b, 0x75, 0x6d, 0x61,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x2e, 0x64, 0x6f, 0x6e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x6f, 0x4e, 0x6f, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x42,
	0x04, 0x88, 0xb5, 0x18, 0x01, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x3a, 0x08,
	0xb2, 0x8c, 0x89, 0xa6, 0x01, 0x02, 0x08, 0x01, 0x42, 0x74, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x6d, 0x61, 0x68, 0x71, 0x2f, 0x6b, 0x75,
	0x6d, 0x61, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x64, 0x6f, 0x6e, 0x6f, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x8a, 0xb5, 0x18, 0x26, 0x50, 0x01, 0xa2, 0x01, 0x0f, 0x44, 0x6f, 0x4e,
	0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0xf2, 0x01, 0x0f, 0x64,
	0x6f, 0x6e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_rawDescOnce sync.Once
	file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_rawDescData = file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_rawDesc
)

func file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_rawDescGZIP() []byte {
	file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_rawDescOnce.Do(func() {
		file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_rawDescData)
	})
	return file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_rawDescData
}

var file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_goTypes = []interface{}{
	(*DoNothingPolicy)(nil),      // 0: kuma.plugins.policies.donothingpolicy.v1alpha1.DoNothingPolicy
	(*DoNothingPolicy_Conf)(nil), // 1: kuma.plugins.policies.donothingpolicy.v1alpha1.DoNothingPolicy.Conf
	(*DoNothingPolicy_To)(nil),   // 2: kuma.plugins.policies.donothingpolicy.v1alpha1.DoNothingPolicy.To
	(*DoNothingPolicy_From)(nil), // 3: kuma.plugins.policies.donothingpolicy.v1alpha1.DoNothingPolicy.From
	(*v1alpha1.TargetRef)(nil),   // 4: kuma.common.v1alpha1.TargetRef
}
var file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_depIdxs = []int32{
	4, // 0: kuma.plugins.policies.donothingpolicy.v1alpha1.DoNothingPolicy.targetRef:type_name -> kuma.common.v1alpha1.TargetRef
	2, // 1: kuma.plugins.policies.donothingpolicy.v1alpha1.DoNothingPolicy.to:type_name -> kuma.plugins.policies.donothingpolicy.v1alpha1.DoNothingPolicy.To
	3, // 2: kuma.plugins.policies.donothingpolicy.v1alpha1.DoNothingPolicy.from:type_name -> kuma.plugins.policies.donothingpolicy.v1alpha1.DoNothingPolicy.From
	4, // 3: kuma.plugins.policies.donothingpolicy.v1alpha1.DoNothingPolicy.To.targetRef:type_name -> kuma.common.v1alpha1.TargetRef
	1, // 4: kuma.plugins.policies.donothingpolicy.v1alpha1.DoNothingPolicy.To.default:type_name -> kuma.plugins.policies.donothingpolicy.v1alpha1.DoNothingPolicy.Conf
	4, // 5: kuma.plugins.policies.donothingpolicy.v1alpha1.DoNothingPolicy.From.targetRef:type_name -> kuma.common.v1alpha1.TargetRef
	1, // 6: kuma.plugins.policies.donothingpolicy.v1alpha1.DoNothingPolicy.From.default:type_name -> kuma.plugins.policies.donothingpolicy.v1alpha1.DoNothingPolicy.Conf
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_init() }
func file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_init() {
	if File_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoNothingPolicy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoNothingPolicy_Conf); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoNothingPolicy_To); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoNothingPolicy_From); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_goTypes,
		DependencyIndexes: file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_depIdxs,
		MessageInfos:      file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_msgTypes,
	}.Build()
	File_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto = out.File
	file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_rawDesc = nil
	file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_goTypes = nil
	file_pkg_plugins_policies_donothingpolicy_api_v1alpha1_donothingpolicy_proto_depIdxs = nil
}
