package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

// Response struct holds the url
// of json response
type Response struct {
	URL string `json:"url"`
}

func ImagesWaifu(nsfwURL string, tags string) string {
	geturl := getImageURL("https://waifu.pics/api/" + nsfwURL + "/" + tags)
	return geturl
}

func getImageURL(waifuURL string) string {

	resp, err := http.Get(waifuURL)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	return result.URL
}

func RandomTags() string {
	x := []string{"waifu", "neko", "shinobu", "bully",
		"cry", "hug", "kiss", "lick", "pat", "smug",
		"highfive", "nom", "bite", "slap", "wink", "poke",
		"dance", "cringe", "blush"}

	rand.Seed(time.Now().Unix())
	num := rand.Int()
	z := num % len(x)
	return x[z]
}
