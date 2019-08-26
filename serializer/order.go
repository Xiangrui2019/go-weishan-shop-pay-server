package serializer

import "go-weishan-shop-pay-server/models"

type Order struct {
	ID          uint    `json:"id"`
	Goodname    string  `json:"good_name"`
	GoodId      string  `json:"good_id"`
	Realname    string  `json:"real_name"`
	Address     string  `json:"address"`
	PhoneNumber string  `json:"phone_number"`
	ExtInfo     string  `json:"ext_info"`
	BuyCount    int     `json:"buy_count"`
	BuyPrice    float64 `json:"buy_price"`
	CreatedAt   int64   `json:"created_at"`
}

func BuildOrder(item *models.Order) *Order {
	return &Order{
		ID:          item.ID,
		Goodname:    item.Goodname,
		GoodId:      item.GoodId,
		Realname:    item.Realname,
		Address:     item.Address,
		PhoneNumber: item.Phonenumber,
		ExtInfo:     item.ExtInfo,
		BuyCount:    item.BuyCount,
		BuyPrice:    item.BuyPrice,
		CreatedAt:   item.CreatedAt.Unix(),
	}
}

func BuildOrders(items []models.Order) (orders []*Order) {
	for _, item := range items {
		order := BuildOrder(&item)
		orders = append(orders, order)
	}
	return orders
}
