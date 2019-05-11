// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package go_micro_srv_dataservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateUserReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserReq) Reset()         { *m = CreateUserReq{} }
func (m *CreateUserReq) String() string { return proto.CompactTextString(m) }
func (*CreateUserReq) ProtoMessage()    {}
func (*CreateUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_c77c0ae8db71ee0e, []int{0}
}
func (m *CreateUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserReq.Unmarshal(m, b)
}
func (m *CreateUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserReq.Marshal(b, m, deterministic)
}
func (dst *CreateUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserReq.Merge(dst, src)
}
func (m *CreateUserReq) XXX_Size() int {
	return xxx_messageInfo_CreateUserReq.Size(m)
}
func (m *CreateUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserReq proto.InternalMessageInfo

func (m *CreateUserReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateUserResp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResp) Reset()         { *m = CreateUserResp{} }
func (m *CreateUserResp) String() string { return proto.CompactTextString(m) }
func (*CreateUserResp) ProtoMessage()    {}
func (*CreateUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_c77c0ae8db71ee0e, []int{1}
}
func (m *CreateUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResp.Unmarshal(m, b)
}
func (m *CreateUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResp.Marshal(b, m, deterministic)
}
func (dst *CreateUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResp.Merge(dst, src)
}
func (m *CreateUserResp) XXX_Size() int {
	return xxx_messageInfo_CreateUserResp.Size(m)
}
func (m *CreateUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResp proto.InternalMessageInfo

func (m *CreateUserResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateUserReq)(nil), "go.micro.srv.dataservice.CreateUserReq")
	proto.RegisterType((*CreateUserResp)(nil), "go.micro.srv.dataservice.CreateUserResp")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_c77c0ae8db71ee0e) }

var fileDescriptor_user_c77c0ae8db71ee0e = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x48, 0xcf, 0xd7, 0xcb, 0xcd, 0x4c, 0x2e, 0xca,
	0xd7, 0x2b, 0x2e, 0x2a, 0xd3, 0x4b, 0x49, 0x2c, 0x49, 0x2c, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e,
	0x55, 0x52, 0xe6, 0xe2, 0x75, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x0d, 0x2d, 0x4e, 0x2d, 0x0a, 0x4a,
	0x2d, 0x14, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x02, 0xb3, 0x95, 0x94, 0xb8, 0xf8, 0x90, 0x15, 0x15, 0x17, 0x08, 0x09, 0x70, 0x31, 0xe7, 0x16,
	0xa7, 0x43, 0x15, 0x81, 0x98, 0x46, 0x99, 0x5c, 0x2c, 0x20, 0x59, 0xa1, 0x44, 0x2e, 0x2e, 0x84,
	0x5a, 0x21, 0x75, 0x3d, 0x5c, 0x36, 0xeb, 0xa1, 0x58, 0x2b, 0xa5, 0x41, 0x9c, 0xc2, 0xe2, 0x02,
	0x25, 0x86, 0x24, 0x36, 0xb0, 0xa7, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x15, 0x19, 0x1a,
	0xf8, 0xe2, 0x00, 0x00, 0x00,
}