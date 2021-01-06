package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewHeading(n *html.Node) (Element, error) {
	h := &remark.Heading{}

	if err := set(h, n); err != nil {
		return nil, err
	}

	return h, nil
}
