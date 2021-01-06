package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewTableBody(n *html.Node) (Element, error) {
	tbody := &remark.TableBody{}

	if err := set(tbody, n); err != nil {
		return nil, err
	}

	return tbody, nil
}
