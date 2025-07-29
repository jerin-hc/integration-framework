package tfgrpc

import (
	"context"
	"errors"
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"github.com/jerin-hc/integration-framework/integration/schema"
	"google.golang.org/grpc"
)

func GetHandShakeConfig() plugin.HandshakeConfig {
	return plugin.HandshakeConfig{
		ProtocolVersion:  5,
		MagicCookieKey:   "TF_RUNTASK_MAGIC_COOKIE",
		MagicCookieValue: "5c3e2dc2f6b7701f988703046fdbc24eb2e4689f3a81c6af1037d41b8eb063c8",
	}
}

type HandlerPlugin struct {
	IntegrationServer IntegrationServer
}

type integrationClient struct {
	cc *grpc.ClientConn
}

func (c *integrationClient) HandlePrePlan(ctx context.Context, req *schema.Request) (*schema.Response, error) {
	resp := new(schema.Response)
	err := c.cc.Invoke(ctx, "/IntegrationService/HandlePrePlan", req, resp, grpc.CallContentSubtype("json"))
	return resp, err
}

func (c *integrationClient) HandlePostPlan(ctx context.Context, req *schema.Request) (*schema.Response, error) {
	resp := new(schema.Response)
	err := c.cc.Invoke(ctx, "/IntegrationService/HandlePostPlan", req, resp, grpc.CallContentSubtype("json"))
	return resp, err
}

func (c *integrationClient) HandlePreApply(ctx context.Context, req *schema.Request) (*schema.Response, error) {
	resp := new(schema.Response)
	err := c.cc.Invoke(ctx, "/IntegrationService/HandlePreApply", req, resp, grpc.CallContentSubtype("json"))
	return resp, err
}

func (c *integrationClient) HandlePostApply(ctx context.Context, req *schema.Request) (*schema.Response, error) {
	resp := new(schema.Response)
	err := c.cc.Invoke(ctx, "/IntegrationService/HandlePostApply", req, resp, grpc.CallContentSubtype("json"))
	return resp, err
}

func (c *integrationClient) HandleTest(ctx context.Context, req *schema.Request) (*schema.Response, error) {
	resp := new(schema.Response)
	err := c.cc.Invoke(ctx, "/IntegrationService/HandleTest", req, resp, grpc.CallContentSubtype("json"))
	return resp, err
}

func (c *integrationClient) Trigger(ctx context.Context, req *schema.Request) (*schema.Response, error) {
	resp := new(schema.Response)
	err := c.cc.Invoke(ctx, "/IntegrationService/Trigger", req, resp, grpc.CallContentSubtype("json"))
	return resp, err
}

func (p *HandlerPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	s.RegisterService(&Integration_ServiceDesc, p.IntegrationServer)
	return nil
}

func (p *HandlerPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, cc *grpc.ClientConn) (any, error) {
	return &integrationClient{cc: cc}, nil
}

func (p *HandlerPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return nil, errors.New("terraform-plugin-go only implements gRPC servers")
}

func (p *HandlerPlugin) Client(*plugin.MuxBroker, *rpc.Client) (interface{}, error) {
	return nil, errors.New("terraform-plugin-go only implements gRPC servers")
}

type IntegrationServer interface {
	HandlePrePlan(context.Context, *schema.Request) (*schema.Response, error)
	HandlePostPlan(context.Context, *schema.Request) (*schema.Response, error)
	HandlePreApply(context.Context, *schema.Request) (*schema.Response, error)
	HandlePostApply(context.Context, *schema.Request) (*schema.Response, error)
	HandleTest(context.Context, *schema.Request) (*schema.Response, error)
	Trigger(context.Context, *schema.Request) (*schema.Response, error)
}

var Integration_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "IntegrationService",
	HandlerType: (*IntegrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandlePrePlan",
			Handler: func(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
				var hr schema.Request

				if err := dec(&hr); err != nil {
					return nil, err
				}

				if interceptor == nil {
					return srv.(IntegrationServer).HandlePrePlan(ctx, &hr)
				}

				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/IntegrationService/HandlePrePlan",
				}

				handler := func(ctx context.Context, req any) (any, error) {
					return srv.(IntegrationServer).HandlePrePlan(ctx, req.(*schema.Request))
				}

				return interceptor(ctx, &hr, info, handler)
			},
		},
		{
			MethodName: "HandlePostPlan",
			Handler: func(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
				var hr schema.Request

				if err := dec(&hr); err != nil {
					return nil, err
				}

				if interceptor == nil {
					return srv.(IntegrationServer).HandlePostPlan(ctx, &hr)
				}

				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/IntegrationService/HandlePostPlan",
				}

				handler := func(ctx context.Context, req any) (any, error) {
					return srv.(IntegrationServer).HandlePostPlan(ctx, req.(*schema.Request))
				}

				return interceptor(ctx, &hr, info, handler)
			},
		},
		{
			MethodName: "HandlePreApply",
			Handler: func(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
				var hr schema.Request

				if err := dec(&hr); err != nil {
					return nil, err
				}

				if interceptor == nil {
					return srv.(IntegrationServer).HandlePreApply(ctx, &hr)
				}

				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/IntegrationService/HandlePreApply",
				}

				handler := func(ctx context.Context, req any) (any, error) {
					return srv.(IntegrationServer).HandlePreApply(ctx, req.(*schema.Request))
				}

				return interceptor(ctx, &hr, info, handler)
			},
		},
		{
			MethodName: "HandlePostApply",
			Handler: func(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
				var hr schema.Request

				if err := dec(&hr); err != nil {
					return nil, err
				}

				if interceptor == nil {
					return srv.(IntegrationServer).HandlePostApply(ctx, &hr)
				}

				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/IntegrationService/HandlePostApply",
				}

				handler := func(ctx context.Context, req any) (any, error) {
					return srv.(IntegrationServer).HandlePostApply(ctx, req.(*schema.Request))
				}

				return interceptor(ctx, &hr, info, handler)
			},
		},
		{
			MethodName: "HandleTest",
			Handler: func(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
				var hr schema.Request

				if err := dec(&hr); err != nil {
					return nil, err
				}

				if interceptor == nil {
					return srv.(IntegrationServer).HandleTest(ctx, &hr)
				}

				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/IntegrationService/HandleTest",
				}

				handler := func(ctx context.Context, req any) (any, error) {
					return srv.(IntegrationServer).HandleTest(ctx, req.(*schema.Request))
				}

				return interceptor(ctx, &hr, info, handler)
			},
		},
		{
			MethodName: "Trigger",
			Handler: func(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
				var hr schema.Request

				if err := dec(&hr); err != nil {
					return nil, err
				}

				if interceptor == nil {
					return srv.(IntegrationServer).Trigger(ctx, &hr)
				}

				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/IntegrationService/Trigger",
				}

				handler := func(ctx context.Context, req any) (any, error) {
					return srv.(IntegrationServer).Trigger(ctx, req.(*schema.Request))
				}

				return interceptor(ctx, &hr, info, handler)
			},
		},
	},
}
