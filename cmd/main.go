package main

import (
	"context"
	"flag"
	"github.com/romaxa83/hra/config"
	"github.com/romaxa83/hra/internal/order/delivery/grpc"
	"github.com/romaxa83/hra/internal/order/repository"
	"github.com/romaxa83/hra/internal/order/service"
	"github.com/romaxa83/hra/internal/server"
	"github.com/romaxa83/hra/pkg/logger"
	"github.com/romaxa83/hra/pkg/mongodb"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	grpcPort = "50051"
	restPort = "9991"
)

// Оболочка app отлично подходит для всех элементов, необходимых для запуска
// и завершения работы микросервиса Order
type app struct {
	restServer server.RestServer
	grpcServer *server.GrpcServer
	/* Listens for an application termination signal
	   Ex. (Ctrl X, Docker container shutdown, etc) */
	shutdownCh chan os.Signal
	//config     config.Config
}

// start запускает сервера REST и gRPC в фоновом режиме
func (a app) start() {
	a.restServer.Start() // non blocking now
	a.grpcServer.Start() // also non blocking :-)
}

// stop останавливает сервера
func (a app) shutdown() error {
	a.grpcServer.Stop()
	return a.restServer.Stop()
}

// newApp создает новое приложение с серверами REST и gRPC
// Эта функция выполняет все необходимые для приложения инициализации
func newApp() (app, error) {

	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()
	appLogger.Info("Init logger")

	ctx := context.Background()
	mongoDBConn, err := mongodb.NewMongoDBConn(ctx, cfg.Mongo)
	if err != nil {
		log.Fatal(err)
	}
	mongoClient := mongoDBConn
	//defer mongoDBConn.Disconnect(ctx) // nolint: errcheck
	appLogger.Infof("Mongo connected: %v", mongoDBConn.NumberSessionsInProgress())

	mongoRepo := repository.NewMongoRepository(appLogger, cfg, mongoClient)

	ser := service.NewOrderService(appLogger, mongoRepo)
	orderService := grpc.NewGrpcOrderService(appLogger, ser, mongoRepo)

	gs, err := server.NewGrpcServer(orderService, grpcPort, appLogger, *cfg)
	if err != nil {
		return app{}, err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	return app{
		restServer: server.NewRestServer(orderService, restPort, appLogger, *cfg),
		grpcServer: gs,
		shutdownCh: quit,
	}, nil
}

// run запускает приложение, обрабатывая любые ошибки серверов REST и gRPC
// и сигналы завершения работы
func run() error {
	app, err := newApp()
	if err != nil {
		return err
	}

	app.start()
	defer app.shutdown()

	select {
	case restErr := <-app.restServer.Error():
		return restErr
	case grpcErr := <-app.grpcServer.Error():
		return grpcErr
	case <-app.shutdownCh:
		return nil
	}
}

func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
