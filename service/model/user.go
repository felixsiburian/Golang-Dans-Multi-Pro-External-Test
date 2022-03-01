package model

import "time"

type User struct {
	ID           int64     `gorm:"primary_key;auto_increment" json:"id"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	PhoneNumber  string    `json:"phone_number"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	IsActive     bool      `json:"is_active"`
	CreatedBy    string    `json:"created_by"`
	CreatedDate  time.Time `json:"created_date"`
	ModifiedBy   string    `json:"modified_by"`
	ModifiedDate time.Time `json:"modified_date"`
	DeletedBy    string    `json:"deleted_by"`
	DeletedDate  time.Time `json:"deleted_date"`
}
