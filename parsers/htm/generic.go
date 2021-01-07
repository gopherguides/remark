package htm

import (
	"strings"

	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func NewGeneric(n *html.Node) (remark.Tag, error) {
	t := &remark.Generic{
		TagName: strings.ToLower(n.Data),
	}
	err := set(t, n)
	return t, err
}
