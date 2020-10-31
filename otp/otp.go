package otp

import (
	"bytes"
	"encoding/base32"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"

	"github.com/MuhammadSuryono1997/framework-okta/utils"
	"github.com/hgfischer/go-otp"
	"github.com/xlzd/gotp"
)

const URL_OTP = "http://localhost:5005/request-otp"

var (
	secret   = flag.String("secret", "OTPOktaPOS", "Secret key")
	isBase32 = flag.Bool("base32", true, "If true, the secret is interpreted as a Base32 string")
	length   = flag.Uint("length", 4, "OTP length")
	period   = flag.Uint("period", otp.DefaultPeriod, "Period in seconds")
	counter  = flag.Uint64("counter", 0, "Counter")
)

var totp = &otp.TOTP{
	Secret:         *secret,
	Length:         uint8(*length),
	Period:         uint8(*period),
	IsBase32Secret: true,
}

func GenerateOTP(expired uint8) string {
	flag.Parse()
	key := *secret

	if !*isBase32 {
		key = base32.StdEncoding.EncodeToString([]byte(*secret))
	}

	key = strings.ToUpper(key)
	if !isGoogleAuthenticatorCompatible(key) {
		fmt.Println(string(utils.ColorYellow()), "WARN: Google Authenticator requires 16 chars base32 secret, without padding")
	}

	fmt.Println(string(utils.ColorCyan()), "Secret Base32 Encoded Key: ", key)

	totp.Secret = key
	totp.Period = expired

	return totp.Now().Get()
}

func ValidateOTP(otp string) bool {
	return totp.Now().Verify(otp)
}

func isGoogleAuthenticatorCompatible(base32Secret string) bool {
	cleaned := strings.Replace(base32Secret, "=", "", -1)
	cleaned = strings.Replace(cleaned, " ", "", -1)
	return len(cleaned) == 16
}

func GenerateHOTP() (string, string, int) {
	secret := gotp.RandomSecret(32)
	rand := rand.Intn(100)
	otp := gotp.NewHOTP(secret, 4, nil)

	return otp.At(rand), secret, rand
}

func RequestOTP(nohp string) (string, int) {
	otp, secret, rand := GenerateHOTP()
	send, err := SendToWA(nohp, otp)

	if err != nil {
		fmt.Println(string(utils.ColorYellow()), err)
	}
	fmt.Println(string(utils.ColorYellow()), "OTP SUCCESS SENDING TO "+send)

	return secret, rand
}

func ValidateHOTP(secret string, rand int, otp string) bool {
	hotp := gotp.NewHOTP(secret, 4, nil)
	valid := hotp.Verify(otp, rand)

	return valid
}

func ResendOTP(nohp string) (string, error) {
	jsonReq, err := json.Marshal(map[string]interface{}{"phone_number": nohp})
	resp, err := http.NewRequest("POST", URL_OTP, bytes.NewBuffer(jsonReq))
	client := &http.Client{}
	req, err := client.Do(resp)

	if err != nil {
		fmt.Println(string(utils.ColorYellow()), err)
		return "", err
	}
	body, _ := ioutil.ReadAll(req.Body)
	fmt.Println(string(utils.ColorCyan()), string(body))

	return "Success request", nil

}
