package main

import (
	"flag"
	"fmt"
	"linktree/crawl"
	"log"
)

func main() {
	urlFlag := flag.String("url", "https://www.netomi.com/", "Starting URL")
	flag.Parse()

	links, err := crawl.Crawl(*urlFlag)
	if err != nil {
		log.Fatal(err)
	}

	for _, l := range links {
		fmt.Println(l)
	}
}