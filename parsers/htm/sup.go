package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewSup(n *html.Node) (Element, error) {
	sup := &remark.Sup{}

	if err := set(sup, n); err != nil {
		return nil, err
	}

	return sup, nil
}
