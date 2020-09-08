package models

type Product struct {
	ID int `json:"idproduct"`
	Code string `json:"code"`
	Name string `json:"name"`
}