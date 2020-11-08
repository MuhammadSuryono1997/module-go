package services

import (
	"fmt"

	db "github.com/MuhammadSuryono1997/module-go/base/database"
	"github.com/MuhammadSuryono1997/module-go/register/models"
)

type RegisterService interface {
	RegisterUser(credential *models.TMerchant) bool
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
