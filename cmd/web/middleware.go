package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// func writeToConsole(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("Hit the page...")
// 		next.ServeHTTP(w, r)
// 	})
// }

// Adds CSRF protection to all POST requests 
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction ,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// Makes the web application state aware
// Loads & saves the session uon every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
