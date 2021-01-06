package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewCode(n *html.Node) (Element, error) {
	code := &remark.Code{}

	if err := set(code, n); err != nil {
		return nil, err
	}

	return code, nil
}
