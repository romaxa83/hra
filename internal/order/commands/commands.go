package commands

import "time"

type OrderCommands struct {
	CreateOrder CreateOrderCmdHandler
}

func NewOrderCommands(createOrder CreateOrderCmdHandler) *OrderCommands {
	return &OrderCommands{
		CreateOrder: createOrder,
	}
}

type CreateOrderCmd struct {
	OrderID   string    `json:"orderId" bson:"_id,omitempty"`
	Name      string    `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=3,max=250"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

func NewCreateOrderCmd(orderId, name string, createdAt time.Time) *CreateOrderCmd {
	return &CreateOrderCmd{
		OrderID:   orderId,
		Name:      name,
		CreatedAt: createdAt,
	}
}
