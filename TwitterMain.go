package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	//"bufio"
	//"os"
)

	

func main() {
	var 
	(
	 consumerKey string   // = "yVUaVdBYnsAzIqkz3Yt7skHIY"
	 consumerSecret string  = "ELGvyaVaxiamrMLgtilB0tqCcw1oyxAvBX6a75F7lWXP8X7G0j"
	  accessToken string     = "744602175540592640-nKlNJ2CBh5J0knOZTd90OzKDDeV1n69"
	 accessSecret string    = "6yctXw0WBUiaDKlu1hspuYybkldLBqCoyoEkJMlOmdoq2"
	)


	fmt.Println(" Enter your consumerKey")
	//reader := bufio.NewReader(os.Stdin)
	//consumerKey, _ = reader.ReadString('\n')
	fmt.Scanf("%s", &consumerKey)


	config := oauth1.NewConfig(consumerKey,consumerSecret)
	token := oauth1.NewToken(accessToken,accessSecret)

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

	//Sending a Tweet
	_,_,_ = client.Statuses.Update("Learning twitter API using GOlang #GOlang", nil)
	 

	// Getting 20 tweets from wall
	timeline,_,_ := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
    Count: 20,
	})
	fmt.Println(timeline)
	
	//status

	tweet, _, _ := client.Statuses.Show(585613041028431872, nil)
	fmt.Println("******* STATUS IS ******* \n", tweet)

	 // var age int
  //   fmt.Println("What is your age?")
  //   _, err: fmt.Scan(&age)

  //   //reading a string
  //   reader := bufio.newReader(os.Stdin)
  //   var name string
  //   fmt.Println("What is your name?")
  //   name, _ := reader.readString("\n")

  //   fmt.Println("Your name is ", name, " and you are age ", age)

}
