package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DalvinCodes/cars/model"
	"github.com/DalvinCodes/cars/service"
	"github.com/DalvinCodes/cars/utils"
)

type CarController interface {
	CreateCarHandler(w http.ResponseWriter, r *http.Request)
	DeleteCarHandler(w http.ResponseWriter, r *http.Request)
	GetCarHandler(w http.ResponseWriter, r *http.Request)
	GetCarsHandler(w http.ResponseWriter, r *http.Request)
	UpdateCarHandler(w http.ResponseWriter, r *http.Request)
}

type CarsController struct {
	service service.CarService
}

func NewCarController(service service.CarService) *CarsController {
	return &CarsController{
		service: service,
	}
}

func (c *CarsController) CreateCarHandler(w http.ResponseWriter, r *http.Request) {
	var car *model.Car

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Printf("error occured due to invalid request method. Method: %s\n", r.Method)
		return
	}

	// decode the request body into the car struct
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		log.Printf("error while decoding the car data. err: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(utils.ErrCreatingObiect.Error()))
		return
	}

	// create the car
	if err := c.service.CreateCar(car); err != nil {
		log.Printf("error while creating the car. err: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(utils.ErrCreatingObiect.Error()))
		return
	}

	// return the success status code
	w.WriteHeader(http.StatusCreated)

}

func (c *CarsController) DeleteCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Printf("error occured due to invalid request method. Method: %s\n", r.Method)
		return
	}

	// get the car id from the url
	id := r.URL.Query().Get("id")

	// delete the car
	if err := c.service.DeleteCar(id); err != nil {
		log.Printf("error while deleting the car. err: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(utils.ErrDeletingObject.Error()))
		return
	}

	// return the success status code
	w.WriteHeader(http.StatusNoContent)
}

func (c *CarsController) GetCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Printf("error occured due to invalid request method. Method: %s\n", r.Method)
		return
	}

	// get the car id from the url
	id := r.URL.Query().Get("id")

	// get the car
	car, err := c.service.GetCar(id)
	if err != nil {
		log.Printf("error while getting the car. err: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(utils.ErrRetrievingObject.Error()))
		return
	}

	// encode the car into the response body
	if err := json.NewEncoder(w).Encode(car); err != nil {
		log.Printf("error while encoding the car into the response body. err: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(utils.ErrRetrievingObject.Error()))
		return
	}

	// return the success status code
	w.WriteHeader(http.StatusOK)
}

func (c *CarsController) GetCarsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Printf("error occured due to invalid request method. Method: %s\n", r.Method)
		return
	}

	// get all the cars
	cars, err := c.service.GetCars()
	if err != nil {
		log.Printf("error while getting the cars. err: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// encode the cars into the response body
	if err := json.NewEncoder(w).Encode(cars); err != nil {
		log.Printf("error while encoding the cars into the response body. err: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// return the success status code
	w.WriteHeader(http.StatusOK)
}

func (c *CarsController) UpdateCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Printf("error occured due to invalid request method. Method: %s\n", r.Method)
		return
	}

	var car *model.Car

	// get the car id from the url
	id := r.URL.Query().Get("id")

	// decode the request body into the car struct
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		log.Printf("error while decoding the car data. err: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// update the car
	if err := c.service.UpdateCar(id, car); err != nil {
		log.Printf("error while updating the car. err: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// return the success status code
	w.WriteHeader(http.StatusAccepted)
}
