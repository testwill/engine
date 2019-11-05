package grpc

import (
	"context"
	"net"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/mesg-foundation/engine/config"
	protobuf_api "github.com/mesg-foundation/engine/protobuf/api"
	"github.com/mesg-foundation/engine/sdk"
	"github.com/mesg-foundation/engine/server/grpc/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

// Server contains the server config.
type Server struct {
	instance *grpc.Server
	sdk      *sdk.SDK
	cfg      *config.Config
}

// New returns a new gRPC server.
func New(sdk *sdk.SDK, cfg *config.Config) *Server {
	return &Server{sdk: sdk, cfg: cfg}
}

// Serve listens for connections.
func (s *Server) Serve(address string) error {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	// Keep alive prevents Docker network to drop TCP idle connections after 15 minutes.
	// See: https://forum.mesg.com/t/solution-summary-for-docker-dropping-connections-after-15-min/246
	keepaliveOpt := grpc.KeepaliveParams(keepalive.ServerParameters{
		Time: 1 * time.Minute,
	})
	s.instance = grpc.NewServer(
		keepaliveOpt,
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_logrus.StreamServerInterceptor(logrus.StandardLogger().WithField("module", "grpc")),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_logrus.UnaryServerInterceptor(logrus.StandardLogger().WithField("module", "grpc")),
			func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
				// inject credentials from config
				ctx = metadata.NewIncomingContext(ctx, map[string][]string{
					"credential_username":   {s.cfg.Account.Name},
					"credential_passphrase": {s.cfg.Account.Password},
				})
				return handler(ctx, req)
			},
		)),
	)
	s.register()
	logrus.WithField("module", "grpc").Info("server listens on ", ln.Addr())
	return s.instance.Serve(ln)
}

// Close gracefully closes the server.
func (s *Server) Close() {
	s.instance.GracefulStop()
}

// register all server
func (s *Server) register() {
	protobuf_api.RegisterEventServer(s.instance, api.NewEventServer(s.sdk))
	protobuf_api.RegisterExecutionServer(s.instance, api.NewExecutionServer(s.sdk))
	protobuf_api.RegisterInstanceServer(s.instance, api.NewInstanceServer(s.sdk))
	protobuf_api.RegisterServiceServer(s.instance, api.NewServiceServer(s.sdk))
	protobuf_api.RegisterProcessServer(s.instance, api.NewProcessServer(s.sdk))
	protobuf_api.RegisterAccountServer(s.instance, api.NewAccountServer(s.sdk))
	protobuf_api.RegisterOwnershipServer(s.instance, api.NewOwnershipServer(s.sdk))
	protobuf_api.RegisterRunnerServer(s.instance, api.NewRunnerServer(s.sdk, s.cfg))

	reflection.Register(s.instance)
}
