package shops

import (
	"gorm.io/gorm"
)

/**
Structures representing the data model of each of the shops
*/
type (
	SuperMarket struct {
		gorm.Model
		Name    string `gorm:"not null" json:"name"`
		Gender  string `gorm:"not null" json:"gender"`
		Email   string `json:"email"`
		PhoneNo string `gorm:"not null" json:"phoneNo"`
	}

	Electronic struct {
		gorm.Model
		Name     string `gorm:"not null" json:"name"`
		Email    string `json:"email"`
		PhoneNo  string `gorm:"not null" json:"phoneNo"`
		Address  string `gorm:"not null" json:"address"`
		Location string `json:"location"`
	}

	Mobile struct {
		gorm.Model
		Name     string `gorm:"not null" json:"name"`
		Gender   string `gorm:"not null" json:"gender"`
		Email    string `json:"email"`
		PhoneNo  string `gorm:"not null"  json:"phoneNo"`
		Address  string `gorm:"not null"  json:"address"`
		Location string `json:"location"`
		Hobby    string `json:"Hobby"`
		Interest string `json:"interest"`
	}

	Bakery struct {
		gorm.Model
		Name    string `gorm:"not null" json:"name"`
		Gender  string `gorm:"not null" json:"gender"`
		Email   string `json:"email"`
		PhoneNo string `gorm:"not null" json:"phoneNo"`
	}
)
