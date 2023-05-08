package utils_test

import (
	"log"
	"testing"

	utils "github.com/DalvinCodes/cars/utils"
)

func TestGenerateID(t *testing.T) {
	id := utils.GenerateID()
	if id == "" {
		t.Error("Error generating ID")
	}

	log.Printf("ID: %s", id)

	if len(id) > 19 {
		t.Error("ID not 19 characters long or less")
	}

	
}
