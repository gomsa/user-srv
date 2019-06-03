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
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName          string   `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
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

func (m *Permission) GetName() string {
	if m != nil {
		return m.Name
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
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName          string   `protobuf:"bytes,5,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
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
	return fileDescriptor_fe0f260493db3286, []int{2}
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
	proto.RegisterType((*Response)(nil), "permission.Response")
}

func init() { proto.RegisterFile("proto/permission/permission.proto", fileDescriptor_fe0f260493db3286) }

var fileDescriptor_fe0f260493db3286 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4f, 0x4b, 0xfb, 0x40,
	0x14, 0xfc, 0xe5, 0x4f, 0x4b, 0xfb, 0xf2, 0xc3, 0xc3, 0x52, 0x25, 0x08, 0x42, 0x9a, 0x53, 0x4f,
	0x15, 0x2a, 0x95, 0x5e, 0x8b, 0x82, 0x17, 0x91, 0xba, 0xe0, 0xb9, 0xac, 0xdd, 0x45, 0x16, 0xd2,
	0xec, 0x92, 0xdd, 0x0a, 0xbd, 0xf9, 0x71, 0xf4, 0xec, 0x17, 0x94, 0x7d, 0x49, 0x9b, 0x05, 0x69,
	0xc1, 0xde, 0xde, 0x9b, 0x79, 0x43, 0x66, 0x66, 0x09, 0x0c, 0x75, 0xa5, 0xac, 0xba, 0xd6, 0xa2,
	0x5a, 0x4b, 0x63, 0xa4, 0x2a, 0xbd, 0x71, 0x8c, 0x1c, 0x81, 0x16, 0xc9, 0xbf, 0x03, 0x80, 0xc5,
	0x7e, 0x25, 0x67, 0x10, 0x4a, 0x9e, 0x06, 0x59, 0x30, 0x8a, 0x68, 0x28, 0x39, 0x21, 0x10, 0x97,
	0x6c, 0x2d, 0xd2, 0x30, 0x0b, 0x46, 0x7d, 0x8a, 0x33, 0x19, 0xc2, 0x7f, 0x2e, 0x8d, 0x2e, 0xd8,
	0x76, 0x89, 0x5c, 0x84, 0x5c, 0xd2, 0x60, 0x4f, 0xee, 0x24, 0x83, 0x84, 0x0b, 0xb3, 0xaa, 0xa4,
	0xb6, 0x52, 0x95, 0x69, 0xdc, 0x5c, 0xb4, 0x10, 0xb9, 0x02, 0x58, 0x55, 0x82, 0x59, 0xc1, 0x97,
	0xcc, 0xa6, 0x1d, 0x3c, 0xe8, 0x37, 0xc8, 0xdc, 0x3a, 0x7a, 0xa3, 0xf9, 0x8e, 0xee, 0xd6, 0x74,
	0x83, 0xcc, 0x6d, 0xfe, 0x11, 0x40, 0xff, 0x51, 0x1a, 0xfb, 0xbc, 0x11, 0xd5, 0x96, 0x0c, 0xa0,
	0x53, 0xc8, 0xb5, 0xb4, 0x8d, 0xef, 0x7a, 0x71, 0xd6, 0x35, 0x7b, 0xab, 0xad, 0x47, 0x14, 0x67,
	0x87, 0x19, 0x55, 0xd9, 0xc6, 0x32, 0xce, 0xfb, 0x88, 0xf1, 0x91, 0x88, 0x9d, 0x5f, 0x11, 0xf3,
	0xcf, 0x00, 0x7a, 0x54, 0x18, 0xad, 0x4a, 0x23, 0xc8, 0x2d, 0x78, 0x9d, 0xa2, 0x8d, 0x64, 0x72,
	0x31, 0xf6, 0x8a, 0x6f, 0x2b, 0xa6, 0xde, 0x25, 0x99, 0x41, 0xd2, 0x6e, 0x26, 0x0d, 0xb3, 0xe8,
	0x88, 0xd0, 0x3f, 0x75, 0x99, 0xad, 0xb2, 0xac, 0xc0, 0x28, 0x11, 0xad, 0x17, 0x87, 0xbe, 0xb3,
	0x42, 0x72, 0x0c, 0xd3, 0xa3, 0xf5, 0x32, 0xf9, 0x0a, 0x21, 0x59, 0x78, 0xda, 0x29, 0xc4, 0xae,
	0x3c, 0x72, 0xee, 0x7f, 0x68, 0x5f, 0xe7, 0xe5, 0xc0, 0x87, 0x77, 0x11, 0xf3, 0x7f, 0x64, 0x0a,
	0xd1, 0x83, 0xb0, 0xe4, 0x80, 0xbd, 0x83, 0xb2, 0x19, 0x74, 0xef, 0xf0, 0x5d, 0x4f, 0x51, 0xbe,
	0xe0, 0x93, 0x9f, 0xa2, 0xbc, 0x17, 0x85, 0xf8, 0xbb, 0xf2, 0xb5, 0x8b, 0xbf, 0xc8, 0xcd, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x6a, 0x2d, 0x99, 0x47, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Permissions service

type PermissionsClient interface {
	// 权限验证授权
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

// Server API for Permissions service

type PermissionsHandler interface {
	// 权限验证授权
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
}

func RegisterPermissionsHandler(s server.Server, hdlr PermissionsHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Permissions{hdlr}, opts...))
}

type Permissions struct {
	PermissionsHandler
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
