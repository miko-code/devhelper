package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type Oauth struct {
	Oconf *oauth2.Config
}

func NewOauthConfig(config *Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  config.RedirectURL,
		Scopes:       config.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://stackoverflow.com/oauth",
			TokenURL: "https://stackoverflow.com/oauth/access_token",
		},
	}
}

func NewOauth(oconf *oauth2.Config) *Oauth {
	return &Oauth{Oconf: oconf}
}

var stackOauthConfig = &oauth2.Config{

	RedirectURL: "http://localhost:8989/callback",

	ClientID: "19181",

	ClientSecret: "VB1EP)vujvApAHpZCamO*w((",

	Scopes: []string{"no_expiry"},

	Endpoint: oauth2.Endpoint{

		AuthURL: "https://stackoverflow.com/oauth",

		TokenURL: "https://stackoverflow.com/oauth/access_token",
	},
}

func (o *Oauth) Login(c *gin.Context) {
	state := generateStateOauthCookie(c)
	url := o.Oconf.AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, url)

}

func (o *Oauth) Callback(c *gin.Context) {
	ctx := context.Background()
	oauthState, _ := c.Cookie("oauthstate")
	code := c.Query("code")
	if c.Query("state") != oauthState || len(code) == 0 {
		log.Println("invalid oauth stackoverflow state or code")
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	tkn, err := o.Oconf.Exchange(ctx, code)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"token": tkn,
	})

}
func generateStateOauthCookie(c *gin.Context) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour).Nanosecond()

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	c.SetCookie("oauthstate", state, expiration, "/", "localhost", true, true)

	return state
}
