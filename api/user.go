package handler

import (
	"fmt"
	"net/http"
)

func User(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go! from user Page!!!</h1>")
}
