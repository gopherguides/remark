package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewTR(n *html.Node) (Element, error) {
	tr := &remark.TR{}

	if err := set(tr, n); err != nil {
		return nil, err
	}

	return tr, nil
}
