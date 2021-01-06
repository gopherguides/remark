package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewDiv(n *html.Node) (Element, error) {
	div := &remark.Div{}

	if err := set(div, n); err != nil {
		return nil, err
	}

	return div, nil
}
