package interfaces

//SMS ...
type SMS interface {
	SendSMS(message, phoneNumber string) error
}
