package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DalvinCodes/cars/controller"
	"github.com/DalvinCodes/cars/model"
	"github.com/DalvinCodes/cars/repository"
	"github.com/DalvinCodes/cars/service"
	"github.com/DalvinCodes/cars/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateController(t *testing.T) {
	t.Run("Create car from controller", func(t *testing.T) {
		// Given
		rr := httptest.NewRecorder()
		mockStorage := repository.NewRepo()
		mockService := service.NewCarService(mockStorage)
		carController := controller.NewCarController(mockService)

		car := &model.Car{
			Make:     "Tesla",
			Model:    "Model 3",
			Package:  "LONG RANGE",
			Color:    "White",
			Year:     2023,
			Category: "Sedan",
			Mileage:  10,
			Price:    4500000,
		}
		car.ID = utils.GenerateID()

		testCar, err := json.Marshal(car)
		assert.NoError(t, err)

		// When
		req, err := http.NewRequest("POST", "/cars", bytes.NewBuffer(testCar))
		assert.NoError(t, err)

		handler := http.HandlerFunc(carController.CreateCarHandler)
		handler.ServeHTTP(rr, req)

		// Then
		assert.Equal(t, http.StatusCreated, rr.Code)

	})
}

func TestGetController(t *testing.T) {
	t.Run("Get car from controller", func(t *testing.T) {
		// Given
		rr := httptest.NewRecorder()
		mockStorage := repository.NewRepo()
		mockService := service.NewCarService(mockStorage)
		carController := controller.NewCarController(mockService)

		car := &model.Car{
			Make:     "Range Rover",
			Model:    "Sport",
			Package:  "HSE",
			Color:    "Grey",
			Year:     2023,
			Category: "Sedan",
			Mileage:  10,
			Price:    85000000,
		}
		car.ID = utils.GenerateID()
		_ = mockService.CreateCar(car)

		// When
		req, err := http.NewRequest("GET", "/cars?id="+car.ID, nil)
		assert.NoError(t, err)

		handler := http.HandlerFunc(carController.GetCarHandler)
		handler.ServeHTTP(rr, req)

		// Then
		expectedCar, err := json.Marshal(car)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, rr.Body.String(), string(expectedCar)+"\n")

	})
}

func TestGetAllController(t *testing.T) {
	t.Run("Get all cars from controller", func(t *testing.T) {
		// Given
		rr := httptest.NewRecorder()
		mockStorage := repository.NewRepo()
		mockService := service.NewCarService(mockStorage)
		carController := controller.NewCarController(mockService)

		car := &model.Car{
			Make:     "Mercedes",
			Model:    "C300",
			Package:  "AMG",
			Color:    "Black",
			Year:     2022,
			Category: "Sedan",
			Mileage:  100,
			Price:    6500000,
		}
		car.ID = utils.GenerateID()
		_ = mockService.CreateCar(car)

		// When
		req, err := http.NewRequest("GET", "/cars", nil)
		assert.NoError(t, err)

		handler := http.HandlerFunc(carController.GetCarsHandler)
		handler.ServeHTTP(rr, req)

		// Then
		expectedCar, err := json.Marshal([]*model.Car{car})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, rr.Body.String(), string(expectedCar)+"\n")

	})
}

func TestUpdateController(t *testing.T) {
	t.Run("Update car from controller", func(t *testing.T) {
		// Given
		rr := httptest.NewRecorder()
		mockStorage := repository.NewRepo()
		mockService := service.NewCarService(mockStorage)
		carController := controller.NewCarController(mockService)

		car := &model.Car{
			Make:     "Chevrolet",
			Model:    "Corvette",
			Package:  "Stingray",
			Color:    "Red",
			Year:     2022,
			Category: "Coupe",
			Mileage:  99,
			Price:    9900000,
		}
		car.ID = utils.GenerateID()
		_ = mockService.CreateCar(car)

		car.Color = "Black"
		testCar, err := json.Marshal(car)
		assert.NoError(t, err)

		// When
		req, err := http.NewRequest("PUT", "/cars?id="+car.ID, bytes.NewBuffer(testCar))
		assert.NoError(t, err)

		handler := http.HandlerFunc(carController.UpdateCarHandler)
		handler.ServeHTTP(rr, req)

		// Then
		assert.Equal(t, http.StatusAccepted, rr.Code)

	})
}

func TestDeleteController(t *testing.T) {
	t.Run("Delete car from controller", func(t *testing.T) {
		// Given
		rr := httptest.NewRecorder()
		mockStorage := repository.NewRepo()
		mockService := service.NewCarService(mockStorage)
		carController := controller.NewCarController(mockService)

		car := &model.Car{
			Make:     "BMW",
			Model:    "4 Series",
			Package:  "M Sport",
			Color:    "White",
			Year:     2021,
			Category: "Coupe",
			Mileage:  1000,
			Price:    8500000,
		}
		car.ID = utils.GenerateID()
		_ = mockService.CreateCar(car)

		// When
		req, err := http.NewRequest("DELETE", "/cars?id="+car.ID, nil)
		assert.NoError(t, err)

		handler := http.HandlerFunc(carController.DeleteCarHandler)
		handler.ServeHTTP(rr, req)

		// Then
		assert.Equal(t, http.StatusNoContent, rr.Code)

	})
}

