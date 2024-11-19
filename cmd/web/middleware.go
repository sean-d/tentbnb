package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// NoSurf adds CSRF protection to all POST requests.
// returns a CSRF handler that sets a cookie relevant for web traffic only, for the entire site,
// not using https, and some cross-site requests are allowed due to the lax mode.
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad provides state for session data/information.
// Loads and saves session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
