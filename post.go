package main

import (
	"log"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	annualIncomeValue := os.Getenv("ANNUAL_INCOME")
	pattern := regexp.MustCompile(`\A年収\z`)
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	log.Println("Client initialized")

	demux := twitter.NewSwitchDemux()
	demux.DM = func(dm *twitter.DirectMessage) {
		if pattern.MatchString(dm.Text) {
			log.Printf("Recieved DM %s from %s", dm.Text, dm.SenderScreenName)
			client.DirectMessages.New(&twitter.DirectMessageNewParams{
				UserID:     dm.SenderID,
				ScreenName: dm.SenderScreenName,
				Text:       annualIncomeValue + "万円",
			})

		}

	}

	log.Println("Starting Stream...")

	params := &twitter.StreamUserParams{
		With:          "followings",
		StallWarnings: twitter.Bool(true),
	}
	stream, _ := client.Streams.User(params)
	go demux.HandleChan(stream.Messages)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	log.Println("Stopping Stream...")
	stream.Stop()

}
