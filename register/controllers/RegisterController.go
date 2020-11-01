package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/MuhammadSuryono1997/module-go/base/database"
	_http "github.com/MuhammadSuryono1997/module-go/base/http"
	db "github.com/MuhammadSuryono1997/module-go/base/database"
	"github.com/MuhammadSuryono1997/module-go/register/models"
	"github.com/MuhammadSuryono1997/module-go/register/services"
	"github.com/MuhammadSuryono1997/module-go/utils"
	"github.com/gin-gonic/gin"
)

type RegisterController interface {
	RegisterUser(c *gin.Context) (string, string)
}

type RegisterMerchant struct {
	IsRegister bool `json:"is_register"`
	models.TMerchant
}

type registerController struct {
	registerService services.RegisterService
}

func RegisterHandler(registerService services.RegisterService) RegisterController {
	return &registerController{
		registerService: registerService,
	}
}

func (controller *registerController) RegisterUser(c *gin.Context) (string, string) {
	var credential *models.TMerchant
	var merchant RegisterMerchant

	if err := c.ShouldBindJSON(&credential); err != nil {
		return "", "Error input!"
	}

	err := db.GetDb().Table("t_merchants").Where("phone_number = ?", credential.PhoneNumber).Scan(&merchant)
	if err.RowsAffected > 0 {

		if !merchant.IsRegister {
			fmt.Println("Request OTP ....")
			_, er := RequestOTP(credential.PhoneNumber)
			if er != nil {
				return "", er.Error()
			}
			database.GetDb().Table("t_merchants").Where("phone_number = ?", credential.PhoneNumber).Updates(map[string]interface{}{"device_id": credential.DeviceId})

			return credential.PhoneNumber, ""
		}

		return "", _http.MessageIsRegistered
	}

	database.GetDb().Create(&credential)
	_, er := RequestOTP(credential.PhoneNumber)
	if er != nil {
		return "", er.Error()
	}

	return credential.PhoneNumber, ""

}

type RegisterControllerStatic interface {
	RegisterStatic(c *gin.Context) string
}

type registerControllerStatic struct {
	registerService services.RegisterServiceStatic
}

func RegisterHandlerStatic(registerService services.RegisterServiceStatic) RegisterControllerStatic {
	return &registerControllerStatic{
		registerService: registerService,
	}
}

func (controller *registerControllerStatic) RegisterStatic(ctx *gin.Context) string {
	var credential *models.TMerchant
	// resp.Header.Add("Authorization", BEARER)
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "Error input"
	}
	isUserAuthenticated := controller.registerService.RegisterStatic(credential.DeviceId, credential.PhoneNumber)
	if isUserAuthenticated {
		return "Number is registered"
	}
	return utils.MaskedNumber(credential.PhoneNumber)
}

func RequestOTP(nohp string) (string, error) {

	jsonReq, err := json.Marshal(map[string]interface{}{"phone_number": nohp})
	resp, err := http.NewRequest("POST", os.Getenv("URL_OTP"), bytes.NewBuffer(jsonReq))
	client := &http.Client{}
	req, err := client.Do(resp)

	if err != nil {
		fmt.Println(string(utils.ColorRed()), err)
		return "", err
	}
	body, _ := ioutil.ReadAll(req.Body)
	fmt.Println(string(utils.ColorCyan()), string(body))

	return _http.MessageSuccessRequest, nil

}
