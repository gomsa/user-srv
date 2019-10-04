// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/role/role.proto

package role

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

type Role struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Label                string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_26e011caf756e89c, []int{0}
}

func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Role) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Role) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Role) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type ListQuery struct {
	Limit                int64    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 int64    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Sort                 string   `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	Label                string   `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListQuery) Reset()         { *m = ListQuery{} }
func (m *ListQuery) String() string { return proto.CompactTextString(m) }
func (*ListQuery) ProtoMessage()    {}
func (*ListQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_26e011caf756e89c, []int{1}
}

func (m *ListQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListQuery.Unmarshal(m, b)
}
func (m *ListQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListQuery.Marshal(b, m, deterministic)
}
func (m *ListQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListQuery.Merge(m, src)
}
func (m *ListQuery) XXX_Size() int {
	return xxx_messageInfo_ListQuery.Size(m)
}
func (m *ListQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ListQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ListQuery proto.InternalMessageInfo

func (m *ListQuery) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListQuery) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListQuery) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *ListQuery) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *ListQuery) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Request struct {
	ListQuery            *ListQuery `protobuf:"bytes,1,opt,name=list_query,json=listQuery,proto3" json:"list_query,omitempty"`
	Role                 *Role      `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_26e011caf756e89c, []int{2}
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

func (m *Request) GetListQuery() *ListQuery {
	if m != nil {
		return m.ListQuery
	}
	return nil
}

func (m *Request) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type Response struct {
	Valid                bool     `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Total                int64    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Role                 *Role    `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	Roles                []*Role  `protobuf:"bytes,4,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_26e011caf756e89c, []int{3}
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

func (m *Response) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Response) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Response) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *Response) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func init() {
	proto.RegisterType((*Role)(nil), "role.Role")
	proto.RegisterType((*ListQuery)(nil), "role.ListQuery")
	proto.RegisterType((*Request)(nil), "role.Request")
	proto.RegisterType((*Response)(nil), "role.Response")
}

func init() { proto.RegisterFile("proto/role/role.proto", fileDescriptor_26e011caf756e89c) }

var fileDescriptor_26e011caf756e89c = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4b, 0xeb, 0x40,
	0x10, 0xc7, 0x5f, 0x92, 0x4d, 0x5f, 0x33, 0xe5, 0xf5, 0xc1, 0xa2, 0x10, 0x04, 0x25, 0xe4, 0x50,
	0xeb, 0xa5, 0x42, 0xfd, 0x04, 0x45, 0xc1, 0x8b, 0x17, 0x17, 0x3c, 0x78, 0x2a, 0xdb, 0x66, 0x90,
	0x85, 0x6d, 0x37, 0xcd, 0x6e, 0x45, 0x8f, 0x7e, 0x13, 0x3f, 0xaa, 0xec, 0x6c, 0xda, 0x8a, 0x14,
	0xcc, 0x25, 0x99, 0xff, 0xcc, 0x9f, 0xcd, 0xef, 0x3f, 0x59, 0x38, 0xad, 0x1b, 0xe3, 0xcc, 0x75,
	0x63, 0x34, 0xd2, 0x63, 0x42, 0x9a, 0x33, 0x5f, 0x97, 0x9f, 0x11, 0x30, 0x61, 0x34, 0xf2, 0x21,
	0xc4, 0xaa, 0xca, 0xa3, 0x22, 0x1a, 0x27, 0x22, 0x56, 0x15, 0x3f, 0x81, 0x54, 0xcb, 0x05, 0xea,
	0x3c, 0x2e, 0xa2, 0x71, 0x26, 0x82, 0xe0, 0x1c, 0xd8, 0x5a, 0xae, 0x30, 0x4f, 0xa8, 0x49, 0x35,
	0x2f, 0x60, 0x50, 0xa1, 0x5d, 0x36, 0xaa, 0x76, 0xca, 0xac, 0x73, 0x46, 0xa3, 0xef, 0x2d, 0x7e,
	0x0e, 0xb0, 0x6c, 0x50, 0x3a, 0xac, 0xe6, 0xd2, 0xe5, 0x29, 0x19, 0xb2, 0xb6, 0x33, 0x73, 0x7e,
	0xbc, 0xad, 0xab, 0xdd, 0xb8, 0x17, 0xc6, 0x6d, 0x67, 0xe6, 0x4a, 0x0b, 0xd9, 0x83, 0xb2, 0xee,
	0x71, 0x8b, 0xcd, 0x3b, 0x61, 0xa9, 0x95, 0x72, 0x2d, 0x69, 0x10, 0x1e, 0xab, 0x96, 0x2f, 0x48,
	0xac, 0x89, 0xa0, 0xda, 0xf7, 0xac, 0x69, 0xdc, 0x0e, 0xd5, 0xd7, 0x87, 0x50, 0xec, 0x58, 0xa8,
	0xf4, 0x10, 0xaa, 0x7c, 0x86, 0xbf, 0x02, 0x37, 0x5b, 0xb4, 0x8e, 0x4f, 0x00, 0xb4, 0xb2, 0x6e,
	0xbe, 0xf1, 0x00, 0xf4, 0xdd, 0xc1, 0xf4, 0xff, 0x84, 0x36, 0xb9, 0xe7, 0x12, 0x99, 0xde, 0x23,
	0x5e, 0x00, 0xad, 0x96, 0x60, 0x06, 0x53, 0x08, 0x4e, 0xbf, 0x63, 0x11, 0x56, 0xfe, 0x06, 0x7d,
	0x81, 0xb6, 0x36, 0x6b, 0x8b, 0x1e, 0xe8, 0x55, 0xea, 0x76, 0xf1, 0x7d, 0x11, 0x84, 0xef, 0x3a,
	0xe3, 0xa4, 0x6e, 0xf3, 0x04, 0xb1, 0x3f, 0x37, 0x39, 0x7e, 0x2e, 0x2f, 0x20, 0xf5, 0x6f, 0x9b,
	0xb3, 0x22, 0xf9, 0x61, 0x08, 0x83, 0xe9, 0x47, 0x0c, 0xa9, 0xd7, 0x96, 0x8f, 0x20, 0x99, 0x69,
	0xcd, 0xff, 0xb5, 0x9e, 0x90, 0xf4, 0x6c, 0xb8, 0x93, 0x81, 0xae, 0xfc, 0xc3, 0x2f, 0x81, 0xf9,
	0x8c, 0xbf, 0x1b, 0x47, 0x90, 0xdc, 0x63, 0x07, 0xdf, 0x15, 0xf4, 0x6e, 0xe9, 0xc7, 0x77, 0xb2,
	0x3e, 0xd1, 0x25, 0xe8, 0x64, 0xbd, 0x43, 0x8d, 0x1d, 0xac, 0x8b, 0x1e, 0xdd, 0xfe, 0x9b, 0xaf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xa3, 0x61, 0xd9, 0x16, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Roles service

type RolesClient interface {
	// 权限验证授权
	// 全部权限
	All(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 获取权限列表
	List(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 根据 唯一 获取权限
	Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 创建权限
	Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 更新权限
	Update(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 删除权限
	Delete(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type rolesClient struct {
	c           client.Client
	serviceName string
}

func NewRolesClient(serviceName string, c client.Client) RolesClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "role"
	}
	return &rolesClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *rolesClient) All(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Roles.All", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) List(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Roles.List", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Roles.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Roles.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) Update(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Roles.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) Delete(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Roles.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Roles service

type RolesHandler interface {
	// 权限验证授权
	// 全部权限
	All(context.Context, *Request, *Response) error
	// 获取权限列表
	List(context.Context, *Request, *Response) error
	// 根据 唯一 获取权限
	Get(context.Context, *Request, *Response) error
	// 创建权限
	Create(context.Context, *Request, *Response) error
	// 更新权限
	Update(context.Context, *Request, *Response) error
	// 删除权限
	Delete(context.Context, *Request, *Response) error
}

func RegisterRolesHandler(s server.Server, hdlr RolesHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Roles{hdlr}, opts...))
}

type Roles struct {
	RolesHandler
}

func (h *Roles) All(ctx context.Context, in *Request, out *Response) error {
	return h.RolesHandler.All(ctx, in, out)
}

func (h *Roles) List(ctx context.Context, in *Request, out *Response) error {
	return h.RolesHandler.List(ctx, in, out)
}

func (h *Roles) Get(ctx context.Context, in *Request, out *Response) error {
	return h.RolesHandler.Get(ctx, in, out)
}

func (h *Roles) Create(ctx context.Context, in *Request, out *Response) error {
	return h.RolesHandler.Create(ctx, in, out)
}

func (h *Roles) Update(ctx context.Context, in *Request, out *Response) error {
	return h.RolesHandler.Update(ctx, in, out)
}

func (h *Roles) Delete(ctx context.Context, in *Request, out *Response) error {
	return h.RolesHandler.Delete(ctx, in, out)
}
