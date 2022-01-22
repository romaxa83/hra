package repository

import (
	"context"
	"github.com/romaxa83/hra/internal/order/models"
)

type Repository interface {
	CreateOrder(ctx context.Context, order *models.Order) (*models.Order, error)
}
