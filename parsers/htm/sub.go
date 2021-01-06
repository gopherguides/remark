package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewSub(n *html.Node) (Element, error) {
	sub := &remark.Sub{}

	if err := set(sub, n); err != nil {
		return nil, err
	}

	return sub, nil
}
