package controllers
/*
import (
	"fmt"
	"github.com/bluele/gcache"
	"mpesa/model"
	"mpesa/utils"
	"net/http"
)

//RegisterC2BURL registers URLs for C2B
func RegisterC2BURL(onfail string, key string, secret string, baseURL string, cache gcache.Cache) model.GenericResponse {
	endpoint := registerC2BURL

	pushObject := &model.C2BURLRegister{
		ShortCode:       globals.ShortCode,
		ResponseType:    onfail,
		ConfirmationURL: globals.ConfirmationURL,
		ValidationURL:   globals.ValidationURL,
	}

	var headers = []model.Header{
		{
			Key:   "authorization",
			Value: fmt.Sprintf("Bearer %s", getToken(key, secret, baseURL, cache)),
		},
		{
			Key:   "content-type",
			Value: "application/json",
		},
		{
			Key:   "cache-control",
			Value: "no-cache",
		},
	}

	res := http.Post(endpoint, headers, pushObject)

	var balanceResponse model.GenericResponse

	utils.Unmarshal(res, &balanceResponse)

	return balanceResponse

}

//SimulateC2B simulates C2B
func SimulateC2B(amount int, phone int, ref string, command string, key string, secret string, baseURL string, cache gcache.Cache) model.GenericResponse {
	endpoint := simulateC2B

	push := &model.C2BSimulation{
		ShortCode:     globals.ShortCode,
		Amount:        amount,
		MSISDN:        phone,
		BillRefNumber: ref,
		CommandID:     command,
	}
	var headers = []model.Header{
		{
			Key:   "authorization",
			Value: fmt.Sprintf("Bearer %s", getToken(key, secret, baseURL, cache)),
		},
		{
			Key:   "content-type",
			Value: "application/json",
		},
		{
			Key:   "cache-control",
			Value: "no-cache",
		},
	}

	res := http.Post(endpoint, headers, push)

	var simulationResponse model.GenericResponse

	utils.Unmarshal(res, &simulationResponse)

	return simulationResponse

}
*/