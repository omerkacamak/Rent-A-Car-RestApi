package repository

import (
	"github.com/omerkacamak/rentacar-golang/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user entity.AuthUser) error
	Update(user entity.AuthUser) error
	Delete(user entity.AuthUser) error
	FindAll() []entity.AuthUser

	FindByPassEmail(email, password string) (entity.AuthUser, error)
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {

	return &userRepository{
		connection: DB,
	}
}

func (userrep *userRepository) Save(user entity.AuthUser) error {
	result := userrep.connection.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (userrep *userRepository) Update(user entity.AuthUser) error {
	result := userrep.connection.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func (userrep *userRepository) Delete(user entity.AuthUser) error {
	result := userrep.connection.Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func (userrep *userRepository) FindAll() []entity.AuthUser {
	var users []entity.AuthUser
	userrep.connection.Set("gorm:auto_preload", true).Find(&users)
	for _, value := range users {
		println("-->repo " + value.FirstName)
	}
	return users
}
func (userrep *userRepository) FindByPassEmail(email, password string) (entity.AuthUser, error) {
	var user entity.AuthUser
	res := userrep.connection.Where(&entity.AuthUser{Email: email, Password: password}).First(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}
