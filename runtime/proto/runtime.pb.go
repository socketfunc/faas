// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/runtime.proto

package runtime

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto1 "github.com/socketfunc/faas/store/proto"

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
	Cmd_STREAM Cmd = 0
	Cmd_STORE  Cmd = 1
)

var Cmd_name = map[int32]string{
	0: "STREAM",
	1: "STORE",
}
var Cmd_value = map[string]int32{
	"STREAM": 0,
	"STORE":  1,
}

func (x Cmd) String() string {
	return proto.EnumName(Cmd_name, int32(x))
}
func (Cmd) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_runtime_c09f5d517869adfa, []int{0}
}

type Store_Cmd int32

const (
	Store_Cmd_GET    Store_Cmd = 0
	Store_Cmd_PUT    Store_Cmd = 1
	Store_Cmd_MODIFY Store_Cmd = 2
	Store_Cmd_DEL    Store_Cmd = 3
)

var Store_Cmd_name = map[int32]string{
	0: "GET",
	1: "PUT",
	2: "MODIFY",
	3: "DEL",
}
var Store_Cmd_value = map[string]int32{
	"GET":    0,
	"PUT":    1,
	"MODIFY": 2,
	"DEL":    3,
}

func (x Store_Cmd) String() string {
	return proto.EnumName(Store_Cmd_name, int32(x))
}
func (Store_Cmd) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_runtime_c09f5d517869adfa, []int{1}
}

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}
var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_runtime_c09f5d517869adfa, []int{7, 0}
}

type Receive struct {
	Cmd                  Cmd            `protobuf:"varint,1,opt,name=cmd,proto3,enum=runtime.Cmd" json:"cmd,omitempty"`
	StreamRequest        *StreamRequest `protobuf:"bytes,2,opt,name=stream_request,json=streamRequest,proto3" json:"stream_request,omitempty"`
	StoreResponse        *StoreResponse `protobuf:"bytes,3,opt,name=store_response,json=storeResponse,proto3" json:"store_response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Receive) Reset()         { *m = Receive{} }
func (m *Receive) String() string { return proto.CompactTextString(m) }
func (*Receive) ProtoMessage()    {}
func (*Receive) Descriptor() ([]byte, []int) {
	return fileDescriptor_runtime_c09f5d517869adfa, []int{0}
}
func (m *Receive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Receive.Unmarshal(m, b)
}
func (m *Receive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Receive.Marshal(b, m, deterministic)
}
func (dst *Receive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receive.Merge(dst, src)
}
func (m *Receive) XXX_Size() int {
	return xxx_messageInfo_Receive.Size(m)
}
func (m *Receive) XXX_DiscardUnknown() {
	xxx_messageInfo_Receive.DiscardUnknown(m)
}

var xxx_messageInfo_Receive proto.InternalMessageInfo

func (m *Receive) GetCmd() Cmd {
	if m != nil {
		return m.Cmd
	}
	return Cmd_STREAM
}

func (m *Receive) GetStreamRequest() *StreamRequest {
	if m != nil {
		return m.StreamRequest
	}
	return nil
}

func (m *Receive) GetStoreResponse() *StoreResponse {
	if m != nil {
		return m.StoreResponse
	}
	return nil
}

type Send struct {
	Cmd                  Cmd           `protobuf:"varint,1,opt,name=cmd,proto3,enum=runtime.Cmd" json:"cmd,omitempty"`
	StreamSend           *StreamSend   `protobuf:"bytes,2,opt,name=stream_send,json=streamSend,proto3" json:"stream_send,omitempty"`
	StoreRequest         *StoreRequest `protobuf:"bytes,3,opt,name=store_request,json=storeRequest,proto3" json:"store_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Send) Reset()         { *m = Send{} }
func (m *Send) String() string { return proto.CompactTextString(m) }
func (*Send) ProtoMessage()    {}
func (*Send) Descriptor() ([]byte, []int) {
	return fileDescriptor_runtime_c09f5d517869adfa, []int{1}
}
func (m *Send) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Send.Unmarshal(m, b)
}
func (m *Send) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Send.Marshal(b, m, deterministic)
}
func (dst *Send) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Send.Merge(dst, src)
}
func (m *Send) XXX_Size() int {
	return xxx_messageInfo_Send.Size(m)
}
func (m *Send) XXX_DiscardUnknown() {
	xxx_messageInfo_Send.DiscardUnknown(m)
}

var xxx_messageInfo_Send proto.InternalMessageInfo

func (m *Send) GetCmd() Cmd {
	if m != nil {
		return m.Cmd
	}
	return Cmd_STREAM
}

func (m *Send) GetStreamSend() *StreamSend {
	if m != nil {
		return m.StreamSend
	}
	return nil
}

func (m *Send) GetStoreRequest() *StoreRequest {
	if m != nil {
		return m.StoreRequest
	}
	return nil
}

type StreamRequest struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Event                string   `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_runtime_c09f5d517869adfa, []int{2}
}
func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (dst *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(dst, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *StreamRequest) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *StreamRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type StreamSend struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Event                string   `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamSend) Reset()         { *m = StreamSend{} }
func (m *StreamSend) String() string { return proto.CompactTextString(m) }
func (*StreamSend) ProtoMessage()    {}
func (*StreamSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_runtime_c09f5d517869adfa, []int{3}
}
func (m *StreamSend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamSend.Unmarshal(m, b)
}
func (m *StreamSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamSend.Marshal(b, m, deterministic)
}
func (dst *StreamSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamSend.Merge(dst, src)
}
func (m *StreamSend) XXX_Size() int {
	return xxx_messageInfo_StreamSend.Size(m)
}
func (m *StreamSend) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamSend.DiscardUnknown(m)
}

var xxx_messageInfo_StreamSend proto.InternalMessageInfo

func (m *StreamSend) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *StreamSend) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *StreamSend) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type StoreRequest struct {
	Cmd                  Store_Cmd        `protobuf:"varint,1,opt,name=cmd,proto3,enum=runtime.Store_Cmd" json:"cmd,omitempty"`
	Key                  string           `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Entity               *proto1.Entity   `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity,omitempty"`
	Filters              []*proto1.Filter `protobuf:"bytes,4,rep,name=filters,proto3" json:"filters,omitempty"`
	Updates              []*proto1.Update `protobuf:"bytes,5,rep,name=updates,proto3" json:"updates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *StoreRequest) Reset()         { *m = StoreRequest{} }
func (m *StoreRequest) String() string { return proto.CompactTextString(m) }
func (*StoreRequest) ProtoMessage()    {}
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_runtime_c09f5d517869adfa, []int{4}
}
func (m *StoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreRequest.Unmarshal(m, b)
}
func (m *StoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreRequest.Marshal(b, m, deterministic)
}
func (dst *StoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreRequest.Merge(dst, src)
}
func (m *StoreRequest) XXX_Size() int {
	return xxx_messageInfo_StoreRequest.Size(m)
}
func (m *StoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoreRequest proto.InternalMessageInfo

func (m *StoreRequest) GetCmd() Store_Cmd {
	if m != nil {
		return m.Cmd
	}
	return Store_Cmd_GET
}

func (m *StoreRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StoreRequest) GetEntity() *proto1.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *StoreRequest) GetFilters() []*proto1.Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *StoreRequest) GetUpdates() []*proto1.Update {
	if m != nil {
		return m.Updates
	}
	return nil
}

type StoreResponse struct {
	Cmd                  Store_Cmd      `protobuf:"varint,1,opt,name=cmd,proto3,enum=runtime.Store_Cmd" json:"cmd,omitempty"`
	Successful           bool           `protobuf:"varint,2,opt,name=successful,proto3" json:"successful,omitempty"`
	Entity               *proto1.Entity `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StoreResponse) Reset()         { *m = StoreResponse{} }
func (m *StoreResponse) String() string { return proto.CompactTextString(m) }
func (*StoreResponse) ProtoMessage()    {}
func (*StoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_runtime_c09f5d517869adfa, []int{5}
}
func (m *StoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreResponse.Unmarshal(m, b)
}
func (m *StoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreResponse.Marshal(b, m, deterministic)
}
func (dst *StoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreResponse.Merge(dst, src)
}
func (m *StoreResponse) XXX_Size() int {
	return xxx_messageInfo_StoreResponse.Size(m)
}
func (m *StoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StoreResponse proto.InternalMessageInfo

func (m *StoreResponse) GetCmd() Store_Cmd {
	if m != nil {
		return m.Cmd
	}
	return Store_Cmd_GET
}

func (m *StoreResponse) GetSuccessful() bool {
	if m != nil {
		return m.Successful
	}
	return false
}

func (m *StoreResponse) GetEntity() *proto1.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type HealthCheckRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckRequest) Reset()         { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()    {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_runtime_c09f5d517869adfa, []int{6}
}
func (m *HealthCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequest.Unmarshal(m, b)
}
func (m *HealthCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequest.Marshal(b, m, deterministic)
}
func (dst *HealthCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequest.Merge(dst, src)
}
func (m *HealthCheckRequest) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequest.Size(m)
}
func (m *HealthCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequest proto.InternalMessageInfo

func (m *HealthCheckRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type HealthCheckResponse struct {
	Status               HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=runtime.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *HealthCheckResponse) Reset()         { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()    {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_runtime_c09f5d517869adfa, []int{7}
}
func (m *HealthCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckResponse.Unmarshal(m, b)
}
func (m *HealthCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckResponse.Marshal(b, m, deterministic)
}
func (dst *HealthCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckResponse.Merge(dst, src)
}
func (m *HealthCheckResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckResponse.Size(m)
}
func (m *HealthCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckResponse proto.InternalMessageInfo

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

func init() {
	proto.RegisterType((*Receive)(nil), "runtime.Receive")
	proto.RegisterType((*Send)(nil), "runtime.Send")
	proto.RegisterType((*StreamRequest)(nil), "runtime.StreamRequest")
	proto.RegisterType((*StreamSend)(nil), "runtime.StreamSend")
	proto.RegisterType((*StoreRequest)(nil), "runtime.StoreRequest")
	proto.RegisterType((*StoreResponse)(nil), "runtime.StoreResponse")
	proto.RegisterType((*HealthCheckRequest)(nil), "runtime.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "runtime.HealthCheckResponse")
	proto.RegisterEnum("runtime.Cmd", Cmd_name, Cmd_value)
	proto.RegisterEnum("runtime.Store_Cmd", Store_Cmd_name, Store_Cmd_value)
	proto.RegisterEnum("runtime.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RuntimeClient is the client API for Runtime service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RuntimeClient interface {
	Stream(ctx context.Context, opts ...grpc.CallOption) (Runtime_StreamClient, error)
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type runtimeClient struct {
	cc *grpc.ClientConn
}

func NewRuntimeClient(cc *grpc.ClientConn) RuntimeClient {
	return &runtimeClient{cc}
}

func (c *runtimeClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Runtime_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Runtime_serviceDesc.Streams[0], "/runtime.Runtime/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &runtimeStreamClient{stream}
	return x, nil
}

type Runtime_StreamClient interface {
	Send(*Receive) error
	Recv() (*Send, error)
	grpc.ClientStream
}

type runtimeStreamClient struct {
	grpc.ClientStream
}

func (x *runtimeStreamClient) Send(m *Receive) error {
	return x.ClientStream.SendMsg(m)
}

func (x *runtimeStreamClient) Recv() (*Send, error) {
	m := new(Send)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *runtimeClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/runtime.Runtime/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RuntimeServer is the server API for Runtime service.
type RuntimeServer interface {
	Stream(Runtime_StreamServer) error
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

func RegisterRuntimeServer(s *grpc.Server, srv RuntimeServer) {
	s.RegisterService(&_Runtime_serviceDesc, srv)
}

func _Runtime_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RuntimeServer).Stream(&runtimeStreamServer{stream})
}

type Runtime_StreamServer interface {
	Send(*Send) error
	Recv() (*Receive, error)
	grpc.ServerStream
}

type runtimeStreamServer struct {
	grpc.ServerStream
}

func (x *runtimeStreamServer) Send(m *Send) error {
	return x.ServerStream.SendMsg(m)
}

func (x *runtimeStreamServer) Recv() (*Receive, error) {
	m := new(Receive)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Runtime_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runtime.Runtime/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeServer).HealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Runtime_serviceDesc = grpc.ServiceDesc{
	ServiceName: "runtime.Runtime",
	HandlerType: (*RuntimeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _Runtime_HealthCheck_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Runtime_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/runtime.proto",
}

func init() { proto.RegisterFile("proto/runtime.proto", fileDescriptor_runtime_c09f5d517869adfa) }

var fileDescriptor_runtime_c09f5d517869adfa = []byte{
	// 594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xcd, 0xd4, 0x6d, 0xfc, 0xe5, 0x26, 0xe9, 0x67, 0x4d, 0x01, 0x59, 0xa5, 0xaa, 0x2a, 0x0b,
	0x44, 0xd5, 0x45, 0x82, 0x02, 0xab, 0x4a, 0x2c, 0xa0, 0x4d, 0x4b, 0x81, 0x26, 0x68, 0x9c, 0x80,
	0x58, 0x45, 0xc6, 0xbe, 0xa5, 0x56, 0x13, 0x3b, 0x78, 0xc6, 0x95, 0x22, 0xb1, 0xe6, 0x0d, 0x90,
	0x78, 0x03, 0x1e, 0x83, 0x57, 0x43, 0xf3, 0x63, 0xd7, 0x29, 0x15, 0x74, 0xc1, 0x6e, 0xce, 0x3d,
	0xf7, 0x5e, 0x9f, 0x73, 0x3c, 0x1a, 0xd8, 0x98, 0x67, 0xa9, 0x48, 0xbb, 0x59, 0x9e, 0x88, 0x78,
	0x86, 0x1d, 0x85, 0xa8, 0x6d, 0xe0, 0xe6, 0xd6, 0x59, 0x10, 0xf0, 0x2e, 0x17, 0x69, 0x86, 0x5d,
	0xdd, 0xa8, 0xce, 0xba, 0xcd, 0xfb, 0x41, 0xc0, 0x66, 0x18, 0x62, 0x7c, 0x89, 0x74, 0x1b, 0xac,
	0x70, 0x16, 0xb9, 0x64, 0x87, 0xec, 0xae, 0xf7, 0x5a, 0x9d, 0x62, 0xdf, 0xc1, 0x2c, 0x62, 0x92,
	0xa0, 0xcf, 0x60, 0x9d, 0x8b, 0x0c, 0x83, 0xd9, 0x24, 0xc3, 0xcf, 0x39, 0x72, 0xe1, 0xae, 0xec,
	0x90, 0xdd, 0x66, 0xef, 0x5e, 0xd9, 0xea, 0x2b, 0x9a, 0x69, 0x96, 0xb5, 0x79, 0x15, 0xea, 0xf1,
	0x34, 0xc3, 0x49, 0x86, 0x7c, 0x9e, 0x26, 0x1c, 0x5d, 0xeb, 0xb7, 0xf1, 0x34, 0x43, 0x66, 0x58,
	0x39, 0x5e, 0x81, 0xde, 0x77, 0x02, 0xab, 0x3e, 0x26, 0xd1, 0x5f, 0x65, 0x3e, 0x85, 0xa6, 0x91,
	0xc9, 0x31, 0x89, 0x8c, 0xc6, 0x8d, 0x6b, 0x1a, 0xe5, 0x26, 0x06, 0xbc, 0x3c, 0xd3, 0x7d, 0x68,
	0x17, 0xea, 0xb4, 0x37, 0x2d, 0xee, 0xee, 0x75, 0x71, 0xda, 0x5a, 0x8b, 0x57, 0x90, 0x37, 0x86,
	0xf6, 0x92, 0x73, 0x7a, 0x07, 0xd6, 0x44, 0x3a, 0x8f, 0x43, 0x25, 0xb2, 0xc1, 0x34, 0x90, 0x55,
	0xbc, 0xc4, 0x44, 0xc7, 0xd6, 0x60, 0x1a, 0x50, 0x17, 0xec, 0x79, 0xb0, 0x98, 0xa6, 0x41, 0xa4,
	0x3e, 0xd9, 0x62, 0x05, 0xf4, 0x18, 0xc0, 0x95, 0xd8, 0x7f, 0xb4, 0xf3, 0x27, 0x81, 0x56, 0xd5,
	0x09, 0x7d, 0x50, 0x4d, 0x93, 0x2e, 0xbb, 0x9d, 0x94, 0x99, 0x3a, 0x60, 0x5d, 0xe0, 0xc2, 0x7c,
	0x44, 0x1e, 0xe9, 0x43, 0xa8, 0x63, 0x22, 0x62, 0xb1, 0x30, 0x41, 0xb5, 0x3b, 0xfa, 0x5a, 0xf5,
	0x55, 0x91, 0x19, 0x92, 0x3e, 0x02, 0xfb, 0x2c, 0x9e, 0x0a, 0xcc, 0xb8, 0xbb, 0xba, 0x63, 0x55,
	0xfa, 0x8e, 0x54, 0x95, 0x15, 0xac, 0x6c, 0xcc, 0xe7, 0x51, 0x20, 0x90, 0xbb, 0x6b, 0x4b, 0x8d,
	0x63, 0x55, 0x65, 0x05, 0xeb, 0x7d, 0x91, 0x61, 0x57, 0x2e, 0xc6, 0x2d, 0x1d, 0x6c, 0x03, 0xf0,
	0x3c, 0x0c, 0x91, 0xf3, 0xb3, 0x7c, 0xaa, 0x8c, 0xfc, 0xc7, 0x2a, 0x95, 0x5b, 0xfa, 0xf1, 0x3a,
	0x40, 0x5f, 0x62, 0x30, 0x15, 0xe7, 0x07, 0xe7, 0x18, 0x5e, 0x14, 0x21, 0xba, 0x60, 0x73, 0xcc,
	0x2e, 0xe3, 0x10, 0xcd, 0xdf, 0x29, 0xa0, 0xf7, 0x8d, 0xc0, 0xc6, 0xd2, 0x80, 0x11, 0xfd, 0x02,
	0xea, 0x5c, 0x04, 0x22, 0xe7, 0x46, 0xf7, 0x5e, 0xa9, 0xfb, 0x86, 0xee, 0x8e, 0x2f, 0xb7, 0x25,
	0x9f, 0x7c, 0x35, 0xc1, 0xcc, 0xa4, 0xb7, 0x0f, 0xed, 0x25, 0x82, 0x36, 0xc1, 0x1e, 0x0f, 0x5e,
	0x0f, 0x86, 0xef, 0x07, 0x4e, 0x4d, 0x02, 0xbf, 0xcf, 0xde, 0x9d, 0x0c, 0x8e, 0x1d, 0x42, 0xff,
	0x87, 0xe6, 0x60, 0x38, 0x9a, 0x14, 0x85, 0x95, 0xbd, 0x2d, 0xb0, 0x0e, 0x66, 0x11, 0x05, 0xa8,
	0xfb, 0x23, 0xd6, 0x7f, 0x7e, 0xea, 0xd4, 0x68, 0x03, 0xd6, 0xfc, 0xd1, 0x90, 0xf5, 0x1d, 0xb2,
	0xd7, 0x83, 0x46, 0x19, 0x1f, 0xb5, 0xc1, 0x3a, 0xee, 0x8f, 0x9c, 0x9a, 0x3c, 0xbc, 0x1d, 0x8f,
	0x1c, 0x22, 0xa7, 0x4e, 0x87, 0x87, 0x27, 0x47, 0x1f, 0x9c, 0x15, 0x59, 0x3c, 0xec, 0xbf, 0x71,
	0xac, 0xde, 0x57, 0xf9, 0x92, 0x68, 0x0f, 0xb4, 0x0b, 0x75, 0x7d, 0x73, 0xa9, 0x53, 0xfa, 0x32,
	0xaf, 0xcc, 0x66, 0xfb, 0xea, 0x0f, 0x61, 0x12, 0x79, 0xb5, 0x5d, 0xf2, 0x98, 0xd0, 0x57, 0xd0,
	0xac, 0xf8, 0xa6, 0xf7, 0x6f, 0x4e, 0x43, 0x85, 0xbd, 0xb9, 0xf5, 0xa7, 0xa8, 0xbc, 0xda, 0xc7,
	0xba, 0x7a, 0xd9, 0x9e, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x10, 0x22, 0x15, 0x16, 0x17, 0x05,
	0x00, 0x00,
}
