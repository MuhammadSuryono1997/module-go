package utils

import (
	"math/rand"
	"regexp"
	"strings"
)

func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func MessageWA(code string) string {
	message := "Kode anda: *" + code + "* Jangan pernah memberikan kode ini ke siapapun, walaupun mereka mengatakan dari *OKTA POS!!*.\n\nJika anda tidak melakukan permintaan kode registrasi ini, maka abaikan pesan ini."
	return message
}

func MaskedNumber(nohp string) string {
	// mask := len(string) - 2
	re := regexp.MustCompile(`\b(\d{2})\d{10}\b`)
	s := re.ReplaceAllString(nohp, "$1**********")
	return s
}

func ToUpper(string string) string {
	return strings.ToUpper(string)
}
