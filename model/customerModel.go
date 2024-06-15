package model

import "gorm.io/gorm"

type Customer struct {
    gorm.Model
	ID 			uint		`gorm:"primary_key;auto_increment"`
	Username 	string		`gorm:"not null",json:"username"`
	Password 	string		`gorm:"not null",json:"password"`
    FirstName   string		`gorm:"not null",json:"firstname"`
	LastName   	string		`gorm:"not null",json:"lastname"`
	Email   	string		`gorm:"not null",json:"email"`
	Address   	string 		`gorm:"not null",json:"address"`
    Order    	[]Order 
}