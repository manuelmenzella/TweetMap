package twitter

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/manuelmenzella/TweetMap/tools"
	"net/url"
)

type Auth struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func MakeStream(auth Auth) <-chan anaconda.Tweet {
	return MakeStreamRegion(auth, tools.Region{-180, 180, -90, 90})
}

func MakeStreamRegion(auth Auth, region tools.Region) <-chan anaconda.Tweet {
	anaconda.SetConsumerKey(auth.ConsumerKey)
	anaconda.SetConsumerSecret(auth.ConsumerSecret)
	api := anaconda.NewTwitterApi(auth.AccessToken, auth.AccessTokenSecret)

	locationsString := fmt.Sprintf("%f,%f,%f,%f",
		region.LonMin, region.LatMin, region.LonMax, region.LatMax)

	filterParams := url.Values{}
	filterParams.Set("locations", locationsString)

	stream := api.PublicStreamFilter(filterParams)

	ch := make(chan anaconda.Tweet)
	go func() {
		for data := range stream.C {
			switch tweet := data.(type) {
			case anaconda.Tweet:
				ch <- tweet
			default:
			}
		}
	}()
	return ch
}
