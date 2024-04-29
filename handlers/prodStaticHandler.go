//go:build !dev
// +build !dev

package handlers

import (
	"embed"
	"net/http"
)

//go:generate cp -r ../public ./public
//go:embed public/*
var publicFS embed.FS

func Public() http.Handler {
	return http.FileServer(http.FS(publicFS))
}
