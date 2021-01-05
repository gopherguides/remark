package htm

import (
	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func NewParagraph(n *html.Node) (*remark.Paragraph, error) {
	d := &remark.Paragraph{}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		el, err := NewElement(c)
		if err != nil {
			return nil, err
		}
		d.Elements = append(d.Elements, el)
	}

	return d, nil
}
