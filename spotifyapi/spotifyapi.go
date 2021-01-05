package spotifyapi

import "os"

var (
	ClientID     string
	ClientSecret string
)

func InitSpotifyAPI() {

	clientID := os.Getenv("SPOTIFY_CLIENTID")
	clientSecret := os.Getenv("SPOTIFY_CLIENTSECRET")

	ClientID = clientID
	ClientSecret = clientSecret
}
