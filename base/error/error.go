package error

type ErrorCode struct {
	Code    int
	Message string
}

var NOT_FOUND = ErrorCode{
	Code:    0001,
	Message: "",
}

var SUCCESS = ErrorCode{
	Code:    0000,
	Message: "",
}
