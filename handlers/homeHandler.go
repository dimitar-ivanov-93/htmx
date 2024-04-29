package handlers

import (
	"net/http"

	"github.com/dimitar-ivanov-93/htmx/views/home"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index())
}
