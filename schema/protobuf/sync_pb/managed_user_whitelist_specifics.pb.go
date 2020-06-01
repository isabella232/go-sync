// Code generated by protoc-gen-go. DO NOT EDIT.
// source: managed_user_whitelist_specifics.proto

package sync_pb

import (
	fmt "fmt"
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

// Properties of supervised user whitelist sync objects.
// The fields here are a subset of the fields in an ExtensionSpecifics.
type ManagedUserWhitelistSpecifics struct {
	// Globally unique id for this whitelist that identifies it in the Web Store.
	Id *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// The name of the whitelist.
	Name                 *string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManagedUserWhitelistSpecifics) Reset()         { *m = ManagedUserWhitelistSpecifics{} }
func (m *ManagedUserWhitelistSpecifics) String() string { return proto.CompactTextString(m) }
func (*ManagedUserWhitelistSpecifics) ProtoMessage()    {}
func (*ManagedUserWhitelistSpecifics) Descriptor() ([]byte, []int) {
	return fileDescriptor_82bdc4d5f95e4e14, []int{0}
}

func (m *ManagedUserWhitelistSpecifics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManagedUserWhitelistSpecifics.Unmarshal(m, b)
}
func (m *ManagedUserWhitelistSpecifics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManagedUserWhitelistSpecifics.Marshal(b, m, deterministic)
}
func (m *ManagedUserWhitelistSpecifics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManagedUserWhitelistSpecifics.Merge(m, src)
}
func (m *ManagedUserWhitelistSpecifics) XXX_Size() int {
	return xxx_messageInfo_ManagedUserWhitelistSpecifics.Size(m)
}
func (m *ManagedUserWhitelistSpecifics) XXX_DiscardUnknown() {
	xxx_messageInfo_ManagedUserWhitelistSpecifics.DiscardUnknown(m)
}

var xxx_messageInfo_ManagedUserWhitelistSpecifics proto.InternalMessageInfo

func (m *ManagedUserWhitelistSpecifics) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *ManagedUserWhitelistSpecifics) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ManagedUserWhitelistSpecifics)(nil), "sync_pb.ManagedUserWhitelistSpecifics")
}

func init() {
	proto.RegisterFile("managed_user_whitelist_specifics.proto", fileDescriptor_82bdc4d5f95e4e14)
}

var fileDescriptor_82bdc4d5f95e4e14 = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcb, 0x4d, 0xcc, 0x4b,
	0x4c, 0x4f, 0x4d, 0x89, 0x2f, 0x2d, 0x4e, 0x2d, 0x8a, 0x2f, 0xcf, 0xc8, 0x2c, 0x49, 0xcd, 0xc9,
	0x2c, 0x2e, 0x89, 0x2f, 0x2e, 0x48, 0x4d, 0xce, 0x4c, 0xcb, 0x4c, 0x2e, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x2f, 0xae, 0xcc, 0x4b, 0x8e, 0x2f, 0x48, 0x52, 0x72, 0xe6, 0x92, 0xf5,
	0x85, 0x68, 0x09, 0x2d, 0x4e, 0x2d, 0x0a, 0x87, 0x69, 0x08, 0x86, 0xa9, 0x17, 0xe2, 0xe3, 0x62,
	0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe2, 0x62,
	0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x02, 0x8b, 0x80, 0xd9, 0x4e, 0xda, 0x5c, 0xaa, 0xf9, 0x45,
	0xe9, 0x7a, 0xc9, 0x19, 0x45, 0xf9, 0xb9, 0x99, 0xa5, 0xb9, 0x7a, 0xc9, 0xf9, 0xb9, 0x05, 0xf9,
	0x79, 0xa9, 0x79, 0x25, 0xc5, 0x7a, 0x20, 0x7b, 0x20, 0x76, 0x26, 0xe7, 0xe7, 0x78, 0x30, 0x07,
	0x30, 0x02, 0x02, 0x00, 0x00, 0xff, 0xff, 0xea, 0x27, 0x0e, 0x19, 0xa3, 0x00, 0x00, 0x00,
}