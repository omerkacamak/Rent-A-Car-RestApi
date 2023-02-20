package entity

import "github.com/golang-jwt/jwt/v4"

//import "github.com/golang-jwt/jwt/v4"

type AuthUser struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	FirstName string `json:"firstname" binding:"required" validate:"required,lent" `
	LastName  string `json:"lastname" binding:"required" validate:"required,deneme"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Token     string `json:"token"`

	jwt.StandardClaims `gorm:"-"`

	RoleId int   `json:"-"`
	Role   *Role `json:"role,omitempty" gorm:"foreignKey:RoleId"`
}
