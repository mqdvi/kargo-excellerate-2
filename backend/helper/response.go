package helper

type JsonResponse struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Error errorPayload `json:"error"`
}

type errorPayload struct {
	ErrorCode    string `json:"code"`
	ErrorMessage string `json:"message"`
}

func NewJsonResponse(data interface{}) JsonResponse {
	return JsonResponse{Data: data}
}

func NewErrorResponse(errorCode string, errorMessage string) ErrorResponse {
	return ErrorResponse{
		errorPayload{
			ErrorCode:    errorCode,
			ErrorMessage: errorMessage,
		},
	}
}
