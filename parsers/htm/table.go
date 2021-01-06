package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewTable(n *html.Node) (Element, error) {
	table := &remark.Table{}

	if err := set(table, n); err != nil {
		return nil, err
	}

	return table, nil
}
