// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal.proto

package runtime

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

type Cmd int32

const (
	Cmd_SEND      Cmd = 0
	Cmd_REPLY     Cmd = 1
	Cmd_BROADCAST Cmd = 2
)

var Cmd_name = map[int32]string{
	0: "SEND",
	1: "REPLY",
	2: "BROADCAST",
}
var Cmd_value = map[string]int32{
	"SEND":      0,
	"REPLY":     1,
	"BROADCAST": 2,
}

func (x Cmd) String() string {
	return proto.EnumName(Cmd_name, int32(x))
}
func (Cmd) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_internal_9239f88a3baccd72, []int{0}
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_9239f88a3baccd72, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Packet struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Event                string   `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	Id                   int32    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Payload              []byte   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_9239f88a3baccd72, []int{1}
}
func (m *Packet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet.Unmarshal(m, b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
}
func (dst *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(dst, src)
}
func (m *Packet) XXX_Size() int {
	return xxx_messageInfo_Packet.Size(m)
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

func (m *Packet) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Packet) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *Packet) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Packet) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type PipeRequest struct {
	Version              int32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Packet               *Packet  `protobuf:"bytes,3,opt,name=packet,proto3" json:"packet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PipeRequest) Reset()         { *m = PipeRequest{} }
func (m *PipeRequest) String() string { return proto.CompactTextString(m) }
func (*PipeRequest) ProtoMessage()    {}
func (*PipeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_9239f88a3baccd72, []int{2}
}
func (m *PipeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PipeRequest.Unmarshal(m, b)
}
func (m *PipeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PipeRequest.Marshal(b, m, deterministic)
}
func (dst *PipeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PipeRequest.Merge(dst, src)
}
func (m *PipeRequest) XXX_Size() int {
	return xxx_messageInfo_PipeRequest.Size(m)
}
func (m *PipeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PipeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PipeRequest proto.InternalMessageInfo

func (m *PipeRequest) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *PipeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PipeRequest) GetPacket() *Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

type PipeResponse struct {
	Version              int32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Cmd                  Cmd      `protobuf:"varint,2,opt,name=cmd,proto3,enum=runtime.Cmd" json:"cmd,omitempty"`
	Packet               *Packet  `protobuf:"bytes,3,opt,name=packet,proto3" json:"packet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PipeResponse) Reset()         { *m = PipeResponse{} }
func (m *PipeResponse) String() string { return proto.CompactTextString(m) }
func (*PipeResponse) ProtoMessage()    {}
func (*PipeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_9239f88a3baccd72, []int{3}
}
func (m *PipeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PipeResponse.Unmarshal(m, b)
}
func (m *PipeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PipeResponse.Marshal(b, m, deterministic)
}
func (dst *PipeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PipeResponse.Merge(dst, src)
}
func (m *PipeResponse) XXX_Size() int {
	return xxx_messageInfo_PipeResponse.Size(m)
}
func (m *PipeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PipeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PipeResponse proto.InternalMessageInfo

func (m *PipeResponse) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *PipeResponse) GetCmd() Cmd {
	if m != nil {
		return m.Cmd
	}
	return Cmd_SEND
}

func (m *PipeResponse) GetPacket() *Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

type SpecializeRequest struct {
	CodePath             string   `protobuf:"bytes,1,opt,name=codePath,proto3" json:"codePath,omitempty"`
	EntryPoint           string   `protobuf:"bytes,2,opt,name=entryPoint,proto3" json:"entryPoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpecializeRequest) Reset()         { *m = SpecializeRequest{} }
func (m *SpecializeRequest) String() string { return proto.CompactTextString(m) }
func (*SpecializeRequest) ProtoMessage()    {}
func (*SpecializeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_9239f88a3baccd72, []int{4}
}
func (m *SpecializeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecializeRequest.Unmarshal(m, b)
}
func (m *SpecializeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecializeRequest.Marshal(b, m, deterministic)
}
func (dst *SpecializeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecializeRequest.Merge(dst, src)
}
func (m *SpecializeRequest) XXX_Size() int {
	return xxx_messageInfo_SpecializeRequest.Size(m)
}
func (m *SpecializeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecializeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SpecializeRequest proto.InternalMessageInfo

func (m *SpecializeRequest) GetCodePath() string {
	if m != nil {
		return m.CodePath
	}
	return ""
}

func (m *SpecializeRequest) GetEntryPoint() string {
	if m != nil {
		return m.EntryPoint
	}
	return ""
}

type SpecializeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpecializeResponse) Reset()         { *m = SpecializeResponse{} }
func (m *SpecializeResponse) String() string { return proto.CompactTextString(m) }
func (*SpecializeResponse) ProtoMessage()    {}
func (*SpecializeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_9239f88a3baccd72, []int{5}
}
func (m *SpecializeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecializeResponse.Unmarshal(m, b)
}
func (m *SpecializeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecializeResponse.Marshal(b, m, deterministic)
}
func (dst *SpecializeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecializeResponse.Merge(dst, src)
}
func (m *SpecializeResponse) XXX_Size() int {
	return xxx_messageInfo_SpecializeResponse.Size(m)
}
func (m *SpecializeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecializeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SpecializeResponse proto.InternalMessageInfo

type HealthzResponse struct {
	StatusCode           int32    `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthzResponse) Reset()         { *m = HealthzResponse{} }
func (m *HealthzResponse) String() string { return proto.CompactTextString(m) }
func (*HealthzResponse) ProtoMessage()    {}
func (*HealthzResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_9239f88a3baccd72, []int{6}
}
func (m *HealthzResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthzResponse.Unmarshal(m, b)
}
func (m *HealthzResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthzResponse.Marshal(b, m, deterministic)
}
func (dst *HealthzResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthzResponse.Merge(dst, src)
}
func (m *HealthzResponse) XXX_Size() int {
	return xxx_messageInfo_HealthzResponse.Size(m)
}
func (m *HealthzResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthzResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthzResponse proto.InternalMessageInfo

func (m *HealthzResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "runtime.Empty")
	proto.RegisterType((*Packet)(nil), "runtime.Packet")
	proto.RegisterType((*PipeRequest)(nil), "runtime.PipeRequest")
	proto.RegisterType((*PipeResponse)(nil), "runtime.PipeResponse")
	proto.RegisterType((*SpecializeRequest)(nil), "runtime.SpecializeRequest")
	proto.RegisterType((*SpecializeResponse)(nil), "runtime.SpecializeResponse")
	proto.RegisterType((*HealthzResponse)(nil), "runtime.HealthzResponse")
	proto.RegisterEnum("runtime.Cmd", Cmd_name, Cmd_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InternalClient is the client API for Internal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InternalClient interface {
	Pipe(ctx context.Context, in *PipeRequest, opts ...grpc.CallOption) (Internal_PipeClient, error)
	Specialize(ctx context.Context, in *SpecializeRequest, opts ...grpc.CallOption) (*SpecializeResponse, error)
	Healthz(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthzResponse, error)
}

type internalClient struct {
	cc *grpc.ClientConn
}

func NewInternalClient(cc *grpc.ClientConn) InternalClient {
	return &internalClient{cc}
}

func (c *internalClient) Pipe(ctx context.Context, in *PipeRequest, opts ...grpc.CallOption) (Internal_PipeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Internal_serviceDesc.Streams[0], "/runtime.Internal/Pipe", opts...)
	if err != nil {
		return nil, err
	}
	x := &internalPipeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Internal_PipeClient interface {
	Recv() (*PipeResponse, error)
	grpc.ClientStream
}

type internalPipeClient struct {
	grpc.ClientStream
}

func (x *internalPipeClient) Recv() (*PipeResponse, error) {
	m := new(PipeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *internalClient) Specialize(ctx context.Context, in *SpecializeRequest, opts ...grpc.CallOption) (*SpecializeResponse, error) {
	out := new(SpecializeResponse)
	err := c.cc.Invoke(ctx, "/runtime.Internal/Specialize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalClient) Healthz(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthzResponse, error) {
	out := new(HealthzResponse)
	err := c.cc.Invoke(ctx, "/runtime.Internal/Healthz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalServer is the server API for Internal service.
type InternalServer interface {
	Pipe(*PipeRequest, Internal_PipeServer) error
	Specialize(context.Context, *SpecializeRequest) (*SpecializeResponse, error)
	Healthz(context.Context, *Empty) (*HealthzResponse, error)
}

func RegisterInternalServer(s *grpc.Server, srv InternalServer) {
	s.RegisterService(&_Internal_serviceDesc, srv)
}

func _Internal_Pipe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PipeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InternalServer).Pipe(m, &internalPipeServer{stream})
}

type Internal_PipeServer interface {
	Send(*PipeResponse) error
	grpc.ServerStream
}

type internalPipeServer struct {
	grpc.ServerStream
}

func (x *internalPipeServer) Send(m *PipeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Internal_Specialize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpecializeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServer).Specialize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runtime.Internal/Specialize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServer).Specialize(ctx, req.(*SpecializeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Internal_Healthz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServer).Healthz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runtime.Internal/Healthz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServer).Healthz(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Internal_serviceDesc = grpc.ServiceDesc{
	ServiceName: "runtime.Internal",
	HandlerType: (*InternalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Specialize",
			Handler:    _Internal_Specialize_Handler,
		},
		{
			MethodName: "Healthz",
			Handler:    _Internal_Healthz_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Pipe",
			Handler:       _Internal_Pipe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal.proto",
}

func init() { proto.RegisterFile("internal.proto", fileDescriptor_internal_9239f88a3baccd72) }

var fileDescriptor_internal_9239f88a3baccd72 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x5f, 0xab, 0xda, 0x30,
	0x14, 0xb7, 0xd5, 0x5a, 0x3d, 0xba, 0xea, 0x82, 0x83, 0xd2, 0x81, 0x48, 0x5f, 0xe6, 0xf6, 0x20,
	0x9b, 0x63, 0xec, 0xd9, 0x55, 0xd9, 0x06, 0x63, 0x96, 0xb8, 0x97, 0xbd, 0x8c, 0x65, 0x6d, 0xc0,
	0xb0, 0xb6, 0x89, 0x6d, 0x14, 0xf4, 0xeb, 0xdd, 0x2f, 0x76, 0x69, 0x1b, 0xa3, 0x78, 0x2f, 0x97,
	0xfb, 0xf8, 0x3b, 0x39, 0xfc, 0xfe, 0xe5, 0x80, 0xc3, 0x32, 0x49, 0xf3, 0x8c, 0x24, 0x33, 0x91,
	0x73, 0xc9, 0x91, 0x9d, 0xef, 0x33, 0xc9, 0x52, 0xea, 0xdb, 0x60, 0xad, 0x52, 0x21, 0x8f, 0xfe,
	0x1f, 0x68, 0x87, 0x24, 0xfa, 0x4f, 0x25, 0x1a, 0x81, 0x25, 0xb9, 0x60, 0x91, 0x6b, 0x4c, 0x8c,
	0x69, 0x17, 0xd7, 0xa0, 0x9c, 0xd2, 0x03, 0xcd, 0xa4, 0x6b, 0xd6, 0xd3, 0x0a, 0x20, 0x07, 0x4c,
	0x16, 0xbb, 0xcd, 0x89, 0x31, 0xb5, 0xb0, 0xc9, 0x62, 0xe4, 0x82, 0x2d, 0xc8, 0x31, 0xe1, 0x24,
	0x76, 0x5b, 0x13, 0x63, 0xda, 0xc7, 0x67, 0xe8, 0xff, 0x85, 0x5e, 0xc8, 0x04, 0xc5, 0x74, 0xb7,
	0xa7, 0x85, 0x2c, 0x17, 0x0f, 0x34, 0x2f, 0x18, 0xcf, 0x2a, 0x19, 0x0b, 0x9f, 0xa1, 0xa2, 0xac,
	0x55, 0x4a, 0xca, 0x37, 0xd0, 0x16, 0x95, 0xb1, 0x4a, 0xa6, 0x37, 0x1f, 0xcc, 0x94, 0xf7, 0x59,
	0xed, 0x17, 0xab, 0x67, 0x7f, 0x07, 0xfd, 0x5a, 0xa1, 0x10, 0x3c, 0x2b, 0xe8, 0x13, 0x12, 0x63,
	0x68, 0x46, 0x69, 0xad, 0xe1, 0xcc, 0xfb, 0x9a, 0x2f, 0x48, 0x63, 0x5c, 0x3e, 0x3c, 0x5f, 0x72,
	0x0d, 0x2f, 0x37, 0x82, 0x46, 0x8c, 0x24, 0xec, 0xa4, 0xa3, 0x79, 0xd0, 0x89, 0x78, 0x4c, 0x43,
	0x22, 0xb7, 0xaa, 0x42, 0x8d, 0xd1, 0x18, 0x80, 0x66, 0x32, 0x3f, 0x86, 0x9c, 0xe9, 0x2a, 0xaf,
	0x26, 0xfe, 0x08, 0xd0, 0x35, 0x61, 0x9d, 0xc4, 0xff, 0x00, 0x83, 0x6f, 0x94, 0x24, 0x72, 0x7b,
	0xd2, 0xe1, 0xc6, 0x00, 0x85, 0x24, 0x72, 0x5f, 0x04, 0x3c, 0xa6, 0x2a, 0xdf, 0xd5, 0xe4, 0xdd,
	0x5b, 0x68, 0x06, 0x69, 0x8c, 0x3a, 0xd0, 0xda, 0xac, 0x7e, 0x2e, 0x87, 0x0d, 0xd4, 0x05, 0x0b,
	0xaf, 0xc2, 0x1f, 0xbf, 0x87, 0x06, 0x7a, 0x01, 0xdd, 0x2f, 0x78, 0xbd, 0x58, 0x06, 0x8b, 0xcd,
	0xaf, 0xa1, 0x39, 0xbf, 0x33, 0xa0, 0xf3, 0x5d, 0x9d, 0x07, 0xfa, 0x0c, 0xad, 0xb2, 0x44, 0x34,
	0xba, 0x44, 0xbe, 0xfc, 0x9a, 0xf7, 0xea, 0x66, 0xaa, 0xfc, 0x35, 0xde, 0x1b, 0xe8, 0x2b, 0xc0,
	0xc5, 0x39, 0xf2, 0xf4, 0xe2, 0x83, 0x7e, 0xbc, 0xd7, 0x8f, 0xbe, 0x9d, 0xa9, 0xd0, 0x27, 0xb0,
	0x55, 0x58, 0xe4, 0xe8, 0xcd, 0xea, 0x46, 0x3d, 0x57, 0xe3, 0x9b, 0x3a, 0xfc, 0xc6, 0xbf, 0x76,
	0x75, 0xd8, 0x1f, 0xef, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0c, 0xf4, 0xe2, 0x1d, 0xea, 0x02, 0x00,
	0x00,
}
