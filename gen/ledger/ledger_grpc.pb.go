// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ledger

import (
	context "context"

	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// LedgerServiceClient is the client API for LedgerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LedgerServiceClient interface {
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetAccountBalance(ctx context.Context, in *GetAccountBalanceRequest, opts ...grpc.CallOption) (*GetAccountBalanceResponse, error)
	GetAnalyticalData(ctx context.Context, in *GetAnalyticalDataRequest, opts ...grpc.CallOption) (LedgerService_GetAnalyticalDataClient, error)
}

type ledgerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLedgerServiceClient(cc grpc.ClientConnInterface) LedgerServiceClient {
	return &ledgerServiceClient{cc}
}

func (c *ledgerServiceClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ledger.LedgerService/CreateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) GetAccountBalance(ctx context.Context, in *GetAccountBalanceRequest, opts ...grpc.CallOption) (*GetAccountBalanceResponse, error) {
	out := new(GetAccountBalanceResponse)
	err := c.cc.Invoke(ctx, "/ledger.LedgerService/GetAccountBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) GetAnalyticalData(ctx context.Context, in *GetAnalyticalDataRequest, opts ...grpc.CallOption) (LedgerService_GetAnalyticalDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LedgerService_serviceDesc.Streams[0], "/ledger.LedgerService/GetAnalyticalData", opts...)
	if err != nil {
		return nil, err
	}
	x := &ledgerServiceGetAnalyticalDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LedgerService_GetAnalyticalDataClient interface {
	Recv() (*GetAnalyticalDataResponse, error)
	grpc.ClientStream
}

type ledgerServiceGetAnalyticalDataClient struct {
	grpc.ClientStream
}

func (x *ledgerServiceGetAnalyticalDataClient) Recv() (*GetAnalyticalDataResponse, error) {
	m := new(GetAnalyticalDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LedgerServiceServer is the server API for LedgerService service.
// All implementations should embed UnimplementedLedgerServiceServer
// for forward compatibility
type LedgerServiceServer interface {
	CreateTransaction(context.Context, *CreateTransactionRequest) (*empty.Empty, error)
	GetAccountBalance(context.Context, *GetAccountBalanceRequest) (*GetAccountBalanceResponse, error)
	GetAnalyticalData(*GetAnalyticalDataRequest, LedgerService_GetAnalyticalDataServer) error
}

// UnimplementedLedgerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLedgerServiceServer struct {
}

func (UnimplementedLedgerServiceServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedLedgerServiceServer) GetAccountBalance(context.Context, *GetAccountBalanceRequest) (*GetAccountBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountBalance not implemented")
}
func (UnimplementedLedgerServiceServer) GetAnalyticalData(*GetAnalyticalDataRequest, LedgerService_GetAnalyticalDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAnalyticalData not implemented")
}

// UnsafeLedgerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LedgerServiceServer will
// result in compilation errors.
type UnsafeLedgerServiceServer interface {
	mustEmbedUnimplementedLedgerServiceServer()
}

func RegisterLedgerServiceServer(s grpc.ServiceRegistrar, srv LedgerServiceServer) {
	s.RegisterService(&_LedgerService_serviceDesc, srv)
}

func _LedgerService_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.LedgerService/CreateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_GetAccountBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).GetAccountBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.LedgerService/GetAccountBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).GetAccountBalance(ctx, req.(*GetAccountBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_GetAnalyticalData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAnalyticalDataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LedgerServiceServer).GetAnalyticalData(m, &ledgerServiceGetAnalyticalDataServer{stream})
}

type LedgerService_GetAnalyticalDataServer interface {
	Send(*GetAnalyticalDataResponse) error
	grpc.ServerStream
}

type ledgerServiceGetAnalyticalDataServer struct {
	grpc.ServerStream
}

func (x *ledgerServiceGetAnalyticalDataServer) Send(m *GetAnalyticalDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _LedgerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ledger.LedgerService",
	HandlerType: (*LedgerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTransaction",
			Handler:    _LedgerService_CreateTransaction_Handler,
		},
		{
			MethodName: "GetAccountBalance",
			Handler:    _LedgerService_GetAccountBalance_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAnalyticalData",
			Handler:       _LedgerService_GetAnalyticalData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ledger/ledger.proto",
}

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthClient interface {
	// Check - checks the system health.
	Check(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthClient(cc grpc.ClientConnInterface) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/ledger.Health/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServer is the server API for Health service.
// All implementations should embed UnimplementedHealthServer
// for forward compatibility
type HealthServer interface {
	// Check - checks the system health.
	Check(context.Context, *empty.Empty) (*HealthCheckResponse, error)
}

// UnimplementedHealthServer should be embedded to have forward compatible implementations.
type UnimplementedHealthServer struct {
}

func (UnimplementedHealthServer) Check(context.Context, *empty.Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

// UnsafeHealthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthServer will
// result in compilation errors.
type UnsafeHealthServer interface {
	mustEmbedUnimplementedHealthServer()
}

func RegisterHealthServer(s grpc.ServiceRegistrar, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ledger.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ledger/ledger.proto",
}
