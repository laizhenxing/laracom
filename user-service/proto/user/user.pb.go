// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package laracom_service_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
	context "golang.org/x/net/context"
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
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	StripeId             string   `protobuf:"bytes,6,opt,name=stripe_id,json=stripeId,proto3" json:"stripe_id,omitempty"`
	CardBrand            string   `protobuf:"bytes,7,opt,name=card_brand,json=cardBrand,proto3" json:"card_brand,omitempty"`
	CardLastFour         string   `protobuf:"bytes,8,opt,name=card_last_four,json=cardLastFour,proto3" json:"card_last_four,omitempty"`
	TrialEndsAt          string   `protobuf:"bytes,9,opt,name=trial_ends_at,json=trialEndsAt,proto3" json:"trial_ends_at,omitempty"`
	DeletedAt            string   `protobuf:"bytes,10,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	RememberToken        string   `protobuf:"bytes,11,opt,name=remember_token,json=rememberToken,proto3" json:"remember_token,omitempty"`
	CreatedAt            string   `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
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

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *User) GetStripeId() string {
	if m != nil {
		return m.StripeId
	}
	return ""
}

func (m *User) GetCardBrand() string {
	if m != nil {
		return m.CardBrand
	}
	return ""
}

func (m *User) GetCardLastFour() string {
	if m != nil {
		return m.CardLastFour
	}
	return ""
}

func (m *User) GetTrialEndsAt() string {
	if m != nil {
		return m.TrialEndsAt
	}
	return ""
}

func (m *User) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

func (m *User) GetRememberToken() string {
	if m != nil {
		return m.RememberToken
	}
	return ""
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *User) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type UserRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

type UserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UserResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *UserResponse) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Decription           string   `protobuf:"bytes,2,opt,name=decription,proto3" json:"decription,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
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

func (m *Error) GetDecription() string {
	if m != nil {
		return m.Decription
	}
	return ""
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
	return fileDescriptor_9b283a848145d6b7, []int{4}
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

func init() {
	proto.RegisterType((*User)(nil), "laracom.service.user.User")
	proto.RegisterType((*UserRequest)(nil), "laracom.service.user.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "laracom.service.user.UserResponse")
	proto.RegisterType((*Error)(nil), "laracom.service.user.Error")
	proto.RegisterType((*Token)(nil), "laracom.service.user.Token")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcd, 0x8a, 0x13, 0x41,
	0x10, 0xc7, 0xd9, 0x24, 0x33, 0x26, 0x95, 0x8f, 0x43, 0x13, 0xa5, 0xc9, 0xa2, 0xac, 0x83, 0x82,
	0xa7, 0x51, 0x76, 0x8f, 0x1e, 0x24, 0x2b, 0x31, 0x2e, 0x2c, 0x08, 0xf1, 0xe3, 0x3a, 0x74, 0xd2,
	0x25, 0xdb, 0x38, 0x99, 0x1e, 0xbb, 0x6b, 0x56, 0x7c, 0x20, 0x1f, 0xc6, 0x97, 0x12, 0xe9, 0xea,
	0x89, 0x78, 0x58, 0x47, 0x71, 0x2f, 0x49, 0xd7, 0xaf, 0xfe, 0x55, 0x5d, 0xd4, 0xbf, 0x13, 0xb8,
	0x5b, 0x3b, 0x4b, 0xf6, 0x69, 0xe3, 0xd1, 0xf1, 0x47, 0xce, 0xb1, 0x98, 0x97, 0xca, 0xa9, 0x9d,
	0xdd, 0xe7, 0x1e, 0xdd, 0xb5, 0xd9, 0x61, 0x1e, 0x72, 0xd9, 0x8f, 0x1e, 0x0c, 0xde, 0x7b, 0x74,
	0x62, 0x06, 0x3d, 0xa3, 0xe5, 0xd1, 0xc9, 0xd1, 0x93, 0xd1, 0xa6, 0x67, 0xb4, 0x10, 0x30, 0xa8,
	0xd4, 0x1e, 0x65, 0x8f, 0x09, 0x9f, 0xc5, 0x1c, 0x12, 0xdc, 0x2b, 0x53, 0xca, 0x3e, 0xc3, 0x18,
	0x88, 0x05, 0x0c, 0x6b, 0xe5, 0xfd, 0x17, 0xeb, 0xb4, 0x1c, 0x70, 0xe2, 0x57, 0x2c, 0xee, 0x41,
	0xea, 0x49, 0x51, 0xe3, 0x65, 0xc2, 0x99, 0x36, 0x12, 0xc7, 0x30, 0xf2, 0xe4, 0x4c, 0x8d, 0x85,
	0xd1, 0x32, 0x8d, 0x45, 0x11, 0x5c, 0x68, 0x71, 0x1f, 0x60, 0xa7, 0x9c, 0x2e, 0xb6, 0x4e, 0x55,
	0x5a, 0xde, 0xe1, 0xec, 0x28, 0x90, 0xf3, 0x00, 0xc4, 0x23, 0x98, 0x71, 0xba, 0x54, 0x9e, 0x8a,
	0x8f, 0xb6, 0x71, 0x72, 0xc8, 0x92, 0x49, 0xa0, 0x97, 0xca, 0xd3, 0x2b, 0xdb, 0x38, 0x91, 0xc1,
	0x94, 0x9c, 0x51, 0x65, 0x81, 0x95, 0xf6, 0x85, 0x22, 0x39, 0x62, 0xd1, 0x98, 0xe1, 0xaa, 0xd2,
	0x7e, 0x49, 0xe1, 0x22, 0x8d, 0x25, 0x12, 0xea, 0x20, 0x80, 0x78, 0x51, 0x4b, 0x96, 0x24, 0x1e,
	0xc3, 0xcc, 0xe1, 0x1e, 0xf7, 0x5b, 0x74, 0x05, 0xd9, 0x4f, 0x58, 0xc9, 0x31, 0x4b, 0xa6, 0x07,
	0xfa, 0x2e, 0x40, 0x1e, 0xd7, 0xa1, 0x6a, 0xbb, 0x4c, 0xda, 0x71, 0x23, 0x89, 0x97, 0x34, 0xb5,
	0x3e, 0xa4, 0xa7, 0x31, 0xdd, 0x92, 0x25, 0x65, 0x53, 0x18, 0x87, 0xfd, 0x6f, 0xf0, 0x73, 0x83,
	0x9e, 0xb2, 0x6f, 0x47, 0x30, 0x89, 0xb1, 0xaf, 0x6d, 0xe5, 0x51, 0xe4, 0x30, 0x08, 0x46, 0xb1,
	0x33, 0xe3, 0xd3, 0x45, 0x7e, 0x93, 0x8b, 0x39, 0x57, 0xb0, 0x4e, 0x3c, 0x83, 0x24, 0x7c, 0x7b,
	0xd9, 0x3b, 0xe9, 0xff, 0xa5, 0x20, 0x0a, 0xc5, 0x19, 0xa4, 0xe8, 0x9c, 0x75, 0x5e, 0xf6, 0xb9,
	0xe4, 0xf8, 0xe6, 0x92, 0x55, 0xd0, 0x6c, 0x5a, 0x69, 0xf6, 0x1c, 0x12, 0x06, 0xe1, 0x9d, 0xec,
	0xac, 0x46, 0x9e, 0x2f, 0xd9, 0xf0, 0x59, 0x3c, 0x08, 0x7b, 0xdd, 0x39, 0x53, 0x93, 0xb1, 0x55,
	0xfb, 0x82, 0x7e, 0x23, 0xd9, 0x15, 0x24, 0x71, 0x75, 0x73, 0x48, 0xe2, 0x62, 0xe3, 0xbb, 0x8b,
	0x41, 0xa0, 0xd7, 0xaa, 0x34, 0x9a, 0x2b, 0x87, 0x9b, 0x18, 0xfc, 0xd7, 0x98, 0xa7, 0xdf, 0xfb,
	0x71, 0xbd, 0x6f, 0xa3, 0x44, 0xbc, 0x86, 0xf4, 0x25, 0x3b, 0x23, 0x3a, 0x16, 0xb3, 0xc8, 0x3a,
	0x96, 0x76, 0xf0, 0x65, 0x05, 0xfd, 0x35, 0xd2, 0xad, 0xdb, 0xbc, 0x81, 0x74, 0x8d, 0xb4, 0x2c,
	0x4b, 0xf1, 0xb0, 0x4b, 0xcd, 0x8f, 0xe3, 0x9f, 0x1a, 0x5e, 0x02, 0xac, 0x91, 0xce, 0xbf, 0xae,
	0xe2, 0x6f, 0xf3, 0x96, 0xe3, 0xbd, 0x80, 0xc1, 0xb2, 0xa1, 0xab, 0xce, 0x3e, 0x7f, 0x30, 0x22,
	0x3a, 0x7c, 0x01, 0xd3, 0x0f, 0xc1, 0x3e, 0x45, 0x18, 0x41, 0x97, 0xba, 0xb3, 0xd5, 0x36, 0xe5,
	0xff, 0xb1, 0xb3, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x66, 0x71, 0x00, 0x7d, 0xe0, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*UserResponse, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*UserResponse, error)
	GetAll(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error)
	GetByEmail(ctx context.Context, in *User, opts ...client.CallOption) (*UserResponse, error)
	Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "laracom.service.user"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Create", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *User, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Get", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAll(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetAll", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetByEmail(ctx context.Context, in *User, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetByEmail", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Auth", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.ValidateToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Create(context.Context, *User, *UserResponse) error
	Get(context.Context, *User, *UserResponse) error
	GetAll(context.Context, *UserRequest, *UserResponse) error
	GetByEmail(context.Context, *User, *UserResponse) error
	Auth(context.Context, *User, *Token) error
	ValidateToken(context.Context, *Token, *Token) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) Create(ctx context.Context, in *User, out *UserResponse) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *UserService) Get(ctx context.Context, in *User, out *UserResponse) error {
	return h.UserServiceHandler.Get(ctx, in, out)
}

func (h *UserService) GetAll(ctx context.Context, in *UserRequest, out *UserResponse) error {
	return h.UserServiceHandler.GetAll(ctx, in, out)
}

func (h *UserService) GetByEmail(ctx context.Context, in *User, out *UserResponse) error {
	return h.UserServiceHandler.GetByEmail(ctx, in, out)
}

func (h *UserService) Auth(ctx context.Context, in *User, out *Token) error {
	return h.UserServiceHandler.Auth(ctx, in, out)
}

func (h *UserService) ValidateToken(ctx context.Context, in *Token, out *Token) error {
	return h.UserServiceHandler.ValidateToken(ctx, in, out)
}
