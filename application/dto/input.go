package dto

import "github.com/VicOsewe/Order-service/domain"

//OrderInput ...
type OrderInput struct {
	Order        domain.Order          `json:"order"`
	OrderProduct []domain.OrderProduct `json:"order_product"`
}

//SMSResponse ...
type SMSResponse struct {
	SMSMessageData SMSMessageData `json:"SMSMessageData"`
}

//SMSMessageData ...
type SMSMessageData struct {
	Message    string       `json:"Message"`
	Recepients []Recepients `json:"Recipients"`
}

//Recepients ...
type Recepients struct {
	StatusCode int    `json:"statusCode"`
	Number     string `json:"number"`
	Status     string `json:"status"`
	Cost       string `json:"cost"`
	MessageID  string `json:"messageId"`
}
