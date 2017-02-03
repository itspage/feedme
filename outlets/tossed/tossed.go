package tossed

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/itspage/feedme/order"
	"github.com/itspage/feedme/user"
)

const (
	apiKey        = "f856649e-fdb9-491d-bb71-db717b61d312"
	orderEndpoint = "https://api.pointone.co.uk/ooappservice.svc/Order/JSON"
	loginEndpoint = "https://api.pointone.co.uk/ooappservice.svc/UserLogon/JSON"
)

type tossed struct {
}

func Outlet() *tossed {
	return &tossed{}
}

func (t *tossed) NewOrder(prefs *user.Preferences, loggedInUser *user.Profile) (order.Order, error) {
	ob := &orderBody{
		PromoCodeDiscount:    "0.63", // TODO: calculate Absolute discount
		TelNo:                prefs.TelNo,
		UserID:               loggedInUser.UserID,
		Email:                loggedInUser.Email,
		DeliveryCost:         "0",
		Name:                 loggedInUser.Name,
		RequiredDate:         time.Now().Format("02/01/2006"),
		OrderItems:           prefs.OrderItems,
		APIKey:               apiKey,
		Payment:              prefs.Payment,
		OrderTotal:           "6.29", // TODO: calculate
		SiteID:               prefs.SiteID,
		PromoCode:            prefs.PromoCode,
		DeliveryInstructions: prefs.DeliveryInstructions,
		DeliveryAddress:      prefs.DeliveryAddress,
		RequiredTime:         requiredTime(),
		DeliveryPostcode:     prefs.DeliveryPostcode,
		OrderType:            0,
	}
	b, _ := json.Marshal(ob)

	req, _ := http.NewRequest("POST", orderEndpoint, strings.NewReader(string(b)))
	req = addDefaultHeaders(req)

	log.Printf("built request %v", string(b))
	return &tossedOrder{Request: req, body: ob}, nil
}

func (t *tossed) Login(prefs *user.Preferences) (*user.Profile, error) {
	type loginBody struct {
		SiteID   string `json:"SiteID"`
		Password string `json:"Password"`
		APIKey   string `json:"APIKey"`
		UserName string `json:"UserName"`
	}

	lb := &loginBody{
		SiteID:   "dccc6adb-f3fc-4c66-96c7-9c5f08cbea34",
		APIKey:   "f856649e-fdb9-491d-bb71-db717b61d312",
		UserName: prefs.UserName,
		Password: prefs.Password,
	}

	b, _ := json.Marshal(lb)

	req, _ := http.NewRequest("POST", loginEndpoint, strings.NewReader(string(b)))
	req = addDefaultHeaders(req)

	log.Printf("sending request %v", string(b))

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("expected 200 status from login, got %v", rsp.StatusCode)
	}

	profile := &user.Profile{}
	rb, _ := ioutil.ReadAll(rsp.Body)
	return profile, json.Unmarshal(rb, profile)
}

func addDefaultHeaders(req *http.Request) *http.Request {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Tossed%20Online%20Ordering/13 CFNetwork/808.3 Darwin/16.3.0")
	req.Header.Add("Accept-Language", "en-gb")
	return req
}

func requiredTime() string {
	return time.Now().Add(16 * time.Minute).Round(5 * time.Minute).Format("15:04")
}
