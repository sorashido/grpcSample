// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	GetItemRequest
	ListItemRequest
	ListItemResponse
	AddItemResponse
	UpdateItemRequest
	UpdateItemResponse
	DeleteItemRequest
	DeleteItemResponse
	Item
*/
package api

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

type GetItemRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetItemRequest) Reset()                    { *m = GetItemRequest{} }
func (m *GetItemRequest) String() string            { return proto.CompactTextString(m) }
func (*GetItemRequest) ProtoMessage()               {}
func (*GetItemRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetItemRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListItemRequest struct {
	Page  int32 `protobuf:"varint,1,opt,name=page" json:"page,omitempty"`
	Limit int32 `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
}

func (m *ListItemRequest) Reset()                    { *m = ListItemRequest{} }
func (m *ListItemRequest) String() string            { return proto.CompactTextString(m) }
func (*ListItemRequest) ProtoMessage()               {}
func (*ListItemRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListItemRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListItemRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ListItemResponse struct {
	Items    []*Item `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	Total    int32   `protobuf:"varint,2,opt,name=total" json:"total,omitempty"`
	PrevPage int32   `protobuf:"varint,3,opt,name=prevPage" json:"prevPage,omitempty"`
	NextPage int32   `protobuf:"varint,4,opt,name=nextPage" json:"nextPage,omitempty"`
}

func (m *ListItemResponse) Reset()                    { *m = ListItemResponse{} }
func (m *ListItemResponse) String() string            { return proto.CompactTextString(m) }
func (*ListItemResponse) ProtoMessage()               {}
func (*ListItemResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListItemResponse) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ListItemResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListItemResponse) GetPrevPage() int32 {
	if m != nil {
		return m.PrevPage
	}
	return 0
}

func (m *ListItemResponse) GetNextPage() int32 {
	if m != nil {
		return m.NextPage
	}
	return 0
}

type AddItemResponse struct {
	Item *Item `protobuf:"bytes,1,opt,name=item" json:"item,omitempty"`
}

func (m *AddItemResponse) Reset()                    { *m = AddItemResponse{} }
func (m *AddItemResponse) String() string            { return proto.CompactTextString(m) }
func (*AddItemResponse) ProtoMessage()               {}
func (*AddItemResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AddItemResponse) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type UpdateItemRequest struct {
	Item *Item `protobuf:"bytes,1,opt,name=item" json:"item,omitempty"`
}

func (m *UpdateItemRequest) Reset()                    { *m = UpdateItemRequest{} }
func (m *UpdateItemRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateItemRequest) ProtoMessage()               {}
func (*UpdateItemRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateItemRequest) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type UpdateItemResponse struct {
	Item *Item `protobuf:"bytes,1,opt,name=item" json:"item,omitempty"`
}

func (m *UpdateItemResponse) Reset()                    { *m = UpdateItemResponse{} }
func (m *UpdateItemResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateItemResponse) ProtoMessage()               {}
func (*UpdateItemResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateItemResponse) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type DeleteItemRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteItemRequest) Reset()                    { *m = DeleteItemRequest{} }
func (m *DeleteItemRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteItemRequest) ProtoMessage()               {}
func (*DeleteItemRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteItemRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteItemResponse struct {
}

func (m *DeleteItemResponse) Reset()                    { *m = DeleteItemResponse{} }
func (m *DeleteItemResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteItemResponse) ProtoMessage()               {}
func (*DeleteItemResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type Item struct {
	Id          string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Title       string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Price       int32  `protobuf:"varint,5,opt,name=price" json:"price,omitempty"`
	Pv          int32  `protobuf:"varint,6,opt,name=pv" json:"pv,omitempty"`
	Status      bool   `protobuf:"varint,7,opt,name=status" json:"status,omitempty"`
}

func (m *Item) Reset()                    { *m = Item{} }
func (m *Item) String() string            { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()               {}
func (*Item) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Item) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Item) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Item) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Item) GetPv() int32 {
	if m != nil {
		return m.Pv
	}
	return 0
}

func (m *Item) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func init() {
	proto.RegisterType((*GetItemRequest)(nil), "GetItemRequest")
	proto.RegisterType((*ListItemRequest)(nil), "ListItemRequest")
	proto.RegisterType((*ListItemResponse)(nil), "ListItemResponse")
	proto.RegisterType((*AddItemResponse)(nil), "AddItemResponse")
	proto.RegisterType((*UpdateItemRequest)(nil), "UpdateItemRequest")
	proto.RegisterType((*UpdateItemResponse)(nil), "UpdateItemResponse")
	proto.RegisterType((*DeleteItemRequest)(nil), "DeleteItemRequest")
	proto.RegisterType((*DeleteItemResponse)(nil), "DeleteItemResponse")
	proto.RegisterType((*Item)(nil), "Item")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for API service

type APIClient interface {
	ListItem(ctx context.Context, in *ListItemRequest, opts ...grpc.CallOption) (*ListItemResponse, error)
	GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*Item, error)
	AddItem(ctx context.Context, in *Item, opts ...grpc.CallOption) (*AddItemResponse, error)
	UpdateItem(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*UpdateItemResponse, error)
	DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*DeleteItemResponse, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) ListItem(ctx context.Context, in *ListItemRequest, opts ...grpc.CallOption) (*ListItemResponse, error) {
	out := new(ListItemResponse)
	err := grpc.Invoke(ctx, "/API/ListItem", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := grpc.Invoke(ctx, "/API/GetItem", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) AddItem(ctx context.Context, in *Item, opts ...grpc.CallOption) (*AddItemResponse, error) {
	out := new(AddItemResponse)
	err := grpc.Invoke(ctx, "/API/AddItem", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) UpdateItem(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*UpdateItemResponse, error) {
	out := new(UpdateItemResponse)
	err := grpc.Invoke(ctx, "/API/UpdateItem", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*DeleteItemResponse, error) {
	out := new(DeleteItemResponse)
	err := grpc.Invoke(ctx, "/API/DeleteItem", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	ListItem(context.Context, *ListItemRequest) (*ListItemResponse, error)
	GetItem(context.Context, *GetItemRequest) (*Item, error)
	AddItem(context.Context, *Item) (*AddItemResponse, error)
	UpdateItem(context.Context, *UpdateItemRequest) (*UpdateItemResponse, error)
	DeleteItem(context.Context, *DeleteItemRequest) (*DeleteItemResponse, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_ListItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/ListItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListItem(ctx, req.(*ListItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetItem(ctx, req.(*GetItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/AddItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).AddItem(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_UpdateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).UpdateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/UpdateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).UpdateItem(ctx, req.(*UpdateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_DeleteItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).DeleteItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/DeleteItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).DeleteItem(ctx, req.(*DeleteItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListItem",
			Handler:    _API_ListItem_Handler,
		},
		{
			MethodName: "GetItem",
			Handler:    _API_GetItem_Handler,
		},
		{
			MethodName: "AddItem",
			Handler:    _API_AddItem_Handler,
		},
		{
			MethodName: "UpdateItem",
			Handler:    _API_UpdateItem_Handler,
		},
		{
			MethodName: "DeleteItem",
			Handler:    _API_DeleteItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcb, 0xaa, 0xdb, 0x30,
	0x14, 0xc4, 0x8e, 0x9d, 0xc4, 0xe7, 0xc2, 0x4d, 0x72, 0x12, 0x8a, 0xeb, 0x6e, 0x5c, 0x77, 0x93,
	0x45, 0x51, 0x20, 0x25, 0xab, 0xae, 0x02, 0x85, 0x12, 0xe8, 0x22, 0x18, 0xfa, 0x01, 0x6e, 0x2c,
	0x8a, 0xc0, 0x0f, 0xd5, 0x52, 0x42, 0x37, 0xfd, 0x97, 0xfe, 0x5f, 0x7f, 0xa2, 0xe8, 0x28, 0x0f,
	0x3f, 0x4a, 0xef, 0x4e, 0x33, 0x67, 0xa4, 0x49, 0x66, 0x8e, 0x21, 0xc8, 0xa4, 0x60, 0xb2, 0xa9,
	0x75, 0x9d, 0xc4, 0xf0, 0xfc, 0x99, 0xeb, 0x83, 0xe6, 0x65, 0xca, 0x7f, 0x9c, 0xb9, 0xd2, 0xf8,
	0x0c, 0xae, 0xc8, 0x43, 0x27, 0x76, 0xd6, 0x41, 0xea, 0x8a, 0x3c, 0xf9, 0x08, 0xb3, 0x2f, 0x42,
	0x75, 0x24, 0x08, 0x9e, 0xcc, 0xbe, 0x73, 0x12, 0xf9, 0x29, 0x9d, 0x71, 0x05, 0x7e, 0x21, 0x4a,
	0xa1, 0x43, 0x97, 0x48, 0x0b, 0x92, 0x5f, 0x30, 0x7f, 0x5c, 0x56, 0xb2, 0xae, 0x14, 0xc7, 0x37,
	0xe0, 0x0b, 0xcd, 0x4b, 0x15, 0x3a, 0xf1, 0x68, 0xfd, 0xb4, 0xf5, 0x19, 0x4d, 0x2d, 0x67, 0x9e,
	0xd1, 0xb5, 0xce, 0x8a, 0xdb, 0x33, 0x04, 0x30, 0x82, 0xa9, 0x6c, 0xf8, 0xe5, 0x68, 0x4c, 0x47,
	0x34, 0xb8, 0x63, 0x33, 0xab, 0xf8, 0x4f, 0x4d, 0x33, 0xcf, 0xce, 0x6e, 0x38, 0x79, 0x0f, 0xb3,
	0x7d, 0x9e, 0x77, 0xdc, 0x5f, 0x83, 0x67, 0x9c, 0xe8, 0xb7, 0xdf, 0xcd, 0x89, 0x4a, 0x18, 0x2c,
	0xbe, 0xca, 0x3c, 0xd3, 0xbc, 0xfd, 0x5f, 0xff, 0xa3, 0xdf, 0x00, 0xb6, 0xf5, 0x2f, 0x1b, 0xbc,
	0x83, 0xc5, 0x27, 0x5e, 0xf0, 0xae, 0x41, 0x3f, 0xef, 0x15, 0x60, 0x5b, 0x64, 0x5f, 0x4d, 0x7e,
	0x3b, 0xe0, 0x19, 0xa2, 0x2f, 0x37, 0x5d, 0x54, 0x59, 0xc9, 0x29, 0xaf, 0x20, 0xa5, 0x33, 0x85,
	0x28, 0x74, 0x61, 0xb3, 0x0a, 0x52, 0x0b, 0x30, 0x86, 0xa7, 0x9c, 0xab, 0x53, 0x23, 0xa4, 0x16,
	0x75, 0x45, 0x59, 0x05, 0x69, 0x9b, 0x32, 0xf7, 0x64, 0x23, 0x4e, 0x3c, 0xf4, 0x6d, 0xf8, 0x04,
	0x8c, 0xa3, 0xbc, 0x84, 0x63, 0xa2, 0x5c, 0x79, 0xc1, 0x57, 0x30, 0x56, 0x3a, 0xd3, 0x67, 0x15,
	0x4e, 0x62, 0x67, 0x3d, 0x4d, 0xaf, 0x68, 0xfb, 0xc7, 0x81, 0xd1, 0xfe, 0x78, 0xc0, 0x0d, 0x4c,
	0x6f, 0x9d, 0xe3, 0x9c, 0xf5, 0x76, 0x27, 0x5a, 0xb0, 0xc1, 0x42, 0xbc, 0x85, 0xc9, 0x75, 0x07,
	0x71, 0xc6, 0xba, 0xdb, 0x18, 0xd9, 0xfc, 0x30, 0x81, 0xc9, 0xb5, 0x48, 0xb4, 0x4c, 0x34, 0x67,
	0xfd, 0x66, 0x77, 0x00, 0x8f, 0x3a, 0x10, 0xd9, 0xa0, 0xcb, 0x68, 0xc9, 0xfe, 0xd1, 0xd7, 0x0e,
	0xe0, 0x91, 0x37, 0x22, 0x1b, 0x34, 0x14, 0x2d, 0xd9, 0xb0, 0x90, 0x6f, 0x63, 0xfa, 0x7e, 0x3e,
	0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xab, 0xb1, 0x33, 0x4c, 0x03, 0x00, 0x00,
}
