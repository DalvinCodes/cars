package model

type Car struct {
	ID       string `json:"id"`
	Make     string `json:"make"`
	Model    string `json:"model"`
	Package  string `json:"package"`
	Color    string `json:"color"`
	Category string `json:"category"`
	Year     int    `json:"year"`
	Mileage  int    `json:"mileage"`
	Price    int    `json:"price"`
}
