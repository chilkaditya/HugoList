package utils

import (
    "context"
    //"fmt"
    "log"
	"os"

    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "google.golang.org/api/youtube/v3"
)

var (
    oauthConfig *oauth2.Config
    token       *oauth2.Token
)

func InitOAuthConfig() {
    oauthConfig = &oauth2.Config{
        RedirectURL:  os.Getenv("REDIRECT_URL"),
        ClientID:     os.Getenv("CLIENT_ID"),
        ClientSecret: os.Getenv("CLIENT_SECRET"),
        Scopes:       []string{youtube.YoutubeScope},
        Endpoint:     google.Endpoint,
    }
}

func GetAuthURL() string {
    return oauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

func ExchangeCode(code string) (*oauth2.Token, error) {
    ctx := context.Background()
    tok, err := oauthConfig.Exchange(ctx, code)
    if err != nil {
        return nil, err
    }
    return tok, nil
}

func SaveToken(tok *oauth2.Token) {
    token = tok
}

func GetClient() *youtube.Service {
    ctx := context.Background()
    client := oauthConfig.Client(ctx, token)

    service, err := youtube.New(client)
    if err != nil {
        log.Fatalf("Error creating YouTube client: %v", err)
    }
    return service
}
