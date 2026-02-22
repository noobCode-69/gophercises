package crawl

import (
	"linktree/linkParser"
	"net/http"
)
func Crawl(start string) ([]string, error) {
	queue := []string{start}
	visited := map[string]bool{}
	var results []string

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		if visited[u] {
			continue
		}
		visited[u] = true

		resp, err := http.Get(u)
		if err != nil {
			continue
		}

		links, err := linkParser.Parse(resp.Body, start)
		resp.Body.Close()
		if err != nil {
			continue
		}

		results = append(results, u)

		for _, l := range links {
			if !visited[l] {
				queue = append(queue, l)
			}
		}
	}

	return results, nil
}