package routes

import (
	"fmt"
	"net/http"
	"solid/src"
)

func Routes() {
	r := src.NewRouter()

	r.Get("/", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
