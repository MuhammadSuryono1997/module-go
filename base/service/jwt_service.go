package service

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	privKeyPath = "keys/app.rsa"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func init_key() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)

}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
		fmt.Println(err.Error())
	}
}

func (service *jwtServices) ValidateKey() *rsa.PublicKey {
	return verifyKey
}

//jwt service
type JWTService interface {
	GenerateToken(obj interface{}) string
	ValidateToken(token string) (*jwt.Token, error)
}
type authCustomClaims struct {
	Data interface{}
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "OktaPos",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(obj interface{}) string {
	init_key()
	claims := &authCustomClaims{
		obj,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)

	t, err := token.SignedString(signKey)
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})

}
