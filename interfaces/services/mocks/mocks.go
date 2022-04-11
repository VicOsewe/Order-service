package mocks

// SMSMocks mocks the sms layer
type SMSMocks struct {
	MockSendSMS func(message, phoneNumber string) error
}

// NewSMSMocks inits a new instance of SMS mocks with happy cases pre-defined
func NewSMSMocks() *SMSMocks {
	return &SMSMocks{
		MockSendSMS: func(message, phoneNumber string) error {
			return nil
		},
	}
}

func (s *SMSMocks) SendSMS(message, phoneNumber string) error {
	return s.MockSendSMS(message, phoneNumber)
}
