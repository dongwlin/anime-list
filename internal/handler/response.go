package handler

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

type Pagination struct {
	TotalItems  int `json:"total_items"`
	TotalPages  int `json:"total_pages"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
}

func SuccessResponse(message string, data interface{}, meta interface{}) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}

func ErrorResponse(message string, errors interface{}) Response {
	return Response{
		Success: false,
		Message: message,
		Errors:  errors,
	}
}
