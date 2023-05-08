package main

import (
	"log"
	"net/http"

	"github.com/DalvinCodes/cars/controller"
	"github.com/DalvinCodes/cars/middleware"
	"github.com/DalvinCodes/cars/repository"
	"github.com/DalvinCodes/cars/service"
)

func main() {
	log.Println("Starting the application...")

	// create the dependencies
	repo := repository.NewRepo()
	carService := service.NewCarService(repo)
	carController := controller.NewCarController(carService)

	// register the handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/cars/create", middleware.AddHeaders(middleware.LoggingMiddleware(carController.CreateCarHandler)))
	mux.HandleFunc("/api/v1/cars/", middleware.AddHeaders(middleware.LoggingMiddleware(carController.GetCarHandler)))
	mux.HandleFunc("/api/v1/cars/all", middleware.AddHeaders(middleware.LoggingMiddleware(carController.GetCarsHandler)))
	mux.HandleFunc("/api/v1/cars/update", middleware.AddHeaders(middleware.LoggingMiddleware(carController.UpdateCarHandler)))
	mux.HandleFunc("/api/v1/cars/delete", middleware.AddHeaders(middleware.LoggingMiddleware(carController.DeleteCarHandler)))

	// start the server
	log.Println("Starting server and listening on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Println("server stopped...")
		panic(err)
	}
}
