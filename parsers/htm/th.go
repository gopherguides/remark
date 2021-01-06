package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewTH(n *html.Node) (Element, error) {
	th := &remark.TH{}

	if err := set(th, n); err != nil {
		return nil, err
	}

	return th, nil
}
