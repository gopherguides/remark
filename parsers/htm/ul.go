package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewUnorderedList(n *html.Node) (Element, error) {
	ul := &remark.UnorderedList{}

	if err := set(ul, n); err != nil {
		return nil, err
	}

	return ul, nil
}
