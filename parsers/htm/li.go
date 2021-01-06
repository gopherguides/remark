package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewListItem(n *html.Node) (Element, error) {
	li := &remark.ListItem{}

	if err := set(li, n); err != nil {
		return nil, err
	}

	return li, nil
}
