//go:build dev
// +build dev

package handlers

import (
	"fmt"
	"net/http"
	"os"
)

func Public() http.Handler {
	fmt.Println("building static files for development")
	return http.StripPrefix("/public/", http.FileServerFS(os.DirFS("public")))
}
