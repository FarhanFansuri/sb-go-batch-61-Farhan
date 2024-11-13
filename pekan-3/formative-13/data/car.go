package data

type Car struct {
	CarID string `json:car_id`
	Brand string `json:brand`
	Model string `json:model`
	Price int    `json:price`
}

var Cars_data = []Car{}
