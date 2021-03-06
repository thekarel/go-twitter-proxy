// An example Twitter API 1.1 proxy
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

// tweet is the struct for building JSON output for GetUserTimeline
type tweet struct {
	ScreenName string `json:"screen_name"`
	Text       string `json:"text"`
	CreatedAt  string `json:"created_at"`
}

// main configures and starts the server
// This example exposes 1 API method on `/GetUserTimeline` (GET),
// eg. http://192.168.59.103:7179/GetUserTimeline?screen_name=golang
func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":7179"
	}

	http.Handle("/GetUserTimeline", newServer())
	log.Println("Starting server on " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

type server struct {
	api *anaconda.TwitterApi
}

// newServer picks up configuration from ENV, creates TwitterApi and returns a new server
func newServer() *server {
	consumerKey := os.Getenv("TWITTER_CONSUMERKEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMERSECRET")
	accessToken := os.Getenv("TWITTER_ACCESSTOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESSTOKENSECRET")

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessTokenSecret == "" {
		log.Fatal("You need to pass consumerKey, consumerSecret, accessToken, accessTokenSecret but some are missing. Try --help.")
	}

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	return &server{api: api}
}

// ServeHTTP handles incoming queries.
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	screen_name := query.Get("screen_name")

	if screen_name == "" {
		http.Error(w, "screen_name not set, use /GetUserTimeline?screen_name=...", http.StatusBadRequest)
		return
	}

	v := url.Values{}
	v.Set("screen_name", screen_name)

	tweets, err := s.api.GetUserTimeline(v)
	if err != nil {
		log.Print(err)
		http.Error(w, "GetUserTimeline failed", http.StatusInternalServerError)
		return
	}

	tweetList := make([]tweet, 0)

	for _, t := range tweets {
		tweetList = append(tweetList, tweet{t.User.ScreenName, t.Text, t.CreatedAt})
	}

	out, err := json.Marshal(tweetList)
	if err != nil {
		log.Print(err)
		http.Error(w, "JSON marshal failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	fmt.Fprint(w, string(out))
}
