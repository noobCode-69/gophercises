package main

import (
	"cyoa/story"
	"net/http"
)


func main() {
	

	s , err := story.GetStory("gopher.json");

	if err != nil {
		panic(err)
	}

	handler, err := story.GetHandler(s);


	if err != nil {
		panic(err)
	}




	http.ListenAndServe(":8080", handler);

}



