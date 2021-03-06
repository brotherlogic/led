// Code generated by protoc-gen-go. DO NOT EDIT.
// source: led.proto

/*
Package led is a generated protocol buffer package.

It is generated from these files:
	led.proto

It has these top-level messages:
	DisplayRequest
	DisplayResponse
*/
package led

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

type DisplayRequest struct {
	TopLine    string `protobuf:"bytes,1,opt,name=top_line,json=topLine" json:"top_line,omitempty"`
	BottomLine string `protobuf:"bytes,2,opt,name=bottom_line,json=bottomLine" json:"bottom_line,omitempty"`
}

func (m *DisplayRequest) Reset()                    { *m = DisplayRequest{} }
func (m *DisplayRequest) String() string            { return proto.CompactTextString(m) }
func (*DisplayRequest) ProtoMessage()               {}
func (*DisplayRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DisplayRequest) GetTopLine() string {
	if m != nil {
		return m.TopLine
	}
	return ""
}

func (m *DisplayRequest) GetBottomLine() string {
	if m != nil {
		return m.BottomLine
	}
	return ""
}

type DisplayResponse struct {
}

func (m *DisplayResponse) Reset()                    { *m = DisplayResponse{} }
func (m *DisplayResponse) String() string            { return proto.CompactTextString(m) }
func (*DisplayResponse) ProtoMessage()               {}
func (*DisplayResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*DisplayRequest)(nil), "led.DisplayRequest")
	proto.RegisterType((*DisplayResponse)(nil), "led.DisplayResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LedService service

type LedServiceClient interface {
	Display(ctx context.Context, in *DisplayRequest, opts ...grpc.CallOption) (*DisplayResponse, error)
}

type ledServiceClient struct {
	cc *grpc.ClientConn
}

func NewLedServiceClient(cc *grpc.ClientConn) LedServiceClient {
	return &ledServiceClient{cc}
}

func (c *ledServiceClient) Display(ctx context.Context, in *DisplayRequest, opts ...grpc.CallOption) (*DisplayResponse, error) {
	out := new(DisplayResponse)
	err := grpc.Invoke(ctx, "/led.LedService/Display", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LedService service

type LedServiceServer interface {
	Display(context.Context, *DisplayRequest) (*DisplayResponse, error)
}

func RegisterLedServiceServer(s *grpc.Server, srv LedServiceServer) {
	s.RegisterService(&_LedService_serviceDesc, srv)
}

func _LedService_Display_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisplayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedServiceServer).Display(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/led.LedService/Display",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedServiceServer).Display(ctx, req.(*DisplayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "led.LedService",
	HandlerType: (*LedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Display",
			Handler:    _LedService_Display_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "led.proto",
}

func init() { proto.RegisterFile("led.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x49, 0x4d, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x49, 0x4d, 0x51, 0xf2, 0xe1, 0xe2, 0x73, 0xc9,
	0x2c, 0x2e, 0xc8, 0x49, 0xac, 0x0c, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe4, 0xe2,
	0x28, 0xc9, 0x2f, 0x88, 0xcf, 0xc9, 0xcc, 0x4b, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62,
	0x2f, 0xc9, 0x2f, 0xf0, 0xc9, 0xcc, 0x4b, 0x15, 0x92, 0xe7, 0xe2, 0x4e, 0xca, 0x2f, 0x29, 0xc9,
	0xcf, 0x85, 0xc8, 0x32, 0x81, 0x65, 0xb9, 0x20, 0x42, 0x20, 0x05, 0x4a, 0x82, 0x5c, 0xfc, 0x70,
	0xd3, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x8d, 0x5c, 0xb8, 0xb8, 0x7c, 0x52, 0x53, 0x82, 0x53,
	0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0xcc, 0xb8, 0xd8, 0xa1, 0x0a, 0x84, 0x84, 0xf5, 0x40, 0x4e,
	0x41, 0xb5, 0x5c, 0x4a, 0x04, 0x55, 0x10, 0x62, 0x86, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0xc9, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xfb, 0xf4, 0x9a, 0xbf, 0x00, 0x00, 0x00,
}
