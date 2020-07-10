package pay

import (
	"encoding/json"
	"github.com/nyumbapoa/mpesa/mpesa/model"
)

func (s Service) B2C() (string, error) {
	b2cObject := &model.B2C{
		InitiatorName:      "",
		SecurityCredential: "",
		CommandID:          "",
		Amount:             0,
		PartyA:             0,
		PartyB:             0,
		Remarks:            "",
		QueueTimeOutURL:    "",
		ResultURL:          "",
		Occassion:          "",
	}

	body, err := json.Marshal(b2cObject)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := s.baseURL() + registerC2BURL
	return s.newReq(url, body, headers)
}
