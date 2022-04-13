package dto

type APIResponse struct {
	Message string `json:"message"`
	Body    interface{}
}

type APIFailureResponse struct {
	APIResponse APIResponse `json:"api_response"`
	Error       string      `json:"error"`
	//to change to enum
	ErrorType string `json:"error_type"`
}
