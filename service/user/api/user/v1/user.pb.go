// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/v1/user.proto

package v1

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

type CreateUserInfo struct {
	NickName             string   `protobuf:"bytes,1,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserInfo) Reset()         { *m = CreateUserInfo{} }
func (m *CreateUserInfo) String() string { return proto.CompactTextString(m) }
func (*CreateUserInfo) ProtoMessage()    {}
func (*CreateUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eeaf71090dd90b5, []int{0}
}

func (m *CreateUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserInfo.Unmarshal(m, b)
}
func (m *CreateUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserInfo.Marshal(b, m, deterministic)
}
func (m *CreateUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserInfo.Merge(m, src)
}
func (m *CreateUserInfo) XXX_Size() int {
	return xxx_messageInfo_CreateUserInfo.Size(m)
}
func (m *CreateUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserInfo proto.InternalMessageInfo

func (m *CreateUserInfo) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *CreateUserInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateUserInfo) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

type UserInfoResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	NickName             string   `protobuf:"bytes,4,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Birthday             int64    `protobuf:"varint,5,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Gender               string   `protobuf:"bytes,6,opt,name=gender,proto3" json:"gender,omitempty"`
	Role                 int32    `protobuf:"varint,7,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoResponse) Reset()         { *m = UserInfoResponse{} }
func (m *UserInfoResponse) String() string { return proto.CompactTextString(m) }
func (*UserInfoResponse) ProtoMessage()    {}
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eeaf71090dd90b5, []int{1}
}

func (m *UserInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoResponse.Unmarshal(m, b)
}
func (m *UserInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoResponse.Marshal(b, m, deterministic)
}
func (m *UserInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoResponse.Merge(m, src)
}
func (m *UserInfoResponse) XXX_Size() int {
	return xxx_messageInfo_UserInfoResponse.Size(m)
}
func (m *UserInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoResponse proto.InternalMessageInfo

func (m *UserInfoResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfoResponse) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserInfoResponse) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *UserInfoResponse) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *UserInfoResponse) GetBirthday() int64 {
	if m != nil {
		return m.Birthday
	}
	return 0
}

func (m *UserInfoResponse) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *UserInfoResponse) GetRole() int32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateUserInfo)(nil), "api.user.v1.CreateUserInfo")
	proto.RegisterType((*UserInfoResponse)(nil), "api.user.v1.UserInfoResponse")
}

func init() { proto.RegisterFile("user/v1/user.proto", fileDescriptor_6eeaf71090dd90b5) }

var fileDescriptor_6eeaf71090dd90b5 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x18, 0xc5, 0x49, 0xd7, 0x55, 0xfd, 0x84, 0x21, 0x9f, 0x20, 0x61, 0x22, 0x8c, 0xe1, 0x61, 0xa7,
	0x96, 0xea, 0xd1, 0x9b, 0x5e, 0xf4, 0x22, 0xa3, 0xe0, 0xc5, 0x93, 0xa9, 0xf9, 0xd4, 0xe0, 0xd6,
	0x84, 0x2f, 0xb5, 0xe2, 0xbf, 0xe6, 0x5f, 0x27, 0x4d, 0xb7, 0xb9, 0xea, 0x6d, 0xa7, 0xe4, 0xc7,
	0x83, 0x97, 0x97, 0xf7, 0x00, 0x3f, 0x3c, 0x71, 0xd6, 0xe4, 0x59, 0x7b, 0xa6, 0x8e, 0x6d, 0x6d,
	0xf1, 0x50, 0x39, 0x93, 0x06, 0x6e, 0xf2, 0xe9, 0x13, 0x8c, 0x6e, 0x98, 0x54, 0x4d, 0x0f, 0x9e,
	0xf8, 0xae, 0x7a, 0xb1, 0x38, 0x86, 0xfd, 0xca, 0x3c, 0xbf, 0xdf, 0xab, 0x25, 0x49, 0x31, 0x11,
	0xb3, 0x83, 0x62, 0xc3, 0xad, 0xe6, 0x94, 0xf7, 0x9f, 0x96, 0xb5, 0x8c, 0x3a, 0x6d, 0xcd, 0x78,
	0x02, 0xc9, 0xd2, 0x96, 0x66, 0x41, 0x72, 0x10, 0x94, 0x15, 0x4d, 0xbf, 0x05, 0x1c, 0xad, 0xcd,
	0x0b, 0xf2, 0xce, 0x56, 0x9e, 0x70, 0x04, 0x91, 0xd1, 0xc1, 0x7e, 0x50, 0x44, 0x46, 0xef, 0x62,
	0xdc, 0x0b, 0x1a, 0xff, 0x0f, 0x5a, 0x1a, 0xae, 0xdf, 0xb4, 0xfa, 0x92, 0xc3, 0xf0, 0xca, 0x86,
	0x5b, 0xbf, 0x57, 0xaa, 0x34, 0xb1, 0x4c, 0x3a, 0xbf, 0x8e, 0x10, 0x21, 0x66, 0xbb, 0x20, 0xb9,
	0x37, 0x11, 0xb3, 0x61, 0x11, 0xee, 0x17, 0x73, 0x88, 0xdb, 0xec, 0x78, 0x0b, 0xf0, 0x5b, 0x13,
	0x9e, 0xa6, 0x5b, 0x15, 0xa6, 0xfd, 0xfe, 0xc6, 0x67, 0x3d, 0xf1, 0xef, 0xcf, 0xaf, 0xcf, 0x61,
	0xbb, 0xff, 0xb9, 0x78, 0x3c, 0x0e, 0x13, 0x29, 0x67, 0xb2, 0xd5, 0x56, 0x57, 0x4d, 0x5e, 0x26,
	0x61, 0xaa, 0xcb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x90, 0x82, 0xdd, 0x41, 0xc0, 0x01, 0x00,
	0x00,
}
