// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gprotobuf.hello.proto

package gprotobuf

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

//消息体
type Hello struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hello) Reset()         { *m = Hello{} }
func (m *Hello) String() string { return proto.CompactTextString(m) }
func (*Hello) ProtoMessage()    {}
func (*Hello) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd5d0dc595c59aee, []int{0}
}

func (m *Hello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hello.Unmarshal(m, b)
}
func (m *Hello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hello.Marshal(b, m, deterministic)
}
func (m *Hello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hello.Merge(m, src)
}
func (m *Hello) XXX_Size() int {
	return xxx_messageInfo_Hello.Size(m)
}
func (m *Hello) XXX_DiscardUnknown() {
	xxx_messageInfo_Hello.DiscardUnknown(m)
}

var xxx_messageInfo_Hello proto.InternalMessageInfo

func (m *Hello) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Hello)(nil), "gprotobuf.Hello")
}

func init() { proto.RegisterFile("gprotobuf.hello.proto", fileDescriptor_dd5d0dc595c59aee) }

var fileDescriptor_dd5d0dc595c59aee = []byte{
	// 77 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0x2f, 0x28, 0xca,
	0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0xcb, 0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x03, 0x73, 0x85, 0x38,
	0xe1, 0xc2, 0x4a, 0xb2, 0x5c, 0xac, 0x1e, 0x20, 0x19, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c,
	0xd2, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x27, 0x89, 0x0d, 0xac, 0xd0, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x80, 0x74, 0xe7, 0xb5, 0x49, 0x00, 0x00, 0x00,
}
