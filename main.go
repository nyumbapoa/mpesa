package main

import (
	"fmt"
	"github.com/nyumbapoa/mpesa/mpesa/pay"
	"log"
)

const (
	appKey    = ""
	appSecret = ""
)

func main() {
	fmt.Println("Welcome to the world of mpesa")

	svc, err := pay.New(appKey, appSecret, pay.SANDBOX)

	if err != nil {
		panic(err)
	}



	res, err := svc.C2BRegisterURL(
		600157,
		"Completed",
		"https://www.pesaswap.com/callback.php",
		"https://www.pesaswap.com/result.php",
	)

	if err != nil {
		panic(err)
	}

	log.Println(res)

}
