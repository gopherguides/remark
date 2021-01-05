package htm

import (
	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func NewImage(n *html.Node) (*remark.Image, error) {
	i := &remark.Image{}

	if err := set(i, n); err != nil {
		return nil, err
	}

	return i, nil
}
