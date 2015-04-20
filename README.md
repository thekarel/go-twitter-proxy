# Hank Morgan

A Twitter API 1.1 proxy *example/starting point* written in Go that exposes
[GET statuses/user_timeline](https://dev.twitter.com/rest/reference/get/statuses/user_timeline) on a selected port.

Once customized, can be used to serve clients (eg. mobile apps) with JSON feed
from the Twitter API.

Uses [ChimeraCoder/anaconda](https://github.com/ChimeraCoder/anaconda) for API access.

To run the proxy you will need credentials from [apps.twitter.com](https://apps.twitter.com/).

## Run

```
export TWITTER_CONSUMERKEY="..."
export TWITTER_CONSUMERSECRET="..."
export TWITTER_ACCESSTOKEN="..."
export TWITTER_ACCESSTOKENSECRET="..."
export ADDR=":7179"

go run main.go
```

then visit `http://localhost:7179/GetUserTimeline?screen_name=golang`

The `ADDR` env var is used for setting the listen address in `http.ListenAndServe`.

## Docker

The repo includes a trivial Dockerfile so it can be deployed with ease (once
again, this is just an example, not available via Docker Registry).

Build the image with

```
docker build -t "hankmorgan" .
```

Run the image and expose the server on port `7179`:

```
docker run \
--rm \
--env TWITTER_CONSUMERKEY="..." \
--env TWITTER_CONSUMERSECRET="..." \
--env TWITTER_ACCESSTOKEN="..." \
--env TWITTER_ACCESSTOKENSECRET="..." \
--env ADDR=":7179" \
-p 7179:7179 \
hankmorgan
```

On OS X or Windows, check the IP address of docker using `boot2docker ip`.

## By

Charles Szilagyi [charlesagile.com](http://charlesagile.com) [@karolysz](https://twitter.com/karolysz)