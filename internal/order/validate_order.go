package order

import (
	"context"
	"errors"
	orders "github.com/romaxa83/hra/proto"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	ErrPreAuthorizationTimeout = errors.New("pre-authorization request timeout")
	ErrInventoryRequestTimeout = errors.New("check inventory request timeout")
	ErrItemOutOfStock          = errors.New("sorry one or more items in your order is out of stock")
)

// preAuthorizePayment выполняет предварительную авторизацию метода оплаты
// и возвращает ошибку. nil возвращается при успешной предварительной авторизации
func preAuthorizePayment(ctx context.Context, payment *orders.PaymentMethod, orderAmount float32) error {

	// Здесь выполняется дорогостоящая логика авторизации - для этого примера задействуем режим сна :-)
	// и вернем nil, чтобы указать успешную авторизацию
	timer := time.NewTimer(3 * time.Second)

	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		return ErrPreAuthorizationTimeout
	}
}

// checkInventory возвращает логическое значение и ошибку, указывающую,
// есть ли все товары на складе. (true, nil) возвращается, если
// все товары есть на складе, и не возникло никаких ошибок
func checkInventory(ctx context.Context, items []*orders.Item) (bool, error) {

	// Здесь выполняется дорогостоящая логика инвентаризации - для этого примера задействуем режим сна :-)
	timer := time.NewTimer(2 * time.Second)

	select {
	case <-timer.C:
		return true, nil
	case <-ctx.Done():
		return false, ErrInventoryRequestTimeout
	}
}

// getOrderTotal высчитывает общую сумму заказа
func getOrderTotal(items []*orders.Item) float32 {
	var total float32

	for _, item := range items {
		total += item.Price
	}

	return total
}

func validateOrder(ctx context.Context, items []*orders.Item, payment *orders.PaymentMethod) error {
	g, errCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return preAuthorizePayment(errCtx, payment, getOrderTotal(items))
	})

	g.Go(func() error {
		itemsInStock, err := checkInventory(errCtx, items)
		if err != nil {
			return err
		}
		if !itemsInStock {
			return ErrItemOutOfStock
		}
		return nil
	})

	return g.Wait()
}
