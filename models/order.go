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

func CreateOrder(order *Order) (*Order, error) {
	res := DB.Create(order)
	if res.Error != nil {
		return nil, res.Error
	}

	return res.Value.(*Order), nil
}

func ListOrder() ([]Order, error) {
	var orders []Order

	err := DB.Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
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
