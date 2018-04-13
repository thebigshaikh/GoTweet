package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/ChimeraCoder/anaconda"
)

func main() {
	consumerKey := "yVUaVdBYnsAzIqkz3Yt7skHIY"
	consumerSecret := "ELGvyaVaxiamrMLgtilB0tqCcw1oyxAvBX6a75F7lWXP8X7G0j"
	accessToken := "744602175540592640-nKlNJ2CBh5J0knOZTd90OzKDDeV1n69"
	accessSecret := "6yctXw0WBUiaDKlu1hspuYybkldLBqCoyoEkJMlOmdoq2"

	//Anaconda APi
	
	anaconda.SetConsumerKey(consumerkey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessSecret)

	searchResult, _ := api.GetSearch("golang", nil)
	for _ , tweet := range searchResult.Statuses {
    	fmt.Print(tweet.Text)
	}


}
