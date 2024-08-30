package main

import "net/http"

func (app *application) ValidUser(next http.Handler) http.Handler {
	// next
	return next
}
