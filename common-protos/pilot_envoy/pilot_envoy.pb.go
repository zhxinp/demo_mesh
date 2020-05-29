// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common-protos/pilot_envoy/pilot_envoy.proto

package pilot_envoy

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

type DiscoveryRequest struct {
	Discoveryrequest     string   `protobuf:"bytes,1,opt,name=discoveryrequest,proto3" json:"discoveryrequest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscoveryRequest) Reset()         { *m = DiscoveryRequest{} }
func (m *DiscoveryRequest) String() string { return proto.CompactTextString(m) }
func (*DiscoveryRequest) ProtoMessage()    {}
func (*DiscoveryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_419c5d37c17750e3, []int{0}
}

func (m *DiscoveryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryRequest.Unmarshal(m, b)
}
func (m *DiscoveryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryRequest.Marshal(b, m, deterministic)
}
func (m *DiscoveryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryRequest.Merge(m, src)
}
func (m *DiscoveryRequest) XXX_Size() int {
	return xxx_messageInfo_DiscoveryRequest.Size(m)
}
func (m *DiscoveryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryRequest proto.InternalMessageInfo

func (m *DiscoveryRequest) GetDiscoveryrequest() string {
	if m != nil {
		return m.Discoveryrequest
	}
	return ""
}

type DiscoveryResponse struct {
	Discoveryresponse    string   `protobuf:"bytes,1,opt,name=discoveryresponse,proto3" json:"discoveryresponse,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscoveryResponse) Reset()         { *m = DiscoveryResponse{} }
func (m *DiscoveryResponse) String() string { return proto.CompactTextString(m) }
func (*DiscoveryResponse) ProtoMessage()    {}
func (*DiscoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_419c5d37c17750e3, []int{1}
}

func (m *DiscoveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryResponse.Unmarshal(m, b)
}
func (m *DiscoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryResponse.Marshal(b, m, deterministic)
}
func (m *DiscoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryResponse.Merge(m, src)
}
func (m *DiscoveryResponse) XXX_Size() int {
	return xxx_messageInfo_DiscoveryResponse.Size(m)
}
func (m *DiscoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryResponse proto.InternalMessageInfo

func (m *DiscoveryResponse) GetDiscoveryresponse() string {
	if m != nil {
		return m.Discoveryresponse
	}
	return ""
}

func init() {
	proto.RegisterType((*DiscoveryRequest)(nil), "DiscoveryRequest")
	proto.RegisterType((*DiscoveryResponse)(nil), "DiscoveryResponse")
}

func init() {
	proto.RegisterFile("common-protos/pilot_envoy/pilot_envoy.proto", fileDescriptor_419c5d37c17750e3)
}

var fileDescriptor_419c5d37c17750e3 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xcd, 0x0a, 0x82, 0x50,
	0x10, 0x85, 0xb9, 0x9b, 0xa0, 0x59, 0xe9, 0x5d, 0x95, 0xab, 0x70, 0x25, 0xfd, 0x58, 0xd4, 0x3e,
	0x10, 0xa2, 0x07, 0xd0, 0x07, 0x08, 0xbb, 0x0e, 0x22, 0xa4, 0x63, 0x33, 0x57, 0xc1, 0xb7, 0x0f,
	0xb4, 0x1f, 0xc9, 0xdd, 0xf0, 0x9d, 0x33, 0x87, 0x73, 0x60, 0x63, 0xa8, 0x2c, 0xa9, 0xda, 0xd5,
	0x4c, 0x96, 0x64, 0x5f, 0x17, 0x0f, 0xb2, 0x37, 0xac, 0x5a, 0xea, 0xc6, 0x77, 0xd8, 0xcb, 0xfe,
	0x19, 0x9c, 0x4b, 0x21, 0x86, 0x5a, 0xe4, 0x2e, 0xc6, 0x67, 0x83, 0x62, 0xf5, 0x1a, 0x9c, 0xec,
	0xc3, 0x78, 0x60, 0x0b, 0xb5, 0x52, 0xc1, 0x3c, 0x9e, 0x70, 0x3f, 0x02, 0x77, 0xf4, 0x2f, 0x35,
	0x55, 0x82, 0x7a, 0x0b, 0xee, 0xc8, 0x38, 0xc0, 0x77, 0xc2, 0x54, 0x38, 0x66, 0xe0, 0x45, 0x79,
	0xce, 0x98, 0xa7, 0x16, 0xb3, 0x6f, 0x58, 0x82, 0xdc, 0x16, 0x06, 0xf5, 0x15, 0x96, 0x89, 0x65,
	0x4c, 0xcb, 0x9f, 0x27, 0x46, 0xa1, 0x86, 0x0d, 0x8a, 0x76, 0xc3, 0xff, 0xf2, 0x9e, 0x0e, 0x27,
	0x7d, 0x02, 0x75, 0x50, 0xf7, 0x59, 0xbf, 0xf7, 0xf4, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x56, 0xa2,
	0xac, 0xc9, 0x1e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AggregatedDiscoveryServiceClient is the client API for AggregatedDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AggregatedDiscoveryServiceClient interface {
	StreamAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_StreamAggregatedResourcesClient, error)
}

type aggregatedDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAggregatedDiscoveryServiceClient(cc grpc.ClientConnInterface) AggregatedDiscoveryServiceClient {
	return &aggregatedDiscoveryServiceClient{cc}
}

func (c *aggregatedDiscoveryServiceClient) StreamAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_StreamAggregatedResourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AggregatedDiscoveryService_serviceDesc.Streams[0], "/AggregatedDiscoveryService/StreamAggregatedResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &aggregatedDiscoveryServiceStreamAggregatedResourcesClient{stream}
	return x, nil
}

type AggregatedDiscoveryService_StreamAggregatedResourcesClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type aggregatedDiscoveryServiceStreamAggregatedResourcesClient struct {
	grpc.ClientStream
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AggregatedDiscoveryServiceServer is the server API for AggregatedDiscoveryService service.
type AggregatedDiscoveryServiceServer interface {
	StreamAggregatedResources(AggregatedDiscoveryService_StreamAggregatedResourcesServer) error
}

// UnimplementedAggregatedDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAggregatedDiscoveryServiceServer struct {
}

func (*UnimplementedAggregatedDiscoveryServiceServer) StreamAggregatedResources(srv AggregatedDiscoveryService_StreamAggregatedResourcesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamAggregatedResources not implemented")
}

func RegisterAggregatedDiscoveryServiceServer(s *grpc.Server, srv AggregatedDiscoveryServiceServer) {
	s.RegisterService(&_AggregatedDiscoveryService_serviceDesc, srv)
}

func _AggregatedDiscoveryService_StreamAggregatedResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AggregatedDiscoveryServiceServer).StreamAggregatedResources(&aggregatedDiscoveryServiceStreamAggregatedResourcesServer{stream})
}

type AggregatedDiscoveryService_StreamAggregatedResourcesServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type aggregatedDiscoveryServiceStreamAggregatedResourcesServer struct {
	grpc.ServerStream
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AggregatedDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AggregatedDiscoveryService",
	HandlerType: (*AggregatedDiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAggregatedResources",
			Handler:       _AggregatedDiscoveryService_StreamAggregatedResources_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "common-protos/pilot_envoy/pilot_envoy.proto",
}