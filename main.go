package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"gopkg.in/telegram-bot-api.v4"
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

	telegramToken, err := ioutil.ReadFile("./secrets/telegram-token")
	if err != nil {
		log.Fatal("Error while loading telegram token.\n", err)
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

	bot, err := tgbotapi.NewBotAPI(string(telegramToken))
	if err != nil {
		log.Fatal("Error while creating new Telegram bot client. ", err)
	}

	log.Print(bot.Self.UserName)

	update := tgbotapi.UpdateConfig{}
	update.Timeout = 60
	update.Offset = 295741457

	updates, err := bot.GetUpdates(update)
	if err != nil {
		log.Fatal("Error while fetching last updates. ", err)
	}

	for _, u := range updates {
		if u.Message == nil {
			continue
		}
		log.Printf("[%s] Update ID: %d", u.Message.Text, u.UpdateID)
		if err = popular(bot, &u, u.Message.Text); err != nil {
			log.Print(err)
		}
	}
}

func help(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "/popular word : get last popular tweets about `word`")
	_, err := bot.Send(msg)
	return err
}

func popular(bot *tgbotapi.BotAPI, update *tgbotapi.Update, text string) error {
	if len(strings.Split(text, " ")) != 2 {
		return help(bot, update)
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Popular: %s", strings.Split(text, " ")[1]))
	_, err := bot.Send(msg)
	return err
}
