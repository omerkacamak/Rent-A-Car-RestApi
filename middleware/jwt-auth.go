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
		if authHeader != "" {
			tokenString := authHeader[len(BEARER_SCHEMA):] //bearerın sona git ve sonrasını al
			println("tokenstrttr: " + tokenString)
			token, err := service.NewAuthService().ValidateToken(tokenString)
			if err != nil {
				//ctx.AbortWithStatus(http.StatusUnauthorized)

				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})

				return
			}
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
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
