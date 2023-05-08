package service

import (
	"github.com/DalvinCodes/cars/model"
	"github.com/DalvinCodes/cars/repository"
	"github.com/DalvinCodes/cars/utils"
)

type CarService interface {
	CreateCar(car *model.Car) error
	UpdateCar(id string, car *model.Car) error
	DeleteCar(id string) error
	GetCar(id string) (*model.Car, error)
	GetCars() ([]*model.Car, error)
}

type carService struct {
	repo repository.Storage
}

func NewCarService(repo repository.Storage) *carService {
	return &carService{
		repo: repo,
	}
}

func (c *carService) GetCar(id string) (*model.Car, error) {
	return c.repo.Get(id)
}

func (c *carService) GetCars() ([]*model.Car, error) {
	return c.repo.GetAll()
}

func (c *carService) CreateCar(car *model.Car) error {
	car.ID = utils.GenerateID()
	return c.repo.Save(car.ID, car)
}

func (c *carService) UpdateCar(id string, car *model.Car) error {
	car.ID = id
	return c.repo.Update(id, car)
}

func (c *carService) DeleteCar(id string) error {
	return c.repo.Delete(id)
}
