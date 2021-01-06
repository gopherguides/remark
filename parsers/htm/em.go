package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewEm(n *html.Node) (Element, error) {
	em := &remark.Em{}

	if err := set(em, n); err != nil {
		return nil, err
	}

	return em, nil
}
