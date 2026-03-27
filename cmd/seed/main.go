package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/shivsperfect/uwe/db"
	"github.com/shivsperfect/uwe/types"
)

func main() {
	fmt.Println("Seeding database...")
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}

	db := db.Create()

	customer := types.Customer{
		ID: uuid.New(),
	}

	if err := db.CreateCustomer(&customer); err != nil {
		log.Fatalf("Error creating customer: %v\n", err)
	}

	fileUpload := types.FileUpload{
		ID:         uuid.New(),
		CustomerID: customer.ID,
		Mapping:    map[string]int{"amount": 0},
	}
	if err := db.CreateFileUpload(&fileUpload); err != nil {
		log.Fatalf("Error creating file upload: %v\n", err)
	}

}
