package run

import (
	"encoding/json"
	"github.com/manuelmenzella/TweetMap/twitter"
	"log"
	"os"
)

const (
	TWITTER_CONSUMER_KEY        = ""
	TWITTER_CONSUMER_SECRET     = ""
	TWITTER_ACCESS_TOKEN        = ""
	TWITTER_ACCESS_TOKEN_SECRET = ""
)

func Collect(tweetFilePath string, maxTweets int64) {
	tweetFile, err := os.Create(tweetFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer tweetFile.Close()

	tweetEncoder := json.NewEncoder(tweetFile)

	twitterAuth := twitter.Auth{
		TWITTER_CONSUMER_KEY,
		TWITTER_CONSUMER_SECRET,
		TWITTER_ACCESS_TOKEN,
		TWITTER_ACCESS_TOKEN_SECRET}

	var tweetIndex int64 = 0
	for tweet := range twitter.MakeStream(twitterAuth) {
		tweetEncoder.Encode(tweet)

		tweetIndex++
		log.Printf("Collected %d tweets.", tweetIndex)

		if maxTweets > 0 && tweetIndex >= maxTweets {
			break
		}
	}
}
