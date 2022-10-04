package dto

import domain "github.com/VicOsewe/Order-service/domain/dao"

type OrderInput struct {
	Order        domain.Order          `json:"order"`
	OrderProduct []domain.OrderProduct `json:"order_product"`
}

type SMSResponse struct {
	SMSMessageData struct {
		Message string `xml:"Message"`
	} `xml:"SMSMessageData"`
}

type SMSMessageData struct {
	Message string `json:"Message"`
}

type OtpInput struct {
	PhoneNumber string `json:"phone_number"`
}
