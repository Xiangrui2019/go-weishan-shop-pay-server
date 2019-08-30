package global

type OrderCache struct {
	Goodname    string  `json:"good_name"`
	GoodId      string  `json:"good_id"`
	Realname    string  `json:"real_name"`
	Address     string  `json:"address"`
	Phonenumber string  `json:"phone_number"`
	ExtInfo     string  `json:"extinfo"`
	BuyCount    int     `json:"buy_count"`
	BuyPrice    float64 `json:"buy_price"`
	SelfMention bool    `json:"self_mention"`
}
