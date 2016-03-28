# TweetMap

Add Twitter API keys in `run/collect.go`:
* `TWITTER_CONSUMER_KEY`
* `TWITTER_CONSUMER_SECRET`
*	`TWITTER_ACCESS_TOKEN`
*	`TWITTER_ACCESS_TOKEN_SECRET`

To collect Tweets, run:
`go run github.com/manuelmenzella/TweetMap/main.go -collect`

To generate the map, run:
`go run github.com/manuelmenzella/TweetMap/main.go -collect`

Map coordinate range can be changed by editing `run/draw.go`.
