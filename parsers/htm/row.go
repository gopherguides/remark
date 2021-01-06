package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewRow(n *html.Node) (Element, error) {
	row := &remark.Row{}

	if err := set(row, n); err != nil {
		return nil, err
	}

	return row, nil
}
