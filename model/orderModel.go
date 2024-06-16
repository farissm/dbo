package model

import "gorm.io/gorm"

type Order struct {
    gorm.Model
    ID 			    uint		`gorm:"primary_key;auto_increment"`
    Name   			string 		`gorm:"not null",json:"name"`
    Quantity 		int 		`gorm:"not null",json:"quantity"`
    CustomerID  	uint   		`gorm:"index;not null",json:"customer_id"` // Foreign key
    Customer    	Customer   	`gorm:"foreignkey:CustomerID"`
}

type OrderResponse struct {
    ID              int         `json:"id"`
    Name   			string 		`json:"name"`
    Quantity 		int 		`json:"quantity"`
}