package main

import (
	"github.com/higashi000/spovimfy/oauth2config"
	"github.com/higashi000/spovimfy/router"
	"github.com/higashi000/spovimfy/spotifyapi"
)

func main() {
	spotifyapi.InitSpotifyAPI()
	oauth2config.InitOauth2Config()

	e := router.NewRouter()

	e.Logger.Fatal(e.Start(":5000"))
}
