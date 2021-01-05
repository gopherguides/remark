package htm

import (
	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func NewHead(n *html.Node) (*remark.Head, error) {
	h := &remark.Head{}

	if err := set(h, n); err != nil {
		return nil, err
	}

	return h, nil
}
