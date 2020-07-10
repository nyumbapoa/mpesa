package model

// Identity Request Object
type Identity struct {
	Initiator         string `json:"Initiator"`
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	TransactionType   string `json:"TransactionType"`
	PhoneNumber       int    `json:"PhoneNumber"`
	CallBackURL       string `json:"CallBackURL"`
	TransactionDesc   string `json:"TransactionDesc"`
}
