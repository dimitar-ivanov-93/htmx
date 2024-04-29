package handlers

import (
	"log"
	"net/http"

	"github.com/dimitar-ivanov-93/htmx/middleware"
	"github.com/dimitar-ivanov-93/htmx/services"
)

func HandleSignin(w http.ResponseWriter, r *http.Request) error {
	authService := services.NewAuthService()
	url := authService.OAuthConfig.AuthCodeURL("random-string")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return nil
}

func HandleCallback(w http.ResponseWriter, r *http.Request) error {
	authService := services.NewAuthService()
	userData, err := authService.GetUserdata(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		log.Printf("Error getting user data: %v\n", err)
		return err
	}

	session, _ := middleware.Store.Get(r, "session-name")
	session.Values["authenticated"] = true
	session.Values["user"] = userData
	if err := session.Save(r, w); err != nil {
		log.Printf("Error saving session: %v\n", err)
		return err
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}
