package service

import (
	"github.com/omerkacamak/rentacar-golang/repository"
)

// type GenericService interface {
// }
type IGenericService[Tentity any] interface {
	Save(tent Tentity) error
	Update(id int, tent Tentity) error
	Delete(tent Tentity) error
	FindAll() ([]Tentity, error)

	GetById(id int) (*Tentity, error)
}

type GenericService[Tentity any] struct {
	genericRepo repository.IGenericRepository[Tentity]
}

func NewGenericService[Tentity any]() IGenericService[Tentity] {
	return &GenericService[Tentity]{
		genericRepo: repository.NewGenericRepository[Tentity](),
	}
}

func (service *GenericService[Tentity]) Save(tent Tentity) error {
	return service.genericRepo.Save(tent)

}
func (service *GenericService[Tentity]) Update(id int, tent Tentity) error {
	return service.genericRepo.Update(id, tent)

}
func (service *GenericService[Tentity]) Delete(tent Tentity) error {
	return service.genericRepo.Delete(tent)

}
func (service *GenericService[Tentity]) FindAll() ([]Tentity, error) {
	return service.genericRepo.FindAll()

}
func (service *GenericService[Tentity]) GetById(id int) (*Tentity, error) {

	return service.genericRepo.GetById(id)
}
