package htm

import (
	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func NewDiv(n *html.Node) (*remark.Div, error) {
	d := &remark.Div{}

	if err := set(d, n); err != nil {
		return nil, err
	}

	return d, nil
}
