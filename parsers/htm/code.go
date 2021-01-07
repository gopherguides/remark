package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)

func NewCode(n *html.Node) (Element, error) {
	code := &remark.Code{}

	err := set(code, n)
	return code, err
}
