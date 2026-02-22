package main

import (
	"fmt"
	linkParser "link/lp"
	"strings"
)



var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>
`


func main() {
	r := strings.NewReader(exampleHtml)

	links, err := linkParser.Parse(r);

	if err != nil {
		panic(err)
	}

	for _ , link := range links {
		fmt.Println(link)
	}

	
}

