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

type SignUpInput struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	OTP         string `json:"otp"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
}
