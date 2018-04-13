package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	//consumerKey := "KEY"
	//consumerSecret := "SECRET"
	//accessToken := "TOKEN"
	//accessSecret := "SECRET"

	config := oauth1.NewConfig("yVUaVdBYnsAzIqkz3Yt7skHIY","ELGvyaVaxiamrMLgtilB0tqCcw1oyxAvBX6a75F7lWXP8X7G0j")
	token := oauth1.NewToken("744602175540592640-nKlNJ2CBh5J0knOZTd90OzKDDeV1n69","6yctXw0WBUiaDKlu1hspuYybkldLBqCoyoEkJMlOmdoq2")

	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		IncludeEmail: twitter.Bool(true),
	}
	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	fmt.Printf("User's Name:%+v\n", user.ScreenName)
}
