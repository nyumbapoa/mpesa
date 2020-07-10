package pay

import (
	"encoding/json"
	"github.com/nyumbapoa/mpesa/mpesa/model"
)

func (s Service) B2B() (string, error) {
	b2b := &model.B2B{
		CommandID:              "",
		Amount:                 0,
		PartyA:                 0,
		SenderIdentifierType:   0,
		PartyB:                 0,
		RecieverIdentifierType: 0,
		Remarks:                "",
		Initiator:              "",
		SecurityCredential:     "",
		QueueTimeOutURL:        "",
		ResultURL:              "",
		AccountReference:       "",
	}
	body, err := json.Marshal(b2b)

	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["content-type"] = "application/json"
	//headers["Authorization"] = "Bearer " + auth
	headers["cache-control"] = "no-cache"

	url := s.baseURL() + processB2B
	return s.newReq(url, body, headers)
}
