package services

import (
	"fmt"

	db "github.com/MuhammadSuryono1997/framework-okta/base/database"

	"github.com/MuhammadSuryono1997/framework-okta/register/models"
)

type RegisterService interface {
	RegisterUser(credential *models.TMerchant) bool
}

type RegisterServiceStatic interface {
	RegisterStatic(devid string, nohp string) bool
}

type registerInformation struct {
	device_id string
	no_hp     string
}

func StaticRegisterService() RegisterServiceStatic {
	return &registerInformation{
		device_id: "123456789",
		no_hp:     "0895355698652",
	}
}

func RegisterUser(credential *models.TMerchant) bool {
	var merchant []models.TMerchant

	err := db.GetDb().Where("no_hp = ?", credential.PhoneNumber).First(&merchant)
	if err == nil {
		return false
	}

	fmt.Println(merchant)

	return true
}

func (info *registerInformation) RegisterStatic(devid string, nohp string) bool {
	return "123456789" == devid && "0895355698652" == nohp
}
