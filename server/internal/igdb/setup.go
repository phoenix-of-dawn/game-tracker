package igdb

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type apiKeys struct {
	Client_id 	string
	Client_secret string
}

type accessToken struct {
	Access_token string `json:"access_token"`
	Expires_in int  `json:"expires_in"`
	Token_type string `json:"token_type"`
}

var ApiKeys apiKeys
var AccessToken accessToken

func loadEnvVariables() {
	ApiKeys.Client_id = os.Getenv("CLIENT_ID")
	ApiKeys.Client_secret = os.Getenv("CLIENT_SECRET")
}

func loadAccessToken() {
	url := fmt.Sprintf("https://id.twitch.tv/oauth2/token?client_id=%s&client_secret=%s&grant_type=client_credentials", 
		ApiKeys.Client_id, 
		ApiKeys.Client_secret)
	resp, err := http.Post(url, "application/json", nil)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	byte_value, _ := io.ReadAll(resp.Body)

	json.Unmarshal(byte_value, &AccessToken)
}

func Setup() {
	loadEnvVariables()
	loadAccessToken()
}

