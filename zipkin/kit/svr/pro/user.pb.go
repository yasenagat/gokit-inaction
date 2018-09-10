// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LoginReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Pwd                  string   `protobuf:"bytes,2,opt,name=pwd,proto3" json:"pwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_8634bc355a54d72e, []int{0}
}
func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (dst *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(dst, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginReq) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type LoginResBody struct {
	Userid               string   `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`
	UnreadCount          int64    `protobuf:"varint,2,opt,name=unread_count,json=unreadCount,proto3" json:"unread_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResBody) Reset()         { *m = LoginResBody{} }
func (m *LoginResBody) String() string { return proto.CompactTextString(m) }
func (*LoginResBody) ProtoMessage()    {}
func (*LoginResBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_8634bc355a54d72e, []int{1}
}
func (m *LoginResBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResBody.Unmarshal(m, b)
}
func (m *LoginResBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResBody.Marshal(b, m, deterministic)
}
func (dst *LoginResBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResBody.Merge(dst, src)
}
func (m *LoginResBody) XXX_Size() int {
	return xxx_messageInfo_LoginResBody.Size(m)
}
func (m *LoginResBody) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResBody.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResBody proto.InternalMessageInfo

func (m *LoginResBody) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

func (m *LoginResBody) GetUnreadCount() int64 {
	if m != nil {
		return m.UnreadCount
	}
	return 0
}

type LoginRes struct {
	Code                 int64         `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string        `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Body                 *LoginResBody `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LoginRes) Reset()         { *m = LoginRes{} }
func (m *LoginRes) String() string { return proto.CompactTextString(m) }
func (*LoginRes) ProtoMessage()    {}
func (*LoginRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_8634bc355a54d72e, []int{2}
}
func (m *LoginRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRes.Unmarshal(m, b)
}
func (m *LoginRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRes.Marshal(b, m, deterministic)
}
func (dst *LoginRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRes.Merge(dst, src)
}
func (m *LoginRes) XXX_Size() int {
	return xxx_messageInfo_LoginRes.Size(m)
}
func (m *LoginRes) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRes.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRes proto.InternalMessageInfo

func (m *LoginRes) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *LoginRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *LoginRes) GetBody() *LoginResBody {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginReq)(nil), "pb.LoginReq")
	proto.RegisterType((*LoginResBody)(nil), "pb.LoginResBody")
	proto.RegisterType((*LoginRes)(nil), "pb.LoginRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error) {
	out := new(LoginRes)
	err := c.cc.Invoke(ctx, "/pb.User/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	Login(context.Context, *LoginReq) (*LoginRes, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_8634bc355a54d72e) }

var fileDescriptor_user_8634bc355a54d72e = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x50, 0x3d, 0x4f, 0x85, 0x30,
	0x14, 0x95, 0x57, 0x7c, 0x79, 0xde, 0xc7, 0x40, 0xee, 0x60, 0x08, 0x13, 0x36, 0x9a, 0xb0, 0xc8,
	0x80, 0x8b, 0xb3, 0x4e, 0x26, 0x4e, 0x4d, 0x74, 0x35, 0x94, 0x36, 0x84, 0x81, 0xb6, 0xb4, 0x10,
	0xc3, 0xbf, 0x37, 0x2d, 0x48, 0xd8, 0xee, 0xf9, 0xea, 0x39, 0x29, 0xc0, 0xec, 0xa4, 0xad, 0x8c,
	0xd5, 0x93, 0xc6, 0x93, 0xe1, 0xf4, 0x15, 0x2e, 0x9f, 0xba, 0xeb, 0x15, 0x93, 0x23, 0xe6, 0x70,
	0xf1, 0xaa, 0x6a, 0x06, 0x99, 0x45, 0x45, 0x54, 0xde, 0xb1, 0x1d, 0x63, 0x0a, 0xc4, 0xfc, 0x8a,
	0xec, 0x14, 0x68, 0x7f, 0xd2, 0x0f, 0x48, 0xb6, 0xa4, 0x7b, 0xd3, 0x62, 0xc1, 0x7b, 0x38, 0x7b,
	0x77, 0x2f, 0xb6, 0xec, 0x86, 0xf0, 0x01, 0x92, 0x59, 0x59, 0xd9, 0x88, 0x9f, 0x56, 0xcf, 0x6a,
	0x0a, 0x4f, 0x10, 0x76, 0x5d, 0xb9, 0x77, 0x4f, 0xd1, 0xef, 0x7d, 0x84, 0x43, 0x84, 0xb8, 0xd5,
	0x62, 0x1d, 0x40, 0x58, 0xb8, 0x7d, 0xf9, 0xe0, 0xba, 0xff, 0xf2, 0xc1, 0x75, 0xf8, 0x08, 0x31,
	0xd7, 0x62, 0xc9, 0x48, 0x11, 0x95, 0xd7, 0x3a, 0xad, 0x0c, 0xaf, 0x8e, 0x63, 0x58, 0x50, 0xeb,
	0x67, 0x88, 0xbf, 0x9c, 0xb4, 0xf8, 0x04, 0xb7, 0x41, 0xc5, 0xe4, 0x60, 0x1c, 0xf3, 0x23, 0x72,
	0xf4, 0x86, 0x9f, 0xc3, 0xb7, 0xbc, 0xfc, 0x05, 0x00, 0x00, 0xff, 0xff, 0xde, 0xff, 0xe2, 0x15,
	0x24, 0x01, 0x00, 0x00,
}