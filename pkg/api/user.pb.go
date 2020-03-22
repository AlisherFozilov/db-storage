// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package storage

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

type Messages struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 int64    `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	SenderId             int64    `protobuf:"varint,3,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	ReceiverId           int64    `protobuf:"varint,4,opt,name=receiver_id,json=receiverId,proto3" json:"receiver_id,omitempty"`
	Data                 string   `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	Timestamp            int64    `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Removed              bool     `protobuf:"varint,7,opt,name=removed,proto3" json:"removed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Messages) Reset()         { *m = Messages{} }
func (m *Messages) String() string { return proto.CompactTextString(m) }
func (*Messages) ProtoMessage()    {}
func (*Messages) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *Messages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Messages.Unmarshal(m, b)
}
func (m *Messages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Messages.Marshal(b, m, deterministic)
}
func (m *Messages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Messages.Merge(m, src)
}
func (m *Messages) XXX_Size() int {
	return xxx_messageInfo_Messages.Size(m)
}
func (m *Messages) XXX_DiscardUnknown() {
	xxx_messageInfo_Messages.DiscardUnknown(m)
}

var xxx_messageInfo_Messages proto.InternalMessageInfo

func (m *Messages) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Messages) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Messages) GetSenderId() int64 {
	if m != nil {
		return m.SenderId
	}
	return 0
}

func (m *Messages) GetReceiverId() int64 {
	if m != nil {
		return m.ReceiverId
	}
	return 0
}

func (m *Messages) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Messages) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Messages) GetRemoved() bool {
	if m != nil {
		return m.Removed
	}
	return false
}

type Response struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
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

func (m *Response) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*Messages)(nil), "storage.Messages")
	proto.RegisterType((*Response)(nil), "storage.Response")
}

func init() {
	proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf)
}

var fileDescriptor_116e343673f7ffaf = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0x4d, 0xbb, 0xf6, 0xcf, 0x28, 0x82, 0x73, 0x90, 0xa0, 0x82, 0xa5, 0xa7, 0x9e, 0x7a,
	0x50, 0xf0, 0xee, 0x71, 0x0f, 0x5e, 0xb2, 0x1f, 0x40, 0xa2, 0x19, 0x96, 0x1c, 0xba, 0x29, 0x99,
	0x6c, 0xc1, 0x6f, 0xe6, 0xc7, 0x13, 0xa6, 0x9b, 0xdd, 0xdb, 0xcc, 0xef, 0x91, 0x97, 0x79, 0x0f,
	0xe0, 0xc8, 0x14, 0xc7, 0x39, 0x86, 0x14, 0xb0, 0xe6, 0x14, 0xa2, 0xdd, 0x53, 0xff, 0xa7, 0xa0,
	0xf9, 0x24, 0x66, 0xbb, 0x27, 0xc6, 0x3b, 0x28, 0xbc, 0xd3, 0xaa, 0x53, 0x43, 0x6b, 0x0a, 0xef,
	0x10, 0x61, 0x93, 0x7e, 0x67, 0xd2, 0x45, 0xa7, 0x86, 0xd2, 0xc8, 0x8c, 0x4f, 0xd0, 0x32, 0x1d,
	0x1c, 0xc5, 0x2f, 0xef, 0x74, 0x29, 0x42, 0xb3, 0x82, 0xad, 0xc3, 0x17, 0xb8, 0x89, 0xf4, 0x43,
	0x7e, 0x59, 0xe5, 0x8d, 0xc8, 0x90, 0xd1, 0x56, 0x1c, 0x9d, 0x4d, 0x56, 0x5f, 0xcb, 0x1f, 0x32,
	0xe3, 0x33, 0xb4, 0xc9, 0x4f, 0xc4, 0xc9, 0x4e, 0xb3, 0xae, 0xe4, 0xc9, 0x05, 0xa0, 0x86, 0x3a,
	0xd2, 0x14, 0x16, 0x72, 0xba, 0xee, 0xd4, 0xd0, 0x98, 0xbc, 0xf6, 0x3d, 0x34, 0x86, 0x78, 0x0e,
	0x07, 0x26, 0x7c, 0x80, 0x8a, 0x93, 0x4d, 0x47, 0x96, 0xeb, 0x4b, 0x73, 0xda, 0x5e, 0x3f, 0xa0,
	0xde, 0xad, 0x49, 0xf1, 0x1d, 0x6e, 0x77, 0x76, 0xa1, 0x73, 0xd8, 0xfb, 0xf1, 0xd4, 0xc1, 0x98,
	0xd1, 0xe3, 0x05, 0x65, 0xe3, 0xfe, 0xea, 0xbb, 0x92, 0xc6, 0xde, 0xfe, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x3d, 0xdd, 0x76, 0x99, 0x3f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StorageClient interface {
	SaveMessages(ctx context.Context, in *Messages, opts ...grpc.CallOption) (*Response, error)
}

type storageClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageClient(cc grpc.ClientConnInterface) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) SaveMessages(ctx context.Context, in *Messages, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/storage.Storage/SaveMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServer is the server API for Storage service.
type StorageServer interface {
	SaveMessages(context.Context, *Messages) (*Response, error)
}

// UnimplementedStorageServer can be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (*UnimplementedStorageServer) SaveMessages(ctx context.Context, req *Messages) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveMessages not implemented")
}

func RegisterStorageServer(s *grpc.Server, srv StorageServer) {
	s.RegisterService(&_Storage_serviceDesc, srv)
}

func _Storage_SaveMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Messages)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).SaveMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.Storage/SaveMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).SaveMessages(ctx, req.(*Messages))
	}
	return interceptor(ctx, in, info, handler)
}

var _Storage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "storage.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveMessages",
			Handler:    _Storage_SaveMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}