package main

import (
	"io/ioutil"
	"log"
)

func main() {
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

	log.Print(string(consumerKey))
	log.Print(string(consumerSecret))
	log.Print(string(accessTokenSecret))
	log.Print(string(accessToken))
}
