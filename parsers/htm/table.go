package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)

func NewTable(n *html.Node) (Element, error) {
	table := &remark.Table{}

	err := set(table, n)
	return table, err
}
