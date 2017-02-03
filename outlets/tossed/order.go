package tossed

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"
	"io/ioutil"

	"github.com/itspage/feedme/order"
)

type orderBody struct {
	PromoCodeDiscount string        `json:"PromoCodeDiscount"`
	TelNo             string        `json:"TelNo"`
	UserID            string        `json:"UserID"`
	Email             string        `json:"Email"`
	DeliveryCost      string        `json:"DeliveryCost"`
	Name              string        `json:"Name"`
	RequiredDate      string        `json:"RequiredDate"`
	OrderItems        []*order.Item `json:"OrderItems"`
	APIKey            string        `json:"APIKey"`
	Payment struct {
		CardNo         string `json:"CardNo"`
		CardAddress    string `json:"CardAddress"`
		CardCVV        string `json:"CardCVV"`
		CardStartDate  string `json:"CardStartDate"`
		CardExpiryDate string `json:"CardExpiryDate"`
		NameOnCard     string `json:"NameOnCard"`
		PaymentType    int    `json:"PaymentType"`
		CardType       string `json:"CardType"`
		CardPostcode   string `json:"CardPostcode"`
	} `json:"Payment"`
	OrderTotal           string `json:"OrderTotal"`
	SiteID               string `json:"SiteID"`
	PromoCode            string `json:"PromoCode"`
	DeliveryInstructions string `json:"DeliveryInstructions"`
	DeliveryAddress      string `json:"DeliveryAddress"`
	RequiredTime         string `json:"RequiredTime"`
	DeliveryPostcode     string `json:"DeliveryPostcode"`
	OrderType            int    `json:"OrderType"`
}

type tossedOrder struct {
	Request *http.Request
	body    *orderBody
}

func (o *tossedOrder) Submit() error {
	log.Printf("sending request %v", o.Request)
	rsp, err := http.DefaultClient.Do(o.Request)
	if err != nil {
		return err
	}
	if rsp.StatusCode != 200 {
		return fmt.Errorf("expected 200 status from order, got %v", rsp.StatusCode)
	}

	type orderResponse struct {
		ResultString string `json:"ResultString"`
		OrderID      int    `json:"OrderID"`
	}

	rb, _ := ioutil.ReadAll(rsp.Body)
	orderPlaced := &orderResponse{}
	json.Unmarshal(rb, orderPlaced)

	fmt.Printf("Success! Order #%v will be ready to collect at %v, enjoy!\n", orderPlaced.OrderID, o.body.RequiredTime)

	return nil
}

func (o *tossedOrder) String() string {
	display := ""

	for _, item := range o.body.OrderItems {
		display += fmt.Sprintf("%v x %v @ £%v\n", item.Quantity, item.ItemName, item.Price)
		for _, modifier := range item.ModifierOptions {
			display += fmt.Sprintf("\t -- %v\n", modifier.Name)
		}
	}
	display += "\n\n"
	display += fmt.Sprintf("Total (before discounts):\t£%v\n", o.body.OrderTotal)
	display += fmt.Sprintf("Discount:\t\t\t£%v", o.body.PromoCodeDiscount)

	return display
}
