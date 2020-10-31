package models

//Login credential
type TMerchant struct {
	DeviceId    string `json:"device_id"`
	PhoneNumber string `json:"phone_number"`
}

type TMerchantSecret struct {
	Secret       string `json:"secret"`
	RandomString int    `json:"random_string"`
	ExpiredTime  string `json:"expired_time"`
	CreatedOtp   string `json:"created_otp`
}
