package user

type Profile struct {
	Address  string `json:"Address"`
	Email    string `json:"Email"`
	Name     string `json:"Name"`
	PostCode string `json:"Postcode"`
	UserID   string `json:"UserID"`
}
