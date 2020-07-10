package model

//AccessToken object received from successfully performing client credential grant type
type AuthResponse struct {
	AccessToken string `json:"access_token"`
}
