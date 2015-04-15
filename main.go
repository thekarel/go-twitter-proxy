package main

import (
	"flag"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"log"
	"net/http"
	"net/url"
)

var (
	consumerKey       = flag.String("ck", "", "Consumer key")
	consumerSecret    = flag.String("cs", "", "Consumer secret")
	accessToken       = flag.String("at", "", "Access token")
	accessTokenSecret = flag.String("as", "", "Access token secret")
)

func main() {
	flag.Parse()
	http.Handle("/GetUserTimeline", newServer())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type server struct {
	api *anaconda.TwitterApi
}

func newServer() *server {
	if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessTokenSecret == "" {
		log.Fatal("You need to pass consumerKey, consumerSecret, accessToken, accessTokenSecret but some are missing. Try --help.")
	}

	anaconda.SetConsumerKey(*consumerKey)
	anaconda.SetConsumerSecret(*consumerSecret)
	api := anaconda.NewTwitterApi(*accessToken, *accessTokenSecret)

	return &server{api: api}
}

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

	for _, tweet := range tweets {
		fmt.Fprintf(w, "\n\n<div>%s</div>", tweet.Text)
	}
}
