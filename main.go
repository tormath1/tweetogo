package main

import (
	"io/ioutil"
	"log"
)

func main() {
	consumerKey, err := ioutil.ReadFile("./consumer.key")
	if err != nil {
		log.Fatal("Error while loading consumer key.\n", err)
	}

	consumerSecret, err := ioutil.ReadFile("./consumer.secret")
	if err != nil {
		log.Fatal("Error while loading consumer secret.\n", err)
	}

	log.Print(string(consumerKey))
	log.Print(string(consumerSecret))
}
