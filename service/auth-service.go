package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

//import "github.com/golang-jwt/jwt/v4"

type AuthService interface {
	GenerateToken(email string, password string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

// type customClaims struct {
// 	Email string `json:"email"`
// 	jwt.StandardClaims
// }

type authService struct {
	secretKey string
	issuer    string
	service   UserService
}

func NewAuthService() AuthService {
	return &authService{
		secretKey: getSecretKey(),
		issuer:    "omerkacamak.com",
		service:   NewUserService(),
	}
}

func getSecretKey() string {
	secret := "secretjwtsecretjwt"
	if secret == "" {
		secret = "secretjwtsecretjwt"
	}
	return secret
}

func (authServ *authService) GenerateToken(email, password string) (string, error) {
	service := authServ.service
	result, err := service.FindByPassEmail(email, password)
	if err != nil {
		return "Kullanıcı yok", err
	} else {

		result.ExpiresAt = time.Now().Add(time.Minute * 2).Unix()
		result.IssuedAt = time.Now().Unix()
		result.Issuer = authServ.issuer

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, result)
		tokenString, err := token.SignedString([]byte(authServ.secretKey))
		if err != nil {
			println("--> " + err.Error())
			//return "hata var",err
			panic(err)

		}
		return tokenString, nil
	}

}

func (authServ *authService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(authServ.secretKey), nil
	})
}
