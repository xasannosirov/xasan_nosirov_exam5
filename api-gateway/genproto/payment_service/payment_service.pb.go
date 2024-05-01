// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: payment_service.proto

package payment_service

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

func init() { proto.RegisterFile("payment_service.proto", fileDescriptor_e7aa2556c2d08659) }

var fileDescriptor_e7aa2556c2d08659 = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x48, 0xac, 0xcc,
	0x4d, 0xcd, 0x2b, 0x89, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x47, 0x13, 0x96, 0x12, 0x86, 0x09, 0xe4, 0xe6, 0xa7, 0xa4, 0xe6, 0x40, 0x54,
	0x19, 0x79, 0x72, 0xf1, 0x05, 0x40, 0x84, 0x83, 0x21, 0xca, 0x84, 0xcc, 0xb9, 0xd8, 0x1d, 0x53,
	0x52, 0x9c, 0x13, 0x8b, 0x52, 0x84, 0x44, 0xf5, 0xd0, 0x8d, 0x06, 0x09, 0x4b, 0x61, 0x17, 0x76,
	0xd2, 0x3a, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x67, 0x3c,
	0x96, 0x63, 0x88, 0x92, 0x48, 0x4f, 0xcd, 0x03, 0x5b, 0xa3, 0x8f, 0xa6, 0x21, 0x89, 0x0d, 0x2c,
	0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x99, 0x13, 0xc1, 0xbc, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaymentServiceClient interface {
	AddCard(ctx context.Context, in *Card, opts ...grpc.CallOption) (*Card, error)
}

type paymentServiceClient struct {
	cc *grpc.ClientConn
}

func NewPaymentServiceClient(cc *grpc.ClientConn) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) AddCard(ctx context.Context, in *Card, opts ...grpc.CallOption) (*Card, error) {
	out := new(Card)
	err := c.cc.Invoke(ctx, "/payment_service.PaymentService/AddCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
type PaymentServiceServer interface {
	AddCard(context.Context, *Card) (*Card, error)
}

// UnimplementedPaymentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPaymentServiceServer struct {
}

func (*UnimplementedPaymentServiceServer) AddCard(ctx context.Context, req *Card) (*Card, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCard not implemented")
}

func RegisterPaymentServiceServer(s *grpc.Server, srv PaymentServiceServer) {
	s.RegisterService(&_PaymentService_serviceDesc, srv)
}

func _PaymentService_AddCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Card)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).AddCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment_service.PaymentService/AddCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).AddCard(ctx, req.(*Card))
	}
	return interceptor(ctx, in, info, handler)
}

var _PaymentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "payment_service.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCard",
			Handler:    _PaymentService_AddCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment_service.proto",
}