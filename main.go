package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	var hashtag string

	hashtag = "TheVoice"
	consumerKey, err := ioutil.ReadFile("./secrets/consumer-key")
	if err != nil {
		log.Fatal("Error while loading consumer key.\n", err)
	}

	consumerSecret, err := ioutil.ReadFile("./secrets/consumer-secret")
	if err != nil {
		log.Fatal("Error while loading consumer secret.\n", err)
	}

	accessToken, err := ioutil.ReadFile("./secrets/access-token")
	if err != nil {
		log.Fatal("Error while loading access token.\n", err)
	}

	accessTokenSecret, err := ioutil.ReadFile("./secrets/access-token-secret")
	if err != nil {
		log.Fatal("Error while loading access token secret.\n", err)
	}

	config := oauth1.NewConfig(string(consumerKey), string(consumerSecret))
	token := oauth1.NewToken(string(accessToken), string(accessTokenSecret))

	httpClient := oauth1.NewClient(oauth1.NoContext, config, token)

	client := twitter.NewClient(httpClient)

	hashtagSearch := &twitter.SearchTweetParams{
		Query:      fmt.Sprintf("#%s", hashtag),
		Count:      5,
		ResultType: "popular",
		Lang:       "en",
	}

	tweets, _, err := client.Search.Tweets(hashtagSearch)
	if err != nil {
		log.Fatal(err)
	}

	for _, tweet := range tweets.Statuses {
		log.Print(tweet.Text, tweet.Entities.Urls)
	}
}
