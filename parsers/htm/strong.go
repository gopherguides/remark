package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewStrong(n *html.Node) (Element, error) {
	strong := &remark.Strong{}

	if err := set(strong, n); err != nil {
		return nil, err
	}

	return strong, nil
}
