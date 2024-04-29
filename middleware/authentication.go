package middleware

import (
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := Store.Get(r, "session-name")
		if err != nil {
			log.Printf("Error getting session: %v\n", err)
			http.Redirect(w, r, "/signin", http.StatusTemporaryRedirect)
			return
		}

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/signin", http.StatusTemporaryRedirect)
			return
		}

		next.ServeHTTP(w, r)
	})
}
