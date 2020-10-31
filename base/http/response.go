package http

import "net/http"

type ErrorCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//0000: No Error
//1001: User not found
//1002: User no permissions
type BaseResponse struct {
	//Success
	IsSuccess bool `json:"is_success"`

	//Error
	Error ErrorCode   `json:"error"`
	Data  interface{} `json:"data"`
}

func (err ErrorCode) AsInvalidResponse() BaseResponse {

	return BaseResponse{
		IsSuccess: false,
		Error:     err,
		Data:      nil,
	}
}
func (err ErrorCode) AsValidResponse(data interface{}) BaseResponse {

	return BaseResponse{
		IsSuccess: true,
		Error:     err,
		Data:      data,
	}
}

var NOT_FOUND = ErrorCode{
	Code:    http.StatusBadRequest,
	Message: "Not method allowed!",
}
