package service

import (
	"Employee/internal/employee/entity"
	"Employee/internal/employee/repository"
	"errors"
	"log"
)

type EmployeeService struct {
	repository repository.Employer
}

func NewEmployeeService(repository repository.Employer) *EmployeeService {
	return &EmployeeService{repository: repository}
}

type Employer interface {
	Add(in *entity.Employee) (uint, error)
	GetAll() ([]entity.Employee, error)
}

func (e *EmployeeService) Add(in *entity.Employee) (uint, error) {
	if err := e.repository.Exist(in.ID); err == nil {
		log.Printf("Employee %v already exists\n", in.ID)
		return 0, errors.New("employee already exists")
	}
	return e.repository.Add(in)
}

func (e *EmployeeService) GetAll() ([]entity.Employee, error) {
	employees, err := e.repository.GetAll()
	if err != nil {
		log.Println("Error getting all employees")
		return nil, err
	}
	return employees, err
}
