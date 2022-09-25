package views

import "net/http"

type Response struct {
	Status  int
	Message string
	Payload interface{}
	Error   interface{}
}

func BadRequestError(err error) *Response {
	return &Response{
		Status:  http.StatusBadRequest,
		Message: "BAD_REQUEST",
		Error:   err.Error(),
	}
}

func SuccessCreateResponse(payload interface{}, message string) *Response {
	return &Response{
		Status:  http.StatusCreated,
		Message: message,
		Payload: payload,
	}
}



func InternalServerError(err error) *Response {
	return &Response{
		Status:  http.StatusInternalServerError,
		Message: "INTERNAL_SERVER_ERROR",
		Error:   err.Error(),
	}
}

func DataNotFound(err error) *Response {
	return &Response{
		Status:  http.StatusNotFound,
		Message: "NO_DATA",
		Error:   err.Error(),
	}
}

func DataConflict(err error) *Response {
	return &Response{
		Status:  http.StatusConflict,
		Message: "DUPLICATE_DATA",
		Error:   err.Error(),
	}
}



func SuccessFindAllResponse(payload interface{}, message string) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: message,
		Payload: payload,
	}
}


func SuccessDeleteResponse(payload interface{}, message string) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: message,
		Payload: payload,
	}
}

func SuccessUpdateResponse(payload interface{}, message string) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: message,
		Payload: payload,
	}
}
