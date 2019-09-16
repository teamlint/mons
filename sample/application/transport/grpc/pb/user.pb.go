// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

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

type FindUserRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindUserRequest) Reset()         { *m = FindUserRequest{} }
func (m *FindUserRequest) String() string { return proto.CompactTextString(m) }
func (*FindUserRequest) ProtoMessage()    {}
func (*FindUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *FindUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindUserRequest.Unmarshal(m, b)
}
func (m *FindUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindUserRequest.Marshal(b, m, deterministic)
}
func (m *FindUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindUserRequest.Merge(m, src)
}
func (m *FindUserRequest) XXX_Size() int {
	return xxx_messageInfo_FindUserRequest.Size(m)
}
func (m *FindUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindUserRequest proto.InternalMessageInfo

func (m *FindUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type FindUserReply struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Nickname             string   `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Intro                string   `protobuf:"bytes,5,opt,name=intro,proto3" json:"intro,omitempty"`
	IsApproved           bool     `protobuf:"varint,6,opt,name=is_approved,json=isApproved,proto3" json:"is_approved,omitempty"`
	CreatedAt            string   `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindUserReply) Reset()         { *m = FindUserReply{} }
func (m *FindUserReply) String() string { return proto.CompactTextString(m) }
func (*FindUserReply) ProtoMessage()    {}
func (*FindUserReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *FindUserReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindUserReply.Unmarshal(m, b)
}
func (m *FindUserReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindUserReply.Marshal(b, m, deterministic)
}
func (m *FindUserReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindUserReply.Merge(m, src)
}
func (m *FindUserReply) XXX_Size() int {
	return xxx_messageInfo_FindUserReply.Size(m)
}
func (m *FindUserReply) XXX_DiscardUnknown() {
	xxx_messageInfo_FindUserReply.DiscardUnknown(m)
}

var xxx_messageInfo_FindUserReply proto.InternalMessageInfo

func (m *FindUserReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FindUserReply) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *FindUserReply) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *FindUserReply) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *FindUserReply) GetIntro() string {
	if m != nil {
		return m.Intro
	}
	return ""
}

func (m *FindUserReply) GetIsApproved() bool {
	if m != nil {
		return m.IsApproved
	}
	return false
}

func (m *FindUserReply) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *FindUserReply) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type UpdateUserRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Nickname             string   `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Intro                string   `protobuf:"bytes,5,opt,name=intro,proto3" json:"intro,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(m, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UpdateUserRequest) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UpdateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UpdateUserRequest) GetIntro() string {
	if m != nil {
		return m.Intro
	}
	return ""
}

type UpdateUserReply struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserReply) Reset()         { *m = UpdateUserReply{} }
func (m *UpdateUserReply) String() string { return proto.CompactTextString(m) }
func (*UpdateUserReply) ProtoMessage()    {}
func (*UpdateUserReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UpdateUserReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserReply.Unmarshal(m, b)
}
func (m *UpdateUserReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserReply.Marshal(b, m, deterministic)
}
func (m *UpdateUserReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserReply.Merge(m, src)
}
func (m *UpdateUserReply) XXX_Size() int {
	return xxx_messageInfo_UpdateUserReply.Size(m)
}
func (m *UpdateUserReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserReply proto.InternalMessageInfo

func (m *UpdateUserReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*FindUserRequest)(nil), "pb.FindUserRequest")
	proto.RegisterType((*FindUserReply)(nil), "pb.FindUserReply")
	proto.RegisterType((*UpdateUserRequest)(nil), "pb.UpdateUserRequest")
	proto.RegisterType((*UpdateUserReply)(nil), "pb.UpdateUserReply")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0xb4, 0x0d, 0xe9, 0x20, 0x28, 0x75, 0x41, 0xb2, 0x22, 0x21, 0x4a, 0xd8, 0x74,
	0x95, 0x05, 0x70, 0x81, 0x6c, 0x38, 0x40, 0xa4, 0xae, 0x2b, 0xa7, 0xf6, 0xc2, 0x22, 0xc4, 0xc6,
	0x76, 0x40, 0x3d, 0x04, 0xf7, 0xe4, 0x18, 0x68, 0xec, 0x84, 0x9f, 0x20, 0xb1, 0x64, 0x97, 0xf7,
	0xbe, 0xc9, 0xd3, 0xfc, 0x18, 0xa0, 0xb3, 0xc2, 0x14, 0xda, 0x28, 0xa7, 0x48, 0xac, 0xeb, 0xfc,
	0x1a, 0x16, 0x0f, 0xb2, 0xe5, 0x5b, 0x2b, 0x4c, 0x25, 0x9e, 0x3b, 0x61, 0x1d, 0x39, 0x85, 0x58,
	0x72, 0x1a, 0xad, 0xa3, 0xcd, 0xbc, 0x8a, 0x25, 0xcf, 0xdf, 0x23, 0x38, 0xf9, 0xaa, 0xd1, 0xcd,
	0x61, 0x5c, 0x41, 0x32, 0x48, 0x31, 0xb6, 0x65, 0x4f, 0x82, 0xc6, 0xde, 0xfd, 0xd4, 0xc8, 0x5a,
	0xb9, 0x7f, 0xf4, 0x6c, 0x12, 0xd8, 0xa0, 0x91, 0x69, 0x66, 0xed, 0xab, 0x32, 0x9c, 0x4e, 0x03,
	0x1b, 0x34, 0x39, 0x87, 0x99, 0x6c, 0x9d, 0x51, 0x74, 0xe6, 0x41, 0x10, 0xe4, 0x0a, 0x8e, 0xa5,
	0xdd, 0x31, 0xad, 0x8d, 0x7a, 0x11, 0x9c, 0x26, 0xeb, 0x68, 0x93, 0x56, 0x20, 0x6d, 0xd9, 0x3b,
	0xe4, 0x12, 0x60, 0x6f, 0x04, 0x73, 0x82, 0xef, 0x98, 0xa3, 0x47, 0xfe, 0xdf, 0x79, 0xef, 0x94,
	0x0e, 0x71, 0xa7, 0xf9, 0x80, 0xd3, 0x80, 0x7b, 0xa7, 0x74, 0xf9, 0x5b, 0x04, 0xcb, 0xad, 0x57,
	0x7f, 0x2c, 0xe4, 0xff, 0xc6, 0xcd, 0x6f, 0x60, 0xf1, 0xbd, 0x1d, 0xdc, 0xfd, 0x19, 0x4c, 0x84,
	0x31, 0x7d, 0x37, 0xf8, 0x79, 0xdb, 0xc0, 0x14, 0x31, 0x29, 0x60, 0x8a, 0x67, 0x22, 0xab, 0x42,
	0xd7, 0xc5, 0xe8, 0xa8, 0xd9, 0xf2, 0xa7, 0x89, 0x49, 0xf7, 0x90, 0x84, 0x70, 0x72, 0x81, 0xf0,
	0xd7, 0xdc, 0xd9, 0x6a, 0x6c, 0xeb, 0xe6, 0x50, 0x27, 0xfe, 0xed, 0xdc, 0x7d, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x2d, 0xb1, 0x84, 0x20, 0x49, 0x02, 0x00, 0x00,
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
	Find(ctx context.Context, in *FindUserRequest, opts ...grpc.CallOption) (*FindUserReply, error)
	Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserReply, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Find(ctx context.Context, in *FindUserRequest, opts ...grpc.CallOption) (*FindUserReply, error) {
	out := new(FindUserReply)
	err := c.cc.Invoke(ctx, "/pb.User/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserReply, error) {
	out := new(UpdateUserReply)
	err := c.cc.Invoke(ctx, "/pb.User/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	Find(context.Context, *FindUserRequest) (*FindUserReply, error)
	Update(context.Context, *UpdateUserRequest) (*UpdateUserReply, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) Find(ctx context.Context, req *FindUserRequest) (*FindUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (*UnimplementedUserServer) Update(ctx context.Context, req *UpdateUserRequest) (*UpdateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Find(ctx, req.(*FindUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Update(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Find",
			Handler:    _User_Find_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _User_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
