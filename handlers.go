package main

import (
	"fmt"
	"net/http"
)

func (app *application) user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World! User")
}
