package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/romaxa83/hra/internal/order/commands"
	"github.com/romaxa83/hra/internal/order/repository"
	"github.com/romaxa83/hra/internal/order/service"
	"github.com/romaxa83/hra/pkg/logger"
	orders "github.com/romaxa83/hra/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type grpcOrderService struct {
	Logger       logger.Logger
	OrderService *service.OrderService
	Repo         repository.Repository
}

func NewGrpcOrderService(
	logger logger.Logger,
	orderService *service.OrderService,
	repo repository.Repository,
) *grpcOrderService {
	return &grpcOrderService{
		Logger:       logger,
		OrderService: orderService,
		Repo:         repo,
	}
}

func (s *grpcOrderService) CreateTest(ctx context.Context, req *orders.CreateTestRequest) (*orders.TestResponse, error) {
	
	// создаем команду (аналог dto)
	id, _ := uuid.NewRandom()
	command := commands.NewCreateOrderCmd(id.String(), req.Name, time.Now())
	// запускаем команду, создания заявки
	if err := s.OrderService.Commands.CreateOrder.Handle(ctx, *command); err != nil {
		s.Logger.Error("CreateOrder.Handle ", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	res := orders.TestResponse{Name: id.String()}

	return &res, nil
}

func (s *grpcOrderService) errResponse(c codes.Code, err error) error {
	return status.Error(c, err.Error())
}

// Создает новый заказ

func (s *grpcOrderService) Create(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	panic("not implement")
}

// Извлекает существующий заказ

func (s *grpcOrderService) Retrieve(ctx context.Context, req *orders.RetrieveOrderRequest) (*orders.RetrieveOrderResponse, error) {
	panic("not implement")
}

// Изменяет существующий заказ

func (s *grpcOrderService) Update(ctx context.Context, req *orders.UpdateOrderRequest) (*orders.UpdateOrderResponse, error) {
	panic("not implement")
}

// Отменяет существующий заказ

func (s *grpcOrderService) Delete(ctx context.Context, req *orders.DeleteOrderRequest) (*orders.DeleteOrderResponse, error) {
	panic("not implement")
}

// Выдает список текущих заказов

func (s *grpcOrderService) List(ctx context.Context, req *orders.ListOrderRequest) (*orders.ListOrderResponse, error) {
	panic("not implement")
}
