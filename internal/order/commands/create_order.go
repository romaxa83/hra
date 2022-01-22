package commands

import (
	"context"
	"github.com/romaxa83/hra/internal/order/models"
	"github.com/romaxa83/hra/internal/order/repository"
	"github.com/romaxa83/hra/pkg/logger"
)

type CreateOrderCmdHandler interface {
	Handle(ctx context.Context, cmd CreateOrderCmd) error
}

type createOrderHandler struct {
	log       logger.Logger
	mongoRepo repository.Repository
}

func NewCreateOrderHandler(logger logger.Logger, mongoRepo repository.Repository) *createOrderHandler {
	return &createOrderHandler{
		logger,
		mongoRepo,
	}
}

func (h *createOrderHandler) Handle(ctx context.Context, cmd CreateOrderCmd) error {

	h.log.Warn("COMMAND CREATE ORDER")

	order := &models.Order{
		OrderID:   cmd.OrderID,
		Name:      cmd.Name,
		CreatedAt: cmd.CreatedAt,
	}

	created, err := h.mongoRepo.CreateOrder(ctx, order)
	if err != nil {
		h.log.Errorf("Error %+v", err)
		return err
	}
	h.log.Info("CREATED ORDER [%s]", created.OrderID)

	return nil
}
