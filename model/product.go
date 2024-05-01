package model

import "time"

// stucture of product
type Product struct {
	Id          string `json:"id"`
	Name        string `json:"product_name"`
	Description string `json:"product_desc"`
	Location    string `json:"product_location"`
	Price       int    `json:"price"`
	Sales       []Sale `json:"sales"`
}

// structure of sale, this will be add into product when the owner want to add discount or other program
type Sale struct {
	Name    string    `json:"name"`
	Type    TypeSale  `json:"sale_type"`
	Expired time.Time `json:"sale_expiredAt"`
}

// thos structure will use to create sales type. Like will discount in percent or substract price
// example : if Name is discount in percent, price will multiply by the discount
// if product have price 50,000, and the product have sale type discount_in_percent at 5%, the price will :
// 50,000 * 5% = 47,500

// but when have discount_in_nominal at 2,000. the price will :
// 50,000 - 2,000 = 48,000

type TypeSale struct {
	Name    string    `json:"name"`
	Disc    float64   `json:"discount"`
	EndDate time.Time `json:"end_date"`
}

var Products []Product = []Product{}
var Sales []Sale = []Sale{}
