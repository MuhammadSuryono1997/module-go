package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MuhammadSuryono1997/module-go/utils"
	"io/ioutil"
	"net/http"
)

func RequestPost(body interface{}, url string)(interface{}, error)  {
	jsonReq, err := json.Marshal(body)
	resp, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	client := &http.Client{}
	req, err := client.Do(resp)

	if err != nil {
		fmt.Println(string(utils.ColorYellow()), err)
		return nil, err
	}
	body, _ = ioutil.ReadAll(req.Body)
	fmt.Println(string(utils.ColorYellow()), body)
	return body, nil

}
