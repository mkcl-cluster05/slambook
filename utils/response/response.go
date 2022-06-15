package r

type ErrorResponse struct {
	Status  int         `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type SuccessResponse struct {
	Status  int         `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"data,omitempty"`
}
