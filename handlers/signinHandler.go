package handlers

import (
	"net/http"

	"github.com/dimitar-ivanov-93/htmx/views/auth"
)

func HandleSigninIndex(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, auth.Signin())
}
