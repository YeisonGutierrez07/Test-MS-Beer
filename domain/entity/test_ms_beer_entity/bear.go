package test_ms_beer_entity

type Beer struct {
	ID       int64   `json:"id" binding:"required"`
	Name     string  `json:"name" binding:"required"`
	Brewery  string  `json:"brewery" binding:"required"`
	Country  string  `json:"country" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
}

type BoxPriceBeer struct {
	Bear       Beer    `json:"bear"`
	PriceTotal float64 `json:"price_total"`
}
