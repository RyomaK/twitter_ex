package twitter

import (
	"log"

	"github.com/dghubble/oauth1"
	twauth "github.com/dghubble/oauth1/twitter"
)

const outOfBand = "oob"

type Approval struct {
	RequestToken string
	URL          string
}

func GetConfig() oauth1.Config {
	t := Token()
	consumerKey := t.ck
	consumerSecret := t.cs
	if consumerKey == "" || consumerSecret == "" {
		log.Fatal("Required environment variable missing.")
	}

	return oauth1.Config{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		CallbackURL:    outOfBand,
		Endpoint:       twauth.AuthorizeEndpoint}
}

func GetRequestToken(config oauth1.Config) (string, string) {
	requestToken, url, err := login(config)
	if err != nil {
		log.Fatalf("Request Token Phase: %s", err.Error())
	}
	return requestToken, url
}

func login(config oauth1.Config) (requestToken string, url string, err error) {
	requestToken, _, err = config.RequestToken()
	if err != nil {
		return "", "", err
	}
	authorizationURL, err := config.AuthorizationURL(requestToken)
	if err != nil {
		return "", "", err
	}
	return requestToken, authorizationURL.String(), err
}

func ReceivePIN(config oauth1.Config, requestToken string, verifier string) (*oauth1.Token, error) {
	accessToken, accessSecret, err := config.AccessToken(requestToken, "secret does not matter", verifier)
	if err != nil {
		return nil, err
	}
	return oauth1.NewToken(accessToken, accessSecret), err
}
