package register

import (
	"github.com/MuhammadSuryono1997/framework-okta/register/controllers"
	"github.com/MuhammadSuryono1997/framework-okta/register/services"
)

type RegisterService interface {
	services.RegisterServiceStatic
	services.RegisterService
}

type RegisterController interface {
	controllers.RegisterControllerStatic
}
