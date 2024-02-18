package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// WriteToConsole is a middleware that writes to the console
// func WriteToConsole(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("Hit the page")
// 		next.ServeHTTP(w, r)
// 	})
// }

/*
nosurf is an HTTP package for Go that helps you prevent Cross-Site Request Forgery attacks.
It acts like a middleware and therefore is compatible with basically any Go HTTP application.
*/
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

// sessionLoad is a middleware that loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
