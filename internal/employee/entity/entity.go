package entity

type Employee struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Year string `json:"year"`
}
