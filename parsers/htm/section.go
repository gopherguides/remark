package htm

import (
	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func NewSection(n *html.Node) (*remark.Section, error) {
	s := &remark.Section{}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		el, err := NewElement(c)
		if err != nil {
			return nil, err
		}
		s.Elements = append(s.Elements, el)
	}

	return s, nil
}
