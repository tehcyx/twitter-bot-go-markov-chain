package main

import (
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/tehcyx/gomarkov"
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

	dict := markov.TrainFromFolder("training", 10000000)

	// dict := markov.Train("Tourists airlifted from snowbound Swiss ski resort. The Swiss Alpine resort of Zermatt airlifted guests by helicopter on Tuesday after heavy snow and a power cut stranded thousands of visitors.", 10000)

	dict = markov.BulkAdjustFactors(dict, 10000, []markov.FitnessFunc{markov.FitnessFunction})

	message := markov.Generate(dict, 20, "")
	fmt.Println()
	fmt.Println(message)
	fmt.Println()
}
