package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewLink(n *html.Node) (Element, error) {
	a := &remark.Link{}

	if err := set(a, n); err != nil {
		return nil, err
	}

	return a, nil
}
