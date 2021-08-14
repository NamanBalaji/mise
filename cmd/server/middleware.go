package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf is a middleware that that handles CSRF attacks
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
