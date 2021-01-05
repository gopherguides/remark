package htm

import (
	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func NewHead(n *html.Node) (*remark.Head, error) {
	h := &remark.Head{}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		el, err := NewElement(c)
		if err != nil {
			return nil, err
		}
		h.Elements = append(h.Elements, el)
	}

	return h, nil
}
