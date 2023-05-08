package repository_test

import (
	"testing"

	"github.com/DalvinCodes/cars/model"
	"github.com/DalvinCodes/cars/repository"
	"github.com/DalvinCodes/cars/utils"
	"github.com/stretchr/testify/assert"
)

var tesla = &model.Car{
	Make:     "Tesla",
	Model:    "Model 3",
	Package:  "Standard Range",
	Color:    "Mythical Black",
	Year:     2022,
	Category: "Sedan",
	Mileage:  999,
	Price:    51000000,
}
var mustang = &model.Car{
	Make:     "Ford",
	Model:    "Mustang",
	Package:  "Cobra",
	Color:    "Grey",
	Year:     2023,
	Category: "Coupe",
	Mileage:  14888,
	Price:    570000000,
}

var mercedes = &model.Car{
	Make:     "Mercedes Benz",
	Model:    "G63",
	Package:  "AMG",
	Color:    "Matte Black",
	Year:     2023,
	Category: "Sedan",
	Mileage:  74,
	Price:    1000000000,
}

func TestSave(t *testing.T) {
	repo := repository.NewRepo()

	car1 := tesla

	car1.ID = utils.GenerateID()

	if err := repo.Save(car1.ID, car1); err != nil {
		assert.NoError(t, err)
	}

	object, err := repo.Get(car1.ID)
	if err != nil {
		assert.NoError(t, err)
	}

	assert.Equal(t, object, car1)
}

func TestGet(t *testing.T) {
	repo := repository.NewRepo()

	t.Run("Test Get Object", func(t *testing.T) {
		car1 := mustang
		car1.ID = utils.GenerateID()

		if err := repo.Save(car1.ID, car1); err != nil {
			assert.NoError(t, err)
		}

		object, err := repo.Get(car1.ID)
		if err != nil {
			assert.NoError(t, err)
		}

		assert.Equal(t, object, car1)
	})

}

func TestUpdate(t *testing.T) {
	repo := repository.NewRepo()

	t.Run("Test Update Object", func(t *testing.T) {
		oldCar := mercedes
		oldCar.ID = utils.GenerateID()
		newCar := tesla

		if err := repo.Save(oldCar.ID, oldCar); err != nil {
			assert.NoError(t, err)
		}

		if err := repo.Update(oldCar.ID, newCar); err != nil {
			assert.NoError(t, err)
		}

		object, err := repo.Get(oldCar.ID)
		if err != nil {
			assert.NoError(t, err)
		}

		assert.Equal(t, object, newCar)
	})
}

func TestDelete(t *testing.T) {
	repo := repository.NewRepo()

	mockCar := tesla
	mockCar.ID = utils.GenerateID()

	if err := repo.Save(mockCar.ID, mockCar); err != nil {
		assert.NoError(t, err)
	}

	if err := repo.Delete(mockCar.ID); err != nil {
		assert.NoError(t, err)
	}

	object, err := repo.Get(mockCar.ID)
	if err != nil {
		assert.NoError(t, err)
	}

	assert.Nil(t, object)
}

func TestGetAll(t *testing.T) {
	repo := repository.NewRepo()

	carList := []*model.Car{tesla, mustang, mercedes}

	for _, car := range carList {
		car.ID = utils.GenerateID()
		if err := repo.Save(car.ID, car); err != nil {
			assert.NoError(t, err)
		}
	}

	cars, err := repo.GetAll()
	if err != nil {
		assert.NoError(t, err)
	}

	assert.Equal(t, len(cars), len(carList))
}
