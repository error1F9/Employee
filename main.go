package main

import (
	"Employee/internal/config"
	"Employee/internal/dbase"
	"Employee/internal/employee/controller"
	"Employee/internal/employee/repository"
	"Employee/internal/employee/service"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := config.NewConfig()

	db, err := dbase.NewPostgersDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	rep := repository.NewEmployeeRepository(db)
	ser := service.NewEmployeeService(rep)
	contr := controller.NewEmployeeController(ser)

	http.HandleFunc("/employees/add", contr.Add)
	http.HandleFunc("/employees/getall", contr.GetAll)

	fmt.Println("Server is running on port 8080...")
	if err = http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
