package dto

import "github.com/VicOsewe/Order-service/domain"

type OrderInput struct {
	Order        domain.Order          `json:"order"`
	OrderProduct []domain.OrderProduct `json:"order_product"`
}
