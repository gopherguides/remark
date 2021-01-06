package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewHR(n *html.Node) (Element, error) {
	hr := &remark.HR{}

	if err := set(hr, n); err != nil {
		return nil, err
	}

	return hr, nil
}
