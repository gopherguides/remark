package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)

func NewLink(n *html.Node) (Element, error) {
	a := &remark.Link{}

	err := set(a, n)
	return a, err
}
