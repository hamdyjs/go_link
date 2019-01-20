package link

import (
	"io"

	"golang.org/x/net/html"
)

// Link is a type that describes parsed links
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
	parseNode(doc, &links)

	return links, nil
}

func parseNode(node *html.Node, links *[]Link) {
	if node.Type == html.ElementNode && node.Data == "a" {
		var href string
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				href = attr.Val
				break
			}
		}
		var text string
		child := node.FirstChild
		for child != nil {
			if child.Type == html.TextNode {
				text += child.Data
			} else if child.Type == html.ElementNode && child.FirstChild != nil {
				text += child.FirstChild.Data
			}
			child = child.NextSibling
		}

		*links = append(*links, Link{href, text})
	}

	if node.FirstChild != nil {
		parseNode(node.FirstChild, links)
	}
	if node.NextSibling != nil {
		parseNode(node.NextSibling, links)
	}
}
