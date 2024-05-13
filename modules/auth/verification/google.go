package verification

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/qnify/api-server/utils/chttp"
	"github.com/qnify/api-server/utils/consts"
)

type GoogleAuthConfig struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	RedirectURI  string `yaml:"redirect_uri"`
}

type GoogleUserResult struct {
	Id            string `json:"id,omitempty"`
	Email         string `json:"email,omitempty"`
	VerifiedEmail bool   `json:"email_verified,omitempty"`
	Name          string `json:"name,omitempty"`
	GivenName     string `json:"given_name,omitempty"`
	FamilyName    string `json:"family_name,omitempty"`
	Picture       string `json:"picture,omitempty"`
	Locale        string `json:"locale,omitempty"`
}

var gApiClient = chttp.NewClient("https://www.googleapis.com", true)

func GetGoogleUser(access_token string) (*GoogleUserResult, error) {

	res, resBody, err := gApiClient.Get(fmt.Sprintf("/oauth2/v3/userinfo?access_token=%s", access_token), nil)

	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve user")
	}

	var userBody GoogleUserResult
	if err := json.Unmarshal(resBody, &userBody); err != nil {
		return nil, err
	}
	return &userBody, nil
}

type GoogleOauthToken struct {
	AccessToken string `json:"access_token,omitempty"`
	IdToken     string `json:"id_token,omitempty"`
}

var oAuthClient = chttp.NewClient("https://oauth2.googleapis.com", true)

func GetGoogleOauthToken(code string, config GoogleAuthConfig) (*GoogleOauthToken, error) {
	values := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"client_id":     {config.ClientID},
		"client_secret": {config.ClientSecret},
		"redirect_uri":  {config.RedirectURI},
	}
	headers := map[string]string{consts.ContentType: consts.UrlEncoded}
	res, resBody, err := oAuthClient.Post("/token", headers, bytes.NewBufferString(values.Encode()))

	if err != nil {
		return nil, errors.New("error while making request")
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("invalid token")
	}

	var oAuthToken GoogleOauthToken
	if err := json.Unmarshal(resBody, &oAuthToken); err != nil {
		return nil, err
	}

	return &oAuthToken, nil
}
