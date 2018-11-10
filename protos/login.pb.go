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
	FirstName            string   `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
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

func (m *SignUpRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *SignUpRequest) GetLastName() string {
	if m != nil {
		return m.LastName
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

type LogInRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogInRequest) Reset()         { *m = LogInRequest{} }
func (m *LogInRequest) String() string { return proto.CompactTextString(m) }
func (*LogInRequest) ProtoMessage()    {}
func (*LogInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_40fa32f243b417c0, []int{2}
}

func (m *LogInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogInRequest.Unmarshal(m, b)
}
func (m *LogInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogInRequest.Marshal(b, m, deterministic)
}
func (m *LogInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogInRequest.Merge(m, src)
}
func (m *LogInRequest) XXX_Size() int {
	return xxx_messageInfo_LogInRequest.Size(m)
}
func (m *LogInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogInRequest proto.InternalMessageInfo

func (m *LogInRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LogInRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LogInResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogInResponse) Reset()         { *m = LogInResponse{} }
func (m *LogInResponse) String() string { return proto.CompactTextString(m) }
func (*LogInResponse) ProtoMessage()    {}
func (*LogInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_40fa32f243b417c0, []int{3}
}

func (m *LogInResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogInResponse.Unmarshal(m, b)
}
func (m *LogInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogInResponse.Marshal(b, m, deterministic)
}
func (m *LogInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogInResponse.Merge(m, src)
}
func (m *LogInResponse) XXX_Size() int {
	return xxx_messageInfo_LogInResponse.Size(m)
}
func (m *LogInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogInResponse proto.InternalMessageInfo

func (m *LogInResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ForgotPasswordRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForgotPasswordRequest) Reset()         { *m = ForgotPasswordRequest{} }
func (m *ForgotPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*ForgotPasswordRequest) ProtoMessage()    {}
func (*ForgotPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_40fa32f243b417c0, []int{4}
}

func (m *ForgotPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForgotPasswordRequest.Unmarshal(m, b)
}
func (m *ForgotPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForgotPasswordRequest.Marshal(b, m, deterministic)
}
func (m *ForgotPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForgotPasswordRequest.Merge(m, src)
}
func (m *ForgotPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_ForgotPasswordRequest.Size(m)
}
func (m *ForgotPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ForgotPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ForgotPasswordRequest proto.InternalMessageInfo

func (m *ForgotPasswordRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type ForgotPasswordResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForgotPasswordResponse) Reset()         { *m = ForgotPasswordResponse{} }
func (m *ForgotPasswordResponse) String() string { return proto.CompactTextString(m) }
func (*ForgotPasswordResponse) ProtoMessage()    {}
func (*ForgotPasswordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_40fa32f243b417c0, []int{5}
}

func (m *ForgotPasswordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForgotPasswordResponse.Unmarshal(m, b)
}
func (m *ForgotPasswordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForgotPasswordResponse.Marshal(b, m, deterministic)
}
func (m *ForgotPasswordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForgotPasswordResponse.Merge(m, src)
}
func (m *ForgotPasswordResponse) XXX_Size() int {
	return xxx_messageInfo_ForgotPasswordResponse.Size(m)
}
func (m *ForgotPasswordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ForgotPasswordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ForgotPasswordResponse proto.InternalMessageInfo

func (m *ForgotPasswordResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*SignUpRequest)(nil), "login.SignUpRequest")
	proto.RegisterType((*SignUpResponse)(nil), "login.SignUpResponse")
	proto.RegisterType((*LogInRequest)(nil), "login.LogInRequest")
	proto.RegisterType((*LogInResponse)(nil), "login.LogInResponse")
	proto.RegisterType((*ForgotPasswordRequest)(nil), "login.ForgotPasswordRequest")
	proto.RegisterType((*ForgotPasswordResponse)(nil), "login.ForgotPasswordResponse")
}

func init() { proto.RegisterFile("protos/login.proto", fileDescriptor_40fa32f243b417c0) }

var fileDescriptor_40fa32f243b417c0 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x65, 0xab, 0xfd, 0x1a, 0xdb, 0x8a, 0xe3, 0x56, 0x96, 0xb5, 0x05, 0xc9, 0x49, 0x0b, 0x76,
	0xa1, 0xde, 0x7a, 0xd2, 0x83, 0x42, 0xa1, 0x48, 0xa9, 0x78, 0x96, 0x58, 0xd3, 0x10, 0xd8, 0x26,
	0xeb, 0x66, 0x8b, 0x27, 0x2f, 0x5e, 0xfc, 0x01, 0xfe, 0x34, 0xff, 0x82, 0x3f, 0x44, 0x36, 0xc9,
	0x8a, 0x2d, 0x6a, 0xc1, 0xe3, 0x9b, 0x37, 0xf3, 0x5e, 0x32, 0x6f, 0x00, 0x93, 0x54, 0x65, 0x4a,
	0x47, 0xb1, 0xe2, 0x42, 0xf6, 0x0d, 0xc0, 0xb2, 0x01, 0x61, 0x87, 0x2b, 0xc5, 0x63, 0x16, 0xd1,
	0x44, 0x44, 0x54, 0x4a, 0x95, 0xd1, 0x4c, 0x28, 0xa9, 0x6d, 0x13, 0x79, 0x86, 0xe6, 0x8d, 0xe0,
	0xf2, 0x36, 0x99, 0xb2, 0xc7, 0x25, 0xd3, 0x19, 0xfa, 0x50, 0x66, 0x0b, 0x2a, 0xe2, 0xc0, 0x3b,
	0xf2, 0x8e, 0xeb, 0x53, 0x0b, 0x30, 0x84, 0x5a, 0x42, 0xb5, 0x7e, 0x52, 0xe9, 0x43, 0x50, 0x32,
	0xc4, 0x17, 0xc6, 0x2e, 0xc0, 0x5c, 0xa4, 0x3a, 0xbb, 0x93, 0x74, 0xc1, 0x82, 0x2d, 0xc3, 0xd6,
	0x4d, 0xe5, 0x9a, 0x2e, 0x18, 0x1e, 0x42, 0x3d, 0xa6, 0x05, 0xbb, 0x6d, 0x67, 0xf3, 0x42, 0x4e,
	0x92, 0x1e, 0xb4, 0x0a, 0x7b, 0x9d, 0x28, 0xa9, 0x19, 0x06, 0x50, 0xd5, 0xcb, 0xd9, 0x8c, 0x69,
	0x6d, 0x5e, 0x50, 0x9b, 0x16, 0x90, 0x9c, 0x43, 0x63, 0xac, 0xf8, 0x48, 0xfe, 0xfb, 0xa5, 0xe4,
	0x04, 0x9a, 0x4e, 0x61, 0xa3, 0xd9, 0x29, 0xb4, 0xaf, 0x54, 0xca, 0x55, 0x36, 0x71, 0xc3, 0x7f,
	0xba, 0x92, 0x01, 0x1c, 0xac, 0xb7, 0x6f, 0xb2, 0x18, 0xbc, 0x96, 0xa0, 0x36, 0xce, 0x23, 0xba,
	0x98, 0x8c, 0x70, 0x04, 0x15, 0xbb, 0x08, 0xf4, 0xfb, 0x36, 0xc4, 0x95, 0x58, 0xc2, 0xf6, 0x5a,
	0xd5, 0xaa, 0x13, 0x7c, 0x79, 0xff, 0x78, 0x2b, 0x35, 0x48, 0x35, 0xd2, 0x82, 0xcb, 0x65, 0x32,
	0xf4, 0x7a, 0x78, 0x09, 0x65, 0xf3, 0x4b, 0xdc, 0x77, 0x33, 0xdf, 0xb7, 0x16, 0xfa, 0xab, 0x45,
	0xa7, 0xb3, 0x67, 0x74, 0x76, 0x48, 0xc5, 0x5e, 0x50, 0x2e, 0x23, 0xa0, 0xb5, 0xfa, 0x25, 0xec,
	0xb8, 0xd1, 0x1f, 0x17, 0x13, 0x76, 0x7f, 0x61, 0x9d, 0x43, 0x68, 0x1c, 0x7c, 0xb2, 0x1b, 0xcd,
	0x4d, 0x43, 0x11, 0xca, 0xd0, 0xeb, 0xdd, 0x57, 0xcc, 0x2d, 0x9e, 0x7d, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x04, 0xbd, 0xdb, 0x4c, 0xc6, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoginAPIClient is the client API for LoginAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoginAPIClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error)
	ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*ForgotPasswordResponse, error)
}

type loginAPIClient struct {
	cc *grpc.ClientConn
}

func NewLoginAPIClient(cc *grpc.ClientConn) LoginAPIClient {
	return &loginAPIClient{cc}
}

func (c *loginAPIClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, "/login.LoginAPI/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginAPIClient) LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error) {
	out := new(LogInResponse)
	err := c.cc.Invoke(ctx, "/login.LoginAPI/LogIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginAPIClient) ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*ForgotPasswordResponse, error) {
	out := new(ForgotPasswordResponse)
	err := c.cc.Invoke(ctx, "/login.LoginAPI/ForgotPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginAPIServer is the server API for LoginAPI service.
type LoginAPIServer interface {
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	LogIn(context.Context, *LogInRequest) (*LogInResponse, error)
	ForgotPassword(context.Context, *ForgotPasswordRequest) (*ForgotPasswordResponse, error)
}

func RegisterLoginAPIServer(s *grpc.Server, srv LoginAPIServer) {
	s.RegisterService(&_LoginAPI_serviceDesc, srv)
}

func _LoginAPI_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginAPIServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.LoginAPI/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginAPIServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginAPI_LogIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginAPIServer).LogIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.LoginAPI/LogIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginAPIServer).LogIn(ctx, req.(*LogInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginAPI_ForgotPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgotPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginAPIServer).ForgotPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.LoginAPI/ForgotPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginAPIServer).ForgotPassword(ctx, req.(*ForgotPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoginAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "login.LoginAPI",
	HandlerType: (*LoginAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _LoginAPI_SignUp_Handler,
		},
		{
			MethodName: "LogIn",
			Handler:    _LoginAPI_LogIn_Handler,
		},
		{
			MethodName: "ForgotPassword",
			Handler:    _LoginAPI_ForgotPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/login.proto",
}
