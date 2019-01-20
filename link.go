package link

import (
	"io"

	"golang.org/x/net/html"
)

// Link type
type Link struct {
	Href string
	Text string
}

// Parse will take an io.Reader and return a slice of links
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	links := make([]Link, 0)
	var f func(node *html.Node)
	f = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			var href string
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					href = attr.Val
					break
				}
			}
			text := node.FirstChild.Data

			links = append(links, Link{href, text})
		}

		if node.FirstChild != nil {
			f(node.FirstChild)
		}
		if node.NextSibling != nil {
			f(node.NextSibling)
		}
	}
	f(doc)

	return links, nil
}
