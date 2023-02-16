package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/omerkacamak/rentacar-golang/service"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := ctx.GetHeader("Authorization")

		tokenString := authHeader[len(BEARER_SCHEMA):] //bearerın sona git ve sonrasını al

		println("tokennn:: ", tokenString)
		token, err := service.NewAuthService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			println("Claims[Name]: ", claims["name"])
			println("Claims[Admin]: ", claims["admin"])
			println("Claims[Issuer]: ", claims["iss"])
			println("Claims[IssuedAt]: ", claims["iat"])
			println("Claims[ExpiresAt]: ", claims["exp"])
		} else {
			println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
