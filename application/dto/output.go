package dto

//APIResponse ...
type APIResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Body       interface{}
}

//APIFailureResponse ...
type APIFailureResponse struct {
	APIResponse APIResponse `json:"api_response"`
	Error       string      `json:"error"`
	//to change to enum
	ErrorType string `json:"error_type"`
}
