package order

import "fmt"

type Order interface {
	Submit() error
	fmt.Stringer
}

type Item struct {
	ID string `json:"ID"`
	ModifierOptions []struct {
		ID       int    `json:"ID"`
		Name     string `json:"Name"`
		Quantity int    `json:"Quantity"`
		Price    string `json:"Price"`
	} `json:"ModifierOptions"`
	OnSaleNow bool   `json:"OnSaleNow"`
	Quantity  int    `json:"Quantity"`
	ItemName  string `json:"ItemName"`
	Price     string `json:"Price"`
}
