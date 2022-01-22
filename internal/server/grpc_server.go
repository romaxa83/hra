package server

import (
	"github.com/romaxa83/hra/config"
	"github.com/romaxa83/hra/pkg/logger"
	orders "github.com/romaxa83/hra/proto"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"

	"google.golang.org/grpc"
)

const (
	maxConnectionIdle = 5
	gRPCTimeout       = 15
	maxConnectionAge  = 5
	gRPCTime          = 10
)

// GrpcServer реализует сервер gRPC для сервиса заказов
type GrpcServer struct {
	server   *grpc.Server
	errCh    chan error
	listener net.Listener
	logger   logger.Logger
	config   *config.Config
}

// создаем Grps сервер
func NewGrpcServer(
	grpsService orders.OrderServiceServer,
	port string,
	logger logger.Logger,
	cfg config.Config,
) (GrpcServer, error) {
	logger.Infof("Create gRPC-server - [:%s]", port)
	logger.Infof("%+v", cfg.Mongo.URI)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return GrpcServer{}, err
	}

	grpcServer := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: maxConnectionIdle * time.Minute,
		Timeout:           gRPCTimeout * time.Second,
		MaxConnectionAge:  maxConnectionAge * time.Minute,
		Time:              gRPCTime * time.Minute,
	}))

	//orderService := grpc2.NewGrpcOrderService()
	orders.RegisterOrderServiceServer(grpcServer, grpsService)

	return GrpcServer{
		server:   grpcServer,
		listener: lis,
		errCh:    make(chan error),
		logger:   logger,
		config:   &cfg,
	}, nil
}

// Start запускает сервер GRPC в фоновом режиме, отправляя ошибку в канал ошибок
func (g GrpcServer) Start() {
	g.logger.Infof("🚀 Start gRPC-server")
	go func() {
		g.errCh <- g.server.Serve(g.listener)
	}()
}

// Stop останавливает сервер
func (g GrpcServer) Stop() {
	g.server.GracefulStop()
}

// Error возвращает канал ошибок сервера
func (g GrpcServer) Error() chan error {
	return g.errCh
}
