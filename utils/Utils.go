package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorYellow = "\033[33m"
const colorBlue = "\033[34m"
const colorPurple = "\033[35m"
const colorCyan = "\033[36m"
const colorWhite = "\033[37m"
const formattime = time.RFC3339

func ToUpper(string string) string {
	return strings.ToUpper(string)
}

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

func TimeStamp() string {
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	return formatted
}

func CompareTime(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func TimeNow() string {
	return time.Now().Format(time.RFC3339)
}

func TimeAdd(minute time.Duration) string {
	return time.Now().Add(minute).Format(time.RFC3339)
}

func ColorYellow() string {
	return colorYellow
}

func ColorCyan() string {
	return colorCyan
}

func ColorRed() string {
	return colorRed
}
