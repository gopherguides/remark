package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewHead(n *html.Node) (Element, error) {
	head := &remark.Head{}

	if err := set(head, n); err != nil {
		return nil, err
	}

	return head, nil
}
