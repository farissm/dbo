package model

import "gorm.io/gorm"

type Order struct {
    gorm.Model
    ID 			    uint		`gorm:"primary_key;auto_increment"`
    Name   			string 		`gorm:"not null"`
    Quantity 		int 		`gorm:"not null"`
    CustomerID  	uint   		`gorm:"index;not null"` // Foreign key
    Customer    	Customer   	`gorm:"foreignkey:CustomerID"` // Belongs to User
}