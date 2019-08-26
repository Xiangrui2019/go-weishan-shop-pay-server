package models

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	Goodname    string
	GoodId      string
	Realname    string
	Address     string
	Phonenumber string
	ExtInfo     string
	BuyCount    int
	BuyPrice    float64
	Status      bool
}

func GetOrderById(id string) (*Order, error) {
	var order Order

	err := DB.First(&order, id).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func PublishOrder(id string) error {
	var order Order
	err := DB.First(&order, id).Error

	if err != nil {
		return err
	}

	order.Status = true
	err = DB.Save(&order).Error

	if err != nil {
		return err
	}

	return nil
}
