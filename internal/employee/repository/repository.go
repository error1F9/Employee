package repository

import (
	"Employee/internal/employee/entity"
	"errors"
	"gorm.io/gorm"
	"log"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

type Employer interface {
	Add(emp *entity.Employee) (uint, error)
	GetAll() ([]entity.Employee, error)
	Exist(id uint) error
}

func (r *EmployeeRepository) Add(emp *entity.Employee) (uint, error) {
	if err := r.db.Create(&emp).Error; err != nil {
		log.Printf("error adding employee: %v\n", err)
		return 0, err
	}
	return emp.ID, nil
}

func (r *EmployeeRepository) GetAll() ([]entity.Employee, error) {
	employees := make([]entity.Employee, 0)
	err := r.db.Find(&employees).Error
	if err != nil {
		log.Printf("error getting all employees: %v\n", err)
		return nil, err
	}
	return employees, nil
}

func (r *EmployeeRepository) Exist(id uint) error {
	if err := r.db.Where("id = ?", id).First(&entity.Employee{}).Error; err == nil {
		return nil
	}
	return errors.New("employee already exist")
}
