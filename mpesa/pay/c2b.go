package pay

import (
	"encoding/json"
	"github.com/nyumbapoa/mpesa/mpesa/model"
)

// C2BRegisterURL requests
func (s Service) C2BRegisterURL(shortcode int, response string, confUrl string, valUrl string) (string, error) {

	c2bRegisterObject := &model.C2BURLRegister{
		ShortCode:       shortcode,
		ResponseType:    response,
		ConfirmationURL: confUrl,
		ValidationURL:   valUrl,
	}

	body, err := json.Marshal(c2bRegisterObject)

	if err != nil {
		return "", err
	}

	auth, err := s.auth()
	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + auth
	headers["Cache-Control"] = "no-cache"

	url := s.baseURL() + "mpesa/c2b/v1/registerurl"
	return s.newReq(url, body, headers)
}

// C2BSimulation sends a new request
func (s Service) C2BSimulation(shortcode int, command string, amount int, phone int, billref string) (string, error) {
	c2bObject := &model.C2BSimulation{
		ShortCode:     shortcode,
		CommandID:     command,
		Amount:        amount,
		MSISDN:        phone,
		BillRefNumber: billref,
	}
	body, err := json.Marshal(c2bObject)
	if err != nil {
		return "", err
	}

	auth, err := s.auth()
	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + auth
	headers["cache-control"] = "no-cache"

	url := s.baseURL() + "mpesa/c2b/v1/simulate"
	return s.newReq(url, body, headers)
}
