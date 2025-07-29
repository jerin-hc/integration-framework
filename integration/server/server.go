package server

import (
	"os"

	"github.com/hashicorp/go-plugin"
	"github.com/jerin-hc/integration-framework/integration/jsoncodec"
	"github.com/jerin-hc/integration-framework/integration/schema"
	"google.golang.org/grpc"
)

const (
	grpcMaxMessageSize        = 256 << 20
	protocolVersionMajor uint = 5
)

type Server struct {
	Ctx    schema.Ctx
	hander schema.IntegrationServer
}

func Init() *Server {
	jsoncodec.Init()
	os.Setenv("TF_RUNTASK_MAGIC_COOKIE", "5c3e2dc2f6b7701f988703046fdbc24eb2e4689f3a81c6af1037d41b8eb063c8")
	return &Server{}
}

func (s *Server) HandleFunc(hander schema.IntegrationServer) {
	s.hander = hander
}

func (s *Server) Run(integrationServer IntegrationServer) {
	serveConfig := &plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  protocolVersionMajor,
			MagicCookieKey:   "TF_RUNTASK_MAGIC_COOKIE",
			MagicCookieValue: "5c3e2dc2f6b7701f988703046fdbc24eb2e4689f3a81c6af1037d41b8eb063c8",
		},
		Plugins: plugin.PluginSet{
			"integration": &HandlerPlugin{
				IntegrationServer: integrationServer,
			},
		},
		GRPCServer: func(opts []grpc.ServerOption) *grpc.Server {
			opts = append(opts, grpc.MaxRecvMsgSize(grpcMaxMessageSize))
			opts = append(opts, grpc.MaxSendMsgSize(grpcMaxMessageSize))

			return grpc.NewServer(opts...)
		},
	}

	plugin.Serve(serveConfig)
}
