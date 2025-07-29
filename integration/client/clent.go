package client

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/go-plugin"
	"github.com/jerin-hc/integration-framework/integration/jsoncodec"
	"github.com/jerin-hc/integration-framework/integration/schema"
	"github.com/jerin-hc/integration-framework/integration/tfgrpc"
)

type Clent struct {
	Ctx  schema.Ctx
	path string
}

func Init() *Clent {
	jsoncodec.Init()
	os.Setenv("TF_RUNTASK_MAGIC_COOKIE", "5c3e2dc2f6b7701f988703046fdbc24eb2e4689f3a81c6af1037d41b8eb063c8")
	return &Clent{
		path: "/home/tfc-agent/integration/terraform-%s",
	}
}

func (s *Clent) RunTask(ctx context.Context, event schema.Event, integrationPlugin string) (*schema.Response, error) {
	pluginPath := fmt.Sprintf(s.path, integrationPlugin)

	var resp *schema.Response
	var err error

	pluginMap := map[string]plugin.Plugin{
		"integration": &tfgrpc.HandlerPlugin{},
	}

	pluginClient := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  tfgrpc.GetHandShakeConfig(),
		Plugins:          pluginMap,
		Cmd:              exec.Command(pluginPath),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	})
	rpcClient, err := pluginClient.Client()
	if err != nil {
		log.Panic(err)
	}

	raw, err := rpcClient.Dispense("integration")
	if err != nil {
		log.Panic(err)
	}

	integration, ok := raw.(tfgrpc.IntegrationServer)

	if !ok {
		log.Panic("invalid IntegrationClient")
	}

	req := &schema.Request{
		Resources: nil,
	}

	switch event {
	case schema.PrePlan:
		req.Event = schema.PrePlan
		resp, err = integration.HandlePrePlan(ctx, req)
	case schema.PostPlan:
		req.Event = schema.PostPlan
		resp, err = integration.HandlePostPlan(ctx, req)
	case schema.PreApply:
		req.Event = schema.PreApply
		resp, err = integration.HandlePreApply(ctx, req)
	case schema.PostApply:
		req.Event = schema.PostApply
		resp, err = integration.HandlePostApply(ctx, req)
	case schema.Test:
		req.Event = schema.Test
		resp, err = integration.HandleTest(ctx, req)
	default:
		resp = nil
		err = nil
	}

	return resp, err
}
