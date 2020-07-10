package model

//C2BURLRegister used to create validation and confirmation URLS
type C2BURLRegister struct {
	ShortCode       int    `json:"ShortCode"`
	ResponseType    string `json:"ResponseType"`
	ConfirmationURL string `json:"ConfirmationURL"`
	ValidationURL   string `json:"ValidationURL"`
}

//C2BSimulation used for simulating transactions
type C2BSimulation struct {
	ShortCode     int    `json:"ShortCode"`
	CommandID     string `json:"CommandID"`
	Amount        int    `json:"Amount"`
	MSISDN        int    `json:"Msisdn"`
	BillRefNumber string `json:"BillRefNumber"`
}
