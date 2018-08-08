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
	Cmd_Stream Cmd = 0
	Cmd_Store  Cmd = 1
)

var Cmd_name = map[int32]string{
	0: "Stream",
	1: "Store",
}
var Cmd_value = map[string]int32{
	"Stream": 0,
	"Store":  1,
}

func (x Cmd) String() string {
	return proto.EnumName(Cmd_name, int32(x))
}
func (Cmd) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_runtime_aadcca0a1b3e6e7d, []int{0}
}

type Store_Cmd int32

const (
	Store_Cmd_Get    Store_Cmd = 0
	Store_Cmd_Put    Store_Cmd = 1
	Store_Cmd_Modify Store_Cmd = 2
	Store_Cmd_Del    Store_Cmd = 3
)

var Store_Cmd_name = map[int32]string{
	0: "Get",
	1: "Put",
	2: "Modify",
	3: "Del",
}
var Store_Cmd_value = map[string]int32{
	"Get":    0,
	"Put":    1,
	"Modify": 2,
	"Del":    3,
}

func (x Store_Cmd) String() string {
	return proto.EnumName(Store_Cmd_name, int32(x))
}
func (Store_Cmd) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_runtime_aadcca0a1b3e6e7d, []int{1}
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
	return fileDescriptor_runtime_aadcca0a1b3e6e7d, []int{0}
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
	return Cmd_Stream
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
	return fileDescriptor_runtime_aadcca0a1b3e6e7d, []int{1}
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
	return Cmd_Stream
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
	return fileDescriptor_runtime_aadcca0a1b3e6e7d, []int{2}
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
	return fileDescriptor_runtime_aadcca0a1b3e6e7d, []int{3}
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
	return fileDescriptor_runtime_aadcca0a1b3e6e7d, []int{4}
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
	return Store_Cmd_Get
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
	return fileDescriptor_runtime_aadcca0a1b3e6e7d, []int{5}
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
	return Store_Cmd_Get
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

func init() {
	proto.RegisterType((*Receive)(nil), "runtime.Receive")
	proto.RegisterType((*Send)(nil), "runtime.Send")
	proto.RegisterType((*StreamRequest)(nil), "runtime.StreamRequest")
	proto.RegisterType((*StreamSend)(nil), "runtime.StreamSend")
	proto.RegisterType((*StoreRequest)(nil), "runtime.StoreRequest")
	proto.RegisterType((*StoreResponse)(nil), "runtime.StoreResponse")
	proto.RegisterEnum("runtime.Cmd", Cmd_name, Cmd_value)
	proto.RegisterEnum("runtime.Store_Cmd", Store_Cmd_name, Store_Cmd_value)
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

// RuntimeServer is the server API for Runtime service.
type RuntimeServer interface {
	Stream(Runtime_StreamServer) error
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

var _Runtime_serviceDesc = grpc.ServiceDesc{
	ServiceName: "runtime.Runtime",
	HandlerType: (*RuntimeServer)(nil),
	Methods:     []grpc.MethodDesc{},
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

func init() { proto.RegisterFile("proto/runtime.proto", fileDescriptor_runtime_aadcca0a1b3e6e7d) }

var fileDescriptor_runtime_aadcca0a1b3e6e7d = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0xf6, 0x46, 0xb1, 0x55, 0x4f, 0xac, 0x20, 0x36, 0x6d, 0x11, 0x21, 0x04, 0x23, 0x5a, 0x6a,
	0x72, 0xb0, 0x8b, 0xdb, 0x53, 0xa0, 0xa7, 0xf4, 0xe7, 0x54, 0x28, 0x1b, 0x72, 0x36, 0xaa, 0x76,
	0x0c, 0xa2, 0x96, 0x56, 0xdd, 0x5d, 0x05, 0x04, 0x7d, 0x90, 0xbe, 0x41, 0x1f, 0xa3, 0xaf, 0x56,
	0x76, 0x57, 0x52, 0xd6, 0xe9, 0xa1, 0x3e, 0xf4, 0x36, 0xdf, 0x7c, 0xf3, 0xf3, 0x7d, 0x23, 0x2d,
	0x9c, 0xd5, 0x52, 0x68, 0xb1, 0x92, 0x4d, 0xa5, 0x8b, 0x12, 0x97, 0x16, 0xd1, 0xb0, 0x83, 0xe7,
	0x17, 0xdb, 0x2c, 0x53, 0x2b, 0xa5, 0x85, 0xc4, 0x95, 0x2b, 0xb4, 0xb1, 0x2b, 0x4b, 0x7f, 0x11,
	0x08, 0x19, 0xe6, 0x58, 0xdc, 0x23, 0xbd, 0x84, 0x20, 0x2f, 0x79, 0x42, 0xe6, 0x64, 0x71, 0xba,
	0x9e, 0x2d, 0xfb, 0x79, 0x37, 0x25, 0x67, 0x86, 0xa0, 0xef, 0xe0, 0x54, 0x69, 0x89, 0x59, 0xb9,
	0x91, 0xf8, 0xbd, 0x41, 0xa5, 0x93, 0xa3, 0x39, 0x59, 0x9c, 0xac, 0x9f, 0x0f, 0xa5, 0xb7, 0x96,
	0x66, 0x8e, 0x65, 0x91, 0xf2, 0xa1, 0x6b, 0x17, 0x12, 0x37, 0x12, 0x55, 0x2d, 0x2a, 0x85, 0x49,
	0xf0, 0x57, 0xbb, 0x90, 0xc8, 0x3a, 0xd6, 0xb4, 0x7b, 0x30, 0xfd, 0x49, 0xe0, 0xf8, 0x16, 0x2b,
	0xfe, 0x4f, 0x99, 0x6f, 0xe1, 0xa4, 0x93, 0xa9, 0xb0, 0xe2, 0x9d, 0xc6, 0xb3, 0x47, 0x1a, 0xcd,
	0x24, 0x06, 0x6a, 0x88, 0xe9, 0x35, 0x44, 0xbd, 0x3a, 0xe7, 0xcd, 0x89, 0x7b, 0xf6, 0x58, 0x9c,
	0xb3, 0x36, 0x53, 0x1e, 0x4a, 0xef, 0x20, 0xda, 0x73, 0x4e, 0x9f, 0xc2, 0x58, 0x8b, 0xba, 0xc8,
	0xad, 0xc8, 0x29, 0x73, 0xc0, 0x64, 0xf1, 0x1e, 0x2b, 0x77, 0xb6, 0x29, 0x73, 0x80, 0x26, 0x10,
	0xd6, 0x59, 0xbb, 0x13, 0x19, 0xb7, 0x2b, 0x67, 0xac, 0x87, 0x29, 0x03, 0x78, 0x10, 0xfb, 0x9f,
	0x66, 0xfe, 0x26, 0x30, 0xf3, 0x9d, 0xd0, 0x17, 0xfe, 0x35, 0xe9, 0xbe, 0xdb, 0xcd, 0x70, 0xd3,
	0x18, 0x82, 0x6f, 0xd8, 0x76, 0x4b, 0x4c, 0x48, 0x5f, 0xc2, 0x04, 0x2b, 0x5d, 0xe8, 0xb6, 0x3b,
	0x54, 0xb4, 0x74, 0xbf, 0xd5, 0x07, 0x9b, 0x64, 0x1d, 0x49, 0x5f, 0x41, 0xb8, 0x2d, 0x76, 0x1a,
	0xa5, 0x4a, 0x8e, 0xe7, 0x81, 0x57, 0xf7, 0xd1, 0x66, 0x59, 0xcf, 0x9a, 0xc2, 0xa6, 0xe6, 0x99,
	0x46, 0x95, 0x8c, 0xf7, 0x0a, 0xef, 0x6c, 0x96, 0xf5, 0x6c, 0xfa, 0xc3, 0x1c, 0xdb, 0xfb, 0x31,
	0x0e, 0x74, 0x70, 0x09, 0xa0, 0x9a, 0x3c, 0x47, 0xa5, 0xb6, 0xcd, 0xce, 0x1a, 0x79, 0xc2, 0xbc,
	0xcc, 0x81, 0x7e, 0xae, 0x2e, 0x20, 0xb8, 0x29, 0x39, 0x05, 0x98, 0xb8, 0x4f, 0x13, 0x8f, 0xe8,
	0x14, 0xc6, 0x76, 0x57, 0x4c, 0xae, 0xd6, 0x30, 0x1d, 0xd6, 0xd2, 0x10, 0x82, 0x4f, 0xa8, 0xe3,
	0x91, 0x09, 0xbe, 0x34, 0x3a, 0x26, 0xa6, 0xeb, 0xb3, 0xe0, 0xc5, 0xb6, 0x8d, 0x8f, 0x4c, 0xf2,
	0x3d, 0xee, 0xe2, 0x60, 0x7d, 0x0d, 0x21, 0x73, 0x92, 0xe9, 0xaa, 0x9f, 0x4a, 0xe3, 0xc1, 0x46,
	0xf7, 0x38, 0xcf, 0xa3, 0x07, 0x63, 0x58, 0xf1, 0x74, 0xb4, 0x20, 0xaf, 0xc9, 0xd7, 0x89, 0x7d,
	0xc4, 0x6f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x88, 0xf0, 0x75, 0x02, 0x04, 0x00, 0x00,
}