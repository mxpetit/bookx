package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mxpetit/bookx/model"
	"golang.org/x/oauth2"
	people "google.golang.org/api/people/v1"
	"net/http"
	"os"
)

// oAuth2Code represents an oAuth2 code for offline access, e.g
// 4/yU4cQZTMnnMtetyFcIWNItG32eKxxxgXXX-Z4yyJJJo.4qHskT-UtugceFc0ZRONyF4z7U4UmAI
type oAuth2Code struct {
	Code string `json:"code"`
}

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://ui.book.xyz",
		ClientID:     os.Getenv("GOOGLE_KEY"),
		ClientSecret: os.Getenv("GOOGLE_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	}

	ErrOAuthCodeMissing      = model.NewDatastoreError(http.StatusBadRequest, "oauth_code_missing")
	ErrUnableExchangeCode    = model.NewDatastoreError(http.StatusInternalServerError, "unable_exchange_code")
	ErrUnableGetGoogleClient = model.NewDatastoreError(http.StatusInternalServerError, "unable_get_google_client")
	ErrPersonNotFound        = model.NewDatastoreError(http.StatusNotFound, "person_not_found")
	ErrGivenNameNotFound     = model.NewDatastoreError(http.StatusNotFound, "given_name_not_found")
)

// Authenticate authenticates an user by exchanging his code with
// an access token.
func Authenticate(c *gin.Context) {
	oAuth2Code := &oAuth2Code{}
	err := c.BindJSON(oAuth2Code)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrOAuthCodeMissing.Error(),
		})

		return
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, oAuth2Code.Code)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": ErrUnableExchangeCode.Error(),
		})

		return
	}

	client := googleOauthConfig.Client(oauth2.NoContext, token)
	service, err := people.New(client)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": ErrUnableGetGoogleClient.Error(),
		})

		return
	}

	person, err := service.People.Get("people/me").RequestMaskIncludeField("person.names").Do()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": ErrPersonNotFound.Error(),
		})

		return
	}

	if len(person.Names) == 0 || person.Names[0].GivenName == "" {
		c.JSON(http.StatusOK, gin.H{
			"error": ErrGivenNameNotFound.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"given_name":   person.Names[0].GivenName,
		"access_token": token.AccessToken,
		"expires":      token.Expiry,
		"token_type":   token.TokenType,
	})
}
