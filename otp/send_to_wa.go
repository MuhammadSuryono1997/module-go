package otp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/MuhammadSuryono1997/framework-okta/utils"
)

const URL_WA = "http://139.59.241.153:6780/message/send"
const BEARER = "Bearer " + "a4546515d2955d582da64ae295bb8c22871b3ac481cfd7b0b0c29cbf16fa48fd"
const SHORT_NAME = "creative-ocean"

type StructWA struct {
	To        string `json:"to"`
	ShortName string `json:"short_name"`
	Message   string `json:"message"`
}

func SendToWA(nohp string, otp string) (string, error) {
	pesan := &StructWA{
		To:        nohp,
		ShortName: SHORT_NAME,
		Message:   utils.MessageWA(otp),
	}

	jsonReq, err := json.Marshal(pesan)
	resp, err := http.NewRequest("POST", URL_WA, bytes.NewBuffer(jsonReq))
	resp.Header.Add("Authorization", BEARER)
	client := &http.Client{}
	req, err := client.Do(resp)

	if err != nil {
		fmt.Println(string(utils.ColorYellow()), err)
		return "", err
	}
	body, _ := ioutil.ReadAll(req.Body)
	fmt.Println(string(utils.ColorCyan()), string(body))
	return utils.MaskedNumber(nohp), nil
}
