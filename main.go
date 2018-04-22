package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	telegramToken := os.Getenv("TELEGRAM_TOKEN")

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)

	httpClient := oauth1.NewClient(oauth1.NoContext, config, token)

	client := twitter.NewClient(httpClient)

	bot, err := tgbotapi.NewBotAPI(string(telegramToken))
	if err != nil {
		log.Fatal("Error while creating new Telegram bot client. ", err)
	}

	update := tgbotapi.UpdateConfig{}
	update.Timeout = 60
	update.Limit = 3

	for {
		updates, err := bot.GetUpdates(update)
		if err != nil {
			log.Fatal("Error while fetching last updates. ", err)
		}
		for _, u := range updates {
			if u.Message == nil {
				continue
			}
			if err = popular(bot, &u, u.Message.Text, client); err != nil {
				log.Print(err)
			}
			update.Offset = u.UpdateID + 1
		}
	}
}

func help(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "/popular word : get last popular tweets about `word`")
	_, err := bot.Send(msg)
	return err
}

func popular(bot *tgbotapi.BotAPI, update *tgbotapi.Update, text string, client *twitter.Client) error {
	if len(strings.Split(text, " ")) != 2 {
		return help(bot, update)
	}
	tweets, err := getPopularTweets(client, strings.Split(text, " ")[1])
	if err != nil {
		return err
	}
	for _, tweet := range tweets {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("@%s: %s", tweet.User.ScreenName, tweet.Text))
		_, err = bot.Send(msg)
		if err != nil {
			return err
		}
	}
	return err
}

func getPopularTweets(client *twitter.Client, text string) ([]twitter.Tweet, error) {
	hashtagSearch := &twitter.SearchTweetParams{
		Query:      fmt.Sprintf("#%s", text),
		Count:      5,
		ResultType: "popular",
		Lang:       "en",
	}
	tweets, _, err := client.Search.Tweets(hashtagSearch)
	if err != nil {
		return nil, err
	}
	return tweets.Statuses, nil
}
