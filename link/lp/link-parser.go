package linkParser

import (
	"io"

	"golang.org/x/net/html"
)


type Link struct {
	Href string
	Text string
}


func Parse(r io.Reader) ([]Link, error) {


	var ret []Link
	doc, err := html.Parse(r);
	if err != nil {
		return nil, err
	}

	nodes := linkNodes(doc)
	for _, node := range nodes {
		ret = append(ret, bulidLink(node))
	}

	return ret, nil
}

func bulidLink(n *html.Node) Link {

	var ret Link;
	for _ , attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
		}
	}
	ret.Text = text(n);
	return ret;
}


func text (n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string;
	for c := n.FirstChild; c != nil ; c = c.NextSibling {
		ret += text(c) + " ";
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

