package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"github.com/itspage/feedme/order"
)

type Preferences struct {
	SiteID       string `json:"SiteID"`
	TelNo        string `json:"TelNo"`
	UserName     string `json:"UserName"`
	Password     string `json:"Password"`
	RequiredDate string `json:"RequiredDate"`
	OrderItems   []*order.Item `json:"OrderItems"`
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
	PromoCode            string `json:"PromoCode"`
	DeliveryInstructions string `json:"DeliveryInstructions"`
	DeliveryAddress      string `json:"DeliveryAddress"`
	DeliveryPostcode     string `json:"DeliveryPostcode"`
}

func LoadPreferences() (*Preferences, error) {
	home := os.Getenv("HOME")
	path := fmt.Sprintf("%v/.feedme", home)
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	prefs := &Preferences{}
	return prefs, json.Unmarshal(f, prefs)
}
