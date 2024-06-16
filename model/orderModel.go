package model

import "gorm.io/gorm"

type Order struct {
    gorm.Model
    ID 			    uint		`gorm:"primary_key;auto_increment" json:"id"`
    Name   			string 		`gorm:"not null" json:"name"`
    Quantity 		int 		`gorm:"not null" json:"quantity"`
    Price           float64     `gorm:"not null" json:"price"`
    Status          string      `gorm:"not null" json:"status"`
    CustomerID  	uint   		`gorm:"index;not null" json:"customer_id"` // Foreign key
    Customer    	Customer   	`gorm:"foreignkey:CustomerID" json:"customer"`
}

type OrderResponse struct {
    ID              int         `json:"id"`
    Name   			string 		`json:"name"`
    Quantity 		int 		`json:"quantity"`
    Price           float64     `json:"price"`
    Status          string      `json:"status"`
}