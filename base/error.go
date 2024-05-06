package base

type ErrorResponse struct {
	Status    string `json:"status"`
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}
