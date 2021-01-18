package oauth2config

import (
	"github.com/higashi000/spovimfy/spotifyapi"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

var Config oauth2.Config

func InitOauth2Config() {
	config := oauth2.Config{
		ClientID:     spotifyapi.ClientID,
		ClientSecret: spotifyapi.ClientSecret,
		Endpoint:     spotify.Endpoint,
		RedirectURL:  "https://localhost:5000/callback/",
		Scopes: []string{
			"user-read-email",
			"user-read-private",
			"user-modify-playback-state",
			"playlist-read-private",
			"playlist-read-collaborative",
		},
	}

	Config = config
}
