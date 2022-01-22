package models

import "time"

type Order struct {
	OrderID   string    `json:"orderId" bson:"_id,omitempty"`
	Name      string    `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=3,max=250"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
