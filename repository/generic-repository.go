package repository

import "gorm.io/gorm"

type IGenericRepository[Tentity any] interface {
	Save(tent Tentity) error
	Update(id int, tent Tentity) error
	Delete(tent Tentity) error
	FindAll() ([]Tentity, error)

	GetById(id int) (*Tentity, error)
}

type GenericRepository[Tentity any] struct {
	connection *gorm.DB
}

func NewGenericRepository[Tentity any]() IGenericRepository[Tentity] {
	return &GenericRepository[Tentity]{
		connection: DB,
	}
}
func (repo *GenericRepository[Tentity]) Save(tent Tentity) error {
	err := repo.connection.Create(&tent).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *GenericRepository[Tentity]) Update(id int, tent Tentity) error {
	err := repo.connection.Where("id = ?", id).Updates(&tent).Error
	if err != nil {
		return err
	}
	return nil
}
func (repo *GenericRepository[Tentity]) Delete(tent Tentity) error {
	err := repo.connection.Delete(&tent).Error
	if err != nil {
		return err
	}
	return nil
}
func (repo *GenericRepository[Tentity]) FindAll() ([]Tentity, error) {
	var entities []Tentity
	err := repo.connection.Set("gorm:auto_preload", true).Find(&entities).Error
	if err != nil {
		return entities, err
	}
	return entities, nil
}

func (repo *GenericRepository[Tentity]) GetById(id int) (*Tentity, error) {
	var entity Tentity

	err := repo.connection.First(&entity, id).Error
	if err != nil {
		println("genericrepo hataaaa")
		return &entity, err
	}
	println("genericrepo doğruuuuuu")
	return &entity, nil
}
