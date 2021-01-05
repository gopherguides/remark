package htm

import (
	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func NewDiv(n *html.Node) (*remark.Div, error) {
	d := &remark.Div{}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		el, err := NewElement(c)
		if err != nil {
			return nil, err
		}
		d.Elements = append(d.Elements, el)
	}

	return d, nil
}
