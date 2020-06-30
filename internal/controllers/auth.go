package controllers

import (
	"fmt"
	"github.com/bluele/gcache"
	"log"
	"mpesa/internal/http"
	"mpesa/internal/model"
	"mpesa/internal/utils"
	"strconv"
	"time"
)

//generateToken generates an AccessToken
func generateToken(key, secret, baseURL string, cache gcache.Cache) string {

	endpoint := oauthGenerate

	var headers = []model.Header{
		model.Header{
			Key:   "Authorization",
			Value: fmt.Sprintf("Basic %s", utils.EncodeConsumerKey(key, secret)),
		},
	}

	params := &model.OAuthString{
		GrantType: "client_credentials",
	}

	res := http.Get(endpoint, baseURL, headers, params)

	var token model.AccessToken

	utils.Unmarshal(res, &token)
	cacheToken(token, cache)
	return token.AccessToken

}

//getToken either retrives token from Cache or Get a new one
func getToken(key, secret, baseURL string, cache gcache.Cache) string {
	token, err := fetchToken(cache)
	if err != nil {
		return generateToken(key, secret, baseURL, cache)
	}
	return fmt.Sprintf("%s", token)
}

//cacheToken caches token with the expiry time received from API
func cacheToken(token model.AccessToken, cache gcache.Cache) {
	expiry, err := strconv.Atoi(token.ExpiresIn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error converting expiry to string, error:%s", err))
	}
	cache.SetWithExpire("Token", token.AccessToken, time.Duration(expiry-100)*time.Second)
}

//fetches stored token if not expired
func fetchToken(cache gcache.Cache) (interface{}, error) {
	value, err := cache.Get("Token")
	return value, err
}

