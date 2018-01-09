package main

import (
	"fmt"
	"os"

	"github.com/tehcyx/twitter-bot-go-markov-chain/markov"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	consumerKey := os.Getenv("MARKOV_TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("MARKOV_TWITTER_CONSUMER_SECRET")
	accessTokenKey := os.Getenv("MARKOV_TWITTER_ACCESS_TOKEN_KEY")
	accessTokenSecret := os.Getenv("MARKOV_TWITTER_ACCESS_TOKEN_SECRET")

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessTokenKey, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	// Home Timeline
	client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 20,
	})

	// Send a Tweet
	//tweet, resp, err := client.Statuses.Update("just setting up my twttr", nil)

	dict := markov.Train("hallo dies ist ein neuer text, was geht denn so ab hier? hallo peter, schoen dich zu sehen, wie geht es dir?", 10000)
	fmt.Println(dict)
}
