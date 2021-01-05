package htm

import (
	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func NewParagraph(n *html.Node) (*remark.Paragraph, error) {
	d := &remark.Paragraph{}

	if err := set(d, n); err != nil {
		return nil, err
	}

	return d, nil
}
