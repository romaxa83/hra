package service

import (
	"github.com/romaxa83/hra/internal/order/commands"
	"github.com/romaxa83/hra/internal/order/repository"
	"github.com/romaxa83/hra/pkg/logger"
)

type OrderService struct {
	Commands *commands.OrderCommands
}

func NewOrderService(
	log logger.Logger,
	mongoRepo repository.Repository,
) *OrderService {
	createOrderHandler := commands.NewCreateOrderHandler(log, mongoRepo)

	orderCommands := commands.NewOrderCommands(createOrderHandler)

	return &OrderService{
		Commands: orderCommands,
	}
}
