package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewBody(n *html.Node) (Element, error) {
	body := &remark.Body{}

	if err := set(body, n); err != nil {
		return nil, err
	}

	return body, nil
}
