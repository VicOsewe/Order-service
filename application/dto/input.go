package dto

import "github.com/VicOsewe/Order-service/domain"

type OrderInput struct {
	Order        domain.Order          `json:"order"`
	OrderProduct []domain.OrderProduct `json:"order_product"`
}

type SMSResponse struct {
	SMSMessageData SMSMessageData `json:"SMSMessageData"`
}
type SMSMessageData struct {
	Message    string       `json:"Message"`
	Recepients []Recepients `json:"Recipients"`
}
type Recepients struct {
	StatusCode int    `json:"statusCode"`
	Number     string `json:"number"`
	Status     string `json:"status"`
	Cost       string `json:"cost"`
	MessageID  string `json:"messageId"`
}
