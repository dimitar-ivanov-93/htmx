package middleware

import (
	"encoding/gob"
	"log"
	"os"

	"github.com/dimitar-ivanov-93/htmx/types"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var Store *sessions.CookieStore

func init() {
	gob.Register(&types.UserData{})

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	secret := os.Getenv("SESSION_SECRET")
	if secret == "" {
		log.Fatal("SESSION_SECRET is not set")
	}

	Store = sessions.NewCookieStore([]byte(secret))
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 24, // 24 hours
		HttpOnly: true,
	}
}
