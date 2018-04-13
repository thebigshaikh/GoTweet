package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	fmt.Printf(" **** Executing **** ")
	consumerKey := " 01C5LitvQlxCtbZum3SYK3LS5"
	consumerSecret := "nhGTEkMaUgbWqzygrF9JNw2meRbQaBnj6JDgiuZwkii6znnf50"
	accessToken := "744602175540592640-nKlNJ2CBh5J0knOZTd90OzKDDeV1n69"
	accessSecret := " 6yctXw0WBUiaDKlu1hspuYybkldLBqCoyoEkJMlOmdoq2"

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Search tweets to retweet
	searchParams := &twitter.SearchTweetParams{
		Query:      "#manchester",
		Count:      5,
		ResultType: "recent",
		Lang:       "en",
	}

	searchResult, _, _ := client.Search.Tweets(searchParams)

        // Retweet
	for _, tweet := range searchResult.Statuses {
		tweet_id := tweet.ID
		client.Statuses.Retweet(tweet_id, &twitter.StatusRetweetParams{})

		fmt.Printf("RETWEETED: %+v\n", tweet.Text)
	}
}
