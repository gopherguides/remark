package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewParagraph(n *html.Node) (Element, error) {
	p := &remark.Paragraph{}

	if err := set(p, n); err != nil {
		return nil, err
	}

	return p, nil
}
