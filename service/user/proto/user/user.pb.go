// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PrettyId             string   `protobuf:"bytes,2,opt,name=pretty_id,json=prettyId,proto3" json:"pretty_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	CreateTime           int64    `protobuf:"varint,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           int64    `protobuf:"varint,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	DeleteTime           int64    `protobuf:"varint,7,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetPrettyId() string {
	if m != nil {
		return m.PrettyId
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *User) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *User) GetDeleteTime() int64 {
	if m != nil {
		return m.DeleteTime
	}
	return 0
}

type Request struct {
	PageNum              int64    `protobuf:"varint,1,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	PageSize             int64    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetPageNum() int64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *Request) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*Request)(nil), "user.Request")
	proto.RegisterType((*Response)(nil), "user.Response")
	proto.RegisterType((*Token)(nil), "user.Token")
	proto.RegisterType((*Error)(nil), "user.Error")
}

func init() {
	proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf)
}

var fileDescriptor_116e343673f7ffaf = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6a, 0x14, 0x41,
	0x10, 0xc6, 0xb3, 0xf3, 0x3f, 0x35, 0x24, 0x87, 0xc2, 0x43, 0x1b, 0x41, 0xc7, 0x09, 0x48, 0x24,
	0x90, 0x43, 0xc4, 0x07, 0x58, 0x54, 0x82, 0x17, 0x0f, 0x9d, 0x28, 0xde, 0x96, 0x71, 0xbb, 0x88,
	0x8d, 0xf3, 0x2f, 0xdd, 0x3d, 0x11, 0x73, 0xf4, 0xe1, 0x7c, 0x2e, 0xe9, 0xea, 0x19, 0xc9, 0x4a,
	0x40, 0x4f, 0x5b, 0xf5, 0xd5, 0x6f, 0x8b, 0xa9, 0xef, 0x9b, 0x01, 0x98, 0x2c, 0x99, 0xb3, 0xd1,
	0x0c, 0x6e, 0xc0, 0xc4, 0xd7, 0xf5, 0xaf, 0x15, 0x24, 0x1f, 0x2d, 0x19, 0x3c, 0x84, 0x48, 0x2b,
	0xb1, 0xaa, 0x56, 0x27, 0xfb, 0x32, 0xd2, 0x0a, 0x9f, 0xc0, 0xfe, 0x68, 0xc8, 0xb9, 0x1f, 0x1b,
	0xad, 0x44, 0xc4, 0x72, 0x11, 0x84, 0xf7, 0x0a, 0x11, 0x92, 0xbe, 0xe9, 0x48, 0xc4, 0xac, 0x73,
	0x8d, 0x47, 0x50, 0x8c, 0x8d, 0xb5, 0xdf, 0x07, 0xa3, 0x44, 0x32, 0xf3, 0x73, 0x8f, 0xcf, 0xa0,
	0xdc, 0x1a, 0x6a, 0x1c, 0x6d, 0x9c, 0xee, 0x48, 0xa4, 0xd5, 0xea, 0x24, 0x96, 0x10, 0xa4, 0x2b,
	0xdd, 0x91, 0x07, 0xa6, 0x51, 0xfd, 0x01, 0xb2, 0x00, 0x04, 0x69, 0x01, 0x14, 0xb5, 0xb4, 0x00,
	0x79, 0x00, 0x82, 0xe4, 0x81, 0x7a, 0x0d, 0xb9, 0xa4, 0x9b, 0x89, 0xac, 0xc3, 0xc7, 0xfe, 0x49,
	0xae, 0x69, 0xd3, 0x4f, 0x1d, 0x1f, 0x14, 0xcb, 0xdc, 0xf7, 0x1f, 0xa6, 0x8e, 0xaf, 0xf2, 0x23,
	0xab, 0xef, 0x88, 0xaf, 0x8a, 0x25, 0xb3, 0x97, 0xfa, 0x8e, 0xea, 0x1b, 0x28, 0x24, 0xd9, 0x71,
	0xe8, 0x2d, 0xe1, 0x53, 0x60, 0x7f, 0xf8, 0xff, 0xe5, 0x39, 0x9c, 0xb1, 0x71, 0xde, 0x28, 0xc9,
	0x3a, 0x56, 0x90, 0xfa, 0x5f, 0x2b, 0xa2, 0x2a, 0xfe, 0x0b, 0x08, 0x03, 0x3c, 0x86, 0x8c, 0x8c,
	0x19, 0x8c, 0x15, 0x31, 0x23, 0x65, 0x40, 0xde, 0x79, 0x4d, 0xce, 0xa3, 0xfa, 0x33, 0xa4, 0x57,
	0xc3, 0x37, 0xea, 0xf1, 0x11, 0xa4, 0xce, 0x17, 0x73, 0x02, 0xa1, 0xf1, 0xea, 0x6d, 0xd3, 0xce,
	0x01, 0x14, 0x32, 0x34, 0xff, 0xb7, 0xf9, 0x35, 0xa4, 0x2c, 0xf8, 0xac, 0xb6, 0x83, 0x22, 0x5e,
	0x9c, 0x4a, 0xae, 0x51, 0x40, 0xde, 0x91, 0xb5, 0xcd, 0x35, 0xcd, 0xd1, 0x2e, 0xed, 0xf9, 0xcf,
	0x08, 0x4a, 0x7f, 0xc5, 0x25, 0x99, 0x5b, 0xbd, 0x25, 0x7c, 0x01, 0xd9, 0x1b, 0x8e, 0x09, 0xef,
	0x9d, 0x78, 0x74, 0x18, 0xea, 0xc5, 0xad, 0x7a, 0x0f, 0x8f, 0x21, 0xbe, 0x20, 0xf7, 0x0f, 0xe8,
	0x25, 0x64, 0x17, 0xe4, 0xd6, 0x6d, 0x8b, 0x07, 0xcb, 0x8c, 0x13, 0x7b, 0x00, 0x7d, 0x0e, 0xc9,
	0x7a, 0x72, 0x5f, 0x77, 0x16, 0xce, 0x77, 0xb2, 0x61, 0xf5, 0x1e, 0x9e, 0x02, 0xbc, 0xe5, 0xfc,
	0x19, 0xbc, 0x3f, 0x7c, 0x60, 0xdf, 0x29, 0x1c, 0x7c, 0xf2, 0xe6, 0xf9, 0xf7, 0x89, 0xad, 0xdd,
	0xe1, 0x77, 0x37, 0x7f, 0xc9, 0xf8, 0x0b, 0x79, 0xf5, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xa8,
	0x57, 0x6d, 0x2f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	Get(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Auth(ctx context.Context, in *User, opts ...grpc.CallOption) (*Token, error)
	DeleteAuth(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Response, error)
	ValidateToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/user.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/user.UserService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAll(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/user.UserService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Auth(ctx context.Context, in *User, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/user.UserService/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteAuth(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/user.UserService/DeleteAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ValidateToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/user.UserService/ValidateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Create(context.Context, *User) (*Response, error)
	Get(context.Context, *User) (*Response, error)
	GetAll(context.Context, *Request) (*Response, error)
	Auth(context.Context, *User) (*Token, error)
	DeleteAuth(context.Context, *Token) (*Response, error)
	ValidateToken(context.Context, *Token) (*Token, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Create(ctx context.Context, req *User) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedUserServiceServer) Get(ctx context.Context, req *User) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedUserServiceServer) GetAll(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedUserServiceServer) Auth(ctx context.Context, req *User) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (*UnimplementedUserServiceServer) DeleteAuth(ctx context.Context, req *Token) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuth not implemented")
}
func (*UnimplementedUserServiceServer) ValidateToken(ctx context.Context, req *Token) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAll(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Auth(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/DeleteAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteAuth(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ValidateToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _UserService_GetAll_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _UserService_Auth_Handler,
		},
		{
			MethodName: "DeleteAuth",
			Handler:    _UserService_DeleteAuth_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _UserService_ValidateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
