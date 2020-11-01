package http

const (
	NoError = 0000

	// CODE REGISTER 1
	SuccessRegister    = 1000
	NumberIsRegistered = 1001

	// CODE REQUEST
	PhoneNumberNotFound = 2001
	DeviceIdNotFound = 2002


	// CODE ERROR SYSTEM
	MetodeNotAllowed = 4005
	BadRequest       = 4004
)

const (
	// Eror Message
	MessageMetodeNotAllowed = "Metode not allowed"
	MessageErrorInput       = "Error input"
	MessageErrorRequest       = "Error request"
	MessageErrorLoadEnv     = ".env file not found!"
	MessageTokenInvalid     = "Token invalid"
	MessageIsRegistered = "Number is registered"
	MessagePhoneNumberNotFound = "Phone number not found"
	MessageDeviceIdNotFound = "Device id not found"
	MessageErrorRequestOtp = "Error request otp"

	// General Message
	MessageInformation = "Informasi service"

	// Success Message
	MessageSuccessRequest = "Success request"
	MessageSuccessRegister = "Success register"
)

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
	Code:    MetodeNotAllowed,
	Message: MessageMetodeNotAllowed,
}
