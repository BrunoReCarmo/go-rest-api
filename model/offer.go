package model 
// A model structure
// It looks like interface from JS
type Offer struct {
	//In the DB the column name is id_offer, so you need specify it
	ID int `json:"id_offer"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}