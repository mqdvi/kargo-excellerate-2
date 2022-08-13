package helper

type DataResponse struct {
	Data interface{} `json:"data"`
}

type ItemsReponse struct {
	Data Data `json:"data"`
}

type Data struct {
	Items interface{} `json:"items,omitempty"`
}

type ErrorResponse struct {
	Error errorPayload `json:"error"`
}

type errorPayload struct {
	ErrorCode    string `json:"code"`
	ErrorMessage string `json:"message"`
}

func NewJsonResponse(data interface{}) DataResponse {
	return DataResponse{Data: data}
}

func NewErrorResponse(errorCode string, errorMessage string) ErrorResponse {
	return ErrorResponse{
		errorPayload{
			ErrorCode:    errorCode,
			ErrorMessage: errorMessage,
		},
	}
}
