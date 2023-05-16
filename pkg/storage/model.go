package storage

import (
	"gorm.io/gorm"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser()*User{
   return &User{} 
}


type Customer struct {
    gorm.Model
    FirstName    string    `gorm:"first_name"`
    LastName     string    `gorm:"last_name"`
    Email        string    `gorm:"email"`
    Phone        string    `gorm:"phone"`
    Token        string    `gorm:"token"`
    AddressLine1 string    `gorm:"address_line1"`
    AddressLine2 string    `gorm:"address_line2"`
    City         string    `gorm:"city"`
    State        string    `gorm:"state"`
    PostalCode   string    `gorm:"postal_code"`
    Country      string    `gorm:"country"`
    Product      []Product  `gorm:"foreignKey:CustomerID"`
}


type Product struct {
    gorm.Model
    Name        string      `gorm:"name"`
    Description string      `gorm:"description"`
    Price       float64     `gorm:"price"`
    CustomerID  uint            
}
