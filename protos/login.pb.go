// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/login.proto

package login

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SignUpRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpRequest) Reset()         { *m = SignUpRequest{} }
func (m *SignUpRequest) String() string { return proto.CompactTextString(m) }
func (*SignUpRequest) ProtoMessage()    {}
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_40fa32f243b417c0, []int{0}
}

func (m *SignUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpRequest.Unmarshal(m, b)
}
func (m *SignUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpRequest.Marshal(b, m, deterministic)
}
func (m *SignUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpRequest.Merge(m, src)
}
func (m *SignUpRequest) XXX_Size() int {
	return xxx_messageInfo_SignUpRequest.Size(m)
}
func (m *SignUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpRequest proto.InternalMessageInfo

func (m *SignUpRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignUpRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SignUpResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpResponse) Reset()         { *m = SignUpResponse{} }
func (m *SignUpResponse) String() string { return proto.CompactTextString(m) }
func (*SignUpResponse) ProtoMessage()    {}
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_40fa32f243b417c0, []int{1}
}

func (m *SignUpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpResponse.Unmarshal(m, b)
}
func (m *SignUpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpResponse.Marshal(b, m, deterministic)
}
func (m *SignUpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpResponse.Merge(m, src)
}
func (m *SignUpResponse) XXX_Size() int {
	return xxx_messageInfo_SignUpResponse.Size(m)
}
func (m *SignUpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpResponse proto.InternalMessageInfo

func (m *SignUpResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*SignUpRequest)(nil), "login.SignUpRequest")
	proto.RegisterType((*SignUpResponse)(nil), "login.SignUpResponse")
}

func init() { proto.RegisterFile("protos/login.proto", fileDescriptor_40fa32f243b417c0) }

var fileDescriptor_40fa32f243b417c0 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0xcf, 0xc9, 0x4f, 0xcf, 0xcc, 0xd3, 0x03, 0x73, 0x84, 0x58, 0xc1, 0x1c, 0x29,
	0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd, 0xc4, 0x82, 0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc,
	0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0x22, 0x25, 0x47, 0x2e, 0xde, 0xe0, 0xcc, 0xf4,
	0xbc, 0xd0, 0x82, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x11, 0x2e, 0xd6, 0xd4, 0xdc,
	0xc4, 0xcc, 0x1c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x48, 0x8a, 0x8b, 0xa3,
	0x20, 0xb1, 0xb8, 0xb8, 0x3c, 0xbf, 0x28, 0x45, 0x82, 0x09, 0x2c, 0x01, 0xe7, 0x2b, 0x69, 0x71,
	0xf1, 0xc1, 0x8c, 0x28, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x92, 0xe0, 0x62, 0x2f, 0x2e, 0x4d,
	0x4e, 0x4e, 0x2d, 0x2e, 0x06, 0x9b, 0xc2, 0x11, 0x04, 0xe3, 0x1a, 0x05, 0x73, 0xb1, 0x41, 0xd4,
	0x0a, 0x79, 0xc2, 0x59, 0x22, 0x7a, 0x10, 0x57, 0xa3, 0xb8, 0x43, 0x4a, 0x14, 0x4d, 0x14, 0x62,
	0xb4, 0x92, 0x50, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0x78, 0x94, 0xd8, 0xf5, 0x8b, 0x33, 0xd3, 0xf3,
	0x4a, 0x0b, 0xac, 0x18, 0xb5, 0x92, 0xd8, 0xc0, 0x5e, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x2e, 0x31, 0x01, 0x77, 0x05, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SignUpClient is the client API for SignUp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SignUpClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
}

type signUpClient struct {
	cc *grpc.ClientConn
}

func NewSignUpClient(cc *grpc.ClientConn) SignUpClient {
	return &signUpClient{cc}
}

func (c *signUpClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, "/login.SignUp/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignUpServer is the server API for SignUp service.
type SignUpServer interface {
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
}

func RegisterSignUpServer(s *grpc.Server, srv SignUpServer) {
	s.RegisterService(&_SignUp_serviceDesc, srv)
}

func _SignUp_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignUpServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.SignUp/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignUpServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SignUp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "login.SignUp",
	HandlerType: (*SignUpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _SignUp_SignUp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/login.proto",
}
