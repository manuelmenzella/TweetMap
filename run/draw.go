package run

import (
	"encoding/json"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/manuelmenzella/TweetMap/mapper"
	"github.com/manuelmenzella/TweetMap/tools"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"log"
	"os"
)

const (
	// WORLD
	MAP_SIZE_X  = 3840
	MAP_SIZE_Y  = 2160
	MAP_LON_MIN = -180
	MAP_LON_MAX = 180
	MAP_LAT_MIN = -75
	MAP_LAT_MAX = 75

	// USA
	// MAP_SIZE_X = 3840
	// MAP_SIZE_Y = 2844
	// MAP_LON_MIN = -130
	// MAP_LON_MAX = -65
	// MAP_LAT_MIN = 10
	// MAP_LAT_MAX = 60

	// EUROPE
	// MAP_SIZE_X  = 3840
	// MAP_SIZE_Y  = 2160
	// MAP_LON_MIN = -20
	// MAP_LON_MAX = 45
	// MAP_LAT_MIN = 30
	// MAP_LAT_MAX = 60

	// SOUTH AMERICA
	// MAP_SIZE_X = 2160
	// MAP_SIZE_Y = 3840
	// MAP_LON_MIN = -88
	// MAP_LON_MAX = -30
	// MAP_LAT_MIN = -60
	// MAP_LAT_MAX = 25
)

func Draw(tweetFilePath, imageFilePath string, maxTweets int64) {
	tweetFile, err := os.Open(tweetFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer tweetFile.Close()

	tweetDecoder := json.NewDecoder(tweetFile)

	var numWithCoor int64 = 0
	var numWithNoCoor int64 = 0

	cBlue := color.RGBA{9, 15, 39, 255}
	cYellow := color.RGBA{255, 230, 137, 80}

	m := mapper.NewMapperRegion(
		image.Point{MAP_SIZE_X, MAP_SIZE_Y},
		tools.Region{MAP_LON_MIN, MAP_LON_MAX, MAP_LAT_MIN, MAP_LAT_MAX})
	m.Clear(cBlue)

	for {
		var tweet anaconda.Tweet
		decodeError := tweetDecoder.Decode(&tweet)
		if decodeError == io.EOF {
			break
		} else if decodeError != nil {
			log.Fatal(err)
		}

		lon, lat, coorErr := getLonLat(&tweet)
		if coorErr != nil {
			numWithNoCoor++
		} else {
			numWithCoor++
			m.DrawCircleCoor(lon, lat, 2.4, cYellow)
		}

		log.Printf(
			"Decoded tweets: %d with coordinates, %d without.",
			numWithCoor, numWithNoCoor)

		if maxTweets > 0 && numWithCoor >= maxTweets {
			break
		}
	}

	imageFile, imageErr := os.Create(imageFilePath)
	if imageErr != nil {
		log.Fatal(imageErr)
	}
	defer imageFile.Close()

	jpeg.Encode(imageFile, m.MapImage, &jpeg.Options{100})
}

func getLonLat(tweet *anaconda.Tweet) (float64, float64, error) {
	lon, lonErr := tweet.Longitude()
	lat, latErr := tweet.Latitude()
	if lonErr == nil && latErr == nil {
		return lon, lat, nil
	}

	boundingBox := tweet.Place.BoundingBox
	allCoor := boundingBox.Coordinates
	if boundingBox.Type != "" && len(allCoor) > 0 {
		coor := allCoor[0]
		sumLon, sumLat, pointCount := float64(0), float64(0), 0
		for _, point := range coor {
			if len(point) == 2 {
				sumLon += point[0]
				sumLat += point[1]
				pointCount++
			}
		}
		if pointCount > 0 {
			lon, lat = sumLon/float64(pointCount), sumLat/float64(pointCount)
			return lon, lat, nil
		}
	}

	return 0, 0, fmt.Errorf("No coordinates in this tweet.")
}
