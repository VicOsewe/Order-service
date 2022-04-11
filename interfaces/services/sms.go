package interfaces

type SMS interface {
	SendSMS(message, phoneNumber string) error
}
