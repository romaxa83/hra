package grpc

import (
	"context"
	orders "github.com/romaxa83/hra/proto"
)

type grpcOrderService struct {
}

func NewGrpcOrderService() *grpcOrderService {
	return &grpcOrderService{}
}

func (s *grpcOrderService) CreateTest(ctx context.Context, req *orders.CreateTestRequest) (*orders.TestResponse, error) {
	res := orders.TestResponse{Name: "WOP OOOOO"}
	//logger.Warnf("GRPC_CTX %+v", ctx)
	//logger.Warnf("GRPC_REQ %+v", req)

	return &res, nil
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
