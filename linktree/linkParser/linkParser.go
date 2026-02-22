package linkParser

import (
	"io"
	"net/url"
	"strings"
	"golang.org/x/net/html"
)

func Parse(r io.Reader, baseURL string) ([]string, error) {
	var links []string
	doc, err := html.Parse(r);
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	for _, node := range nodes {
		links = append(links, bulidLink(node))
	}
	ret := filterLinks(links, baseURL);
	return ret, nil
}




func filterLinks(links []string, baseURL string) []string {


	var ret []string
	seen := make(map[string]struct{})

	base, err := url.Parse(baseURL)
	if err != nil || base.Scheme == "" || base.Host == "" {
		return ret
	}

	for _, href := range links {
		href = strings.TrimSpace(href)
		if href == "" || strings.HasPrefix(href, "#") {
			continue
		}

		u, err := url.Parse(href)
		if err != nil {
			continue
		}

		u = base.ResolveReference(u)


		if u.Hostname() != base.Hostname() {
			continue
		}

		u.Fragment = ""
		final := u.String()
		if _, ok := seen[final]; !ok {
			seen[final] = struct{}{}
			ret = append(ret, final)
		}
	}

	return ret
}

func bulidLink(n *html.Node) string {
	var ret string;
	for _ , attr := range n.Attr {
		if attr.Key == "href" {
			ret = attr.Val
		}
	}
	return ret;
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode &&  n.Data == "a" {
		return []*html.Node{n}
	}
	var ret  []*html.Node
	for c := n.FirstChild; c != nil ; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret;
}

