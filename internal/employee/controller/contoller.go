package controller

import (
	"Employee/internal/employee/entity"
	"Employee/internal/employee/service"
	"encoding/json"
	"log"
	"net/http"
)

type EmployeeController struct {
	service service.Employer
}

func NewEmployeeController(service service.Employer) *EmployeeController {
	return &EmployeeController{service}
}

func (c *EmployeeController) Add(w http.ResponseWriter, r *http.Request) {
	var employee entity.Employee
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "Invalid input",
			Data:    nil,
		})
		log.Println("Invalid input")
		return
	}

	id, err := c.service.Add(&employee)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "error adding employee",
			Data:    nil,
		})
		log.Println("error adding employee")
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: "employee successfully added",
		Data:    id,
	})
	log.Println("employee successfully added")
}

func (c *EmployeeController) GetAll(w http.ResponseWriter, r *http.Request) {

	employees, err := c.service.GetAll()
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "error getting employees",
			Data:    nil,
		})
		log.Println("error getting employees")
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: "employees list",
		Data:    employees,
	})
	log.Println("employees list successfully")
}
