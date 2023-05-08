package service_test

import (
	"testing"

	"github.com/DalvinCodes/cars/model"
	"github.com/DalvinCodes/cars/repository"
	"github.com/DalvinCodes/cars/service"
	"github.com/DalvinCodes/cars/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Run("Create car from service", func(t *testing.T) {
		mockRepo := repository.NewRepo()
		carService := service.NewCarService(mockRepo)

		car := &model.Car{
			Make:     "Toyota",
			Model:    "Camry",
			Package:  "LE",
			Color:    "White",
			Year:     2019,
			Category: "Sedan",
			Mileage:  10000,
			Price:    1500000,
		}

		car.ID = utils.GenerateID()

		if err := carService.CreateCar(car); err != nil {
			t.Error("error creating car")
		}

		obj, err := carService.GetCar(car.ID)
		if err != nil {
			t.Error("error getting car")
		}

		assert.Equal(t, obj.Make, "Toyota")
		assert.Equal(t, obj.Model, "Camry")
		assert.Equal(t, obj.Package, "LE")
		assert.Equal(t, obj.Color, "White")
		assert.Equal(t, obj.Year, 2019)
		assert.Equal(t, obj.Category, "Sedan")
		assert.Equal(t, obj.Mileage, 10000)
		assert.Equal(t, obj.Price, 1500000)

	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update car from service", func(t *testing.T) {
		mockRepo := repository.NewRepo()
		carService := service.NewCarService(mockRepo)

		oldCar := &model.Car{
			Make:     "Toyota",
			Model:    "Camry",
			Package:  "LE",
			Color:    "White",
			Year:     2019,
			Category: "Sedan",
			Mileage:  10000,
			Price:    1500000,
		}
		oldCar.ID = utils.GenerateID()

		if err := carService.CreateCar(oldCar); err != nil {
			t.Error("error creating car")
		}

		newCar := &model.Car{
			Make:     "Tesla",
			Model:    "Model 3",
			Package:  "Standard Range",
			Color:    "Red",
			Year:     2019,
			Category: "Sedan",
			Mileage:  567,
			Price:    45000000,
		}

		if err := carService.UpdateCar(oldCar.ID, newCar); err != nil {
			t.Error("error updating car")
		}

		object, err := carService.GetCar(oldCar.ID)
		if err != nil {
			t.Error("error getting car")
		}

		assert.Equal(t, object.Make, newCar.Make)
		assert.Equal(t, object.Model, newCar.Model)
		assert.Equal(t, object.Package, newCar.Package)
		assert.Equal(t, object.Color, newCar.Color)
		assert.Equal(t, object.Year, newCar.Year)
		assert.Equal(t, object.Category, newCar.Category)
		assert.Equal(t, object.Mileage, newCar.Mileage)
		assert.Equal(t, object.Price, newCar.Price)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete car from service", func(t *testing.T) {
		mockRepo := repository.NewRepo()
		carService := service.NewCarService(mockRepo)

		car := &model.Car{
			Make:     "Toyota",
			Model:    "Camry",
			Package:  "LE",
			Color:    "White",
			Year:     2019,
			Category: "Sedan",
			Mileage:  10000,
			Price:    1500000,
		}
		car.ID = utils.GenerateID()

		if err := carService.CreateCar(car); err != nil {
			t.Error("error creating car")
		}

		if err := carService.DeleteCar(car.ID); err != nil {
			t.Error("error deleting car")
		}

		delCar, err := carService.GetCar(car.ID)
		if err != nil {
			t.Error("error getting car")
		}

		assert.Nil(t, delCar)
	})

}

func TestGet(t *testing.T) {
	mockRepo := repository.NewRepo()
	carService := service.NewCarService(mockRepo)

	car := &model.Car{
		Make:     "Toyota",
		Model:    "Camry",
		Package:  "LE",
		Color:    "White",
		Year:     2019,
		Category: "Sedan",
		Mileage:  10000,
		Price:    1500000,
	}
	car.ID = utils.GenerateID()

	if err := carService.CreateCar(car); err != nil {
		t.Error("error creating car")
	}

	object, err := carService.GetCar(car.ID)
	if err != nil {
		t.Error("error getting car")
	}

	assert.Equal(t, object.Make, car.Make)
	assert.Equal(t, object.Model, car.Model)
	assert.Equal(t, object.Package, car.Package)
	assert.Equal(t, object.Color, car.Color)
	assert.Equal(t, object.Year, car.Year)
	assert.Equal(t, object.Category, car.Category)
	assert.Equal(t, object.Mileage, car.Mileage)
	assert.Equal(t, object.Price, car.Price)
}

func TestGetAll(t *testing.T) {
	mockRepo := repository.NewRepo()
	carService := service.NewCarService(mockRepo)

	car1 := &model.Car{
		Make:     "Toyota",
		Model:    "Camry",
		Package:  "LE",
		Color:    "White",
		Year:     2019,
		Category: "Sedan",
		Mileage:  10000,
		Price:    1500000,
	}
	car1.ID = utils.GenerateID()

	car2 := &model.Car{
		Make:     "Toyota",
		Model:    "Camry",
		Package:  "LE",
		Color:    "White",
		Year:     2019,
		Category: "Sedan",
		Mileage:  10000,
		Price:    1500000,
	}
	car2.ID = utils.GenerateID()

	_ = carService.CreateCar(car1)
	_ = carService.CreateCar(car2)

	cars, err := carService.GetCars()
	if err != nil {
		t.Error("error getting cars")
	}

	assert.Equal(t, len(cars), 2)
}
