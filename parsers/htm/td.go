package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewTD(n *html.Node) (Element, error) {
	td := &remark.TD{}

	if err := set(td, n); err != nil {
		return nil, err
	}

	return td, nil
}
