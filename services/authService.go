package services

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/dimitar-ivanov-93/htmx/types"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthService struct {
	OAuthConfig *oauth2.Config
}

func NewAuthService() *AuthService {
	return &AuthService{
		OAuthConfig: &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
			Endpoint:     google.Endpoint,
		},
	}
}

func (s *AuthService) GetUserdata(state, code string) (*types.UserData, error) {
	if state != "random-string" {
		return nil, errors.New("invalid user state")
	}

	token, err := s.OAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var userData types.UserData
	if err := json.Unmarshal(data, &userData); err != nil {
		return nil, err
	}

	return &userData, nil
}
