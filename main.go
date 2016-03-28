package main

import (
	"flag"
	"github.com/manuelmenzella/Image/run"
)

func main() {
	shouldCollect := flag.Bool(
		"collect",
		false,
		"Switch to collection mode.")
	tweetFilePath := flag.String(
		"tweets",
		"tweets.txt",
		"Text file to use for collecting and reading tweets.")
	imageFilePath := flag.String(
		"map",
		"tweets.jpeg",
		"Output path for map image.")
	maxTweets := flag.Int64(
		"max",
		-1,
		"Maximum number of tweets to process.")
	flag.Parse()

	if *shouldCollect {
		run.Collect(*tweetFilePath, *maxTweets)
	} else {
		run.Draw(*tweetFilePath, *imageFilePath, *maxTweets)
	}
}
