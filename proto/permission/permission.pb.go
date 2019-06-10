// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/permission/permission.proto

package permission

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

type Permission struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Service              string   `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Method               string   `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	DisplayName          string   `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Description          string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Permission) Reset()         { *m = Permission{} }
func (m *Permission) String() string { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()    {}
func (*Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe0f260493db3286, []int{0}
}

func (m *Permission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission.Unmarshal(m, b)
}
func (m *Permission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission.Marshal(b, m, deterministic)
}
func (m *Permission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission.Merge(m, src)
}
func (m *Permission) XXX_Size() int {
	return xxx_messageInfo_Permission.Size(m)
}
func (m *Permission) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission.DiscardUnknown(m)
}

var xxx_messageInfo_Permission proto.InternalMessageInfo

func (m *Permission) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Permission) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Permission) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Permission) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Permission) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Permission) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Permission) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type ListQuery struct {
	Limit                int64    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 int64    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Sort                 string   `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	Service              string   `protobuf:"bytes,4,opt,name=service,proto3" json:"service,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName          string   `protobuf:"bytes,6,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListQuery) Reset()         { *m = ListQuery{} }
func (m *ListQuery) String() string { return proto.CompactTextString(m) }
func (*ListQuery) ProtoMessage()    {}
func (*ListQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe0f260493db3286, []int{1}
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

func (m *ListQuery) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *ListQuery) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListQuery) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe0f260493db3286, []int{2}
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

type Response struct {
	Permission           *Permission   `protobuf:"bytes,1,opt,name=permission,proto3" json:"permission,omitempty"`
	Permissions          []*Permission `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions,omitempty"`
	Total                int64         `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Valid                bool          `protobuf:"varint,4,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe0f260493db3286, []int{3}
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

func (m *Response) GetPermission() *Permission {
	if m != nil {
		return m.Permission
	}
	return nil
}

func (m *Response) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *Response) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Response) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func init() {
	proto.RegisterType((*Permission)(nil), "permission.Permission")
	proto.RegisterType((*ListQuery)(nil), "permission.ListQuery")
	proto.RegisterType((*Request)(nil), "permission.Request")
	proto.RegisterType((*Response)(nil), "permission.Response")
}

func init() { proto.RegisterFile("proto/permission/permission.proto", fileDescriptor_fe0f260493db3286) }

var fileDescriptor_fe0f260493db3286 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x25, 0xeb, 0x34, 0xdb, 0x9d, 0xa0, 0x1e, 0x86, 0x52, 0x59, 0x48, 0x48, 0xdb, 0x9c, 0x7a,
	0x2a, 0xd2, 0x22, 0x50, 0x8f, 0xac, 0x40, 0xe2, 0x82, 0xa0, 0x58, 0xe2, 0x5c, 0x99, 0xf5, 0x08,
	0x2c, 0x25, 0xeb, 0x60, 0x7b, 0x2b, 0xf5, 0x57, 0xb8, 0x70, 0xe5, 0x77, 0xf8, 0x23, 0x14, 0xdb,
	0x6d, 0x5c, 0x55, 0xdd, 0x43, 0x6e, 0xf3, 0xde, 0x9b, 0xa7, 0xcc, 0x9b, 0x71, 0xe0, 0xb4, 0xb7,
	0xc6, 0x9b, 0x57, 0x3d, 0xd9, 0x4e, 0x3b, 0xa7, 0xcd, 0x36, 0x2b, 0xcf, 0x83, 0x86, 0x30, 0x32,
	0xcd, 0xbf, 0x02, 0xe0, 0xf2, 0x0e, 0xe2, 0x11, 0xcc, 0xb4, 0xe2, 0xc5, 0xb2, 0x38, 0x63, 0x62,
	0xa6, 0x15, 0x72, 0x98, 0x3b, 0xb2, 0xd7, 0x7a, 0x43, 0x7c, 0xb6, 0x2c, 0xce, 0x16, 0xe2, 0x16,
	0xe2, 0x09, 0x54, 0x1d, 0xf9, 0x9f, 0x46, 0x71, 0x16, 0x84, 0x84, 0xf0, 0x14, 0x9e, 0x2a, 0xed,
	0xfa, 0x56, 0xde, 0x5c, 0x6d, 0x65, 0x47, 0xbc, 0x0c, 0x6a, 0x9d, 0xb8, 0xcf, 0xb2, 0x23, 0x5c,
	0x42, 0xad, 0xc8, 0x6d, 0xac, 0xee, 0xbd, 0x36, 0x5b, 0x7e, 0x90, 0x3a, 0x46, 0x0a, 0x5f, 0x02,
	0x6c, 0x2c, 0x49, 0x4f, 0xea, 0x4a, 0x7a, 0x5e, 0x85, 0x86, 0x45, 0x62, 0xd6, 0x7e, 0x90, 0x77,
	0xbd, 0xba, 0x95, 0xe7, 0x51, 0x4e, 0xcc, 0xda, 0x37, 0xbf, 0x0b, 0x58, 0x7c, 0xd2, 0xce, 0x7f,
	0xdd, 0x91, 0xbd, 0xc1, 0x63, 0x38, 0x68, 0x75, 0xa7, 0x7d, 0x4a, 0x15, 0x01, 0x22, 0x94, 0xbd,
	0xfc, 0x11, 0x53, 0x31, 0x11, 0xea, 0x81, 0x73, 0xc6, 0xfa, 0x14, 0x28, 0xd4, 0xf9, 0x02, 0xca,
	0xfb, 0x0b, 0x40, 0x28, 0x43, 0xc0, 0x38, 0x7e, 0xa8, 0x1f, 0x84, 0xaf, 0x1e, 0x84, 0x6f, 0x16,
	0x30, 0x17, 0xf4, 0x6b, 0x47, 0xce, 0x37, 0x7f, 0x0b, 0x38, 0x14, 0xe4, 0x7a, 0xb3, 0x75, 0x84,
	0x6f, 0x21, 0x3b, 0x4b, 0x98, 0xb5, 0x5e, 0x9d, 0x9c, 0x67, 0xb7, 0x1b, 0xaf, 0x24, 0xb2, 0x4e,
	0xbc, 0x80, 0x7a, 0x44, 0x8e, 0xcf, 0x96, 0x6c, 0x8f, 0x31, 0x6f, 0x1d, 0x16, 0xe3, 0x8d, 0x97,
	0x6d, 0xc8, 0xcb, 0x44, 0x04, 0x03, 0x7b, 0x2d, 0x5b, 0xad, 0x42, 0xdc, 0x43, 0x11, 0xc1, 0xea,
	0x0f, 0x83, 0xfa, 0x32, 0xf3, 0xae, 0x80, 0xad, 0xdb, 0x16, 0x9f, 0xe5, 0xdf, 0x49, 0xb1, 0x5e,
	0x1c, 0xdf, 0x27, 0x63, 0xbe, 0xe6, 0x09, 0xbe, 0x81, 0x72, 0xb8, 0x0a, 0x3e, 0xcf, 0xf5, 0xbb,
	0x3b, 0xed, 0xb1, 0xb1, 0x8f, 0xe4, 0xf1, 0x91, 0x48, 0x8f, 0xda, 0x2e, 0xa0, 0x7a, 0x1f, 0x1e,
	0xcc, 0x14, 0xe7, 0xb7, 0xf0, 0x96, 0xa6, 0x38, 0x3f, 0x50, 0x4b, 0x13, 0x9c, 0xef, 0xe0, 0x28,
	0x7e, 0xf3, 0x8b, 0x9d, 0x36, 0xf5, 0xf7, 0x2a, 0xfc, 0xdb, 0xaf, 0xff, 0x07, 0x00, 0x00, 0xff,
	0xff, 0x3f, 0x10, 0x0f, 0xa1, 0x00, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Permissions service

type PermissionsClient interface {
	// 权限验证授权
	// 全部权限
	All(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 获取权限列表
	List(ctx context.Context, in *ListQuery, opts ...client.CallOption) (*Response, error)
	// 根据 唯一 获取权限
	Get(ctx context.Context, in *Permission, opts ...client.CallOption) (*Response, error)
	// 创建权限
	Create(ctx context.Context, in *Permission, opts ...client.CallOption) (*Response, error)
	// 更新权限
	Update(ctx context.Context, in *Permission, opts ...client.CallOption) (*Response, error)
	// 删除权限
	Delete(ctx context.Context, in *Permission, opts ...client.CallOption) (*Response, error)
	// 微服务内部调用
	// 同步
	UpdateOrCreate(ctx context.Context, in *Permission, opts ...client.CallOption) (*Response, error)
}

type permissionsClient struct {
	c           client.Client
	serviceName string
}

func NewPermissionsClient(serviceName string, c client.Client) PermissionsClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "permission"
	}
	return &permissionsClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *permissionsClient) All(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Permissions.All", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsClient) List(ctx context.Context, in *ListQuery, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Permissions.List", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsClient) Get(ctx context.Context, in *Permission, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Permissions.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsClient) Create(ctx context.Context, in *Permission, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Permissions.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsClient) Update(ctx context.Context, in *Permission, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Permissions.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsClient) Delete(ctx context.Context, in *Permission, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Permissions.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsClient) UpdateOrCreate(ctx context.Context, in *Permission, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Permissions.UpdateOrCreate", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Permissions service

type PermissionsHandler interface {
	// 权限验证授权
	// 全部权限
	All(context.Context, *Request, *Response) error
	// 获取权限列表
	List(context.Context, *ListQuery, *Response) error
	// 根据 唯一 获取权限
	Get(context.Context, *Permission, *Response) error
	// 创建权限
	Create(context.Context, *Permission, *Response) error
	// 更新权限
	Update(context.Context, *Permission, *Response) error
	// 删除权限
	Delete(context.Context, *Permission, *Response) error
	// 微服务内部调用
	// 同步
	UpdateOrCreate(context.Context, *Permission, *Response) error
}

func RegisterPermissionsHandler(s server.Server, hdlr PermissionsHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Permissions{hdlr}, opts...))
}

type Permissions struct {
	PermissionsHandler
}

func (h *Permissions) All(ctx context.Context, in *Request, out *Response) error {
	return h.PermissionsHandler.All(ctx, in, out)
}

func (h *Permissions) List(ctx context.Context, in *ListQuery, out *Response) error {
	return h.PermissionsHandler.List(ctx, in, out)
}

func (h *Permissions) Get(ctx context.Context, in *Permission, out *Response) error {
	return h.PermissionsHandler.Get(ctx, in, out)
}

func (h *Permissions) Create(ctx context.Context, in *Permission, out *Response) error {
	return h.PermissionsHandler.Create(ctx, in, out)
}

func (h *Permissions) Update(ctx context.Context, in *Permission, out *Response) error {
	return h.PermissionsHandler.Update(ctx, in, out)
}

func (h *Permissions) Delete(ctx context.Context, in *Permission, out *Response) error {
	return h.PermissionsHandler.Delete(ctx, in, out)
}

func (h *Permissions) UpdateOrCreate(ctx context.Context, in *Permission, out *Response) error {
	return h.PermissionsHandler.UpdateOrCreate(ctx, in, out)
}
