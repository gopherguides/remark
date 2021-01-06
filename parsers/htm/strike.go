package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewStrike(n *html.Node) (Element, error) {
	strike := &remark.Strike{}

	if err := set(strike, n); err != nil {
		return nil, err
	}

	return strike, nil
}
