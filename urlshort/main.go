package main

import (
	"fmt"
	"net/http"
	"urlshort/shortner"
)

func main() {

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := shortner.MapHandler(pathsToUrls, hello)

	http.ListenAndServe(":8080", mapHandler)
}


func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
