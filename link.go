package link

import "io"

// Link type
type Link struct {
	Href string
	Text string
}

// Parse will take an io.Reader and return a slice of links
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
