// Code generated by protoc-gen-go. DO NOT EDIT.
// source: svrbootstrap.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PreJoinRequest struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Host   string `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	Client string `protobuf:"bytes,3,opt,name=client" json:"client,omitempty"`
}

func (m *PreJoinRequest) Reset()                    { *m = PreJoinRequest{} }
func (m *PreJoinRequest) String() string            { return proto1.CompactTextString(m) }
func (*PreJoinRequest) ProtoMessage()               {}
func (*PreJoinRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *PreJoinRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PreJoinRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *PreJoinRequest) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

type PreJoinResponse struct {
	Result     int64             `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	Bootexpect int64             `protobuf:"varint,2,opt,name=bootexpect" json:"bootexpect,omitempty"`
	Message    string            `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Jointarget string            `protobuf:"bytes,4,opt,name=jointarget" json:"jointarget,omitempty"`
	Curnodes   map[string]string `protobuf:"bytes,5,rep,name=curnodes" json:"curnodes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *PreJoinResponse) Reset()                    { *m = PreJoinResponse{} }
func (m *PreJoinResponse) String() string            { return proto1.CompactTextString(m) }
func (*PreJoinResponse) ProtoMessage()               {}
func (*PreJoinResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *PreJoinResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *PreJoinResponse) GetBootexpect() int64 {
	if m != nil {
		return m.Bootexpect
	}
	return 0
}

func (m *PreJoinResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PreJoinResponse) GetJointarget() string {
	if m != nil {
		return m.Jointarget
	}
	return ""
}

func (m *PreJoinResponse) GetCurnodes() map[string]string {
	if m != nil {
		return m.Curnodes
	}
	return nil
}

func init() {
	proto1.RegisterType((*PreJoinRequest)(nil), "proto.PreJoinRequest")
	proto1.RegisterType((*PreJoinResponse)(nil), "proto.PreJoinResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DiscoveryAsBoot service

type DiscoveryAsBootClient interface {
	PreJoin(ctx context.Context, in *PreJoinRequest, opts ...grpc.CallOption) (*PreJoinResponse, error)
}

type discoveryAsBootClient struct {
	cc *grpc.ClientConn
}

func NewDiscoveryAsBootClient(cc *grpc.ClientConn) DiscoveryAsBootClient {
	return &discoveryAsBootClient{cc}
}

func (c *discoveryAsBootClient) PreJoin(ctx context.Context, in *PreJoinRequest, opts ...grpc.CallOption) (*PreJoinResponse, error) {
	out := new(PreJoinResponse)
	err := grpc.Invoke(ctx, "/proto.DiscoveryAsBoot/PreJoin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DiscoveryAsBoot service

type DiscoveryAsBootServer interface {
	PreJoin(context.Context, *PreJoinRequest) (*PreJoinResponse, error)
}

func RegisterDiscoveryAsBootServer(s *grpc.Server, srv DiscoveryAsBootServer) {
	s.RegisterService(&_DiscoveryAsBoot_serviceDesc, srv)
}

func _DiscoveryAsBoot_PreJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreJoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryAsBootServer).PreJoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DiscoveryAsBoot/PreJoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryAsBootServer).PreJoin(ctx, req.(*PreJoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DiscoveryAsBoot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DiscoveryAsBoot",
	HandlerType: (*DiscoveryAsBootServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PreJoin",
			Handler:    _DiscoveryAsBoot_PreJoin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svrbootstrap.proto",
}

func init() { proto1.RegisterFile("svrbootstrap.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xcd, 0x4e, 0xf3, 0x40,
	0x0c, 0xfc, 0xda, 0xf4, 0xe7, 0xc3, 0x08, 0x8a, 0x2c, 0xa8, 0x56, 0x3d, 0xa0, 0xaa, 0xe2, 0xd0,
	0x53, 0x0f, 0xe5, 0x82, 0xca, 0x85, 0xdf, 0x0b, 0x12, 0x52, 0x95, 0x37, 0x48, 0x83, 0x55, 0x02,
	0xe9, 0x3a, 0xac, 0x37, 0x11, 0xb9, 0xf2, 0xe4, 0x68, 0x37, 0x4b, 0x55, 0x2a, 0x4e, 0xeb, 0x19,
	0xdb, 0xe3, 0xd9, 0x01, 0x94, 0xca, 0xac, 0x98, 0xad, 0x58, 0x93, 0x14, 0xb3, 0xc2, 0xb0, 0x65,
	0xec, 0xfa, 0x67, 0xb2, 0x84, 0xe3, 0xa5, 0xa1, 0x27, 0xce, 0x74, 0x4c, 0x1f, 0x25, 0x89, 0x45,
	0x84, 0x8e, 0x4e, 0x36, 0xa4, 0x5a, 0xe3, 0xd6, 0xf4, 0x20, 0xf6, 0xb5, 0xe3, 0x5e, 0x59, 0xac,
	0x6a, 0x37, 0x9c, 0xab, 0x71, 0x08, 0xbd, 0x34, 0xcf, 0x48, 0x5b, 0x15, 0x79, 0x36, 0xa0, 0xc9,
	0x57, 0x1b, 0x06, 0x5b, 0x49, 0x29, 0x58, 0x0b, 0xb9, 0x59, 0x43, 0x52, 0xe6, 0xd6, 0xab, 0x46,
	0x71, 0x40, 0x78, 0x0e, 0xe0, 0x7c, 0xd1, 0x67, 0x41, 0x69, 0xa3, 0x1e, 0xc5, 0x3b, 0x0c, 0x2a,
	0xe8, 0x6f, 0x48, 0x24, 0x59, 0x53, 0x38, 0xf2, 0x03, 0xdd, 0xe6, 0x1b, 0x67, 0xda, 0x26, 0x66,
	0x4d, 0x56, 0x75, 0x7c, 0x73, 0x87, 0xc1, 0x1b, 0xf8, 0x9f, 0x96, 0x46, 0xf3, 0x0b, 0x89, 0xea,
	0x8e, 0xa3, 0xe9, 0xe1, 0xfc, 0xa2, 0xf9, 0xf8, 0x6c, 0xcf, 0xdb, 0xec, 0x3e, 0x8c, 0x3d, 0x6a,
	0x6b, 0xea, 0x78, 0xbb, 0x35, 0xba, 0x86, 0xa3, 0x5f, 0x2d, 0x3c, 0x81, 0xe8, 0x9d, 0xea, 0x90,
	0x8b, 0x2b, 0xf1, 0x14, 0xba, 0x55, 0x92, 0x97, 0x14, 0x72, 0x69, 0xc0, 0xa2, 0x7d, 0xd5, 0x9a,
	0x3f, 0xc3, 0xe0, 0x21, 0x93, 0x94, 0x2b, 0x32, 0xf5, 0xad, 0xdc, 0x31, 0x5b, 0x5c, 0x40, 0x3f,
	0x9c, 0xc6, 0xb3, 0x7d, 0x2b, 0x3e, 0xf9, 0xd1, 0xf0, 0x6f, 0x87, 0x93, 0x7f, 0xab, 0x9e, 0x6f,
	0x5c, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x99, 0x60, 0x20, 0x09, 0xc9, 0x01, 0x00, 0x00,
}