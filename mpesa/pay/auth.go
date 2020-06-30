package pay

import (
	"encoding/json"
	"fmt"
	"mpesa/mpesa/model"
	"net/http"
)

func (s Service) auth() (string, error) {
	url := s.baseURL() + "oauth/v1/generate?grant_type=client_credentials"
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return "", err
	}

	req.SetBasicAuth(s.AppKey, s.AppSecret)
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("could not send auth request: %v", err)
	}

	if res != nil {
		defer res.Body.Close()
	}

	var authResp model.AccessToken

	err = json.NewDecoder(res.Body).Decode(&authResp)
	if err != nil {
		return "", fmt.Errorf("could not decode auth response: %v", err)
	}

	accessToken := authResp.AccessToken
	return accessToken, nil
}
