package repository

import (
	"sync"

	"github.com/DalvinCodes/cars/model"
)

type Storage interface {
	Save(key string, object *model.Car) error
	Get(key string) (*model.Car, error)
	GetAll() ([]*model.Car, error)
	Delete(key string) error
	Update(key string, value *model.Car) error
}

type Repo struct {
	db map[string]*model.Car
	sync.RWMutex
}

func NewRepo() *Repo {
	return &Repo{
		db: make(map[string]*model.Car),
	}
}

func (r *Repo) Save(key string, object *model.Car) error {
	r.Lock()
	defer r.Unlock()
	r.db[key] = object
	return nil
}

func (r *Repo) Get(key string) (*model.Car, error) {
	r.RLock()
	defer r.RUnlock()
	return r.db[key], nil
}

func (r *Repo) GetAll() ([]*model.Car, error) {
	r.RLock()
	defer r.RUnlock()
	var cars []*model.Car
	for _, car := range r.db {
		cars = append(cars, car)
	}
	return cars, nil
}

func (r *Repo) Delete(key string) error {
	r.Lock()
	defer r.Unlock()
	delete(r.db, key)
	return nil
}

func (r *Repo) Update(key string, object *model.Car) error {
	car, err := r.Get(key)
	if err != nil {
		return err
	}

	defer r.Unlock()
	r.Lock()

	r.db[car.ID] = object
	return nil
}
